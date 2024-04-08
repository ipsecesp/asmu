package asmu

import (
	"unsafe"
)

func memcopyGeneric(dst, src *byte, nbytes int) {
	const (
		wsize = int(unsafe.Sizeof(uintptr(0)))
		wmask = int(wsize - 1)
	)
	if nbytes <= 0 || dst == src {
		return
	}
	var (
		dstptr = unsafe.Pointer(dst)
		srcptr = unsafe.Pointer(src)
		n      = nbytes & -wsize
	)
	if uintptr(dstptr) < uintptr(srcptr) {
		// Copies forward
		for ; n > 0; n -= wsize {
			*(*uintptr)(dstptr) = *(*uintptr)(srcptr)
			dstptr = unsafe.Add(dstptr, wsize)
			srcptr = unsafe.Add(srcptr, wsize)
		}
		for n = nbytes & wmask; n > 0; n-- {
			*(*byte)(dstptr) = *(*byte)(srcptr)
			dstptr = unsafe.Add(dstptr, 1)
			srcptr = unsafe.Add(srcptr, 1)
		}
	} else {
		// Copies backward
		dstptr = unsafe.Add(dstptr, nbytes)
		srcptr = unsafe.Add(srcptr, nbytes)

		for ; n > 0; n -= wsize {
			dstptr = unsafe.Add(dstptr, -wsize)
			srcptr = unsafe.Add(srcptr, -wsize)
			*(*uintptr)(dstptr) = *(*uintptr)(srcptr)
		}
		for n = nbytes & wmask; n > 0; n-- {
			dstptr = unsafe.Add(dstptr, -1)
			srcptr = unsafe.Add(srcptr, -1)
			*(*byte)(dstptr) = *(*byte)(srcptr)
		}
	}
}

func memcopyNaif(dst, src *byte, nbytes int) {
	if nbytes <= 0 || dst == src {
		return
	}
	dstptr := unsafe.Pointer(dst)
	srcptr := unsafe.Pointer(src)
	d := *(*[]byte)(unsafe.Pointer(&[3]uintptr{
		uintptr(dstptr),
		uintptr(nbytes),
		uintptr(nbytes),
	}))
	s := *(*[]byte)(unsafe.Pointer(&[3]uintptr{
		uintptr(srcptr),
		uintptr(nbytes),
		uintptr(nbytes),
	}))
	nbytes--
	_, _ = d[nbytes], s[nbytes]

	if uintptr(dstptr) < uintptr(srcptr) {
		for i := 0; i <= nbytes; i++ {
			d[i] = s[i]
		}
	} else {
		for ; nbytes >= 0; nbytes-- {
			d[nbytes] = s[nbytes]
		}
	}
}
