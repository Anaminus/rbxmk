package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Cookie) }
func Cookie() Reflector {
	return Reflector{
		Name:     "Cookie",
		PushTo:   rbxmk.PushTypeTo("Cookie"),
		PullFrom: rbxmk.PullTypeFrom("Cookie"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Cookie").(rtypes.Cookie)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Cookie").(rtypes.Cookie)
				op := s.Pull(2, "Cookie").(rtypes.Cookie)
				s.L.Push(lua.LBool(v.Name == op.Name && v.Value == op.Value))
				return 1
			},
		},
		Members: Members{
			"Name": Member{
				Get: func(s State, v types.Value) int {
					cookie := v.(rtypes.Cookie)
					return s.Push(types.String(cookie.Name))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string"), ReadOnly: true} },
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
