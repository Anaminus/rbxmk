package library

import (
	"os"
	"path/filepath"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
)

func init() { register(OS, 10) }

var OS = rbxmk.Library{Name: "os", Open: openOS}

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
		j := make([]string, s.L.GetTop())
		for i := 1; i <= s.L.GetTop(); i++ {
			j[i-1] = s.CheckString(i)
		}
		filename := filepath.Join(j...)
		s.L.Push(lua.LString(filename))
		return 1
	}))
	lib.RawSetString("split", s.WrapFunc(func(s rbxmk.State) int {
		path := s.CheckString(1)
		n := s.L.GetTop()
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
