package formats

import (
	"io"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

const F_Text = "txt"

func init() { register(Text) }
func Text() rbxmk.Format {
	return rbxmk.Format{
		Name:       F_Text,
		MediaTypes: []string{"text/plain"},
		CanDecode: func(g rtypes.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == rtypes.T_String
		},
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			b, err := io.ReadAll(r)
			if err != nil {
				return nil, err
			}
			return types.String(b), nil
		},
		Encode: func(g rtypes.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			b, ok := rtypes.Stringable(v)
			if !ok {
				return cannotEncode(v)
			}
			_, err := w.Write([]byte(b))
			return err
		},
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/txt:Summary",
				Description: "Formats/txt:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			reflect.String,
		},
	}
}
