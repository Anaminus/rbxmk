package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/anaminus/cobra"
	"github.com/anaminus/pflag"
	"github.com/anaminus/rbxmk"
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

func initRunCommand(c *RunCommand) *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "run",
		RunE: c.Run,
	}
	c.SetFlags(cmd.PersistentFlags())
	cmd.FParseErrWhitelist.UnknownFlags = true
	cmd.Flags().KeepUnknownFlags = true
	return cmd
}

func init() {
	var c RunCommand
	cmd := initRunCommand(&c)
	Program.AddCommand(cmd)
}

type RunCommand struct {
	WorldFlags
	DescFlags
	Init func(c *RunCommand, s rbxmk.State)
}

func (c *RunCommand) SetFlags(flags *pflag.FlagSet) {
	c.WorldFlags.SetFlags(flags)
	c.DescFlags.SetFlags(flags)
}

// Run is the entrypoint to the command for running scripts. init runs after the
// World envrionment is fully initialized and arguments have been pushed, and
// before the script runs.
func (c *RunCommand) Run(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cmd.Usage()
	}
	file := args[0]
	args = args[1:]

	// Initialize world.
	world, err := InitWorld(WorldOpt{
		WorldFlags: c.WorldFlags,
		Args:       args,
	})
	if err != nil {
		return err
	}
	injectSSLKeyLogFile(world, cmd.ErrOrStderr())
	if c.Init != nil {
		c.Init(c, world.State())
	}

	// Initialize global descriptor.
	world.Desc, err = c.DescFlags.Resolve(world.Client)
	if err != nil {
		return err
	}
	world.SetEnumGlobal()

	// Run stdin as script.
	if file == "-" {
		stdin := cmd.InOrStdin()
		if stdin == nil {
			return fmt.Errorf("no file handle")
		}
		if f, ok := stdin.(fs.File); ok {
			return world.DoFileHandle(f, "", len(args))
		} else {
			b, err := io.ReadAll(stdin)
			if err != nil {
				return fmt.Errorf("read stdin: %w", err)
			}
			return world.DoString(string(b), "", len(args))
		}
	}

	// Run file as script.
	filename := shortenPath(filepath.Clean(file))
	return world.DoFile(filename, len(args))
}
