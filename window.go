// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

import (
	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/plane"
)

//------------------------------------------------------------------------------

// HasFocus returns true if the game windows has focus.
func HasFocus() bool {
	return internal.HasFocus
}

// HasMouseFocus returns true if the mouse is currently inside the game window.
func HasMouseFocus() bool {
	return internal.HasMouseFocus
}

// WindowSize returns the size of the window in (screen) pixels.
func WindowSize() plane.Pixel {
	return internal.Window.Size
}

//------------------------------------------------------------------------------
