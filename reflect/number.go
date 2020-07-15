package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
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
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.(float32); ok {
				return rbxfile.ValueFloat(v), nil
			}
			return nil, TypeError(nil, 0, "float32")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueFloat); ok {
				return float32(sv), nil
			}
			return nil, TypeError(nil, 0, "ValueFloat")
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
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.(float64); ok {
				return rbxfile.ValueDouble(v), nil
			}
			return nil, TypeError(nil, 0, "float64")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueDouble); ok {
				return float64(sv), nil
			}
			return nil, TypeError(nil, 0, "ValueDouble")
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
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.(float64); ok {
				return rbxfile.ValueDouble(v), nil
			}
			return nil, TypeError(nil, 0, "float64")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueDouble); ok {
				return float64(sv), nil
			}
			return nil, TypeError(nil, 0, "ValueDouble")
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
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.(int); ok {
				return rbxfile.ValueInt(v), nil
			}
			return nil, TypeError(nil, 0, "int")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueInt); ok {
				return int(sv), nil
			}
			return nil, TypeError(nil, 0, "ValueInt")
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
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.(int64); ok {
				return rbxfile.ValueInt64(v), nil
			}
			return nil, TypeError(nil, 0, "int64")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueInt64); ok {
				return int64(sv), nil
			}
			return nil, TypeError(nil, 0, "ValueInt64")
		},
	}
}
