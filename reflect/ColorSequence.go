package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func ColorSequence() Type {
	return Type{
		Name:     "ColorSequence",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString(s.Pull(1, "ColorSequence").(types.ColorSequence).String()))
				return 1
			},
			"__eq": func(s State) int {
				u := s.Pull(1, "ColorSequence").(types.ColorSequence)
				op := s.Pull(2, "ColorSequence").(types.ColorSequence)
				if len(op) != len(u) {
					return s.Push(types.False)
				}
				for i, v := range u {
					if v != op[i] {
						return s.Push(types.False)
					}
				}
				return s.Push(types.True)
			},
		},
		Members: map[string]Member{
			"Keypoints": {Get: func(s State, v types.Value) int {
				u := v.(types.ColorSequence)
				keypointType := s.Type("ColorSequenceKeypoint")
				table := s.L.CreateTable(len(u), 0)
				for i, v := range u {
					lv, err := keypointType.PushTo(s, keypointType, v)
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
						keypointType := s.Type("ColorSequenceKeypoint")
						for i := 1; i <= n; i++ {
							k, err := keypointType.PullFrom(s, keypointType, c.RawGetInt(i))
							if err != nil {
								return s.RaiseError(err.Error())
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
		},
	}
}
