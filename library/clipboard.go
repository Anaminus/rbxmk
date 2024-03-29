package library

import (
	"bytes"
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/library/internal/clipboard"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Clipboard) }

var Clipboard = rbxmk.Library{
	Name:     "clipboard",
	Import:   []string{"clipboard"},
	Priority: 10,
	Open:     openClipboard,
	Dump:     dumpClipboard,
	Types: []func() rbxmk.Reflector{
		reflect.FormatSelector,
		reflect.Variant,
	},
}

func openClipboard(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 2)
	lib.RawSetString("read", s.WrapFunc(clipboardRead))
	lib.RawSetString("write", s.WrapFunc(clipboardWrite))
	return lib
}

// getFormatSelectors produces a list of FormatSelectors from arguments.
func getFormatSelectors(s rbxmk.State, n int) (selectors []rtypes.FormatSelector) {
	c := s.Count()
	selectors = make([]rtypes.FormatSelector, 0, c-n+1)
	for i := n; i <= c; i++ {
		selector := s.Pull(i, rtypes.T_FormatSelector).(rtypes.FormatSelector)
		selectors = append(selectors, selector)
	}
	return selectors
}

func clipboardRead(s rbxmk.State) int {
	selectors := getFormatSelectors(s, 1)
	v, err := ClipboardSource{World: s.World}.Read(selectors...)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	if v == nil {
		return s.Push(rtypes.Nil)
	}
	return s.Push(v)
}

func clipboardWrite(s rbxmk.State) int {
	value := s.Pull(1, rtypes.T_Variant)
	selectors := getFormatSelectors(s, 2)
	err := ClipboardSource{World: s.World}.Write(value, selectors...)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return 0
}

func dumpClipboard(s rbxmk.State) dump.Library {
	return dump.Library{
		Struct: dump.Struct{
			Fields: dump.Fields{
				"read": dump.Function{
					Parameters: dump.Parameters{
						{Name: "...", Type: dt.Prim(rtypes.T_FormatSelector)},
					},
					Returns: dump.Parameters{
						{Name: "value", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
					},
					CanError:    true,
					Summary:     "Libraries/clipboard:Fields/read/Summary",
					Description: "Libraries/clipboard:Fields/read/Description",
				},
				"write": dump.Function{
					Parameters: dump.Parameters{
						{Name: "value", Type: dt.Prim(rtypes.T_Any)},
						{Name: "...", Type: dt.Prim(rtypes.T_FormatSelector)},
					},
					CanError:    true,
					Summary:     "Libraries/clipboard:Fields/write/Summary",
					Description: "Libraries/clipboard:Fields/write/Description",
				},
			},
			Summary:     "Libraries/clipboard:Summary",
			Description: "Libraries/clipboard:Description",
		},
	}
}

// ClipboardSource provides access to the clipboard of the operating system.
type ClipboardSource struct {
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

// Read reads a value from the clipboard according to the given formats. If no
// formats are given, or no data is found, then nil is returned with no error.
func (s ClipboardSource) Read(formats ...rtypes.FormatSelector) (v types.Value, err error) {
	if len(formats) == 0 {
		return nil, nil
	}
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
		if clipboard.IsNoData(err) {
			return nil, nil
		}
		return nil, err
	}
	if f < 0 {
		return nil, nil
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

// Write writes a value to the clipboard according to the given formats. If no
// formats are given, then the clipboard is cleared.
func (s ClipboardSource) Write(value types.Value, formats ...rtypes.FormatSelector) error {
	if len(formats) == 0 {
		if err := clipboard.Clear(); err != nil {
			if clipboard.IsNoData(err) {
				return nil
			}
			return err
		}
	}
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
		if clipboard.IsNoData(err) {
			return nil
		}
		return err
	}
	return nil
}
