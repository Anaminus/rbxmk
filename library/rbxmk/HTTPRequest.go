package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/sources"
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
				v := s.Pull(1, "HTTPRequest").(*sources.HTTPRequest)
				op := s.Pull(2, "HTTPRequest").(*sources.HTTPRequest)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: Members{
			"Resolve": Member{Method: true, Get: func(s State, v types.Value) int {
				req := v.(*sources.HTTPRequest)
				resp, err := req.Resolve()
				if err != nil {
					return s.RaiseError("%s", err)
				}
				return s.Push(*resp)
			}},
			"Cancel": Member{Method: true, Get: func(s State, v types.Value) int {
				req := v.(*sources.HTTPRequest)
				req.Cancel()
				return 0
			}},
		},
	}
}
