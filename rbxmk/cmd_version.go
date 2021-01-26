package main

import (
	"github.com/anaminus/but"
	"github.com/anaminus/rbxmk/rbxmk/cmds"
)

func init() {
	Commands.Register(cmds.Command{
		Name:    "version",
		Summary: "Display the version.",
		Usage:   `rbxmk version`,
		Description: `
Displays the current version of rbxmk.`,
		Func: VersionCommand,
	})
}

// VersionCommand executes the version command.
func VersionCommand(flags cmds.Flags) {
	but.IfFatal(flags.Parse(), "parse flags")
	but.Log(Version)
}
