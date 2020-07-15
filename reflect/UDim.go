package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func UDim() Type {
	return Type{
		Name:        "UDim",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				s.L.Push(lua.LString(v.(types.UDim).String()))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "UDim").(types.UDim)
				return s.Push("bool", v.(types.UDim) == op)
			},
			"__add": func(s State, v Value) int {
				op := s.Pull(2, "UDim").(types.UDim)
				return s.Push("UDim", v.(types.UDim).Add(op))
			},
			"__sub": func(s State, v Value) int {
				op := s.Pull(2, "UDim").(types.UDim)
				return s.Push("UDim", v.(types.UDim).Sub(op))
			},
			"__unm": func(s State, v Value) int {
				return s.Push("UDim", v.(types.UDim).Neg())
			},
		},
		Members: map[string]Member{
			"Scale": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.UDim).Scale)
			}},
			"Offset": {Get: func(s State, v Value) int {
				return s.Push("int", v.(types.UDim).Offset)
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				return s.Push("UDim", types.UDim{
					Scale:  s.Pull(1, "float").(float32),
					Offset: int32(s.Pull(2, "int").(int)),
				})
			},
		},
	}
}
