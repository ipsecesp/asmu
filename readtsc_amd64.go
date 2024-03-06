package asmu

// ReadTSC returns he current 64-bit value of the processor's time-stamp counter.
// The processor monotonically increments the time-stamp counter
// every clock cycle and resets it to 0 whenever the processor is reset.
//
//go:noescape
func ReadTSC() int64
