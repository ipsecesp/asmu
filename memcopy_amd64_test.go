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

func TestMemCopyAVX_Forward(t *testing.T) {
	testMemCopy(t, 1, memcopyAVX)
}
func TestMemCopyAVX_Backward(t *testing.T) {
	testMemCopy(t, -1, memcopyAVX)
}

func BenchmarkMemCopyAVX_Forward(b *testing.B) {
	benchMemCopy(b, 1, memcopyAVX)
}

func BenchmarkMemCopyAVX_Backward(b *testing.B) {
	benchMemCopy(b, -1, memcopyAVX)
}
