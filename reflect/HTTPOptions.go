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
		PushTo: func(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			options, ok := v.(rtypes.HTTPOptions)
			if !ok {
				return nil, rbxmk.TypeError{Want: "HTTPOptions", Got: v.Type()}
			}
			table := s.CreateTable(0, 7)
			if err := s.PushToDictionary(table, "URL", types.String(options.URL)); err != nil {
				return nil, err
			}
			if err := s.PushToDictionary(table, "Method", types.String(options.Method)); err != nil {
				return nil, err
			}
			if err := s.PushToDictionary(table, "RequestFormat", options.RequestFormat); err != nil {
				return nil, err
			}
			if err := s.PushToDictionary(table, "ResponseFormat", options.ResponseFormat); err != nil {
				return nil, err
			}
			if err := s.PushToDictionary(table, "Headers", options.Headers); err != nil {
				return nil, err
			}
			if err := s.PushToDictionary(table, "Cookies", options.Cookies); err != nil {
				return nil, err
			}
			if err := s.PushToDictionary(table, "Body", options.Body); err != nil {
				return nil, err
			}
			return table, nil
		},
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lv.Type().String()}
			}
			url, err := s.PullFromDictionary(table, "URL", "string")
			if err != nil {
				return nil, err
			}
			method, err := s.PullFromDictionaryOpt(table, "Method", types.String("GET"), "string")
			if err != nil {
				return nil, err
			}
			responseFormat, err := s.PullFromDictionaryOpt(table, "ResponseFormat", rtypes.FormatSelector{}, "FormatSelector")
			if err != nil {
				return nil, err
			}
			headers, err := s.PullFromDictionaryOpt(table, "Headers", rtypes.HTTPHeaders(nil), "HTTPHeaders")
			if err != nil {
				return nil, err
			}
			cookies, err := s.PullFromDictionaryOpt(table, "Cookies", rtypes.Cookies(nil), "Cookies")
			if err != nil {
				return nil, err
			}
			options := rtypes.HTTPOptions{
				URL:            string(url.(types.String)),
				Method:         string(method.(types.String)),
				ResponseFormat: responseFormat.(rtypes.FormatSelector),
				Headers:        headers.(rtypes.HTTPHeaders),
				Cookies:        cookies.(rtypes.Cookies),
			}
			requestFormat, err := s.PullFromDictionaryOpt(table, "RequestFormat", rtypes.FormatSelector{}, "FormatSelector")
			if err != nil {
				return nil, err
			}
			options.RequestFormat = requestFormat.(rtypes.FormatSelector)
			if format := s.Format(options.RequestFormat.Format); format.Name != "" {
				options.Body, err = s.PullAnyFromDictionaryOpt(table, "Body", nil, format.EncodeTypes...)
			} else {
				options.Body, err = s.PullFromDictionaryOpt(table, "Body", nil, "Variant")
			}
			if err != nil {
				return nil, err
			}
			return options, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.HTTPOptions:
				*p = v.(rtypes.HTTPOptions)
			default:
				return setPtrErr(p, v)
			}
			return nil
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
