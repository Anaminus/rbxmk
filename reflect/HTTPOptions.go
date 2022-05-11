package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

const T_HTTPOptions = "HTTPOptions"

func init() { register(HTTPOptions) }
func HTTPOptions() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: T_HTTPOptions,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			options, ok := v.(rtypes.HTTPOptions)
			if !ok {
				return nil, rbxmk.TypeError{Want: T_HTTPOptions, Got: v.Type()}
			}
			table := c.CreateTable(0, 7)
			if err := c.PushToDictionary(table, "URL", types.String(options.URL)); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "Method", types.String(options.Method)); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "RequestFormat", options.RequestFormat); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "ResponseFormat", options.ResponseFormat); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "Headers", options.Headers); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "Cookies", options.Cookies); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "Body", options.Body); err != nil {
				return nil, err
			}
			return table, nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: T_Table, Got: lv.Type().String()}
			}
			url, err := c.PullFromDictionary(table, "URL", T_String)
			if err != nil {
				return nil, err
			}
			method, err := c.PullFromDictionaryOpt(table, "Method", types.String("GET"), T_String)
			if err != nil {
				return nil, err
			}
			responseFormat, err := c.PullFromDictionaryOpt(table, "ResponseFormat", rtypes.FormatSelector{}, T_FormatSelector)
			if err != nil {
				return nil, err
			}
			headers, err := c.PullFromDictionaryOpt(table, "Headers", rtypes.HTTPHeaders(nil), T_HTTPHeaders)
			if err != nil {
				return nil, err
			}
			cookies, err := c.PullFromDictionaryOpt(table, "Cookies", rtypes.Cookies(nil), T_Cookies)
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
			requestFormat, err := c.PullFromDictionaryOpt(table, "RequestFormat", rtypes.FormatSelector{}, T_FormatSelector)
			if err != nil {
				return nil, err
			}
			options.RequestFormat = requestFormat.(rtypes.FormatSelector)
			options.Body, err = c.PullEncodedFromDict(table, "Body", options.RequestFormat)
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
					"URL":            dt.Prim(T_String),
					"Method":         dt.Optional{T: dt.Prim(T_String)},
					"RequestFormat":  dt.Optional{T: dt.Prim(T_FormatSelector)},
					"ResponseFormat": dt.Optional{T: dt.Prim(T_FormatSelector)},
					"Headers":        dt.Optional{T: dt.Prim(T_HTTPHeaders)},
					"Cookies":        dt.Optional{T: dt.Prim(T_Cookies)},
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
