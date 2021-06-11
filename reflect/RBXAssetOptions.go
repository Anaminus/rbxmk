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
		PushTo: func(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			options, ok := v.(rtypes.RBXAssetOptions)
			if !ok {
				return nil, rbxmk.TypeError{Want: "RBXAssetOptions", Got: v.Type()}
			}
			if options.AssetID <= 0 {
				return nil, fmt.Errorf("field AssetID (%d) must be greater than 0", options.AssetID)
			}
			table := s.CreateTable(0, 4)
			if err := s.PushToDictionary(table, "AssetID", types.Int64(options.AssetID)); err != nil {
				return nil, err
			}
			if err := s.PushToDictionary(table, "Format", options.Format); err != nil {
				return nil, err
			}
			if err := s.PushToDictionary(table, "Cookies", options.Cookies); err != nil {
				return nil, err
			}
			if err := s.PushToDictionary(table, "Body", options.Body); err != nil {
				return nil, err
			}
			return table, nil
		},
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lv.Type().String()}
			}
			assetID, err := s.PullFromDictionary(table, "AssetID", "int64")
			if err != nil {
				return nil, err
			}
			cookies, err := s.PullFromDictionaryOpt(table, "Cookies", rtypes.Cookies(nil), "Cookies")
			if err != nil {
				return nil, err
			}
			options := rtypes.RBXAssetOptions{
				AssetID: int64(assetID.(types.Int64)),
				Cookies: cookies.(rtypes.Cookies),
			}
			format, err := s.PullFromDictionary(table, "Format", "FormatSelector")
			if err != nil {
				return nil, err
			}
			options.Format = format.(rtypes.FormatSelector)
			if format := s.Format(options.Format.Format); format.Name != "" {
				options.Body, err = s.PullAnyFromDictionaryOpt(table, "Body", nil, format.EncodeTypes...)
			} else {
				options.Body, err = s.PullFromDictionaryOpt(table, "Body", nil, "Variant")
			}
			if err != nil {
				return nil, err
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
