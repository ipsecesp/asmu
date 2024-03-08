package asmu

import (
	"testing"
)

func TestReadTSC(t *testing.T) {
	c1 := ReadTSC()
	if c1 < 1 {
		t.Fatalf("invalid first value: %d", c1)
	}
	if c2 := ReadTSC(); c1 >= c2 {
		t.Fatalf("first value must be less than the second value: %d, %d", c1, c2)
	}
}
