package library

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/anaminus/rbxmk"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(OS, 10) }

var OS = rbxmk.Library{
	Name: "os",
	Open: func(s rbxmk.State) *lua.LTable {
		lib := s.L.CreateTable(0, 6)
		lib.RawSetString("split", s.WrapFunc(osSplit))
		lib.RawSetString("join", s.WrapFunc(osJoin))
		lib.RawSetString("expand", s.WrapFunc(osExpand))
		lib.RawSetString("getenv", s.WrapFunc(osGetenv))
		lib.RawSetString("dir", s.WrapFunc(osDir))
		lib.RawSetString("stat", s.WrapFunc(osStat))
		return lib
	},
}

func osSplit(s rbxmk.State) int {
	path := s.CheckString(1)
	n := s.L.GetTop()
	for i := 2; i <= n; i++ {
		var result string
		switch comp := s.CheckString(i); comp {
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
			result = s.Ext(path)
			if result != "" && result != "." {
				result = "." + result
			}
		case "fstem":
			ext := s.Ext(path)
			if ext != "" && ext != "." {
				ext = "." + ext
			}
			result = filepath.Base(path)
			result = result[:len(result)-len(ext)]
		default:
			return s.RaiseError("unknown argument %q", comp)
		}
		s.L.Push(lua.LString(result))
	}
	return n - 1
}

func osJoin(s rbxmk.State) int {
	j := make([]string, s.L.GetTop())
	for i := 1; i <= s.L.GetTop(); i++ {
		j[i-1] = s.CheckString(i)
	}
	filename := filepath.Join(j...)
	s.L.Push(lua.LString(filename))
	return 1
}

func osExpand(s rbxmk.State) int {
	expanded := os.Expand(s.CheckString(1), func(v string) string {
		switch v {
		case "script_name", "sn":
			if fi, ok := s.PeekFile(); ok {
				path, _ := filepath.Abs(fi.Path)
				return filepath.Base(path)
			}
		case "script_directory", "script_dir", "sd":
			if fi, ok := s.PeekFile(); ok {
				path, _ := filepath.Abs(fi.Path)
				return filepath.Dir(path)
			}
		case "working_directory", "working_dir", "wd":
			wd, _ := os.Getwd()
			return wd
		case "temp_directory", "temp_dir", "tmp":
			return os.TempDir()
		}
		return ""
	})
	s.L.Push(lua.LString(expanded))
	return 1
}

func osGetenv(s rbxmk.State) int {
	value, ok := os.LookupEnv(s.CheckString(1))
	if ok {
		s.L.Push(lua.LString(value))
	} else {
		s.L.Push(lua.LNil)
	}
	return 1
}

func osDir(s rbxmk.State) int {
	dirname := s.CheckString(1)
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	tfiles := s.L.CreateTable(len(files), 0)
	for _, info := range files {
		tinfo := s.L.CreateTable(0, 4)
		tinfo.RawSetString("Name", lua.LString(info.Name()))
		tinfo.RawSetString("IsDir", lua.LBool(info.IsDir()))
		tinfo.RawSetString("Size", lua.LNumber(info.Size()))
		tinfo.RawSetString("ModTime", lua.LNumber(info.ModTime().Unix()))
		tfiles.Append(tinfo)
	}
	s.L.Push(tfiles)
	return 1
}

func osStat(s rbxmk.State) int {
	filename := s.CheckString(1)
	info, err := os.Stat(filename)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	tinfo := s.L.CreateTable(0, 4)
	tinfo.RawSetString("Name", lua.LString(info.Name()))
	tinfo.RawSetString("IsDir", lua.LBool(info.IsDir()))
	tinfo.RawSetString("Size", lua.LNumber(info.Size()))
	tinfo.RawSetString("ModTime", lua.LNumber(info.ModTime().Unix()))
	s.L.Push(tinfo)
	return 1
}
