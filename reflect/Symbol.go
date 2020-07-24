package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

type SymbolType struct {
	Name string
}

func (SymbolType) Type() string {
	return "Symbol"
}

func (s SymbolType) String() string {
	return "Symbol<" + s.Name + ">"
}

var SymbolReference = SymbolType{Name: "Reference"}
var SymbolIsService = SymbolType{Name: "IsService"}

func Symbol() Type {
	return Type{
		Name:        "Symbol",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				return s.Push("string", types.String(s.Pull(1, "Symbol").(SymbolType).String()))
			},
			"__eq": func(s State) int {
				op := s.Pull(2, "Symbol").(SymbolType)
				return s.Push("bool", types.Bool(s.Pull(1, "Symbol").(SymbolType) == op))
			},
		},
		Environment: func(s State) {
			typ := s.Type("Symbol")

			v, _ := typ.ReflectTo(s, typ, SymbolReference)
			s.L.SetGlobal("Reference", v[0])

			v, _ = typ.ReflectTo(s, typ, SymbolIsService)
			s.L.SetGlobal("IsService", v[0])
		},
	}
}
