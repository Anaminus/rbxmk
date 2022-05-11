package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

const T_Dictionary = "Dictionary"

func init() { register(Dictionary) }
func Dictionary() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: T_Dictionary,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			if c.CycleGuard() {
				defer c.CycleClear()
			}
			dict, ok := v.(rtypes.Dictionary)
			if !ok {
				return nil, rbxmk.TypeError{Want: T_Dictionary, Got: v.Type()}
			}
			if c.CycleMark(&dict) {
				return nil, fmt.Errorf("dictionaries cannot be cyclic")
			}
			variantRfl := c.MustReflector(T_Variant)
			table := c.CreateTable(0, len(dict))
			for k, v := range dict {
				lv, err := variantRfl.PushTo(c, v)
				if err != nil {
					return nil, err
				}
				table.RawSetString(k, lv)
			}
			return table, nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			if c.CycleGuard() {
				defer c.CycleClear()
			}
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: T_Table, Got: lv.Type().String()}
			}
			if c.CycleMark(table) {
				return nil, fmt.Errorf("tables cannot be cyclic")
			}
			variantRfl := c.MustReflector(T_Variant)
			dict := make(rtypes.Dictionary)
			err = table.ForEach(func(k, lv lua.LValue) error {
				v, err := variantRfl.PullFrom(c, lv)
				if err != nil {
					return err
				}
				dict[k.String()] = v
				return nil
			})
			if err != nil {
				return nil, err
			}
			return dict, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.Dictionary:
				*p = v.(rtypes.Dictionary)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/Dictionary:Summary",
				Description: "Types/Dictionary:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Variant,
		},
	}
}
