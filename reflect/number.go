package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/yuin/gopher-lua"
)

func Float() Type {
	return Type{
		Name: "float",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(float32))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LNumber); ok {
				return float32(n), nil
			}
			return nil, TypeError(nil, 0, "float")
		},
	}
}

func Double() Type {
	return Type{
		Name: "double",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(float64))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LNumber); ok {
				return float64(n), nil
			}
			return nil, TypeError(nil, 0, "double")
		},
	}
}

func Number() Type {
	return Type{
		Name: "number",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(float64))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LNumber); ok {
				return float64(n), nil
			}
			return nil, TypeError(nil, 0, "number")
		},
	}
}

func Int() Type {
	return Type{
		Name: "int",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(int))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LNumber); ok {
				return int(n), nil
			}
			return nil, TypeError(nil, 0, "int")
		},
	}
}

func Int64() Type {
	return Type{
		Name: "int64",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(int64))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LNumber); ok {
				return int64(n), nil
			}
			return nil, TypeError(nil, 0, "int64")
		},
	}
}
