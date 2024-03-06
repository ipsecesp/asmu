#include "textflag.h"

// func MemCopy(dst, src *byte, nbytes int)
TEXT ·MemCopy(SB), NOSPLIT|NOFRAME, $0-24
    MOVQ    dst+(0*8)(FP), DI
    MOVQ    src+(1*8)(FP), SI
    MOVQ    nbytes+(2*8)(FP), CX

    TESTQ   CX, CX
    JLE     _done
    TESTQ   DI, SI
    JZ      _done

    CMPQ    CX, $16
    JA      _gt_16
    MOVQ    $memcpy_t<>(SB), DX
    JMP     -8(DX)(CX*8)

_gt_16:
    CMPQ    CX, $32
    JA      _gt_32

    MOVOU   (SI), X0
    MOVOU   -16(SI)(CX*1), X1
    MOVOU   X0, (DI)
    MOVOU   X1, -16(DI)(CX*1)
    RET

_gt_32:
    CMPQ    CX, $256
    JA      _gt_256

_gt_32_256:
    VMOVUPS (0*32)(SI), Y0
    VMOVUPS (-1*32)(SI)(CX*1), Y7
    CMPQ    CX, $64
    JBE     _gt_32_64

    VMOVUPS (1*32)(SI), Y1
    VMOVUPS (-2*32)(SI)(CX*1), Y6
    CMPQ    CX, $128
    JBE     _gt_64_128

    VMOVUPS (2*32)(SI), Y2
    VMOVUPS (-3*32)(SI)(CX*1), Y5
    CMPQ    CX, $192
    JBE     _gt_128_192

    VMOVUPS (3*32)(SI), Y3
    VMOVUPS (-4*32)(SI)(CX*1), Y4
    CMPQ    CX, $256
    JBE     _gt_192_256

_gt_256:
    CMPQ    DI, SI
    JA      _bwd

    MOVQ    CX, R9
    ANDQ    $-128, R9
    XORQ    R8, R8

_fwd_loop128:
    VMOVUPS (0*32)(SI)(R8*1), Y0
    VMOVUPS (1*32)(SI)(R8*1), Y1
    VMOVUPS (2*32)(SI)(R8*1), Y2
    VMOVUPS (3*32)(SI)(R8*1), Y3
    VMOVUPS Y0, (0*32)(DI)(R8*1)
    VMOVUPS Y1, (1*32)(DI)(R8*1)
    VMOVUPS Y2, (2*32)(DI)(R8*1)
    VMOVUPS Y3, (3*32)(DI)(R8*1)
    LEAQ    128(R8), R8
    CMPQ    R8, R9
    JNE     _fwd_loop128

    ANDQ    $127, CX
    JZ      _done
    
    LEAQ    (DI)(R9*1), DI
    LEAQ    (SI)(R9*1), SI
    
    CMPQ    CX, $32
    JA      _gt_32_256

    VMOVUPS -32(SI)(CX*1), Y4
    VMOVUPS Y4, -32(DI)(CX*1)
    VZEROUPPER
    RET 

_bwd:
    MOVQ    CX, R9
    ANDQ    $-128, R9
    NEGQ    R9
    XORQ    R8, R8

    LEAQ    -32(DI)(CX*1), DI
    LEAQ    -32(SI)(CX*1), SI

_bwd_loop128:
    VMOVUPS (-3*32)(SI)(R8*1), Y0
    VMOVUPS (-2*32)(SI)(R8*1), Y1
    VMOVUPS (-1*32)(SI)(R8*1), Y2
    VMOVUPS (-0*32)(SI)(R8*1), Y3
    VMOVUPS Y3, (-0*32)(DI)(R8*1)
    VMOVUPS Y2, (-1*32)(DI)(R8*1)
    VMOVUPS Y1, (-2*32)(DI)(R8*1)
    VMOVUPS Y0, (-3*32)(DI)(R8*1)
    LEAQ    -128(R8), R8
    CMPQ    R8, R9
    JNE     _bwd_loop128

    ANDQ    $127, CX
    JZ      _done

    SUBQ    CX, R8
    LEAQ    (DI)(R8*1), DI
    LEAQ    (SI)(R8*1), SI

    CMPQ    CX, $32
    JA      _gt_32_256

    VMOVUPS -32(SI)(CX*1), Y4
    VMOVUPS Y4, -32(DI)(CX*1)
    VZEROUPPER
    RET

