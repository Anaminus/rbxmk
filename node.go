package rbxmk

import (
	"errors"
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
	var ext string
	if node.Data != nil {
		data = node.Data
		ext = node.Format
	} else {
		if len(ref) < 1 {
			return nil, &NodeError{"input", "", errors.New("node requires at least one reference argument")}
		}
		schemeName, nextPart := parseScheme(ref[0])
		if schemeName == "" {
			// Assume file:// scheme.
			schemeName = "file"
		}
		scheme := opt.Schemes.Input(schemeName)
		if scheme == nil {
			return nil, &NodeError{"input", "", errors.New("input scheme \"" + schemeName + "\" has not been registered")}
		}
		modref := make([]string, len(ref))
		copy(modref[1:], ref[1:])
		modref[0] = nextPart
		if ext, ref, data, err = scheme.Handler(opt, node, modref); err != nil {
			err = &NodeError{"input", fmt.Sprintf("%s scheme, Handler", schemeName), err}
			return nil, err
		}
	}

	if !opt.Formats.Registered(ext) {
		return nil, &NodeError{"input", "", errors.New("unknown format \"" + ext + "\"")}
	}
	for i, drill := range opt.Formats.InputDrills(ext) {
		if data, ref, err = drill(opt, data, ref); err != nil && err != EOD {
			err = &NodeError{"input", fmt.Sprintf("%s format, Drill #%d", opt.Formats.Name(ext), i+1), err}
			return nil, err
		} else if err == EOD {
			break
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

func (node *OutputNode) ResolveReference(data Data) (err error) {
	opt := node.Options
	ref := node.Reference
	var ext string
	if node.Data != nil {
		ext = node.Format
		_, err = node.drillOutput(opt, node.Data, ref, ext, data)
		return err
	}

	if len(ref) < 1 {
		return &NodeError{"output", "", errors.New("node requires at least one reference argument")}
	}
	schemeName, nextPart := parseScheme(ref[0])
	if schemeName == "" {
		// Assume file:// scheme.
		schemeName = "file"
	}
	scheme := opt.Schemes.Output(schemeName)
	if scheme == nil {
		return &NodeError{"output", "", errors.New("output scheme \"" + schemeName + "\" has not been registered")}
	}
	modref := make([]string, len(ref))
	copy(modref[1:], ref[1:])
	modref[0] = nextPart
	var outdata Data
	if ext, ref, outdata, err = scheme.Handler(opt, node, modref); err != nil {
		err = &NodeError{"output", fmt.Sprintf("%s scheme, Handler", schemeName), err}
		return err
	}
	if outdata, err = node.drillOutput(opt, outdata, ref, ext, data); err != nil {
		return err
	}
	if err = scheme.Finalizer(opt, node, modref, ext, outdata); err != nil {
		err = &NodeError{"output", fmt.Sprintf("%s scheme, Finalizer", schemeName), err}
	}
	return err
}

func (node *OutputNode) drillOutput(opt Options, indata Data, ref []string, ext string, data Data) (outdata Data, err error) {
	if !opt.Formats.Registered(ext) {
		return nil, &NodeError{"output", "", errors.New("invalid format \"" + ext + "\"")}
	}
	for i, drill := range opt.Formats.OutputDrills(ext) {
		if indata, ref, err = drill(opt, indata, ref); err != nil && err != EOD {
			err = &NodeError{"output", fmt.Sprintf("%s format, Drill #%d", opt.Formats.Name(ext), i+1), err}
			return nil, err
		} else if err == EOD {
			break
		}
	}
	resolver := opt.Formats.Resolver(ext)
	if outdata, err = resolver(opt, indata, data); err != nil {
		err = &NodeError{"output", fmt.Sprintf("%s format, Resolver", opt.Formats.Name(ext)), err}
	}
	return outdata, err
}
