package formats

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func Text() rbxmk.Format {
	return rbxmk.Format{
		Name: "txt",
		Decode: func(b []byte) (v rbxmk.Value, err error) {
			return types.String(b), nil
		},
		Encode: func(v rbxmk.Value) (b []byte, err error) {
			s := rtypes.Stringlike{Value: v}
			if !s.IsStringlike() {
				return nil, cannotEncode(v)
			}
			return []byte(s.Stringlike()), nil
		},
	}
}
