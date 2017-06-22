package rbxmk

import (
	"errors"
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
	Reference []string // Raw strings that refer to data.
	Data      Data     // Pre-resolved Data.
	Format    string   // Forced file format.
}

func (node *InputNode) ResolveReference(opt *Options) (data Data, err error) {
	ref := node.Reference
	var ext string
	if node.Data != nil {
		data = node.Data
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
		scheme := opt.Schemes.Input(schemeName)
		if scheme == nil {
			return nil, errors.New("input scheme \"" + schemeName + "\" has not been registered")
		}
		modref := make([]string, len(ref))
		copy(modref[1:], ref[1:])
		modref[0] = nextPart
		if ext, ref, data, err = scheme.Handler(opt, node, modref); err != nil {
			return nil, err
		}
	}

	if !opt.Formats.Registered(ext) {
		return nil, errors.New("unknown format \"" + ext + "\"")
	}
	for _, drill := range opt.Formats.InputDrills(ext) {
		if data, ref, err = drill(opt, data, ref); err != nil && err != EOD {
			return nil, err
		}
	}
	return data, nil
}

type OutputNode struct {
	Reference []string // Raw string that refers to data.
	Data      Data     // Pre-resolved Data.
	Format    string   // Forced file format. If empty, it is filled in after being guessed.
}

func (node *OutputNode) ResolveReference(opt *Options, data Data) (err error) {
	ref := node.Reference
	var ext string
	if node.Data != nil {
		ext = node.Format
		_, err = node.drillOutput(opt, node.Data, ref, ext, data)
		return err
	}

	if len(ref) < 1 {
		return errors.New("node requires at least one reference argument")
	}
	schemeName, nextPart := parseScheme(ref[0])
	if schemeName == "" {
		// Assume file:// scheme.
		schemeName = "file"
	}
	scheme := opt.Schemes.Output(schemeName)
	if scheme == nil {
		return errors.New("output scheme \"" + schemeName + "\" has not been registered")
	}
	modref := make([]string, len(ref))
	copy(modref[1:], ref[1:])
	modref[0] = nextPart
	var outdata Data
	if ext, ref, outdata, err = scheme.Handler(opt, node, modref); err != nil {
		return err
	}
	if outdata, err = node.drillOutput(opt, outdata, ref, ext, data); err != nil {
		return err
	}
	return scheme.Finalizer(opt, node, modref, ext, outdata)
}

func (node *OutputNode) drillOutput(opt *Options, indata Data, ref []string, ext string, data Data) (outdata Data, err error) {
	if !opt.Formats.Registered(ext) {
		return nil, errors.New("invalid format \"" + ext + "\"")
	}
	for _, drill := range opt.Formats.OutputDrills(ext) {
		if indata, ref, err = drill(opt, indata, ref); err != nil && err != EOD {
			return nil, err
		}
	}
	resolver := opt.Formats.Resolver(ext)
	return resolver(opt, indata, data)
}
