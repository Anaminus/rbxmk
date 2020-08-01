package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func Ray() Reflector {
	return Reflector{
		Name:     "Ray",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Ray").(types.Ray)
				return s.Push(types.String(v.String()))
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Ray").(types.Ray)
				op := s.Pull(2, "Ray").(types.Ray)
				return s.Push(types.Bool(v == op))
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
