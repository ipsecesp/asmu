#include "textflag.h"

#define N_TRIES 20

// func RandUint64n(n uint64) (uint64, bool)
TEXT ·RandUint64n(SB), NOSPLIT|NOFRAME, $0-18
    MOVQ    n+(0*8)(FP), AX

    TESTQ   AX, AX
    JLE     _fail

    LEAQ    -1(AX), AX
    MOVQ    AX, CX
    ORQ     $1, CX
    BSRQ    CX, CX
    LEAQ    -63(CX), CX
    NEGQ    CX

    MOVQ    $-1, DX
    SHRQ    CX, DX

_loop:
    MOVQ    $N_TRIES, R8

_retry:
    RDRANDQ SI
    JC      _clause
    DECQ    R8
    JZ      _fail
    JMP     _retry

_clause:
    ANDQ    DX, SI
    CMPQ    SI, AX
    JA      _loop

_done:
    MOVQ    SI, ret+(1*8)(FP)
    MOVQ    $1, ret+(2*8)(FP)
    RET

_fail:
    MOVQ    $0, ret+(1*8)(FP)
    MOVQ    $0, ret+(2*8)(FP)
    RET

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

    MOVQ    $rndset_t<>(SB), DX
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


TEXT rndset1<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVB    SI, (DI)
    RET

TEXT rndset2<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVW    SI, (DI)
    RET

TEXT rndset3<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVW    SI, (DI)
    BSWAPQ  SI
    MOVB    SI, 2(DI)
    RET

TEXT rndset4<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVL    SI, (DI)
    RET

TEXT rndset5_7<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVL    SI, (DI)
    BSWAPQ  SI
    MOVL    SI, -4(DI)(CX*1)
    RET

TEXT rndset8<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVQ    SI, (DI)
    RET

DATA    rndset_t<>+( 0*8)(SB)/8,    $rndset1<>(SB)
DATA    rndset_t<>+( 1*8)(SB)/8,    $rndset2<>(SB)
DATA    rndset_t<>+( 2*8)(SB)/8,    $rndset3<>(SB)
DATA    rndset_t<>+( 3*8)(SB)/8,    $rndset4<>(SB)
DATA    rndset_t<>+( 4*8)(SB)/8,    $rndset5_7<>(SB)
DATA    rndset_t<>+( 5*8)(SB)/8,    $rndset5_7<>(SB)
DATA    rndset_t<>+( 6*8)(SB)/8,    $rndset5_7<>(SB)
DATA    rndset_t<>+( 7*8)(SB)/8,    $rndset8<>(SB)
GLOBL   rndset_t<>(SB),             (RODATA|NOPTR), $(8*8)
