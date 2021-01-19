package library

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	reflect "github.com/anaminus/rbxmk/library/roblox"
)

func init() { register(Roblox, 1) }

var Roblox = rbxmk.Library{
	Name: "",
	Open: func(s rbxmk.State) *lua.LTable {
		lib := s.L.CreateTable(0, 1)
		lib.RawSetString("typeof", s.L.NewFunction(robloxTypeof))

		for _, f := range reflect.All() {
			r := f()
			s.RegisterReflector(r)
			s.ApplyReflector(r, lib)
		}

		return lib
	},
}

func robloxTypeof(l *lua.LState) int {
	v := l.CheckAny(1)
	u, ok := v.(*lua.LUserData)
	if !ok {
		l.Push(lua.LString(v.Type().String()))
		return 1
	}
	t, ok := l.GetMetaField(u, "__type").(lua.LString)
	if !ok {
		l.Push(lua.LString(u.Type().String()))
		return 1
	}
	l.Push(lua.LString(t))
	return 1
}
