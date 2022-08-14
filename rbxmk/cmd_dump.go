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
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dumpformats"
	"github.com/anaminus/rbxmk/formats"
	"github.com/anaminus/rbxmk/library"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

// GenerateDump produces a dump.Root for the given world configuration.
func GenerateDump(opt WorldOpt) (root dump.Root, err error) {
	env := &dump.EnvRef{}
	opt.EventHook = func(e rbxmk.EnvEvent) {
		node := env
		for _, name := range e.EnvPath {
			sub, ok := node.Fields[name]
			if !ok {
				sub = &dump.EnvRef{}
				if node.Fields == nil {
					node.Fields = map[string]*dump.EnvRef{}
				}
				node.Fields[name] = sub
			}
			node = sub
		}
		node.Path = e.DumpPath
	}
	world, err := InitWorld(opt)
	if err != nil {
		return root, err
	}
	root = DumpWorld(world)
	root.Environment = env
	return root, nil
}

func init() {
	var Dump = Register.NewCommand(dump.Command{
		Arguments:   "Commands/dump:Arguments",
		Summary:     "Commands/dump:Summary",
		Description: "Commands/dump:Description",
	}, &cobra.Command{
		Use:  "dump",
		Args: cobra.NoArgs,
	})

	for _, format := range dumpformats.All() {
		name := format.Name
		fn := format.Func
		opts := dumpformats.Options{}
		cmd := Register.NewCommand(dump.Command{
			Summary:     "Commands/dump/" + name + ":Summary",
			Description: "Commands/dump/" + name + ":Description",
		}, &cobra.Command{
			Use:  name,
			Args: cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				root, err := GenerateDump(WorldOpt{
					WorldFlags:       WorldFlags{Debug: false},
					IncludeLibraries: library.All(),
					ExcludeRoots:     true,
				})
				if err != nil {
					return err
				}
				return fn(cmd.OutOrStdout(), root, opts)
			},
		})
		flags := cmd.Flags()
		for flag, value := range format.Options {
			switch value := value.(type) {
			case string:
				opts[flag] = flags.String(flag, value, "")
			default:
				panic("unimplemented dump format option type")
			}
			Register.NewFlag(dump.Flag{Description: "Commands/dump/" + name + ":Flags/" + flag}, flags, flag)
		}
		Dump.AddCommand(cmd)
	}

	var plugin = Register.NewCommand(dump.Command{
		Arguments:   "Commands/dump/plugin:Arguments",
		Summary:     "Commands/dump/plugin:Summary",
		Description: "Commands/dump/plugin:Description",
	}, &cobra.Command{
		Use:  "plugin",
		Args: cobra.ExactArgs(1),
		RunE: runDumpPluginCommand,
	})
	Dump.AddCommand(plugin)

	Program.AddCommand(Dump)
}

func runDumpPluginCommand(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cmd.Usage()
	}
	file := args[0]
	args = args[1:]

	// Generate dump.
	root, err := GenerateDump(WorldOpt{
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
