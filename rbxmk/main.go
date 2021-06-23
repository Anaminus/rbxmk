package main

import (
	"os"

	"github.com/anaminus/snek"
)

var Program = snek.NewProgram("rbxmk", os.Args).Usage(Doc("Commands"))

func main() {
	DocumentCommands()
	UnresolvedFragments()
	Program.Main()
}