_gt_192_256:
    VMOVUPS Y3, (3*32)(DI)
    VMOVUPS Y4, (-4*32)(DI)(CX*1)

_gt_128_192:
    VMOVUPS Y2, (2*32)(DI)
    VMOVUPS Y5, (-3*32)(DI)(CX*1)

_gt_64_128:
    VMOVUPS Y1, (1*32)(DI)
    VMOVUPS Y6, (-2*32)(DI)(CX*1)

_gt_32_64:
    VMOVUPS Y0, (0*32)(DI)
    VMOVUPS Y7, (-1*32)(DI)(CX*1)

_done:
    VZEROUPPER
    RET


TEXT memcpy1<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVB    (SI), AL
    MOVB    AL, (DI)
    RET

TEXT memcpy2<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVW    (SI), AX
    MOVW    AX, (DI)
    RET

TEXT memcpy3<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVW    (SI), AX
    MOVW    -2(SI)(CX*1), BX
    MOVW    AX, (DI)
    MOVW    BX, -2(DI)(CX*1)
    RET

TEXT memcpy4<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVL    (SI), AX
    MOVL    AX, (DI)
    RET

TEXT memcpy5_7<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVL    (SI), AX
    MOVL    -4(SI)(CX*1), BX
    MOVL    AX, (DI)
    MOVL    BX, -4(DI)(CX*1)
    RET

TEXT memcpy8<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVQ    (SI), AX
    MOVQ    AX, (DI)
    RET

TEXT memcpy9_15<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVQ    (SI), AX
    MOVQ    -8(SI)(CX*1), BX
    MOVQ    AX, (DI)
    MOVQ    BX, -8(DI)(CX*1)
    RET

TEXT memcpy16<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVOU   (SI), X0
    MOVOU   X0, (DI)
    RET

DATA    memcpy_t<>+( 0*8)(SB)/8,    $memcpy1<>(SB)
DATA    memcpy_t<>+( 1*8)(SB)/8,    $memcpy2<>(SB)
DATA    memcpy_t<>+( 2*8)(SB)/8,    $memcpy3<>(SB)
DATA    memcpy_t<>+( 3*8)(SB)/8,    $memcpy4<>(SB)
DATA    memcpy_t<>+( 4*8)(SB)/8,    $memcpy5_7<>(SB)
DATA    memcpy_t<>+( 5*8)(SB)/8,    $memcpy5_7<>(SB)
DATA    memcpy_t<>+( 6*8)(SB)/8,    $memcpy5_7<>(SB)
DATA    memcpy_t<>+( 7*8)(SB)/8,    $memcpy8<>(SB)
DATA    memcpy_t<>+( 8*8)(SB)/8,    $memcpy9_15<>(SB)
DATA    memcpy_t<>+( 9*8)(SB)/8,    $memcpy9_15<>(SB)
DATA    memcpy_t<>+(10*8)(SB)/8,    $memcpy9_15<>(SB)
DATA    memcpy_t<>+(11*8)(SB)/8,    $memcpy9_15<>(SB)
DATA    memcpy_t<>+(12*8)(SB)/8,    $memcpy9_15<>(SB)
DATA    memcpy_t<>+(13*8)(SB)/8,    $memcpy9_15<>(SB)
DATA    memcpy_t<>+(14*8)(SB)/8,    $memcpy9_15<>(SB)
DATA    memcpy_t<>+(15*8)(SB)/8,    $memcpy16<>(SB)
GLOBL   memcpy_t<>(SB),             (RODATA|NOPTR), $(16*8)
