package library

import (
	"fmt"
	"os"
	"path/filepath"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
)

func init() { register(OS, 10) }

var OS = rbxmk.Library{
	Name: "os",
	Open: func(s rbxmk.State) *lua.LTable {
		lib := s.L.CreateTable(0, 4)
		lib.RawSetString("expand", s.WrapFunc(func(s rbxmk.State) int {
			path := s.CheckString(1)
			expanded := OSLibrary{World: s.World}.Expand(path)
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
			components, err := OSLibrary{World: s.World}.Split(path, components...)
			if err != nil {
				return s.RaiseError("%s", err)
			}
			for _, comp := range components {
				s.L.Push(lua.LString(comp))
			}
			return n - 1
		}))
		return lib
	},
}

// OSLibrary provides additional functions to the standard os library.
type OSLibrary struct {
	*rbxmk.World
}

// Expand expands a string containing predefined variables.
func (l OSLibrary) Expand(path string) string {
	return os.Expand(path, func(v string) string {
		switch v {
		case "script_name", "sn":
			if entry, ok := l.PeekFile(); ok {
				if entry.Path == "" {
					return ""
				}
				path, _ := filepath.Abs(entry.Path)
				return filepath.Base(path)
			}
		case "script_directory", "script_dir", "sd":
			if entry, ok := l.PeekFile(); ok {
				if entry.Path == "" {
					return ""
				}
				path, _ := filepath.Abs(entry.Path)
				return filepath.Dir(path)
			}
		case "root_script_directory", "root_script_dir", "rsd":
			return l.RootDir()
		case "working_directory", "working_dir", "wd":
			wd, _ := os.Getwd()
			return wd
		case "temp_directory", "temp_dir", "tmp":
			return l.TempDir()
		}
		return ""
	})
}

// Split returns the components of a file path.
func (l OSLibrary) Split(path string, components ...string) ([]string, error) {
	parts := make([]string, len(components))
	for i, comp := range components {
		var result string
		switch comp {
		case "dir":
			result = filepath.Dir(path)
		case "base":
			result = filepath.Base(path)
		case "ext":
			result = filepath.Ext(path)
		case "stem":
			result = filepath.Base(path)
			result = result[:len(result)-len(filepath.Ext(path))]
		case "fext":
			result = l.Ext(path)
			if result != "" && result != "." {
				result = "." + result
			}
		case "fstem":
			ext := l.Ext(path)
			if ext != "" && ext != "." {
				ext = "." + ext
			}
			result = filepath.Base(path)
			result = result[:len(result)-len(ext)]
		default:
			return nil, fmt.Errorf("unknown argument %q", comp)
		}
		parts[i] = result
	}
	return parts, nil
}
