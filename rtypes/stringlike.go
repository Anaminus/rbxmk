package rtypes

import "github.com/anaminus/rbxmk"

// Stringlike implements rbxmk.Stringlike for a number of types.
type Stringlike struct {
	Value interface{}
}

func (s Stringlike) IsStringlike() bool {
	switch v := s.Value.(type) {
	case string, []byte, []rune, rbxmk.Stringlike:
		return true
	case *Instance:
		switch v.ClassName {
		case "Script", "LocalScript", "ModuleScript":
			return Stringlike{Value: v.Get("Source")}.IsStringlike()
		}
	}
	return false
}

func (s Stringlike) Stringlike() string {
	switch v := s.Value.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case []rune:
		return string(v)
	case rbxmk.Stringlike:
		return v.Stringlike()
	case *Instance:
		switch v.ClassName {
		case "Script", "LocalScript", "ModuleScript":
			return Stringlike{Value: v.Get("Source")}.Stringlike()
		}
	}
	return ""
}
