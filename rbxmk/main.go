package main

import (
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/filter"
	"github.com/anaminus/rbxmk/format"
	"github.com/anaminus/rbxmk/scheme"
	"github.com/jessevdk/go-flags"
	"github.com/yuin/gopher-lua"
	"os"
	"path/filepath"
	"strconv"
)

func Fatalf(f string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, f, v...)
	os.Exit(2)
}

const CommandName = "rbxmk"
const CommandUsage = `[ -h ] [ -f FILE ] [ ARGS... ]

Options after any valid flags will be passed to the script as arguments.
Numbers, bools, and nil are parsed into their respective types in Lua, and any
other values are read as strings.
`

type FlagOptions struct {
	File string `short:"f" long:"file" description:"A file to be executed as a Lua script. If not specified, then the script will be read from stdin instead." long-description:"" value-name:"FILE"`
}

func main() {
	var flagOptions FlagOptions
	fp := flags.NewParser(&flagOptions, flags.Default|flags.PassAfterNonOption)
	fp.Usage = CommandUsage
	args, err := fp.Parse()
	if err != nil {
		if err, ok := err.(*flags.Error); ok && err.Type == flags.ErrHelp {
			fmt.Fprintln(os.Stdout, err)
			return
		}
		Fatalf("flag parser error: %s", err)
	}
	if stat, _ := os.Stdin.Stat(); stat == nil && len(os.Args) < 2 {
		fp.WriteHelp(os.Stderr)
		return
	}

	options := rbxmk.NewOptions()
	if err := options.Schemes.Register(scheme.Schemes.List()...); err != nil {
		Fatalf("%s", err)
	}
	if err := options.Formats.Register(format.Formats.List()...); err != nil {
		Fatalf("%s", err)
	}
	if err := options.Filters.Register(filter.Filters.List()...); err != nil {
		Fatalf("%s", err)
	}

	state := NewLuaState(options)

	for _, arg := range args {
		number, err := strconv.ParseFloat(arg, 64)
		switch {
		case err == nil:
			state.state.Push(lua.LNumber(number))
		case arg == "true":
			state.state.Push(lua.LTrue)
		case arg == "false":
			state.state.Push(lua.LFalse)
		case arg == "nil":
			state.state.Push(lua.LNil)
		default:
			state.state.Push(lua.LString(arg))
		}
	}

	if flagOptions.File != "" {
		filename := shortenPath(filepath.Clean(flagOptions.File))
		if err := state.DoFile(filename, len(args)); err != nil {
			Fatalf("%s", err)
		}
	} else {
		if err := state.DoFileHandle(os.Stdin, len(args)); err != nil {
			Fatalf("%s", err)
		}
	}
}
