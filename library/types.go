package library

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
)

func setUserdata(s rbxmk.State, t string) int {
	v := s.Pull(1, t)
	u := s.L.NewUserData(v)
	s.L.SetMetatable(u, s.L.GetTypeMetatable(t))
	s.L.Push(u)
	return 1
}

func init() { register(Types) }

var Types = rbxmk.Library{
	Name:       "types",
	ImportedAs: "types",
	Priority:   10,
	Open:       openTypes,
	Dump:       dumpTypes,
}

func openTypes(s rbxmk.State) *lua.LTable {
	exprims := s.Reflectors(rbxmk.Exprim)
	lib := s.L.CreateTable(0, len(exprims))
	for _, t := range exprims {
		name := t.Name
		if name == "Optional" {
			lib.RawSetString("some", s.WrapFunc(func(s rbxmk.State) int {
				return setUserdata(s, name)
			}))
			lib.RawSetString("none", s.WrapFunc(func(s rbxmk.State) int {
				v := s.CheckString(1)
				u := s.L.NewUserData(rtypes.None(v))
				s.L.SetMetatable(u, s.L.GetTypeMetatable(name))
				s.L.Push(u)
				return 1
			}))
			continue
		}
		lib.RawSetString(t.Name, s.WrapFunc(func(s rbxmk.State) int {
			return setUserdata(s, name)
		}))
	}
	return lib
}

func dumpTypes(s rbxmk.State) dump.Library {
	exprims := s.Reflectors(rbxmk.Exprim)
	lib := dump.Library{
		Struct: dump.Struct{
			Fields:      make(dump.Fields, len(exprims)),
			Summary:     "Libraries/types:Summary",
			Description: "Libraries/types:Description",
		},
	}
	for _, t := range exprims {
		typ := dt.Prim(t.Name)
		if t.Dump != nil {
			if d := t.Dump(); d.Underlying != nil {
				if prim, ok := d.Underlying.(dt.Prim); ok {
					typ = prim
				}
			}
		}
		if t.Name == "Optional" {
			lib.Struct.Fields["some"] = dump.Function{
				Parameters:  []dt.Parameter{{Name: "type", Type: dt.Prim("any")}},
				Returns:     []dt.Parameter{{Type: dt.Prim(t.Name)}},
				Summary:     "Libraries/types:Fields/" + t.Name + "/Summary",
				Description: "Libraries/types:Fields/" + t.Name + "/Description",
			}
			lib.Struct.Fields["none"] = dump.Function{
				Parameters:  []dt.Parameter{{Name: "type", Type: dt.Prim("string")}},
				Returns:     []dt.Parameter{{Type: dt.Prim(t.Name)}},
				Summary:     "Libraries/types:Fields/" + t.Name + "/Summary",
				Description: "Libraries/types:Fields/" + t.Name + "/Description",
			}
			continue
		}
		lib.Struct.Fields[t.Name] = dump.Function{
			Parameters:  []dt.Parameter{{Name: "value", Type: typ}},
			Returns:     []dt.Parameter{{Type: dt.Prim(t.Name)}},
			Summary:     "Libraries/types:Fields/" + t.Name + "/Summary",
			Description: "Libraries/types:Fields/" + t.Name + "/Description",
		}
	}
	return lib
}
