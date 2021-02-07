package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Tuple) }
func Tuple() Reflector {
	return Reflector{
		Name:  "Tuple",
		Count: -1,
		PushTo: func(s State, v types.Value) (lvs []lua.LValue, err error) {
			values := v.(rtypes.Tuple)
			lvs = make([]lua.LValue, len(values))
			variantRfl := s.Reflector("Variant")
			for i, value := range values {
				lv, err := variantRfl.PushTo(s, value)
				if err != nil {
					return nil, err
				}
				lvs[i] = lv[0]
			}
			return lvs, nil
		},
		PullFrom: func(s State, lvs ...lua.LValue) (v types.Value, err error) {
			vs := make(rtypes.Tuple, len(lvs))
			variantRfl := s.Reflector("Variant")
			for i, lv := range lvs {
				v, err := variantRfl.PullFrom(s, lv)
				if err != nil {
					return nil, err
				}
				vs[i] = v
			}
			return vs, nil
		},
	}
}
