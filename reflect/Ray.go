package reflect

import (
	"strconv"

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
			"__tostring": func(s State, v Value) int {
				u := v.(types.Ray)
				var b string
				b += strconv.FormatFloat(float64(u.Origin.X), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Origin.Y), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Origin.Z), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Direction.X), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Direction.Y), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Direction.Z), 'g', -1, 32)
				s.L.Push(lua.LString(b))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "Ray").(types.Ray)
				return s.Push("bool", v.(types.Ray) == op)
			},
		},
		Members: map[string]Member{
			"Origin": {Get: func(s State, v Value) int {
				return s.Push("Vector", v.(types.Ray).Origin)
			}},
			"Direction": {Get: func(s State, v Value) int {
				return s.Push("Vector", v.(types.Ray).Direction)
			}},
			"ClosestPoint": {Method: true, Get: func(s State, v Value) int {
				point := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push("Vector3", v.(types.Ray).ClosestPoint(point))
			}},
			"Distance": {Method: true, Get: func(s State, v Value) int {
				point := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push("number", v.(types.Ray).Distance(point))
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
