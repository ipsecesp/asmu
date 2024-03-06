package asmu

import (
	"strconv"
	"testing"
	"unsafe"
)

func TestMemXor(t *testing.T) {
	cases := [...]struct {
		nbytes int
		chr1   byte
		chr2   byte
	}{
		{0, 1, 2},
		{1, 1, 2},
		{3, 3, 3 - 1},
		{8, 8, 8 - 1},
		{9, 9, 9 - 1},
		{15, 15, 15 - 1},
		{16, 16, 16 - 1},
		{17, 17, 17 - 1},
		{32, 32, 32 - 1},
		{33, 33, 33 - 1},
		{64, 64, 64 - 1},
		{65, 65, 65 - 1},
		{127, 127, 127 - 1},
		{128, 128, 128 - 1},
		{200, 200, 200 - 1},
	}
	for _, c := range cases {
		t.Run(strconv.Itoa(c.nbytes), func(t *testing.T) {
			dst := make([]byte, c.nbytes)
			src1 := make([]byte, len(dst))
			src2 := make([]byte, len(dst))
			for i := 0; i < len(dst); i++ {
				src1[i] = c.chr1
				src2[i] = c.chr2
			}
			MemXor(
				*(**byte)(unsafe.Pointer(&dst)),
				*(**byte)(unsafe.Pointer(&src1)),
				*(**byte)(unsafe.Pointer(&src2)),
				c.nbytes,
			)
			for i := 0; i < len(dst); i++ {
				if exp := c.chr1 ^ c.chr2; dst[i] != exp {
					t.Fatalf("byte[%d] mismatch (exp: %d; got: %d)", i, exp, dst[i])
				}
			}
		})
	}
}

func BenchmarkMemXor(b *testing.B) {
	cases := [...]struct {
		nbytes int
		chr1   byte
		chr2   byte
	}{
		{1 * 1024, 1, 2},
		{4 * 1024, 2, 3},
		{16 * 1024, 3, 4},
		{64 * 1024, 4, 5},
	}
	for _, c := range cases {
		b.Run(strconv.Itoa(c.nbytes), func(b *testing.B) {
			dst := make([]byte, c.nbytes)
			src1 := make([]byte, c.nbytes)
			src2 := make([]byte, c.nbytes)
			for i := 0; i < c.nbytes; i++ {
				src1[i] = c.chr1 + byte(i)
				src2[i] = c.chr2 + byte(i)
			}
			b.SetBytes(int64(c.nbytes))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				MemXor(
					*(**byte)(unsafe.Pointer(&dst)),
					*(**byte)(unsafe.Pointer(&src1)),
					*(**byte)(unsafe.Pointer(&src2)),
					c.nbytes,
				)
			}
		})
	}
}
