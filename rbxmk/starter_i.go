//go:build interactive_commands

package main

import (
	"context"

	"github.com/anaminus/cobra"
	"github.com/inconshreveable/mousetrap"
)

func Starter() func(context.Context) error {
	// If program was started by Windows Explorer, run in interactive mode.
	// Cobra uses this to display a message, so it must be disabled first.
	cobra.MousetrapHelpText = ""
	if mousetrap.StartedByExplorer() {
		return InteractiveMode
	}
	return Program.ExecuteContext
}
