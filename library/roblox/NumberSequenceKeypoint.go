package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(NumberSequenceKeypoint) }
func NumberSequenceKeypoint() Reflector {
	return Reflector{
		Name:     "NumberSequenceKeypoint",
		PushTo:   rbxmk.PushTypeTo("NumberSequenceKeypoint"),
		PullFrom: rbxmk.PullTypeFrom("NumberSequenceKeypoint"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "NumberSequenceKeypoint").(types.NumberSequenceKeypoint)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "NumberSequenceKeypoint").(types.NumberSequenceKeypoint)
				op := s.Pull(2, "NumberSequenceKeypoint").(types.NumberSequenceKeypoint)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: map[string]Member{
			"Time": {
				Get: func(s State, v types.Value) int {
					return s.Push(types.Float(v.(types.NumberSequenceKeypoint).Time))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"Value": {
				Get: func(s State, v types.Value) int {
					return s.Push(types.Float(v.(types.NumberSequenceKeypoint).Value))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"Envelope": {
				Get: func(s State, v types.Value) int {
					return s.Push(types.Float(v.(types.NumberSequenceKeypoint).Envelope))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
		},
		Constructors: Constructors{
			"new": {
				Func: func(s State) int {
					var v types.NumberSequenceKeypoint
					switch s.Count() {
					case 2:
						v.Time = float32(s.Pull(1, "float").(types.Float))
						v.Value = float32(s.Pull(2, "float").(types.Float))
					case 3:
						v.Time = float32(s.Pull(1, "float").(types.Float))
						v.Value = float32(s.Pull(2, "float").(types.Float))
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
								{Name: "value", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("ColorSequenceKeypoint")},
							},
						},
						{
							Parameters: dump.Parameters{
								{Name: "time", Type: dt.Prim("float")},
								{Name: "value", Type: dt.Prim("float")},
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
