package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(HTTPResponse) }
func HTTPResponse() Reflector {
	return Reflector{
		Name: "HTTPResponse",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			resp, ok := v.(rtypes.HTTPResponse)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "HTTPResponse")
			}
			table := s.L.CreateTable(0, 5)
			s.PushToTable(table, lua.LString("Success"), types.Bool(resp.Success))
			s.PushToTable(table, lua.LString("StatusCode"), types.Int(resp.StatusCode))
			s.PushToTable(table, lua.LString("StatusMessage"), types.String(resp.StatusMessage))
			s.PushToTable(table, lua.LString("Body"), resp.Body)
			table.RawSetString("Headers", pushHTTPHeaders(s, resp.Headers))
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "table")
			}
			resp := rtypes.HTTPResponse{
				Success:       bool(s.PullFromTable(table, lua.LString("Success"), "bool").(types.Bool)),
				StatusCode:    int(s.PullFromTable(table, lua.LString("StatusCode"), "int").(types.Int)),
				StatusMessage: string(s.PullFromTable(table, lua.LString("StatusMessage"), "string").(types.String)),
				Body:          s.PullFromTableOpt(table, lua.LString("Body"), "Variant", nil),
			}
			switch table := table.RawGetString("Headers").(type) {
			case *lua.LNilType:
			case *lua.LTable:
				if resp.Headers, err = pullHTTPHeaders(table); err != nil {
					return nil, err
				}
			default:
				return nil, fmt.Errorf("Headers must be a table")
			}
			return resp, nil
		},
	}
}
