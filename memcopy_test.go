package asmu

import (
	"strconv"
	"testing"
	"unsafe"
)

func memcpyInit(length, offset int) (exp, dst, src []byte) {
	fwd := true
	if offset < 0 {
		fwd = false
		offset *= -1
	}
	exp = make([]byte, length)
	dst = make([]byte, length+offset)
	if dst, src = dst[:length], dst[offset:]; !fwd {
		dst, src = src, dst
	}
	for i := 0; i < len(src); i++ {
		src[i] = byte(length + i)
		exp[i] = src[i]
	}
	return
}

func testMemCopy(t *testing.T, offset int, memcpy func(*byte, *byte, int)) {
	cases := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
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
			exp, dst, src := memcpyInit(nbytes, offset)
			memcpy((*byte)(&dst[0]), (*byte)(&src[0]), nbytes)

			for i := 0; i < nbytes; i++ {
				if dst[i] != exp[i] {
					t.Fatalf("byte mismatch (at: %d; exp: %d; got: %d)", i, exp[i], dst[i])
				}
			}
		})
	}
}

func TestMemCopyGeneric_Forward(t *testing.T) {
	testMemCopy(t, 1, memcopyGeneric)
}
func TestMemCopyGeneric_Backward(t *testing.T) {
	testMemCopy(t, -1, memcopyGeneric)
}

func TestMemCopyNaif_Forward(t *testing.T) {
	testMemCopy(t, 1, memcopyNaif)
}

func TestMemCopyNaif_Backward(t *testing.T) {
	testMemCopy(t, -1, memcopyNaif)
}

func benchMemCopy(b *testing.B, offset int, memcpy func(*byte, *byte, int)) {
	cases := []int{255, 256, 1023, 1024, 4095, 4096, 8191, 8192}
	for _, c := range cases {
		nbytes := c
		b.Run(strconv.Itoa(nbytes), func(b *testing.B) {
			b.StopTimer()
			_, dst, src := memcpyInit(nbytes, offset)
			b.SetBytes(int64(nbytes))
			b.ResetTimer()
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				memcpy(
					*(**byte)(unsafe.Pointer(&dst)),
					*(**byte)(unsafe.Pointer(&src)),
					nbytes,
				)
			}
		})
	}
}

func BenchmarkMemCopyGeneric_Forward(b *testing.B) {
	benchMemCopy(b, 1, memcopyGeneric)
}

func BenchmarkMemCopyGeneric_Backward(b *testing.B) {
	benchMemCopy(b, -1, memcopyGeneric)
}

func BenchmarkMemCopyNaif_Forward(b *testing.B) {
	benchMemCopy(b, 1, memcopyNaif)
}

func BenchmarkMemCopyNaif_Backward(b *testing.B) {
	benchMemCopy(b, -1, memcopyNaif)
}
