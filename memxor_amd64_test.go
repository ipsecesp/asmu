package asmu

import (
	"testing"
)

func TestMemXor(t *testing.T) {
	testMemXor(t, MemXor)
}

func TestMemXorAVX(t *testing.T) {
	testMemXor(t, memxorAVX)
}

func BenchmarkMemXorAVX(b *testing.B) {
	benchMemXor(b, memxorAVX)
}
