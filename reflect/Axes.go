package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Axes() Type {
	return Type{
		Name:        "Axes",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString(s.Pull(1, "Axes").(types.Axes).String()))
				return 1
			},
			"__eq": func(s State) int {
				op := s.Pull(2, "Axes").(types.Axes)
				return s.Push("bool", types.Bool(s.Pull(1, "Axes").(types.Axes) == op))
			},
		},
		Members: map[string]Member{
			"X": {Get: func(s State, v types.Value) int {
				return s.Push("bool", types.Bool(v.(types.Axes).X))
			}},
			"Y": {Get: func(s State, v types.Value) int {
				return s.Push("bool", types.Bool(v.(types.Axes).Y))
			}},
			"Z": {Get: func(s State, v types.Value) int {
				return s.Push("bool", types.Bool(v.(types.Axes).Z))
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
					s.L.RaiseError("expected 0 or 3 arguments")
					return 0
				}
				return s.Push("Axes", v)
			},
		},
	}
}
