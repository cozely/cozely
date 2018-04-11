// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"image"
	_ "image/png"
	"os"

	"github.com/drakmaniso/cozely"
	"github.com/drakmaniso/cozely/colour"
	"github.com/drakmaniso/cozely/plane"
	"github.com/drakmaniso/cozely/space"
	"github.com/drakmaniso/cozely/x/gl"
)

//------------------------------------------------------------------------------

func main() {
	err := cozely.Run(loop{})
	if err != nil {
		cozely.ShowError(err)
		return
	}
}

//------------------------------------------------------------------------------

// OpenGL objects
var (
	pipeline    *gl.Pipeline
	perFrameUBO gl.UniformBuffer
	sampler     gl.Sampler
	diffuse     gl.Texture2D
)

// Uniform buffer
var perObject struct {
	screenFromObject space.Matrix
}

// Vertex buffer
type mesh []struct {
	position space.Coord `layout:"0"`
	uv       plane.Coord `layout:"1"`
}

// Transformation matrices
var (
	screenFromView  space.Matrix // projection matrix
	viewFromWorld   space.Matrix // view matrix
	worldFromObject space.Matrix // model matrix
)

// Cube state
var (
	position   space.Coord
	yaw, pitch float32
)

//------------------------------------------------------------------------------

type loop struct {
	cozely.EmptyLoop
}

//------------------------------------------------------------------------------

func (loop) Enter() error {
	// Create and configure the pipeline
	pipeline = gl.NewPipeline(
		gl.Shader(cozely.Path()+"shader.vert"),
		gl.Shader(cozely.Path()+"shader.frag"),
		gl.VertexFormat(0, mesh{}),
		gl.Topology(gl.Triangles),
		gl.CullFace(false, true),
		gl.DepthTest(true),
	)

	// Create the uniform buffer
	perFrameUBO = gl.NewUniformBuffer(&perObject, gl.DynamicStorage)

	// Create and fill the vertex buffer
	vbo := gl.NewVertexBuffer(cube(), gl.StaticStorage)

	// Create and bind the sampler
	sampler = gl.NewSampler(
		gl.Minification(gl.LinearMipmapLinear),
		gl.Anisotropy(16.0),
	)

	// Create and load the textures
	diffuse = gl.NewTexture2D(8, gl.SRGBA8, 512, 512)
	r, err := os.Open(cozely.Path() + "../../shared/testpattern.png")
	if err != nil {
		return cozely.Error("opening texture", err)
	}
	defer r.Close()
	img, _, err := image.Decode(r)
	if err != nil {
		return cozely.Error("decoding texture", err)
	}
	diffuse.SubImage(0, 0, 0, img)
	diffuse.GenerateMipmap()

	// Initialize worldFromObject and viewFromWorld matrices
	position = space.Coord{0, 0, 0}
	yaw = -0.6
	pitch = 0.3
	computeWorldFromObject()
	computeViewFromWorld()

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	vbo.Bind(0, 0)
	pipeline.Unbind()

	return cozely.Error("gl", gl.Err())
}

//------------------------------------------------------------------------------

func computeWorldFromObject() {
	r := space.EulerZXY(pitch, yaw, 0)
	worldFromObject = space.Translation(position).Times(r)
}

func computeViewFromWorld() {
	viewFromWorld = space.LookAt(
		space.Coord{0, 0, 3},
		space.Coord{0, 0, 0},
		space.Coord{0, 1, 0},
	)
}

//------------------------------------------------------------------------------

func (loop) Update() error {
	return nil
}

//------------------------------------------------------------------------------

func (loop) Draw() error {
	pipeline.Bind()
	gl.ClearDepthBuffer(1.0)
	gl.ClearColorBuffer(colour.LRGBA{0.9, 0.9, 0.9, 1.0})
	gl.Enable(gl.FramebufferSRGB)

	perObject.screenFromObject =
		screenFromView.
			Times(viewFromWorld).
			Times(worldFromObject)
	perFrameUBO.SubData(&perObject, 0)
	perFrameUBO.Bind(0)

	diffuse.Bind(0)
	sampler.Bind(0)
	gl.Draw(0, 6*2*3)

	pipeline.Unbind()

	return gl.Err()
}

//------------------------------------------------------------------------------
