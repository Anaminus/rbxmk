package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(FormatSelector) }
func FormatSelector() Reflector {
	return Reflector{
		Name: "FormatSelector",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			switch v := v.(type) {
			case types.Stringlike:
				table := s.L.CreateTable(0, 1)
				table.RawSetString("Format", lua.LString(v.Stringlike()))
				return []lua.LValue{table}, nil
			case rtypes.FormatSelector:
				if s.Cycle == nil {
					s.Cycle = &rbxmk.Cycle{}
					defer func() { s.Cycle = nil }()
				}
				if s.Cycle.Mark(&v) {
					return nil, fmt.Errorf("format selectors cannot be cyclic")
				}
				format := s.Format(v.Format)
				if format.Name == "" {
					return nil, fmt.Errorf("unknown format")
				}
				if len(format.Options) == 0 {
					table := s.L.CreateTable(0, 1)
					table.RawSetString("Format", lua.LString(format.Name))
					return []lua.LValue{table}, nil
				}
				table := s.L.CreateTable(0, len(format.Options))
				for field, typ := range format.Options {
					rfl := s.Reflector(typ)
					if rfl.Name == "" {
						return nil, fmt.Errorf("unknown type %q for option %s of format %s", typ, field, format.Name)
					}
					value, ok := v.Options[field]
					if ok {
						v, err := rfl.PushTo(s, rfl, value)
						if err != nil {
							return nil, fmt.Errorf("field %s for format %s: %w", field, format.Name, err)
						}
						table.RawSetString(field, v[0])
					}
				}
				return []lua.LValue{table}, nil
			default:
				return nil, rbxmk.TypeError(nil, 0, "FormatSelector or string")
			}
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LString:
				format := s.Format(string(v))
				if format.Name == "" {
					return nil, fmt.Errorf("unknown format %q", v)
				}
				return rtypes.FormatSelector{Format: format.Name}, nil
			case *lua.LTable:
				if s.Cycle == nil {
					s.Cycle = &rbxmk.Cycle{}
					defer func() { s.Cycle = nil }()
				}
				if s.Cycle.Mark(v) {
					return nil, fmt.Errorf("tables cannot be cyclic")
				}
				name, ok := v.RawGetString("Format").(lua.LString)
				if !ok {
					return nil, fmt.Errorf("Format field must be a string")
				}
				format := s.Format(string(name))
				if format.Name == "" {
					return nil, fmt.Errorf("unknown format %q", name)
				}
				if len(format.Options) == 0 {
					return rtypes.FormatSelector{Format: format.Name}, nil
				}
				sel := rtypes.FormatSelector{
					Format:  format.Name,
					Options: make(rtypes.Dictionary),
				}
				for field, typ := range format.Options {
					rfl := s.Reflector(typ)
					if rfl.Name == "" {
						return nil, fmt.Errorf("unknown type %q for option %s of format %s", typ, field, format.Name)
					}
					v, err := rfl.PullFrom(s, rfl, v.RawGetString(field))
					if err != nil {
						return nil, fmt.Errorf("field %s for format %s: %w", field, format.Name, err)
					}
					sel.Options[field] = v
				}
				return sel, nil
			default:
				return nil, rbxmk.TypeError(nil, 0, "string or table")
			}
		},
	}
}
