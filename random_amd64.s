#include "textflag.h"

#define N_TRIES 20

// func RandomBytes(dst *byte, nbytes int) (int, bool)
TEXT ·RandomBytes(SB), NOSPLIT|NOFRAME, $0-24
    MOVQ    dst+(0*8)(FP), DI
    MOVQ    nbytes+(1*8)(FP), CX

    TESTQ   CX, CX
    JLE     _fail

    CMPQ    CX, $8
    JA      _pre

    XORQ    R9, R9
    MOVQ    $N_TRIES, R8

_retry0:
    RDRANDQ SI
    JC      _short
    DECQ    R8
    JZ      _fail
    JMP     _retry0

_short:
    MOVQ    CX, ret+(2*8)(FP)
    MOVQ    $1, ret+(3*8)(FP)

    MOVQ    $rdset_t<>(SB), DX
    JMP     -8(DX)(CX*8)

_pre:
    LEAQ    (DI)(CX*1), DX
    MOVQ    DX, R9
    ANDQ    $-8, R9
    SUBQ    R9, DI
    JNB     _tail

_loop:
    MOVQ    $N_TRIES, R8

_retry1:
    RDRANDQ SI
    JC      _clause
    DECQ    R8
    JZ      _fail
    JMP     _retry1

_clause:
    MOVQ    SI, (DI)(R9*1)
    ADDQ    $8, DI
    JNZ     _loop
    TESTQ   R9, DX
    JZ      _done

_tail:
    MOVQ    $N_TRIES, R8

_retry2:
    RDRANDQ SI
    JC      _tailend
    DECQ    R8
    JZ      _fail
    JMP     _retry2

_tailend:
    MOVQ    SI, -8(DX) 

_done:
    MOVQ    CX, ret+(2*8)(FP)
    MOVQ    $1, ret+(3*8)(FP)
    RET

_fail:
    MOVQ    R9, ret+(2*8)(FP)
    MOVQ    $0, ret+(3*8)(FP)
    RET


TEXT rdset1<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVB    SI, (DI)
    RET

TEXT rdset2<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVW    SI, (DI)
    RET

TEXT rdset3<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVW    SI, (DI)
    BSWAPQ  SI
    MOVB    SI, 2(DI)
    RET

TEXT rdset4<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVL    SI, (DI)
    RET

TEXT rdset5_7<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVL    SI, (DI)
    BSWAPQ  SI
    MOVL    SI, -4(DI)(CX*1)
    RET

TEXT rdset8<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVQ    SI, (DI)
    RET

DATA    rdset_t<>+( 0*8)(SB)/8,    $rdset1<>(SB)
DATA    rdset_t<>+( 1*8)(SB)/8,    $rdset2<>(SB)
DATA    rdset_t<>+( 2*8)(SB)/8,    $rdset3<>(SB)
DATA    rdset_t<>+( 3*8)(SB)/8,    $rdset4<>(SB)
DATA    rdset_t<>+( 4*8)(SB)/8,    $rdset5_7<>(SB)
DATA    rdset_t<>+( 5*8)(SB)/8,    $rdset5_7<>(SB)
DATA    rdset_t<>+( 6*8)(SB)/8,    $rdset5_7<>(SB)
DATA    rdset_t<>+( 7*8)(SB)/8,    $rdset8<>(SB)
GLOBL   rdset_t<>(SB),             (RODATA|NOPTR), $(8*8)
