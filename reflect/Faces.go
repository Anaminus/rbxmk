package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Faces() Type {
	return Type{
		Name:        "Faces",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString(s.Pull(1, "Faces").(types.Faces).String()))
				return 1
			},
			"__eq": func(s State) int {
				op := s.Pull(2, "Faces").(types.Faces)
				return s.Push(types.Bool(s.Pull(1, "Faces").(types.Faces) == op))
			},
		},
		Members: map[string]Member{
			"Right": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Faces).Right))
			}},
			"Top": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Faces).Top))
			}},
			"Back": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Faces).Back))
			}},
			"Left": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Faces).Left))
			}},
			"Bottom": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Faces).Bottom))
			}},
			"Front": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Faces).Front))
			}},
		},
		Constructors: Constructors{
			// TODO: match API.
			"new": func(s State) int {
				var v types.Faces
				switch s.Count() {
				case 6:
					v.Right = bool(s.Pull(1, "bool").(types.Bool))
					v.Top = bool(s.Pull(2, "bool").(types.Bool))
					v.Back = bool(s.Pull(3, "bool").(types.Bool))
					v.Left = bool(s.Pull(4, "bool").(types.Bool))
					v.Bottom = bool(s.Pull(5, "bool").(types.Bool))
					v.Front = bool(s.Pull(6, "bool").(types.Bool))
				default:
					s.L.RaiseError("expected 0 or 6 arguments")
					return 0
				}
				return s.Push(v)
			},
		},
	}
}
