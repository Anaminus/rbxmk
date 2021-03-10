package library

import (
	"os"
	"path/filepath"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
)

func init() { register(OS, 10) }

var OS = rbxmk.Library{Name: "os", Open: openOS, Dump: dumpOS}

func openOS(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 4)
	lib.RawSetString("expand", s.WrapFunc(func(s rbxmk.State) int {
		path := s.CheckString(1)
		expanded := s.World.Expand(path)
		s.L.Push(lua.LString(expanded))
		return 1

	}))
	lib.RawSetString("getenv", s.WrapFunc(func(s rbxmk.State) int {
		value, ok := os.LookupEnv(s.CheckString(1))
		if ok {
			s.L.Push(lua.LString(value))
		} else {
			s.L.Push(lua.LNil)
		}
		return 1
	}))
	lib.RawSetString("join", s.WrapFunc(func(s rbxmk.State) int {
		j := make([]string, s.Count())
		for i := 1; i <= s.Count(); i++ {
			j[i-1] = s.CheckString(i)
		}
		filename := filepath.Join(j...)
		s.L.Push(lua.LString(filename))
		return 1
	}))
	lib.RawSetString("split", s.WrapFunc(func(s rbxmk.State) int {
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
	}))
	return lib
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
				},
				"getenv": dump.Function{
					Parameters: dump.Parameters{
						{Name: "name", Type: dt.Optional{T: dt.Prim("string")}},
					},
					Returns: dump.Parameters{
						{Type: dt.Or{dt.Prim("string"), dt.Array{T: dt.Prim("string")}}},
					},
				},
				"join": dump.Function{
					Parameters: dump.Parameters{
						{Name: "...", Type: dt.Prim("string")},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("string")},
					},
				},
				"split": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim("string")},
						{Name: "...", Type: dt.Prim("string")},
					},
					Returns: dump.Parameters{
						{Name: "...", Type: dt.Prim("string")},
					},
					CanError: true,
				},
			},
		},
	}
}
