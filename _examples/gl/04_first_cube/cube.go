// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/colour"
	"github.com/drakmaniso/glam/space"
)

//------------------------------------------------------------------------------

var (
	purple = colour.RGB{R: 0.2, G: 0, B: 0.6}
	orange = colour.RGB{R: 0.8, G: 0.3, B: 0}
	green  = colour.RGB{R: 0, G: 0.3, B: 0.1}
)

//------------------------------------------------------------------------------

func cube() mesh {
	return mesh{
		// Front Face
		{space.Coord{-0.5, -0.5, +0.5}, purple},
		{space.Coord{+0.5, +0.5, +0.5}, purple},
		{space.Coord{-0.5, +0.5, +0.5}, purple},
		{space.Coord{-0.5, -0.5, +0.5}, purple},
		{space.Coord{+0.5, -0.5, +0.5}, purple},
		{space.Coord{+0.5, +0.5, +0.5}, purple},
		// Back Face
		{space.Coord{-0.5, -0.5, -0.5}, purple},
		{space.Coord{-0.5, +0.5, -0.5}, purple},
		{space.Coord{+0.5, +0.5, -0.5}, purple},
		{space.Coord{-0.5, -0.5, -0.5}, purple},
		{space.Coord{+0.5, +0.5, -0.5}, purple},
		{space.Coord{+0.5, -0.5, -0.5}, purple},
		// Right Face
		{space.Coord{+0.5, -0.5, +0.5}, green},
		{space.Coord{+0.5, +0.5, -0.5}, green},
		{space.Coord{+0.5, +0.5, +0.5}, green},
		{space.Coord{+0.5, -0.5, +0.5}, green},
		{space.Coord{+0.5, -0.5, -0.5}, green},
		{space.Coord{+0.5, +0.5, -0.5}, green},
		// Left Face
		{space.Coord{-0.5, -0.5, +0.5}, green},
		{space.Coord{-0.5, +0.5, +0.5}, green},
		{space.Coord{-0.5, +0.5, -0.5}, green},
		{space.Coord{-0.5, -0.5, +0.5}, green},
		{space.Coord{-0.5, +0.5, -0.5}, green},
		{space.Coord{-0.5, -0.5, -0.5}, green},
		// Bottom Face
		{space.Coord{-0.5, -0.5, +0.5}, orange},
		{space.Coord{-0.5, -0.5, -0.5}, orange},
		{space.Coord{+0.5, -0.5, +0.5}, orange},
		{space.Coord{-0.5, -0.5, -0.5}, orange},
		{space.Coord{+0.5, -0.5, -0.5}, orange},
		{space.Coord{+0.5, -0.5, +0.5}, orange},
		// Top Face
		{space.Coord{-0.5, +0.5, +0.5}, orange},
		{space.Coord{+0.5, +0.5, +0.5}, orange},
		{space.Coord{-0.5, +0.5, -0.5}, orange},
		{space.Coord{-0.5, +0.5, -0.5}, orange},
		{space.Coord{+0.5, +0.5, +0.5}, orange},
		{space.Coord{+0.5, +0.5, -0.5}, orange},
	}
}

//------------------------------------------------------------------------------
