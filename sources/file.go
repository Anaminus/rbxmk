package sources

import (
	"os"
	"path/filepath"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(File) }
func File() rbxmk.Source {
	return rbxmk.Source{
		Name: "file",
		Library: rbxmk.Library{
			Open: func(s rbxmk.State) *lua.LTable {
				lib := s.L.CreateTable(0, 2)
				lib.RawSetString("read", s.WrapFunc(fileRead))
				lib.RawSetString("write", s.WrapFunc(fileWrite))
				return lib
			},
		},
	}
}

func fileRead(s rbxmk.State) int {
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

	r, err := os.Open(fileName)
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

func fileWrite(s rbxmk.State) int {
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

	w, err := os.Create(fileName)
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
