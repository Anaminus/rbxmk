package reflect

import (
	"fmt"

	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/types"
	"github.com/yuin/gopher-lua"
)

func ReflectVariantTo(s State, v Value) (lv lua.LValue, t Type, err error) {
	switch v := v.(type) {
	case nil:
		return lua.LNil, s.Type("nil"), nil
	case bool:
		return lua.LBool(v), s.Type("bool"), nil
	case uint8:
		return lua.LNumber(v), s.Type("number"), nil
	case uint16:
		return lua.LNumber(v), s.Type("number"), nil
	case uint32:
		return lua.LNumber(v), s.Type("number"), nil
	case uint64:
		return lua.LNumber(v), s.Type("number"), nil
	case uint:
		return lua.LNumber(v), s.Type("number"), nil
	case int8:
		return lua.LNumber(v), s.Type("number"), nil
	case int16:
		return lua.LNumber(v), s.Type("number"), nil
	case int32:
		return lua.LNumber(v), s.Type("number"), nil
	case int64:
		return lua.LNumber(v), s.Type("number"), nil
	case int:
		return lua.LNumber(v), s.Type("number"), nil
	case float32:
		return lua.LNumber(v), s.Type("number"), nil
	case float64:
		return lua.LNumber(v), s.Type("number"), nil
	case string:
		return lua.LString(v), s.Type("string"), nil
	case []byte:
		return lua.LString(v), s.Type("string"), nil
	case []rune:
		return lua.LString(v), s.Type("string"), nil
	case []Value:
		typ := s.Type("Array")
		values, err := typ.ReflectTo(s, typ, v)
		if err != nil {
			return nil, typ, err
		}
		return values[0], typ, nil
	case map[string]Value:
		typ := s.Type("Dictionary")
		values, err := typ.ReflectTo(s, typ, v)
		if err != nil {
			return nil, typ, err
		}
		return values[0], typ, nil
	case types.Stringlike:
		if v, ok := v.Stringlike(); ok {
			return lua.LString(v), s.Type("string"), nil
		}
	case types.Floatlike:
		if v, ok := v.Floatlike(); ok {
			return lua.LNumber(v), s.Type("number"), nil
		}
	case types.Intlike:
		if v, ok := v.Intlike(); ok {
			return lua.LNumber(v), s.Type("int"), nil
		}
	case TValue:
		typ := s.Type(v.Type)
		if typ.ReflectTo == nil {
			return nil, typ, fmt.Errorf("unknown type %q", v.Type)
		}
		values, err := typ.ReflectTo(s, typ, v.Value)
		if err != nil {
			return nil, typ, err
		}
		if len(values) > 0 {
			return values[0], typ, nil
		}
	}
	return nil, s.Type("Variant"), fmt.Errorf("unable to cast value to Variant")
}

func ReflectVariantFrom(s State, lv lua.LValue) (v Value, t Type, err error) {
	switch lv := lv.(type) {
	case *lua.LNilType:
		return nil, s.Type("nil"), nil
	case lua.LBool:
		return bool(lv), s.Type("bool"), nil
	case lua.LNumber:
		return float64(lv), s.Type("number"), nil
	case lua.LString:
		return string(lv), s.Type("string"), nil
	case *lua.LTable:
		if lv.Len() > 0 {
			arrayType := s.Type("Array")
			if v, err = arrayType.ReflectFrom(s, arrayType, lv); err == nil {
				return v, arrayType, nil
			}
		}
		dictType := s.Type("Dictionary")
		v, err := dictType.ReflectFrom(s, dictType, lv)
		return v, dictType, err
	case *lua.LUserData:
		name, ok := s.L.GetMetaField(lv, "__type").(lua.LString)
		if !ok {
			return nil, s.Type("Variant"), fmt.Errorf("unable to cast %s to Variant", lv.Type().String())
		}
		typ := s.Type(string(name))
		if typ.Name == "" {
			return nil, s.Type("Variant"), fmt.Errorf("unknown type %q", string(name))
		}
		if typ.ReflectFrom == nil {
			return nil, typ, fmt.Errorf("unable to cast %s to Variant", typ.Name)
		}
		v, err := typ.ReflectFrom(s, typ, lv)
		return v, typ, err
	}
	return nil, s.Type("Variant"), fmt.Errorf("unable to cast %s to Variant", lv.Type().String())
}

// PullVariant gets from the Lua state the value at n, and reflects a value from
// it according to the Variant type.
func PullVariant(s State, n int) (v Value, t Type) {
	v, t, err := ReflectVariantFrom(s, s.L.CheckAny(n))
	if err != nil {
		s.L.ArgError(n, err.Error())
		return nil, t
	}
	return v, t
}

func Variant() Type {
	return Type{
		Name: "Variant",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			lv, _, err := ReflectVariantTo(s, v)
			if err != nil {
				return nil, err
			}
			return []lua.LValue{lv}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			v, _, err = ReflectVariantFrom(s, lvs[0])
			return v, err
		},
	}
}
