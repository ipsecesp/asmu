package asmu

import "unsafe"

func memsetGeneric(dst *byte, chr byte, nbytes int) {
	const wsize = int(unsafe.Sizeof(uintptr(0)))

	if nbytes == 0 {
		return
	}
	dstptr := unsafe.Pointer(dst)

	if n := nbytes & -wsize; n > 0 {
		p := uintptr(0x0101010101010101) * uintptr(chr)
		*(*uintptr)(unsafe.Add(dstptr, nbytes-wsize)) = p

		for ; n > 0; n -= wsize {
			*(*uintptr)(dstptr) = p
			dstptr = unsafe.Add(dstptr, wsize)
		}
	} else {
		for ; nbytes > 0; nbytes-- {
			*(*byte)(dstptr) = chr
			dstptr = unsafe.Add(dstptr, 1)
		}
	}
}

func memsetNaif(dst *byte, chr byte, nbytes int) {
	if nbytes > 0 {
		d := *(*[]byte)(unsafe.Pointer(&[3]uintptr{
			uintptr(unsafe.Pointer(dst)),
			uintptr(nbytes),
			uintptr(nbytes),
		}))
		nbytes--

		for _ = d[nbytes]; nbytes >= 0; nbytes-- {
			d[nbytes] = chr
		}
	}
}
