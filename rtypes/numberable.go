package rtypes

import (
	"github.com/robloxapi/types"
)

const T_Number = "Number"

// Numberable returns v as a floating-point number. ok is false if the value
// could not be converted. Types that can be converted are the built-in numeric
// types, as well as any value implementing types.Numberlike or types.Intlike.
func Numberable(v interface{}) (n float64, ok bool) {
	switch v := v.(type) {
	case uint8:
		return float64(v), true
	case uint16:
		return float64(v), true
	case uint32:
		return float64(v), true
	case uint64:
		return float64(v), true
	case uint:
		return float64(v), true
	case int8:
		return float64(v), true
	case int16:
		return float64(v), true
	case int32:
		return float64(v), true
	case int64:
		return float64(v), true
	case int:
		return float64(v), true
	case float32:
		return float64(v), true
	case float64:
		return float64(v), true
	case types.Numberlike:
		return v.Numberlike(), true
	case types.Intlike:
		return float64(v.Intlike()), true
	}
	return 0, false
}
