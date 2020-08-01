package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Axes() Reflector {
	return Reflector{
		Name:     "Axes",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Axes").(types.Axes)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Axes").(types.Axes)
				op := s.Pull(2, "Axes").(types.Axes)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: map[string]Member{
			"X": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Axes).X))
			}},
			"Y": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Axes).Y))
			}},
			"Z": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Axes).Z))
			}},
		},
		Constructors: Constructors{
			// TODO: match API.
			"new": func(s State) int {
				var v types.Axes
				switch s.Count() {
				case 3:
					v.X = bool(s.Pull(1, "bool").(types.Bool))
					v.Y = bool(s.Pull(2, "bool").(types.Bool))
					v.Z = bool(s.Pull(3, "bool").(types.Bool))
				default:
					return s.RaiseError("expected 0 or 3 arguments")
				}
				return s.Push(v)
			},
		},
	}
}
