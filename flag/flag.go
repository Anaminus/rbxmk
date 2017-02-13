package flag

import (
	"fmt"
	"strconv"
	"time"
)

// Set is a set of Nodes parsed from a list of arguments.
type Set struct {
	parsed      bool
	args        []string
	flagDefs    map[string]*Def
	currentNode *Node
	finalNodes  []*Node
}

// Nodes returns a list of parsed nodes.
func (f *Set) Nodes() []*Node {
	return f.finalNodes
}

func parseValue(s string, value interface{}) (err error) {
	switch value := value.(type) {
	case *bool:
		v, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		*value = v
	case *int:
		v, err := strconv.ParseInt(s, 0, 64)
		if err != nil {
			return err
		}
		*value = int(v)
	case *int64:
		v, err := strconv.ParseInt(s, 0, 64)
		if err != nil {
			return err
		}
		*value = v
	case *uint:
		v, err := strconv.ParseUint(s, 0, 64)
		if err != nil {
			return err
		}
		*value = uint(v)
	case *uint64:
		v, err := strconv.ParseUint(s, 0, 64)
		if err != nil {
			return err
		}
		*value = v
	case *float64:
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		*value = v
	case *string:
		*value = s
	case *time.Duration:
		v, err := time.ParseDuration(s)
		if err != nil {
			return err
		}
		*value = v
	}
	return nil
}

// Def is a definition of a flag.
type Def struct {
	Name      string // Name as it appears on command line.
	Default   string // Default value (as text).
	SingleArg string // Value when flag is single (bool-like). Empty means the flag is not single.
	IsNode    bool   // Whether the flag is a node.
	Usage     string // Help message.
}

// ParseDefault parses the default value string provided by the flag's Def.
// The string is interpreted depending on the type of the given value. The
// value type must be a pointer.
func (f Def) ParseDefault(value interface{}) (err error) {
	return parseValue(f.Default, value)
}

// Flag is the state of a flag.
type Flag struct {
	*Def
	Value string
}

// ParseValue parses the Value string. The string is interpreted depending on
// the type of the given value. The value type must be a pointer.
func (f Flag) ParseValue(value interface{}) (err error) {
	return parseValue(f.Value, value)
}

// Node is a flag that also contains a list of other flags.
type Node struct {
	Flag
	Flags []*Flag
}

// Lookup returns the parsed value from a flag of the given name, which is
// under the node. If a flag of the given name does not exist, then the value
// is not set, and no error is returned.
func (n *Node) Lookup(name string, value interface{}) error {
	for _, flag := range n.Flags {
		if flag.Name != name {
			continue
		}
		if err := flag.ParseValue(value); err != nil {
			return fmt.Errorf("invalid value %q for flag -%s: %v", value, name, err)
		}
	}
	return nil
}

// Define adds a number of flag definitions to the set.
func (f *Set) Define(defs ...Def) {
	for _, def := range defs {
		if _, alreadythere := f.flagDefs[def.Name]; alreadythere {
			// Happens only if flags are declared with identical names.
			panic(fmt.Sprintf("flag redefined: %s", def.Name))
		}
		if f.flagDefs == nil {
			f.flagDefs = make(map[string]*Def)
		}

		d := def
		f.flagDefs[def.Name] = &d
	}
}

// parseOne parses one flag. It reports whether a flag was seen.
func (f *Set) parseOne() (bool, error) {
	if len(f.args) == 0 {
		return false, nil
	}
	s := f.args[0]
	if len(s) == 0 || s[0] != '-' || len(s) == 1 {
		return false, nil
	}
	numMinuses := 1
	if s[1] == '-' {
		numMinuses++
		if len(s) == 2 { // "--" terminates the flags
			f.args = f.args[1:]
			return false, nil
		}
	}
	name := s[numMinuses:]
	if len(name) == 0 || name[0] == '-' || name[0] == '=' {
		return false, fmt.Errorf("bad flag syntax: %s", s)
	}

	// it's a flag. does it have an argument?
	f.args = f.args[1:]
	hasValue := false
	value := ""
	for i := 1; i < len(name); i++ { // equals cannot be first
		if name[i] == '=' {
			value = name[i+1:]
			hasValue = true
			name = name[0:i]
			break
		}
	}
	def, alreadythere := f.flagDefs[name] // BUG
	if !alreadythere {
		return false, fmt.Errorf("flag provided but not defined: -%s", name)
	}

	flag := Flag{Def: def}
	if def.SingleArg != "" { // special case: doesn't need an arg
		if hasValue {
			flag.Value = value
		} else {
			flag.Value = def.SingleArg
		}
	} else {
		// It must have a value, which might be the next argument.
		if !hasValue && len(f.args) > 0 {
			// value is the next arg
			hasValue = true
			value, f.args = f.args[0], f.args[1:]
		}
		if !hasValue {
			return false, fmt.Errorf("flag needs an argument: -%s", name)
		}
		flag.Value = value
	}

	if f.currentNode == nil {
		f.currentNode = new(Node)
	}
	if def.IsNode {
		f.currentNode.Flag = flag
		f.finalNodes = append(f.finalNodes, f.currentNode)
		f.currentNode = new(Node)
	} else {
		f.currentNode.Flags = append(f.currentNode.Flags, &flag)
	}
	return true, nil
}

// Parse parses a given list of arguments according to its definitions.
func (f *Set) Parse(arguments []string) error {
	f.args = arguments
	for {
		seen, err := f.parseOne()
		if seen {
			continue
		}
		if err == nil {
			break
		}
		return err
	}
	if len(f.currentNode.Flags) > 0 {
		f.finalNodes = append(f.finalNodes, f.currentNode)
	}
	f.currentNode = nil
	f.parsed = true
	return nil
}

// Lookup returns the parsed value of the last flag in the set of the given
// name. Is may be used to search for global flags.
func (f *Set) Lookup(name string, value interface{}) error {
	if !f.parsed {
		return fmt.Errorf("flags have not been parsed")
	}
	def, ok := f.flagDefs[name]
	if !ok {
		return fmt.Errorf("undefined flag", name)
	}
	if err := parseValue(def.Default, value); err != nil {
		panic(fmt.Sprintf("invalid default %q: %v", value, err))
	}
	if def.IsNode {
		for _, node := range f.finalNodes {
			if err := node.ParseValue(value); err != nil {
				return fmt.Errorf("invalid value %q: %v", value, err)
			}
		}
		return nil
	}
	for _, node := range f.finalNodes {
		for _, flag := range node.Flags {
			if flag.Name != name {
				continue
			}
			if err := flag.ParseValue(value); err != nil {
				return fmt.Errorf("invalid value %q: %v", value, err)
			}
		}
	}
	return nil
}
