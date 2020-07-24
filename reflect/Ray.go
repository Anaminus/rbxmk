package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Ray() Type {
	return Type{
		Name:        "Ray",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v types.Value) int {
				s.L.Push(lua.LString(v.(types.Ray).String()))
				return 1
			},
			"__eq": func(s State, v types.Value) int {
				op := s.Pull(2, "Ray").(types.Ray)
				return s.Push("bool", types.Bool(v.(types.Ray) == op))
			},
		},
		Members: map[string]Member{
			"Origin": {Get: func(s State, v types.Value) int {
				return s.Push("Vector", v.(types.Ray).Origin)
			}},
			"Direction": {Get: func(s State, v types.Value) int {
				return s.Push("Vector", v.(types.Ray).Direction)
			}},
			"ClosestPoint": {Method: true, Get: func(s State, v types.Value) int {
				point := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push("Vector3", v.(types.Ray).ClosestPoint(point))
			}},
			"Distance": {Method: true, Get: func(s State, v types.Value) int {
				point := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push("number", types.Double(v.(types.Ray).Distance(point)))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				return s.Push("Ray", types.Ray{
					Origin:    s.Pull(1, "Vector3").(types.Vector3),
					Direction: s.Pull(2, "Vector3").(types.Vector3),
				})
			},
		},
	}
}
