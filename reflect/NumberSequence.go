package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func NumberSequence() Type {
	return Type{
		Name:        "NumberSequence",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString(s.Pull(1, "NumberSequence").(types.NumberSequence).String()))
				return 1
			},
			"__eq": func(s State) int {
				u := s.Pull(1, "NumberSequence").(types.NumberSequence)
				op := s.Pull(2, "NumberSequence").(types.NumberSequence)
				if len(op) != len(u) {
					return s.Push("bool", types.False)
				}
				for i, v := range u {
					if v != op[i] {
						return s.Push("bool", types.False)
					}
				}
				return s.Push("bool", types.True)
			},
		},
		Members: map[string]Member{
			"Keypoints": {Get: func(s State, v types.Value) int {
				u := v.(types.NumberSequence)
				keypointType := s.Type("NumberSequenceKeypoint")
				table := s.L.CreateTable(len(u), 0)
				for i, v := range u {
					lv, err := keypointType.ReflectTo(s, keypointType, v)
					if err != nil {
						s.L.RaiseError(err.Error())
						return 0
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
							s.L.RaiseError("NumberSequence requires at least 2 keypoints")
							return 0
						}
						v = make(types.NumberSequence, n)
						keypointType := s.Type("NumberSequenceKeypoint")
						for i := 1; i <= n; i++ {
							k, err := keypointType.ReflectFrom(s, keypointType, c.RawGetInt(i))
							if err != nil {
								s.L.RaiseError(err.Error())
								return 0
							}
							v[i] = k.(types.NumberSequenceKeypoint)
						}
						const epsilon = 1e-4
						if t := v[len(v)-1].Time; t < 1-epsilon || t > 1+epsilon {
							s.L.RaiseError("NumberSequence time must end at 1.0")
							return 0
						}
						if t := v[0].Time; t < -epsilon || t > epsilon {
							s.L.RaiseError("NumberSequence time must start at 0.0")
							return 0
						}
					}
				case 2:
					v = types.NumberSequence{
						types.NumberSequenceKeypoint{Time: 0, Value: float32(s.Pull(1, "float").(types.Float))},
						types.NumberSequenceKeypoint{Time: 1, Value: float32(s.Pull(2, "float").(types.Float))},
					}
				default:
					s.L.RaiseError("expected 1 or 2 arguments")
					return 0
				}
				return s.Push("NumberSequence", v)
			},
		},
	}
}
