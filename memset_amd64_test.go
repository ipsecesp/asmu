package asmu

import (
	"testing"
)

func TestMemSet(t *testing.T) {
	testMemSet(t, MemSet)
}

func TestMemSetAVX(t *testing.T) {
	testMemSet(t, memsetAVX)
}

func BenchmarkMemSetAVX(b *testing.B) {
	benchMemSet(b, memsetAVX)
}
