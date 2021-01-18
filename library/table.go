package library

import (
	"github.com/anaminus/rbxmk"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(Table, 10) }

var Table = rbxmk.Library{
	Name: "table",
	Open: func(s rbxmk.State) *lua.LTable {
		lib := s.L.CreateTable(0, 6)
		lib.RawSetString("clear", s.WrapFunc(tableClear))
		lib.RawSetString("create", s.WrapFunc(tableCreate))
		lib.RawSetString("find", s.WrapFunc(tableFind))
		lib.RawSetString("move", s.WrapFunc(tableMove))
		lib.RawSetString("pack", s.WrapFunc(tablePack))
		lib.RawSetString("unpack", s.WrapFunc(tableUnpack))
		return lib
	},
}

func tableClear(s rbxmk.State) int {
	t := s.L.CheckTable(1)
	// TODO: Implement native clear function.
	t.ForEach(func(k, v lua.LValue) {
		t.RawSet(k, lua.LNil)
	})
	return 0
}

func tableCreate(s rbxmk.State) int {
	cap := int(s.L.CheckNumber(1))
	value := s.L.Get(2)
	t := s.L.CreateTable(cap, 0)
	if value != lua.LNil {
		for i := 1; i <= cap; i++ {
			t.RawSetInt(i, value)
		}
	}
	s.L.Push(t)
	return 1
}

func tableFind(s rbxmk.State) int {
	t := s.L.CheckTable(1)
	v := s.L.Get(2)
	init := s.L.OptInt(3, 1)
	if v != lua.LNil {
		for i, n := init, t.Len(); i <= n; i++ {
			if t.RawGetInt(i) == v {
				s.L.Push(lua.LNumber(i))
				return 1
			}
		}
	}
	s.L.Push(lua.LNil)
	return 1
}

func tableMove(s rbxmk.State) int {
	a1 := s.L.CheckTable(1)
	f := s.L.CheckInt(2)
	e := s.L.CheckInt(3)
	t := s.L.CheckInt(4)
	var a2 *lua.LTable
	if s.L.Get(5) == lua.LNil {
		a2 = a1
	} else {
		a2 = s.L.CheckTable(5)
	}
	if e >= f {
		const LUA_MAXINTEGER = 1<<31 - 1
		if !(f > 0 || e < LUA_MAXINTEGER+f) {
			s.L.ArgError(3, "too many elements to move")
		}
		n := e - f + 1
		if !(t <= LUA_MAXINTEGER-n+1) {
			s.L.ArgError(4, "destination wrap around")
		}
		if t > e || t <= f || a1 != a2 {
			for i := 0; i < n; i++ {
				v := a1.RawGetInt(f + i)
				a2.RawSetInt(t+i, v)
			}
		} else {
			for i := n - 1; i >= 0; i-- {
				v := a1.RawGetInt(f + i)
				a2.RawSetInt(t+i, v)
			}
		}
	}
	s.L.Push(a2)
	return 1
}

func tablePack(s rbxmk.State) int {
	n := s.L.GetTop()
	t := s.L.CreateTable(n, 1)
	for i := n; i >= 1; i-- {
		t.RawSetInt(i, s.L.Get(i))
	}
	t.RawSetString("n", lua.LNumber(n))
	s.L.Push(t)
	return 1
}

func tableUnpack(s rbxmk.State) int {
	t := s.L.CheckTable(1)
	i := s.L.OptInt(2, 1)
	j := s.L.OptInt(3, t.Len())
	for k := i; k <= j; k++ {
		s.L.Push(t.RawGetInt(k))
	}
	r := j - i + 1
	if r < 0 {
		return 0
	}
	return r
}