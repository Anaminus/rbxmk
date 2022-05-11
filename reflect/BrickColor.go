package reflect

import (
	"math/rand"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

const T_BrickColor = "BrickColor"

func init() { register(BrickColor) }
func BrickColor() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     T_BrickColor,
		PushTo:   rbxmk.PushTypeTo(T_BrickColor),
		PullFrom: rbxmk.PullTypeFrom(T_BrickColor),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.BrickColor:
				*p = v.(types.BrickColor)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, T_BrickColor).(types.BrickColor)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, T_BrickColor).(types.BrickColor)
				op := s.Pull(2, T_BrickColor).(types.BrickColor)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.String(v.(types.BrickColor).Name()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_String),
						ReadOnly:    true,
						Summary:     "Types/BrickColor:Properties/Name/Summary",
						Description: "Types/BrickColor:Properties/Name/Description",
					}
				},
			},
			"Number": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int(v.(types.BrickColor).Number()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Int),
						ReadOnly:    true,
						Summary:     "Types/BrickColor:Properties/Number/Summary",
						Description: "Types/BrickColor:Properties/Number/Description",
					}
				},
			},
			"R": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.BrickColor).R()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Float),
						ReadOnly:    true,
						Summary:     "Types/BrickColor:Properties/R/Summary",
						Description: "Types/BrickColor:Properties/R/Description",
					}
				},
			},
			"G": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.BrickColor).G()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Float),
						ReadOnly:    true,
						Summary:     "Types/BrickColor:Properties/G/Summary",
						Description: "Types/BrickColor:Properties/G/Description",
					}
				},
			},
			"B": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.BrickColor).B()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Float),
						ReadOnly:    true,
						Summary:     "Types/BrickColor:Properties/B/Summary",
						Description: "Types/BrickColor:Properties/B/Description",
					}
				},
			},
			"Color": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.BrickColor).Color())
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Color3),
						ReadOnly:    true,
						Summary:     "Types/BrickColor:Properties/Color/Summary",
						Description: "Types/BrickColor:Properties/Color/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.BrickColor
					switch s.Count() {
					case 1:
						switch arg := s.PullAnyOf(1, T_Int, T_String, T_Color3).(type) {
						case types.Int:
							v = types.NewBrickColor(int(arg))
						case types.String:
							v = types.NewBrickColorFromName(string(arg))
						case types.Color3:
							v = types.NewBrickColorFromColor3(arg)
						default:
							return s.ReflectorError(1)
						}
					case 3:
						v = types.NewBrickColorFromColor(
							float64(s.Pull(1, T_Float).(types.Float)),
							float64(s.Pull(2, T_Float).(types.Float)),
							float64(s.Pull(3, T_Float).(types.Float)),
						)
					default:
						return s.RaiseError("expected 1 or 3 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "value", Type: dt.Prim(T_Int)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_BrickColor)},
							},
							Summary:     "Types/BrickColor:Constructors/new/Number/Summary",
							Description: "Types/BrickColor:Constructors/new/Number/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "r", Type: dt.Prim(T_Float)},
								{Name: "g", Type: dt.Prim(T_Float)},
								{Name: "b", Type: dt.Prim(T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_BrickColor)},
							},
							Summary:     "Types/BrickColor:Constructors/new/Components/Summary",
							Description: "Types/BrickColor:Constructors/new/Components/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "name", Type: dt.Prim(T_String)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_BrickColor)},
							},
							Summary:     "Types/BrickColor:Constructors/new/Name/Summary",
							Description: "Types/BrickColor:Constructors/new/Name/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "color", Type: dt.Prim(T_Color3)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_BrickColor)},
							},
							Summary:     "Types/BrickColor:Constructors/new/Color/Summary",
							Description: "Types/BrickColor:Constructors/new/Color/Description",
						},
					}
				},
			},
			"palette": {
				Func: func(s rbxmk.State) int {
					index := int(s.Pull(1, T_Int).(types.Int))
					return s.Push(types.NewBrickColorFromPalette(index))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "index", Type: dt.Prim(T_Int)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_BrickColor)},
							},
							Summary:     "Types/BrickColor:Constructors/palette/Summary",
							Description: "Types/BrickColor:Constructors/palette/Description",
						},
					}
				},
			},
			"random": {
				Func: func(s rbxmk.State) int {
					index := rand.Intn(types.BrickColorIndexSize)
					return s.Push(types.NewBrickColorFromIndex(index))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim(T_BrickColor)},
							},
							Summary:     "Types/BrickColor:Constructors/random/Summary",
							Description: "Types/BrickColor:Constructors/random/Description",
						},
					}
				},
			},
			"White": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("White"))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim(T_BrickColor)},
							},
							Summary:     "Types/BrickColor:Constructors/White/Summary",
							Description: "Types/BrickColor:Constructors/White/Description",
						},
					}
				},
			},
			"Gray": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("Medium stone grey"))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim(T_BrickColor)},
							},
							Summary:     "Types/BrickColor:Constructors/Gray/Summary",
							Description: "Types/BrickColor:Constructors/Gray/Description",
						},
					}
				},
			},
			"DarkGray": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("Dark stone grey"))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim(T_BrickColor)},
							},
							Summary:     "Types/BrickColor:Constructors/DarkGray/Summary",
							Description: "Types/BrickColor:Constructors/DarkGray/Description",
						},
					}
				},
			},
			"Black": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("Black"))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim(T_BrickColor)},
							},
							Summary:     "Types/BrickColor:Constructors/Black/Summary",
							Description: "Types/BrickColor:Constructors/Black/Description",
						},
					}
				},
			},
			"Red": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("Bright red"))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim(T_BrickColor)},
							},
							Summary:     "Types/BrickColor:Constructors/Red/Summary",
							Description: "Types/BrickColor:Constructors/Red/Description",
						},
					}
				},
			},
			"Yellow": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("Bright yellow"))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim(T_BrickColor)},
							},
							Summary:     "Types/BrickColor:Constructors/Yellow/Summary",
							Description: "Types/BrickColor:Constructors/Yellow/Description",
						},
					}
				},
			},
			"Green": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("Dark green"))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim(T_BrickColor)},
							},
							Summary:     "Types/BrickColor:Constructors/Green/Summary",
							Description: "Types/BrickColor:Constructors/Green/Description",
						},
					}
				},
			},
			"Blue": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("Bright blue"))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim(T_BrickColor)},
							},
							Summary:     "Types/BrickColor:Constructors/Blue/Summary",
							Description: "Types/BrickColor:Constructors/Blue/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/BrickColor:Operators/Eq/Summary",
						Description: "Types/BrickColor:Operators/Eq/Description",
					},
				},
				Summary:     "Types/BrickColor:Summary",
				Description: "Types/BrickColor:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Color3,
			Float,
			Int,
			String,
		},
	}
}
