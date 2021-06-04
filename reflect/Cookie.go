package reflect

import (
	"net/http"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Cookie) }
func Cookie() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "Cookie",
		PushTo:   rbxmk.PushTypeTo("Cookie"),
		PullFrom: rbxmk.PullTypeFrom("Cookie"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "Cookie").(rtypes.Cookie)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "Cookie").(rtypes.Cookie)
				op := s.Pull(2, "Cookie").(rtypes.Cookie)
				s.L.Push(lua.LBool(v.Name == op.Name && v.Value == op.Value))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					cookie := v.(rtypes.Cookie)
					return s.Push(types.String(cookie.Name))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						ReadOnly:    true,
						Summary:     "Types/Cookie:Properties/Name/Summary",
						Description: "Types/Cookie:Properties/Name/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"from": rbxmk.Constructor{
				Func: func(s rbxmk.State) int {
					location := string(s.Pull(1, "string").(types.String))
					cookies, err := rbxmk.CookiesFrom(location)
					if err != nil {
						return s.RaiseError("unknown location %q", location)
					}
					if len(cookies) == 0 {
						s.Push(rtypes.Nil)
					}
					return s.Push(cookies)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						dump.Function{
							Parameters: dump.Parameters{
								{Name: "location", Type: dt.Prim("string"),
									Enums: dt.Enums{
										`"studio"`,
									},
								},
							},
							Returns: dump.Parameters{
								{Name: "cookies", Type: dt.Prim("Cookies")},
							},
							CanError:    true,
							Summary:     "Types/Cookie:Constructors/from/Summary",
							Description: "Types/Cookie:Constructors/from/Description",
						},
					}
				},
			},
			"new": rbxmk.Constructor{
				Func: func(s rbxmk.State) int {
					name := string(s.Pull(1, "string").(types.String))
					value := string(s.Pull(2, "string").(types.String))
					cookie := rtypes.Cookie{Cookie: &http.Cookie{Name: name, Value: value}}
					return s.Push(cookie)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						dump.Function{
							Parameters: dump.Parameters{
								{Name: "name", Type: dt.Prim("string")},
								{Name: "value", Type: dt.Prim("string")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("Cookie")},
							},
							Summary:     "Types/Cookie:Constructors/new/Summary",
							Description: "Types/Cookie:Constructors/new/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/Cookie:Operators/Eq/Summary",
						Description: "Types/Cookie:Operators/Eq/Description",
					},
				},
				Summary:     "Types/Cookie:Summary",
				Description: "Types/Cookie:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			String,
		},
	}
}
