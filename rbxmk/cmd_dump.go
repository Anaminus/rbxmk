package main

import (
	"github.com/anaminus/cobra"
	"github.com/anaminus/rbxmk/dumpformats"
)

func init() {
	var Dump = &cobra.Command{
		Use:  "dump",
		Args: cobra.NoArgs,
	}

	for _, format := range dumpformats.All() {
		name := format.Name
		dump := format.Func
		opts := dumpformats.Options{}
		cmd := &cobra.Command{
			Use:   name,
			Short: Doc("Commands/dump/" + name + ":Summary"),
			Long:  Doc("Commands/dump/" + name + ":Description"),
			Args:  cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
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
				return dump(cmd.OutOrStdout(), root, opts)
			},
		}
		flags := cmd.PersistentFlags()
		for flag, value := range format.Options {
			usage := Doc("Commands/dump/" + name + ":Flags/" + flag)
			switch value := value.(type) {
			case string:
				opts[flag] = flags.String(flag, value, usage)
			default:
				panic("unimplemented dump format option type")
			}
		}
		Dump.AddCommand(cmd)
	}
	Program.AddCommand(Dump)
}
