package reflect

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
)

func init() { register(JsonValue) }
func JsonValue() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     rtypes.T_JsonValue,
		PushTo:   PushBasicTo,
		PullFrom: PullBasicFrom,
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Or{
					dt.Prim(rtypes.T_Nil),
					dt.Prim(rtypes.T_Bool),
					dt.Prim(rtypes.T_Number),
					dt.Prim(rtypes.T_String),
					dt.Prim(rtypes.T_Array),
					dt.Prim(rtypes.T_Dictionary),
				},
				Summary:     "Types/JsonValue:Summary",
				Description: "Types/JsonValue:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			Bool,
			Dictionary,
			Nil,
			Number,
			String,
		},
	}
}
