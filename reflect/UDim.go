package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func UDim() Reflector {
	return Reflector{
		Name:     "UDim",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString(s.Pull(1, "UDim").(types.UDim).String()))
				return 1
			},
			"__eq": func(s State) int {
				op := s.Pull(2, "UDim").(types.UDim)
				return s.Push(types.Bool(s.Pull(1, "UDim").(types.UDim) == op))
			},
			"__add": func(s State) int {
				op := s.Pull(2, "UDim").(types.UDim)
				return s.Push(s.Pull(1, "UDim").(types.UDim).Add(op))
			},
			"__sub": func(s State) int {
				op := s.Pull(2, "UDim").(types.UDim)
				return s.Push(s.Pull(1, "UDim").(types.UDim).Sub(op))
			},
			"__unm": func(s State) int {
				return s.Push(s.Pull(1, "UDim").(types.UDim).Neg())
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
