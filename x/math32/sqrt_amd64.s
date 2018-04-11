// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

// Based on code from the Go standard library.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the ORIGINAL_LICENSE file.

#include "textflag.h"

////////////////////////////////////////////////////////////////////////////////

// func Sqrt(x float32) float32
TEXT ·Sqrt(SB), NOSPLIT, $0
	XORPS  X0, X0 // break dependency
	SQRTSS     x+0(FP), X0
	MOVSS      X0, ret+8(FP)
	RET

////////////////////////////////////////////////////////////////////////////////
