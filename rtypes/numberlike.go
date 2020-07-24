package rtypes

import (
	"github.com/anaminus/rbxmk"
)

// Numberlike implements rbxmk.Numberlike for a number of types.
type Numberlike struct {
	Value interface{}
}

func (n Numberlike) IsIntlike() bool {
	switch n.Value.(type) {
	case uint8, uint16, uint32, uint64, uint,
		int8, int16, int32, int64, int,
		float32, float64,
		rbxmk.Numberlike, rbxmk.Intlike:
		return true
	}
	return false
}

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
	case rbxmk.Numberlike:
		return v.Numberlike()
	case rbxmk.Intlike:
		return float64(v.Intlike())
	}
	return 0
}
