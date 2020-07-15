package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Faces() Type {
	return Type{
		Name:        "Faces",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.(types.Faces); ok {
				return rbxfile.ValueFaces(v), nil
			}
			return nil, TypeError(nil, 0, "Faces")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueFaces); ok {
				return types.Faces(sv), nil
			}
			return nil, TypeError(nil, 0, "Faces")
		},
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				s.L.Push(lua.LString(v.(types.Faces).String()))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "Faces").(types.Faces)
				return s.Push("bool", v.(types.Faces) == op)
			},
		},
		Members: map[string]Member{
			"Right": {Get: func(s State, v Value) int {
				return s.Push("bool", v.(types.Faces).Right)
			}},
			"Top": {Get: func(s State, v Value) int {
				return s.Push("bool", v.(types.Faces).Top)
			}},
			"Back": {Get: func(s State, v Value) int {
				return s.Push("bool", v.(types.Faces).Back)
			}},
			"Left": {Get: func(s State, v Value) int {
				return s.Push("bool", v.(types.Faces).Left)
			}},
			"Bottom": {Get: func(s State, v Value) int {
				return s.Push("bool", v.(types.Faces).Bottom)
			}},
			"Front": {Get: func(s State, v Value) int {
				return s.Push("bool", v.(types.Faces).Front)
			}},
		},
		Constructors: Constructors{
			// TODO: match API.
			"new": func(s State) int {
				var v types.Faces
				switch s.Count() {
				case 6:
					v.Right = s.Pull(1, "bool").(bool)
					v.Top = s.Pull(2, "bool").(bool)
					v.Back = s.Pull(3, "bool").(bool)
					v.Left = s.Pull(4, "bool").(bool)
					v.Bottom = s.Pull(5, "bool").(bool)
					v.Front = s.Pull(6, "bool").(bool)
				default:
					s.L.RaiseError("expected 0 or 6 arguments")
					return 0
				}
				return s.Push("Faces", v)
			},
		},
	}
}
