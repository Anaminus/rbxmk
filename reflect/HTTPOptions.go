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
		PushTo: func(s rbxmk.State, v types.Value) (lv lua.LValue, err error) {
			options, ok := v.(rtypes.HTTPOptions)
			if !ok {
				return nil, rbxmk.TypeError{Want: "HTTPOptions", Got: v.Type()}
			}
			table := s.L.CreateTable(0, 7)
			s.PushToDictionary(table, "URL", types.String(options.URL))
			s.PushToDictionary(table, "Method", types.String(options.Method))
			s.PushToDictionary(table, "RequestFormat", options.RequestFormat)
			s.PushToDictionary(table, "ResponseFormat", options.ResponseFormat)
			s.PushToDictionary(table, "Headers", options.Headers)
			s.PushToDictionary(table, "Cookies", options.Cookies)
			s.PushToDictionary(table, "Body", options.Body)
			return table, nil
		},
		PullFrom: func(s rbxmk.State, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lv.Type().String()}
			}
			options := rtypes.HTTPOptions{
				URL:            string(s.PullFromDictionary(table, "URL", "string").(types.String)),
				Method:         string(s.PullFromDictionaryOpt(table, "Method", types.String("GET"), "string").(types.String)),
				ResponseFormat: s.PullFromDictionaryOpt(table, "ResponseFormat", rtypes.FormatSelector{}, "FormatSelector").(rtypes.FormatSelector),
				Headers:        s.PullFromDictionaryOpt(table, "Headers", rtypes.HTTPHeaders(nil), "HTTPHeaders").(rtypes.HTTPHeaders),
				Cookies:        s.PullFromDictionaryOpt(table, "Cookies", rtypes.Cookies(nil), "Cookies").(rtypes.Cookies),
			}
			options.RequestFormat = s.PullFromDictionaryOpt(table, "RequestFormat", rtypes.FormatSelector{}, "FormatSelector").(rtypes.FormatSelector)
			if format := s.Format(options.RequestFormat.Format); format.Name != "" {
				options.Body = s.PullAnyFromDictionaryOpt(table, "Body", nil, format.EncodeTypes...)
			} else {
				options.Body = s.PullFromDictionaryOpt(table, "Body", nil, "Variant")
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
