package reflect

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(HttpRequest) }
func HttpRequest() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     rtypes.T_HttpRequest,
		PushTo:   rbxmk.PushPtrTypeTo(rtypes.T_HttpRequest),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_HttpRequest),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case **rbxmk.HttpRequest:
				*p = v.(*rbxmk.HttpRequest)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Methods: rbxmk.Methods{
			"Resolve": {
				Func: func(s rbxmk.State, v types.Value) int {
					req := v.(*rbxmk.HttpRequest)
					resp, err := req.Resolve()
					if err != nil {
						return s.RaiseError("%s", err)
					}
					return s.Push(*resp)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "resp", Type: dt.Prim(rtypes.T_HttpResponse)},
						},
						CanError:    true,
						Summary:     "Types/HttpRequest:Methods/Resolve/Summary",
						Description: "Types/HttpRequest:Methods/Resolve/Description",
					}
				},
			},
			"Cancel": {
				Func: func(s rbxmk.State, v types.Value) int {
					req := v.(*rbxmk.HttpRequest)
					req.Cancel()
					return 0
				},
				Dump: func() dump.Function {
					return dump.Function{
						Summary:     "Types/HttpRequest:Methods/Cancel/Summary",
						Description: "Types/HttpRequest:Methods/Cancel/Description",
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/HttpRequest:Summary",
				Description: "Types/HttpRequest:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			HttpResponse,
		},
	}
}
