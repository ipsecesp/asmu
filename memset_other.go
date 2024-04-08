//go:build !amd64

package asmu

func MemSet(dst *byte, chr byte, nbytes int) {
	memsetGeneric(dst, chr, nbytes)
}
