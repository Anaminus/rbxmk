package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Tuple() Type {
	return Type{
		Name:  "Tuple",
		Count: -1,
		ReflectTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			values := v.(rtypes.Tuple)
			lvs = make([]lua.LValue, len(values))
			variantType := s.Type("Variant")
			for i, value := range values {
				lv, err := variantType.ReflectTo(s, variantType, value)
				if err != nil {
					return nil, err
				}
				lvs[i] = lv[0]
			}
			return lvs, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			vs := make(rtypes.Tuple, len(lvs))
			variantType := s.Type("Variant")
			for i, lv := range lvs {
				v, err := variantType.ReflectFrom(s, variantType, lv)
				if err != nil {
					return nil, err
				}
				vs[i] = v
			}
			return vs, nil
		},
	}
}
