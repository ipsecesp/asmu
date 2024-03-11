package asmu

//go:noescape
func RandUint64n(n uint64) (uint64, bool)

//go:noescape
func RandomBytes(dst *byte, nbytes int) (int, bool)
