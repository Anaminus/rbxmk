package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(HTTPOptions) }
func HTTPOptions() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "HTTPOptions",
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			options, ok := v.(rtypes.HTTPOptions)
			if !ok {
				return nil, rbxmk.TypeError{Want: "HTTPOptions", Got: v.Type()}
			}
			table := s.L.CreateTable(0, 7)
			s.PushToTable(table, lua.LString("URL"), types.String(options.URL))
			s.PushToTable(table, lua.LString("Method"), types.String(options.Method))
			s.PushToTable(table, lua.LString("RequestFormat"), options.RequestFormat)
			s.PushToTable(table, lua.LString("ResponseFormat"), options.ResponseFormat)
			s.PushToTable(table, lua.LString("Headers"), options.Headers)
			s.PushToTable(table, lua.LString("Cookies"), options.Cookies)
			s.PushToTable(table, lua.LString("Body"), options.Body)
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lvs[0].Type().String()}
			}
			options := rtypes.HTTPOptions{
				URL:            string(s.PullFromTable(table, lua.LString("URL"), "string").(types.String)),
				Method:         string(s.PullFromTableOpt(table, lua.LString("Method"), "string", types.String("GET")).(types.String)),
				RequestFormat:  s.PullFromTableOpt(table, lua.LString("RequestFormat"), "FormatSelector", rtypes.FormatSelector{}).(rtypes.FormatSelector),
				ResponseFormat: s.PullFromTableOpt(table, lua.LString("ResponseFormat"), "FormatSelector", rtypes.FormatSelector{}).(rtypes.FormatSelector),
				Headers:        s.PullFromTableOpt(table, lua.LString("Headers"), "HTTPHeaders", rtypes.HTTPHeaders(nil)).(rtypes.HTTPHeaders),
				Cookies:        s.PullFromTableOpt(table, lua.LString("Cookies"), "Cookies", rtypes.Cookies(nil)).(rtypes.Cookies),
				Body:           s.PullFromTableOpt(table, lua.LString("Body"), "Variant", nil),
			}
			return options, nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"URL":            dt.Prim("string"),
					"Method":         dt.Optional{T: dt.Prim("string")},
					"RequestFormat":  dt.Optional{T: dt.Prim("FormatSelector")},
					"ResponseFormat": dt.Optional{T: dt.Prim("FormatSelector")},
					"Headers":        dt.Optional{T: dt.Prim("HTTPHeaders")},
					"Cookies":        dt.Optional{T: dt.Prim("Cookies")},
					"Body":           dt.Optional{T: dt.Prim("any")},
				},
				Summary:     "Types/HTTPOptions:Summary",
				Description: "Types/HTTPOptions:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Cookies,
			FormatSelector,
			HTTPHeaders,
			String,
			Variant,
		},
	}
}
