package library

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
)

func init() { register(HTTP) }

var HTTP = rbxmk.Library{
	Name:       "http",
	ImportedAs: "http",
	Priority:   10,
	Open:       openHTTP,
	Dump:       dumpHTTP,
	Types: []func() rbxmk.Reflector{
		reflect.HTTPHeaders,
		reflect.HTTPOptions,
		reflect.HTTPRequest,
		reflect.HTTPResponse,
	},
}

func openHTTP(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 1)
	lib.RawSetString("request", s.WrapFunc(httpRequest))
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
					Summary:     "Libraries/http:Fields/request/Summary",
					Description: "Libraries/http:Fields/request/Description",
				},
			},
			Summary:     "Libraries/http:Summary",
			Description: "Libraries/http:Description",
		},
		Types: dump.TypeDefs{},
	}
	return lib
}
