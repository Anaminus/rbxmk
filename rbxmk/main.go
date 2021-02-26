package main

import (
	"os"

	"github.com/anaminus/snek"
)

var Program = snek.NewProgram("rbxmk", os.Args).Usage(
	`rbxmk is a tool for managing Roblox projects.

Usage:

	%[1]s <command> [options]

Commands:

%[2]s
Run "%[1]s help <command>" for more information about a command.
`)

func main() {
	Program.Main()
}
