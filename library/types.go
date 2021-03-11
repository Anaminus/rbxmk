package library

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
)

func setUserdata(s rbxmk.State, t string) int {
	v := s.Pull(1, t)
	u := s.L.NewUserData(v)
	s.L.SetMetatable(u, s.L.GetTypeMetatable(t))
	s.L.Push(u)
	return 1
}

func init() { register(Types, 10) }

var Types = rbxmk.Library{Name: "types", Open: openTypes, Dump: dumpTypes}

func openTypes(s rbxmk.State) *lua.LTable {
	exprims := s.Reflectors(rbxmk.Exprim)
	lib := s.L.CreateTable(0, len(exprims))
	for _, t := range exprims {
		name := t.Name
		lib.RawSetString(t.Name, s.WrapFunc(func(s rbxmk.State) int {
			return setUserdata(s, name)
		}))
	}
	return lib
}

func dumpTypes(s rbxmk.State) dump.Library {
	exprims := s.Reflectors(rbxmk.Exprim)
	lib := dump.Library{Struct: dump.Struct{Fields: make(dump.Fields, len(exprims))}}
	for _, t := range exprims {
		lib.Struct.Fields[t.Name] = dump.Property{ValueType: dt.Prim("exprim"), ReadOnly: true}
	}
	lib.Types = dump.TypeDefs{
		"exprim": {
			Properties: dump.Properties{
				"Value": dump.Property{
					ValueType: dt.Prim("any"),
					ReadOnly:  true,
				},
			},
		},
	}

	return lib
}
