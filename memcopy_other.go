//go:build !amd64

package asmu

func MemCopy(dst, src *byte, nbytes int) {
	memcopyGeneric(dst, src, nbytes)
}
