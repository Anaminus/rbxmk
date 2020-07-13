package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
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
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.(string); ok {
				return rbxfile.ValueString(v), nil
			}
			return nil, TypeError(nil, 0, "string")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueString); ok {
				return string(sv), nil
			}
			return nil, TypeError(nil, 0, "ValueString")
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
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.([]byte); ok {
				return rbxfile.ValueBinaryString(v), nil
			}
			return nil, TypeError(nil, 0, "[]byte")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueBinaryString); ok {
				return []byte(sv), nil
			}
			return nil, TypeError(nil, 0, "ValueBinaryString")
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
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.(string); ok {
				return rbxfile.ValueProtectedString(v), nil
			}
			return nil, TypeError(nil, 0, "string")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueProtectedString); ok {
				return string(sv), nil
			}
			return nil, TypeError(nil, 0, "ValueProtectedString")
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
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.(string); ok {
				return rbxfile.ValueContent(v), nil
			}
			return nil, TypeError(nil, 0, "string")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueContent); ok {
				return string(sv), nil
			}
			return nil, TypeError(nil, 0, "ValueContent")
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
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.([]byte); ok {
				return rbxfile.ValueSharedString(v), nil
			}
			return nil, TypeError(nil, 0, "[]byte")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueSharedString); ok {
				return []byte(sv), nil
			}
			return nil, TypeError(nil, 0, "ValueSharedString")
		},
	}
}
