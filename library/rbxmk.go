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
