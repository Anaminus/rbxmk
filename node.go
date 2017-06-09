package rbxmk

import (
	"fmt"

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
		scheme := opt.Schemes.Input(schemeName)
		if scheme == nil {
			return nil, errors.New("input scheme \"" + schemeName + "\" has not been registered")
		}
		if ext, src, err = scheme.Handler(opt, node, nextPart); err != nil {
			return nil, err
		}
		ref = ref[1:]
	}

	drills, _ := opt.Formats.InputDrills(ext)
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

type addrSource struct {
	src *Source
}

func (a addrSource) Get() (v interface{}, err error) {
	return a.src, nil
}

func (a addrSource) Set(v interface{}) (err error) {
	switch v := v.(type) {
	case *Source:
		*a.src = *v
	default:
		return fmt.Errorf("received unexpected type")
	}
	return nil
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
	scheme := opt.Schemes.Output(schemeName)
	if scheme == nil {
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
	drills, exists := opt.Formats.OutputDrills(ext)
	if !exists {
		return errors.New("invalid format \"" + ext + "\"")
	}
	for _, drill := range drills {
		if addr, ref, err = drill(opt, addr, ref); err != nil && err != EOD {
			return err
		}
	}
	resolver, _ := opt.Formats.OutputResolver(ext)
	if err = resolver(node.Reference, addr, src); err != nil {
		return err
	}
	return nil
}
