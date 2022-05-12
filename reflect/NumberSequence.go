package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(NumberSequence) }
func NumberSequence() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     rtypes.T_NumberSequence,
		PushTo:   rbxmk.PushTypeTo(rtypes.T_NumberSequence),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_NumberSequence),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.NumberSequence:
				*p = v.(types.NumberSequence)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_NumberSequence).(types.NumberSequence)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_NumberSequence).(types.NumberSequence)
				op := s.Pull(2, rtypes.T_NumberSequence).(types.NumberSequence)
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
		Properties: rbxmk.Properties{
			"Keypoints": {
				Get: func(s rbxmk.State, v types.Value) int {
					u := v.(types.NumberSequence)
					keypointRfl := s.MustReflector(rtypes.T_NumberSequenceKeypoint)
					table := s.L.CreateTable(len(u), 0)
					for i, v := range u {
						lv, err := keypointRfl.PushTo(s.Context(), v)
						if err != nil {
							return s.RaiseError("%s", err)
						}
						table.RawSetInt(i, lv)
					}
					s.L.Push(table)
					return 1
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Array{T: dt.Prim(rtypes.T_NumberSequenceKeypoint)},
						ReadOnly:    true,
						Summary:     "Types/NumberSequence:Properties/Keypoints/Summary",
						Description: "Types/NumberSequence:Properties/Keypoints/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.NumberSequence
					switch s.Count() {
					case 1:
						switch c := s.PullAnyOf(1, rtypes.T_Float, rtypes.T_Table).(type) {
						case types.Float:
							v = types.NumberSequence{
								types.NumberSequenceKeypoint{Time: 0, Value: float32(c)},
								types.NumberSequenceKeypoint{Time: 1, Value: float32(c)},
							}
						case rtypes.Table:
							n := c.Len()
							if n < 2 {
								return s.RaiseError("NumberSequence requires at least 2 keypoints")
							}
							v = make(types.NumberSequence, n)
							keypointRfl := s.MustReflector(rtypes.T_NumberSequenceKeypoint)
							for i := 1; i <= n; i++ {
								k, err := keypointRfl.PullFrom(s.Context(), c.RawGetInt(i))
								if err != nil {
									return s.RaiseError("%s", err)
								}
								v[i-1] = k.(types.NumberSequenceKeypoint)
							}
							for i := 1; i < len(v); i++ {
								if v[i].Time < v[i-1].Time {
									return s.RaiseError("all NumberSequenceKeypoints must be ordered by time")
								}
							}
							const epsilon = 1e-4
							if t := v[len(v)-1].Time; t < 1-epsilon || t > 1+epsilon {
								return s.RaiseError("NumberSequence time must end at 1.0")
							}
							if t := v[0].Time; t < -epsilon || t > epsilon {
								return s.RaiseError("NumberSequence time must start at 0.0")
							}
						default:
							return s.ReflectorError(1)
						}
					case 2:
						v = types.NumberSequence{
							types.NumberSequenceKeypoint{Time: 0, Value: float32(s.Pull(1, rtypes.T_Float).(types.Float))},
							types.NumberSequenceKeypoint{Time: 1, Value: float32(s.Pull(2, rtypes.T_Float).(types.Float))},
						}
					default:
						return s.RaiseError("expected 1 or 2 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "value", Type: dt.Prim(rtypes.T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_NumberSequence)},
							},
							Summary:     "Types/NumberSequence:Constructors/new/Single/Summary",
							Description: "Types/NumberSequence:Constructors/new/Single/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "value0", Type: dt.Prim(rtypes.T_Float)},
								{Name: "value1", Type: dt.Prim(rtypes.T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_NumberSequence)},
							},
							Summary:     "Types/NumberSequence:Constructors/new/Range/Summary",
							Description: "Types/NumberSequence:Constructors/new/Range/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "keypoints", Type: dt.Array{T: dt.Prim(rtypes.T_NumberSequenceKeypoint)}},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_NumberSequence)},
							},
							CanError:    true,
							Summary:     "Types/NumberSequence:Constructors/new/Keypoints/Summary",
							Description: "Types/NumberSequence:Constructors/new/Keypoints/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/NumberSequence:Operators/Eq/Summary",
						Description: "Types/NumberSequence:Operators/Eq/Description",
					},
				},
				Summary:     "Types/NumberSequence:Summary",
				Description: "Types/NumberSequence:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Float,
			NumberSequenceKeypoint,
			Table,
		},
	}
}
