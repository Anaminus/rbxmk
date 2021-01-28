package formats

import (
	"io"
	"io/ioutil"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Binary) }
func Binary() rbxmk.Format {
	return rbxmk.Format{
		Name:       "bin",
		MediaTypes: []string{"application/octet-stream"},
		CanDecode: func(f rbxmk.FormatOptions, typeName string) bool {
			return typeName == "BinaryString"
		},
		Decode: func(f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			b, err := ioutil.ReadAll(r)
			if err != nil {
				return nil, err
			}
			return types.BinaryString(b), nil
		},
		Encode: func(f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			s := rtypes.Stringlike{Value: v}
			if !s.IsStringlike() {
				return cannotEncode(v)
			}
			_, err := w.Write([]byte(s.Stringlike()))
			return err
		},
	}
}
