package main

import (
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/filter"
	"github.com/anaminus/rbxmk/format"
	"github.com/anaminus/rbxmk/luautil"
	"github.com/anaminus/rbxmk/scheme"
	"github.com/jessevdk/go-flags"
	"github.com/yuin/gopher-lua"
	"os"
	"path/filepath"
)

func Fatalf(f string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, f, v...)
	os.Exit(2)
}

const CommandName = "rbxmk"
const CommandUsage = `[ -h ] [ -f FILE ] [ -d NAME:VALUE ] [ ARGS... ]

Options after any valid flags will be passed to the script as arguments.
Numbers, bools, and nil are parsed into their respective types in Lua, and any
other values are read as strings.
`

type FlagOptions struct {
	File   string            `short:"f" long:"file" description:"A file to be executed as a Lua script. If not specified, then the script will be read from stdin instead." long-description:"" value-name:"FILE"`
	Define map[string]string `short:"d" long:"define" description:"Defines a variable to be used by the preprocessor. Can be specified multiple times for multiple variables. The value may be a Lua bool, number, string, or nil." long-description:"" value-name:"NAME:VALUE"`
}

// Order of preprocessor variable environments.
const (
	ppEnvScript  = iota // Defined via script (rbxmk.configure).
	ppEnvCommand        // Defined via --define option.
	ppEnvLen            // Number of environments.
)

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

	options.Config.PreprocessorEnvs = make([]*lua.LTable, ppEnvLen)
	for i := range options.Config.PreprocessorEnvs {
		options.Config.PreprocessorEnvs[i] = &lua.LTable{Metatable: lua.LNil}
	}

	for k, v := range flagOptions.Define {
		if !luautil.CheckStringVar(k) {
			Fatalf("invalid variable name %q", k)
		}
		options.Config.PreprocessorEnvs[ppEnvCommand].RawSetString(k, luautil.ParseLuaValue(v))
	}

	ctx := luautil.NewLuaContext(options)
	luautil.OpenFilteredLibs(ctx.State(), luautil.GetFilteredStdLib())
	uctx := ctx.State().NewUserData()
	uctx.Value = ctx
	luautil.OpenFilteredLibs(ctx.State(), []luautil.LibFilter{
		{MainLibName, OpenMain, nil},
	}, uctx)

	for _, arg := range args {
		ctx.State().Push(luautil.ParseLuaValue(arg))
	}

	if flagOptions.File != "" {
		filename := shortenPath(filepath.Clean(flagOptions.File))
		if err := ctx.DoFile(filename, len(args)); err != nil {
			Fatalf("%s", err)
		}
	} else {
		if err := ctx.DoFileHandle(os.Stdin, len(args)); err != nil {
			Fatalf("%s", err)
		}
	}
}
