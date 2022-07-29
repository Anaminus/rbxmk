package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(DirEntry) }
func DirEntry() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_DirEntry,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			options, ok := v.(rtypes.DirEntry)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_DirEntry, Got: v.Type()}
			}
			table := c.CreateTable(0, 2)
			if err := c.PushToDictionary(table, "Name", types.String(options.Name())); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "IsDir", types.Bool(options.IsDir())); err != nil {
				return nil, err
			}
			return table, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.DirEntry:
				*p = v.(rtypes.DirEntry)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category: "rbxmk",
				Underlying: dt.P(dt.Struct(dt.KindStruct{
					"Name":  dt.Prim(rtypes.T_String),
					"IsDir": dt.Prim(rtypes.T_Bool),
				})),
				Summary:     "Types/DirEntry:Summary",
				Description: "Types/DirEntry:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Bool,
			String,
		},
	}
}
