// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

import (
	"fmt"
	"unsafe"
)

//------------------------------------------------------------------------------

/*
#include <stdlib.h>
#include "sdl.h"

static inline void SwapWindow(SDL_Window* w) {
	SDL_GL_SwapWindow(w);
}
*/
import "C"

//------------------------------------------------------------------------------

// OpenWindow creates the game window and its associated OpenGL context.
func OpenWindow(
	title string,
	width, height int16,
	display int,
	fullscreen bool,
	fullscreenMode string,
	vsync bool,
	debug bool,
) error {
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MAJOR_VERSION, 4)
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MINOR_VERSION, 6)
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_PROFILE_MASK, C.SDL_GL_CONTEXT_PROFILE_CORE)
	C.SDL_GL_SetAttribute(C.SDL_GL_DOUBLEBUFFER, 1)
	if Window.Multisample > 0 {
		C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLEBUFFERS, 1)
		C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLESAMPLES, C.int(Window.Multisample))
	}

	if debug {
		C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_FLAGS, C.SDL_GL_CONTEXT_DEBUG_FLAG)
	}

	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))

	Window.Width, Window.Height = width, height

	var fs uint32
	if fullscreen {
		if fullscreenMode == "Desktop" {
			fs = C.SDL_WINDOW_FULLSCREEN_DESKTOP
		} else {
			fs = C.SDL_WINDOW_FULLSCREEN
		}
	}
	fl := C.SDL_WINDOW_OPENGL | C.SDL_WINDOW_RESIZABLE | C.Uint32(fs)

	Window.window = C.SDL_CreateWindow(
		t,
		C.int(C.SDL_WINDOWPOS_CENTERED_MASK|display),
		C.int(C.SDL_WINDOWPOS_CENTERED_MASK|display),
		C.int(Window.Width),
		C.int(Window.Height),
		fl,
	)
	if Window.window == nil {
		err := GetSDLError()
		return fmt.Errorf("could not open window: %s", err)
	}

	ctx := C.SDL_GL_CreateContext(Window.window)
	if ctx == nil {
		err := GetSDLError()
		return fmt.Errorf("could not create OpenGL context: %s", err)
	}
	Window.context = ctx

	var si C.int
	if vsync {
		si = 1
	}
	C.SDL_GL_SetSwapInterval(si)

	logOpenGLInfos()

	return nil
}

// logOpenGLInfos displays information about the OpenGL context
func logOpenGLInfos() {
	s := "OpenGL: "
	maj, err1 := sdlGLAttribute(C.SDL_GL_CONTEXT_MAJOR_VERSION)
	min, err2 := sdlGLAttribute(C.SDL_GL_CONTEXT_MINOR_VERSION)
	if err1 == nil && err2 == nil {
		s += fmt.Sprintf("%d.%d", maj, min)
	}

	db, err1 := sdlGLAttribute(C.SDL_GL_DOUBLEBUFFER)
	if err1 == nil {
		if db != 0 {
			s += ", double buffer"
		} else {
			s += ", NO double buffer"
		}
	}

	av, err1 := sdlGLAttribute(C.SDL_GL_ACCELERATED_VISUAL)
	if err1 == nil {
		if av != 0 {
			s += ", accelerated"
		} else {
			s += ", NOT accelerated"
		}
	}

	sw := C.SDL_GL_GetSwapInterval()
	if sw > 0 {
		if sw != 0 {
			s += ", vsync"
		} else {
			s += ", NO vsync"
		}
	}
	Debug.Println(s)
}

func sdlGLAttribute(attr C.SDL_GLattr) (int, error) {
	var v C.int
	errcode := C.SDL_GL_GetAttribute(attr, &v)
	if errcode < 0 {
		return 0, GetSDLError()
	}
	return int(v), nil
}

//------------------------------------------------------------------------------

// SwapWindow swaps the double-buffer.
func SwapWindow() {
	C.SwapWindow(Window.window)
}

//------------------------------------------------------------------------------

func SetFullscreen(f bool) {
	var fs C.Uint32
	if f {
		if Config.FullscreenMode == "Desktop" {
			fs = C.SDL_WINDOW_FULLSCREEN_DESKTOP
		} else {
			fs = C.SDL_WINDOW_FULLSCREEN
		}
	}
	C.SDL_SetWindowFullscreen(Window.window, fs)
}

func GetFullscreen() bool {
	fs := C.SDL_GetWindowFlags(Window.window)
	fs &= (C.SDL_WINDOW_FULLSCREEN_DESKTOP | C.SDL_WINDOW_FULLSCREEN)
	return fs != 0
}

func ToggleFullscreen() {
	fs := !GetFullscreen()
	SetFullscreen(fs)
}

//------------------------------------------------------------------------------

func SetWindowTitle(title string) {
	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))
	C.SDL_SetWindowTitle(Window.window, t)
}

//------------------------------------------------------------------------------

// destroyWindow closes the game window and delete the OpenGL context
func destroyWindow() {
	C.SDL_GL_DeleteContext(Window.context)
	C.SDL_DestroyWindow(Window.window)
}

//------------------------------------------------------------------------------
