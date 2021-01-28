package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(NumberSequence) }
func NumberSequence() Reflector {
	return Reflector{
		Name:     "NumberSequence",
		PushTo:   rbxmk.PushTypeTo,
		PullFrom: rbxmk.PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "NumberSequence").(types.NumberSequence)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
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
		Members: map[string]Member{
			"Keypoints": {Get: func(s State, v types.Value) int {
				u := v.(types.NumberSequence)
				keypointRfl := s.Reflector("NumberSequenceKeypoint")
				table := s.L.CreateTable(len(u), 0)
				for i, v := range u {
					lv, err := keypointRfl.PushTo(s, keypointRfl, v)
					if err != nil {
						return s.RaiseError(err.Error())
					}
					table.RawSetInt(i, lv[0])
				}
				s.L.Push(table)
				return 1
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
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
						keypointRfl := s.Reflector("NumberSequenceKeypoint")
						for i := 1; i <= n; i++ {
							k, err := keypointRfl.PullFrom(s, keypointRfl, c.RawGetInt(i))
							if err != nil {
								return s.RaiseError(err.Error())
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
		},
	}
}
