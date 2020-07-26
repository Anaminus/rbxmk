package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func pullString(s State, n int) string {
	switch v := s.L.CheckAny(n).(type) {
	case lua.LString:
		return string(v)
	case *lua.LUserData:
		if v, ok := v.Value.(types.Stringlike); ok {
			return string(v.Stringlike())
		}
	}
	panic("expected string-like type")
}

func String() Type {
	return Type{
		Name: "string",
		ReflectTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.String))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			if n, ok := lvs[0].(lua.LString); ok {
				return types.String(n), nil
			}
			return nil, TypeError(nil, 0, "string")
		},
	}
}

func BinaryString() Type {
	return Type{
		Name:  "BinaryString",
		Flags: Exprim,
		ReflectTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.BinaryString))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LString:
				return types.BinaryString(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("BinaryString") {
					if v, ok := v.Value.(types.BinaryString); ok {
						return v, nil
					}
				}
			}
			return nil, TypeError(nil, 0, "BinaryString")
		},
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString("BinaryString: " + pullString(s, 1)))
				return 1
			},
			"__call": func(s State) int {
				return s.Push(types.String(pullString(s, 1)))
			},
			"__eq": func(s State) int {
				return s.Push(types.Bool(pullString(s, 1) == pullString(s, 2)))
			},
			"__lt": func(s State) int {
				return s.Push(types.Bool(pullString(s, 1) < pullString(s, 2)))
			},
			"__le": func(s State) int {
				return s.Push(types.Bool(pullString(s, 1) <= pullString(s, 2)))
			},
			"__concat": func(s State) int {
				return s.Push(types.String(pullString(s, 1) + pullString(s, 2)))
			},
		},
	}
}

func ProtectedString() Type {
	return Type{
		Name:  "ProtectedString",
		Flags: Exprim,
		ReflectTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.ProtectedString))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LString:
				return types.ProtectedString(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("ProtectedString") {
					if v, ok := v.Value.(types.ProtectedString); ok {
						return v, nil
					}
				}
			}
			return nil, TypeError(nil, 0, "ProtectedString")
		},
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString("ProtectedString: " + pullString(s, 1)))
				return 1
			},
			"__call": func(s State) int {
				return s.Push(types.String(pullString(s, 1)))
			},
			"__eq": func(s State) int {
				return s.Push(types.Bool(pullString(s, 1) == pullString(s, 2)))
			},
			"__lt": func(s State) int {
				return s.Push(types.Bool(pullString(s, 1) < pullString(s, 2)))
			},
			"__le": func(s State) int {
				return s.Push(types.Bool(pullString(s, 1) <= pullString(s, 2)))
			},
			"__concat": func(s State) int {
				return s.Push(types.String(pullString(s, 1) + pullString(s, 2)))
			},
		},
	}
}

func Content() Type {
	return Type{
		Name:  "Content",
		Flags: Exprim,
		ReflectTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.Content))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LString:
				return types.Content(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("Content") {
					if v, ok := v.Value.(types.Content); ok {
						return v, nil
					}
				}
			}
			return nil, TypeError(nil, 0, "Content")
		},
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString("Content: " + pullString(s, 1)))
				return 1
			},
			"__call": func(s State) int {
				return s.Push(types.String(pullString(s, 1)))
			},
			"__eq": func(s State) int {
				return s.Push(types.Bool(pullString(s, 1) == pullString(s, 2)))
			},
			"__lt": func(s State) int {
				return s.Push(types.Bool(pullString(s, 1) < pullString(s, 2)))
			},
			"__le": func(s State) int {
				return s.Push(types.Bool(pullString(s, 1) <= pullString(s, 2)))
			},
			"__concat": func(s State) int {
				return s.Push(types.String(pullString(s, 1) + pullString(s, 2)))
			},
		},
	}
}

func SharedString() Type {
	return Type{
		Name:  "SharedString",
		Flags: Exprim,
		ReflectTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.SharedString))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LString:
				return types.SharedString(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("SharedString") {
					if v, ok := v.Value.(types.SharedString); ok {
						return v, nil
					}
				}
			}
			return nil, TypeError(nil, 0, "SharedString")
		},
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString("SharedString: " + pullString(s, 1)))
				return 1
			},
			"__call": func(s State) int {
				return s.Push(types.String(pullString(s, 1)))
			},
			"__eq": func(s State) int {
				return s.Push(types.Bool(pullString(s, 1) == pullString(s, 2)))
			},
			"__lt": func(s State) int {
				return s.Push(types.Bool(pullString(s, 1) < pullString(s, 2)))
			},
			"__le": func(s State) int {
				return s.Push(types.Bool(pullString(s, 1) <= pullString(s, 2)))
			},
			"__concat": func(s State) int {
				return s.Push(types.String(pullString(s, 1) + pullString(s, 2)))
			},
		},
	}
}
