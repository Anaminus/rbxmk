package library

import (
	"path/filepath"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Path) }

var Path = rbxmk.Library{
	Name:     "path",
	Import:   []string{"path"},
	Priority: 10,
	Open:     openPath,
	Dump:     dumpPath,
}

func openPath(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 6)
	lib.RawSetString("clean", s.WrapFunc(pathClean))
	lib.RawSetString("expand", s.WrapFunc(pathExpand))
	lib.RawSetString("explode", s.WrapFunc(pathExplode))
	lib.RawSetString("join", s.WrapFunc(pathJoin))
	lib.RawSetString("rel", s.WrapFunc(pathRel))
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
	expanded, err := s.World.Expand(path)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	s.L.Push(lua.LString(expanded))
	return 1
}

func splitPathAll(s string) []string {
	if filepath.IsAbs(s) || s == "" || filepath.VolumeName(s) == s {
		return nil
	}
	a := []string{}
	for {
		s = filepath.Clean(s)
		d, b := filepath.Split(s)
		a = append([]string{b}, a...)
		if d == "" {
			// Terminate relative path.
			break
		} else if d == s {
			// Terminate absolute path.
			return nil
		}
		s = d
	}
	return a
}

func pathExplode(s rbxmk.State) int {
	path := s.CheckString(1)
	elements := splitPathAll(path)
	a := make(rtypes.Array, len(elements))
	for i, element := range elements {
		a[i] = types.String(element)
	}
	return s.PushTuple(a...)
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

func pathRel(s rbxmk.State) int {
	basePath := s.CheckString(1)
	targetPath := s.CheckString(2)
	relPath, err := filepath.Rel(basePath, targetPath)
	if err != nil {
		s.L.Push(lua.LNil)
		return 1
	}
	s.L.Push(lua.LString(relPath))
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
						{Name: "path", Type: dt.Prim(rtypes.T_LuaString)},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_LuaString)},
					},
					Summary:     "Libraries/path:Fields/clean/Summary",
					Description: "Libraries/path:Fields/clean/Description",
				},
				"expand": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim(rtypes.T_LuaString)},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_LuaString)},
					},
					CanError:    true,
					Summary:     "Libraries/path:Fields/expand/Summary",
					Description: "Libraries/path:Fields/expand/Description",
				},
				"explode": dump.Function{
					Hidden: true,
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim(rtypes.T_LuaString)},
					},
					Returns: dump.Parameters{
						{Name: "...", Type: dt.Prim(rtypes.T_LuaString)},
					},
					Summary:     "Libraries/path:Fields/explode/Summary",
					Description: "Libraries/path:Fields/explode/Description",
				},
				"join": dump.Function{
					Parameters: dump.Parameters{
						{Name: "...", Type: dt.Prim(rtypes.T_LuaString)},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_LuaString)},
					},
					Summary:     "Libraries/path:Fields/join/Summary",
					Description: "Libraries/path:Fields/join/Description",
				},
				"rel": dump.Function{
					Parameters: dump.Parameters{
						{Name: "basePath", Type: dt.Prim(rtypes.T_LuaString)},
						{Name: "targetPath", Type: dt.Prim(rtypes.T_LuaString)},
					},
					Returns: dump.Parameters{
						{Type: dt.Optional(dt.Prim(rtypes.T_LuaString))},
					},
					Summary:     "Libraries/path:Fields/rel/Summary",
					Description: "Libraries/path:Fields/rel/Description",
				},
				"split": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim(rtypes.T_LuaString)},
						{Name: "...", Type: dt.Prim(rtypes.T_LuaString)},
					},
					Returns: dump.Parameters{
						{Name: "...", Type: dt.Prim(rtypes.T_LuaString)},
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
