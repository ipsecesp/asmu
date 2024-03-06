package asmu

//go:noescape
func MemCopy(dst, src *byte, nbytes int)

//go:noescape
func MemCopyWrap(dst, src *byte, nbytes int)
