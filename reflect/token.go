package reflect

import (
	"math"
	"strconv"

	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Token() Type {
	return Type{
		Name: "token",
		ReflectTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Token))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LNumber:
				return types.Token(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("token") {
					if v, ok := v.Value.(types.Token); ok {
						return v, nil
					}
				}
			}
			return nil, TypeError(nil, 0, "token")
		},
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString("token: " + strconv.FormatInt(int64(types.Token(pullInt(s, 1))), 10)))
				return 1
			},
			"__call": func(s State) int {
				return s.Push("number", types.Double(pullInt(s, 1)))
			},
			"__eq": func(s State) int {
				return s.Push("bool", types.Bool(types.Token(pullInt(s, 1)) == types.Token(pullInt(s, 2))))
			},
			"__lt": func(s State) int {
				return s.Push("bool", types.Bool(types.Token(pullInt(s, 1)) < types.Token(pullInt(s, 2))))
			},
			"__le": func(s State) int {
				return s.Push("bool", types.Bool(types.Token(pullInt(s, 1)) <= types.Token(pullInt(s, 2))))
			},
			"__add": func(s State) int {
				return s.Push("token", types.Token(pullInt(s, 1))+types.Token(pullInt(s, 2)))
			},
			"__sub": func(s State) int {
				return s.Push("token", types.Token(pullInt(s, 1))-types.Token(pullInt(s, 2)))
			},
			"__mul": func(s State) int {
				return s.Push("token", types.Token(pullInt(s, 1))*types.Token(pullInt(s, 2)))
			},
			"__div": func(s State) int {
				a := types.Token(pullInt(s, 1))
				b := types.Token(pullInt(s, 1))
				if b == 0 {
					if a == 0 {
						return s.Push("double", types.Double(math.NaN()))
					}
					return s.Push("double", types.Double(math.Inf(int(a))))
				}
				return s.Push("token", a/b)
			},
			"__mod": func(s State) int {
				return s.Push("token", types.Token(pullInt(s, 1))%types.Token(pullInt(s, 2)))
			},
			"__pow": func(s State) int {
				return s.Push("token", types.Token(math.Pow(float64(types.Token(pullInt(s, 1))), float64(types.Token(pullInt(s, 2))))))
			},
			"__unm": func(s State) int {
				return s.Push("token", -types.Token(pullInt(s, 1)))
			},
		},
	}
}
