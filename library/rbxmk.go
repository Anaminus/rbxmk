package library

import (
	"os"
	"path/filepath"

	"github.com/anaminus/rbxmk"
	"github.com/yuin/gopher-lua"
)

func RBXMK(s rbxmk.State) {
	lib := s.L.CreateTable(0, 1)
	lib.RawSetString("load", s.WrapFunc(rbxmkLoad))
	lib.RawSetString("encodeformat", s.WrapFunc(rbxmkEncodeFormat))
	lib.RawSetString("decodeformat", s.WrapFunc(rbxmkDecodeFormat))
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
	name := s.Pull(1, "string").(string)
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
	return s.Push("BinaryString", b)
}

func rbxmkDecodeFormat(s rbxmk.State) int {
	name := s.Pull(1, "string").(string)
	format := s.Format(name)
	if format.Name == "" {
		s.L.RaiseError("unknown format %q", name)
		return 0
	}
	if format.Decode == nil {
		s.L.RaiseError("cannot decode with format %s", name)
		return 0
	}
	v, err := format.Decode(s.Pull(2, "BinaryString").([]byte))
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	return s.Push("Variant", v)
}
