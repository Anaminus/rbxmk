package cmds

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"sort"
	"strings"
)

// FileReader represents a file that can be read from.
type ReadFile = fs.File

// FileWriter represents a file that can be written to.
type WriteFile interface {
	fs.File
	io.Writer
}

// Flags bundles a FlagSet with arguments and file descriptors.
type Flags struct {
	*flag.FlagSet
	Stdin  ReadFile
	Stdout WriteFile
	Stderr WriteFile

	args []string
}

// NewFlags returns an initialized Flags. name is the name of the program. args
// are the arguments to be parsed with the embedded FlagSet. The Flags' file
// descriptors are set to the standard file descriptors.
func NewFlags(name string, args []string) Flags {
	return Flags{
		FlagSet: flag.NewFlagSet(name, flag.ExitOnError),
		args:    args,
		Stdin:   os.Stdin,
		Stdout:  os.Stdout,
		Stderr:  os.Stderr,
	}
}

// ShiftArg attempts to return the first argument of Flags. If successful, the
// first argument is removed, shifting down the remaining arguments.
func (f *Flags) ShiftArg() (arg string, ok bool) {
	if len(f.args) == 0 {
		return "", false
	}
	arg = f.args[0]
	f.args = f.args[1:]
	return arg, true
}

// formatDesc formats a command description for readability.
func formatDesc(s string) string {
	s = strings.TrimSpace(s)
	//TODO: Wrap to 80 characters.
	return s
}

// UsageOf returns a Usage function constructed from cmd.
func (f *Flags) UsageOf(cmd Command) func() {
	var usage string
	var desc string
	if cmd.Usage != "" {
		usage = cmd.Usage
	}
	if cmd.Description != "" {
		desc = formatDesc(cmd.Description)
	}
	return func() {
		if usage != "" {
			fmt.Fprintf(f.Output(), "Usage: %s\n\n", usage)
		}
		if desc != "" {
			fmt.Fprintf(f.Output(), "%s", desc)
		}
		f.PrintDefaults()
	}
}

// Parse parses the Flags' arguments with the FlagSet.
func (f Flags) Parse() error {
	return f.FlagSet.Parse(f.args)
}

// Command describes a subcommand to be run within the program.
type Command struct {
	// Name is the name of the command.
	Name string

	// Summary is a short description of the command.
	Summary string

	// Usage describes the structure of the command.
	Usage string

	// Description is a detailed description of the command.
	Description string

	// Func is the function that runs when the command is invoked.
	Func func(Flags)
}

// Commands maps a name to a Command.
type Commands struct {
	// Name is the name of the program.
	Name string

	m map[string]Command
}

// NewCommands returns an initialized commands.
func NewCommands(name string) Commands {
	return Commands{
		Name: name,
		m:    map[string]Command{},
	}
}

// Register registers cmd as cmd.Name.
func (c Commands) Register(cmd Command) {
	c.m[cmd.Name] = cmd
}

// Has returns whether name is a registered command.
func (c Commands) Has(name string) bool {
	_, ok := c.m[name]
	return ok
}

// Get returns the Command mapped to the given name.
func (c Commands) Get(name string) Command {
	return c.m[name]
}

// List returns a list of commands, sorted by name.
func (c Commands) List() []Command {
	list := make([]Command, 0, len(c.m))
	for _, cmd := range c.m {
		list = append(list, cmd)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})
	return list
}

// Do executes the Command mapped to the given name. Does nothing if the name is
// not defined.
func (c Commands) Do(name string, args []string) {
	cmd := c.m[name]
	if cmd.Func == nil {
		return
	}
	flags := NewFlags(c.Name, args)
	flags.Usage = flags.UsageOf(cmd)
	cmd.Func(flags)
}
