package library

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	reflect "github.com/anaminus/rbxmk/library/roblox"
)

func init() { register(Roblox, 1) }

var Roblox = rbxmk.Library{Name: "", Open: openRoblox}

func openRoblox(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 1)
	lib.RawSetString("typeof", s.WrapFunc(func(s rbxmk.State) int {
		v := s.L.CheckAny(1)
		t := s.Typeof(v)
		s.L.Push(lua.LString(t))
		return 1
	}))

	for _, f := range reflect.All() {
		r := f()
		s.RegisterReflector(r)
		s.ApplyReflector(r, lib)
	}

	return lib
}
