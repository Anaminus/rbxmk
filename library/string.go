package library

import (
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
)

func init() { register(String, 10) }

var String = rbxmk.Library{Name: "string", Open: openString}

func openString(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 1)
	lib.RawSetString("split", s.WrapFunc(stringSplit))
	return lib
}

func stringSplit(s rbxmk.State) int {
	str := s.L.CheckString(1)
	if str == "" && s.L.Get(2) == lua.LNil {
		t := s.L.CreateTable(1, 0)
		t.RawSetInt(1, lua.LString(""))
		s.L.Push(t)
		return 1
	}
	sep := s.L.OptString(2, "")
	a := strings.Split(str, sep)
	t := s.L.CreateTable(len(a), 0)
	for i, v := range a {
		t.RawSetInt(i+1, lua.LString(v))
	}
	s.L.Push(t)
	return 1
}
