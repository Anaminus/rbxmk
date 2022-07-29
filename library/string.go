package library

import (
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
)

func init() { register(String) }

var String = rbxmk.Library{
	Name:       "string",
	ImportedAs: "string",
	Priority:   10,
	Open:       openString,
	Dump:       dumpString,
}

func openString(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 1)
	lib.RawSetString("split", s.WrapFunc(stringSplit))
	return lib
}

func stringSplit(s rbxmk.State) int {
	str := s.CheckString(1)
	if str == "" && s.L.Get(2) == lua.LNil {
		t := s.L.CreateTable(1, 0)
		t.RawSetInt(1, lua.LString(""))
		s.L.Push(t)
		return 1
	}
	sep := s.OptString(2, "")
	a := strings.Split(str, sep)
	t := s.L.CreateTable(len(a), 0)
	for i, v := range a {
		t.RawSetInt(i+1, lua.LString(v))
	}
	s.L.Push(t)
	return 1
}

func dumpString(s rbxmk.State) dump.Library {
	return dump.Library{
		Struct: dump.Struct{
			Fields: dump.Fields{
				"split": dump.Function{
					Parameters: dump.Parameters{
						{Name: "s", Type: dt.Prim(rtypes.T_LuaString)},
						{Name: "sep", Type: dt.Optional(dt.Prim(rtypes.T_LuaString)), Default: `","`},
					},
					Returns: dump.Parameters{
						{Type: dt.Array(dt.Prim(rtypes.T_LuaString))},
					},
					Summary:     "Libraries/string:Fields/split/Summary",
					Description: "Libraries/string:Fields/split/Description",
				},
			},
			Summary:     "Libraries/string:Summary",
			Description: "Libraries/string:Description",
		},
	}
}
