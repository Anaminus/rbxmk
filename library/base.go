package library

import (
	"github.com/anaminus/rbxmk"
	"github.com/yuin/gopher-lua"
)

func init() { register(Base) }

var Base = rbxmk.Library{
	Name: "",
	Open: func(s rbxmk.State) *lua.LTable {
		lib := s.L.CreateTable(0, 1)
		lib.RawSetString("typeof", s.L.NewFunction(baseTypeof))

		for _, r := range s.Reflectors(0) {
			if mt := s.CreateTypeMetatable(r); mt != nil {
				s.L.SetField(s.L.Get(lua.RegistryIndex), r.Name, mt)
			}
			if r.Constructors != nil {
				ctors := s.L.CreateTable(0, len(r.Constructors))
				for name, ctor := range r.Constructors {
					c := ctor
					ctors.RawSetString(name, s.L.NewFunction(func(l *lua.LState) int {
						return c(s)
					}))
				}
				lib.RawSetString(r.Name, ctors)
			}
			if r.Environment != nil {
				r.Environment(s, lib)
			}
		}

		return lib
	},
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
