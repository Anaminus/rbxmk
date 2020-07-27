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
		Name:     "Symbol",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				return s.Push(types.String(s.Pull(1, "Symbol").(SymbolType).String()))
			},
			"__eq": func(s State) int {
				op := s.Pull(2, "Symbol").(SymbolType)
				return s.Push(types.Bool(s.Pull(1, "Symbol").(SymbolType) == op))
			},
		},
		Environment: func(s State) {
			typ := s.Type("Symbol")

			v, _ := typ.PushTo(s, typ, SymbolReference)
			s.L.SetGlobal("Reference", v[0])

			v, _ = typ.PushTo(s, typ, SymbolIsService)
			s.L.SetGlobal("IsService", v[0])
		},
	}
}
