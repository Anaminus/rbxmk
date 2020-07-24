package library

import (
	"path/filepath"

	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func File(s rbxmk.State) {
	lib := s.L.CreateTable(0, 1)
	lib.RawSetString("read", s.WrapFunc(fileRead))
	lib.RawSetString("write", s.WrapFunc(fileWrite))
	s.L.SetGlobal("file", lib)
}

func fileRead(s rbxmk.State) int {
	fileName := string(s.Pull(1, "string").(types.String))
	formatName := string(s.PullOpt(2, "string", types.String("")).(types.String))
	if formatName == "" {
		f := guessExt(s, fileName)
		if f == "" {
			s.L.RaiseError("unknown format from %s", filepath.Base(fileName))
			return 0
		}
		formatName = f
	}

	format := s.Format(formatName)
	if format.Name == "" {
		s.L.RaiseError("unknown format %q", formatName)
		return 0
	}
	if format.Decode == nil {
		s.L.RaiseError("cannot decode with format %s", format.Name)
		return 0
	}

	b, err := s.Source("file").Read(fileName)
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	v, err := format.Decode(b)
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	return s.Push("Variant", v)
}

func fileWrite(s rbxmk.State) int {
	var fileName string
	var formatName string
	var value rbxmk.Value
	switch s.L.GetTop() {
	case 2:
		fileName = string(s.Pull(1, "string").(types.String))
		value = s.Pull(2, "Variant")
		f := guessExt(s, fileName)
		if f == "" {
			s.L.RaiseError("unknown format from %s", filepath.Base(fileName))
			return 0
		}
		formatName = f
	case 3:
		fileName = string(s.Pull(1, "string").(types.String))
		formatName = string(s.Pull(2, "string").(types.String))
		value = s.Pull(3, "Variant")
	default:
		s.L.RaiseError("expected 2 or 3 arguments")
		return 0
	}
	format := s.Format(formatName)
	if format.Name == "" {
		s.L.RaiseError("unknown format %q", formatName)
		return 0
	}
	if format.Encode == nil {
		s.L.RaiseError("cannot encode with format %s", format.Name)
		return 0
	}

	b, err := format.Encode(value)
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	if err := s.Source("file").Write(b, fileName); err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	return 0
}
