package library

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	reflect "github.com/anaminus/rbxmk/library/http"
	"github.com/anaminus/rbxmk/rtypes"
)

func init() { register(HTTPSource, 10) }

var HTTPSource = rbxmk.Library{
	Name: "http",
	Open: func(s rbxmk.State) *lua.LTable {
		lib := s.L.CreateTable(0, 1)
		lib.RawSetString("request", s.WrapFunc(func(s rbxmk.State) int {
			options := s.Pull(1, "HTTPOptions").(rtypes.HTTPOptions)
			request, err := rbxmk.BeginHTTPRequest(s.World, options)
			if err != nil {
				return s.RaiseError("%s", err)
			}
			return s.Push(request)
		}))

		for _, f := range reflect.All() {
			r := f()
			s.RegisterReflector(r)
			s.ApplyReflector(r, lib)
		}

		return lib
	},
	Dump: func(s rbxmk.State) dump.Library {
		return dump.Library{
			Struct: dump.Struct{
				Fields: dump.Fields{
					"request": dump.Function{
						Parameters: dump.Parameters{
							{Name: "options", Type: dt.Prim("HTTPOptions")},
						},
						Returns: dump.Parameters{
							{Name: "req", Type: dt.Prim("HTTPRequest")},
						},
					},
				},
			},
		}
	},
}
