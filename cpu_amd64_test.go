package asmu

import (
	"fmt"
	"testing"
)

func TestCpuID(t *testing.T) {
	eax, _, _, _ := CpuID(0, 0)
	fmt.Printf("%d\n", eax)
}

func TestCpuXCR(t *testing.T) {
	xcr := CpuXCR(1)
	fmt.Printf("%d\n", xcr)
}

func TestCpuTSC(t *testing.T) {
	c1 := CpuTSC()
	if c1 == 0 {
		t.Fatalf("invalid first value: %d", c1)
	}
	c2 := CpuTSC()
	if c1 >= c2 {
		t.Fatalf("first value must be less than the second value: %d, %d", c1, c2)
	}
	t.Logf("TSC1 = %d, TSC2 = %d", c1, c2)
}
