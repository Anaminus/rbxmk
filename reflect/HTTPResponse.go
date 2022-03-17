package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(HTTPResponse) }
func HTTPResponse() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "HTTPResponse",
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			resp, ok := v.(rtypes.HTTPResponse)
			if !ok {
				return nil, rbxmk.TypeError{Want: "HTTPResponse", Got: v.Type()}
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
				return nil, rbxmk.TypeError{Want: "table", Got: lv.Type().String()}
			}
			success, err := c.PullFromDictionary(table, "Success", "bool")
			if err != nil {
				return nil, err
			}
			statusCode, err := c.PullFromDictionary(table, "StatusCode", "int")
			if err != nil {
				return nil, err
			}
			statusMessage, err := c.PullFromDictionary(table, "StatusMessage", "string")
			if err != nil {
				return nil, err
			}
			headers, err := c.PullFromDictionaryOpt(table, "Headers", rtypes.HTTPHeaders(nil), "HTTPHeaders")
			if err != nil {
				return nil, err
			}
			cookies, err := c.PullFromDictionaryOpt(table, "Cookies", rtypes.Cookies(nil), "Cookies")
			if err != nil {
				return nil, err
			}
			body, err := c.PullFromDictionaryOpt(table, "Body", nil, "Variant")
			if err != nil {
				return nil, err
			}
			resp := rtypes.HTTPResponse{
				Success:       bool(success.(types.Bool)),
				StatusCode:    int(statusCode.(types.Int)),
				StatusMessage: string(statusMessage.(types.String)),
				Headers:       headers.(rtypes.HTTPHeaders),
				Cookies:       cookies.(rtypes.Cookies),
				Body:          body,
			}
			return resp, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.HTTPResponse:
				*p = v.(rtypes.HTTPResponse)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"Success":       dt.Prim("bool"),
					"StatusCode":    dt.Prim("int"),
					"StatusMessage": dt.Prim("string"),
					"Headers":       dt.Prim("HTTPHeaders"),
					"Cookies":       dt.Prim("Cookies"),
					"Body":          dt.Optional{T: dt.Prim("any")},
				},
				Summary:     "Types/HTTPResponse:Summary",
				Description: "Types/HTTPResponse:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Bool,
			Cookies,
			HTTPHeaders,
			Int,
			String,
			Variant,
		},
	}
}
