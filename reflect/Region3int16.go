package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

const T_Region3int16 = "Region3int16"

func init() { register(Region3int16) }
func Region3int16() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     T_Region3int16,
		PushTo:   rbxmk.PushTypeTo(T_Region3int16),
		PullFrom: rbxmk.PullTypeFrom(T_Region3int16),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Region3int16:
				*p = v.(types.Region3int16)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, T_Region3int16).(types.Region3int16)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, T_Region3int16).(types.Region3int16)
				op := s.Pull(2, T_Region3int16).(types.Region3int16)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Min": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.Region3int16).Min)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Vector3int16),
						ReadOnly:    true,
						Summary:     "Types/Region3int16:Properties/Min/Summary",
						Description: "Types/Region3int16:Properties/Min/Description",
					}
				},
			},
			"Max": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.Region3int16).Max)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Vector3int16),
						ReadOnly:    true,
						Summary:     "Types/Region3int16:Properties/Max/Summary",
						Description: "Types/Region3int16:Properties/Max/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.Region3int16{
						Min: s.Pull(1, T_Vector3int16).(types.Vector3int16),
						Max: s.Pull(2, T_Vector3int16).(types.Vector3int16),
					})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "min", Type: dt.Prim(T_Vector3int16)},
								{Name: "max", Type: dt.Prim(T_Vector3int16)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_Region3int16)},
							},
							Summary:     "Types/Region3int16:Constructors/new/Summary",
							Description: "Types/Region3int16:Constructors/new/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/Region3int16:Operators/Eq/Summary",
						Description: "Types/Region3int16:Operators/Eq/Description",
					},
				},
				Summary:     "Types/Region3int16:Summary",
				Description: "Types/Region3int16:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Vector3int16,
		},
	}
}
