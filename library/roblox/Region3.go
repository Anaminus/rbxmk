package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(Region3) }
func Region3() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "Region3",
		PushTo:   rbxmk.PushTypeTo("Region3"),
		PullFrom: rbxmk.PullTypeFrom("Region3"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "Region3").(types.Region3)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "Region3").(types.Region3)
				op := s.Pull(2, "Region3").(types.Region3)
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
						ValueType:   dt.Prim("CFrame"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/Region3:Properties/CFrame/Summary",
						Description: "libraries/roblox/types/Region3:Properties/CFrame/Description",
					}
				},
			},
			"Size": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.Region3).Size())
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("Vector3"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/Region3:Properties/Size/Summary",
						Description: "libraries/roblox/types/Region3:Properties/Size/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"ExpandToGrid": {
				Func: func(s rbxmk.State, v types.Value) int {
					region := float64(s.Pull(2, "float").(types.Float))
					return s.Push(v.(types.Region3).ExpandToGrid(region))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "region", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Region3")},
						},
						Summary:     "libraries/roblox/types/Region3:Methods/ExpandToGrid/Summary",
						Description: "libraries/roblox/types/Region3:Methods/ExpandToGrid/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.Region3{
						Min: s.Pull(1, "Vector3").(types.Vector3),
						Max: s.Pull(2, "Vector3").(types.Vector3),
					})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "min", Type: dt.Prim("Vector3")},
								{Name: "max", Type: dt.Prim("Vector3")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("Region3")},
							},
							Summary:     "libraries/roblox/types/Region3:Constructors/new/Summary",
							Description: "libraries/roblox/types/Region3:Constructors/new/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators:   &dump.Operators{Eq: true},
				Summary:     "libraries/roblox/types/Region3:Summary",
				Description: "libraries/roblox/types/Region3:Description",
			}
		},
	}
}
