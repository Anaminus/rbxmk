package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(TypeDesc) }
func TypeDesc() Reflector {
	return Reflector{
		Name:     "TypeDesc",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__eq": func(s State) int {
				v := s.Pull(1, "TypeDesc").(rtypes.TypeDesc)
				op := s.Pull(2, "TypeDesc").(rtypes.TypeDesc)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: Members{
			"Category": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.TypeDesc)
					return s.Push(types.String(desc.Embedded.Category))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.TypeDesc)
					desc.Embedded.Category = string(s.Pull(3, "string").(types.String))
				},
			},
			"Name": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.TypeDesc)
					return s.Push(types.String(desc.Embedded.Name))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.TypeDesc)
					desc.Embedded.Name = string(s.Pull(3, "string").(types.String))
				},
			},
		},
	}
}
