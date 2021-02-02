package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(RBXWebOptions) }
func RBXWebOptions() Reflector {
	return Reflector{
		Name: "RBXWebOptions",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			options, ok := v.(rtypes.RBXWebOptions)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "RBXWebOptions")
			}
			if options.AssetID <= 0 {
				return nil, fmt.Errorf("field AssetID (%d) must be greater than 0", options.AssetID)
			}
			table := s.L.CreateTable(0, 4)
			s.PushToTable(table, lua.LString("AssetID"), types.Int64(options.AssetID))
			s.PushToTable(table, lua.LString("Format"), options.Format)
			s.PushToTable(table, lua.LString("Body"), options.Body)
			cookies := s.L.CreateTable(len(options.Cookies), 0)
			for _, cookie := range options.Cookies {
				cookies.Append(lua.LString(cookie))
			}
			table.RawSetString("Cookies", cookies)
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "table")
			}
			options := rtypes.RBXWebOptions{
				AssetID: int64(s.PullFromTable(table, lua.LString("AssetID"), "int64").(types.Int64)),
				Format:  s.PullFromTable(table, lua.LString("Format"), "FormatSelector").(rtypes.FormatSelector),
				Body:    s.PullFromTableOpt(table, lua.LString("Body"), "Variant", nil),
			}
			if options.AssetID <= 0 {
				return nil, fmt.Errorf("field AssetID (%d) must be greater than 0", options.AssetID)
			}
			if options.Cookies, err = pullStringArray(table.RawGetString("Cookies")); err != nil {
				return nil, fmt.Errorf("field Cookies: %w", err)
			}
			return options, nil
		},
	}
}
