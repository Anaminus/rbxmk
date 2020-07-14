package reflect

import (
	. "github.com/anaminus/rbxmk"
)

type SymbolType struct {
	Name string
}

var SymbolReference = SymbolType{Name: "Reference"}
var SymbolIsService = SymbolType{Name: "IsService"}

func Symbol() Type {
	return Type{
		Name:        "Symbol",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				return s.Push("string", "Symbol<"+v.(SymbolType).Name+">")
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "Symbol").(SymbolType)
				return s.Push("bool", v.(SymbolType) == op)
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
