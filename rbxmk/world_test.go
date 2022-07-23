package main

import (
	"testing"
)

// TestCommandDefs verifies that commands and flags attached to Program all have
// associated dump definitions.
func TestCommandDefs(t *testing.T) {
	defer func() {
		if v := recover(); v != nil {
			t.Fatal(v)
		}
	}()

	walkCommands(*Register.Command[Program], Register, Program)
}
