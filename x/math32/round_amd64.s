// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

////////////////////////////////////////////////////////////////////////////////

// SLOWER than the Go function
// func round_asm(s float32) float32
TEXT ·roundAsm(SB),7,$0
	CVTSS2SL  s+0(FP), BX
	MOVL       BX, ret+8(FP)
	RET

////////////////////////////////////////////////////////////////////////////////
