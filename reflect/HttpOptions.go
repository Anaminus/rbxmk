package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(HttpOptions) }
func HttpOptions() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_HttpOptions,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			options, ok := v.(rtypes.HttpOptions)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_HttpOptions, Got: v.Type()}
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
				return nil, rbxmk.TypeError{Want: rtypes.T_Table, Got: lv.Type().String()}
			}
			url, err := c.PullFromDictionary(table, "URL", rtypes.T_String)
			if err != nil {
				return nil, err
			}
			method, err := c.PullFromDictionaryOpt(table, "Method", types.String("GET"), rtypes.T_String)
			if err != nil {
				return nil, err
			}
			responseFormat, err := c.PullFromDictionaryOpt(table, "ResponseFormat", rtypes.FormatSelector{}, rtypes.T_FormatSelector)
			if err != nil {
				return nil, err
			}
			headers, err := c.PullFromDictionaryOpt(table, "Headers", rtypes.HttpHeaders(nil), rtypes.T_HttpHeaders)
			if err != nil {
				return nil, err
			}
			cookies, err := c.PullFromDictionaryOpt(table, "Cookies", rtypes.Cookies(nil), rtypes.T_Cookies)
			if err != nil {
				return nil, err
			}
			options := rtypes.HttpOptions{
				URL:            string(url.(types.String)),
				Method:         string(method.(types.String)),
				ResponseFormat: responseFormat.(rtypes.FormatSelector),
				Headers:        headers.(rtypes.HttpHeaders),
				Cookies:        cookies.(rtypes.Cookies),
			}
			requestFormat, err := c.PullFromDictionaryOpt(table, "RequestFormat", rtypes.FormatSelector{}, rtypes.T_FormatSelector)
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
			case *rtypes.HttpOptions:
				*p = v.(rtypes.HttpOptions)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"URL":            dt.Prim(rtypes.T_String),
					"Method":         dt.Optional{T: dt.Prim(rtypes.T_String)},
					"RequestFormat":  dt.Optional{T: dt.Prim(rtypes.T_FormatSelector)},
					"ResponseFormat": dt.Optional{T: dt.Prim(rtypes.T_FormatSelector)},
					"Headers":        dt.Optional{T: dt.Prim(rtypes.T_HttpHeaders)},
					"Cookies":        dt.Optional{T: dt.Prim(rtypes.T_Cookies)},
					"Body":           dt.Optional{T: dt.Prim(rtypes.T_Any)},
				},
				Summary:     "Types/HttpOptions:Summary",
				Description: "Types/HttpOptions:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Cookies,
			FormatSelector,
			HttpHeaders,
			String,
			Variant,
		},
	}
}
