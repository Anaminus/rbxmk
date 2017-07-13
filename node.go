package rbxmk

import (
	"fmt"
)

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
	Options   Options
	Reference []string // Raw strings that refer to data.
	Data      Data     // Pre-resolved Data.
	Format    string   // Forced file format.
}

type NodeError struct {
	Type string
	Func string
	Err  error
}

func (err *NodeError) Error() string {
	if err.Func == "" {
		return fmt.Sprintf("%s node error: %s", err.Type, err.Err.Error())
	}
	return fmt.Sprintf("%s node error: %s: %s", err.Type, err.Func, err.Err.Error())
}

func (node *InputNode) ResolveReference() (data Data, err error) {
	opt := node.Options
	ref := node.Reference
	if node.Data != nil {
		data = node.Data
	} else {
		if len(ref) < 1 {
			return nil, &NodeError{"input", "", fmt.Errorf("node requires at least one reference argument")}
		}
		schemeName, nextPart := parseScheme(ref[0])
		if schemeName == "" {
			// Assume file:// scheme.
			schemeName = "file"
		}
		scheme := opt.Schemes.Input(schemeName)
		if scheme == nil {
			return nil, &NodeError{"input", "", fmt.Errorf("input scheme %s has not been registered", schemeName)}
		}
		modref := make([]string, len(ref))
		copy(modref[1:], ref[1:])
		modref[0] = nextPart
		if ref, data, err = scheme.Handler(opt, node, modref); err != nil {
			err = &NodeError{"input", fmt.Sprintf("%s scheme, Handler", schemeName), err}
			return nil, err
		}
	}

	for i := 1; err != EOD; i++ {
		if data, ref, err = data.Drill(opt, ref); err != nil && err != EOD {
			return nil, &NodeError{"input", fmt.Sprintf("Drill #%d", i), err}
		}
	}
	return data, nil
}

type OutputNode struct {
	Options   Options
	Reference []string // Raw string that refers to data.
	Data      Data     // Pre-resolved Data.
	Format    string   // Forced file format. If empty, it is filled in after being guessed.
}

func (node *OutputNode) ResolveReference(indata Data) (err error) {
	opt := node.Options
	ref := node.Reference
	if node.Data != nil {
		_, err = node.drillOutput(opt, node.Data, ref, indata)
		return err
	}

	if len(ref) < 1 {
		return &NodeError{"output", "", fmt.Errorf("node requires at least one reference argument")}
	}
	schemeName, nextPart := parseScheme(ref[0])
	if schemeName == "" {
		// Assume file:// scheme.
		schemeName = "file"
	}
	scheme := opt.Schemes.Output(schemeName)
	if scheme == nil {
		return &NodeError{"output", "", fmt.Errorf("output scheme %s has not been registered", schemeName)}
	}
	modref := make([]string, len(ref))
	copy(modref[1:], ref[1:])
	modref[0] = nextPart
	var rootdata Data
	var ext string
	if ext, ref, rootdata, err = scheme.Handler(opt, node, modref); err != nil {
		err = &NodeError{"output", fmt.Sprintf("%s scheme, Handler", schemeName), err}
		return err
	}
	var outdata Data
	if outdata, err = node.drillOutput(opt, rootdata, ref, indata); err != nil {
		return err
	}
	if err = scheme.Finalizer(opt, node, modref, ext, outdata); err != nil {
		err = &NodeError{"output", fmt.Sprintf("%s scheme, Finalizer", schemeName), err}
	}
	return err
}

func (node *OutputNode) drillOutput(opt Options, rootdata Data, ref []string, indata Data) (outdata Data, err error) {
	drilldata := rootdata
	for i := 1; err != EOD; i++ {
		if drilldata, ref, err = drilldata.Drill(opt, ref); err != nil && err != EOD {
			return nil, &NodeError{"output", fmt.Sprintf("Drill #%d", i), err}
		}
	}
	if outdata, err = indata.Merge(opt, rootdata, drilldata); err != nil {
		err = &NodeError{"output", "Merge", err}
	}
	return outdata, err
}
