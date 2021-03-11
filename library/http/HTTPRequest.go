package reflect

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(HTTPRequest) }
func HTTPRequest() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "HTTPRequest",
		PushTo:   rbxmk.PushPtrTypeTo("HTTPRequest"),
		PullFrom: rbxmk.PullTypeFrom("HTTPRequest"),
		Members: rbxmk.Members{
			"Resolve": {Method: true,
				Get: func(s rbxmk.State, v types.Value) int {
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
						CanError: true,
					}
				},
			},
			"Cancel": {Method: true,
				Get: func(s rbxmk.State, v types.Value) int {
					req := v.(*rbxmk.HTTPRequest)
					req.Cancel()
					return 0
				},
				Dump: func() dump.Value { return dump.Function{} },
			},
		},
	}
}
