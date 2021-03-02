package sources

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(FSSource) }
func FSSource() rbxmk.Source {
	return rbxmk.Source{
		Name: "fs",
		Library: rbxmk.Library{
			Open: func(s rbxmk.State) *lua.LTable {
				lib := s.L.CreateTable(0, 6)
				lib.RawSetString("dir", s.WrapFunc(func(s rbxmk.State) int {
					dirname := s.CheckString(1)
					files, err := FS{World: s.World}.Dir(dirname)
					if err != nil {
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

				}))
				lib.RawSetString("mkdir", s.WrapFunc(func(s rbxmk.State) int {
					path := string(s.Pull(1, "string").(types.String))
					all := bool(s.PullOpt(1, "bool", types.Bool(false)).(types.Bool))
					ok, err := FS{World: s.World}.MkDir(path, all)
					if err != nil {
						return s.RaiseError("%s", err)
					}
					if ok {
						s.L.Push(lua.LTrue)
					} else {
						s.L.Push(lua.LFalse)
					}
					return 1
				}))
				lib.RawSetString("read", s.WrapFunc(func(s rbxmk.State) int {
					filename := string(s.Pull(1, "string").(types.String))
					selector := s.PullOpt(2, "FormatSelector", rtypes.FormatSelector{}).(rtypes.FormatSelector)
					v, err := FS{World: s.World}.Read(filename, selector)
					if err != nil {
						return s.RaiseError("%s", err)
					}
					return s.Push(v)
				}))
				lib.RawSetString("remove", s.WrapFunc(func(s rbxmk.State) int {
					path := string(s.Pull(1, "string").(types.String))
					all := bool(s.PullOpt(1, "bool", types.Bool(false)).(types.Bool))
					ok, err := FS{World: s.World}.Remove(path, all)
					if err != nil {
						return s.RaiseError("%s", err)
					}
					if ok {
						s.L.Push(lua.LTrue)
					} else {
						s.L.Push(lua.LFalse)
					}
					return 1
				}))
				lib.RawSetString("rename", s.WrapFunc(func(s rbxmk.State) int {
					from := string(s.Pull(1, "string").(types.String))
					to := string(s.Pull(1, "string").(types.String))
					ok, err := FS{World: s.World}.Rename(from, to)
					if err != nil {
						return s.RaiseError("%s", err)
					}
					if ok {
						s.L.Push(lua.LTrue)
					} else {
						s.L.Push(lua.LFalse)
					}
					return 1
				}))
				lib.RawSetString("stat", s.WrapFunc(func(s rbxmk.State) int {
					filename := s.CheckString(1)
					info, err := FS{World: s.World}.Stat(filename)
					if err != nil {
						return s.RaiseError("%s", err)
					}
					if info == nil {
						s.L.Push(lua.LNil)
						return 1
					}
					tinfo := s.L.CreateTable(0, 4)
					tinfo.RawSetString("Name", lua.LString(info.Name()))
					tinfo.RawSetString("IsDir", lua.LBool(info.IsDir()))
					tinfo.RawSetString("Size", lua.LNumber(info.Size()))
					tinfo.RawSetString("ModTime", lua.LNumber(info.ModTime().Unix()))
					s.L.Push(tinfo)
					return 1
				}))
				lib.RawSetString("write", s.WrapFunc(func(s rbxmk.State) int {
					filename := string(s.Pull(1, "string").(types.String))
					value := s.Pull(2, "Variant")
					selector := s.PullOpt(3, "FormatSelector", rtypes.FormatSelector{}).(rtypes.FormatSelector)
					err := FS{World: s.World}.Write(filename, value, selector)
					if err != nil {
						return s.RaiseError("%s", err)
					}
					return 0
				}))
				return lib
			},
		},
	}
}

// FS provides access to the file system.
type FS struct {
	*rbxmk.World
}

// Dir returns a list of files in the given directory.
func (s FS) Dir(dirname string) (files []fs.FileInfo, err error) {
	entries, err := s.FS.ReadDir(dirname)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	files = make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		files = append(files, info)
	}
	return files, nil
}

// MkDir creates a new directory.
func (s FS) MkDir(path string, all bool) (ok bool, err error) {
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
func (s FS) Read(filename string, selector rtypes.FormatSelector) (v types.Value, err error) {
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
func (s FS) Remove(path string, all bool) (ok bool, err error) {
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
func (s FS) Rename(from, to string) (ok bool, err error) {
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
func (s FS) Stat(filename string) (info fs.FileInfo, err error) {
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
func (s FS) Write(filename string, value types.Value, selector rtypes.FormatSelector) error {
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
