package main

import (
	"os"
	"path/filepath"
	"strconv"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/formats"
	"github.com/anaminus/rbxmk/library"
	"github.com/anaminus/rbxmk/sources"
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

// ParseLuaValue parses a string into a Lua value. Numbers, bools, and nil are
// parsed into their respective types, and any other value is interpreted as a
// string.
func ParseLuaValue(s string) lua.LValue {
	switch s {
	case "true":
		return lua.LTrue
	case "false":
		return lua.LFalse
	case "nil":
		return lua.LNil
	}
	if number, err := strconv.ParseFloat(s, 64); err == nil {
		return lua.LNumber(number)
	}
	return lua.LString(s)
}

func init() {
	Program.Register(snek.Def{
		Name:      "run",
		Summary:   "Execute a script.",
		Arguments: `[ FILE ] [ ...VALUE ]`,
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

type RunCommand struct {
	Debug bool
	Init  func(rbxmk.State)
}

func (c *RunCommand) SetFlags(flags snek.FlagSet) {
	flags.BoolVar(&c.Debug, "debug", false, "Display stack traces when an error occurs.")
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
		opt.Usage()
		return nil
	}
	file := args[0]
	args = args[1:]

	// Initialize world.
	world := rbxmk.NewWorld(lua.NewState(lua.Options{
		SkipOpenLibs:        true,
		IncludeGoStackTrace: c.Debug,
	}))
	if wd, err := os.Getwd(); err == nil {
		// Working directory is an accessible root.
		world.FS.AddRoot(wd)
	}
	for _, f := range formats.All() {
		world.RegisterFormat(f())
	}
	for _, s := range sources.All() {
		world.RegisterSource(s())
	}
	for _, lib := range library.All() {
		if err := world.Open(lib); err != nil {
			return err
		}
	}

	world.State().SetGlobal("_RBXMK_VERSION", lua.LString(Version))

	// Add script arguments.
	for _, arg := range args {
		world.State().Push(ParseLuaValue(arg))
	}

	if c.Init != nil {
		c.Init(rbxmk.State{World: world, L: world.State()})
	}

	// Run stdin as script.
	if file == "-" {
		return world.DoFileHandle(opt.Stdin, "", len(args))
	}

	// Run file as script.
	filename := shortenPath(filepath.Clean(file))
	return world.DoFile(filename, len(args))
}
