package asmu

import (
	"testing"
)

func TestMemXor(t *testing.T) {
	testMemXor(t, MemXor)
}

func BenchmarkMemXor(b *testing.B) {
	benchMemXor(b, MemXor)
}
