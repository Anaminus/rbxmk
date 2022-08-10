package library

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
)

func init() { register(Sym) }

var Sym = rbxmk.Library{
	Name:     "sym",
	Import:   []string{"sym"},
	Priority: 10,
	Open:     openSym,
	Dump:     dumpSym,
	Types: []func() rbxmk.Reflector{
		reflect.Symbol,
	},
}

func openSym(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 7)
	lib.RawSetString(rtypes.T_AttrConfig, s.UserDataOf(rtypes.Symbol{Name: rtypes.T_AttrConfig}, rtypes.T_Symbol))
	lib.RawSetString(rtypes.T_Desc, s.UserDataOf(rtypes.Symbol{Name: rtypes.T_Desc}, rtypes.T_Symbol))
	lib.RawSetString("IsService", s.UserDataOf(rtypes.Symbol{Name: "IsService"}, rtypes.T_Symbol))
	lib.RawSetString("Metadata", s.UserDataOf(rtypes.Symbol{Name: "Metadata"}, rtypes.T_Symbol))
	lib.RawSetString("Properties", s.UserDataOf(rtypes.Symbol{Name: "Properties"}, rtypes.T_Symbol))
	lib.RawSetString("Raw"+rtypes.T_AttrConfig, s.UserDataOf(rtypes.Symbol{Name: "Raw" + rtypes.T_AttrConfig}, rtypes.T_Symbol))
	lib.RawSetString("Raw"+rtypes.T_Desc, s.UserDataOf(rtypes.Symbol{Name: "Raw" + rtypes.T_Desc}, rtypes.T_Symbol))
	lib.RawSetString("Reference", s.UserDataOf(rtypes.Symbol{Name: "Reference"}, rtypes.T_Symbol))
	return lib
}

func dumpSym(s rbxmk.State) dump.Library {
	return dump.Library{
		Struct: dump.Struct{
			Fields: dump.Fields{
				rtypes.T_AttrConfig: dump.Property{
					ValueType:   dt.Prim(rtypes.T_Symbol),
					ReadOnly:    true,
					Summary:     "Libraries/sym:Fields/AttrConfig/Summary",
					Description: "Libraries/sym:Fields/AttrConfig/Description",
				},
				rtypes.T_Desc: dump.Property{
					ValueType:   dt.Prim(rtypes.T_Symbol),
					ReadOnly:    true,
					Summary:     "Libraries/sym:Fields/Desc/Summary",
					Description: "Libraries/sym:Fields/Desc/Description",
				},
				"IsService": dump.Property{
					ValueType:   dt.Prim(rtypes.T_Symbol),
					ReadOnly:    true,
					Summary:     "Libraries/sym:Fields/IsService/Summary",
					Description: "Libraries/sym:Fields/IsService/Description",
				},
				"Metadata": dump.Property{
					ValueType:   dt.Prim(rtypes.T_Symbol),
					ReadOnly:    true,
					Summary:     "Libraries/sym:Fields/Metadata/Summary",
					Description: "Libraries/sym:Fields/Metadata/Description",
				},
				"Properties": dump.Property{
					ValueType:   dt.Prim(rtypes.T_Symbol),
					ReadOnly:    true,
					Summary:     "Libraries/sym:Fields/Properties/Summary",
					Description: "Libraries/sym:Fields/Properties/Description",
				},
				"Raw" + rtypes.T_AttrConfig: dump.Property{
					ValueType:   dt.Prim(rtypes.T_Symbol),
					ReadOnly:    true,
					Summary:     "Libraries/sym:Fields/RawAttrConfig/Summary",
					Description: "Libraries/sym:Fields/RawAttrConfig/Description",
				},
				"Raw" + rtypes.T_Desc: dump.Property{
					ValueType:   dt.Prim(rtypes.T_Symbol),
					ReadOnly:    true,
					Summary:     "Libraries/sym:Fields/RawDesc/Summary",
					Description: "Libraries/sym:Fields/RawDesc/Description",
				},
				"Reference": dump.Property{
					ValueType:   dt.Prim(rtypes.T_Symbol),
					ReadOnly:    true,
					Summary:     "Libraries/sym:Fields/Reference/Summary",
					Description: "Libraries/sym:Fields/Reference/Description",
				},
			},
			Summary:     "Libraries/sym:Summary",
			Description: "Libraries/sym:Description",
		},
	}
}
