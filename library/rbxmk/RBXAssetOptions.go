package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(RBXAssetOptions) }
func RBXAssetOptions() Reflector {
	return Reflector{
		Name: "RBXAssetOptions",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			options, ok := v.(rtypes.RBXAssetOptions)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "RBXAssetOptions")
			}
			if options.AssetID <= 0 {
				return nil, fmt.Errorf("field AssetID (%d) must be greater than 0", options.AssetID)
			}
			table := s.L.CreateTable(0, 4)
			s.PushToTable(table, lua.LString("AssetID"), types.Int64(options.AssetID))
			s.PushToTable(table, lua.LString("Format"), options.Format)
			s.PushToTable(table, lua.LString("Cookies"), options.Cookies)
			s.PushToTable(table, lua.LString("Body"), options.Body)
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "table")
			}
			options := rtypes.RBXAssetOptions{
				AssetID: int64(s.PullFromTable(table, lua.LString("AssetID"), "int64").(types.Int64)),
				Format:  s.PullFromTable(table, lua.LString("Format"), "FormatSelector").(rtypes.FormatSelector),
				Cookies: s.PullFromTable(table, lua.LString("Cookies"), "Cookies").(rtypes.Cookies),
				Body:    s.PullFromTableOpt(table, lua.LString("Body"), "Variant", nil),
			}
			if options.AssetID <= 0 {
				return nil, fmt.Errorf("field AssetID (%d) must be greater than 0", options.AssetID)
			}
			return options, nil
		},
	}
}
