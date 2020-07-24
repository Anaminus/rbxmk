package library

import (
	"os"
	"path/filepath"

	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func RBXMK(s rbxmk.State) {
	lib := s.L.CreateTable(0, 1)
	lib.RawSetString("load", s.WrapFunc(rbxmkLoad))
	lib.RawSetString("encodeformat", s.WrapFunc(rbxmkEncodeFormat))
	lib.RawSetString("decodeformat", s.WrapFunc(rbxmkDecodeFormat))
	lib.RawSetString("readsource", s.WrapFunc(rbxmkReadSource))
	lib.RawSetString("writesource", s.WrapFunc(rbxmkWriteSource))
	s.L.SetGlobal("rbxmk", lib)
}

func rbxmkLoad(s rbxmk.State) int {
	fileName := filepath.Clean(s.L.CheckString(1))
	fi, err := os.Stat(fileName)
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	if err = s.PushFile(rbxmk.FileInfo{Path: fileName, FileInfo: fi}); err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}

	// Load file as function.
	fn, err := s.L.LoadFile(fileName)
	if err != nil {
		s.PopFile()
		s.L.RaiseError(err.Error())
		return 0
	}
	s.L.Push(fn) // +function

	// Push extra arguments as arguments to loaded function.
	nt := s.L.GetTop()
	for i := 2; i <= nt; i++ {
		s.L.Push(s.L.Get(i)) // function, ..., +arg
	}
	// function, +args...

	// Call loaded function.
	err = s.L.PCall(nt-1, lua.MultRet, nil) // -function, -args..., +returns...
	s.PopFile()
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	return s.L.GetTop() - 1
}

func rbxmkEncodeFormat(s rbxmk.State) int {
	name := string(s.Pull(1, "string").(types.String))
	format := s.Format(name)
	if format.Name == "" {
		s.L.RaiseError("unknown format %q", name)
		return 0
	}
	if format.Encode == nil {
		s.L.RaiseError("cannot encode with format %s", name)
		return 0
	}
	b, err := format.Encode(s.Pull(2, "Variant"))
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	return s.Push(types.BinaryString(b))
}

func rbxmkDecodeFormat(s rbxmk.State) int {
	name := string(s.Pull(1, "string").(types.String))
	format := s.Format(name)
	if format.Name == "" {
		s.L.RaiseError("unknown format %q", name)
		return 0
	}
	if format.Decode == nil {
		s.L.RaiseError("cannot decode with format %s", name)
		return 0
	}
	v, err := format.Decode([]byte(s.Pull(2, "BinaryString").(types.BinaryString)))
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	return s.Push(v)
}

func rbxmkReadSource(s rbxmk.State) int {
	name := string(s.Pull(1, "string").(types.String))
	source := s.Source(name)
	if source.Name == "" {
		s.L.RaiseError("unknown source %q", name)
		return 0
	}
	if source.Read == nil {
		s.L.RaiseError("cannot read with format %s", name)
		return 0
	}
	options := make([]interface{}, s.L.GetTop()-1)
	for i := 2; i <= s.L.GetTop(); i++ {
		options[i-2] = s.Pull(i, "Variant")
	}
	b, err := source.Read(options...)
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	return s.Push(types.BinaryString(b))
}

func rbxmkWriteSource(s rbxmk.State) int {
	name := string(s.Pull(1, "string").(types.String))
	source := s.Source(name)
	if source.Name == "" {
		s.L.RaiseError("unknown source %q", name)
		return 0
	}
	if source.Write == nil {
		s.L.RaiseError("cannot write with format %s", name)
		return 0
	}
	b := []byte(s.Pull(2, "BinaryString").(types.BinaryString))
	options := make([]interface{}, s.L.GetTop()-2)
	for i := 3; i <= s.L.GetTop(); i++ {
		options[i-3] = s.Pull(i, "Variant")
	}
	if err := source.Write(b, options...); err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	return 0
}
