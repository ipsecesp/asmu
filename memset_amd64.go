package asmu

// Memset fills the first bytes of the nbytes of the array
// pointed to by dst with the value of chr.
//
//go:noescape
func MemSet(dst *byte, chr byte, nbytes int)
