package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Region3int16() Reflector {
	return Reflector{
		Name:     "Region3int16",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Region3int16").(types.Region3int16)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Region3int16").(types.Region3int16)
				op := s.Pull(2, "Region3int16").(types.Region3int16)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: map[string]Member{
			"Min": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.Region3int16).Min)
			}},
			"Max": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.Region3int16).Max)
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				return s.Push(types.Region3int16{
					Min: s.Pull(1, "Vector3int16").(types.Vector3int16),
					Max: s.Pull(2, "Vector3int16").(types.Vector3int16),
				})
			},
		},
	}
}
