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
	FileName   string
	Instances  []*rbxfile.Instance
	Properties map[string]rbxfile.Value
	Values     []rbxfile.Value
	Sources    []*Source
}

func (src *Source) Copy() *Source {
	dst := &Source{
		FileName:   src.FileName,
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
	Source         *Source  // Pre-resolved Source.
	Format         string   // Forced file format. If empty, it is filled in after being guessed.
	ParsedProtocol string   // The protocol parsed from the reference.
}

func (node *InputNode) ResolveReference(opt *Options) (src *Source, err error) {
	ref := node.Reference
	var ext string
	if node.Source != nil {
		src = node.Source
		ext = node.Format
	} else {
		if len(ref) < 1 {
			return nil, errors.New("node requires at least one reference argument")
		}
		schemeName, nextPart := parseScheme(ref[0])
		if schemeName == "" {
			// Assume file:// scheme.
			schemeName = "file"
		}
		scheme, exists := registeredInputSchemes[schemeName]
		if !exists {
			return nil, errors.New("input scheme \"" + schemeName + "\" has not been registered")
		}
		if ext, src, err = scheme.Handler(opt, node, nextPart); err != nil {
			return nil, err
		}
		ref = ref[1:]
	}

	drills, _ := DefaultFormats.InputDrills(ext)
	for _, drill := range drills {
		if src, ref, err = drill(opt, src, ref); err != nil && err != EOD {
			return nil, err
		}
	}
	return src, nil
}

type OutputNode struct {
	Reference      []string // Raw string that refers to a source.
	Source         *Source  // Pre-resolved Source.
	Format         string   // Forced file format. If empty, it is filled in after being guessed.
	ParsedProtocol string   // The protocol parsed from the reference.
}

func (node *OutputNode) ResolveReference(opt *Options, src *Source) (err error) {
	ref := node.Reference
	var ext string
	if node.Source != nil {
		ext = node.Format
		return node.drillOutput(opt, addrSource{src: node.Source}, ref, ext, src)
	}

	if len(ref) < 1 {
		return errors.New("node requires at least one reference argument")
	}
	schemeName, nextPart := parseScheme(ref[0])
	if schemeName == "" {
		// Assume file:// scheme.
		schemeName = "file"
	}
	scheme, exists := registeredOutputSchemes[schemeName]
	if !exists {
		return errors.New("output scheme \"" + schemeName + "\" has not been registered")
	}
	var outsrc *Source
	if ext, outsrc, err = scheme.Handler(opt, node, nextPart); err != nil {
		return err
	}
	ref = ref[1:]
	if err = node.drillOutput(opt, addrSource{src: outsrc}, ref, ext, src); err != nil {
		return err
	}
	return scheme.Finalizer(opt, node, nextPart, ext, outsrc)
}

func (node *OutputNode) drillOutput(opt *Options, addr SourceAddress, ref []string, ext string, src *Source) (err error) {
	drills, exists := DefaultFormats.OutputDrills(ext)
	if !exists {
		return errors.New("invalid format \"" + ext + "\"")
	}
	for _, drill := range drills {
		if addr, ref, err = drill(opt, addr, ref); err != nil && err != EOD {
			return err
		}
	}
	resolver, _ := DefaultFormats.OutputResolver(ext)
	if err = resolver(node.Reference, addr, src); err != nil {
		return err
	}
	return nil
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
