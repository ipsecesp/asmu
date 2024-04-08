package asmu

//go:noescape
func MemCopy(dst, src *byte, nbytes int)

//go:noescape
func memcopyAVX(dst, src *byte, nbytes int)

var (
	memcopyImpl func(dst, src *byte, nbytes int) = memcopyGeneric
)

func init() {
	detect()

	if hasAVX {
		memcopyImpl = memcopyAVX
	}
}
