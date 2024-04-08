package asmu

//go:noescape
func MemXor(dst, src1, src2 *byte, nbytes int)

//go:noescape
func memxorAVX(dst, src1, src2 *byte, nbytes int)

var (
	memxorImpl func(dst, src1, src2 *byte, nbytes int) = memxorGeneric
)

func init() {
	detect()

	if hasAVX {
		memxorImpl = memxorAVX
	}
}
