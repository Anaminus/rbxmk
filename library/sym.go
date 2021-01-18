package library

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(Sym, 10) }

var Sym = rbxmk.Library{
	Name: "sym",
	Open: func(s rbxmk.State) *lua.LTable {
		lib := s.L.CreateTable(0, 4)
		lib.RawSetString("Reference", s.UserDataOf(rtypes.Symbol{Name: "Reference"}, "Symbol"))
		lib.RawSetString("IsService", s.UserDataOf(rtypes.Symbol{Name: "IsService"}, "Symbol"))
		lib.RawSetString("Desc", s.UserDataOf(rtypes.Symbol{Name: "Desc"}, "Symbol"))
		lib.RawSetString("RawDesc", s.UserDataOf(rtypes.Symbol{Name: "RawDesc"}, "Symbol"))
		lib.RawSetString("AttrConfig", s.UserDataOf(rtypes.Symbol{Name: "AttrConfig"}, "Symbol"))
		lib.RawSetString("RawAttrConfig", s.UserDataOf(rtypes.Symbol{Name: "RawAttrConfig"}, "Symbol"))
		return lib
	},
}
