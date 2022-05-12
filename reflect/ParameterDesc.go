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
		Name: rtypes.T_ParameterDesc,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			param, ok := v.(rtypes.ParameterDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_ParameterDesc, Got: v.Type()}
			}
			var table *lua.LTable
			if param.Optional {
				table = c.CreateTable(0, 3)
			} else {
				table = c.CreateTable(0, 2)
			}
			if err := c.PushToDictionary(table, "Type", rtypes.TypeDesc{Embedded: param.Parameter.Type}); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "Name", types.String(param.Name)); err != nil {
				return nil, err
			}
			if param.Optional {
				if err := c.PushToDictionary(table, "Default", types.String(param.Default)); err != nil {
					return nil, err
				}
			}
			return table, nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_Table, Got: lv.Type().String()}
			}
			typ, err := c.PullFromDictionary(table, "Type", rtypes.T_TypeDesc)
			if err != nil {
				return nil, err
			}
			name, err := c.PullFromDictionary(table, "Name", rtypes.T_String)
			if err != nil {
				return nil, err
			}
			param := rtypes.ParameterDesc{
				Parameter: rbxdump.Parameter{
					Type: typ.(rtypes.TypeDesc).Embedded,
					Name: string(name.(types.String)),
				},
			}
			def, err := c.PullFromDictionaryOpt(table, "Default", rtypes.Nil, rtypes.T_String)
			if err != nil {
				return nil, err
			}
			switch def := def.(type) {
			case rtypes.NilType:
			case types.String:
				param.Optional = true
				param.Default = string(def)
			default:
				c.ReflectorError()
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
					"Type":    dt.Prim(rtypes.T_TypeDesc),
					"Name":    dt.Prim(rtypes.T_String),
					"Default": dt.Optional{T: dt.Prim(rtypes.T_String)},
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
