package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func String() Type {
	return Type{
		Name: "string",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.String))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LString); ok {
				return types.String(n), nil
			}
			return nil, TypeError(nil, 0, "string")
		},
	}
}

func BinaryString() Type {
	return Type{
		Name: "BinaryString",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.BinaryString))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LString); ok {
				return types.BinaryString(n), nil
			}
			return nil, TypeError(nil, 0, "string")
		},
	}
}

func ProtectedString() Type {
	return Type{
		Name: "ProtectedString",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.ProtectedString))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LString); ok {
				return types.ProtectedString(n), nil
			}
			return nil, TypeError(nil, 0, "string")
		},
	}
}

func Content() Type {
	return Type{
		Name: "Content",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.Content))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LString); ok {
				return types.Content(n), nil
			}
			return nil, TypeError(nil, 0, "string")
		},
	}
}

func SharedString() Type {
	return Type{
		Name: "SharedString",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.SharedString))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LString); ok {
				return types.SharedString(n), nil
			}
			return nil, TypeError(nil, 0, "string")
		},
	}
}
