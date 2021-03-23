package main

import (
	"os"
	"strconv"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/formats"
	"github.com/anaminus/rbxmk/library"
	"github.com/anaminus/snek"
)

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

// WorldFlags are common command flags involved in initializing a World.
type WorldFlags struct {
	IncludedRoots []string
	InsecurePaths bool
	Debug         bool
}

func (f *WorldFlags) SetFlags(flags snek.FlagSet) {
	flags.Var((*repeatedString)(&f.IncludedRoots), "include-root", Doc("commands/world_flags.md/include-root"))
	flags.BoolVar(&f.InsecurePaths, "allow-insecure-paths", false, Doc("commands/world_flags.md/allow-insecure-paths"))
	flags.BoolVar(&f.Debug, "debug", false, Doc("commands/world_flags.md/debug"))
}

// WorldOpt are options to InitWorld.
type WorldOpt struct {
	WorldFlags
	ExcludeRoots     bool
	ExcludeFormats   bool
	ExcludeLibraries bool
	ExcludeVersion   bool
	Args             []string
}

// InitWorld initializes an rbxmk.World with a common structure.
func InitWorld(opt WorldOpt) (world *rbxmk.World, err error) {
	world = rbxmk.NewWorld(lua.NewState(lua.Options{
		SkipOpenLibs:        true,
		IncludeGoStackTrace: opt.Debug,
	}))
	if !opt.ExcludeRoots {
		if opt.InsecurePaths {
			world.FS.SetSecured(false)
		}
		if wd, err := os.Getwd(); err == nil {
			// Working directory is an accessible root.
			world.FS.AddRoot(wd)
		}
		for _, root := range opt.IncludedRoots {
			world.FS.AddRoot(root)
		}
	}
	if !opt.ExcludeFormats {
		for _, f := range formats.All() {
			world.RegisterFormat(f())
		}
	}
	if !opt.ExcludeLibraries {
		for _, lib := range library.All() {
			if err := world.Open(lib); err != nil {
				return nil, err
			}
		}
	}
	if !opt.ExcludeVersion {
		world.State().SetGlobal("_RBXMK_VERSION", lua.LString(VersionString()))
	}
	for _, arg := range opt.Args {
		world.State().Push(ParseLuaValue(arg))
	}
	return world, nil
}
