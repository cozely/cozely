// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"errors"
	"image"
	_ "image/png" // Activate PNG support
	"os"
	"path/filepath"

	"github.com/drakmaniso/glam/internal"
	"github.com/drakmaniso/glam/x/atlas"
)

//------------------------------------------------------------------------------

var (
	fntFiles []atlas.Image
	fntAtlas *atlas.Atlas
)

//------------------------------------------------------------------------------

func (f Font) load() error {
	//TODO: support other image formats?
	n := fontPaths[f]
	path := filepath.FromSlash(internal.Path + n + ".png")
	path, err := filepath.EvalSymlinks(path)
	if err != nil {
		return internal.Error("in path while loading font", err)
	}

	fl, err := os.Open(path)
	if err != nil {
		return internal.Error(`while opening font file "`+path+`"`, err)
	}
	defer fl.Close() //TODO: error handling

	img, _, err := image.Decode(fl)
	switch err {
	case nil:
	case image.ErrFormat:
		return nil
	default:
		return internal.Error("decoding font file", err)
	}

	p, ok := img.(*image.Paletted)
	if !ok {
		internal.Debug.Println(`ignoring image: "` + path + `" (color model not supported)`)
		return nil
	}
	gh := p.Bounds().Dy()

	h := gh - 1
	fonts[f].height = int16(h)
	gly := uint16(len(glyphMap))
	fonts[f].first = gly

	for y := 0; y < p.Bounds().Dy(); y++ {
		if p.Pix[0+y*p.Stride] != 0 {
			fonts[f].baseline = int16(y)
			break
		}
	}

	// Create images and reserve mapping for each rune

	for x, g := 1, 0; x < p.Bounds().Dx(); g++ {
		w := 0
		for x+w < p.Bounds().Dx() && p.Pix[x+w+h*p.Stride] != 0 {
			w++
		}
		m := p.SubImage(image.Rect(x, 0, x+w, h))
		mm, ok := m.(*image.Paletted)
		if !ok {
			return errors.New("unexpected subimage in Loadfont")
		}
		glyphMap = append(glyphMap, mapping{w: int16(w), h: int16(h)})
		fntFiles = append(
			fntFiles,
			fntrune{
				glyph: gly + uint16(g),
				img:   mm,
			},
		)
		x += w
		for x < p.Bounds().Dx() && p.Pix[x+h*p.Stride] == 0 {
			x++
		}
	}

	// Pack them into the atlas

	fntAtlas.Pack(fntFiles)

	return nil
}

//------------------------------------------------------------------------------

type fntrune struct {
	glyph uint16
	img   *image.Paletted
}

func (fr fntrune) Size() (width, height int16) {
	w, h := fr.img.Bounds().Dx(), fr.img.Bounds().Dy()
	return int16(w), int16(h)
}

func (fr fntrune) Put(bin int16, x, y int16) {
	glyphMap[fr.glyph].bin = bin
	glyphMap[fr.glyph].x = x
	glyphMap[fr.glyph].y = y
}

func (fr fntrune) Paint(dest interface{}) error {
	fx, fy := glyphMap[fr.glyph].x, glyphMap[fr.glyph].y
	fw, fh := glyphMap[fr.glyph].w, glyphMap[fr.glyph].h

	dm, ok := dest.(*image.Paletted)
	if !ok {
		return errors.New("unexpected dest argument to fntrune paint method")
	}
	for y := 0; y < int(fh); y++ {
		for x := 0; x < int(fw); x++ {
			w := dm.Bounds().Dx()
			ci := fr.img.Pix[x+y*fr.img.Stride]
			dm.Pix[int(fx)+x+w*(int(fy)+y)] = uint8(ci)
		}
	}

	return nil
}

//------------------------------------------------------------------------------
