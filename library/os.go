package library

import (
	"os"
	"path/filepath"
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
)

func init() { register(OS, 10) }

var OS = rbxmk.Library{Name: "os", Open: openOS, Dump: dumpOS}

func openOS(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 4)
	lib.RawSetString("expand", s.WrapFunc(osExpand))
	lib.RawSetString("getenv", s.WrapFunc(osGetenv))
	lib.RawSetString("join", s.WrapFunc(osJoin))
	lib.RawSetString("split", s.WrapFunc(osSplit))
	return lib
}

func osExpand(s rbxmk.State) int {
	path := s.CheckString(1)
	expanded := s.World.Expand(path)
	s.L.Push(lua.LString(expanded))
	return 1

}

func osGetenv(s rbxmk.State) int {
	switch lv := s.L.Get(1).(type) {
	case *lua.LNilType:
		vars := os.Environ()
		table := s.L.CreateTable(0, len(vars))
		for _, v := range vars {
			if i := strings.IndexByte(v, '='); i >= 0 {
				table.RawSetString(v[:i], lua.LString(v[i+1:]))
				continue
			}
			// Shouldn't happen, but just in case, set the whole variable to an
			// empty string.
			table.RawSetString(v, lua.LString(""))
		}
		s.L.Push(table)
		return 1
	case lua.LString:
		if value, ok := os.LookupEnv(s.CheckString(1)); ok {
			s.L.Push(lua.LString(value))
			return 1
		}
		s.L.Push(lua.LNil)
		return 1
	default:
		return s.TypeError(1, "string or nil", lv.Type().String())
	}
}

func osJoin(s rbxmk.State) int {
	j := make([]string, s.Count())
	for i := 1; i <= s.Count(); i++ {
		j[i-1] = s.CheckString(i)
	}
	filename := filepath.Join(j...)
	s.L.Push(lua.LString(filename))
	return 1
}

func osSplit(s rbxmk.State) int {
	path := s.CheckString(1)
	n := s.Count()
	components := make([]string, n-2)
	for i := 2; i <= n; i++ {
		components[i-2] = s.CheckString(i)
	}
	components, err := s.World.Split(path, components...)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	for _, comp := range components {
		s.L.Push(lua.LString(comp))
	}
	return n - 1
}

func dumpOS(s rbxmk.State) dump.Library {
	return dump.Library{
		Struct: dump.Struct{
			Fields: dump.Fields{
				"expand": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim("string")},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("string")},
					},
					Summary:     "libraries/os:Fields/path/Summary",
					Description: "libraries/os:Fields/path/Description",
				},
				"getenv": dump.Function{
					Parameters: dump.Parameters{
						{Name: "name", Type: dt.Optional{T: dt.Prim("string")}},
					},
					Returns: dump.Parameters{
						{Type: dt.Or{dt.Optional{T: dt.Prim("string")}, dt.Dictionary{V: dt.Prim("string")}}},
					},
					Summary:     "libraries/os:Fields/getenv/Summary",
					Description: "libraries/os:Fields/getenv/Description",
				},
				"join": dump.Function{
					Parameters: dump.Parameters{
						{Name: "...", Type: dt.Prim("string")},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("string")},
					},
					Summary:     "libraries/os:Fields/join/Summary",
					Description: "libraries/os:Fields/join/Description",
				},
				"split": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim("string")},
						{Name: "...", Type: dt.Prim("string")},
					},
					Returns: dump.Parameters{
						{Name: "...", Type: dt.Prim("string")},
					},
					CanError:    true,
					Summary:     "libraries/os:Fields/split/Summary",
					Description: "libraries/os:Fields/split/Description",
				},
			},
			Summary:     "libraries/os:Summary",
			Description: "libraries/os:Description",
		},
	}
}
