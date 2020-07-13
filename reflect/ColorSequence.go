package reflect

import (
	"strconv"
	"strings"

	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func ColorSequence() Type {
	return Type{
		Name:        "ColorSequence",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				u := v.(types.ColorSequence)
				var b strings.Builder
				for i, v := range u {
					if i > 0 {
						b.WriteString("; ")
					}
					b.WriteString(strconv.FormatFloat(float64(v.Time), 'g', -1, 32))
					b.WriteString(", (")
					b.WriteString(strconv.FormatFloat(float64(v.Value.R), 'g', -1, 32))
					b.WriteString(", ")
					b.WriteString(strconv.FormatFloat(float64(v.Value.G), 'g', -1, 32))
					b.WriteString(", ")
					b.WriteString(strconv.FormatFloat(float64(v.Value.B), 'g', -1, 32))
					b.WriteString("), ")
					b.WriteString(strconv.FormatFloat(float64(v.Envelope), 'g', -1, 32))
				}
				s.L.Push(lua.LString(b.String()))
				return 1
			},
			"__eq": func(s State, v Value) int {
				u := v.(types.ColorSequence)
				op := s.Pull(2, "ColorSequence").(types.ColorSequence)
				if len(op) != len(u) {
					return s.Push("bool", false)
				}
				for i, v := range u {
					if v != op[i] {
						return s.Push("bool", false)
					}
				}
				return s.Push("bool", true)
			},
		},
		Members: map[string]Member{
			"Keypoints": {Get: func(s State, v Value) int {
				u := v.(types.ColorSequence)
				keypointType := s.Type("ColorSequenceKeypoint")
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
				var v types.ColorSequence
				switch s.Count() {
				case 1:
					switch c := s.PullAnyOf(1, "Color3", "table").(type) {
					case types.Color3:
						v = types.ColorSequence{
							types.ColorSequenceKeypoint{Time: 0, Value: c},
							types.ColorSequenceKeypoint{Time: 1, Value: c},
						}
					case *lua.LTable:
						n := c.Len()
						if n < 2 {
							s.L.RaiseError("ColorSequence requires at least 2 keypoints")
							return 0
						}
						v = make(types.ColorSequence, n)
						keypointType := s.Type("ColorSequenceKeypoint")
						for i := 1; i <= n; i++ {
							k, err := keypointType.ReflectFrom(s, keypointType, c.RawGetInt(i))
							if err != nil {
								s.L.RaiseError(err.Error())
								return 0
							}
							v[i] = k.(types.ColorSequenceKeypoint)
						}
						const epsilon = 1e-4
						if t := v[len(v)-1].Time; t < 1-epsilon || t > 1+epsilon {
							s.L.RaiseError("ColorSequence time must end at 1.0")
							return 0
						}
						if t := v[0].Time; t < -epsilon || t > epsilon {
							s.L.RaiseError("ColorSequence time must start at 0.0")
							return 0
						}
					}
				case 2:
					v = types.ColorSequence{
						types.ColorSequenceKeypoint{Time: 0, Value: s.Pull(1, "Color3").(types.Color3)},
						types.ColorSequenceKeypoint{Time: 1, Value: s.Pull(2, "Color3").(types.Color3)},
					}
				default:
					s.L.RaiseError("expected 1 or 2 arguments")
					return 0
				}
				return s.Push("ColorSequence", v)
			},
		},
	}
}
