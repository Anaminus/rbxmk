package main

import (
	"bytes"
	"fmt"

	"github.com/anaminus/but"
	"github.com/anaminus/rbxmk/rbxmk/cmds"
)

const globalUsage = `rbxmk is a tool for managing Roblox projects.

Usage:

	rbxmk <command> [options]

Commands:

%s
Run "rbxmk help <command>" for more information about a command.
`

func init() {
	Commands.Register(cmds.Command{
		Name:    "help",
		Summary: "Display help.",
		Usage:   `rbxmk help [command]`,
		Description: `
Displays help for a command, or general help if no command is given.`,
		Func: HelpCommand,
	})
}

// HelpCommand executes the help command.
func HelpCommand(flags cmds.Flags) {
	name, ok := flags.ShiftArg()
	but.IfFatal(flags.Parse(), "parse flags")
	if ok {
		if Commands.Has(name) {
			cmd := Commands.Get(name)
			flags.UsageOf(cmd)()
			return
		}
		but.Logf("unknown command %q\n\n", name)
	}
	list := Commands.List()
	width := 0
	for _, cmd := range list {
		if len(cmd.Name) > width {
			width = len(cmd.Name)
		}
	}
	var buf bytes.Buffer
	for _, cmd := range list {
		fmt.Fprintf(&buf, "\t%-*s    %s\n", width, cmd.Name, cmd.Summary)
	}
	but.Logf(globalUsage, buf.String())
}
