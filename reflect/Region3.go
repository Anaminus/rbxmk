package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func Region3() Reflector {
	return Reflector{
		Name:     "Region3",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Region3").(types.Region3)
				return s.Push(types.String(v.String()))
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Region3").(types.Region3)
				op := s.Pull(2, "Region3").(types.Region3)
				return s.Push(types.Bool(v == op))
			},
		},
		Members: map[string]Member{
			"CFrame": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.Region3).CFrame())
			}},
			"Size": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.Region3).Size())
			}},
			"ExpandToGrid": {Method: true, Get: func(s State, v types.Value) int {
				region := int(s.Pull(2, "int").(types.Int))
				return s.Push(v.(types.Region3).ExpandToGrid(region))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				return s.Push(types.Region3{
					Min: s.Pull(1, "Vector3").(types.Vector3),
					Max: s.Pull(2, "Vector3").(types.Vector3),
				})
			},
		},
	}
}
