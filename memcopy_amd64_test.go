package asmu

import (
	"testing"
)

func TestMemCopy_Forward(t *testing.T) {
	testMemCopy(t, 1, MemCopy)
}
func TestMemCopy_Backward(t *testing.T) {
	testMemCopy(t, -1, MemCopy)
}

func BenchmarkMemCopy_Forward(b *testing.B) {
	benchMemCopy(b, 1, MemCopy)
}

func BenchmarkMemCopy_Backward(b *testing.B) {
	benchMemCopy(b, -1, MemCopy)
}
