package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/snek"
)

// shortenPath transforms the given path so that it is relative to the working
// directory. Returns the original path if that fails.
func shortenPath(filename string) string {
	if wd, err := os.Getwd(); err == nil {
		if abs, err := filepath.Abs(filename); err == nil {
			if r, err := filepath.Rel(wd, abs); err == nil {
				filename = r
			}
		}
	}
	return filename
}

func init() {
	Program.Register(snek.Def{
		Name:      "run",
		Summary:   "Execute a script.",
		Arguments: `[ FLAGS ] FILE [ VALUE... ]`,
		Description: `
Receives a file to be executed as a Lua script. If "-" is given, then the script
will be read from stdin instead.

Remaining arguments are Lua values to be passed to the file. Numbers, bools, and
nil are parsed into their respective types in Lua, and any other value is
interpreted as a string. Within the script, these arguments can be received from
the ... operator.`,
		New: func() snek.Command { return &RunCommand{} },
	})
}

// repeatedString is a string flag that can be specified multiple times.
type repeatedString []string

func (s repeatedString) String() string {
	return strings.Join(s, ",")
}

func (s *repeatedString) Set(v string) error {
	*s = append(*s, v)
	return nil
}

type RunCommand struct {
	WorldFlags
	Init func(rbxmk.State)
}

func (c *RunCommand) SetFlags(flags snek.FlagSet) {
	c.WorldFlags.SetFlags(flags)
}

// Run is the entrypoint to the command for running scripts. init runs after the
// World envrionment is fully initialized and arguments have been pushed, and
// before the script runs.
func (c *RunCommand) Run(opt snek.Options) error {
	// Parse flags.
	if err := opt.ParseFlags(); err != nil {
		return err
	}
	args := opt.Args()
	if len(args) == 0 {
		opt.WriteUsageOf(opt.Stderr, opt.Def)
		return nil
	}
	file := args[0]
	args = args[1:]

	// Initialize world.
	world, err := InitWorld(WorldOpt{
		WorldFlags: c.WorldFlags,
		Args:       args,
	})
	if err != nil {
		return err
	}
	if c.Init != nil {
		c.Init(rbxmk.State{World: world, L: world.State()})
	}

	// Run stdin as script.
	if file == "-" {
		if opt.Stdin == nil {
			return fmt.Errorf("no file handle")
		}
		return world.DoFileHandle(opt.Stdin, "", len(args))
	}

	// Run file as script.
	filename := shortenPath(filepath.Clean(file))
	return world.DoFile(filename, len(args))
}
