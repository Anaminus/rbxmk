package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxapi/dump"
	"github.com/robloxapi/rbxfile"
	"os"
)

type Source struct {
	Instances  []*rbxfile.Instance
	Properties map[string]rbxfile.Value
	Values     []rbxfile.Value
}

func (src *Source) Copy() *Source {
	dst := &Source{
		Instances:  make([]*rbxfile.Instance, len(src.Instances)),
		Properties: make(map[string]rbxfile.Value, len(src.Properties)),
		Values:     make([]rbxfile.Value, len(src.Values)),
	}
	for i, inst := range src.Instances {
		dst.Instances[i] = inst.Clone()
	}
	for name, value := range src.Properties {
		dst.Properties[name] = value.Copy()
	}
	for i, value := range src.Values {
		dst.Values[i] = value.Copy()
	}
	return dst
}

// parseScheme separates a string into scheme and path parts.
func parseScheme(s string) (scheme, path string) {
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z':
		case '0' <= c && c <= '9' || c == '+' || c == '-' || c == '.':
			if i == 0 {
				return "", s
			}
		case c == ':':
			if i > 0 && i+2 < len(s) && s[i+1] == '/' && s[i+2] == '/' {
				return s[:i], s[i+3:]
			}
			return "", s
		default:
			return "", s
		}
	}
	return "", s
}

type InputNode struct {
	Reference      []string // Raw strings that refer to a source.
	Format         string   // Forced file format. If empty, it is filled in after being guessed.
	ParsedProtocol string   // The protocol parsed from the reference.
}

func (node *InputNode) ResolveReference(opt *Options) (src *Source, err error) {
	if len(node.Reference) < 1 {
		return nil, errors.New("node requires at least one reference argument")
	}
	scheme, nextPart := parseScheme(node.Reference[0])
	handler, exists := registeredInputSchemes[scheme]
	if !exists {
		// Assume file:// scheme.
		handler = registeredInputSchemes["file"]
	}
	if handler == nil {
		panic("\"file\" input scheme has not been registered")
	}
	return handler(opt, node, nextPart)
}

type OutputNode struct {
	Reference      []string // Raw string that refers to a source.
	Format         string   // Forced file format. If empty, it is filled in after being guessed.
	ParsedProtocol string   // The protocol parsed from the reference.
}

func (node *OutputNode) ResolveReference(opt *Options, src *Source) (err error) {
	if len(node.Reference) < 1 {
		return errors.New("node requires at least one reference argument")
	}
	scheme, nextPart := parseScheme(node.Reference[0])
	handler, exists := registeredOutputSchemes[scheme]
	if !exists {
		// Assume file:// scheme.
		handler = registeredOutputSchemes["file"]
	}
	if handler == nil {
		panic("\"file\" output scheme has not been registered")
	}
	return handler(opt, node, nextPart, src)
}

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

type Options struct {
	API *rbxapi.API
}

const CommandName = "rbxmk"
const CommandUsage = `[OPTIONS...] [- | SCRIPT]

rbxmk uses Lua to perform actions. Lua scripts can be executed from several
places.

- From files: Specifying one or more "-f" or "--file" options will execute Lua
  from the given files. Files are executed in the order they are specified.
- From SCRIPT: The first non-flag option signals the rest of the command to be
  read as a Lua script.
- From stdin: Specifying "-" will begin parsing Lua from stdin.
`

type FlagOptions struct {
	Files []string `short:"f" long:"file" description:"A file to be executed as a Lua script." long-description:"" value-name:"FILE"`
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

	options := &Options{}
	state := NewLuaState(options)

	for _, f := range flagOptions.Files {
		if err := state.DoFile(f); err != nil {
			Fatalf("error running file %q: %s", f, err)
		}
	}

	if len(args) == 1 && args[0] == "-" {
		if err := state.DoFileHandle(os.Stdin); err != nil {
			Fatalf("error running stdin: %s", err)
		}
		return
	}

	if len(args) > 1 {
		Fatalf("single SCRIPT argument expected (is your script quoted?)")
	}

	if err := state.DoString(args[0], "command argument"); err != nil {
		Fatalf("error running command argument: %s", err)
	}
}
