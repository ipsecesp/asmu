package asmu

import (
	"strconv"
	"testing"
	"unsafe"
)

func testMemSet(t *testing.T, memset func(*byte, byte, int)) {
	cases := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		31, 32, 63, 64, 127, 128, 255, 256, 511, 512,
	}
	for _, c := range cases {
		nbytes := c
		t.Run(strconv.Itoa(nbytes), func(t *testing.T) {
			dst, chr := make([]byte, nbytes), byte(nbytes)
			memset((*byte)(&dst[0]), chr, nbytes)

			for i := 0; i < len(dst); i++ {
				if dst[i] != chr {
					t.Fatalf("byte mismatch (at: %d; exp: %d; got: %d)", i, chr, dst[i])
				}
			}
		})
	}
}

func TestMemSetGeneric(t *testing.T) {
	testMemSet(t, memsetGeneric)
}

func TestMemSetNaif(t *testing.T) {
	testMemSet(t, memsetNaif)
}

func benchMemSet(b *testing.B, memset func(*byte, byte, int)) {
	cases := []int{1023, 1024, 4095, 4096}
	for _, c := range cases {
		nbytes := c
		b.Run(strconv.Itoa(nbytes), func(b *testing.B) {
			b.StopTimer()
			dst := make([]byte, nbytes)
			b.SetBytes(int64(nbytes))
			b.ResetTimer()
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				memset(
					*(**byte)(unsafe.Pointer(&dst)),
					byte(nbytes),
					nbytes,
				)
			}
		})
	}
}

func BenchmarkMemSetGeneric(b *testing.B) {
	benchMemSet(b, memsetGeneric)
}

func BenchmarkMemSetNaif(b *testing.B) {
	benchMemSet(b, memsetNaif)
}
