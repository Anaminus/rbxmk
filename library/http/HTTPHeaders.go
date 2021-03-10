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

func init() { register(HTTPHeaders) }
func HTTPHeaders() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "HTTPHeaders",
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			headers, ok := v.(rtypes.HTTPHeaders)
			if !ok {
				return nil, rbxmk.TypeError{Want: "HTTPHeaders", Got: v.Type()}
			}
			table := s.L.CreateTable(0, len(headers))
			for name, values := range headers {
				vs := s.L.CreateTable(len(values), 0)
				for _, value := range values {
					vs.Append(lua.LString(value))
				}
				table.RawSetString(name, vs)
			}
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lvs[0].Type().String()}
			}
			headers := make(rtypes.HTTPHeaders)
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
		Dump: func() dump.TypeDef {
			return dump.TypeDef{Underlying: dt.Map{K: dt.Prim("string"), V: dt.Or{dt.Prim("string"), dt.Array{T: dt.Prim("string")}}}}
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
