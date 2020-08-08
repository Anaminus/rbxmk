package library

import (
	"github.com/anaminus/rbxmk"
	lua "github.com/yuin/gopher-lua"
)

func setUserdata(s rbxmk.State, t string) int {
	v := s.Pull(1, t)
	u := s.UserDataOf(v, t)
	s.L.Push(u)
	return 1
}

func init() { register(Types, 10) }

var Types = rbxmk.Library{
	Name: "types",
	Open: func(s rbxmk.State) *lua.LTable {
		exprims := s.Reflectors(rbxmk.Exprim)
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
