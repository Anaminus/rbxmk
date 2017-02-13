package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/anaminus/rbxmk/flag"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxapi/dump"
	"github.com/robloxapi/rbxfile"
	"os"
	"sort"
	"strconv"
	"strings"
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
	API   *rbxapi.API
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

func MakeNodes(fnodes []*flag.Node) (nodes *Nodes, err error) {
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
		if fnode.Flag.Def == nil {
			for _, flag := range fnode.Flags {
				if flag.Name != "map" {
					continue
				}
				unresolvedMaps = append(unresolvedMaps, unresolvedMap{nodeEmpty, "", flag.Value})
			}
			continue
		}
		switch fnode.Name {
		case "i":
			node := InputNode{
				Reference: fnode.Value,
			}
			if err := fnode.Lookup("id", &node.ID); err != nil {
				return nil, err
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

			for _, flag := range fnode.Flags {
				if flag.Name != "map" {
					continue
				}
				unresolvedMaps = append(unresolvedMaps, unresolvedMap{nodeInput, node.ID, flag.Value})
			}

			inNodeIDs[node.ID] = len(nodes.In)
			nodes.In = append(nodes.In, &node)
		case "o":
			node := OutputNode{
				Reference: fnode.Value,
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

			for _, flag := range fnode.Flags {
				if flag.Name != "map" {
					continue
				}
				unresolvedMaps = append(unresolvedMaps, unresolvedMap{nodeOutput, node.ID, flag.Value})
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

type Config struct {
	API string `json:"api"`
}

func (c *Config) Load(flagset *flag.Set) {
	// Try to load defaults from file.
	var path string
	var file *os.File
	if err := flagset.Lookup("config", &path); err != nil {
		Fatalf("failed to parse -config flag: %s", err)
	}
	if path != "" {
		var err error
		if file, err = os.Open(path); err != nil {
			Fatalf("failed to open config file: %s", err)
		}
		defer file.Close()
		jd := json.NewDecoder(file)
		if err := jd.Decode(c); err != nil {
			Fatalf("failed to decode config file: %s", err)
		}
	}
	// Override from flags.
	{
		var flag string
		var err error

		flag = "api"
		if err = flagset.Lookup(flag, &c.API); err != nil {
			goto Error
		}

		return
	Error:
		Fatalf("failed to configure -%s flag: %s", flag, err)
	}

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

func main() {
	// Define and parse flags.
	flagset := new(flag.Set)
	flagset.Define(
		flag.Def{Name: "i", IsNode: true, Usage: "Define the reference of an input node."},
		flag.Def{Name: "o", IsNode: true, Usage: "Define the reference of an output node."},
		flag.Def{Name: "id", Usage: "Force the ID of the current node."},
		flag.Def{Name: "map", Usage: "Map nodes to another node."},
		flag.Def{Name: "format", Usage: "Force the format of the current node."},

		flag.Def{Name: "config", Default: "", Usage: "Set default config with a file."},
		flag.Def{Name: "api", Default: "", Usage: "Get API data from a file for more accurate format decoding."},
	)
	flagset.Parse(os.Args[1:])

	config := new(Config)
	config.Load(flagset)

	nodes, err := MakeNodes(flagset.Nodes())
	if err != nil {
		Fatalf("failed to parse flag nodes: %s", err)
	}

	options := &Options{
		API: LoadAPI(config.API),
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
