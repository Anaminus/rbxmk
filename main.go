package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxapi/dump"
	"github.com/robloxapi/rbxfile"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Source struct {
	Instances  []*rbxfile.Instance
	Properties map[string]rbxfile.Value
	Values     []rbxfile.Value
}

// Receives a Node and a reference string. ref is n.Reference after it has been
// parsed by the protocol detector, and excludes the scheme ("scheme://")
// portion of the string, if it was given.
type InputSchemeHandler func(opt *Options, node *InputNode, ref string) (src *Source, err error)

var registeredInputSchemes = map[string]InputSchemeHandler{}

func RegisterInputScheme(name string, handler InputSchemeHandler) {
	if handler == nil {
		panic("cannot register nil scheme handler")
	}
	if _, registered := registeredInputSchemes[name]; registered {
		panic("scheme already registered")
	}
	registeredInputSchemes[name] = handler
}

// Receives a Node and a reference string. ref is n.Reference after it has been
// parsed by the protocol detector, and excludes the scheme ("scheme://")
// portion of the string, if it was given. Also receives an input source.
type OutputSchemeHandler func(opt *Options, node *OutputNode, ref string, src *Source) (err error)

var registeredOutputSchemes = map[string]OutputSchemeHandler{}

func RegisterOutputScheme(name string, handler OutputSchemeHandler) {
	if handler == nil {
		panic("cannot register nil scheme handler")
	}
	if _, registered := registeredOutputSchemes[name]; registered {
		panic("scheme already registered")
	}
	registeredOutputSchemes[name] = handler
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
	ID             string // Name used to identify the node.
	Reference      string // Raw string that refers to a source.
	Format         string // Forced file format. If empty, it is filled in after being guessed.
	ParsedProtocol string // The protocol parsed from the reference.
}

func (node *InputNode) ResolveReference(opt *Options) (src *Source, err error) {
	scheme, nextPart := parseScheme(node.Reference)
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
	ID             string // Name used to identify the node.
	Reference      string // Raw string that refers to a source.
	Format         string // Forced file format. If empty, it is filled in after being guessed.
	ParsedProtocol string // The protocol parsed from the reference.
}

func (node *OutputNode) ResolveReference(opt *Options, src *Source) (err error) {
	scheme, nextPart := parseScheme(node.Reference)
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

type Nodes struct {
	In    []*InputNode
	Out   []*OutputNode
	Graph [][2]int
}

func IsAlnum(s string) bool {
	for _, r := range s {
		if (r >= '0' && r <= '9') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= 'a' && r <= 'z') ||
			r == '_' {
			continue
		}
		return false
	}
	return true
}
func IsDigit(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			continue
		}
		return false
	}
	return true
}

