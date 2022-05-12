package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/anaminus/cobra"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dumpformats"
	"github.com/anaminus/rbxmk/formats"
	"github.com/anaminus/rbxmk/library"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
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
					WorldFlags:       WorldFlags{Debug: false},
					IncludeLibraries: library.All(),
					ExcludeRoots:     true,
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

	var plugin = &cobra.Command{
		Use:   "plugin " + Doc("Commands/dump/plugin:Arguments"),
		Short: Doc("Commands/dump/plugin:Summary"),
		Long:  Doc("Commands/dump/plugin:Description"),
		Args:  cobra.ExactArgs(1),
		RunE:  runDumpPluginCommand,
	}
	Dump.AddCommand(plugin)

	Program.AddCommand(Dump)
}

func runDumpPluginCommand(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cmd.Usage()
	}
	file := args[0]
	args = args[1:]

	// Initialize world for dumping.
	dumpWorld, err := InitWorld(WorldOpt{
		ExcludeRoots:     true,
		IncludeLibraries: library.All(),
	})
	if err != nil {
		return err
	}

	// Initialize world for plugin environment.
	world, err := InitWorld(WorldOpt{
		ExcludeRoots: true,
		IncludeLibraries: rbxmk.Libraries{
			library.Base,
			library.Math,
			library.String,
			library.Table,
		},
	})
	if err != nil {
		return err
	}

	// Push dump as table.
	root := DumpWorld(dumpWorld)
	var buf bytes.Buffer
	je := json.NewEncoder(&buf)
	je.SetEscapeHTML(false)
	je.SetIndent("", "")
	je.Encode(root)
	table, _ := formats.JSON().Decode(rtypes.Global{}, nil, &buf)
	world.RegisterReflector(reflect.Dictionary())
	world.State().Push(table)

	// Push write function.
	world.RegisterReflector(reflect.Tuple())
	var builder strings.Builder
	world.LuaState().Push(world.Context().WrapFunc(func(s rbxmk.State) int {
		for _, v := range s.PullTuple(1) {
			switch v := v.(type) {
			case types.Bool:
				builder.WriteString(strconv.FormatBool(bool(v)))
			case types.Numberlike:
				builder.WriteString(strconv.FormatFloat(v.Numberlike(), 'g', -1, 64))
			case types.Intlike:
				builder.WriteString(strconv.FormatInt(v.Intlike(), 10))
			case types.Stringlike:
				builder.WriteString(v.Stringlike())
			}
		}
		return 0
	}))

	nargs := 2

	if file == "-" {
		// Run stdin as script.
		stdin := cmd.InOrStdin()
		if stdin == nil {
			err = fmt.Errorf("no file handle")
		}
		if f, ok := stdin.(fs.File); ok {
			err = world.DoFileHandle(f, "", nargs)
		} else {
			b, err := io.ReadAll(stdin)
			if err != nil {
				err = fmt.Errorf("read stdin: %w", err)
			}
			err = world.DoString(string(b), "", nargs)
		}
	} else {
		// Run file as script.
		filename := shortenPath(filepath.Clean(file))
		err = world.DoFile(filename, nargs)
	}
	if err != nil {
		return err
	}

	cmd.Print(builder.String())

	return nil
}
