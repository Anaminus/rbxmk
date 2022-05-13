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

func init() { register(RbxAssetOptions) }
func RbxAssetOptions() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_RbxAssetOptions,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			options, ok := v.(rtypes.RbxAssetOptions)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_RbxAssetOptions, Got: v.Type()}
			}
			if options.AssetId <= 0 {
				return nil, fmt.Errorf("field AssetId (%d) must be greater than 0", options.AssetId)
			}
			table := c.CreateTable(0, 4)
			if err := c.PushToDictionary(table, "AssetId", types.Int64(options.AssetId)); err != nil {
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
				return nil, rbxmk.TypeError{Want: rtypes.T_Table, Got: lv.Type().String()}
			}
			assetID, err := c.PullFromDictionary(table, "AssetId", rtypes.T_Int64)
			if err != nil {
				return nil, err
			}
			cookies, err := c.PullFromDictionaryOpt(table, "Cookies", rtypes.Cookies(nil), rtypes.T_Cookies)
			if err != nil {
				return nil, err
			}
			options := rtypes.RbxAssetOptions{
				AssetId: int64(assetID.(types.Int64)),
				Cookies: cookies.(rtypes.Cookies),
			}
			format, err := c.PullFromDictionary(table, "Format", rtypes.T_FormatSelector)
			if err != nil {
				return nil, err
			}
			options.Format = format.(rtypes.FormatSelector)
			options.Body, err = c.PullEncodedFromDict(table, "Body", options.Format)
			if err != nil {
				return nil, err
			}
			if options.AssetId <= 0 {
				return nil, fmt.Errorf("field AssetId (%d) must be greater than 0", options.AssetId)
			}
			return options, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.RbxAssetOptions:
				*p = v.(rtypes.RbxAssetOptions)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"AssetId": dt.Prim(rtypes.T_Int64),
					"Cookies": dt.Optional{T: dt.Prim(rtypes.T_Cookies)},
					"Format":  dt.Prim(rtypes.T_FormatSelector),
					"Body":    dt.Optional{T: dt.Prim(rtypes.T_Any)},
				},
				Summary:     "Types/RbxAssetOptions:Summary",
				Description: "Types/RbxAssetOptions:Description",
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