func MakeNodes(fnodes []FlagNode) (nodes *Nodes, err error) {
	nodes = new(Nodes)

	type unresolvedMap struct {
		nodeType   uint8  // Type of the parent node
		parentNode string // ID of the parent node
		mapStr     string // Unparsed map string
	}
	const (
		nodeEmpty  uint8 = iota // no parent
		nodeInput               // -i type node
		nodeOutput              // -o type node
	)

	unresolvedMaps := []unresolvedMap{}
	inNumericID, outNumericID := 0, 0
	inNodeIDs := map[string]int{}
	outNodeIDs := map[string]int{}

	for _, fnode := range fnodes {
		switch fnode.Type {
		case NodeTypeNone:
			for _, mapping := range fnode.Mapping {
				unresolvedMaps = append(unresolvedMaps, unresolvedMap{nodeEmpty, "", mapping})
			}
		case NodeTypeInput:
			node := InputNode{
				Reference: fnode.Reference,
				ID:        fnode.ID,
				Format:    fnode.Format,
			}
			if node.ID != "" {
				// Manually assigned ID; check for integrity.
				if !IsAlnum(node.ID) {
					return nil, fmt.Errorf("ID %q contains non-alphanumeric characters", node.ID)
				}
				if _, exists := inNodeIDs[node.ID]; exists {
					return nil, fmt.Errorf("input node with ID %q already exists", node.ID)
				}
			} else {
				// Automatically assigned ID; make sure it doesn't conflict
				// with a previous node.
				for {
					node.ID = strconv.Itoa(inNumericID)
					inNumericID++
					if _, exists := inNodeIDs[node.ID]; !exists {
						break
					}
				}
			}

			for _, mapping := range fnode.Mapping {
				unresolvedMaps = append(unresolvedMaps, unresolvedMap{nodeInput, "", mapping})
			}

			inNodeIDs[node.ID] = len(nodes.In)
			nodes.In = append(nodes.In, &node)
		case NodeTypeOutput:
			node := OutputNode{
				Reference: fnode.Reference,
				ID:        fnode.ID,
				Format:    fnode.Format,
			}
			if node.ID != "" {
				if !IsAlnum(node.ID) {
					return nil, fmt.Errorf("ID %q contains non-alphanumeric characters", node.ID)
				}
				if _, exists := outNodeIDs[node.ID]; exists {
					return nil, fmt.Errorf("output node with ID %q already exists", node.ID)
				}
			} else {
				for {
					node.ID = strconv.Itoa(outNumericID)
					outNumericID++
					if _, exists := outNodeIDs[node.ID]; !exists {
						break
					}
				}
			}

			for _, mapping := range fnode.Mapping {
				unresolvedMaps = append(unresolvedMaps, unresolvedMap{nodeOutput, "", mapping})
			}

			outNodeIDs[node.ID] = len(nodes.Out)
			nodes.Out = append(nodes.Out, &node)
		}
	}

	if len(unresolvedMaps) == 0 {
		// map each input to each output
		for i := range nodes.In {
			for o := range nodes.Out {
				nodes.Graph = append(nodes.Graph, [2]int{i, o})
			}
		}
		return
	}

	var ErrInvalid = errors.New("invalid")
	var ErrSyntax = errors.New("syntax")

	// Parse a string used to map inputs to outputs.
	parseMapping := func(m unresolvedMap, inNodeIDs, outNodeIDs map[string]int) (in, out map[string]bool, err error) {
		node := m.nodeType
		v := m.mapStr

		in = make(map[string]bool)
		out = make(map[string]bool)

		i := 0
		add := true

		const (
			stateOrphan uint8 = iota // Mapping only has one side defined.
			stateInput               // The input side of the mapping.
			stateOutput              // The output side of the mapping.
		)

		// Sides are separated by a ':'. If there is no separator, then the
		// mapping is orphaned, and must be derived from context.
		state := stateOrphan
		if strings.IndexByte(v, ':') > -1 {
			state = stateInput
		}

		// Parse data that identifies one or more nodes.
	ParseSelector:
		add = true
	ParseID:
		if i >= len(v) {
			goto Finish
		}
		switch c := v[i]; {
		case c == ' ', c == '\t', c == '\f', c == '\r', c == '\n':
			// Skip whitespace
			i++
			goto ParseSelector
		case c == '*':
			// Wildcard: select all nodes.
			switch state {
			case stateOrphan:
				// Select mappings based on the parent node.
				switch node {
				case nodeInput:
					// Node is an input, so map it to each output node.
					for id := range outNodeIDs {
						out[id] = add
					}
				case nodeOutput:
					// Node is an output, so map each input node to it.
					for id := range inNodeIDs {
						in[id] = add
					}
				default:
					// Orphaned mappings with no parent node are invalid.
					goto Invalid
				}
			case stateInput:
				// Currently parsing the input side, so select each input node.
				for id := range inNodeIDs {
					in[id] = add
				}
			case stateOutput:
				// Currently parsing the output side, so select each output node.
				for id := range outNodeIDs {
					out[id] = add
				}
			default:
				return nil, nil, ErrSyntax
			}
			i++
		case c == '-':
			// Negate the next selection.
			if !add {
				// '--' is bad syntax.
				return nil, nil, ErrSyntax
			}
			add = false
			i++
			goto ParseID
		case IsAlnum(string(c)):
			// Parse a literal identifier.
			j := i
			for ; j < len(v); j++ {
				if !IsAlnum(string(v[j])) {
					break
				}
			}
			switch state {
			case stateOrphan:
				// Select mapping based on the parent node.
				switch node {
				case nodeInput:
					if _, exists := outNodeIDs[v[i:j]]; !exists {
						goto Invalid
					}
					out[v[i:j]] = add
				case nodeOutput:
					if _, exists := inNodeIDs[v[i:j]]; !exists {
						goto Invalid
					}
					in[v[i:j]] = add
				}
			case stateInput:
				if _, exists := inNodeIDs[v[i:j]]; !exists {
					goto Invalid
				}
				in[v[i:j]] = add
			case stateOutput:
				if _, exists := outNodeIDs[v[i:j]]; !exists {
					goto Invalid
				}
				out[v[i:j]] = add
			default:
				return nil, nil, ErrSyntax
			}
			i = j
		}

		// Parse a separator inbetween selectors.
	ParseSeperator:
		if i >= len(v) {
			goto Finish
		}
		switch c := v[i]; c {
		case ' ', '\t', '\f', '\r', '\n':
			// Skip whitespace.
			i++
			goto ParseSeperator
		case ',':
			// Parse another selectors.
			i++
		case ':':
			// Switch side from input to output.
			if state == stateOutput {
				return nil, nil, ErrSyntax
			}
			state = stateOutput
			i++
		default:
			return nil, nil, ErrSyntax
		}
		goto ParseSelector

	Finish:
		// Make sure results are valid.
		switch {
		case
			// Mapping with two sides must have selections on both sides.
			state != stateOrphan && (len(in) == 0 || len(out) == 0),
			// Orphaned mapping to an input node must have at least one output mapping.
			state == stateOrphan && node == nodeInput && len(out) == 0,
			// Orphaned mapping to an output node must have at least one input mapping.
			state == stateOrphan && node == nodeOutput && len(in) == 0,
			// Empty node cannot have an orphaned mapping.
			state == stateOrphan && node == nodeEmpty:
			goto Invalid
		}

		// If the mapping was orphaned, then the adjacent side will be empty.
		// Fill it in with the parent node.
		if state == stateOrphan {
			switch m.nodeType {
			case nodeInput:
				in[m.parentNode] = true
			case nodeOutput:
				out[m.parentNode] = true
			}
		}

		return in, out, nil

	Invalid:
		return nil, nil, ErrInvalid
	}

	for _, m := range unresolvedMaps {
		in, out, err := parseMapping(m, inNodeIDs, outNodeIDs)
		if err != nil {
			return nil, err
		}

		isorted := make([]string, 0, len(in))
		for id := range in {
			isorted = append(isorted, id)
		}
		sort.Strings(isorted)
		osorted := make([]string, 0, len(out))
		for id := range out {
			osorted = append(osorted, id)
		}
		sort.Strings(osorted)

		for _, inID := range isorted {
			i := inNodeIDs[inID]
			for _, outID := range osorted {
				o := outNodeIDs[outID]
				// Remove duplicates.
				g := [2]int{i, o}
				for n, m := range nodes.Graph {
					if m == g {
						nodes.Graph = nodes.Graph[:n+copy(nodes.Graph[n:], nodes.Graph[n+1:])]
					}
				}
				// Negated mappings are not added back in.
				if in[inID] && out[outID] {
					nodes.Graph = append(nodes.Graph, g)
				}
			}
		}
	}
	return
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
const CommandUsage = `[OPTIONS...]

rbxmk options are grouped together as "nodes". Certain flags delimit nodes.
For example, the -i flag delimits an input node, and also specifies a
reference for that node. The -o flag delimits an output node, also defining a
reference. All flags given before a delimiting flag are counted as being apart
of the node. All flags after a delimiter will be apart of the next node.

Several flags, like --id, specify information for the node they are apart of.

Other flags, like --options, are global; they do not belong to any particular
node, and may be specified anywhere.

In general, any flag may be specified multiple times. If the flag requires a
single value, then only the last flag will be counted.`

type FlagOptions struct {
	InputReference  func(string) `short:"i" long:"input" description:"Define the reference of an input node. Delimits an input node." long-description:"" value-name:"REF"`
	OutputReference func(string) `short:"o" long:"output" description:"Define the reference of an output node. Delimits an output node." long-description:"" value-name:"REF"`
	NodeID          func(string) `short:"" long:"id" description:"Force the ID of the current node." long-description:"" value-name:"STRING"`
	NodeMap         func(string) `short:"" long:"map" description:"Map input nodes to output nodes." long-description:"" value-name:"MAPPING"`
	NodeFormat      func(string) `short:"" long:"format" description:"Force the format of the current node." long-description:"" value-name:"STRING"`
	OptionsFile     func(string) `short:"" long:"options" description:"Set options from a file." long-description:"" value-name:"FILE"`
	APIFile         string       `short:"" long:"api" description:"Get API data from a file for more accurate format decoding." long-description:"" value-name:"FILE"`
}

type NodeType uint8

const (
	NodeTypeNone NodeType = iota
	NodeTypeInput
	NodeTypeOutput
)

type FlagNode struct {
	Type      NodeType
	Reference string
	ID        string
	Mapping   []string
	Format    string
}

func ParseOptionsFile(r io.Reader) (args []string, err error) {
	buf := bufio.NewReader(r)
	currentLine := make([]byte, 0, 1024)
	for {
		part, isPrefix, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		currentLine = append(currentLine, part...)
		if !isPrefix {
			if currentLine[0] != '#' {
				line := strings.TrimFunc(string(currentLine), unicode.IsSpace)
				if line != "" {
					var i int
					var r rune
					for i, r = range line {
						if unicode.IsSpace(r) {
							break
						}
					}
					name := line[:i]
					line = strings.TrimLeftFunc(line[i:], unicode.IsSpace)
					if line != "" {
						args = append(args, "--"+name)
						args = append(args, line)
					} else {
						args = append(args, "-"+name)
					}
				}
			}
			currentLine = currentLine[:0]
		}
	}
	return args, nil
}

func LoadOptionsFile(files *[]os.FileInfo, options *FlagOptions, path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	for _, fi := range *files {
		if os.SameFile(fileInfo, fi) {
			return fmt.Errorf("detected recursive file: %s", path)
		}
	}
	*files = append(*files, fileInfo)

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	args, err := ParseOptionsFile(file)
	file.Close()
	if err != nil {
		return err
	}

	fp := flags.NewParser(options, flags.HelpFlag)
	fp.Usage = CommandUsage
	if _, err := fp.ParseArgs(args); err != nil {
		return err
	}
	return nil
}

func main() {
	flagNodes := make([]FlagNode, 1)
	parsedOptionFiles := make([]os.FileInfo, 0, 1)
	var flagOptions FlagOptions
	flagOptions = FlagOptions{
		InputReference: func(s string) {
			flagNodes[len(flagNodes)-1].Reference = s
			flagNodes[len(flagNodes)-1].Type = NodeTypeInput
			flagNodes = append(flagNodes, FlagNode{})
		},
		OutputReference: func(s string) {
			flagNodes[len(flagNodes)-1].Reference = s
			flagNodes[len(flagNodes)-1].Type = NodeTypeOutput
			flagNodes = append(flagNodes, FlagNode{})
		},
		NodeID: func(s string) {
			flagNodes[len(flagNodes)-1].ID = s
		},
		NodeMap: func(s string) {
			i := len(flagNodes) - 1
			flagNodes[i].Mapping = append(flagNodes[i].Mapping, s)
		},
		NodeFormat: func(s string) {
			flagNodes[len(flagNodes)-1].Format = s
		},
		OptionsFile: func(s string) {
			if err := LoadOptionsFile(&parsedOptionFiles, &flagOptions, s); err != nil {
				Fatalf("failed to load options file: %s", err)
			}
		},
	}

	fp := flags.NewParser(&flagOptions, flags.HelpFlag)
	fp.Usage = CommandUsage
	if _, err := fp.Parse(); err != nil {
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

	nodes, err := MakeNodes(flagNodes)
	if err != nil {
		Fatalf("failed to parse flag nodes: %s", err)
	}

	options := &Options{
		API: LoadAPI(flagOptions.APIFile),
	}

	// Gather Sources from inputs.
	sources := make([]*Source, len(nodes.In))
	for i, node := range nodes.In {
		var err error
		if sources[i], err = node.ResolveReference(options); err != nil {
			Fatalf("error resolving reference of input %q: %s", node.ID, err)
		}
	}

	// Map inputs to outputs.
	for _, m := range nodes.Graph {
		node := nodes.Out[m[1]]
		if err := node.ResolveReference(options, sources[m[0]]); err != nil {
			Fatalf("error resolving reference of output %q: %s", node.ID, err)
		}
	}
}
