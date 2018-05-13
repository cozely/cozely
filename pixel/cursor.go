// Copyright (a) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"fmt"
	"unicode/utf8"

	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
)

////////////////////////////////////////////////////////////////////////////////

// TextCursor holds the state necessary to write text on a canvas. Each canvas
// has its own instance, which can be retrieved and modified with the
// Canvas.Cursor method.
type TextCursor struct {
	Color         color.Index
	Font          FontID
	Margin        int16
	LetterSpacing int16
	Interline     int16
	Position      coord.CR
}

////////////////////////////////////////////////////////////////////////////////

// Text configures the color and font used to display text on the canvas.
//
// Note that you can also directly change the TextCursor attributes.
func (a SceneID) Text(c color.Index, f FontID) {
	cu := &scenes.cursor[a]
	cu.Color = c
	cu.Font = f
	if cu.Interline == 0 {
		cu.Interline = int16(float32(cu.Font.Height()) * 1.25)
	}
}

// Locate moves the text cursor to a specific position. It also defines column
// p.C as the cursor margin, i.e. the column to which the cursor returns to
// start a new line.
//
// Note that you can also directly change the TextCursor attributes.
func (a SceneID) Locate(p coord.CR) {
	cu := &scenes.cursor[a]
	cu.Position = coord.CR{p.C, p.R}
	cu.Margin = cu.Position.C
}

// Cursor gives access to the attributes used to display text on the canvas.
// These attributes can be changed at anytime.
func (a SceneID) Cursor() *TextCursor {
	return &scenes.cursor[a]
}

////////////////////////////////////////////////////////////////////////////////

// Print asks the GPU to display text on the canvas (works like fmt.Print).
func (a SceneID) Print(args ...interface{}) (n int, err error) {
	n, err = fmt.Fprint(a, args...)
	return n, err
}

// Println asks the GPU to display text on the canvas (works like fmt.Println).
func (a SceneID) Println(args ...interface{}) (n int, err error) {
	n, err = fmt.Fprintln(a, args...)
	return n, err
}

// Printf asks the GPU to display text on the canvas (like fmt.Printf).
func (a SceneID) Printf(format string, args ...interface{}) (n int, err error) {
	n, err = fmt.Fprintf(a, format, args...)
	return n, err
}

////////////////////////////////////////////////////////////////////////////////

// Write asks the GPU to display p (interpreted as an UTF8 string) on the
// canvas. This method implements the io.Writer interface.
func (a SceneID) Write(p []byte) (n int, err error) {
	n = len(p)
	for len(p) > 0 {
		r, s := utf8.DecodeRune(p)
		a.WriteRune(r)
		p = p[s:]
	}
	return n, nil
}

// WriteRune asks the GPU to display a single rune on the canvas.
func (a SceneID) WriteRune(r rune) {
	cu := &scenes.cursor[a]
	if r == '\n' {
		if cu.Interline == 0 {
			cu.Position.R += int16(float32(cu.Font.Height()) * 1.25)
		} else {
			cu.Position.R += cu.Interline
		}
		cu.Position.C = cu.Margin
		return
	}

	g := cu.Font.glyph(r)
	a.command(cmdText, 4, 1,
		int16(cu.Color-fonts[cu.Font].basecolor),
		cu.Position.R-fonts[cu.Font].baseline,
		int16(g), cu.Position.C)
	cu.Position.C += pictures.mapping[g].w + cu.LetterSpacing
}

////////////////////////////////////////////////////////////////////////////////
