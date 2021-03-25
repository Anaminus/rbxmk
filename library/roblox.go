package library

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	reflect "github.com/anaminus/rbxmk/library/roblox"
)

func init() { register(Roblox, 1) }

var Roblox = rbxmk.Library{Name: "", Open: openRoblox, Dump: dumpRoblox}

func openRoblox(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 1)
	lib.RawSetString("typeof", s.WrapFunc(robloxTypeof))

	for _, f := range reflect.All() {
		r := f()
		s.RegisterReflector(r)
		s.ApplyReflector(r, lib)
	}

	return lib
}

func robloxTypeof(s rbxmk.State) int {
	v := s.CheckAny(1)
	t := s.Typeof(v)
	s.L.Push(lua.LString(t))
	return 1
}

func dumpRoblox(s rbxmk.State) dump.Library {
	lib := dump.Library{
		Name: "roblox",
		Struct: dump.Struct{
			Fields: dump.Fields{
				"typeof": dump.Function{
					Parameters: dump.Parameters{
						{Name: "value", Type: dt.Prim("any")},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("string")},
					},
					Summary:     "libraries/roblox:Fields/typeof/Summary",
					Description: "libraries/roblox:Fields/typeof/Description",
				},
			},
			Summary:     "libraries/roblox:Summary",
			Description: "libraries/roblox:Description",
		},
		Types: dump.TypeDefs{
			"DataModel": dump.TypeDef{
				Underlying: dt.Prim("Instance"),
				Symbols: dump.Properties{
					"Metadata": dump.Property{
						ValueType:   dt.Dictionary{V: dt.Prim("string")},
						Summary:     "libraries/roblox/types/DataModel:Symbols/Metadata/Summary",
						Description: "libraries/roblox/types/DataModel:Symbols/Metadata/Description",
					},
				},
				Methods: dump.Methods{
					"GetService": dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Instance")},
						},
						Summary:     "libraries/roblox/types/DataModel:Methods/GetService/Summary",
						Description: "libraries/roblox/types/DataModel:Methods/GetService/Description",
					},
				},
				Constructors: dump.Constructors{
					"new": dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "descriptor", Type: dt.Optional{T: dt.Group{T: dt.Or{dt.Prim("RootDesc"), dt.Prim("bool")}}}},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("DataModel")},
							},
							Summary:     "libraries/roblox/types/DataModel:Constructors/new/Summary",
							Description: "libraries/roblox/types/DataModel:Constructors/new/Description",
						},
					},
				},
				Summary:     "libraries/roblox/types/DataModel:Summary",
				Description: "libraries/roblox/types/DataModel:Description",
			},
		},
	}
	for _, f := range reflect.All() {
		r := f()
		lib.Types[r.Name] = r.DumpAll()
	}
	return lib
}
