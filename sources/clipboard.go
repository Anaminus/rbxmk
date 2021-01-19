package sources

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/anaminus/rbxmk/sources/internal/clipboard"
	"github.com/robloxapi/types"
)

// Get list of each unique format from arguments.
func getFormats(s rbxmk.State, n int, decode bool) (formats []rbxmk.Format, err error) {
	formatNames := s.Pull(n, "Tuple").(rtypes.Tuple)
	formats = make([]rbxmk.Format, 0, len(formatNames))
loop:
	for i, formatName := range formatNames {
		formatName, ok := formatName.(types.String)
		if !ok {
			return nil, rbxmk.TypeError(s.L, i+n, "string")
		}
		format := s.Format(string(formatName))
		if format.Name == "" {
			return nil, fmt.Errorf("unknown format %q", string(formatName))
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
		for _, f := range formats {
			if format.Name == f.Name {
				// Skip duplicate formats.
				continue loop
			}
		}
		formats = append(formats, format)
	}
	return formats, nil
}

func init() { register(Clipboard) }
func Clipboard() rbxmk.Source {
	return rbxmk.Source{
		Name: "clipboard",
		Read: func(s rbxmk.State) (b []byte, err error) {
			formats, err := getFormats(s, 1, true)
			if err != nil {
				return nil, err
			}

			mediaTypes := []string{}
			mediaDefined := map[string]struct{}{}
			for _, format := range formats {
				for _, mediaType := range format.MediaTypes {
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
			formats, err := getFormats(s, 1, false)
			if err != nil {
				return err
			}

			clipboardFormats := []clipboard.Format{}
			mediaDefined := map[string]struct{}{}
			for _, format := range formats {
				for _, mediaType := range format.MediaTypes {
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
	formats, err := getFormats(s, 1, true)
	if err != nil {
		return s.RaiseError(err.Error())
	}

	// Get list of media types from each format.
	mediaTypes := []string{}
	mediaFormats := []rbxmk.Format{}
	mediaDefined := map[string]struct{}{}
	for _, format := range formats {
		for _, mediaType := range format.MediaTypes {
			if _, ok := mediaDefined[mediaType]; ok {
				continue
			}
			mediaTypes = append(mediaTypes, mediaType)
			mediaFormats = append(mediaFormats, format)
			mediaDefined[mediaType] = struct{}{}
		}
	}

	// Read and decode.
	f, b, err := clipboard.Read(mediaTypes...)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	format := mediaFormats[f]
	v, err := format.Decode(rbxmk.FormatOptions{}, b)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	return s.Push(v)
}

func clipboardWrite(s rbxmk.State) int {
	value := s.Pull(1, "Variant")

	formats, err := getFormats(s, 2, false)
	if err != nil {
		return s.RaiseError(err.Error())
	}

	// Get list of media types and content from each format. The same content is
	// written for each media type defined by a format. Only the first content
	// for each media type is written.
	clipboardFormats := []clipboard.Format{}
	mediaDefined := map[string]struct{}{}
	for _, format := range formats {
		var b []byte
		for _, mediaType := range format.MediaTypes {
			if _, ok := mediaDefined[mediaType]; ok {
				continue
			}
			if b == nil {
				var err error
				if b, err = format.Encode(rbxmk.FormatOptions{}, value); err != nil {
					return s.RaiseError(err.Error())
				}
			}
			clipboardFormats = append(clipboardFormats, clipboard.Format{
				Name:    mediaType,
				Content: b,
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
