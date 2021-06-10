package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
)

func init() { register(ClassDesc) }
func ClassDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "ClassDesc",
		PushTo: func(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			desc, ok := v.(rtypes.ClassDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: "ClassDesc", Got: v.Type()}
			}
			class := rbxdump.Class(desc)
			return s.MustReflector("DescFields").PushTo(s, rtypes.DescFields(class.Fields()))
		},
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			fields, err := s.MustReflector("DescFields").PullFrom(s, lv)
			if err != nil {
				return nil, err
			}
			class := rbxdump.Class{}
			class.SetFields(rbxdump.Fields(fields.(rtypes.DescFields)))
			return rtypes.ClassDesc(class), nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"Name":           dt.Prim("string"),
					"Superclass":     dt.Prim("string"),
					"MemoryCategory": dt.Prim("string"),
					"Tags":           dt.Array{T: dt.Prim("string")},
				},
				Summary:     "Types/ClassDesc:Summary",
				Description: "Types/ClassDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			String,
		},
	}
}
