package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func UDim() Reflector {
	return Reflector{
		Name:     "UDim",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "UDim").(types.UDim)
				return s.Push(types.String(v.String()))
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "UDim").(types.UDim)
				op := s.Pull(2, "UDim").(types.UDim)
				return s.Push(types.Bool(v == op))
			},
			"__add": func(s State) int {
				v := s.Pull(1, "UDim").(types.UDim)
				op := s.Pull(2, "UDim").(types.UDim)
				return s.Push(v.Add(op))
			},
			"__sub": func(s State) int {
				v := s.Pull(1, "UDim").(types.UDim)
				op := s.Pull(2, "UDim").(types.UDim)
				return s.Push(v.Sub(op))
			},
			"__unm": func(s State) int {
				v := s.Pull(1, "UDim").(types.UDim)
				return s.Push(v.Neg())
			},
		},
		Members: map[string]Member{
			"Scale": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.UDim).Scale))
			}},
			"Offset": {Get: func(s State, v types.Value) int {
				return s.Push(types.Int(v.(types.UDim).Offset))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				return s.Push(types.UDim{
					Scale:  float32(s.Pull(1, "float").(types.Float)),
					Offset: int32(s.Pull(2, "int").(types.Int)),
				})
			},
		},
	}
}
