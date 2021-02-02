package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(HTTPHeaders) }
func HTTPHeaders() Reflector {
	return Reflector{
		Name: "HTTPHeaders",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			headers, ok := v.(rtypes.HTTPHeaders)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "HTTPHeaders")
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
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "table")
			}
			headers := make(rtypes.HTTPHeaders)
			table.ForEach(func(k, lv lua.LValue) {
				if err != nil {
					return
				}
				name, ok := k.(lua.LString)
				if !ok {
					return
				}
				values, e := pullStringArray(lv)
				if e != nil {
					err = fmt.Errorf("header %q: %w", string(name), e)
					return
				}
				headers[string(name)] = values
			})
			if err != nil {
				return nil, err
			}
			return headers, nil
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
