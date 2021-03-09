package main

import (
	"fmt"
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/dumpformats"
	"github.com/anaminus/rbxmk/formats"
	"github.com/anaminus/rbxmk/library"
	"github.com/anaminus/snek"
)

func init() {
	dumpfmts := dumpformats.All()
	Program.Register(snek.Def{
		Name:      "dump",
		Summary:   "Dump the script API.",
		Arguments: `FORMAT`,
		Description: `
Dumps the API of the rbxmk Lua environment. The following formats are supported:

%s`,
		New: func() snek.Command {
			return &DumpCommand{Formats: dumpfmts}
		},
		Init: func(def snek.Def) snek.Def {
			// Populate description with dump formats.
			var buf strings.Builder
			dumpfmts.WriteTo(&buf)
			def.Description = fmt.Sprintf(def.Description, buf.String())
			return def
		},
	})
}

type DumpCommand struct {
	Formats dumpformats.Formats
}

func (c *DumpCommand) Run(opt snek.Options) error {
	// Parse flags.
	if err := opt.ParseFlags(); err != nil {
		return err
	}

	args := opt.Args()
	if len(args) == 0 {
		opt.WriteUsageOf(opt.Stderr, opt.Def)
		return nil
	}

	format, ok := c.Formats.Get(args[0])
	if !ok {
		return fmt.Errorf("unknown format %q", args[0])
	}

	// Populate dump.
	world := rbxmk.NewWorld(lua.NewState(lua.Options{
		SkipOpenLibs: true,
	}))
	state := rbxmk.State{World: world, L: world.State()}
	var root dump.Root
	for _, f := range formats.All() {
		world.RegisterFormat(f())
	}
	for _, l := range library.All() {
		if err := world.Open(l); err != nil {
			return err
		}
		if l.Dump == nil {
			continue
		}
		lib := l.Dump(state)
		lib.ImportedAs = l.Name
		if lib.Name == "" {
			lib.Name = l.Name
		}
		root.Libraries = append(root.Libraries, lib)
	}
	root.Libraries = append(root.Libraries, dump.Library{
		Name:       "executable",
		ImportedAs: "",
		Struct: dump.Struct{Fields: dump.Fields{
			"_RBXMK_VERSION": dump.Property{ValueType: dt.Prim("string")},
		}},
	})

	// Dump format.
	return format.Func(opt.Stdout, root)
}
