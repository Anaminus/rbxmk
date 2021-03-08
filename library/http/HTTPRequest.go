package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(HTTPRequest) }
func HTTPRequest() Reflector {
	return Reflector{
		Name:     "HTTPRequest",
		PushTo:   rbxmk.PushTypeTo("HTTPRequest"),
		PullFrom: rbxmk.PullTypeFrom("HTTPRequest"),
		Metatable: Metatable{
			"__eq": func(s State) int {
				v := s.Pull(1, "HTTPRequest").(*rbxmk.HTTPRequest)
				op := s.Pull(2, "HTTPRequest").(*rbxmk.HTTPRequest)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: Members{
			"Resolve": Member{Method: true,
				Get: func(s State, v types.Value) int {
					req := v.(*rbxmk.HTTPRequest)
					resp, err := req.Resolve()
					if err != nil {
						return s.RaiseError("%s", err)
					}
					return s.Push(*resp)
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "resp", Type: dt.Prim("HTTPResponse")},
						},
					}
				},
			},
			"Cancel": Member{Method: true,
				Get: func(s State, v types.Value) int {
					req := v.(*rbxmk.HTTPRequest)
					req.Cancel()
					return 0
				},
				Dump: func() dump.Value { return dump.Function{} },
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
