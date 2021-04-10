package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(Ray) }
func Ray() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "Ray",
		PushTo:   rbxmk.PushTypeTo("Ray"),
		PullFrom: rbxmk.PullTypeFrom("Ray"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "Ray").(types.Ray)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "Ray").(types.Ray)
				op := s.Pull(2, "Ray").(types.Ray)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Origin": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.Ray).Origin)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("Vector3"),
						ReadOnly:    true,
						Summary:     "Types/Ray:Properties/Origin/Summary",
						Description: "Types/Ray:Properties/Origin/Description",
					}
				},
			},
			"Direction": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.Ray).Direction)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("Vector3"),
						ReadOnly:    true,
						Summary:     "Types/Ray:Properties/Direction/Summary",
						Description: "Types/Ray:Properties/Direction/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"ClosestPoint": {
				Func: func(s rbxmk.State, v types.Value) int {
					point := s.Pull(2, "Vector3").(types.Vector3)
					return s.Push(v.(types.Ray).ClosestPoint(point))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "point", Type: dt.Prim("Vector3")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Vector3")},
						},
						Summary:     "Types/Ray:Methods/ClosestPoint/Summary",
						Description: "Types/Ray:Methods/ClosestPoint/Description",
					}
				},
			},
			"Distance": {
				Func: func(s rbxmk.State, v types.Value) int {
					point := s.Pull(2, "Vector3").(types.Vector3)
					return s.Push(types.Float(v.(types.Ray).Distance(point)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "point", Type: dt.Prim("Vector3")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("float")},
						},
						Summary:     "Types/Ray:Methods/Distance/Summary",
						Description: "Types/Ray:Methods/Distance/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.Ray{
						Origin:    s.Pull(1, "Vector3").(types.Vector3),
						Direction: s.Pull(2, "Vector3").(types.Vector3),
					})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "origin", Type: dt.Prim("Vector3")},
								{Name: "direction", Type: dt.Prim("Vector3")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("Ray")},
							},
							Summary:     "Types/Ray:Constructors/new/Summary",
							Description: "Types/Ray:Constructors/new/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/Ray:Operators/Eq/Summary",
						Description: "Types/Ray:Operators/Eq/Description",
					},
				},
				Summary:     "Types/Ray:Summary",
				Description: "Types/Ray:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Vector3,
		},
	}
}
