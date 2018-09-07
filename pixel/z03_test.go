// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/window"
)

////////////////////////////////////////////////////////////////////////////////

var ()

type loop3 struct {
	bg, fg  pixel.Color

	tinela9, monozela10, simpela10, simpela12,
	cozela10, cozela12, chaotela12, font pixel.FontID
}

////////////////////////////////////////////////////////////////////////////////

func TestTest3(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		l := loop3{}
		l.declare()

		input.Load(bindings)
		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (a *loop3) declare() {
	pixel.SetZoom(2)
	//TODO:
	a.bg = 8
	a.fg = 1

	a.tinela9 = pixel.Font("fonts/tinela9")
	a.monozela10 = pixel.Font("fonts/monozela10")
	a.simpela10 = pixel.Font("fonts/simpela10")
	a.simpela12 = pixel.Font("fonts/simpela12")
	a.cozela10 = pixel.Font("fonts/cozela10")
	a.cozela12 = pixel.Font("fonts/cozela12")
	a.chaotela12 = pixel.Font("fonts/chaotela12")
	a.font = a.monozela10
}

func (a *loop3) Enter() {
	println(a.bg, a.fg)
	pixel.Cursor.Color = a.fg - 1
}

func (loop3) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop3) React() {
	if quit.Started(0) {
		cozely.Stop(nil)
	}
}

func (loop3) Update() {
}

func (a *loop3) Render() {
	pixel.Clear(a.bg)

	pixel.Text(a.fg, pixel.Monozela10)

	pixel.Locate(pixel.XY{2, 8})
	pixel.Println("a quick brown fox \"jumps\" over the (lazy) dog.")
	pixel.Println("A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	pixel.Println("0123456789!@#$^&*()-+=_~[]{}|\\;:'\",.<>/?%")
	pixel.Println("12+34 56-7.8 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")
	pixel.Println()

	pixel.Locate(pixel.XY{16, 100})
	pixel.Cursor.Write([]byte("Foo"))
	pixel.Cursor.Position = pixel.Cursor.Position.Plus(pixel.XY{1, 3})
	pixel.Cursor.WriteRune('B')
	pixel.Cursor.Position = pixel.Cursor.Position.Plus(pixel.XY{2, 2})
	pixel.Cursor.WriteRune('a')
	pixel.Cursor.Position = pixel.Cursor.Position.Plus(pixel.XY{3, 1})
	pixel.Cursor.WriteRune('r')
	pixel.Cursor.Position = pixel.XY{32, 132}
	pixel.Cursor.Write([]byte("Boo\n"))
	pixel.Cursor.Write([]byte("Choo"))

	pixel.Locate(pixel.XY{16, 200})
	pixel.Cursor.Font = a.tinela9
	pixel.Print("Tinela")
	pixel.Cursor.Font = a.simpela10
	pixel.Print("Simpela10")
	pixel.Cursor.Font = a.simpela12
	pixel.Print("Simpela12")
	pixel.Cursor.Font = a.cozela10
	pixel.Print("Cozela10")
	pixel.Cursor.Font = a.cozela12
	pixel.Print("Cozela12")
	pixel.Cursor.Font = a.chaotela12
	pixel.Print("Chaotela12")

	pixel.Locate(pixel.XY{pixel.Resolution().X - 200, 9})
	pixel.Cursor.Font = pixel.FontID(0)
	m := pixel.ToCanvas(window.XYof(cursor.XY(0)))
	pixel.Printf("Position x=%d, y=%d\n", m.X, m.Y)
}
