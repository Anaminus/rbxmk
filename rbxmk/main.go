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
)

func Fatalf(f string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, f, v...)
	os.Exit(2)
}

const CommandName = "rbxmk"
const CommandUsage = `[ -h ] [ -a VALUE ] [ -d NAME:VALUE ] [ FILE ]

Receives a file to be executed as a Lua script. If not specified, then the
script will be read from stdin instead.

When specifying an argument or definition, a Lua value is received. Numbers,
bools, and nil are parsed into their respective types in Lua, and any other
value is read as a string. Either option may be given more than once to
provide multiple values.`

type FlagOptions struct {
	Arguments []string          `short:"a" long:"arg" description:"An argument to be passed to the script." long-description:"" value-name:"VALUE"`
	Define    map[string]string `short:"d" long:"define" description:"A variable to be used by the preprocessor." long-description:"" value-name:"NAME:VALUE"`
}

const DefaultHost = `roblox.com`

func main() {
	// Parse flags.
	var flagOptions FlagOptions
	fp := flags.NewParser(&flagOptions, flags.Default|flags.PassAfterNonOption)
	fp.Usage = CommandUsage
	args, err := fp.Parse()
	if err != nil {
		if err, ok := err.(*flags.Error); ok && err.Type == flags.ErrHelp {
			return
		}
		Fatalf("flag parser error: %s", err)
	}

	// Display help if user does nothing.
	if stat, _ := os.Stdin.Stat(); stat == nil && len(os.Args) < 2 {
		fp.WriteHelp(os.Stderr)
		return
	}

	// Initialize top-level options.
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

	// Initialize preprocessor.
	if err := options.Filters.Register(rbxmk.Filter{Name: "preprocess", Func: Preprocess}); err != nil {
		Fatalf("%s", err)
	}
	var envs [PPEnvLen]*lua.LTable
	for i := range envs {
		envs[i] = &lua.LTable{Metatable: lua.LNil}
	}

	// Add preprocessor definitions.
	cmdEnv := envs[PPEnvCommand]
	for k, v := range flagOptions.Define {
		if !CheckStringVar(k) {
			Fatalf("invalid variable name %q", k)
		}
		cmdEnv.RawSetString(k, ParseLuaValue(v, true))
	}
	options.Config["PPEnv"] = envs[:]

	// Initialize Host config.
	options.Config["Host"] = DefaultHost

	// Initialize context.
	ctx := NewLuaContext(options)
	OpenFilteredLibs(ctx.State(), GetFilteredStdLib())
	uctx := ctx.State().NewUserData()
	uctx.Value = ctx
	OpenFilteredLibs(ctx.State(), []LibFilter{
		{MainLibName, OpenMain, nil},
	}, uctx)

	// Add script arguments.
	for _, arg := range flagOptions.Arguments {
		ctx.State().Push(ParseLuaValue(arg, false))
	}

	if len(args) > 0 && args[0] != "" {
		// Run file as script.
		filename := shortenPath(filepath.Clean(args[0]))
		if err := ctx.DoFile(filename, len(flagOptions.Arguments)); err != nil {
			Fatalf("%s", err)
		}
	} else {
		// Run stdin as script.
		if err := ctx.DoFileHandle(os.Stdin, len(flagOptions.Arguments)); err != nil {
			Fatalf("%s", err)
		}
	}
}
