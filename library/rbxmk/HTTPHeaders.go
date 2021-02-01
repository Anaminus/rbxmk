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
				switch v := lv.(type) {
				case lua.LString:
					headers[string(name)] = []string{string(v)}
				case *lua.LTable:
					n := v.Len()
					if n == 0 {
						err = fmt.Errorf("header %q must be string or array of strings", string(name))
						return
					}
					values := make([]string, n)
					for i := 1; i <= n; i++ {
						value, ok := v.RawGetInt(i).(lua.LString)
						if !ok {
							err = fmt.Errorf("expected string from index %d of header %q, got %s", i, string(name), value.Type())
							return
						}
						values[i-1] = string(value)
					}
					headers[string(name)] = values
				default:
					err = fmt.Errorf("header %q must be string or array of strings", string(name))
					return
				}
			})
			if err != nil {
				return nil, err
			}
			return headers, nil
		},
	}
}
