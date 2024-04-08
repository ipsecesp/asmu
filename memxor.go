package asmu

import (
	"unsafe"
)

func memxorGeneric(dst, src1, src2 *byte, nbytes int) {
	const wsize = int(unsafe.Sizeof(uintptr(0)))

	if nbytes == 0 {
		return
	}
	var (
		dstptr  = unsafe.Pointer(dst)
		src1ptr = unsafe.Pointer(src1)
		src2ptr = unsafe.Pointer(src2)
	)
	if nbytes >= wsize {
		if nbytes&int(wsize-1) != 0 {
			n := nbytes - wsize
			*(*uintptr)(unsafe.Add(dstptr, n)) =
				*(*uintptr)(unsafe.Add(src1ptr, n)) ^
					*(*uintptr)(unsafe.Add(src2ptr, n))
		}
		for n := nbytes & -wsize; n > 0; n -= wsize {
			*(*uintptr)(dstptr) = *(*uintptr)(src1ptr) ^ *(*uintptr)(src2ptr)
			dstptr = unsafe.Add(dstptr, wsize)
			src1ptr = unsafe.Add(src1ptr, wsize)
			src2ptr = unsafe.Add(src2ptr, wsize)
		}
	} else {
		for ; nbytes > 0; nbytes-- {
			*(*byte)(dstptr) = *(*byte)(src1ptr) ^ *(*byte)(src2ptr)
			dstptr = unsafe.Add(dstptr, 1)
			src1ptr = unsafe.Add(src1ptr, 1)
			src2ptr = unsafe.Add(src2ptr, 1)
		}
	}
}

func memxorNaif(dst, src1, src2 *byte, nbytes int) {
	if nbytes == 0 {
		return
	}
	d := *(*[]byte)(unsafe.Pointer(&[3]uintptr{
		uintptr(unsafe.Pointer(dst)),
		uintptr(nbytes),
		uintptr(nbytes),
	}))
	s1 := *(*[]byte)(unsafe.Pointer(&[3]uintptr{
		uintptr(unsafe.Pointer(src1)),
		uintptr(nbytes),
		uintptr(nbytes),
	}))
	s2 := *(*[]byte)(unsafe.Pointer(&[3]uintptr{
		uintptr(unsafe.Pointer(src2)),
		uintptr(nbytes),
		uintptr(nbytes),
	}))
	nbytes--
	_, _, _ = d[nbytes], s1[nbytes], s2[nbytes]

	for ; nbytes >= 0; nbytes-- {
		d[nbytes] = s1[nbytes] ^ s2[nbytes]
	}
}
