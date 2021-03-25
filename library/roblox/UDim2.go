package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(UDim2) }
func UDim2() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "UDim2",
		PushTo:   rbxmk.PushTypeTo("UDim2"),
		PullFrom: rbxmk.PullTypeFrom("UDim2"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "UDim2").(types.UDim2)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "UDim2").(types.UDim2)
				op := s.Pull(2, "UDim2").(types.UDim2)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s rbxmk.State) int {
				v := s.Pull(1, "UDim2").(types.UDim2)
				op := s.Pull(2, "UDim2").(types.UDim2)
				return s.Push(v.Add(op))
			},
			"__sub": func(s rbxmk.State) int {
				v := s.Pull(1, "UDim2").(types.UDim2)
				op := s.Pull(2, "UDim2").(types.UDim2)
				return s.Push(v.Sub(op))
			},
			"__unm": func(s rbxmk.State) int {
				v := s.Pull(1, "UDim2").(types.UDim2)
				return s.Push(v.Neg())
			},
		},
		Properties: rbxmk.Properties{
			"X": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.UDim2).X)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("UDim"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/UDim2:Properties/X/Summary",
						Description: "Libraries/roblox/Types/UDim2:Properties/X/Description",
					}
				},
			},
			"Y": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.UDim2).Y)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("UDim"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/UDim2:Properties/Y/Summary",
						Description: "Libraries/roblox/Types/UDim2:Properties/Y/Description",
					}
				},
			},
			"Width": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.UDim2).X)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("UDim"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/UDim2:Properties/Width/Summary",
						Description: "Libraries/roblox/Types/UDim2:Properties/Width/Description",
					}
				},
			},
			"Height": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.UDim2).Y)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("UDim"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/UDim2:Properties/Height/Summary",
						Description: "Libraries/roblox/Types/UDim2:Properties/Height/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"Lerp": {
				Func: func(s rbxmk.State, v types.Value) int {
					goal := s.Pull(2, "UDim2").(types.UDim2)
					alpha := float64(s.Pull(3, "float").(types.Float))
					return s.Push(v.(types.UDim2).Lerp(goal, alpha))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "goal", Type: dt.Prim("UDim2")},
							{Name: "alpha", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("UDim2")},
						},
						Summary:     "Libraries/roblox/Types/UDim2:Methods/Lerp/Summary",
						Description: "Libraries/roblox/Types/UDim2:Methods/Lerp/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
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
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "xScale", Type: dt.Prim("float")},
								{Name: "xOffset", Type: dt.Prim("int")},
								{Name: "yScale", Type: dt.Prim("float")},
								{Name: "yOffset", Type: dt.Prim("int")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("UDim2")},
							},
							Summary:     "Libraries/roblox/Types/UDim2:Constructors/new/1/Summary",
							Description: "Libraries/roblox/Types/UDim2:Constructors/new/1/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("UDim")},
								{Name: "y", Type: dt.Prim("UDim")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("UDim2")},
							},
							Summary:     "Libraries/roblox/Types/UDim2:Constructors/new/2/Summary",
							Description: "Libraries/roblox/Types/UDim2:Constructors/new/2/Description",
						},
					}
				},
			},
			"fromScale": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.UDim2{
						X: types.UDim{Scale: float32(s.Pull(1, "float").(types.Float))},
						Y: types.UDim{Scale: float32(s.Pull(2, "float").(types.Float))},
					})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("float")},
								{Name: "y", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("UDim2")},
							},
							Summary:     "Libraries/roblox/Types/UDim2:Constructors/fromScale/Summary",
							Description: "Libraries/roblox/Types/UDim2:Constructors/fromScale/Description",
						},
					}
				},
			},
			"fromOffset": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.UDim2{
						X: types.UDim{Offset: int32(s.Pull(1, "int").(types.Int))},
						Y: types.UDim{Offset: int32(s.Pull(2, "int").(types.Int))},
					})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("int")},
								{Name: "y", Type: dt.Prim("int")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("UDim2")},
							},
							Summary:     "Libraries/roblox/Types/UDim2:Constructors/fromOffset/Summary",
							Description: "Libraries/roblox/Types/UDim2:Constructors/fromOffset/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: true,
					Add: []dump.Binop{
						{
							Operand:     dt.Prim("UDim2"),
							Result:      dt.Prim("UDim2"),
							Summary:     "Libraries/roblox/Types/UDim2:Operators/Add/Summary",
							Description: "Libraries/roblox/Types/UDim2:Operators/Add/Description",
						},
					},
					Sub: []dump.Binop{
						{
							Operand:     dt.Prim("UDim2"),
							Result:      dt.Prim("UDim2"),
							Summary:     "Libraries/roblox/Types/UDim2:Operators/Sub/Summary",
							Description: "Libraries/roblox/Types/UDim2:Operators/Sub/Description",
						},
					},
					Unm: &dump.Unop{
						Result:      dt.Prim("UDim2"),
						Summary:     "Libraries/roblox/Types/UDim2:Operators/Unm/Summary",
						Description: "Libraries/roblox/Types/UDim2:Operators/Unm/Description",
					},
				},
				Summary:     "Libraries/roblox/Types/UDim2:Summary",
				Description: "Libraries/roblox/Types/UDim2:Description",
			}
		},
	}
}
