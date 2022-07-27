package formats

import (
	"encoding/base64"
	"io"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

type lineSplit struct {
	w io.Writer
	s int
	n int
}

func (l *lineSplit) Write(p []byte) (n int, err error) {
	for i := 0; ; {
		var q []byte
		if len(p[i:]) < l.n {
			q = p[i:]
		} else {
			q = p[i : i+l.n]
		}
		n, err = l.w.Write(q)
		if n < len(q) {
			return
		}
		l.n -= len(q)
		i += len(q)
		if i >= len(p) {
			break
		}
		if l.n <= 0 {
			_, e := l.w.Write([]byte{'\n'})
			if e != nil {
				return
			}
			l.n = l.s
		}
	}
	return
}

const F_Base64 = "base64"

func init() { register(Base64) }
func Base64() rbxmk.Format {
	return rbxmk.Format{
		Name:       F_Base64,
		MediaTypes: []string{"application/octet-stream"},
		Options: map[string][]string{
			"Width": {rtypes.T_Int},
		},
		CanDecode: func(g rtypes.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == rtypes.T_BinaryString
		},
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			r = base64.NewDecoder(base64.StdEncoding, r)
			b, err := io.ReadAll(r)
			if err != nil {
				return nil, err
			}
			return types.BinaryString(b), nil
		},
		Encode: func(g rtypes.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			b, ok := rtypes.Stringable(v)
			if !ok {
				return cannotEncode(v)
			}
			if v, ok := intOf(f, "Width"); ok && v > 0 {
				w = &lineSplit{w: w, s: int(v), n: int(v)}
			}
			e := base64.NewEncoder(base64.StdEncoding, w)
			if _, err := e.Write([]byte(b)); err != nil {
				return err
			}
			return e.Close()
		},
		Dump: func() dump.Format {
			return dump.Format{
				Options: dump.FormatOptions{
					"Width": dump.FormatOption{
						Type:        dt.Prim(rtypes.T_Int),
						Default:     "0",
						Description: "Formats/base64:Options/Width",
					},
				},
				Summary:     "Formats/base64:Summary",
				Description: "Formats/base64:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			reflect.String,
		},
	}
}
