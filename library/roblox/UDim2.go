package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(UDim2) }
func UDim2() Reflector {
	return Reflector{
		Name:     "UDim2",
		PushTo:   rbxmk.PushTypeTo("UDim2"),
		PullFrom: rbxmk.PullTypeFrom("UDim2"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "UDim2").(types.UDim2)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "UDim2").(types.UDim2)
				op := s.Pull(2, "UDim2").(types.UDim2)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s State) int {
				v := s.Pull(1, "UDim2").(types.UDim2)
				op := s.Pull(2, "UDim2").(types.UDim2)
				return s.Push(v.Add(op))
			},
			"__sub": func(s State) int {
				v := s.Pull(1, "UDim2").(types.UDim2)
				op := s.Pull(2, "UDim2").(types.UDim2)
				return s.Push(v.Sub(op))
			},
			"__unm": func(s State) int {
				v := s.Pull(1, "UDim2").(types.UDim2)
				return s.Push(v.Neg())
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
					return s.RaiseError("expected 0 or 3 arguments")
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
