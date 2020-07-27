package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

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
	}
}
