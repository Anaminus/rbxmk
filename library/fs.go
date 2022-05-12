package library

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(FS) }

var FS = rbxmk.Library{
	Name:       "fs",
	ImportedAs: "fs",
	Priority:   10,
	Open:       openFS,
	Dump:       dumpFS,
	Types: []func() rbxmk.Reflector{
		reflect.Bool,
		reflect.DirEntry,
		reflect.FileInfo,
		reflect.FormatSelector,
		reflect.String,
		reflect.Variant,
	},
}

func openFS(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 6)
	lib.RawSetString("dir", s.WrapFunc(fsDir))
	lib.RawSetString("mkdir", s.WrapFunc(fsMkdir))
	lib.RawSetString("read", s.WrapFunc(fsRead))
	lib.RawSetString("remove", s.WrapFunc(fsRemove))
	lib.RawSetString("rename", s.WrapFunc(fsRename))
	lib.RawSetString("stat", s.WrapFunc(fsStat))
	lib.RawSetString("write", s.WrapFunc(fsWrite))
	return lib
}

func fsDir(s rbxmk.State) int {
	dirname := string(s.Pull(1, rtypes.T_String).(types.String))
	files, err := FSSource{World: s.World}.Dir(dirname)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	tfiles := s.L.CreateTable(len(files), 0)
	for i, entry := range files {
		s.PushToArray(tfiles, i+1, rtypes.DirEntry{DirEntry: entry})
	}
	s.L.Push(tfiles)
	return 1

}

func fsMkdir(s rbxmk.State) int {
	path := string(s.Pull(1, rtypes.T_String).(types.String))
	all := bool(s.PullOpt(2, types.Bool(false), rtypes.T_Bool).(types.Bool))
	ok, err := FSSource{World: s.World}.MkDir(path, all)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	if ok {
		s.L.Push(lua.LTrue)
	} else {
		s.L.Push(lua.LFalse)
	}
	return 1
}

func fsRead(s rbxmk.State) int {
	filename := string(s.Pull(1, rtypes.T_String).(types.String))
	selector := s.PullOpt(2, rtypes.FormatSelector{}, rtypes.T_FormatSelector).(rtypes.FormatSelector)
	v, err := FSSource{World: s.World}.Read(filename, selector)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return s.Push(v)
}

func fsRemove(s rbxmk.State) int {
	path := string(s.Pull(1, rtypes.T_String).(types.String))
	all := bool(s.PullOpt(2, types.Bool(false), rtypes.T_Bool).(types.Bool))
	ok, err := FSSource{World: s.World}.Remove(path, all)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	if ok {
		s.L.Push(lua.LTrue)
	} else {
		s.L.Push(lua.LFalse)
	}
	return 1
}

func fsRename(s rbxmk.State) int {
	from := string(s.Pull(1, rtypes.T_String).(types.String))
	to := string(s.Pull(2, rtypes.T_String).(types.String))
	ok, err := FSSource{World: s.World}.Rename(from, to)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	if ok {
		s.L.Push(lua.LTrue)
	} else {
		s.L.Push(lua.LFalse)
	}
	return 1
}

func fsStat(s rbxmk.State) int {
	filename := string(s.Pull(1, rtypes.T_String).(types.String))
	info, err := FSSource{World: s.World}.Stat(filename)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	if info == nil {
		s.L.Push(lua.LNil)
		return 1
	}
	return s.Push(rtypes.FileInfo{FileInfo: info})
}

func fsWrite(s rbxmk.State) int {
	filename := string(s.Pull(1, rtypes.T_String).(types.String))
	selector := s.PullOpt(3, rtypes.FormatSelector{}, rtypes.T_FormatSelector).(rtypes.FormatSelector)
	if selector.Format == "" {
		selector.Format = s.Ext(filename)
	}
	value := s.PullEncoded(2, selector)
	err := FSSource{World: s.World}.Write(filename, value, selector)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return 0
}

