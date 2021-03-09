package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Dictionary) }
func Dictionary() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "Dictionary",
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			if s.CycleGuard() {
				defer s.CycleClear()
			}
			dict, ok := v.(rtypes.Dictionary)
			if !ok {
				return nil, rbxmk.TypeError("Dictionary", v.Type())
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
				table.RawSetString(k, lv[0])
			}
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
			if s.CycleGuard() {
				defer s.CycleClear()
			}
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError("table", lvs[0].Type().String())
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
	}
}
