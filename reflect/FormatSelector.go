package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(FormatSelector) }
func FormatSelector() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_FormatSelector,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			switch v := v.(type) {
			case types.Stringlike:
				table := c.CreateTable(0, 1)
				table.RawSetString("Format", lua.LString(v.Stringlike()))
				return table, nil
			case rtypes.FormatSelector:
				if c.CycleGuard() {
					defer c.CycleClear()
				}
				if c.CycleMark(&v) {
					return nil, fmt.Errorf("format selectors cannot be cyclic")
				}
				format := c.Format(v.Format)
				if format.Name == "" {
					return nil, fmt.Errorf("unknown format")
				}
				if len(format.Options) == 0 {
					table := c.CreateTable(0, 1)
					table.RawSetString("Format", lua.LString(format.Name))
					return table, nil
				}
				table := c.CreateTable(0, len(format.Options))
				for field, fieldTypes := range format.Options {
					value, ok := v.Options[field]
					if ok {
						for _, fieldType := range fieldTypes {
							if v.Type() == fieldType {
								rfl := c.Reflector(fieldType)
								if rfl.Name == "" {
									return nil, fmt.Errorf("unknown type %q for option %s of format %s", fieldType, field, format.Name)
								}
								v, err := rfl.PushTo(c, value)
								if err != nil {
									return nil, fmt.Errorf("field %s for format %s: %w", field, format.Name, err)
								}
								table.RawSetString(field, v)
							}
						}
						return nil, fmt.Errorf("expected type %s for option %s of format %s, got %s", c.ListTypes(fieldTypes), field, format.Name, value.Type())
					}
				}
				return table, nil
			default:
				return nil, rbxmk.TypeError{Want: "FormatSelector or string", Got: v.Type()}
			}
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			switch v := lv.(type) {
			case lua.LString:
				format := c.Format(string(v))
				if format.Name == "" {
					return nil, fmt.Errorf("unknown format %q", v)
				}
				return rtypes.FormatSelector{Format: format.Name}, nil
			case *lua.LTable:
				if c.CycleGuard() {
					defer c.CycleClear()
				}
				if c.CycleMark(v) {
					return nil, fmt.Errorf("tables cannot be cyclic")
				}
				name, ok := v.RawGetString("Format").(lua.LString)
				if !ok {
					return nil, fmt.Errorf("Format field must be a string")
				}
				format := c.Format(string(name))
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
				for field, fieldTypes := range format.Options {
					v, err := c.PullAnyFromDictionaryOpt(v, field, nil, fieldTypes...)
					if err != nil {
						return nil, fmt.Errorf("field %s for format %s: %w", field, format.Name, err)
					}
					if v != nil {
						sel.Options[field] = v
					}
				}
				return sel, nil
			default:
				return nil, rbxmk.TypeError{Want: "string or table", Got: v.Type().String()}
			}
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.FormatSelector:
				*p = v.(rtypes.FormatSelector)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category: "rbxmk",
				Underlying: dt.Or{
					dt.Prim(rtypes.T_String),
					dt.Table{
						Struct: dt.Struct{
							"Format": dt.Prim(rtypes.T_String),
						},
						Map: dt.Map{
							K: dt.Prim(rtypes.T_String),
							V: dt.Prim(rtypes.T_Any),
						},
					},
				},
				Summary:     "Types/FormatSelector:Summary",
				Description: "Types/FormatSelector:Description",
			}
		},
	}
}