func dumpFS(s rbxmk.State) dump.Library {
	return dump.Library{
		Struct: dump.Struct{
			Fields: dump.Fields{
				"dir": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim(rtypes.T_String)},
					},
					Returns: dump.Parameters{
						{Type: dt.Optional{T: dt.Array{T: dt.Prim(rtypes.T_DirEntry)}}},
					},
					CanError:    true,
					Summary:     "Libraries/fs:Fields/dir/Summary",
					Description: "Libraries/fs:Fields/dir/Description",
				},
				"mkdir": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim(rtypes.T_String)},
						{Name: "all", Type: dt.Optional{T: dt.Prim(rtypes.T_Bool)}},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_LuaBoolean)},
					},
					CanError:    true,
					Summary:     "Libraries/fs:Fields/mkdir/Summary",
					Description: "Libraries/fs:Fields/mkdir/Description",
				},
				"read": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim(rtypes.T_String)},
						{Name: "format", Type: dt.Optional{T: dt.Prim(rtypes.T_FormatSelector)}},
					},
					Returns: dump.Parameters{
						{Name: "value", Type: dt.Prim("any")},
					},
					CanError:    true,
					Summary:     "Libraries/fs:Fields/read/Summary",
					Description: "Libraries/fs:Fields/read/Description",
				},
				"remove": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim(rtypes.T_String)},
						{Name: "all", Type: dt.Optional{T: dt.Prim(rtypes.T_Bool)}},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_LuaBoolean)},
					},
					CanError:    true,
					Summary:     "Libraries/fs:Fields/remove/Summary",
					Description: "Libraries/fs:Fields/remove/Description",
				},
				"rename": dump.Function{
					Parameters: dump.Parameters{
						{Name: "old", Type: dt.Prim(rtypes.T_String)},
						{Name: "new", Type: dt.Prim(rtypes.T_String)},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_LuaBoolean)},
					},
					CanError:    true,
					Summary:     "Libraries/fs:Fields/rename/Summary",
					Description: "Libraries/fs:Fields/rename/Description",
				},
				"stat": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim(rtypes.T_String)},
					},
					Returns: dump.Parameters{
						{Type: dt.Optional{T: dt.Prim(rtypes.T_FileInfo)}},
					},
					CanError:    true,
					Summary:     "Libraries/fs:Fields/stat/Summary",
					Description: "Libraries/fs:Fields/stat/Description",
				},
				"write": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim(rtypes.T_String)},
						{Name: "value", Type: dt.Prim("any")},
						{Name: "format", Type: dt.Optional{T: dt.Prim(rtypes.T_FormatSelector)}},
					},
					CanError:    true,
					Summary:     "Libraries/fs:Fields/write/Summary",
					Description: "Libraries/fs:Fields/write/Description",
				},
			},
			Summary:     "Libraries/fs:Summary",
			Description: "Libraries/fs:Description",
		},
	}
}

// FSSource provides access to the file system.
type FSSource struct {
	*rbxmk.World
}

// Dir returns a list of files in the given directory.
func (s FSSource) Dir(dirname string) (files []fs.DirEntry, err error) {
	files, err = s.FS.ReadDir(dirname)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return files, nil
}

// MkDir creates a new directory.
func (s FSSource) MkDir(path string, all bool) (ok bool, err error) {
	if all {
		err = s.FS.MkdirAll(path, 0755)
	} else {
		err = s.FS.Mkdir(path, 0755)
	}
	if err != nil {
		if os.IsExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Read reads the content of a file.
func (s FSSource) Read(filename string, selector rtypes.FormatSelector) (v types.Value, err error) {
	if selector.Format == "" {
		selector.Format = s.Ext(filename)
		if selector.Format == "" {
			return nil, fmt.Errorf("unknown format from %s", filepath.Base(filename))
		}
	}

	format := s.Format(selector.Format)
	if format.Name == "" {
		return nil, fmt.Errorf("unknown format %q", selector.Format)
	}
	if format.Decode == nil {
		return nil, fmt.Errorf("cannot decode with format %s", format.Name)
	}

	r, err := s.FS.Open(filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	v, err = format.Decode(s.Global, selector, r)
	if err != nil {
		return nil, err
	}
	if inst, ok := v.(*rtypes.Instance); ok {
		ext := s.Ext(filename)
		if ext != "" && ext != "." {
			ext = "." + ext
		}
		stem := filepath.Base(filename)
		stem = stem[:len(stem)-len(ext)]
		inst.SetName(stem)
	}
	return v, nil
}

// Remove removes a file or directory.
func (s FSSource) Remove(path string, all bool) (ok bool, err error) {
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
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Rename moves a file or directory.
func (s FSSource) Rename(from, to string) (ok bool, err error) {
	if _, err := s.FS.Stat(from); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if err := s.FS.Rename(from, to); err != nil {
		return false, err
	}
	return true, nil
}

// Stat gets metadata of the given file.
func (s FSSource) Stat(filename string) (info fs.FileInfo, err error) {
	info, err = s.FS.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return info, nil
}

// Write writes a value to a file.
func (s FSSource) Write(filename string, value types.Value, selector rtypes.FormatSelector) error {
	if selector.Format == "" {
		selector.Format = s.Ext(filename)
		if selector.Format == "" {
			return fmt.Errorf("unknown format from %s", filepath.Base(filename))
		}
	}

	format := s.Format(selector.Format)
	if format.Name == "" {
		return fmt.Errorf("unknown format %q", selector.Format)
	}
	if format.Encode == nil {
		return fmt.Errorf("cannot encode with format %s", format.Name)
	}

	w, err := s.FS.Create(filename)
	if err != nil {
		return err
	}
	defer w.Close()
	if err := format.Encode(s.Global, selector, w, value); err != nil {
		return err
	}
	if err := w.Sync(); err != nil {
		return err
	}
	return nil
}
