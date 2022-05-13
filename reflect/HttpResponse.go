package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(HttpResponse) }
func HttpResponse() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_HttpResponse,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			resp, ok := v.(rtypes.HttpResponse)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_HttpResponse, Got: v.Type()}
			}
			table := c.CreateTable(0, 5)
			if err := c.PushToDictionary(table, "Success", types.Bool(resp.Success)); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "StatusCode", types.Int(resp.StatusCode)); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "StatusMessage", types.String(resp.StatusMessage)); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "Headers", resp.Headers); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "Cookies", resp.Cookies); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "Body", resp.Body); err != nil {
				return nil, err
			}
			return table, nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_Table, Got: lv.Type().String()}
			}
			success, err := c.PullFromDictionary(table, "Success", rtypes.T_Bool)
			if err != nil {
				return nil, err
			}
			statusCode, err := c.PullFromDictionary(table, "StatusCode", rtypes.T_Int)
			if err != nil {
				return nil, err
			}
			statusMessage, err := c.PullFromDictionary(table, "StatusMessage", rtypes.T_String)
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
			body, err := c.PullFromDictionaryOpt(table, "Body", nil, rtypes.T_Variant)
			if err != nil {
				return nil, err
			}
			resp := rtypes.HttpResponse{
				Success:       bool(success.(types.Bool)),
				StatusCode:    int(statusCode.(types.Int)),
				StatusMessage: string(statusMessage.(types.String)),
				Headers:       headers.(rtypes.HttpHeaders),
				Cookies:       cookies.(rtypes.Cookies),
				Body:          body,
			}
			return resp, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.HttpResponse:
				*p = v.(rtypes.HttpResponse)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"Success":       dt.Prim(rtypes.T_Bool),
					"StatusCode":    dt.Prim(rtypes.T_Int),
					"StatusMessage": dt.Prim(rtypes.T_String),
					"Headers":       dt.Prim(rtypes.T_HttpHeaders),
					"Cookies":       dt.Prim(rtypes.T_Cookies),
					"Body":          dt.Optional{T: dt.Prim(rtypes.T_Variant)},
				},
				Summary:     "Types/HttpResponse:Summary",
				Description: "Types/HttpResponse:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Bool,
			Cookies,
			HttpHeaders,
			Int,
			String,
			Variant,
		},
	}
}
