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

const T_RBXAssetOptions = "RBXAssetOptions"

func init() { register(RBXAssetOptions) }
func RBXAssetOptions() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: T_RBXAssetOptions,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			options, ok := v.(rtypes.RBXAssetOptions)
			if !ok {
				return nil, rbxmk.TypeError{Want: T_RBXAssetOptions, Got: v.Type()}
			}
			if options.AssetID <= 0 {
				return nil, fmt.Errorf("field AssetID (%d) must be greater than 0", options.AssetID)
			}
			table := c.CreateTable(0, 4)
			if err := c.PushToDictionary(table, "AssetID", types.Int64(options.AssetID)); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "Format", options.Format); err != nil {
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
				return nil, rbxmk.TypeError{Want: T_Table, Got: lv.Type().String()}
			}
			assetID, err := c.PullFromDictionary(table, "AssetID", T_Int64)
			if err != nil {
				return nil, err
			}
			cookies, err := c.PullFromDictionaryOpt(table, "Cookies", rtypes.Cookies(nil), T_Cookies)
			if err != nil {
				return nil, err
			}
			options := rtypes.RBXAssetOptions{
				AssetID: int64(assetID.(types.Int64)),
				Cookies: cookies.(rtypes.Cookies),
			}
			format, err := c.PullFromDictionary(table, "Format", T_FormatSelector)
			if err != nil {
				return nil, err
			}
			options.Format = format.(rtypes.FormatSelector)
			options.Body, err = c.PullEncodedFromDict(table, "Body", options.Format)
			if err != nil {
				return nil, err
			}
			if options.AssetID <= 0 {
				return nil, fmt.Errorf("field AssetID (%d) must be greater than 0", options.AssetID)
			}
			return options, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.RBXAssetOptions:
				*p = v.(rtypes.RBXAssetOptions)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"AssetID": dt.Prim(T_Int64),
					"Cookies": dt.Optional{T: dt.Prim(T_Cookies)},
					"Format":  dt.Prim(T_FormatSelector),
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
