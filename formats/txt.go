package formats

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/types"
)

func Text() rbxmk.Format {
	return rbxmk.Format{
		Name: "txt",
		Decode: func(b []byte) (v rbxmk.Value, err error) {
			return string(b), nil
		},
		Encode: func(v rbxmk.Value) (b []byte, err error) {
			b, ok := types.Stringlike{Value: v}.Stringlike()
			if !ok {
				return nil, cannotEncode(v, false)
			}
			return b, nil
		},
	}
}
