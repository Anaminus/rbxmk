package formats

import (
	"io"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

const F_RBXAttr = "rbxattr"

func init() { register(RBXAttr) }
func RBXAttr() rbxmk.Format {
	return rbxmk.Format{
		Name:       F_RBXAttr,
		MediaTypes: []string{"application/octet-stream"},
		CanDecode: func(g rtypes.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == "Instance"
		},
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return rtypes.DecodeAttributes(r)
		},
		Encode: func(g rtypes.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			return rtypes.EncodeAttributes(w, v)
		},
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/rbxattr:Summary",
				Description: "Formats/rbxattr:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			reflect.Bool,
			reflect.BrickColor,
			reflect.Color3,
			reflect.ColorSequence,
			reflect.ColorSequenceKeypoint,
			reflect.Dictionary,
			reflect.Double,
			reflect.Float,
			reflect.Nil,
			reflect.NumberRange,
			reflect.NumberSequence,
			reflect.NumberSequenceKeypoint,
			reflect.Rect,
			reflect.String,
			reflect.UDim,
			reflect.UDim2,
			reflect.Vector2,
			reflect.Vector3,
		},
	}
}
