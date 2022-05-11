package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

const T_Region3 = "Region3"

func init() { register(Region3) }
func Region3() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     T_Region3,
		PushTo:   rbxmk.PushTypeTo(T_Region3),
		PullFrom: rbxmk.PullTypeFrom(T_Region3),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Region3:
				*p = v.(types.Region3)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, T_Region3).(types.Region3)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, T_Region3).(types.Region3)
				op := s.Pull(2, T_Region3).(types.Region3)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"CFrame": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.Region3).CFrame())
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_CFrame),
						ReadOnly:    true,
						Summary:     "Types/Region3:Properties/CFrame/Summary",
						Description: "Types/Region3:Properties/CFrame/Description",
					}
				},
			},
			"Size": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.Region3).Size())
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Vector3),
						ReadOnly:    true,
						Summary:     "Types/Region3:Properties/Size/Summary",
						Description: "Types/Region3:Properties/Size/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"ExpandToGrid": {
				Func: func(s rbxmk.State, v types.Value) int {
					region := float64(s.Pull(2, T_Float).(types.Float))
					return s.Push(v.(types.Region3).ExpandToGrid(region))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "resolution", Type: dt.Prim(T_Float)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(T_Region3)},
						},
						Summary:     "Types/Region3:Methods/ExpandToGrid/Summary",
						Description: "Types/Region3:Methods/ExpandToGrid/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.Region3{
						Min: s.Pull(1, T_Vector3).(types.Vector3),
						Max: s.Pull(2, T_Vector3).(types.Vector3),
					})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "min", Type: dt.Prim(T_Vector3)},
								{Name: "max", Type: dt.Prim(T_Vector3)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_Region3)},
							},
							Summary:     "Types/Region3:Constructors/new/Summary",
							Description: "Types/Region3:Constructors/new/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/Region3:Operators/Eq/Summary",
						Description: "Types/Region3:Operators/Eq/Description",
					},
				},
				Summary:     "Types/Region3:Summary",
				Description: "Types/Region3:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			CFrame,
			Float,
			Vector3,
		},
	}
}
