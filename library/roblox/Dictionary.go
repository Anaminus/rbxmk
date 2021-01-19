package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Dictionary) }
func Dictionary() Reflector {
	return Reflector{
		Name: "Dictionary",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			if s.Cycle == nil {
				s.Cycle = &Cycle{}
				defer func() { s.Cycle = nil }()
			}
			dict, ok := v.(rtypes.Dictionary)
			if !ok {
				return nil, TypeError(nil, 0, "Dictionary")
			}
			if s.Cycle.Has(&dict) {
				return nil, fmt.Errorf("dictionaries cannot be cyclic")
			}
			s.Cycle.Put(&dict)
			variantRfl := s.Reflector("Variant")
			table := s.L.CreateTable(0, len(dict))
			for k, v := range dict {
				lv, err := variantRfl.PushTo(s, variantRfl, v)
				if err != nil {
					return nil, err
				}
				table.RawSetString(k, lv[0])
			}
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			if s.Cycle == nil {
				s.Cycle = &Cycle{}
				defer func() { s.Cycle = nil }()
			}
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, TypeError(nil, 0, "table")
			}
			if s.Cycle.Has(table) {
				return nil, fmt.Errorf("tables cannot be cyclic")
			}
			s.Cycle.Put(table)
			variantRfl := s.Reflector("Variant")
			dict := make(rtypes.Dictionary)
			table.ForEach(func(k, lv lua.LValue) {
				if err != nil {
					return
				}
				var v types.Value
				if v, err = variantRfl.PullFrom(s, variantRfl, lv); err == nil {
					dict[k.String()] = v
				}
			})
			if err != nil {
				return nil, err
			}
			return dict, nil
		},
	}
}
