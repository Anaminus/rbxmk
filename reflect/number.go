package reflect

import (
	"math"
	"strconv"

	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func pullNumber(s State, n int) float64 {
	switch v := s.L.CheckAny(n).(type) {
	case lua.LNumber:
		return float64(v)
	case *lua.LUserData:
		switch v := v.Value.(type) {
		case types.Numberlike:
			return v.Numberlike()
		case types.Intlike:
			return float64(v.Intlike())
		}
	}
	panic("expected number-like type")
}

func pullInt(s State, n int) int64 {
	switch v := s.L.CheckAny(n).(type) {
	case lua.LNumber:
		return int64(v)
	case *lua.LUserData:
		switch v := v.Value.(type) {
		case types.Intlike:
			return v.Intlike()
		case types.Numberlike:
			return int64(v.Numberlike())
		}
	}
	panic("expected int-like type")
}

func Float() Type {
	return Type{
		Name:  "float",
		Flags: Exprim,
		ReflectTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Float))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LNumber:
				return types.Float(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("float") {
					if v, ok := v.Value.(types.Float); ok {
						return v, nil
					}
				}
			}
			return nil, TypeError(nil, 0, "float")
		},
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString("float: " + strconv.FormatFloat(pullNumber(s, 1), 'g', -1, 32)))
				return 1
			},
			"__call": func(s State) int {
				return s.Push(types.Double(pullNumber(s, 1)))
			},
			"__eq": func(s State) int {
				return s.Push(types.Bool(types.Float(pullNumber(s, 1)) == types.Float(pullNumber(s, 2))))
			},
			"__lt": func(s State) int {
				return s.Push(types.Bool(types.Float(pullNumber(s, 1)) < types.Float(pullNumber(s, 2))))
			},
			"__le": func(s State) int {
				return s.Push(types.Bool(types.Float(pullNumber(s, 1)) <= types.Float(pullNumber(s, 2))))
			},
			"__add": func(s State) int {
				return s.Push(types.Float(pullNumber(s, 1)) + types.Float(pullNumber(s, 2)))
			},
			"__sub": func(s State) int {
				return s.Push(types.Float(pullNumber(s, 1)) - types.Float(pullNumber(s, 2)))
			},
			"__mul": func(s State) int {
				return s.Push(types.Float(pullNumber(s, 1)) * types.Float(pullNumber(s, 2)))
			},
			"__div": func(s State) int {
				a := types.Float(pullNumber(s, 1))
				b := types.Float(pullNumber(s, 1))
				if b == 0 {
					if a == 0 {
						return s.Push(types.Float(math.NaN()))
					}
					return s.Push(types.Float(math.Inf(int(a))))
				}
				return s.Push(a / b)
			},
			"__mod": func(s State) int {
				return s.Push(types.Float(math.Mod(pullNumber(s, 1), pullNumber(s, 2))))
			},
			"__pow": func(s State) int {
				return s.Push(types.Float(math.Pow(pullNumber(s, 1), pullNumber(s, 2))))
			},
			"__unm": func(s State) int {
				return s.Push(-types.Float(pullNumber(s, 1)))
			},
		},
	}
}

