package asmu

import (
	"strconv"
	"testing"
	"unsafe"
)

func TestMemSet(t *testing.T) {
	cases := [...]struct {
		nbytes int
		chr    byte
	}{
		{0, '0' + 0},
		{1, '0' + 1},
		{2, '0' + 2},
		{3, '0' + 3},
		{4, '0' + 4},
		{5, '0' + 5},
		{8, '0' + 8},
		{9, '0' + 9},
		{15, 'A' + 0},
		{16, 'A' + 1},
		{17, 'A' + 2},
		{31, 'A' + 3},
		{32, 'A' + 4},
		{64, 'A' + 5},
		{65, 'A' + 6},
		{127, 'A' + 7},
		{128, 'A' + 8},
		{256, 'A' + 9},
	}
	for _, c := range cases {
		t.Run(strconv.Itoa(c.nbytes), func(t *testing.T) {
			dst := make([]byte, c.nbytes)
			MemSet(*(**byte)(unsafe.Pointer(&dst)), c.chr, c.nbytes)

			for i := 0; i < len(dst); i++ {
				if dst[i] != c.chr {
					t.Fatalf("byte[%d] mismatch (exp: %d; got: %d)", i, c.chr, dst[i])
				}
			}
		})
	}
}
