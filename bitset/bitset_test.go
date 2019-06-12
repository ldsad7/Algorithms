package bitset

import (
	"testing"
)

func TestUnit__SetUnsetCount(t *testing.T) {

	bs := New(200)

	// Setting bits
	for i := 0; i < 200; i++ {
		err := bs.Set(i, true)
		if err != nil {
			t.Fatalf("Failed to set a bit: %s", err)
		}
		if isSet, err := bs.Test(i); err != nil || !isSet {
			t.Fatalf("Failed to test a previously set bit")
		}
		if bs.Count() != (i + 1) {
			t.Fatalf("Invalid bit count")
		}
	}

	// Unsetting bits
	for i := 199; i >= 0; i-- {
		bs.Set(i, false)
		if isSet, err := bs.Test(i); err != nil || isSet {
			t.Fatalf("Failed to test a previously unset bit")
		}
		if bs.Count() != i {
			t.Fatalf("Invalid bit count")
		}
	}
}
