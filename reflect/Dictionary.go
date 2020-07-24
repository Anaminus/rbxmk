package reflect

import (
	"fmt"

	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/yuin/gopher-lua"
)

func Dictionary() Type {
	return Type{
		Name: "Dictionary",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			if s.Cycle == nil {
				s.Cycle = &Cycle{}
				defer func() { s.Cycle = nil }()
			}
			dict, ok := v.(rtypes.Dictionary)
			if !ok {
				return nil, TypeError(nil, 0, "map[string]Value")
			}
			if s.Cycle.Has(&dict) {
				return nil, fmt.Errorf("dictionaries cannot be cyclic")
			}
			s.Cycle.Put(&dict)
			variantType := s.Type("Variant")
			table := s.L.CreateTable(0, len(dict))
			for k, v := range dict {
				lv, err := variantType.ReflectTo(s, variantType, v)
				if err != nil {
					return nil, err
				}
				table.RawSetString(k, lv[0])
			}
			return []lua.LValue{table}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
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
			variantType := s.Type("Variant")
			dict := make(rtypes.Dictionary)
			table.ForEach(func(k, lv lua.LValue) {
				if err != nil {
					return
				}
				var v Value
				if v, err = variantType.ReflectFrom(s, variantType, lv); err == nil {
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
