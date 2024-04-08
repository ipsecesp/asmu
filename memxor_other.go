//go:build !amd64

package asmu

func MemXor(dst, src1, src2 *byte, nbytes int) {
	memxorGeneric(dst, src1, src2, nbytes)
}
