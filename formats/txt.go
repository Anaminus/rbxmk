package formats

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Text) }
func Text() rbxmk.Format {
	return rbxmk.Format{
		Name: "txt",
		CanDecode: func(typeName string) bool {
			return typeName == "string"
		},
		Decode: func(f rbxmk.FormatOptions, b []byte) (v types.Value, err error) {
			return types.String(b), nil
		},
		Encode: func(f rbxmk.FormatOptions, v types.Value) (b []byte, err error) {
			s := rtypes.Stringlike{Value: v}
			if !s.IsStringlike() {
				return nil, cannotEncode(v)
			}
			return []byte(s.Stringlike()), nil
		},
	}
}
