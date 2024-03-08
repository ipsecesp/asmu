#include "textflag.h"

// func ReadTSC() int64
TEXT ·ReadTSC(SB), NOSPLIT|NOFRAME, $0-8
    XORQ    CX, CX
    XORQ    AX, AX
    CPUID

    RDTSC
    SHLQ    $32, DX
    ORQ     AX, DX

    MOVQ    DX, ret+(0*8)(FP)
    RET
