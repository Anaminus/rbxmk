package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(Region3) }
func Region3() Reflector {
	return Reflector{
		Name:     "Region3",
		PushTo:   rbxmk.PushTypeTo("Region3"),
		PullFrom: rbxmk.PullTypeFrom("Region3"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Region3").(types.Region3)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Region3").(types.Region3)
				op := s.Pull(2, "Region3").(types.Region3)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: map[string]Member{
			"CFrame": {
				Get: func(s State, v types.Value) int {
					return s.Push(v.(types.Region3).CFrame())
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("CFrame")} },
			},
			"Size": {
				Get: func(s State, v types.Value) int {
					return s.Push(v.(types.Region3).Size())
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("Vector3")} },
			},
			"ExpandToGrid": {Method: true,
				Get: func(s State, v types.Value) int {
					region := float64(s.Pull(2, "float").(types.Float))
					return s.Push(v.(types.Region3).ExpandToGrid(region))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "region", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Region3")},
						},
					}
				},
			},
		},
		Constructors: Constructors{
			"new": {
				Func: func(s State) int {
					return s.Push(types.Region3{
						Min: s.Pull(1, "Vector3").(types.Vector3),
						Max: s.Pull(2, "Vector3").(types.Vector3),
					})
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Parameters: dump.Parameters{
							{Name: "min", Type: dt.Prim("Vector3")},
							{Name: "max", Type: dt.Prim("Vector3")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Region3")},
						},
					}}
				},
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
