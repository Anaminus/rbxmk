package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(Faces) }
func Faces() Reflector {
	return Reflector{
		Name:     "Faces",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Faces").(types.Faces)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Faces").(types.Faces)
				op := s.Pull(2, "Faces").(types.Faces)
				s.L.Push(lua.LBool(v == op))
				return 1
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
					return s.RaiseError("expected 0 or 6 arguments")
				}
				return s.Push(v)
			},
		},
	}
}
