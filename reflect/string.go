package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/yuin/gopher-lua"
)

func String() Type {
	return Type{
		Name: "string",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(string))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LString); ok {
				return string(n), nil
			}
			return nil, TypeError(nil, 0, "string")
		},
	}
}

func BinaryString() Type {
	return Type{
		Name: "BinaryString",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.([]byte))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LString); ok {
				return []byte(n), nil
			}
			return nil, TypeError(nil, 0, "string")
		},
	}
}

func ProtectedString() Type {
	return Type{
		Name: "ProtectedString",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(string))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LString); ok {
				return string(n), nil
			}
			return nil, TypeError(nil, 0, "string")
		},
	}
}

func Content() Type {
	return Type{
		Name: "Content",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(string))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LString); ok {
				return string(n), nil
			}
			return nil, TypeError(nil, 0, "string")
		},
	}
}

func SharedString() Type {
	return Type{
		Name: "SharedString",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.([]byte))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LString); ok {
				return []byte(n), nil
			}
			return nil, TypeError(nil, 0, "string")
		},
	}
}
