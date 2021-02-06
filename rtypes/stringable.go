package rtypes

import (
	"github.com/robloxapi/types"
)

// Stringable returns v as a string. ok is false if the value could not be
// converted. Types that can be converted are the built-in string, []byte, and
// []rune, as well as any value implementing types.Stringlike.
//
// Additionally, an Instance can be converted if it has a particular ClassName,
// and a selected property implements types.Stringlike:
//
//     ClassName         | Property
//     ------------------|---------
//     LocalizationTable | Contents
//     LocalScript       | Source
//     ModuleScript      | Source
//     Script            | Source
//
func Stringable(v interface{}) (s string, ok bool) {
	switch v := v.(type) {
	case string:
		return v, true
	case []byte:
		return string(v), true
	case []rune:
		return string(v), true
	case types.Stringlike:
		return v.Stringlike(), true
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
				return value.Stringlike(), true
			}
		}
	}
	return "", false
}
