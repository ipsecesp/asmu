#include "textflag.h"

// func CpuID(leaf, subleaf uint32) (eax, ebx, ecx, edx uint32)
TEXT ·CpuID(SB), NOSPLIT|NOFRAME, $0-24
    MOVL    leaf+(0*4)(FP), AX
    MOVL    subleaf+(1*4)(FP), CX
    
    CPUID

    MOVL    AX, eax+(2*4)(FP)
    MOVL    BX, ebx+(3*4)(FP)
    MOVL    CX, ecx+(4*4)(FP)
    MOVL    DX, edx+(5*4)(FP)
    RET


// func CpuXCR(index uint32) uint64
TEXT ·CpuXCR(SB), NOSPLIT|NOFRAME, $0-16
    MOVL    index+0(FP), CX

    XGETBV
    SHLQ    $32, DX
    ORQ     AX, DX
    
    MOVQ    DX, ret+8(FP)
    RET


// func CpuTSC() uint64
TEXT ·CpuTSC(SB), NOSPLIT|NOFRAME, $0-8
    XORQ    AX, AX
    XORQ    CX, CX
    CPUID

    RDTSC
    SHLQ    $32, DX
    ORQ     AX, DX

    MOVQ    DX, ret+0(FP)
    RET
