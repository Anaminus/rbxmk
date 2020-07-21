package types

import (
	"github.com/anaminus/rbxmk"
)

// Stringlike implements rbxmk.Stringlike for a number of types.
type Stringlike struct {
	Value interface{}
}

func (s Stringlike) Stringlike() (v []byte, ok bool) {
	switch v := s.Value.(type) {
	case []byte:
		return v, true
	case string:
		return []byte(v), true
	case []rune:
		return []byte(string(v)), true
	case rbxmk.TValue:
		return Stringlike{Value: v.Value}.Stringlike()
	case *Instance:
		switch v.ClassName {
		case "Script", "LocalScript", "ModuleScript":
			return Stringlike{Value: v.Get("Source").Value}.Stringlike()
		}
	}
	return nil, false
}
