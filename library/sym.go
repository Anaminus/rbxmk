package library

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
)

func init() { register(Sym, 10) }

var Sym = rbxmk.Library{Name: "sym", Open: openSym}

func openSym(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 7)
	lib.RawSetString("AttrConfig", s.UserDataOf(rtypes.Symbol{Name: "AttrConfig"}, "Symbol"))
	lib.RawSetString("Desc", s.UserDataOf(rtypes.Symbol{Name: "Desc"}, "Symbol"))
	lib.RawSetString("IsService", s.UserDataOf(rtypes.Symbol{Name: "IsService"}, "Symbol"))
	lib.RawSetString("Metadata", s.UserDataOf(rtypes.Symbol{Name: "Metadata"}, "Symbol"))
	lib.RawSetString("RawAttrConfig", s.UserDataOf(rtypes.Symbol{Name: "RawAttrConfig"}, "Symbol"))
	lib.RawSetString("RawDesc", s.UserDataOf(rtypes.Symbol{Name: "RawDesc"}, "Symbol"))
	lib.RawSetString("Reference", s.UserDataOf(rtypes.Symbol{Name: "Reference"}, "Symbol"))
	return lib
}
