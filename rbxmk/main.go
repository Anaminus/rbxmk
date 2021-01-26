package main

import (
	"os"

	"github.com/anaminus/but"
	"github.com/anaminus/rbxmk/rbxmk/cmds"
)

// GlobalOptions contains flag options that apply to all commands.
type GlobalOptions struct {
}

// SetFlags registers global options for the given flags.
func (opt GlobalOptions) SetFlags(flags cmds.Flags) {
}

// Commands contains the subcommands for the program.
var Commands = cmds.NewCommands("")

func main() {
	Commands.Name = os.Args[0]
	if len(os.Args) <= 1 {
		Commands.Do("help", nil)
		return
	}
	if Commands.Has(os.Args[1]) {
		Commands.Do(os.Args[1], os.Args[2:])
		return
	}
	but.Logf("unknown command %q\n", os.Args[1])
	Commands.Do("help", nil)
}
