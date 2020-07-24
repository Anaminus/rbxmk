package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Float() Type {
	return Type{
		Name: "float",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Float))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LNumber); ok {
				return types.Float(n), nil
			}
			return nil, TypeError(nil, 0, "float")
		},
	}
}

func Double() Type {
	return Type{
		Name: "double",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Double))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LNumber); ok {
				return types.Double(n), nil
			}
			return nil, TypeError(nil, 0, "double")
		},
	}
}

func Number() Type {
	return Type{
		Name: "number",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Double))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LNumber); ok {
				return types.Double(n), nil
			}
			return nil, TypeError(nil, 0, "number")
		},
	}
}

func Int() Type {
	return Type{
		Name: "int",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Int))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LNumber); ok {
				return types.Int(n), nil
			}
			return nil, TypeError(nil, 0, "int")
		},
	}
}

func Int64() Type {
	return Type{
		Name: "int64",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Int64))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LNumber); ok {
				return types.Int64(n), nil
			}
			return nil, TypeError(nil, 0, "int64")
		},
	}
}
