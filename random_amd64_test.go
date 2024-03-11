package asmu

import (
	"strconv"
	"testing"
	"unsafe"
)

func TestRandomBytes_Filling(t *testing.T) {
	cases := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	for i := range [16]struct{}{} {
		i = (i + 1) * 16
		cases = append(cases, i-1, i)
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
				n = -2
				p := unsafe.Add(unsafe.Pointer(&dst[0]), n)
				for ; n >= 0; n-- {
					if *(*int16)(p) == int16(0) {
						t.Logf("two or more zeros in a row (at: %d)", n)
						break
					}
					p = unsafe.Add(p, -1)
				}
			}
		})
	}
}

func TestRandUint64n(t *testing.T) {
	cases := []uint64{10, 100, 1000, 10000}
	for _, c := range cases {
		n := c
		t.Run(strconv.Itoa(int(n)), func(t *testing.T) {
			for i := 0; i < int(n); i++ {
				r, ok := RandUint64n(n)
				if !ok {
					t.Fatal(`invalid result`)
				} else if r > n {
					t.Fatalf("result is greater than the upper bound [%d, %d)", r, n)
				}
			}
		})
	}
}
