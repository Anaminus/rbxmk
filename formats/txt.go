package formats

import (
	"io"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Text) }
func Text() rbxmk.Format {
	return rbxmk.Format{
		Name:       "txt",
		MediaTypes: []string{"text/plain"},
		CanDecode: func(g rbxmk.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == "string"
		},
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			b, err := io.ReadAll(r)
			if err != nil {
				return nil, err
			}
			return types.String(b), nil
		},
		Encode: func(g rbxmk.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			b, ok := rtypes.Stringable(v)
			if !ok {
				return cannotEncode(v)
			}
			_, err := w.Write([]byte(b))
			return err
		},
	}
}
