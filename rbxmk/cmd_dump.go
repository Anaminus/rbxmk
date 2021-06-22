package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/anaminus/rbxmk/dumpformats"
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
	world, err := InitWorld(WorldOpt{
		WorldFlags:   WorldFlags{Debug: false},
		ExcludeRoots: true,
	})
	if err != nil {
		return err
	}
	root := DumpWorld(world)

	// Dump format.
	return format.Func(opt.Stdout, root)
}
