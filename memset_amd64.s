#include "textflag.h"

// func MemSet(dst *byte, chr byte, nbytes int)
TEXT ·MemSet(SB), NOSPLIT|NOFRAME, $0-24
    MOVQ    ·memsetImpl(SB), DX
    JMP     (DX)


// func memsetAVX(dst *byte, chr byte, nbytes int)
TEXT ·memsetAVX(SB), NOSPLIT|NOFRAME, $0-24
    MOVQ    dst+(0*8)(FP), DI
    MOVBQZX chr+(1*8)(FP), SI
    MOVQ    nbytes+(2*8)(FP), CX

    TESTQ   CX, CX
    JLE     _done

    MOVQ    $0x0101010101010101, DX
    IMULQ   DX, SI

    CMPQ    CX, $16
    JA      _gt_16
    MOVQ    $memset_t<>(SB), DX
    JMP     -8(DX)(CX*8)

_gt_16:
    MOVQ    SI, X0
    PSHUFD  $0, X0, X0

    CMPQ    CX, $32
    JBE     _set17_32

    LEAQ    (DI)(CX*1), DX

    MOVUPS  X0, (DI)
    ADDQ    $16, DI
    ANDQ    $-16, DI
    MOVAPS  X0, (DI)

    ADDQ    $16, DI
    ANDQ    $-32, DI

    MOVQ    DX, CX
    ANDQ    $-32, CX
    SUBQ    CX, DI
    JNB     _tail

    VINSERTF128 $1, X0, Y0, Y0

_loop32:
    VMOVAPS Y0, (DI)(CX*1)
    ADDQ    $32, DI
    JNZ     _loop32
    VZEROUPPER

_tail:
    MOVUPS  X0, -32(DX)
    MOVUPS  X0, -16(DX)
    RET

_set17_32:
    MOVUPS  X0, (DI)
    MOVUPS  X0, -16(DI)(CX*1)

_done:
    RET


TEXT memset1<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVB    SI, (DI)
    RET

TEXT memset2<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVW    SI, (DI)
    RET

TEXT memset3<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVB    SI, (DI)
    MOVW    SI, 1(DI)
    RET

TEXT memset4<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVL    SI, (DI)
    RET

TEXT memset5_7<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVL    SI, (DI)
    MOVL    SI, -4(DI)(CX*1)
    RET

TEXT memset8<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVQ    SI, (DI)
    RET

TEXT memset9_16<>(SB), NOSPLIT|NOFRAME, $0-0
    MOVQ    SI, (DI)
    MOVQ    SI, -8(DI)(CX*1)
    RET

DATA    memset_t<>+( 0*8)(SB)/8,    $memset1<>(SB)
DATA    memset_t<>+( 1*8)(SB)/8,    $memset2<>(SB)
DATA    memset_t<>+( 2*8)(SB)/8,    $memset3<>(SB)
DATA    memset_t<>+( 3*8)(SB)/8,    $memset4<>(SB)
DATA    memset_t<>+( 4*8)(SB)/8,    $memset5_7<>(SB)
DATA    memset_t<>+( 5*8)(SB)/8,    $memset5_7<>(SB)
DATA    memset_t<>+( 6*8)(SB)/8,    $memset5_7<>(SB)
DATA    memset_t<>+( 7*8)(SB)/8,    $memset8<>(SB)
DATA    memset_t<>+( 8*8)(SB)/8,    $memset9_16<>(SB)
DATA    memset_t<>+( 9*8)(SB)/8,    $memset9_16<>(SB)
DATA    memset_t<>+(10*8)(SB)/8,    $memset9_16<>(SB)
DATA    memset_t<>+(11*8)(SB)/8,    $memset9_16<>(SB)
DATA    memset_t<>+(12*8)(SB)/8,    $memset9_16<>(SB)
DATA    memset_t<>+(13*8)(SB)/8,    $memset9_16<>(SB)
DATA    memset_t<>+(14*8)(SB)/8,    $memset9_16<>(SB)
DATA    memset_t<>+(15*8)(SB)/8,    $memset9_16<>(SB)
GLOBL   memset_t<>(SB),             (RODATA|NOPTR), $(16*8)
