package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(ColorSequenceKeypoint) }
func ColorSequenceKeypoint() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "ColorSequenceKeypoint",
		PushTo:   rbxmk.PushTypeTo("ColorSequenceKeypoint"),
		PullFrom: rbxmk.PullTypeFrom("ColorSequenceKeypoint"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "ColorSequenceKeypoint").(types.ColorSequenceKeypoint)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "ColorSequenceKeypoint").(types.ColorSequenceKeypoint)
				op := s.Pull(2, "ColorSequenceKeypoint").(types.ColorSequenceKeypoint)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Time": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.ColorSequenceKeypoint).Time))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"Value": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.ColorSequenceKeypoint).Value)
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("Color3"), ReadOnly: true} },
			},
			"Envelope": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.ColorSequenceKeypoint).Envelope))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.ColorSequenceKeypoint
					switch s.Count() {
					case 2:
						v.Time = float32(s.Pull(1, "float").(types.Float))
						v.Value = s.Pull(2, "Color3").(types.Color3)
					case 3:
						v.Time = float32(s.Pull(1, "float").(types.Float))
						v.Value = s.Pull(2, "Color3").(types.Color3)
						v.Envelope = float32(s.Pull(3, "float").(types.Float))
					default:
						return s.RaiseError("expected 2 or 3 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{
						{
							Parameters: dump.Parameters{
								{Name: "time", Type: dt.Prim("float")},
								{Name: "color", Type: dt.Prim("Color3")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("ColorSequenceKeypoint")},
							},
						},
						{
							Parameters: dump.Parameters{
								{Name: "time", Type: dt.Prim("float")},
								{Name: "color", Type: dt.Prim("Color3")},
								{Name: "envelope", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("ColorSequenceKeypoint")},
							},
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
