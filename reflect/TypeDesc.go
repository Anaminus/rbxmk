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

func init() { register(TypeDesc) }
func TypeDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "TypeDesc",
		PushTo: func(s rbxmk.State, v types.Value) (lv lua.LValue, err error) {
			typ, ok := v.(rtypes.TypeDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: "TypeDesc", Got: v.Type()}
			}
			table := s.L.CreateTable(0, 2)
			s.PushToTable(table, lua.LString("Category"), types.String(typ.Embedded.Category))
			s.PushToTable(table, lua.LString("Name"), types.String(typ.Embedded.Name))
			return table, nil
		},
		PullFrom: func(s rbxmk.State, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lv.Type().String()}
			}
			typ := rtypes.TypeDesc{
				Embedded: rbxdump.Type{
					Category: string(s.PullFromTable(table, lua.LString("Category"), "string").(types.String)),
					Name:     string(s.PullFromTable(table, lua.LString("Name"), "string").(types.String)),
				},
			}
			return typ, nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"Category": dt.Prim("string"),
					"Name":     dt.Prim("string"),
				},
				Summary:     "Types/TypeDesc:Summary",
				Description: "Types/TypeDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			String,
		},
	}
}
