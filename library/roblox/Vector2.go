package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(Vector2) }
func Vector2() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "Vector2",
		PushTo:   rbxmk.PushTypeTo("Vector2"),
		PullFrom: rbxmk.PullTypeFrom("Vector2"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector2").(types.Vector2)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector2").(types.Vector2)
				op := s.Pull(2, "Vector2").(types.Vector2)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector2").(types.Vector2)
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push(v.Add(op))
			},
			"__sub": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector2").(types.Vector2)
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push(v.Sub(op))
			},
			"__mul": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector2").(types.Vector2)
				switch op := s.PullAnyOf(2, "number", "Vector2").(type) {
				case types.Double:
					return s.Push(v.MulN(float64(op)))
				case types.Vector2:
					return s.Push(v.Mul(op))
				default:
					return s.ReflectorError(2)
				}
			},
			"__div": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector2").(types.Vector2)
				switch op := s.PullAnyOf(2, "number", "Vector2").(type) {
				case types.Double:
					return s.Push(v.DivN(float64(op)))
				case types.Vector2:
					return s.Push(v.Div(op))
				default:
					return s.ReflectorError(2)
				}
			},
			"__unm": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector2").(types.Vector2)
				return s.Push(v.Neg())
			},
		},
		Properties: rbxmk.Properties{
			"X": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Vector2).X))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("float"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/Vector2:Properties/X/Summary",
						Description: "Libraries/roblox/Types/Vector2:Properties/X/Description",
					}
				},
			},
			"Y": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Vector2).Y))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("float"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/Vector2:Properties/Y/Summary",
						Description: "Libraries/roblox/Types/Vector2:Properties/Y/Description",
					}
				},
			},
			"Magnitude": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Vector2).Magnitude()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("float"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/Vector2:Properties/Magnitude/Summary",
						Description: "Libraries/roblox/Types/Vector2:Properties/Magnitude/Description",
					}
				},
			},
			"Unit": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.Vector2).Unit())
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("Vector2"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/Vector2:Properties/Unit/Summary",
						Description: "Libraries/roblox/Types/Vector2:Properties/Unit/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"Lerp": {
				Func: func(s rbxmk.State, v types.Value) int {
					goal := s.Pull(2, "Vector2").(types.Vector2)
					alpha := float64(s.Pull(3, "float").(types.Float))
					return s.Push(v.(types.Vector2).Lerp(goal, alpha))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "goal", Type: dt.Prim("Vector2")},
							{Name: "alpha", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Vector2")},
						},
						Summary:     "Libraries/roblox/Types/Vector2:Methods/Lerp/Summary",
						Description: "Libraries/roblox/Types/Vector2:Methods/Lerp/Description",
					}
				},
			},
			"Dot": {
				Func: func(s rbxmk.State, v types.Value) int {
					op := s.Pull(2, "Vector2").(types.Vector2)
					return s.Push(types.Double(v.(types.Vector2).Dot(op)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "op", Type: dt.Prim("Vector2")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("float")},
						},
						Summary:     "Libraries/roblox/Types/Vector2:Methods/Dot/Summary",
						Description: "Libraries/roblox/Types/Vector2:Methods/Dot/Description",
					}
				},
			},
			"Cross": {
				Func: func(s rbxmk.State, v types.Value) int {
					op := s.Pull(2, "Vector2").(types.Vector2)
					return s.Push(types.Double(v.(types.Vector2).Cross(op)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "op", Type: dt.Prim("Vector2")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("float")},
						},
						Summary:     "Libraries/roblox/Types/Vector2:Methods/Cross/Summary",
						Description: "Libraries/roblox/Types/Vector2:Methods/Cross/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.Vector2
					switch s.Count() {
					case 0:
					case 2:
						v.X = float32(s.Pull(1, "float").(types.Float))
						v.Y = float32(s.Pull(2, "float").(types.Float))
					default:
						return s.RaiseError("expected 0 or 2 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim("Vector2")},
							},
							Summary:     "Libraries/roblox/Types/Vector2:Constructors/new/1/Summary",
							Description: "Libraries/roblox/Types/Vector2:Constructors/new/1/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("float")},
								{Name: "y", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("Vector2")},
							},
							Summary:     "Libraries/roblox/Types/Vector2:Constructors/new/2/Summary",
							Description: "Libraries/roblox/Types/Vector2:Constructors/new/2/Description",
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
							Operand:     dt.Prim("Vector2"),
							Result:      dt.Prim("Vector2"),
							Summary:     "Libraries/roblox/Types/Vector2:Operators/Add/Summary",
							Description: "Libraries/roblox/Types/Vector2:Operators/Add/Description",
						},
					},
					Sub: []dump.Binop{
						{
							Operand:     dt.Prim("Vector2"),
							Result:      dt.Prim("Vector2"),
							Summary:     "Libraries/roblox/Types/Vector2:Operators/Sub/Summary",
							Description: "Libraries/roblox/Types/Vector2:Operators/Sub/Description",
						},
					},
					Mul: []dump.Binop{
						{
							Operand:     dt.Prim("Vector2"),
							Result:      dt.Prim("Vector2"),
							Summary:     "Libraries/roblox/Types/Vector2:Operators/Mul/1/Summary",
							Description: "Libraries/roblox/Types/Vector2:Operators/Mul/1/Description",
						},
						{
							Operand:     dt.Prim("number"),
							Result:      dt.Prim("Vector2"),
							Summary:     "Libraries/roblox/Types/Vector2:Operators/Mul/2/Summary",
							Description: "Libraries/roblox/Types/Vector2:Operators/Mul/2/Description",
						},
					},
					Div: []dump.Binop{
						{
							Operand:     dt.Prim("Vector2"),
							Result:      dt.Prim("Vector2"),
							Summary:     "Libraries/roblox/Types/Vector2:Operators/Div/1/Summary",
							Description: "Libraries/roblox/Types/Vector2:Operators/Div/1/Description",
						},
						{
							Operand:     dt.Prim("number"),
							Result:      dt.Prim("Vector2"),
							Summary:     "Libraries/roblox/Types/Vector2:Operators/Div/2/Summary",
							Description: "Libraries/roblox/Types/Vector2:Operators/Div/2/Description",
						},
					},
					Unm: &dump.Unop{
						Result:      dt.Prim("Vector2"),
						Summary:     "Libraries/roblox/Types/Vector2:Operators/Unm/Summary",
						Description: "Libraries/roblox/Types/Vector2:Operators/Unm/Description",
					},
				},
				Summary:     "Libraries/roblox/Types/Vector2:Summary",
				Description: "Libraries/roblox/Types/Vector2:Description",
			}
		},
	}
}
