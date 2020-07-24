package reflect

import (
	"fmt"

	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func ReflectVariantTo(s State, v types.Value) (lv lua.LValue, err error) {
	switch v := v.(type) {
	case nil:
		return lua.LNil, nil
	case types.Bool:
		return lua.LBool(v), nil
	case types.Numberlike:
		return lua.LNumber(v.Numberlike()), nil
	case types.Intlike:
		return lua.LNumber(v.Intlike()), nil
	case types.Stringlike:
		return lua.LString(v.Stringlike()), nil
	case rtypes.Array:
		typ := s.Type("Array")
		values, err := typ.ReflectTo(s, typ, v)
		if err != nil {
			return nil, err
		}
		return values[0], nil
	case rtypes.Dictionary:
		typ := s.Type("Dictionary")
		values, err := typ.ReflectTo(s, typ, v)
		if err != nil {
			return nil, err
		}
		return values[0], nil
	}
	typ := s.Type(v.Type())
	if typ.Name == "" {
		return nil, fmt.Errorf("unknown type %q", string(v.Type()))
	}
	if typ.ReflectTo == nil {
		return nil, fmt.Errorf("unable to cast %s to Variant", typ.Name)
	}
	values, err := typ.ReflectTo(s, typ, v)
	if err != nil {
		return nil, err
	}
	return values[0], nil
}

func ReflectVariantFrom(s State, lv lua.LValue) (v types.Value, err error) {
	switch lv := lv.(type) {
	case *lua.LNilType:
		return nil, nil
	case lua.LBool:
		return types.Bool(lv), nil
	case lua.LNumber:
		return types.Double(lv), nil
	case lua.LString:
		return types.String(lv), nil
	case *lua.LTable:
		if lv.Len() > 0 {
			arrayType := s.Type("Array")
			if v, err = arrayType.ReflectFrom(s, arrayType, lv); err == nil {
				return v, nil
			}
		}
		dictType := s.Type("Dictionary")
		v, err := dictType.ReflectFrom(s, dictType, lv)
		return v, err
	case *lua.LUserData:
		name, ok := s.L.GetMetaField(lv, "__type").(lua.LString)
		if !ok {
			return nil, fmt.Errorf("unable to cast %s to Variant", lv.Type().String())
		}
		typ := s.Type(string(name))
		if typ.Name == "" {
			return nil, fmt.Errorf("unknown type %q", string(name))
		}
		if typ.ReflectFrom == nil {
			return nil, fmt.Errorf("unable to cast %s to Variant", typ.Name)
		}
		v, err := typ.ReflectFrom(s, typ, lv)
		return v, err
	}
	return nil, fmt.Errorf("unable to cast %s to Variant", lv.Type().String())
}

// PullVariant gets from the Lua state the value at n, and reflects a value from
// it according to the Variant type.
func PullVariant(s State, n int) (v types.Value) {
	v, err := ReflectVariantFrom(s, s.L.CheckAny(n))
	if err != nil {
		s.L.ArgError(n, err.Error())
		return nil
	}
	return v
}

func Variant() Type {
	return Type{
		Name: "Variant",
		ReflectTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			lv, err := ReflectVariantTo(s, v)
			if err != nil {
				return nil, err
			}
			return []lua.LValue{lv}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			v, err = ReflectVariantFrom(s, lvs[0])
			return v, err
		},
	}
}
