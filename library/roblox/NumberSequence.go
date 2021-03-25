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
		Name:     "NumberSequence",
		PushTo:   rbxmk.PushTypeTo("NumberSequence"),
		PullFrom: rbxmk.PullTypeFrom("NumberSequence"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "NumberSequence").(types.NumberSequence)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "NumberSequence").(types.NumberSequence)
				op := s.Pull(2, "NumberSequence").(types.NumberSequence)
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
					keypointRfl := s.MustReflector("NumberSequenceKeypoint")
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
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Array{T: dt.Prim("NumberSequenceKeypoint")},
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/NumberSequence:Properties/Keypoints/Summary",
						Description: "libraries/roblox/types/NumberSequence:Properties/Keypoints/Description",
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
						switch c := s.PullAnyOf(1, "float", "table").(type) {
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
							keypointRfl := s.MustReflector("NumberSequenceKeypoint")
							for i := 1; i <= n; i++ {
								k, err := keypointRfl.PullFrom(s, c.RawGetInt(i))
								if err != nil {
									return s.RaiseError("%s", err)
								}
								v[i] = k.(types.NumberSequenceKeypoint)
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
							types.NumberSequenceKeypoint{Time: 0, Value: float32(s.Pull(1, "float").(types.Float))},
							types.NumberSequenceKeypoint{Time: 1, Value: float32(s.Pull(2, "float").(types.Float))},
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
								{Name: "value", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("NumberSequence")},
							},
							Summary:     "libraries/roblox/types/NumberSequence:Constructors/new/1/Summary",
							Description: "libraries/roblox/types/NumberSequence:Constructors/new/1/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "value0", Type: dt.Prim("float")},
								{Name: "value1", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("NumberSequence")},
							},
							Summary:     "libraries/roblox/types/NumberSequence:Constructors/new/2/Summary",
							Description: "libraries/roblox/types/NumberSequence:Constructors/new/2/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "keypoints", Type: dt.Array{T: dt.Prim("NumberSequenceKeypoint")}},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("NumberSequence")},
							},
							CanError:    true,
							Summary:     "libraries/roblox/types/NumberSequence:Constructors/new/3/Summary",
							Description: "libraries/roblox/types/NumberSequence:Constructors/new/3/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators:   &dump.Operators{Eq: true},
				Summary:     "libraries/roblox/types/NumberSequence:Summary",
				Description: "libraries/roblox/types/NumberSequence:Description",
			}
		},
	}
}
