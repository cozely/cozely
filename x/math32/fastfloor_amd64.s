// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

////////////////////////////////////////////////////////////////////////////////

// SLOWER than the Go function
// func fastFloorAsm(s float32) int32
TEXT ·fastFloorAsm(SB),7,$0
	CVTTSS2SL  s+0(FP), BX
	MOVL       s+0(FP), AX
	SHRL       $31, AX
	SUBL       AX, BX
	MOVL       BX,ret+8(FP)
	RET

////////////////////////////////////////////////////////////////////////////////