func Double() Type {
	return Type{
		Name: "double",
		ReflectTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Double))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LNumber:
				return types.Double(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("double") {
					if v, ok := v.Value.(types.Double); ok {
						return v, nil
					}
				}
			}
			return nil, TypeError(nil, 0, "double")
		},
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString("double: " + strconv.FormatFloat(pullNumber(s, 1), 'g', -1, 32)))
				return 1
			},
			"__call": func(s State) int {
				return s.Push(types.Double(pullNumber(s, 1)))
			},
			"__eq": func(s State) int {
				return s.Push(types.Bool(types.Double(pullNumber(s, 1)) == types.Double(pullNumber(s, 2))))
			},
			"__lt": func(s State) int {
				return s.Push(types.Bool(types.Double(pullNumber(s, 1)) < types.Double(pullNumber(s, 2))))
			},
			"__le": func(s State) int {
				return s.Push(types.Bool(types.Double(pullNumber(s, 1)) <= types.Double(pullNumber(s, 2))))
			},
			"__add": func(s State) int {
				return s.Push(types.Double(pullNumber(s, 1)) + types.Double(pullNumber(s, 2)))
			},
			"__sub": func(s State) int {
				return s.Push(types.Double(pullNumber(s, 1)) - types.Double(pullNumber(s, 2)))
			},
			"__mul": func(s State) int {
				return s.Push(types.Double(pullNumber(s, 1)) * types.Double(pullNumber(s, 2)))
			},
			"__div": func(s State) int {
				a := types.Double(pullNumber(s, 1))
				b := types.Double(pullNumber(s, 1))
				if b == 0 {
					if a == 0 {
						return s.Push(types.Double(math.NaN()))
					}
					return s.Push(types.Double(math.Inf(int(a))))
				}
				return s.Push(a / b)
			},
			"__mod": func(s State) int {
				return s.Push(types.Double(math.Mod(pullNumber(s, 1), pullNumber(s, 2))))
			},
			"__pow": func(s State) int {
				return s.Push(types.Double(math.Pow(pullNumber(s, 1), pullNumber(s, 2))))
			},
			"__unm": func(s State) int {
				return s.Push(-types.Double(pullNumber(s, 1)))
			},
		},
	}
}

func Number() Type {
	return Type{
		Name: "number",
		ReflectTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Double))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			if n, ok := lvs[0].(lua.LNumber); ok {
				return types.Double(n), nil
			}
			return nil, TypeError(nil, 0, "number")
		},
	}
}

func Int() Type {
	return Type{
		Name:  "int",
		Flags: Exprim,
		ReflectTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Int))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LNumber:
				return types.Int(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("int") {
					if v, ok := v.Value.(types.Int); ok {
						return v, nil
					}
				}
			}
			return nil, TypeError(nil, 0, "int")
		},
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString("int: " + strconv.FormatInt(int64(types.Int(pullInt(s, 1))), 10)))
				return 1
			},
			"__call": func(s State) int {
				return s.Push(types.Double(pullInt(s, 1)))
			},
			"__eq": func(s State) int {
				return s.Push(types.Bool(types.Int(pullInt(s, 1)) == types.Int(pullInt(s, 2))))
			},
			"__lt": func(s State) int {
				return s.Push(types.Bool(types.Int(pullInt(s, 1)) < types.Int(pullInt(s, 2))))
			},
			"__le": func(s State) int {
				return s.Push(types.Bool(types.Int(pullInt(s, 1)) <= types.Int(pullInt(s, 2))))
			},
			"__add": func(s State) int {
				return s.Push(types.Int(pullInt(s, 1)) + types.Int(pullInt(s, 2)))
			},
			"__sub": func(s State) int {
				return s.Push(types.Int(pullInt(s, 1)) - types.Int(pullInt(s, 2)))
			},
			"__mul": func(s State) int {
				return s.Push(types.Int(pullInt(s, 1)) * types.Int(pullInt(s, 2)))
			},
			"__div": func(s State) int {
				a := types.Int(pullInt(s, 1))
				b := types.Int(pullInt(s, 1))
				if b == 0 {
					if a == 0 {
						return s.Push(types.Double(math.NaN()))
					}
					return s.Push(types.Double(math.Inf(int(a))))
				}
				return s.Push(a / b)
			},
			"__mod": func(s State) int {
				return s.Push(types.Int(pullInt(s, 1)) % types.Int(pullInt(s, 2)))
			},
			"__pow": func(s State) int {
				return s.Push(types.Int(math.Pow(float64(types.Int(pullInt(s, 1))), float64(types.Int(pullInt(s, 2))))))
			},
			"__unm": func(s State) int {
				return s.Push(-types.Int(pullInt(s, 1)))
			},
		},
	}
}

