package main

import (
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/filter"
	"github.com/anaminus/rbxmk/format"
	"github.com/anaminus/rbxmk/scheme"
	"github.com/jessevdk/go-flags"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxapi/dump"
	"os"
	"strconv"
)

func Fatalf(f string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, f, v...)
	os.Exit(2)
}

func LoadAPI(path string) (api *rbxapi.API) {
	if path != "" {
		file, err := os.Open(path)
		if err != nil {
			Fatalf("failed to open config file: %s", err)
		}
		defer file.Close()
		if api, err = dump.Decode(file); err != nil {
			Fatalf("failed to decode API file: %s", err)
		}
	}
	return
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
	if len(os.Args) < 2 {
		fp.WriteHelp(os.Stderr)
		return
	}

	options := rbxmk.NewOptions()
	scheme.Register(options.Schemes)
	format.Register(options.Formats)
	filter.Register(options.Filters)

	state := NewLuaState(options)

	for _, arg := range args {
		number, err := strconv.ParseFloat(arg, 64)
		switch {
		case err == nil:
			state.state.PushNumber(number)
		case arg == "true":
			state.state.PushBoolean(true)
		case arg == "false":
			state.state.PushBoolean(false)
		case arg == "nil":
			state.state.PushNil()
		default:
			state.state.PushString(arg)
		}
	}

	if flagOptions.File != "" {
		if err := state.DoFile(flagOptions.File, len(args)); err != nil {
			Fatalf("error running file %q: %s", flagOptions.File, err)
		}
	} else {
		if err := state.DoFileHandle(os.Stdin, len(args)); err != nil {
			Fatalf("error running stdin: %s", err)
		}
	}
}
