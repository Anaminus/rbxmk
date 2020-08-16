package rtypes

import (
	"github.com/robloxapi/types"
)

// Stringlike implements types.Stringlike for a number of types.
type Stringlike struct {
	Value interface{}
}

// IsStringlike returns whether Value can be converted to a string.
func (s Stringlike) IsStringlike() bool {
	switch v := s.Value.(type) {
	case string, []byte, []rune, types.Stringlike:
		return true
	case *Instance:
		switch v.ClassName {
		case "Script", "LocalScript", "ModuleScript":
			return Stringlike{Value: v.Get("Source")}.IsStringlike()
		}
	}
	return false
}

// Stringlike returns Value as a string, or an empty string if the value could
// not be converted. Types that can be converted are the built-in string, []byte
// and []rune, as well as any value implementing types.Stringlike. Additionally,
// an Instance can be converted if its ClassName is either "Script",
// "LocalScript", or "ModuleScript", and it's Source property is a string-like
// type, in which case the Source is returned.
func (s Stringlike) Stringlike() string {
	switch v := s.Value.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case []rune:
		return string(v)
	case types.Stringlike:
		return v.Stringlike()
	case *Instance:
		switch v.ClassName {
		case "Script", "LocalScript", "ModuleScript":
			return Stringlike{Value: v.Get("Source")}.Stringlike()
		}
	}
	return ""
}
