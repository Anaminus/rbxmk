package rtypes

import (
	"github.com/robloxapi/types"
)

// Intlike implements types.Intlike for a number of types.
type Intlike struct {
	Value interface{}
}

// IsIntlike returns whether Value can be converted to an integer.
func (i Intlike) IsIntlike() bool {
	switch i.Value.(type) {
	case uint8, uint16, uint32, uint64, uint,
		int8, int16, int32, int64, int,
		float32, float64,
		types.Intlike, types.Numberlike:
		return true
	}
	return false
}

// Intlike returns Value as an integer, or 0 if the value could not be
// converted. Types that can be converted are the built-in numeric types, as
// well as any value implementing types.Intlike or types.Numberlike.
func (i Intlike) Intlike() int64 {
	switch v := i.Value.(type) {
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case uint:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return int64(v)
	case int:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case types.Intlike:
		return v.Intlike()
	case types.Numberlike:
		return int64(v.Numberlike())
	}
	return 0
}
