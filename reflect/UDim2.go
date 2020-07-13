package reflect

import (
	"strconv"

	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func UDim2() Type {
	return Type{
		Name:        "UDim2",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				u := v.(types.UDim2)
				var b string
				b += strconv.FormatFloat(float64(u.X.Scale), 'g', -1, 32) + ", "
				b += strconv.FormatInt(int64(u.X.Offset), 10) + ", "
				b += strconv.FormatFloat(float64(u.Y.Scale), 'g', -1, 32) + ", "
				b += strconv.FormatInt(int64(u.Y.Offset), 10)
				s.L.Push(lua.LString(b))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "UDim2").(types.UDim2)
				return s.Push("bool", v.(types.UDim2) == op)
			},
			"__add": func(s State, v Value) int {
				op := s.Pull(2, "UDim2").(types.UDim2)
				return s.Push("UDim2", v.(types.UDim2).Add(op))
			},
			"__sub": func(s State, v Value) int {
				op := s.Pull(2, "UDim2").(types.UDim2)
				return s.Push("UDim2", v.(types.UDim2).Sub(op))
			},
			"__unm": func(s State, v Value) int {
				return s.Push("UDim2", v.(types.UDim2).Neg())
			},
		},
		Members: map[string]Member{
			"X": {Get: func(s State, v Value) int {
				return s.Push("UDim", v.(types.UDim2).X)
			}},
			"Y": {Get: func(s State, v Value) int {
				return s.Push("UDim", v.(types.UDim2).Y)
			}},
			"Width": {Get: func(s State, v Value) int {
				return s.Push("UDim", v.(types.UDim2).X)
			}},
			"Height": {Get: func(s State, v Value) int {
				return s.Push("UDim", v.(types.UDim2).Y)
			}},
			"Lerp": {Method: true, Get: func(s State, v Value) int {
				goal := s.Pull(2, "UDim2").(types.UDim2)
				alpha := s.Pull(3, "double").(float64)
				return s.Push("UDim2", v.(types.UDim2).Lerp(goal, alpha))
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
					v.X.Scale = s.Pull(1, "float").(float32)
					v.X.Offset = int32(s.Pull(2, "int").(int))
					v.Y.Scale = s.Pull(3, "float").(float32)
					v.Y.Offset = int32(s.Pull(4, "int").(int))
				default:
					s.L.RaiseError("expected 0 or 3 arguments")
					return 0
				}
				return s.Push("UDim2", v)
			},
			"fromScale": func(s State) int {
				return s.Push("UDim2", types.UDim2{
					X: types.UDim{Scale: s.Pull(1, "float").(float32)},
					Y: types.UDim{Scale: s.Pull(2, "float").(float32)},
				})
			},
			"fromOffset": func(s State) int {
				return s.Push("UDim2", types.UDim2{
					X: types.UDim{Offset: int32(s.Pull(1, "int").(int))},
					Y: types.UDim{Offset: int32(s.Pull(2, "int").(int))},
				})
			},
		},
	}
}
