#include "textflag.h"

// func ReadTSC() int64
TEXT ·ReadTSC(SB), NOSPLIT|NOFRAME, $0-8
    XORQ    CX, CX
    XORQ    AX, AX
    CPUID

    RDTSC
    SHLQ    $32, DX
    ORQ     AX, DX
    MOVQ    DX, R8

    XORQ    AX, AX
    CPUID

    MOVQ    R8, ret+(0*8)(FP)
    RET
