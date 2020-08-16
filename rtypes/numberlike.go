package rtypes

import (
	"github.com/robloxapi/types"
)

// Numberlike implements types.Numberlike for a number of types.
type Numberlike struct {
	Value interface{}
}

// IsIntlike returns whether Value can be converted to a floating-point number.
func (n Numberlike) IsNumberlike() bool {
	switch n.Value.(type) {
	case uint8, uint16, uint32, uint64, uint,
		int8, int16, int32, int64, int,
		float32, float64,
		types.Numberlike, types.Intlike:
		return true
	}
	return false
}

// Numberlike returns Value as a floating-point number, or 0 if the value could
// not be converted. Types that can be converted are the built-in numeric types,
// as well as any value implementing types.Numberlike or types.Intlike.
func (n Numberlike) Numberlike() float64 {
	switch v := n.Value.(type) {
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case uint:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case int:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return float64(v)
	case types.Numberlike:
		return v.Numberlike()
	case types.Intlike:
		return float64(v.Intlike())
	}
	return 0
}