func Int64() Type {
	return Type{
		Name:  "int64",
		Flags: Exprim,
		ReflectTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Int64))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LNumber:
				return types.Int64(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("int64") {
					if v, ok := v.Value.(types.Int64); ok {
						return v, nil
					}
				}
			}
			return nil, TypeError(nil, 0, "int64")
		},
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString("int64: " + strconv.FormatInt(pullInt(s, 1), 10)))
				return 1
			},
			"__call": func(s State) int {
				return s.Push(types.Double(pullInt(s, 1)))
			},
			"__eq": func(s State) int {
				return s.Push(types.Bool(types.Int64(pullInt(s, 1)) == types.Int64(pullInt(s, 2))))
			},
			"__lt": func(s State) int {
				return s.Push(types.Bool(types.Int64(pullInt(s, 1)) < types.Int64(pullInt(s, 2))))
			},
			"__le": func(s State) int {
				return s.Push(types.Bool(types.Int64(pullInt(s, 1)) <= types.Int64(pullInt(s, 2))))
			},
			"__add": func(s State) int {
				return s.Push(types.Int64(pullInt(s, 1)) + types.Int64(pullInt(s, 2)))
			},
			"__sub": func(s State) int {
				return s.Push(types.Int64(pullInt(s, 1)) - types.Int64(pullInt(s, 2)))
			},
			"__mul": func(s State) int {
				return s.Push(types.Int64(pullInt(s, 1)) * types.Int64(pullInt(s, 2)))
			},
			"__div": func(s State) int {
				a := types.Int64(pullInt(s, 1))
				b := types.Int64(pullInt(s, 1))
				if b == 0 {
					if a == 0 {
						return s.Push(types.Double(math.NaN()))
					}
					return s.Push(types.Double(math.Inf(int(a))))
				}
				return s.Push(a / b)
			},
			"__mod": func(s State) int {
				return s.Push(types.Int64(pullInt(s, 1)) % types.Int64(pullInt(s, 2)))
			},
			"__pow": func(s State) int {
				return s.Push(types.Int64(math.Pow(float64(types.Int64(pullInt(s, 1))), float64(types.Int64(pullInt(s, 2))))))
			},
			"__unm": func(s State) int {
				return s.Push(-types.Int64(pullInt(s, 1)))
			},
		},
	}
}

func Token() Type {
	return Type{
		Name:  "token",
		Flags: Exprim,
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
				return s.Push(types.Double(pullInt(s, 1)))
			},
			"__eq": func(s State) int {
				return s.Push(types.Bool(types.Token(pullInt(s, 1)) == types.Token(pullInt(s, 2))))
			},
			"__lt": func(s State) int {
				return s.Push(types.Bool(types.Token(pullInt(s, 1)) < types.Token(pullInt(s, 2))))
			},
			"__le": func(s State) int {
				return s.Push(types.Bool(types.Token(pullInt(s, 1)) <= types.Token(pullInt(s, 2))))
			},
			"__add": func(s State) int {
				return s.Push(types.Token(pullInt(s, 1)) + types.Token(pullInt(s, 2)))
			},
			"__sub": func(s State) int {
				return s.Push(types.Token(pullInt(s, 1)) - types.Token(pullInt(s, 2)))
			},
			"__mul": func(s State) int {
				return s.Push(types.Token(pullInt(s, 1)) * types.Token(pullInt(s, 2)))
			},
			"__div": func(s State) int {
				a := types.Token(pullInt(s, 1))
				b := types.Token(pullInt(s, 1))
				if b == 0 {
					if a == 0 {
						return s.Push(types.Double(math.NaN()))
					}
					return s.Push(types.Double(math.Inf(int(a))))
				}
				return s.Push(a / b)
			},
			"__mod": func(s State) int {
				return s.Push(types.Token(pullInt(s, 1)) % types.Token(pullInt(s, 2)))
			},
			"__pow": func(s State) int {
				return s.Push(types.Token(math.Pow(float64(types.Token(pullInt(s, 1))), float64(types.Token(pullInt(s, 2))))))
			},
			"__unm": func(s State) int {
				return s.Push(-types.Token(pullInt(s, 1)))
			},
		},
	}
}
