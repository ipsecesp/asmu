package asmu

import (
	"strconv"
	"testing"
	"unsafe"
)

func TestRandomBytes_Filling(t *testing.T) {
	cases := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
		15, 16, 31, 32, 63, 64, 127, 128,
	}
	for _, c := range cases {
		nbytes := c
		t.Run(strconv.Itoa(nbytes), func(t *testing.T) {
			dst := make([]byte, nbytes)
			n, ok := RandomBytes(*(**byte)(unsafe.Pointer(&dst)), len(dst))
			if !ok {
				if n == nbytes {
					t.Error(`full filling: false status`)
				} else {
					t.Log(`partial filling`)
				}
			} else if n != nbytes {
				t.Error(`partial filling: incorrect status`)
			}
			if n == 0 || (n == 1 && dst[0] == 0) {
				t.Error(`empty result`)
			} else if n >= 2 {
				p := unsafe.Add(unsafe.Pointer(&dst[0]), n-2)
				for i := 0; i < n; i++ {
					if *(*int16)(p) == int16(0) {
						t.Errorf("two or more zeros in a row (at: %d)", n-i)
						break
					}
					p = unsafe.Add(p, -1)
				}
			}
		})
	}
}
