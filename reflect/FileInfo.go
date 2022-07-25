package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(FileInfo) }
func FileInfo() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_FileInfo,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			options, ok := v.(rtypes.FileInfo)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_FileInfo, Got: v.Type()}
			}
			table := c.CreateTable(0, 4)
			if err := c.PushToDictionary(table, "Name", types.String(options.Name())); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "IsDir", types.Bool(options.IsDir())); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "Size", types.Int64(options.Size())); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "ModTime", types.Int64(options.ModTime().Unix())); err != nil {
				return nil, err
			}
			return table, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.FileInfo:
				*p = v.(rtypes.FileInfo)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category: "rbxmk",
				Underlying: dt.Struct{
					"Name":    dt.Prim(rtypes.T_String),
					"IsDir":   dt.Prim(rtypes.T_Bool),
					"Size":    dt.Prim(rtypes.T_Int64),
					"ModTime": dt.Prim(rtypes.T_Int64),
				},
				Summary:     "Types/FileInfo:Summary",
				Description: "Types/FileInfo:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Bool,
			Int64,
			String,
		},
	}
}
