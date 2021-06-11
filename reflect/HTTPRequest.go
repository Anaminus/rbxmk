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
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case **rbxmk.HTTPRequest:
				*p = v.(*rbxmk.HTTPRequest)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Methods: rbxmk.Methods{
			"Resolve": {
				Func: func(s rbxmk.State, v types.Value) int {
					req := v.(*rbxmk.HTTPRequest)
					resp, err := req.Resolve()
					if err != nil {
						return s.RaiseError("%s", err)
					}
					return s.Push(*resp)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "resp", Type: dt.Prim("HTTPResponse")},
						},
						CanError:    true,
						Summary:     "Types/HTTPRequest:Methods/Resolve/Summary",
						Description: "Types/HTTPRequest:Methods/Resolve/Description",
					}
				},
			},
			"Cancel": {
				Func: func(s rbxmk.State, v types.Value) int {
					req := v.(*rbxmk.HTTPRequest)
					req.Cancel()
					return 0
				},
				Dump: func() dump.Function {
					return dump.Function{
						Summary:     "Types/HTTPRequest:Methods/Cancel/Summary",
						Description: "Types/HTTPRequest:Methods/Cancel/Description",
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/HTTPRequest:Summary",
				Description: "Types/HTTPRequest:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			HTTPResponse,
		},
	}
}
