package sources

import (
	"os"
	"path/filepath"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(FS) }
func FS() rbxmk.Source {
	return rbxmk.Source{
		Name: "fs",
		Library: rbxmk.Library{
			Open: func(s rbxmk.State) *lua.LTable {
				lib := s.L.CreateTable(0, 6)
				lib.RawSetString("dir", s.WrapFunc(fsDir))
				lib.RawSetString("mkdir", s.WrapFunc(fsMkDir))
				lib.RawSetString("read", s.WrapFunc(fsRead))
				lib.RawSetString("remove", s.WrapFunc(fsRemove))
				lib.RawSetString("rename", s.WrapFunc(fsRename))
				lib.RawSetString("stat", s.WrapFunc(fsStat))
				lib.RawSetString("write", s.WrapFunc(fsWrite))
				return lib
			},
		},
	}
}

func fsDir(s rbxmk.State) int {
	dirname := s.CheckString(1)
	files, err := s.FS.ReadDir(dirname)
	if err != nil {
		if os.IsNotExist(err) {
			s.L.Push(lua.LNil)
			return 1
		}
		return s.RaiseError("%s", err)
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

func fsMkDir(s rbxmk.State) int {
	path := string(s.Pull(1, "string").(types.String))
	all := bool(s.PullOpt(1, "bool", types.Bool(false)).(types.Bool))
	var err error
	if all {
		err = s.FS.MkdirAll(path, 0755)
	} else {
		err = s.FS.Mkdir(path, 0755)
	}
	if err != nil {
		if os.IsExist(err) {
			s.L.Push(lua.LFalse)
			return 1
		}
		return s.RaiseError("%s", err)
	}
	s.L.Push(lua.LTrue)
	return 1
}

func fsRead(s rbxmk.State) int {
	fileName := string(s.Pull(1, "string").(types.String))
	selector := s.PullOpt(2, "FormatSelector", rtypes.FormatSelector{}).(rtypes.FormatSelector)
	if selector.Format == "" {
		selector.Format = s.Ext(fileName)
		if selector.Format == "" {
			return s.RaiseError("unknown format from %s", filepath.Base(fileName))
		}
	}

	format := s.Format(selector.Format)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", selector.Format)
	}
	if format.Decode == nil {
		return s.RaiseError("cannot decode with format %s", format.Name)
	}

	r, err := s.FS.Open(fileName)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	defer r.Close()
	v, err := format.Decode(selector, r)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	if inst, ok := v.(*rtypes.Instance); ok {
		ext := s.Ext(fileName)
		if ext != "" && ext != "." {
			ext = "." + ext
		}
		stem := filepath.Base(fileName)
		stem = stem[:len(stem)-len(ext)]
		inst.SetName(stem)
	}
	return s.Push(v)
}

func fsRemove(s rbxmk.State) int {
	path := string(s.Pull(1, "string").(types.String))
	all := bool(s.PullOpt(1, "bool", types.Bool(false)).(types.Bool))
	var err error
	if all {
		// RemoveAll returns nil if file does not exist.
		if _, err = s.FS.Stat(path); err == nil {
			err = s.FS.RemoveAll(path)
		}
	} else {
		err = s.FS.Remove(path)
	}
	if err != nil {
		if os.IsNotExist(err) {
			s.L.Push(lua.LFalse)
			return 1
		}
		return s.RaiseError("%s", err)
	}
	s.L.Push(lua.LTrue)
	return 1
}

func fsRename(s rbxmk.State) int {
	from := string(s.Pull(1, "string").(types.String))
	to := string(s.Pull(1, "string").(types.String))
	if _, err := s.FS.Stat(from); err != nil {
		if os.IsNotExist(err) {
			s.L.Push(lua.LFalse)
			return 1
		}
		return s.RaiseError("%s", err)
	}
	if err := s.FS.Rename(from, to); err != nil {
		return s.RaiseError("%s", err)
	}
	s.L.Push(lua.LTrue)
	return 1
}

func fsStat(s rbxmk.State) int {
	filename := s.CheckString(1)
	info, err := s.FS.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			s.L.Push(lua.LNil)
			return 1
		}
		return s.RaiseError("%s", err)
	}
	tinfo := s.L.CreateTable(0, 4)
	tinfo.RawSetString("Name", lua.LString(info.Name()))
	tinfo.RawSetString("IsDir", lua.LBool(info.IsDir()))
	tinfo.RawSetString("Size", lua.LNumber(info.Size()))
	tinfo.RawSetString("ModTime", lua.LNumber(info.ModTime().Unix()))
	s.L.Push(tinfo)
	return 1
}

func fsWrite(s rbxmk.State) int {
	fileName := string(s.Pull(1, "string").(types.String))
	value := s.Pull(2, "Variant")
	selector := s.PullOpt(3, "FormatSelector", rtypes.FormatSelector{}).(rtypes.FormatSelector)
	if selector.Format == "" {
		selector.Format = s.Ext(fileName)
		if selector.Format == "" {
			return s.RaiseError("unknown format from %s", filepath.Base(fileName))
		}
	}

	format := s.Format(selector.Format)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", selector.Format)
	}
	if format.Encode == nil {
		return s.RaiseError("cannot encode with format %s", format.Name)
	}

	w, err := s.FS.Create(fileName)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	defer w.Close()
	if err := format.Encode(selector, w, value); err != nil {
		return s.RaiseError("%s", err)
	}
	if err := w.Sync(); err != nil {
		return s.RaiseError("%s", err)
	}
	return 0
}
