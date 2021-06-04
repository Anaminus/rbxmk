package library

import (
	"os"
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
)

func init() { register(OS, 10) }

var OS = rbxmk.Library{
	Name:       "os",
	ImportedAs: "os",
	Open:       openOS,
	Dump:       dumpOS,
}

func openOS(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 1)
	lib.RawSetString("getenv", s.WrapFunc(osGetenv))
	return lib
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

func dumpOS(s rbxmk.State) dump.Library {
	return dump.Library{
		Struct: dump.Struct{
			Fields: dump.Fields{
				"getenv": dump.Function{
					Parameters: dump.Parameters{
						{Name: "name", Type: dt.Optional{T: dt.Prim("string")}},
					},
					Returns: dump.Parameters{
						{Type: dt.Or{dt.Optional{T: dt.Prim("string")}, dt.Dictionary{V: dt.Prim("string")}}},
					},
					Summary:     "Libraries/os:Fields/getenv/Summary",
					Description: "Libraries/os:Fields/getenv/Description",
				},
			},
			Summary:     "Libraries/os:Summary",
			Description: "Libraries/os:Description",
		},
	}
}
