package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(RBXAssetOptions) }
func RBXAssetOptions() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "RBXAssetOptions",
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			options, ok := v.(rtypes.RBXAssetOptions)
			if !ok {
				return nil, rbxmk.TypeError{Want: "RBXAssetOptions", Got: v.Type()}
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
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lvs[0].Type().String()}
			}
			options := rtypes.RBXAssetOptions{
				AssetID: int64(s.PullFromTable(table, lua.LString("AssetID"), "int64").(types.Int64)),
				Cookies: s.PullFromTableOpt(table, lua.LString("Cookies"), "Cookies", rtypes.Cookies(nil)).(rtypes.Cookies),
			}
			options.Format = s.PullFromTable(table, lua.LString("Format"), "FormatSelector").(rtypes.FormatSelector)
			if format := s.Format(options.Format.Format); format.Name != "" {
				options.Body = s.PullAnyFromTableOpt(table, lua.LString("Body"), nil, format.EncodeTypes...)
			} else {
				options.Body = s.PullFromTableOpt(table, lua.LString("Body"), "Variant", nil)
			}
			if options.AssetID <= 0 {
				return nil, fmt.Errorf("field AssetID (%d) must be greater than 0", options.AssetID)
			}
			return options, nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"AssetID": dt.Prim("int64"),
					"Cookies": dt.Optional{T: dt.Prim("Cookies")},
					"Format":  dt.Prim("FormatSelector"),
					"Body":    dt.Optional{T: dt.Prim("any")},
				},
				Summary:     "Types/RBXAssetOptions:Summary",
				Description: "Types/RBXAssetOptions:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Cookies,
			FormatSelector,
			Int64,
			Variant,
		},
	}
}
