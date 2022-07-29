package library

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
)

func init() { register(Table) }

var Table = rbxmk.Library{
	Name:       "table",
	ImportedAs: "table",
	Priority:   10,
	Open:       openTable,
	Dump:       dumpTable,
}

func openTable(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 6)
	lib.RawSetString("clear", s.WrapFunc(tableClear))
	lib.RawSetString("create", s.WrapFunc(tableCreate))
	lib.RawSetString("find", s.WrapFunc(tableFind))
	lib.RawSetString("move", s.WrapFunc(tableMove))
	lib.RawSetString("pack", s.WrapFunc(tablePack))
	lib.RawSetString("unpack", s.WrapFunc(tableUnpack))
	return lib
}

func tableClear(s rbxmk.State) int {
	t := s.CheckTable(1)
	t.Clear()
	return 0
}

func tableCreate(s rbxmk.State) int {
	cap := int(s.CheckInt(1))
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
	t := s.CheckTable(1)
	v := s.L.Get(2)
	init := s.OptInt(3, 1)
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
	a1 := s.CheckTable(1)
	f := s.CheckInt(2)
	e := s.CheckInt(3)
	t := s.CheckInt(4)
	var a2 *lua.LTable
	if s.L.Get(5) == lua.LNil {
		a2 = a1
	} else {
		a2 = s.CheckTable(5)
	}
	if e >= f {
		const LUA_MAXINTEGER = 1<<31 - 1
		if !(f > 0 || e < LUA_MAXINTEGER+f) {
			return s.ArgError(3, "too many elements to move")
		}
		n := e - f + 1
		if !(t <= LUA_MAXINTEGER-n+1) {
			return s.ArgError(4, "destination wrap around")
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
	n := s.Count()
	t := s.L.CreateTable(n, 1)
	for i := n; i >= 1; i-- {
		t.RawSetInt(i, s.L.Get(i))
	}
	t.RawSetString("n", lua.LNumber(n))
	s.L.Push(t)
	return 1
}

func tableUnpack(s rbxmk.State) int {
	t := s.CheckTable(1)
	i := s.OptInt(2, 1)
	j := s.OptInt(3, t.Len())
	for k := i; k <= j; k++ {
		s.L.Push(t.RawGetInt(k))
	}
	r := j - i + 1
	if r < 0 {
		return 0
	}
	return r
}

func dumpTable(s rbxmk.State) dump.Library {
	return dump.Library{
		Struct: dump.Struct{
			Fields: dump.Fields{
				"clear": dump.Function{
					Parameters: dump.Parameters{
						{Name: "t", Type: dt.Prim(rtypes.T_LuaTable)},
					},
					Summary:     "Libraries/table:Fields/clear/Summary",
					Description: "Libraries/table:Fields/clear/Description",
				},
				"create": dump.Function{
					Parameters: dump.Parameters{
						{Name: "cap", Type: dt.Prim(rtypes.T_LuaInteger)},
						{Name: "value", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_LuaTable)},
					},
					Summary:     "Libraries/table:Fields/create/Summary",
					Description: "Libraries/table:Fields/create/Description",
				},
				"find": dump.Function{
					Parameters: dump.Parameters{
						{Name: "t", Type: dt.Prim(rtypes.T_LuaTable)},
						{Name: "value", Type: dt.Prim(rtypes.T_Any)},
						{Name: "init", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger)), Default: `1`},
					},
					Returns: dump.Parameters{
						{Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger))},
					},
					Summary:     "Libraries/table:Fields/find/Summary",
					Description: "Libraries/table:Fields/find/Description",
				},
				"move": dump.Function{
					Parameters: dump.Parameters{
						{Name: "a1", Type: dt.Prim(rtypes.T_LuaTable)},
						{Name: "f", Type: dt.Prim(rtypes.T_LuaInteger)},
						{Name: "e", Type: dt.Prim(rtypes.T_LuaInteger)},
						{Name: "t", Type: dt.Prim(rtypes.T_LuaInteger)},
						{Name: "a2", Type: dt.Optional(dt.Prim(rtypes.T_LuaTable))},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_LuaTable)},
					},
					CanError:    true,
					Summary:     "Libraries/table:Fields/move/Summary",
					Description: "Libraries/table:Fields/move/Description",
				},
				"pack": dump.Function{
					Parameters: dump.Parameters{
						{Name: "...", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_LuaTable)},
					},
					Summary:     "Libraries/table:Fields/pack/Summary",
					Description: "Libraries/table:Fields/pack/Description",
				},
				"unpack": dump.Function{
					Parameters: dump.Parameters{
						{Name: "list", Type: dt.Prim(rtypes.T_LuaTable)},
						{Name: "i", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger))},
						{Name: "j", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger))},
					},
					Returns: dump.Parameters{
						{Name: "...", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
					},
					Summary:     "Libraries/table:Fields/unpack/Summary",
					Description: "Libraries/table:Fields/unpack/Description",
				},
			},
			Summary:     "Libraries/table:Summary",
			Description: "Libraries/table:Description",
		},
	}
}
