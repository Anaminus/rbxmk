package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(ColorSequence) }
func ColorSequence() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "ColorSequence",
		PushTo:   rbxmk.PushTypeTo("ColorSequence"),
		PullFrom: rbxmk.PullTypeFrom("ColorSequence"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "ColorSequence").(types.ColorSequence)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "ColorSequence").(types.ColorSequence)
				op := s.Pull(2, "ColorSequence").(types.ColorSequence)
				if len(op) != len(v) {
					s.L.Push(lua.LFalse)
					return 1
				}
				for i, k := range v {
					if k != op[i] {
						s.L.Push(lua.LFalse)
						return 1
					}
				}
				s.L.Push(lua.LTrue)
				return 1
			},
		},
		Members: map[string]rbxmk.Member{
			"Keypoints": {
				Get: func(s rbxmk.State, v types.Value) int {
					u := v.(types.ColorSequence)
					keypointRfl := s.MustReflector("ColorSequenceKeypoint")
					table := s.L.CreateTable(len(u), 0)
					for i, v := range u {
						lv, err := keypointRfl.PushTo(s, v)
						if err != nil {
							return s.RaiseError("%s", err)
						}
						table.RawSetInt(i, lv[0])
					}
					s.L.Push(table)
					return 1
				},
				Dump: func() dump.Value {
					return dump.Property{ValueType: dt.Array{T: dt.Prim("ColorSequenceKeypoint")}, ReadOnly: true}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.ColorSequence
					switch s.Count() {
					case 1:
						switch c := s.PullAnyOf(1, "Color3", "table").(type) {
						case types.Color3:
							v = types.ColorSequence{
								types.ColorSequenceKeypoint{Time: 0, Value: c},
								types.ColorSequenceKeypoint{Time: 1, Value: c},
							}
						case rtypes.Table:
							n := c.Len()
							if n < 2 {
								return s.RaiseError("ColorSequence requires at least 2 keypoints")
							}
							v = make(types.ColorSequence, n)
							keypointRfl := s.MustReflector("ColorSequenceKeypoint")
							for i := 1; i <= n; i++ {
								k, err := keypointRfl.PullFrom(s, c.RawGetInt(i))
								if err != nil {
									return s.RaiseError("%s", err)
								}
								v[i] = k.(types.ColorSequenceKeypoint)
							}
							const epsilon = 1e-4
							if t := v[len(v)-1].Time; t < 1-epsilon || t > 1+epsilon {
								return s.RaiseError("ColorSequence time must end at 1.0")
							}
							if t := v[0].Time; t < -epsilon || t > epsilon {
								return s.RaiseError("ColorSequence time must start at 0.0")
							}
						default:
							return s.ReflectorError(1)
						}
					case 2:
						v = types.ColorSequence{
							types.ColorSequenceKeypoint{Time: 0, Value: s.Pull(1, "Color3").(types.Color3)},
							types.ColorSequenceKeypoint{Time: 1, Value: s.Pull(2, "Color3").(types.Color3)},
						}
					default:
						return s.RaiseError("expected 1 or 2 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{
						{
							Parameters: dump.Parameters{
								{Name: "color", Type: dt.Prim("Color3")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("ColorSequence")},
							},
						},
						{
							Parameters: dump.Parameters{
								{Name: "color0", Type: dt.Prim("Color3")},
								{Name: "color1", Type: dt.Prim("Color3")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("ColorSequence")},
							},
						},
						{
							Parameters: dump.Parameters{
								{Name: "keypoints", Type: dt.Array{T: dt.Prim("ColorSequenceKeypoint")}},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("ColorSequence")},
							},
							CanError: true,
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
