package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Ray() Type {
	return Type{
		Name:     "Ray",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString(s.Pull(1, "Ray").(types.Ray).String()))
				return 1
			},
			"__eq": func(s State) int {
				op := s.Pull(2, "Ray").(types.Ray)
				return s.Push(types.Bool(s.Pull(1, "Ray").(types.Ray) == op))
			},
		},
		Members: map[string]Member{
			"Origin": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.Ray).Origin)
			}},
			"Direction": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.Ray).Direction)
			}},
			"ClosestPoint": {Method: true, Get: func(s State, v types.Value) int {
				point := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(v.(types.Ray).ClosestPoint(point))
			}},
			"Distance": {Method: true, Get: func(s State, v types.Value) int {
				point := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(types.Double(v.(types.Ray).Distance(point)))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				return s.Push(types.Ray{
					Origin:    s.Pull(1, "Vector3").(types.Vector3),
					Direction: s.Pull(2, "Vector3").(types.Vector3),
				})
			},
		},
	}
}
