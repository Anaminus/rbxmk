package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(UDim) }
func UDim() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "UDim",
		PushTo:   rbxmk.PushTypeTo("UDim"),
		PullFrom: rbxmk.PullTypeFrom("UDim"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "UDim").(types.UDim)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "UDim").(types.UDim)
				op := s.Pull(2, "UDim").(types.UDim)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s rbxmk.State) int {
				v := s.Pull(1, "UDim").(types.UDim)
				op := s.Pull(2, "UDim").(types.UDim)
				return s.Push(v.Add(op))
			},
			"__sub": func(s rbxmk.State) int {
				v := s.Pull(1, "UDim").(types.UDim)
				op := s.Pull(2, "UDim").(types.UDim)
				return s.Push(v.Sub(op))
			},
			"__unm": func(s rbxmk.State) int {
				v := s.Pull(1, "UDim").(types.UDim)
				return s.Push(v.Neg())
			},
		},
		Properties: rbxmk.Properties{
			"Scale": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.UDim).Scale))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("float"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/UDim:Properties/Scale/Summary",
						Description: "Libraries/roblox/Types/UDim:Properties/Scale/Description",
					}
				},
			},
			"Offset": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int(v.(types.UDim).Offset))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("int"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/UDim:Properties/Offset/Summary",
						Description: "Libraries/roblox/Types/UDim:Properties/Offset/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.UDim{
						Scale:  float32(s.Pull(1, "float").(types.Float)),
						Offset: int32(s.Pull(2, "int").(types.Int)),
					})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "scale", Type: dt.Prim("float")},
								{Name: "offset", Type: dt.Prim("int")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("UDim")},
							},
							Summary:     "Libraries/roblox/Types/UDim:Constructors/new/Summary",
							Description: "Libraries/roblox/Types/UDim:Constructors/new/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Libraries/roblox/Types/UDim:Operators/Eq/Summary",
						Description: "Libraries/roblox/Types/UDim:Operators/Eq/Description",
					},
					Add: []dump.Binop{
						{
							Operand:     dt.Prim("UDim"),
							Result:      dt.Prim("UDim"),
							Summary:     "Libraries/roblox/Types/UDim:Operators/Add/Summary",
							Description: "Libraries/roblox/Types/UDim:Operators/Add/Description",
						},
					},
					Sub: []dump.Binop{
						{
							Operand:     dt.Prim("UDim"),
							Result:      dt.Prim("UDim"),
							Summary:     "Libraries/roblox/Types/UDim:Operators/Sub/Summary",
							Description: "Libraries/roblox/Types/UDim:Operators/Sub/Description",
						},
					},
					Unm: &dump.Unop{
						Result:      dt.Prim("UDim"),
						Summary:     "Libraries/roblox/Types/UDim:Operators/Unm/Summary",
						Description: "Libraries/roblox/Types/UDim:Operators/Unm/Description",
					},
				},
				Summary:     "Libraries/roblox/Types/UDim:Summary",
				Description: "Libraries/roblox/Types/UDim:Description",
			}
		},
	}
}
