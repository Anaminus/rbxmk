package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
)

func init() { register(DescAction) }
func DescAction() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "DescAction",
		PushTo:   rbxmk.PushPtrTypeTo("DescAction"),
		PullFrom: rbxmk.PullTypeFrom("DescAction"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "DescAction").(*rtypes.DescAction)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "$TODO",
				Description: "$TODO",
			}
		},
	}
}
