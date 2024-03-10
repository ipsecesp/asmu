package asmu

import (
	"strconv"
	"testing"
	"unsafe"
)

func memxorInit(length int) (dst, src1, src2 []byte) {
	dst = make([]byte, 3*length)
	dst, src1, src2 = dst[:length], dst[length:2*length], dst[2*length:]
	for i := 0; i < length; i++ {
		src1[i], src2[i] = byte(i+1), byte(i+2)
	}
	return
}

func testMemXor(t *testing.T, memxor func(*byte, *byte, *byte, int)) {
	cases := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for i := range [16]struct{}{} {
		i = (i + 1) * 16
		cases = append(cases, i-1, i)
	}
	for i := range [8]struct{}{} {
		i = 256 + (i+1)*32
		cases = append(cases, i-1, i)
	}
	for _, c := range cases {
		nbytes := c
		t.Run(strconv.Itoa(nbytes), func(t *testing.T) {
			dst, src1, src2 := memxorInit(nbytes)
			memxor((*byte)(&dst[0]), (*byte)(&src1[0]), (*byte)(&src2[0]), nbytes)

			for i := 0; i < nbytes; i++ {
				if exp := src1[i] ^ src2[i]; exp != dst[i] {
					t.Fatalf("byte mismatch (at: %d; exp: %d; got: %d)", i, exp, dst[i])
				}
			}
		})
	}
}

func TestMemXorGeneric(t *testing.T) {
	testMemXor(t, memxorGeneric)
}

func TestMemXorNaif(t *testing.T) {
	testMemXor(t, memxorNaif)
}

func benchMemXor(b *testing.B, memxor func(*byte, *byte, *byte, int)) {
	cases := []int{255, 256, 1023, 1024, 4095, 4096, 8191, 8192}
	for _, c := range cases {
		nbytes := c
		b.Run(strconv.Itoa(nbytes), func(b *testing.B) {
			b.StopTimer()
			dst, src1, src2 := memxorInit(nbytes)
			b.SetBytes(int64(nbytes))
			b.ResetTimer()
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				memxor(
					*(**byte)(unsafe.Pointer(&dst)),
					*(**byte)(unsafe.Pointer(&src1)),
					*(**byte)(unsafe.Pointer(&src2)),
					nbytes,
				)
			}
		})
	}
}

func BenchmarkMemXorGeneric(b *testing.B) {
	benchMemXor(b, memxorGeneric)
}

func BenchmarkMemXorNaif(b *testing.B) {
	benchMemXor(b, memxorNaif)
}
