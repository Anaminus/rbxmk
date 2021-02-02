package sources

import (
	"bytes"
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/anaminus/rbxmk/sources/internal/clipboard"
	"github.com/robloxapi/types"
)

type formatSelector struct {
	Format  rbxmk.Format
	Options rtypes.Dictionary
}

// ValueOf returns the value of field. Returns nil if the value does not exist.
func (f formatSelector) ValueOf(field string) types.Value {
	return f.Options[field]
}

// Get list of each unique format from arguments.
func getFormatSelectors(s rbxmk.State, n int, decode bool) (selectors []formatSelector, err error) {
	values := s.Pull(n, "Tuple").(rtypes.Tuple)
	selectors = make([]formatSelector, 0, len(values))
loop:
	for i, value := range values {
		selector, ok := value.(rtypes.FormatSelector)
		if !ok {
			return nil, rbxmk.TypeError(s.L, i+n, "FormatSelector")
		}
		format := s.Format(selector.Format)
		if format.Name == "" {
			return nil, fmt.Errorf("unknown format %q", selector.Format)
		}
		if decode {
			if format.Decode == nil {
				return nil, fmt.Errorf("cannot encode with format %s", format.Name)
			}
		} else {
			if format.Encode == nil {
				return nil, fmt.Errorf("cannot encode with format %s", format.Name)
			}
		}
		for _, f := range selectors {
			if format.Name == f.Format.Name {
				// Skip duplicate formats.
				continue loop
			}
		}
		selectors = append(selectors, formatSelector{
			Format:  format,
			Options: selector.Options,
		})
	}
	return selectors, nil
}

func init() { register(Clipboard) }
func Clipboard() rbxmk.Source {
	return rbxmk.Source{
		Name: "clipboard",
		Read: func(s rbxmk.State) (b []byte, err error) {
			selectors, err := getFormatSelectors(s, 1, true)
			if err != nil {
				return nil, err
			}

			mediaTypes := []string{}
			mediaDefined := map[string]struct{}{}
			for _, selector := range selectors {
				for _, mediaType := range selector.Format.MediaTypes {
					if _, ok := mediaDefined[mediaType]; ok {
						continue
					}
					mediaTypes = append(mediaTypes, mediaType)
					mediaDefined[mediaType] = struct{}{}
				}
			}

			_, b, err = clipboard.Read(mediaTypes...)
			return b, err
		},
		Write: func(s rbxmk.State, b []byte) (err error) {
			selectors, err := getFormatSelectors(s, 1, false)
			if err != nil {
				return err
			}

			clipboardFormats := []clipboard.Format{}
			mediaDefined := map[string]struct{}{}
			for _, selector := range selectors {
				for _, mediaType := range selector.Format.MediaTypes {
					if _, ok := mediaDefined[mediaType]; ok {
						continue
					}
					clipboardFormats = append(clipboardFormats, clipboard.Format{
						Name:    mediaType,
						Content: b,
					})
					mediaDefined[mediaType] = struct{}{}
				}
			}

			return clipboard.Write(clipboardFormats)
		},
		Library: rbxmk.Library{
			Open: func(s rbxmk.State) *lua.LTable {
				lib := s.L.CreateTable(0, 2)
				lib.RawSetString("read", s.WrapFunc(clipboardRead))
				lib.RawSetString("write", s.WrapFunc(clipboardWrite))
				return lib
			},
		},
	}
}

func clipboardRead(s rbxmk.State) int {
	selectors, err := getFormatSelectors(s, 1, true)
	if err != nil {
		return s.RaiseError("%w", err)
	}

	// Get list of media types from each format.
	mediaTypes := []string{}
	mediaFormats := []formatSelector{}
	mediaDefined := map[string]struct{}{}
	for _, selector := range selectors {
		for _, mediaType := range selector.Format.MediaTypes {
			if _, ok := mediaDefined[mediaType]; ok {
				continue
			}
			mediaTypes = append(mediaTypes, mediaType)
			mediaFormats = append(mediaFormats, selector)
			mediaDefined[mediaType] = struct{}{}
		}
	}

	// Read and decode.
	f, b, err := clipboard.Read(mediaTypes...)
	if err != nil {
		return s.RaiseError("%w", err)
	}
	selector := mediaFormats[f]
	v, err := selector.Format.Decode(selector, bytes.NewReader(b))
	if err != nil {
		return s.RaiseError("%w", err)
	}
	return s.Push(v)
}

func clipboardWrite(s rbxmk.State) int {
	value := s.Pull(1, "Variant")

	selectors, err := getFormatSelectors(s, 2, false)
	if err != nil {
		return s.RaiseError("%w", err)
	}

	// Get list of media types and content from each format. The same content is
	// written for each media type defined by a format. Only the first content
	// for each media type is written.
	clipboardFormats := []clipboard.Format{}
	mediaDefined := map[string]struct{}{}
	for _, selector := range selectors {
		var w bytes.Buffer
		var written bool
		for _, mediaType := range selector.Format.MediaTypes {
			if _, ok := mediaDefined[mediaType]; ok {
				continue
			}
			if !written {
				if err := selector.Format.Encode(selector, &w, value); err != nil {
					return s.RaiseError("%w", err)
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
		return s.RaiseError("%w", err)
	}
	return 0
}
