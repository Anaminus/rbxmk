package main

import (
	"os"
	"sort"
	"strconv"
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/pflag"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/enums"
	"github.com/anaminus/rbxmk/formats"
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
	Libraries     []string
}

func (f *WorldFlags) SetFlags(flags *pflag.FlagSet) {
	flags.StringArrayVar(&f.IncludedRoots, "include-root", nil, DocFlag("Flags/world:Flags/include-root"))
	flags.StringArrayVar(&f.Libraries, "libraries", nil, DocFlag("Flags/world:Flags/libraries"))
	flags.BoolVar(&f.InsecurePaths, "allow-insecure-paths", false, DocFlag("Flags/world:Flags/allow-insecure-paths"))
	flags.BoolVar(&f.Debug, "debug", false, DocFlag("Flags/world:Flags/debug"))
}

// WorldOpt are options to InitWorld.
type WorldOpt struct {
	WorldFlags
	ExcludeRoots     bool
	ExcludeFormats   bool
	ExcludeEnums     bool
	IncludeLibraries rbxmk.Libraries
	ExcludeProgram   bool
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
	var libraries rbxmk.Libraries
	if !opt.ExcludeProgram {
		libraries = append(libraries, ProgramLibrary)
	}
	libraries = append(libraries, opt.IncludeLibraries...)
	sort.Sort(libraries)
	included := make(map[string]bool, len(libraries))
	for _, lib := range libraries {
		included[lib.Name] = true
	}
	for _, list := range opt.Libraries {
		for _, name := range strings.Split(list, ",") {
			name = strings.TrimSpace(name)
			if name == "" {
				continue
			}
			include := true
			switch name[0] {
			case '-':
				include = false
				name = name[1:]
			case '+':
				include = true
				name = name[1:]
			}
			if name == "*" {
				for lib := range included {
					included[lib] = include
				}
			} else if _, ok := included[name]; ok {
				included[name] = include
			}
		}
	}
	// Load negative-priority libraries before formats.
	for _, lib := range libraries {
		if lib.Priority >= 0 {
			break
		}
		if !included[lib.Name] {
			continue
		}
		if err := world.Open(lib); err != nil {
			return nil, err
		}
	}
	if !opt.ExcludeFormats {
		for _, f := range formats.All() {
			world.RegisterFormat(f())
		}
	}
	if !opt.ExcludeEnums {
		world.RegisterEnums(enums.All()...)
	}
	for _, lib := range libraries {
		if lib.Priority < 0 {
			// Already loaded negative-priority libraries.
			continue
		}
		if !included[lib.Name] {
			continue
		}
		if err := world.Open(lib); err != nil {
			return nil, err
		}
	}
	for _, arg := range opt.Args {
		world.LuaState().Push(ParseLuaValue(arg))
	}
	return world, nil
}

func dumpTypes(dst dump.TypeDefs, src []func() rbxmk.Reflector) {
	for _, t := range src {
		r := t()
		if _, ok := dst[r.Name]; ok {
			continue
		}
		dst[r.Name] = r.DumpAll()
		dumpTypes(dst, r.Types)
	}
}

func DumpWorld(world *rbxmk.World) dump.Root {
	state := world.State()
	root := dump.Root{
		Formats: dump.Formats{},
		Types:   dump.TypeDefs{},
	}
	for _, format := range world.Formats() {
		root.Formats[format.Name] = format.Dump()
	}
	for _, l := range world.Libraries() {
		if l.Dump == nil {
			continue
		}
		lib := l.Dump(state)
		if lib.Name == "" {
			lib.Name = l.Name
		}
		if lib.ImportedAs == "" {
			lib.ImportedAs = l.ImportedAs
		}
		lib.Priority = l.Priority
		if l.Types != nil {
			dumpTypes(root.Types, l.Types)
		}
		root.Libraries = append(root.Libraries, lib)
	}
	root.Fragments = DocFragments()
	root.Description = "Libraries"
	return root
}

var ProgramLibrary = rbxmk.Library{
	Name:       "program",
	ImportedAs: "",
	Priority:   0,
	Open: func(s rbxmk.State) *lua.LTable {
		lib := s.L.CreateTable(0, 1)
		lib.RawSetString("_RBXMK_VERSION", lua.LString(VersionString()))
		return lib
	},
	Dump: func(s rbxmk.State) dump.Library {
		return dump.Library{
			Name:       "program",
			ImportedAs: "",
			Struct: dump.Struct{
				Fields: dump.Fields{
					"_RBXMK_VERSION": dump.Property{
						ValueType:   dt.Prim("string"),
						ReadOnly:    true,
						Summary:     "Libraries/program:Fields/_RBXMK_VERSION/Summary",
						Description: "Libraries/program:Fields/_RBXMK_VERSION/Description",
					},
				},
				Summary:     "Libraries/program:Summary",
				Description: "Libraries/program:Description",
			},
		}
	},
}
