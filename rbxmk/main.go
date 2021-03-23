package main

import (
	"os"

	"github.com/anaminus/snek"
)

var Program = snek.NewProgram("rbxmk", os.Args).Usage(Doc("commands/main.md"))

func main() {
	Program.Main()
}
