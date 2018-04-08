// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
	"unsafe"

	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

type canvas struct {
	buffer        gl.Framebuffer
	texture       gl.Texture2D
	depth         gl.Renderbuffer
	commandsICBO  gl.IndirectBuffer
	parametersTBO gl.BufferTexture
	target        plane.Pixel
	autozoom      bool
	origin        plane.Pixel // Offset when there is a border around the screen
	size          plane.Pixel
	pixel         int16
	commands      []gl.DrawIndirectCommand
	parameters    []int16
}

var canvases []canvas

// A Canvas identifies a surface that can be used to draw, print text or show
// pictures.
type Canvas uint16

//------------------------------------------------------------------------------

// NewCanvas reserves an ID for a new canvas, that will be created by glam.Run.
func NewCanvas(o ...CanvasOption) Canvas {
	if len(canvases) >= 0xFFFF {
		setErr("in NewCanvas", errors.New("too many canvases"))
	}

	cv := Canvas(len(canvases))
	canvases = append(canvases, canvas{})

	s := &canvases[cv]
	s.target = plane.Pixel{640, 360}
	s.pixel = 2
	s.commands = make([]gl.DrawIndirectCommand, 0, maxCommandCount)
	s.parameters = make([]int16, 0, maxParamCount)

	for i := range o {
		o[i](cv)
	}

	//TODO: create textures if not autoresize

	return cv
}

//------------------------------------------------------------------------------

func (cv Canvas) createBuffer() {
	s := &canvases[cv]
	s.buffer = gl.NewFramebuffer()

	s.commandsICBO = gl.NewIndirectBuffer(
		uintptr(cap(s.commands))*unsafe.Sizeof(s.commands[0]),
		gl.DynamicStorage,
	)
	s.parametersTBO = gl.NewBufferTexture(
		uintptr(cap(s.parameters))*unsafe.Sizeof(s.parameters[0]),
		gl.R16I,
		gl.DynamicStorage,
	)
}

//------------------------------------------------------------------------------

func (cv Canvas) autoresize() {
	s := &canvases[cv]

	if s.autozoom {
		// Find best fit for pixel size
		p := internal.Window.Size.Slashcw(s.target)
		if p.X < p.Y {
			s.pixel = p.X
		} else {
			s.pixel = p.Y
		}
		if s.pixel < 1 {
			s.pixel = 1
		}
	}

	// Extend the screen to cover the window
	s.size = internal.Window.Size.Slash(s.pixel)
	cv.createTextures()

	// Compute offset
	sz := s.size.Times(s.pixel)
	s.origin = internal.Window.Size.Minus(sz).Slash(2)
}

//------------------------------------------------------------------------------

func (cv Canvas) createTextures() {
	s := &canvases[cv]

	s.texture.Delete()
	s.texture = gl.NewTexture2D(1, gl.R8UI, int32(s.size.X), int32(s.size.Y))
	s.buffer.Texture(gl.ColorAttachment0, s.texture, 0)

	s.depth.Delete()
	s.depth = gl.NewRenderbuffer(gl.Depth32F, int32(s.size.X), int32(s.size.Y))
	s.buffer.Renderbuffer(gl.DepthAttachment, s.depth)

	s.buffer.DrawBuffer(gl.ColorAttachment0)
	s.buffer.ReadBuffer(gl.NoAttachment)

	st := s.buffer.CheckStatus(gl.DrawReadFramebuffer)
	if st != gl.FramebufferComplete {
		setErr("while creating screen textures", errors.New(st.String()))
	}
}

//------------------------------------------------------------------------------

// Paint executes all pending commands on the canvas. It is automatically called
// by Display; the only reason to call it manually is to be able to read from it
// before display.
func (cv Canvas) Paint() {
	s := &canvases[cv]

	if len(s.commands) == 0 {
		return
	}

	internal.PaletteUpload()

	screenUniforms.PixelSize.X = 1.0 / float32(s.size.X)
	screenUniforms.PixelSize.Y = 1.0 / float32(s.size.Y)
	screenUBO.SubData(&screenUniforms, 0)

	s.buffer.Bind(gl.DrawFramebuffer)
	gl.Viewport(0, 0, int32(s.size.X), int32(s.size.Y))
	pipeline.Bind()
	gl.Disable(gl.Blend)

	screenUBO.Bind(layoutScreen)
	s.commandsICBO.Bind()
	s.parametersTBO.Bind(layoutParameters)
	pictureMapTBO.Bind(layoutPictureMap)
	glyphMapTBO.Bind(layoutGlyphMap)
	picturesTA.Bind(layoutPictures)
	glyphsTA.Bind(layoutGlyphs)

	s.commandsICBO.SubData(s.commands, 0)
	s.parametersTBO.SubData(s.parameters, 0)
	gl.DrawIndirect(0, int32(len(s.commands)))
	s.commands = s.commands[:0]
	s.parameters = s.parameters[:0]
}

//------------------------------------------------------------------------------

// Display first execute all pending commands on the canvas (if any), then
// displays it on the game window.
func (cv Canvas) Display() {
	cv.Paint()

	s := &canvases[cv]

	sz := s.size.Times(s.pixel)

	blitUniforms.ScreenSize.X = float32(s.size.X)
	blitUniforms.ScreenSize.Y = float32(s.size.Y)
	blitUBO.SubData(&blitUniforms, 0)

	blitPipeline.Bind()
	gl.DefaultFramebuffer.Bind(gl.DrawFramebuffer)
	gl.Enable(gl.FramebufferSRGB)
	gl.Disable(gl.Blend)
	gl.Viewport(int32(s.origin.X), int32(s.origin.Y),
		int32(s.origin.X+sz.X), int32(s.origin.Y+sz.Y))
	blitUBO.Bind(0)
	s.texture.Bind(0)
	gl.Draw(0, 4)
}

//------------------------------------------------------------------------------

// Clear sets both the color and peth of all pixels on the canvas. Only the
// color is specified, the depth being initialized to the minimum value.
func (cv Canvas) Clear(color palette.Index) {
	s := &canvases[cv]
	pipeline.Bind() //TODO: find another way to enable depthWrite
	s.buffer.ClearColorUint(uint32(color), 0, 0, 0)
	s.buffer.ClearDepth(-1.0)
}

//------------------------------------------------------------------------------

// Size returns the current dimension of the canvas (in canvas pixels).
func (cv Canvas) Size() plane.Pixel {
	return canvases[cv].size
}

// PixelSize returns the size of one canvas pixel, in window pixels.
func (cv Canvas) PixelSize() int16 {
	return canvases[cv].pixel
}

//------------------------------------------------------------------------------

// Mouse returns the mouse position on the canvas.
func (cv Canvas) Mouse() plane.Pixel {
	mx, my := mouse.Position()
	m := plane.Pixel{int16(mx), int16(my)}
	return m.Minus(canvases[cv].origin).Slash(canvases[cv].pixel)
}

//------------------------------------------------------------------------------
