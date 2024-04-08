package asmu

// Memset fills the first bytes of the nbytes of the array
// pointed to by dst with the value of chr.
//
//go:noescape
func MemSet(dst *byte, chr byte, nbytes int)

//go:noescape
func memsetAVX(dst *byte, chr byte, nbytes int)

var (
	memsetImpl func(dst *byte, chr byte, nbytes int) = memsetGeneric
)

func init() {
	detect()

	if hasAVX {
		memsetImpl = memsetAVX
	}
}
