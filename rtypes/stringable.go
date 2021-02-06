package rtypes

import (
	"github.com/robloxapi/types"
)

// Stringable converts a number of types to a string.
type Stringable struct {
	Value interface{}
}

// IsStringable returns whether Value can be converted to a string.
func (s Stringable) IsStringable() bool {
	switch v := s.Value.(type) {
	case string, []byte, []rune, types.Stringlike:
		return true
	case *Instance:
		var value types.Value
		switch v.ClassName {
		case "Script", "LocalScript", "ModuleScript":
			value = v.Get("Source")
		case "LocalizationTable":
			value = v.Get("Contents")
		}
		if value != nil {
			_, ok := value.(types.Stringlike)
			return ok
		}
	}
	return false
}

// Stringable returns Value as a string, or an empty string if the value could
// not be converted. Types that can be converted are the built-in string,
// []byte, and []rune, as well as any value implementing types.Stringlike.
//
// Additionally, an Instance can be converted if it has a particular ClassName,
// and a selected property has a stringable type that isn't an Instance:
//
//     ClassName         | Property
//     ------------------|---------
//     LocalizationTable | Contents
//     LocalScript       | Source
//     ModuleScript      | Source
//     Script            | Source
//
func (s Stringable) Stringable() string {
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
		var value types.Value
		switch v.ClassName {
		case "Script", "LocalScript", "ModuleScript":
			value = v.Get("Source")
		case "LocalizationTable":
			value = v.Get("Contents")
		}
		if value != nil {
			if value, ok := value.(types.Stringlike); ok {
				return value.Stringlike()
			}
		}
	}
	return ""
}
