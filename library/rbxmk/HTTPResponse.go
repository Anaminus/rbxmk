package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(HTTPResponse) }
func HTTPResponse() Reflector {
	return Reflector{
		Name: "HTTPResponse",
		PushTo: func(s State, v types.Value) (lvs []lua.LValue, err error) {
			resp, ok := v.(rtypes.HTTPResponse)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "HTTPResponse")
			}
			table := s.L.CreateTable(0, 5)
			s.PushToTable(table, lua.LString("Success"), types.Bool(resp.Success))
			s.PushToTable(table, lua.LString("StatusCode"), types.Int(resp.StatusCode))
			s.PushToTable(table, lua.LString("StatusMessage"), types.String(resp.StatusMessage))
			s.PushToTable(table, lua.LString("Headers"), resp.Headers)
			s.PushToTable(table, lua.LString("Body"), resp.Body)
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s State, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "table")
			}
			resp := rtypes.HTTPResponse{
				Success:       bool(s.PullFromTable(table, lua.LString("Success"), "bool").(types.Bool)),
				StatusCode:    int(s.PullFromTable(table, lua.LString("StatusCode"), "int").(types.Int)),
				StatusMessage: string(s.PullFromTable(table, lua.LString("StatusMessage"), "string").(types.String)),
				Headers:       s.PullFromTableOpt(table, lua.LString("Headers"), "HTTPHeaders", rtypes.HTTPHeaders(nil)).(rtypes.HTTPHeaders),
				Body:          s.PullFromTableOpt(table, lua.LString("Body"), "Variant", nil),
			}
			return resp, nil
		},
	}
}
