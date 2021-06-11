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
		PushTo: func(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			param, ok := v.(rtypes.ParameterDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: "ParameterDesc", Got: v.Type()}
			}
			var table *lua.LTable
			if param.Optional {
				table = s.CreateTable(0, 3)
			} else {
				table = s.CreateTable(0, 2)
			}
			if err := s.PushToDictionary(table, "Type", rtypes.TypeDesc{Embedded: param.Parameter.Type}); err != nil {
				return nil, err
			}
			if err := s.PushToDictionary(table, "Name", types.String(param.Name)); err != nil {
				return nil, err
			}
			if param.Optional {
				if err := s.PushToDictionary(table, "Default", types.String(param.Default)); err != nil {
					return nil, err
				}
			}
			return table, nil
		},
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lv.Type().String()}
			}
			typ, err := s.PullFromDictionary(table, "Type", "TypeDesc")
			if err != nil {
				return nil, err
			}
			name, err := s.PullFromDictionary(table, "Name", "string")
			if err != nil {
				return nil, err
			}
			param := rtypes.ParameterDesc{
				Parameter: rbxdump.Parameter{
					Type: typ.(rtypes.TypeDesc).Embedded,
					Name: string(name.(types.String)),
				},
			}
			def, err := s.PullFromDictionaryOpt(table, "Default", rtypes.Nil, "string")
			if err != nil {
				return nil, err
			}
			switch def := def.(type) {
			case rtypes.NilType:
			case types.String:
				param.Optional = true
				param.Default = string(def)
			default:
				s.ReflectorError(0)
			}
			return param, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.ParameterDesc:
				*p = v.(rtypes.ParameterDesc)
			default:
				return setPtrErr(p, v)
			}
			return nil
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
