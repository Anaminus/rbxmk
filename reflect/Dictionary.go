package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Dictionary) }
func Dictionary() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "Dictionary",
		PushTo: func(s rbxmk.State, v types.Value) (lv lua.LValue, err error) {
			if s.CycleGuard() {
				defer s.CycleClear()
			}
			dict, ok := v.(rtypes.Dictionary)
			if !ok {
				return nil, rbxmk.TypeError{Want: "Dictionary", Got: v.Type()}
			}
			if s.CycleMark(&dict) {
				return nil, fmt.Errorf("dictionaries cannot be cyclic")
			}
			variantRfl := s.MustReflector("Variant")
			table := s.L.CreateTable(0, len(dict))
			for k, v := range dict {
				lv, err := variantRfl.PushTo(s, v)
				if err != nil {
					return nil, err
				}
				table.RawSetString(k, lv)
			}
			return table, nil
		},
		PullFrom: func(s rbxmk.State, lv lua.LValue) (v types.Value, err error) {
			if s.CycleGuard() {
				defer s.CycleClear()
			}
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lv.Type().String()}
			}
			if s.CycleMark(table) {
				return nil, fmt.Errorf("tables cannot be cyclic")
			}
			variantRfl := s.MustReflector("Variant")
			dict := make(rtypes.Dictionary)
			err = table.ForEach(func(k, lv lua.LValue) error {
				v, err := variantRfl.PullFrom(s, lv)
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
