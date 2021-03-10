package library

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	reflect "github.com/anaminus/rbxmk/library/http"
	"github.com/anaminus/rbxmk/rtypes"
)

func init() { register(HTTP, 10) }

var HTTP = rbxmk.Library{Name: "http", Open: openHTTP, Dump: dumpHTTP}

func openHTTP(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 1)
	lib.RawSetString("request", s.WrapFunc(httpRequest))

	for _, f := range reflect.All() {
		r := f()
		s.RegisterReflector(r)
		s.ApplyReflector(r, lib)
	}

	return lib
}

func httpRequest(s rbxmk.State) int {
	options := s.Pull(1, "HTTPOptions").(rtypes.HTTPOptions)
	request, err := rbxmk.BeginHTTPRequest(s.World, options)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return s.Push(request)
}

func dumpHTTP(s rbxmk.State) dump.Library {
	lib := dump.Library{
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
		Types: dump.TypeDefs{},
	}
	for _, f := range reflect.All() {
		r := f()
		lib.Types[r.Name] = r.DumpAll()
	}
	return lib
}
