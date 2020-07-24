package formats

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func Binary() rbxmk.Format {
	return rbxmk.Format{
		Name: "bin",
		Decode: func(b []byte) (v types.Value, err error) {
			return types.BinaryString(b), nil
		},
		Encode: func(v types.Value) (b []byte, err error) {
			s := rtypes.Stringlike{Value: v}
			if !s.IsStringlike() {
				return nil, cannotEncode(v)
			}
			return []byte(s.Stringlike()), nil
		},
	}
}
