package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(Ray) }
func Ray() Reflector {
	return Reflector{
		Name:     "Ray",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Ray").(types.Ray)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Ray").(types.Ray)
				op := s.Pull(2, "Ray").(types.Ray)
				s.L.Push(lua.LBool(v == op))
				return 1
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
