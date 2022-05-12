package formats

import (
	"io"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Binary) }
func Binary() rbxmk.Format {
	return rbxmk.Format{
		Name:       "bin",
		MediaTypes: []string{"application/octet-stream"},
		CanDecode: func(g rbxmk.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == rtypes.T_BinaryString
		},
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			b, err := io.ReadAll(r)
			if err != nil {
				return nil, err
			}
			return types.BinaryString(b), nil
		},
		Encode: func(g rbxmk.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			b, ok := rtypes.Stringable(v)
			if !ok {
				return cannotEncode(v)
			}
			_, err := w.Write([]byte(b))
			return err
		},
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/bin:Summary",
				Description: "Formats/bin:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			reflect.String,
		},
	}
}
