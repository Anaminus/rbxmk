package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(NumberRange) }
func NumberRange() Reflector {
	return Reflector{
		Name:     "NumberRange",
		PushTo:   rbxmk.PushTypeTo("NumberRange"),
		PullFrom: rbxmk.PullTypeFrom("NumberRange"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "NumberRange").(types.NumberRange)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "NumberRange").(types.NumberRange)
				op := s.Pull(2, "NumberRange").(types.NumberRange)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: map[string]Member{
			"Min": {
				Get: func(s State, v types.Value) int {
					return s.Push(types.Float(v.(types.NumberRange).Min))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"Max": {
				Get: func(s State, v types.Value) int {
					return s.Push(types.Float(v.(types.NumberRange).Max))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
		},
		Constructors: Constructors{
			"new": {
				Func: func(s State) int {
					var v types.NumberRange
					switch s.Count() {
					case 1:
						v.Min = float32(s.Pull(1, "float").(types.Float))
						v.Max = v.Min
					case 2:
						v.Min = float32(s.Pull(1, "float").(types.Float))
						v.Max = float32(s.Pull(2, "float").(types.Float))
					default:
						return s.RaiseError("expected 1 or 2 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{
						{
							Parameters: dump.Parameters{
								{Name: "value", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("NumberRange")},
							},
						},
						{
							Parameters: dump.Parameters{
								{Name: "minimum", Type: dt.Prim("float")},
								{Name: "maxmimum", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("NumberRange")},
							},
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
