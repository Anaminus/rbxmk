package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func UDim2() Type {
	return Type{
		Name:     "UDim2",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString(s.Pull(1, "UDim2").(types.UDim2).String()))
				return 1
			},
			"__eq": func(s State) int {
				op := s.Pull(2, "UDim2").(types.UDim2)
				return s.Push(types.Bool(s.Pull(1, "UDim2").(types.UDim2) == op))
			},
			"__add": func(s State) int {
				op := s.Pull(2, "UDim2").(types.UDim2)
				return s.Push(s.Pull(1, "UDim2").(types.UDim2).Add(op))
			},
			"__sub": func(s State) int {
				op := s.Pull(2, "UDim2").(types.UDim2)
				return s.Push(s.Pull(1, "UDim2").(types.UDim2).Sub(op))
			},
			"__unm": func(s State) int {
				return s.Push(s.Pull(1, "UDim2").(types.UDim2).Neg())
			},
		},
		Members: map[string]Member{
			"X": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.UDim2).X)
			}},
			"Y": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.UDim2).Y)
			}},
			"Width": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.UDim2).X)
			}},
			"Height": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.UDim2).Y)
			}},
			"Lerp": {Method: true, Get: func(s State, v types.Value) int {
				goal := s.Pull(2, "UDim2").(types.UDim2)
				alpha := float64(s.Pull(3, "number").(types.Double))
				return s.Push(v.(types.UDim2).Lerp(goal, alpha))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.UDim2
				switch s.Count() {
				case 2:
					v.X = s.Pull(1, "UDim").(types.UDim)
					v.Y = s.Pull(2, "UDim").(types.UDim)
				case 4:
					v.X.Scale = float32(s.Pull(1, "float").(types.Float))
					v.X.Offset = int32(s.Pull(2, "int").(types.Int))
					v.Y.Scale = float32(s.Pull(3, "float").(types.Float))
					v.Y.Offset = int32(s.Pull(4, "int").(types.Int))
				default:
					s.L.RaiseError("expected 0 or 3 arguments")
					return 0
				}
				return s.Push(v)
			},
			"fromScale": func(s State) int {
				return s.Push(types.UDim2{
					X: types.UDim{Scale: float32(s.Pull(1, "float").(types.Float))},
					Y: types.UDim{Scale: float32(s.Pull(2, "float").(types.Float))},
				})
			},
			"fromOffset": func(s State) int {
				return s.Push(types.UDim2{
					X: types.UDim{Offset: int32(s.Pull(1, "int").(types.Int))},
					Y: types.UDim{Offset: int32(s.Pull(2, "int").(types.Int))},
				})
			},
		},
	}
}
