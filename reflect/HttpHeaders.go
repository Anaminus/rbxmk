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

func init() { register(HttpHeaders) }
func HttpHeaders() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_HttpHeaders,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			headers, ok := v.(rtypes.HttpHeaders)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_HttpHeaders, Got: v.Type()}
			}
			table := c.CreateTable(0, len(headers))
			for name, values := range headers {
				vs := c.CreateTable(len(values), 0)
				for _, value := range values {
					vs.Append(lua.LString(value))
				}
				table.RawSetString(name, vs)
			}
			return table, nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_Table, Got: lv.Type().String()}
			}
			headers := make(rtypes.HttpHeaders)
			err = table.ForEach(func(k, lv lua.LValue) error {
				name, ok := k.(lua.LString)
				if !ok {
					return nil
				}
				values, err := pullStringArray(lv)
				if err != nil {
					return fmt.Errorf("header %q: %w", string(name), err)
				}
				headers[string(name)] = values
				return nil
			})
			if err != nil {
				return nil, err
			}
			return headers, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.HttpHeaders:
				*p = v.(rtypes.HttpHeaders)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category:    "rbxmk",
				Underlying:  dt.P(dt.Map(dt.Prim(rtypes.T_String), dt.Or(dt.Prim(rtypes.T_String), dt.Array(dt.Prim(rtypes.T_String))))),
				Summary:     "Types/HttpHeaders:Summary",
				Description: "Types/HttpHeaders:Description",
			}
		},
	}
}

// Convert a string or an array of strings.
func pullStringArray(v lua.LValue) ([]string, error) {
	switch v := v.(type) {
	case lua.LString:
		return []string{string(v)}, nil
	case *lua.LTable:
		n := v.Len()
		if n == 0 {
			return nil, fmt.Errorf("expected string or array of strings")
		}
		values := make([]string, n)
		for i := 1; i <= n; i++ {
			value, ok := v.RawGetInt(i).(lua.LString)
			if !ok {
				return nil, fmt.Errorf("index %d: expected string, got %s", i, value.Type())
			}
			values[i-1] = string(value)
		}
		return values, nil
	default:
		return nil, fmt.Errorf("expected string or array of strings")
	}
}
