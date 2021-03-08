package library

import (
	"bytes"
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/library/internal/clipboard"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

// getFormatSelectors produces a list of FormatSelectors from arguments.
func getFormatSelectors(s rbxmk.State, n int) (selectors []rtypes.FormatSelector, err error) {
	values := s.Pull(n, "Tuple").(rtypes.Tuple)
	selectors = make([]rtypes.FormatSelector, 0, len(values))
	for i, value := range values {
		selector, ok := value.(rtypes.FormatSelector)
		if !ok {
			return nil, rbxmk.TypeError(s.L, i+n, "FormatSelector")
		}
		selectors = append(selectors, selector)
	}
	return selectors, nil
}

func init() { register(ClipboardSource, 10) }

var ClipboardSource = rbxmk.Library{
	Name: "clipboard",
	Open: func(s rbxmk.State) *lua.LTable {
		lib := s.L.CreateTable(0, 2)
		lib.RawSetString("read", s.WrapFunc(func(s rbxmk.State) int {
			selectors, err := getFormatSelectors(s, 1)
			if err != nil {
				return s.RaiseError("%s", err)
			}
			v, err := Clipboard{World: s.World}.Read(selectors...)
			if err != nil {
				return s.RaiseError("%s", err)
			}
			return s.Push(v)
		}))
		lib.RawSetString("write", s.WrapFunc(func(s rbxmk.State) int {
			value := s.Pull(1, "Variant")
			selectors, err := getFormatSelectors(s, 2)
			if err != nil {
				return s.RaiseError("%s", err)
			}
			err = Clipboard{World: s.World}.Write(value, selectors...)
			if err != nil {
				return s.RaiseError("%s", err)
			}
			return 0
		}))
		return lib
	},
	Dump: func(s rbxmk.State) dump.Library {
		return dump.Library{
			Struct: dump.Struct{
				Fields: dump.Fields{
					"read": dump.Function{
						Parameters: dump.Parameters{
							{Name: "...", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Name: "value", Type: dt.Prim("any")},
						},
					},
					"write": dump.Function{
						Parameters: dump.Parameters{
							{Name: "value", Type: dt.Prim("any")},
							{Name: "...", Type: dt.Prim("string")},
						},
					},
				},
			},
		}
	},
}

// Clipboard provides access to the clipboard of the operating system.
type Clipboard struct {
	*rbxmk.World
}

// formatOptions implements rbxmk.FormatOptions.
type formatOptions struct {
	Format  rbxmk.Format
	Options rtypes.Dictionary
}

// ValueOf returns the value of field. Returns nil if the value does not exist.
func (f formatOptions) ValueOf(field string) types.Value {
	return f.Options[field]
}

// Read reads a value from the clipboard according to the given formats.
func (s Clipboard) Read(formats ...rtypes.FormatSelector) (v types.Value, err error) {
	options := make([]formatOptions, 0, len(formats))
loop:
	for _, selector := range formats {
		format := s.Format(selector.Format)
		if format.Name == "" {
			return nil, fmt.Errorf("unknown format %q", selector.Format)
		}
		if format.Decode == nil {
			return nil, fmt.Errorf("cannot decode with format %s", format.Name)
		}
		for _, option := range options {
			if format.Name == option.Format.Name {
				// Skip duplicate formats.
				continue loop
			}
		}
		options = append(options, formatOptions{
			Format:  format,
			Options: selector.Options,
		})
	}

	// Get list of media types from each format.
	mediaTypes := []string{}
	mediaFormats := []formatOptions{}
	mediaDefined := map[string]struct{}{}
	for _, option := range options {
		for _, mediaType := range option.Format.MediaTypes {
			if _, ok := mediaDefined[mediaType]; ok {
				continue
			}
			mediaTypes = append(mediaTypes, mediaType)
			mediaFormats = append(mediaFormats, option)
			mediaDefined[mediaType] = struct{}{}
		}
	}

	// Read and decode.
	f, b, err := clipboard.Read(mediaTypes...)
	if err != nil {
		return nil, err
	}
	option := mediaFormats[f]
	if option.Format.Decode == nil {
		return nil, fmt.Errorf("cannot decode with format %s", option.Format.Name)
	}
	v, err = option.Format.Decode(s.Global, option, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Write writes a value to the clipboard according to the given formats.
func (s Clipboard) Write(value types.Value, formats ...rtypes.FormatSelector) error {
	options := make([]formatOptions, 0, len(formats))
loop:
	for _, selector := range formats {
		format := s.Format(selector.Format)
		if format.Name == "" {
			return fmt.Errorf("unknown format %q", selector.Format)
		}
		if format.Encode == nil {
			return fmt.Errorf("cannot encode with format %s", format.Name)
		}
		for _, option := range options {
			if format.Name == option.Format.Name {
				// Skip duplicate formats.
				continue loop
			}
		}
		options = append(options, formatOptions{
			Format:  format,
			Options: selector.Options,
		})
	}

	// Get list of media types and content from each format. The same content is
	// written for each media type defined by a format. Only the first content
	// for each media type is written.
	clipboardFormats := []clipboard.Format{}
	mediaDefined := map[string]struct{}{}
	for _, option := range options {
		var w bytes.Buffer
		var written bool
		for _, mediaType := range option.Format.MediaTypes {
			if _, ok := mediaDefined[mediaType]; ok {
				continue
			}
			if !written {
				if option.Format.Encode == nil {
					return fmt.Errorf("cannot encode with format %s", option.Format.Name)
				}
				if err := option.Format.Encode(s.Global, option, &w, value); err != nil {
					return err
				}
				written = true
			}
			clipboardFormats = append(clipboardFormats, clipboard.Format{
				Name:    mediaType,
				Content: w.Bytes(),
			})
			mediaDefined[mediaType] = struct{}{}
		}
	}

	// Write to clipboard.
	if err := clipboard.Write(clipboardFormats); err != nil {
		return err
	}
	return nil
}
