package reflect

import (
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
			s.PushToTable(table, lua.LString("Headers"), options.Headers)
			s.PushToTable(table, lua.LString("Body"), options.Body)
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
				Headers:        s.PullFromTableOpt(table, lua.LString("Headers"), "HTTPHeaders", rtypes.HTTPHeaders(nil)).(rtypes.HTTPHeaders),
				Body:           s.PullFromTableOpt(table, lua.LString("Body"), "Variant", nil),
			}
			return options, nil
		},
	}
}
