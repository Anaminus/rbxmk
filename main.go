package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxapi/dump"
	"github.com/robloxapi/rbxfile"
	"os"
	"strconv"
)

type Source struct {
	Instances  []*rbxfile.Instance
	Properties map[string]rbxfile.Value
	Values     []rbxfile.Value
	Sources    []*Source
}

func (src *Source) Copy() *Source {
	dst := &Source{
		Instances:  make([]*rbxfile.Instance, len(src.Instances)),
		Properties: make(map[string]rbxfile.Value, len(src.Properties)),
		Values:     make([]rbxfile.Value, len(src.Values)),
		Sources:    make([]*Source, len(src.Sources)),
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
	for i, s := range src.Sources {
		dst.Sources[i] = s.Copy()
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

	options := &Options{}
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
