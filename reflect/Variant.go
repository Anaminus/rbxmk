package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func PushVariantTo(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
	switch v := v.(type) {
	case rtypes.NilType:
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
		rfl := s.MustReflector("Array")
		values, err := rfl.PushTo(s, v)
		if err != nil {
			return nil, err
		}
		return values, nil
	case rtypes.Dictionary:
		rfl := s.MustReflector("Dictionary")
		values, err := rfl.PushTo(s, v)
		if err != nil {
			return nil, err
		}
		return values, nil
	}
	rfl := s.Reflector(v.Type())
	if rfl.Name == "" {
		return nil, fmt.Errorf("unknown type %q", string(v.Type()))
	}
	if rfl.PushTo == nil {
		return nil, fmt.Errorf("unable to cast %s to Variant", rfl.Name)
	}
	values, err := rfl.PushTo(s, v)
	if err != nil {
		return nil, err
	}
	return values, nil
}

func PullVariantFrom(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
	switch lv := lv.(type) {
	case *lua.LNilType:
		return rtypes.Nil, nil
	case lua.LBool:
		return types.Bool(lv), nil
	case lua.LNumber:
		return types.Double(lv), nil
	case lua.LString:
		return types.String(lv), nil
	case *lua.LTable:
		if lv.Len() > 0 {
			arrayRfl := s.MustReflector("Array")
			if v, err = arrayRfl.PullFrom(s, lv); err == nil {
				return v, nil
			}
		}
		dictRfl := s.MustReflector("Dictionary")
		v, err := dictRfl.PullFrom(s, lv)
		return v, err
	case *lua.LUserData:
		name, ok := s.GetMetaField(lv, "__type").(lua.LString)
		if !ok {
			return nil, fmt.Errorf("unable to cast %s to Variant", lv.Type().String())
		}
		rfl := s.Reflector(string(name))
		if rfl.Name == "" {
			return nil, fmt.Errorf("unknown type %q", string(name))
		}
		if rfl.PullFrom == nil {
			return nil, fmt.Errorf("unable to cast %s to Variant", rfl.Name)
		}
		v, err := rfl.PullFrom(s, lv)
		return v, err
	}
	return nil, fmt.Errorf("unable to cast %s to Variant", lv.Type().String())
}

// PullVariant gets from the Lua state the value at n, and reflects a value from
// it according to the Variant type.
func PullVariant(s rbxmk.State, n int) (v types.Value) {
	v, err := PullVariantFrom(s.Context(), s.CheckAny(n))
	if err != nil {
		s.ArgError(n, err.Error())
		return nil
	}
	return v
}

func init() { register(Variant) }
func Variant() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "Variant",
		PushTo:   PushVariantTo,
		PullFrom: PullVariantFrom,
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Value:
				*p = v
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/Variant:Summary",
				Description: "Types/Variant:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			Dictionary,
		},
	}
}
