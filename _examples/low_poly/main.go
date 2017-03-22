// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"os"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/math"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/mtx"
	"github.com/drakmaniso/glam/pbr"
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam/poly"
	"github.com/drakmaniso/glam/space"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

func main() {
	glam.Setup()

	err := setup()
	if err != nil {
		glam.ErrorDialog(err)
		return
	}

	window.Handle = handler{}
	mouse.Handle = handler{}
	key.Handle = handler{}

	glam.TimeStep = 1.0 / 50.0

	// Run the Game Loop
	glam.Loop = looper{}
	err = glam.Run()
	if err != nil {
		glam.ErrorDialog(err)
	}
}

//------------------------------------------------------------------------------

// OpenGL objects
var (
	frameUBO gfx.UniformBuffer
)

// Uniform buffer
var frame struct {
	ProjectionView space.Matrix
	Model          space.Matrix
	CameraPosition space.Coord
	CameraExposure float32
	SunIlluminance color.RGB
	_              byte
}

var meshes poly.Meshes

// Camera
var (
	object struct {
		position         space.Coord
		yaw, pitch, roll float32
		scale            float32
	}
	camera struct {
		position   space.Coord
		velocity   space.Coord
		yaw, pitch float32
	}
	cameraNext struct {
		position   space.Coord
		yaw, pitch float32
	}
)

var (
	projection space.Matrix
	view       space.Matrix
)

//------------------------------------------------------------------------------

func setup() error {
	// Create and configure the pipeline
	f, err := os.Open(glam.Path() + "shader.frag")
	if err != nil {
		return err
	}
	poly.SetupPipeline(
		gfx.FragmentShader(f),
	)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	frameUBO = gfx.NewUniformBuffer(&frame, gfx.DynamicStorage)

	//
	meshes = poly.Meshes{}
	meshes.AddObj(glam.Path() + "../shared/suzanne.obj")
	// meshes.AddObj("E:/objtestfiles/stygirl.obj")
	poly.SetupMeshBuffers(meshes)

	// Initialize view matrix
	object.position = space.Coord{0, 0, -4.0}
	object.scale = 1.0
	camera.position = space.Coord{0, 0, 0.0}
	cameraNext.position = camera.position
	updateModel()
	updateView(camera.position, camera.yaw, camera.pitch)
	frame.SunIlluminance = pbr.DirectionalLightSpectralIlluminance(116400.0, 5400.0)
	frame.CameraExposure = float32(pbr.Exposure(16.0, 1.0/125.0, 100.0))
	print(frame.CameraExposure, "\n")
	print(frame.SunIlluminance.R, "  ", frame.SunIlluminance.G, "  ", frame.SunIlluminance.B, "\n")

	// MTX
	mtx.Color(color.RGB{0.0, 0.05, 0.1}, color.RGB{0.7, 0.6, 0.45})
	mtx.Opaque(false)
	mtx.ShowFrameTime(true, -1, 0, false)

	// Bind the vertex buffer to the pipeline
	poly.BindPipeline()

	// pipeline.Bind()
	// vbo.Bind(0, 0)
	// pipeline.Unbind()

	return gfx.Err()
}

//------------------------------------------------------------------------------

type looper struct{}

func (l looper) Update(_, dt float64) {
	v := camera.velocity.Times(float32(dt))
	camera.position = cameraNext.position
	camera.yaw = cameraNext.yaw
	camera.pitch = cameraNext.pitch

	cameraNext.position.X += v.X*math.Cos(cameraNext.yaw) - v.Z*math.Sin(cameraNext.yaw)
	cameraNext.position.Z += v.X*math.Sin(cameraNext.yaw) + v.Z*math.Cos(cameraNext.yaw)
	mtx.Print(1, 0, "cam: %6.2f,%6.2f", cameraNext.position.X, cameraNext.position.Z)

	if firstPerson {
		mx, my := mouse.Delta().Cartesian()

		const d = 0.5
		smoothedMouse.X += (mx - smoothedMouse.X) * d
		smoothedMouse.Y += (my - smoothedMouse.Y) * d

		sx, sy := window.Size().Cartesian()
		cameraNext.yaw += 2 * smoothedMouse.X / sx
		cameraNext.pitch += 2 * smoothedMouse.Y / sy
		switch {
		case cameraNext.pitch < -math.Pi/2:
			cameraNext.pitch = -math.Pi / 2
		case cameraNext.pitch > +math.Pi/2:
			cameraNext.pitch = +math.Pi / 2
		}
	}
}

var smoothedMouse plane.Coord

//------------------------------------------------------------------------------

func (l looper) Draw(interpolation float64) {
	poly.BindPipeline()
	gfx.ClearColorBuffer(color.RGBA{0.8, 0.8, 0.8, 1.0})

	pos := camera.position.Times(1.0 - float32(interpolation))
	pos = pos.Plus(cameraNext.position.Times(float32(interpolation)))
	y := (1.0-float32(interpolation))*camera.yaw + float32(interpolation)*cameraNext.yaw
	p := (1.0-float32(interpolation))*camera.pitch + float32(interpolation)*cameraNext.pitch
	updateView(pos, y, p)
	frame.ProjectionView = projection.Times(view)
	frame.CameraPosition = pos
	frameUBO.SubData(&frame, 0)
	frameUBO.Bind(0)

	poly.BindMeshBuffers()

	// poly.Draw()
	gfx.Draw(gfx.Triangles, 0, int32(len(meshes.Faces)*6))

	poly.UnbindPipeline()
}

func updateModel() {
	frame.Model = space.Translation(object.position)
	frame.Model = frame.Model.Times(space.EulerXZY(object.pitch, object.yaw, object.roll))
	frame.Model = frame.Model.Times(space.Scaling(space.Coord{object.scale, object.scale, object.scale}))
}

func updateView(pos space.Coord, yaw, pitch float32) {
	view = space.EulerZXY(pitch, yaw, 0)
	view = view.Times(space.Translation(pos.Inverse()))
}

//------------------------------------------------------------------------------