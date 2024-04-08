#include "textflag.h"

// func MemXor(dst, src1, src2 *byte, nbytes int)
TEXT ·MemXor(SB), NOSPLIT|NOFRAME, $0-32
    MOVQ    ·memxorImpl(SB), DX
    JMP     (DX)


// func memxorAVX(dst, src1, src2 *byte, nbytes int)
TEXT ·memxorAVX(SB), NOSPLIT|NOFRAME, $0-32
    MOVQ    dst+(0*8)(FP), DI
    MOVQ    src1+(1*8)(FP), SI
    MOVQ    src2+(2*8)(FP), DX
    MOVQ    nbytes+(3*8)(FP), CX

    TESTQ   CX, CX
    JLE     _done

    MOVQ    CX, AX

    CMPQ    CX, $16
    JAE     _ge_16
    
    XORQ    CX, CX
    JMP     _loop1

_ge_16:
    CMPQ    CX, $128
    JAE     _ge_128
    XORQ    CX, CX
    JMP     _lt_128

_ge_128:
    MOVQ    AX, CX
    ANDQ    $-128, CX
    XORQ    R8, R8

_loop128:
    VMOVUPS (0*32)(DX)(R8*1), Y0
    VMOVUPS (1*32)(DX)(R8*1), Y1
    VMOVUPS (2*32)(DX)(R8*1), Y2
    VMOVUPS (3*32)(DX)(R8*1), Y3
    VXORPS  (0*32)(SI)(R8*1), Y0, Y0
    VXORPS  (1*32)(SI)(R8*1), Y1, Y1
    VXORPS  (2*32)(SI)(R8*1), Y2, Y2
    VXORPS  (3*32)(SI)(R8*1), Y3, Y3
    VMOVUPS Y0, (0*32)(DI)(R8*1)
    VMOVUPS Y1, (1*32)(DI)(R8*1)
    VMOVUPS Y2, (2*32)(DI)(R8*1)
    VMOVUPS Y3, (3*32)(DI)(R8*1)
    LEAQ    128(R8), R8
    CMPQ    R8, CX
    JNE     _loop128
    CMPQ    CX, AX
    JE      _done
    TESTB   $112, AL
    JE      _loop1

_lt_128:
    MOVQ    CX, R8
    MOVQ    AX, CX
    ANDQ    $-16, CX

_loop16:
    VMOVUPS (DX)(R8*1), X0
    VXORPS  (SI)(R8*1), X0, X0
    VMOVUPS X0, (DI)(R8*1)
    
    LEAQ    16(R8), R8
    CMPQ    R8, CX
    JNE     _loop16
    CMPQ    CX, AX
    JE      _done

_loop1:
    MOVBQZX (DX)(CX*1), R8
    XORB    (SI)(CX*1), R8
    MOVB    R8, (DI)(CX*1)
    
    LEAQ    1(CX), CX
    CMPQ    CX, AX
    JNE     _loop1

_done:
    VZEROUPPER
    RET
