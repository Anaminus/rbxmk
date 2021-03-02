package library

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
)

func setUserdata(s rbxmk.State, t string) int {
	v := s.Pull(1, t)
	u := s.L.NewUserData()
	u.Value = v
	s.L.SetMetatable(u, s.L.GetTypeMetatable(t))
	s.L.Push(u)
	return 1
}

func init() { register(Types, 10) }

var Types = rbxmk.Library{Name: "types", Open: openTypes}

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
