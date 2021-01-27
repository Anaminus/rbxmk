package sources

import (
	"io/ioutil"
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
		Read: func(s rbxmk.State) (b []byte, err error) {
			path := string(s.Pull(1, "string").(types.String))
			return ioutil.ReadFile(path)
		},
		Write: func(s rbxmk.State, b []byte) (err error) {
			path := string(s.Pull(1, "string").(types.String))
			return ioutil.WriteFile(path, b, 0666)
		},
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
	formatName := string(s.PullOpt(2, "string", types.String("")).(types.String))
	if formatName == "" {
		if formatName = s.Ext(fileName); formatName == "" {
			return s.RaiseError("unknown format from %s", filepath.Base(fileName))
		}
	}

	format := s.Format(formatName)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", formatName)
	}
	if format.Decode == nil {
		return s.RaiseError("cannot decode with format %s", format.Name)
	}

	r, err := os.Open(fileName)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	defer r.Close()
	v, err := format.Decode(rbxmk.FormatOptions{}, r)
	if err != nil {
		return s.RaiseError(err.Error())
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
	formatName := string(s.PullOpt(3, "string", types.String("")).(types.String))
	if formatName == "" {
		if formatName = s.Ext(fileName); formatName == "" {
			return s.RaiseError("unknown format from %s", filepath.Base(fileName))
		}
	}

	format := s.Format(formatName)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", formatName)
	}
	if format.Encode == nil {
		return s.RaiseError("cannot encode with format %s", format.Name)
	}

	w, err := os.Create(fileName)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	defer w.Close()
	if err := format.Encode(rbxmk.FormatOptions{}, w, value); err != nil {
		return s.RaiseError(err.Error())
	}
	if err := w.Sync(); err != nil {
		return s.RaiseError(err.Error())
	}
	return 0
}
