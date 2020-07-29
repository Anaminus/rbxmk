package library

import (
	"github.com/anaminus/rbxmk"
	"github.com/yuin/gopher-lua"
)

func setUserdata(s rbxmk.State, t string) int {
	u := s.L.NewUserData()
	u.Value = s.Pull(1, t)
	s.L.SetMetatable(u, s.L.GetTypeMetatable(t))
	s.L.Push(u)
	return 1
}

var Types = rbxmk.Library{
	Name: "types",
	Open: func(s rbxmk.State) *lua.LTable {
		exprims := s.Types(rbxmk.Exprim)
		lib := s.L.CreateTable(0, len(exprims))
		for _, t := range exprims {
			name := t.Name
			lib.RawSetString(t.Name, s.WrapFunc(func(s rbxmk.State) int {
				return setUserdata(s, name)
			}))
		}
		return lib
	},
}
