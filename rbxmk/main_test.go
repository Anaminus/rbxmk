package main

import (
	"testing"
)

// TestFragments verifies that fragments have been compiled correctly.
func TestFragments(t *testing.T) {
	defer func() {
		if v := recover(); v != nil {
			t.Fatal(v)
		}
	}()
	DocumentCommands()
	docState.UnresolvedFragments()
}
