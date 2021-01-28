package sources

import (
	"bytes"
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/anaminus/rbxmk/sources/internal/clipboard"
)

// Get list of each unique format from arguments.
func getFormatSelectors(s rbxmk.State, n int, decode bool) (selectors []rbxmk.FormatSelector, err error) {
	values := s.Pull(n, "Tuple").(rtypes.Tuple)
	selectors = make([]rbxmk.FormatSelector, 0, len(values))
loop:
	for i, value := range values {
		selector, ok := value.(rbxmk.FormatSelector)
		if !ok {
			return nil, rbxmk.TypeError(s.L, i+n, "FormatSelector")
		}
		if decode {
			if selector.Format.Decode == nil {
				return nil, fmt.Errorf("cannot encode with format %s", selector.Format.Name)
			}
		} else {
			if selector.Format.Encode == nil {
				return nil, fmt.Errorf("cannot encode with format %s", selector.Format.Name)
			}
		}
		for _, f := range selectors {
			if selector.Format.Name == f.Format.Name {
				// Skip duplicate formats.
				continue loop
			}
		}
		selectors = append(selectors, selector)
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
		return s.RaiseError(err.Error())
	}

	// Get list of media types from each format.
	mediaTypes := []string{}
	mediaFormats := []rbxmk.FormatSelector{}
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
		return s.RaiseError(err.Error())
	}
	selector := mediaFormats[f]
	v, err := selector.Format.Decode(selector, bytes.NewReader(b))
	if err != nil {
		return s.RaiseError(err.Error())
	}
	return s.Push(v)
}

func clipboardWrite(s rbxmk.State) int {
	value := s.Pull(1, "Variant")

	selectors, err := getFormatSelectors(s, 2, false)
	if err != nil {
		return s.RaiseError(err.Error())
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
					return s.RaiseError(err.Error())
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
		return s.RaiseError(err.Error())
	}
	return 0
}
