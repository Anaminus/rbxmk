package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/anaminus/cobra"
)

var Program = &cobra.Command{
	Use:   "rbxmk",
	Short: "rbxmk is a tool for managing Roblox projects.",
	// Long:  Doc("Commands"),
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill)
	defer stop()

	Program.SetIn(os.Stdin)
	Program.SetOut(os.Stdout)
	Program.SetErr(os.Stderr)

	DocumentCommands()
	UnresolvedFragments()
	if err := Program.ExecuteContext(ctx); err != nil {
		Program.PrintErrln(err)
	}
}
