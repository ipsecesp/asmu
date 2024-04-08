package asmu

import (
	"sync"
)

// CpuID returns processor identification and feature information in the EAX, EBX, ECX, and EDX registers.
//
//go:noescape
func CpuID(leaf, subleaf uint32) (eax, ebx, ecx, edx uint32)

// CpuXCR returns the unsigned 64-bit value of the processor's Extended Control Register.
//
//go:noescape
func CpuXCR(index uint32) uint64

// CpuTSC returns the unsigned 64-bit value of the processor's Time-Stamp Counter.
// The processor monotonically increments the Time-Stamp Counter
// every clock cycle and resets it to 0 whenever the processor is reset.
//
//go:noescape
func CpuTSC() uint64

var (
	detectOnce sync.Once

	hasAESNI  bool
	hasAVX    bool
	hasAVX2   bool
	hasBMI1   bool
	hasBMI2   bool
	hasRDRAND bool
	hasRDSEED bool
	hasRDTSC  bool
	hasSHA    bool
	hasSSE    bool
	hasSSE2   bool
	hasSSE3   bool
	hasSSE4   bool
	hasSSE42  bool
	hasSSSE3  bool
)

func detect() {
	detectOnce.Do(func() {
		leafMax, _, _, _ := CpuID(0, 0)
		if leafMax > 0 {
			_, _, ecx, edx := CpuID(1, 0)
			hasSSE3 = ecx&(1<<0) != 0
			hasSSSE3 = ecx&(1<<9) != 0
			hasSSE4 = ecx&(1<<19) != 0
			hasSSE42 = ecx&(1<<20) != 0
			hasAESNI = ecx&(1<<25) != 0
			hasAVX = ecx&(1<<28) != 0
			hasRDRAND = ecx&(1<<30) != 0
			hasRDTSC = edx&(1<<4) != 0
			hasSSE = edx&(1<<25) != 0
			hasSSE2 = edx&(1<<26) != 0
		}
		if leafMax >= 7 {
			_, ebx, _, _ := CpuID(7, 0)
			hasBMI1 = ebx&(1<<3) != 0
			hasAVX2 = ebx&(1<<5) != 0
			hasBMI2 = ebx&(1<<8) != 0
			hasRDSEED = ebx&(1<<18) != 0
			hasSHA = ebx&(1<<29) != 0
		}
	})
}
