package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/anaminus/cobra"
	"github.com/anaminus/rbxmk/dumpformats"
)

func init() {
	dumpfmts := dumpformats.All()
	for i, format := range dumpfmts {
		dumpfmts[i].Description = Doc("DumpFormats/" + format.Name + ":Summary")
	}
	var c DumpCommand
	c.Formats = dumpfmts
	var cmd = &cobra.Command{
		Use:  "dump",
		RunE: c.Run,
	}
	var o sync.Once
	cobra.OnInitialize(func() {
		o.Do(func() {
			// Populate description with dump formats.
			var buf strings.Builder
			dumpfmts.WriteTo(&buf)
			cmd.Long = os.Expand(cmd.Long, func(v string) string {
				switch strings.ToLower(v) {
				case "formats":
					return buf.String()
				default:
					return ""
				}
			})
		})
	})

	Program.AddCommand(cmd)
}

type DumpCommand struct {
	Formats dumpformats.Formats
}

func (c *DumpCommand) Run(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cmd.Usage()
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
	return format.Func(cmd.OutOrStdout(), root)
}
