package formats

import (
	"io"
	"io/ioutil"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Text) }
func Text() rbxmk.Format {
	return rbxmk.Format{
		Name:       "txt",
		MediaTypes: []string{"text/plain"},
		CanDecode: func(f rbxmk.FormatOptions, typeName string) bool {
			return typeName == "string"
		},
		Decode: func(f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			s, err := ioutil.ReadAll(r)
			if err != nil {
				return nil, err
			}
			return types.String(s), nil
		},
		Encode: func(f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			s := rtypes.Stringable{Value: v}
			if !s.IsStringable() {
				return cannotEncode(v)
			}
			_, err := w.Write([]byte(s.Stringable()))
			return err
		},
	}
}
