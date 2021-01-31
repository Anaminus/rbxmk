package reflect

import (
	"fmt"
	"net/http"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(HTTPOptions) }
func HTTPOptions() Reflector {
	return Reflector{
		Name: "HTTPOptions",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			options, ok := v.(rtypes.HTTPOptions)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "HTTPOptions")
			}
			table := s.L.CreateTable(0, 5)
			s.PushToTable(table, lua.LString("URL"), types.String(options.URL))
			s.PushToTable(table, lua.LString("Method"), types.String(options.Method))
			s.PushToTable(table, lua.LString("RequestFormat"), options.RequestFormat)
			s.PushToTable(table, lua.LString("ResponseFormat"), options.ResponseFormat)
			s.PushToTable(table, lua.LString("Body"), options.Body)
			table.RawSetString("Headers", pushHTTPHeaders(s, options.Headers))
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "table")
			}
			options := rtypes.HTTPOptions{
				URL:            string(s.PullFromTable(table, lua.LString("URL"), "string").(types.String)),
				Method:         string(s.PullFromTableOpt(table, lua.LString("Method"), "string", types.String("GET")).(types.String)),
				RequestFormat:  s.PullFromTableOpt(table, lua.LString("RequestFormat"), "FormatSelector", rtypes.FormatSelector{}).(rtypes.FormatSelector),
				ResponseFormat: s.PullFromTableOpt(table, lua.LString("ResponseFormat"), "FormatSelector", rtypes.FormatSelector{}).(rtypes.FormatSelector),
				Body:           s.PullFromTableOpt(table, lua.LString("Body"), "Variant", nil),
			}
			switch table := table.RawGetString("Headers").(type) {
			case *lua.LNilType:
			case *lua.LTable:
				if options.Headers, err = pullHTTPHeaders(table); err != nil {
					return nil, err
				}
			default:
				return nil, fmt.Errorf("Headers must be a table")
			}
			return options, nil
		},
	}
}

func pushHTTPHeaders(s State, headers http.Header) (table *lua.LTable) {
	table = s.L.CreateTable(0, len(headers))
	for name, values := range headers {
		vs := s.L.CreateTable(len(values), 0)
		for _, value := range values {
			vs.Append(lua.LString(value))
		}
		table.RawSetString(name, vs)
	}
	return table
}

func pullHTTPHeaders(table *lua.LTable) (headers http.Header, err error) {
	headers = make(http.Header)
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
}
