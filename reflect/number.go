package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Float() Reflector {
	return Reflector{
		Name:  "float",
		Flags: Exprim,
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Float))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
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
	}
}

func Double() Reflector {
	return Reflector{
		Name: "double",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Double))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
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
	}
}

func Number() Reflector {
	return Reflector{
		Name: "number",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Double))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			if n, ok := lvs[0].(lua.LNumber); ok {
				return types.Double(n), nil
			}
			return nil, TypeError(nil, 0, "number")
		},
	}
}

func Int() Reflector {
	return Reflector{
		Name:  "int",
		Flags: Exprim,
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Int))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
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
	}
}

func Int64() Reflector {
	return Reflector{
		Name:  "int64",
		Flags: Exprim,
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Int64))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
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
	}
}

func Token() Reflector {
	return Reflector{
		Name:  "token",
		Flags: Exprim,
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Token))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
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
	}
}
