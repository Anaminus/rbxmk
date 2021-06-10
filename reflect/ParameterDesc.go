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

func init() { register(ParameterDesc) }
func ParameterDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "ParameterDesc",
		PushTo: func(s rbxmk.State, v types.Value) (lv lua.LValue, err error) {
			param, ok := v.(rtypes.ParameterDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: "ParameterDesc", Got: v.Type()}
			}
			var table *lua.LTable
			if param.Optional {
				table = s.L.CreateTable(0, 3)
			} else {
				table = s.L.CreateTable(0, 2)
			}
			s.PushToDictionary(table, "Type", rtypes.TypeDesc{Embedded: param.Parameter.Type})
			s.PushToDictionary(table, "Name", types.String(param.Name))
			if param.Optional {
				s.PushToDictionary(table, "Default", types.String(param.Default))
			}
			return table, nil
		},
		PullFrom: func(s rbxmk.State, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lv.Type().String()}
			}
			param := rtypes.ParameterDesc{
				Parameter: rbxdump.Parameter{
					Type: s.PullFromDictionary(table, "Type", "TypeDesc").(rtypes.TypeDesc).Embedded,
					Name: string(s.PullFromDictionary(table, "Name", "string").(types.String)),
				},
			}
			switch def := s.PullFromDictionaryOpt(table, "Default", rtypes.Nil, "string").(type) {
			case rtypes.NilType:
			case types.String:
				param.Optional = true
				param.Default = string(def)
			default:
				s.ReflectorError(0)
			}
			return param, nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"Type":    dt.Prim("TypeDesc"),
					"Name":    dt.Prim("string"),
					"Default": dt.Optional{T: dt.Prim("string")},
				},
				Summary:     "Types/ParameterDesc:Summary",
				Description: "Types/ParameterDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			TypeDesc,
			String,
		},
	}
}
