package library

import (
	"github.com/anaminus/rbxmk"
	"github.com/yuin/gopher-lua"
)

var Types = rbxmk.Library{
	Name: "types",
	Open: func(s rbxmk.State) *lua.LTable {
		lib := s.L.CreateTable(0, 8)
		lib.RawSetString("int", s.WrapFunc(typesInt))
		lib.RawSetString("int64", s.WrapFunc(typesInt64))
		lib.RawSetString("float", s.WrapFunc(typesFloat))
		lib.RawSetString("token", s.WrapFunc(typesToken))
		lib.RawSetString("BinaryString", s.WrapFunc(typesBinaryString))
		lib.RawSetString("ProtectedString", s.WrapFunc(typesProtectedString))
		lib.RawSetString("Content", s.WrapFunc(typesContent))
		lib.RawSetString("SharedString", s.WrapFunc(typesSharedString))
		return lib
	},
}

func setUserdata(s rbxmk.State, t string) int {
	u := s.L.NewUserData()
	u.Value = s.Pull(1, t)
	s.L.SetMetatable(u, s.L.GetTypeMetatable(t))
	s.L.Push(u)
	return 1
}

func typesInt(s rbxmk.State) int {
	return setUserdata(s, "int")
}

func typesInt64(s rbxmk.State) int {
	return setUserdata(s, "int64")
}

func typesFloat(s rbxmk.State) int {
	return setUserdata(s, "float")
}

func typesToken(s rbxmk.State) int {
	return setUserdata(s, "token")
}

func typesBinaryString(s rbxmk.State) int {
	return setUserdata(s, "BinaryString")
}

func typesProtectedString(s rbxmk.State) int {
	return setUserdata(s, "ProtectedString")
}

func typesContent(s rbxmk.State) int {
	return setUserdata(s, "Content")
}

func typesSharedString(s rbxmk.State) int {
	return setUserdata(s, "SharedString")
}
