package rtypes

import (
	"github.com/robloxapi/types"
)

// Intable returns Value as an integer. ok is false if the value could not be
// converted. Types that can be converted are the built-in numeric types, as
// well as any value implementing types.Intlike or types.Numberlike.
func Intable(v interface{}) (i int64, ok bool) {
	switch v := v.(type) {
	case uint8:
		return int64(v), true
	case uint16:
		return int64(v), true
	case uint32:
		return int64(v), true
	case uint64:
		return int64(v), true
	case uint:
		return int64(v), true
	case int8:
		return int64(v), true
	case int16:
		return int64(v), true
	case int32:
		return int64(v), true
	case int64:
		return int64(v), true
	case int:
		return int64(v), true
	case float32:
		return int64(v), true
	case float64:
		return int64(v), true
	case types.Intlike:
		return v.Intlike(), true
	case types.Numberlike:
		return int64(v.Numberlike()), true
	}
	return 0, false
}
