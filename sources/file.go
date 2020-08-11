package sources

import (
	"io/ioutil"
	"path/filepath"

	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	lua "github.com/yuin/gopher-lua"
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
		f := s.Ext(fileName)
		if f == "" {
			return s.RaiseError("unknown format from %s", filepath.Base(fileName))
		}
		formatName = f
	}

	format := s.Format(formatName)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", formatName)
	}
	if format.Decode == nil {
		return s.RaiseError("cannot decode with format %s", format.Name)
	}

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	v, err := format.Decode(b)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	return s.Push(v)
}

func fileWrite(s rbxmk.State) int {
	var fileName string
	var formatName string
	var value types.Value
	switch s.L.GetTop() {
	case 2:
		fileName = string(s.Pull(1, "string").(types.String))
		value = s.Pull(2, "Variant")
		f := s.Ext(fileName)
		if f == "" {
			return s.RaiseError("unknown format from %s", filepath.Base(fileName))
		}
		formatName = f
	case 3:
		fileName = string(s.Pull(1, "string").(types.String))
		formatName = string(s.Pull(2, "string").(types.String))
		value = s.Pull(3, "Variant")
	default:
		return s.RaiseError("expected 2 or 3 arguments")
	}
	format := s.Format(formatName)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", formatName)
	}
	if format.Encode == nil {
		return s.RaiseError("cannot encode with format %s", format.Name)
	}

	b, err := format.Encode(value)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	if err := ioutil.WriteFile(fileName, b, 0666); err != nil {
		return s.RaiseError(err.Error())
	}
	return 0
}
