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
		PushTo: func(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			typ, ok := v.(rtypes.TypeDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: "TypeDesc", Got: v.Type()}
			}
			table := s.CreateTable(0, 2)
			if err := s.PushToDictionary(table, "Category", types.String(typ.Embedded.Category)); err != nil {
				return nil, err
			}
			if err := s.PushToDictionary(table, "Name", types.String(typ.Embedded.Name)); err != nil {
				return nil, err
			}
			return table, nil
		},
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lv.Type().String()}
			}
			category, err := s.PullFromDictionary(table, "Category", "string")
			if err != nil {
				return nil, err
			}
			name, err := s.PullFromDictionary(table, "Name", "string")
			if err != nil {
				return nil, err
			}
			typ := rtypes.TypeDesc{
				Embedded: rbxdump.Type{
					Category: string(category.(types.String)),
					Name:     string(name.(types.String)),
				},
			}
			return typ, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.TypeDesc:
				*p = v.(rtypes.TypeDesc)
			default:
				return setPtrErr(p, v)
			}
			return nil
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
