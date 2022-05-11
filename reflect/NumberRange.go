package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

const T_NumberRange = "NumberRange"

func init() { register(NumberRange) }
func NumberRange() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     T_NumberRange,
		PushTo:   rbxmk.PushTypeTo(T_NumberRange),
		PullFrom: rbxmk.PullTypeFrom(T_NumberRange),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.NumberRange:
				*p = v.(types.NumberRange)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, T_NumberRange).(types.NumberRange)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, T_NumberRange).(types.NumberRange)
				op := s.Pull(2, T_NumberRange).(types.NumberRange)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Min": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.NumberRange).Min))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Float),
						ReadOnly:    true,
						Summary:     "Types/NumberRange:Properties/Min/Summary",
						Description: "Types/NumberRange:Properties/Min/Description",
					}
				},
			},
			"Max": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.NumberRange).Max))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Float),
						ReadOnly:    true,
						Summary:     "Types/NumberRange:Properties/Max/Summary",
						Description: "Types/NumberRange:Properties/Max/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.NumberRange
					switch s.Count() {
					case 1:
						v.Min = float32(s.Pull(1, T_Float).(types.Float))
						v.Max = v.Min
					case 2:
						v.Min = float32(s.Pull(1, T_Float).(types.Float))
						v.Max = float32(s.Pull(2, T_Float).(types.Float))
						if v.Min > v.Max {
							return s.RaiseError("invalid range")
						}
					default:
						return s.RaiseError("expected 1 or 2 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "value", Type: dt.Prim(T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_NumberRange)},
							},
							Summary:     "Types/NumberRange:Constructors/new/Single/Summary",
							Description: "Types/NumberRange:Constructors/new/Single/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "minimum", Type: dt.Prim(T_Float)},
								{Name: "maxmimum", Type: dt.Prim(T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_NumberRange)},
							},
							Summary:     "Types/NumberRange:Constructors/new/Range/Summary",
							Description: "Types/NumberRange:Constructors/new/Range/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/NumberRange:Operators/Eq/Summary",
						Description: "Types/NumberRange:Operators/Eq/Description",
					},
				},
				Summary:     "Types/NumberRange:Summary",
				Description: "Types/NumberRange:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Float,
		},
	}
}
