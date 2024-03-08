package asmu

import (
	"testing"
)

func TestMemSet(t *testing.T) {
	testMemSet(t, MemSet)
}

func BenchmarkMemSet(b *testing.B) {
	benchMemSet(b, MemSet)
}
