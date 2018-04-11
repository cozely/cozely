// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package poly

import (
	"math"

	"github.com/drakmaniso/cozely/colour"
	"github.com/drakmaniso/cozely/x/math32"
)

//------------------------------------------------------------------------------

func TemperatureColor(temperature float64) colour.LRGB {
	// Ported by Renaud Bédard (@renaudbedard), from original code
	// by Tanner Helland:
	// http://www.tannerhelland.com/4435/convert-temperature-rgb-algorithm-code/
	// https://www.shadertoy.com/view/lsSXW1
	// licensed and released under Creative Commons 3.0 Attribution
	// https://creativecommons.org/licenses/by/3.0/

	var h colour.LRGB

	temperature = clamp(temperature, 1000.0, 40000.0) / 100.0

	if temperature <= 66.0 {
		h.R = 1.0
		h.G = float32(saturate(
			0.39008157876901960784*math.Log(temperature) - 0.63184144378862745098,
		))
	} else {
		t := temperature - 60.0
		h.R = float32(saturate(1.29293618606274509804 * math.Pow(t, -0.1332047592)))
		h.G = float32(saturate(1.12989086089529411765 * math.Pow(t, -0.0755148492)))
	}

	switch {
	case temperature >= 66.0:
		h.B = 1.0
	case temperature <= 19.0:
		h.B = 0.0
	default:
		h.B = float32(saturate(
			0.54320678911019607843*math.Log(temperature-10.0) - 1.19625408914,
		))
	}

	return h
}

//------------------------------------------------------------------------------

func DirectionalLightSpectralIlluminance(illuminance, temperature float64) colour.LRGB {
	h := TemperatureColor(temperature)
	s := float32(illuminance) * (1.0 / math32.Pi)
	h.R *= s
	h.G *= s
	h.B *= s

	return h
}

//------------------------------------------------------------------------------
