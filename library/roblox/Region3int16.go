package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(Region3int16) }
func Region3int16() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "Region3int16",
		PushTo:   rbxmk.PushTypeTo("Region3int16"),
		PullFrom: rbxmk.PullTypeFrom("Region3int16"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "Region3int16").(types.Region3int16)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "Region3int16").(types.Region3int16)
				op := s.Pull(2, "Region3int16").(types.Region3int16)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Min": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.Region3int16).Min)
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("Vector3int16"), ReadOnly: true} },
			},
			"Max": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.Region3int16).Max)
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("Vector3int16"), ReadOnly: true} },
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.Region3int16{
						Min: s.Pull(1, "Vector3int16").(types.Vector3int16),
						Max: s.Pull(2, "Vector3int16").(types.Vector3int16),
					})
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Parameters: dump.Parameters{
							{Name: "min", Type: dt.Prim("Vector3int16")},
							{Name: "max", Type: dt.Prim("Vector3int16")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Region3int16")},
						},
					}}
				},
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
