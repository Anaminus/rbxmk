package library

import (
	"github.com/anaminus/rbxmk"
	"github.com/yuin/gopher-lua"
)

func Base(s rbxmk.State) {
	s.L.SetGlobal("typeof", s.L.NewFunction(baseTypeof))
}

func baseTypeof(l *lua.LState) int {
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
