package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Color3) }
func Color3() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "Color3",
		PushTo:   rbxmk.PushTypeTo("Color3"),
		PullFrom: rbxmk.PullTypeFrom("Color3"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "Color3").(types.Color3)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "Color3").(types.Color3)
				op := s.Pull(2, "Color3").(types.Color3)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"R": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Color3).R))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("float"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/Color3:Properties/R/Summary",
						Description: "Libraries/roblox/Types/Color3:Properties/R/Description",
					}
				},
			},
			"G": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Color3).G))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("float"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/Color3:Properties/G/Summary",
						Description: "Libraries/roblox/Types/Color3:Properties/G/Description",
					}
				},
			},
			"B": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Color3).B))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("float"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/Color3:Properties/B/Summary",
						Description: "Libraries/roblox/Types/Color3:Properties/B/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"Lerp": {
				Func: func(s rbxmk.State, v types.Value) int {
					goal := s.Pull(2, "Color3").(types.Color3)
					alpha := float64(s.Pull(3, "float").(types.Float))
					return s.Push(v.(types.Color3).Lerp(goal, alpha))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "goal", Type: dt.Prim("Color3")},
							{Name: "alpha", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Color3")},
						},
						Summary:     "Libraries/roblox/Types/Color3:Methods/Lerp/Summary",
						Description: "Libraries/roblox/Types/Color3:Methods/Lerp/Description",
					}
				},
			},
			"ToHSV": {
				Func: func(s rbxmk.State, v types.Value) int {
					hue, sat, val := v.(types.Color3).ToHSV()
					return s.Push(rtypes.Tuple{types.Float(hue), types.Float(sat), types.Float(val)})
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "h", Type: dt.Prim("float")},
							{Name: "s", Type: dt.Prim("float")},
							{Name: "v", Type: dt.Prim("float")},
						},
						Summary:     "Libraries/roblox/Types/Color3:Methods/ToHSV/Summary",
						Description: "Libraries/roblox/Types/Color3:Methods/ToHSV/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.Color3
					switch s.Count() {
					case 0:
					case 3:
						v.R = float32(s.Pull(1, "float").(types.Float))
						v.G = float32(s.Pull(2, "float").(types.Float))
						v.B = float32(s.Pull(3, "float").(types.Float))
					default:
						return s.RaiseError("expected 0 or 3 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim("Color3")},
							},
							Summary:     "Libraries/roblox/Types/Color3:Constructors/new/Zero/Summary",
							Description: "Libraries/roblox/Types/Color3:Constructors/new/Zero/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "r", Type: dt.Prim("float")},
								{Name: "g", Type: dt.Prim("float")},
								{Name: "b", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("Color3")},
							},
							Summary:     "Libraries/roblox/Types/Color3:Constructors/new/Components/Summary",
							Description: "Libraries/roblox/Types/Color3:Constructors/new/Components/Description",
						},
					}
				},
			},
			"fromRGB": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewColor3FromRGB(
						int(s.Pull(1, "int").(types.Int)),
						int(s.Pull(2, "int").(types.Int)),
						int(s.Pull(3, "int").(types.Int)),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "r", Type: dt.Prim("int")},
								{Name: "g", Type: dt.Prim("int")},
								{Name: "b", Type: dt.Prim("int")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("Color3")},
							},
							Summary:     "Libraries/roblox/Types/Color3:Constructors/fromRGB/Summary",
							Description: "Libraries/roblox/Types/Color3:Constructors/fromRGB/Description",
						},
					}
				},
			},
			"fromHSV": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewColor3FromHSV(
						float64(s.Pull(1, "float").(types.Float)),
						float64(s.Pull(2, "float").(types.Float)),
						float64(s.Pull(3, "float").(types.Float)),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "h", Type: dt.Prim("float")},
								{Name: "s", Type: dt.Prim("float")},
								{Name: "v", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("Color3")},
							},
							Summary:     "Libraries/roblox/Types/Color3:Constructors/fromHSV/Summary",
							Description: "Libraries/roblox/Types/Color3:Constructors/fromHSV/Description",
						},
					}
				},
			},
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.Color3:
				return v
			case rtypes.Color3uint8:
				return types.Color3(v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators:   &dump.Operators{Eq: true},
				Summary:     "Libraries/roblox/Types/Color3:Summary",
				Description: "Libraries/roblox/Types/Color3:Description",
			}
		},
	}
}
