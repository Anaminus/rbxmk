package main

import (
	"fmt"
	"os"
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dumpformats"
	"github.com/anaminus/rbxmk/formats"
	"github.com/anaminus/rbxmk/library"
	"github.com/anaminus/snek"
)

func init() {
	dumpfmts := dumpformats.All()
	for i, format := range dumpfmts {
		dumpfmts[i].Description = Doc("DumpFormats/" + format.Name + ":Summary")
	}
	Program.Register(snek.Def{
		Name: "dump",
		New: func() snek.Command {
			return &DumpCommand{Formats: dumpfmts}
		},
		Init: func(def snek.Def) snek.Def {
			// Populate description with dump formats.
			var buf strings.Builder
			dumpfmts.WriteTo(&buf)
			def.Description = os.Expand(def.Description, func(v string) string {
				switch strings.ToLower(v) {
				case "formats":
					return buf.String()
				default:
					return ""
				}
			})
			return def
		},
	})
}

type DumpCommand struct {
	Formats dumpformats.Formats
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
	state := world.State()
	root := dump.Root{
		Formats: dump.Formats{},
		Types:   dump.TypeDefs{},
	}
	for _, f := range formats.All() {
		format := f()
		world.RegisterFormat(format)
		root.Formats[format.Name] = format.Dump()
	}
	libraries := library.All()
	libraries = append(libraries, ProgramLibrary)
	for _, l := range libraries {
		if err := world.Open(l); err != nil {
			return err
		}
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
		if l.Types != nil {
			dumpTypes(root.Types, l.Types)
		}
		root.Libraries = append(root.Libraries, lib)
	}
	root.Fragments = DocFragments()
	root.Description = "Libraries"

	// Dump format.
	return format.Func(opt.Stdout, root)
}
