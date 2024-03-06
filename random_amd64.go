package asmu

//go:noescape
func RandomBytes(dst *byte, nbytes int) (int, bool)
