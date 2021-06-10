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
		PushTo: func(s rbxmk.State, v types.Value) (lv lua.LValue, err error) {
			options, ok := v.(rtypes.RBXAssetOptions)
			if !ok {
				return nil, rbxmk.TypeError{Want: "RBXAssetOptions", Got: v.Type()}
			}
			if options.AssetID <= 0 {
				return nil, fmt.Errorf("field AssetID (%d) must be greater than 0", options.AssetID)
			}
			table := s.L.CreateTable(0, 4)
			s.PushToDictionary(table, "AssetID", types.Int64(options.AssetID))
			s.PushToDictionary(table, "Format", options.Format)
			s.PushToDictionary(table, "Cookies", options.Cookies)
			s.PushToDictionary(table, "Body", options.Body)
			return table, nil
		},
		PullFrom: func(s rbxmk.State, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lv.Type().String()}
			}
			options := rtypes.RBXAssetOptions{
				AssetID: int64(s.PullFromDictionary(table, "AssetID", "int64").(types.Int64)),
				Cookies: s.PullFromDictionaryOpt(table, "Cookies", rtypes.Cookies(nil), "Cookies").(rtypes.Cookies),
			}
			options.Format = s.PullFromDictionary(table, "Format", "FormatSelector").(rtypes.FormatSelector)
			if format := s.Format(options.Format.Format); format.Name != "" {
				options.Body = s.PullAnyFromDictionaryOpt(table, "Body", nil, format.EncodeTypes...)
			} else {
				options.Body = s.PullFromDictionaryOpt(table, "Body", nil, "Variant")
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
