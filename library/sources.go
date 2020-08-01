package library

import (
	"github.com/anaminus/rbxmk"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(Sources) }

var Sources = rbxmk.Library{
	Name: "",
	Open: func(s rbxmk.State) *lua.LTable {
		lib := s.L.CreateTable(0, 1)
		for _, source := range s.Sources() {
			if source.Library.Open != nil {
				name := source.Library.Name
				if name == "" {
					name = source.Name
				}
				src := source.Library.Open(s)
				if err := s.MergeTables(lib, src, name); err != nil {
					panic(err.Error())
				}
			}
		}
		return lib
	},
}
