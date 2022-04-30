package library

import (
	"path/filepath"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
)

func init() { register(Path) }

var Path = rbxmk.Library{
	Name:       "path",
	ImportedAs: "path",
	Priority:   10,
	Open:       openPath,
	Dump:       dumpPath,
}

func openPath(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 4)
	lib.RawSetString("clean", s.WrapFunc(pathClean))
	lib.RawSetString("expand", s.WrapFunc(pathExpand))
	lib.RawSetString("join", s.WrapFunc(pathJoin))
	lib.RawSetString("split", s.WrapFunc(pathSplit))
	return lib
}

func pathClean(s rbxmk.State) int {
	path := s.CheckString(1)
	filename := filepath.Clean(path)
	s.L.Push(lua.LString(filename))
	return 1
}

func pathExpand(s rbxmk.State) int {
	path := s.CheckString(1)
	expanded := s.World.Expand(path)
	s.L.Push(lua.LString(expanded))
	return 1

}

func pathJoin(s rbxmk.State) int {
	j := make([]string, s.Count())
	for i := 1; i <= s.Count(); i++ {
		j[i-1] = s.CheckString(i)
	}
	filename := filepath.Join(j...)
	s.L.Push(lua.LString(filename))
	return 1
}

func pathSplit(s rbxmk.State) int {
	path := s.CheckString(1)
	n := s.Count()
	components := make([]string, n-1)
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

func dumpPath(s rbxmk.State) dump.Library {
	return dump.Library{
		Struct: dump.Struct{
			Fields: dump.Fields{
				"clean": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim("string")},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("string")},
					},
					Summary:     "Libraries/path:Fields/clean/Summary",
					Description: "Libraries/path:Fields/clean/Description",
				},
				"expand": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim("string")},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("string")},
					},
					Summary:     "Libraries/path:Fields/expand/Summary",
					Description: "Libraries/path:Fields/expand/Description",
				},
				"join": dump.Function{
					Parameters: dump.Parameters{
						{Name: "...", Type: dt.Prim("string")},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("string")},
					},
					Summary:     "Libraries/path:Fields/join/Summary",
					Description: "Libraries/path:Fields/join/Description",
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
					Summary:     "Libraries/path:Fields/split/Summary",
					Description: "Libraries/path:Fields/split/Description",
				},
			},
			Summary:     "Libraries/path:Summary",
			Description: "Libraries/path:Description",
		},
	}
}
