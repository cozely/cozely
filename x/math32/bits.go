// Based on code from the Go standard library.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the ORIGINAL_LICENSE file.

package math32

////////////////////////////////////////////////////////////////////////////////

const (
	uvnan    = 0x7F800001
	uvinf    = 0x7F800000
	uvneginf = 0xFF800000
	mask     = 0xFF
	shift    = 32 - 8 - 1
	bias     = 127
)

////////////////////////////////////////////////////////////////////////////////

// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Inf(sign int) float32 {
	var v uint32
	if sign >= 0 {
		v = uvinf
	} else {
		v = uvneginf
	}
	return Float32frombits(v)
}

// NaN returns an IEEE 754 "not-a-number" value.
func NaN() float32 { return Float32frombits(uvnan) }

// IsNaN returns whether f is an IEEE 754 "not-a-number" value.
func IsNaN(f float32) (is bool) {
	// IEEE 754 says that only NaNs satisfy `f != f`.
	// To avoid the floating-point hardware, could use:
	//	`x := Float32bits(f)`
	//	`return uint32(x>>shift)&mask == mask && x != uvinf && x != uvneginf`
	return f != f
}

// IsInf returns whether f is an infinity, according to sign.
// If sign > 0, IsInf returns whether f is positive infinity.
// If sign < 0, IsInf returns whether f is negative infinity.
// If sign == 0, IsInf returns whether f is either infinity.
func IsInf(f float32, sign int) bool {
	// Test for infinity by comparing against maximum float.
	// To avoid the floating-point hardware, could use:
	//	`x := Float32bits(f)`
	//	`return sign >= 0 && x == uvinf || sign <= 0 && x == uvneginf`
	return sign >= 0 && f > MaxFloat32 || sign <= 0 && f < -MaxFloat32
}

// Normalized returns a normal number y and exponent exp
// satisfying x == y × 2**exp. It assumes x is finite and non-zero.
func Normalized(x float32) (y float32, exp int) {
	if Abs(x) < SmallestNormalFloat32 {
		return x * (1 << 23), -23
	}
	return x, 0
}

func normalize(x float32) (y float32, exp int) {
	//TODO: ???
	return Normalized(x)
}

////////////////////////////////////////////////////////////////////////////////
