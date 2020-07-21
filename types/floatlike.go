package types

// Floatlike implements rbxmk.Floatlike for a number of types.
type Floatlike struct {
	Value interface{}
}

func (f Floatlike) Floatlike() (v float64, ok bool) {
	switch v := f.Value.(type) {
	case uint8:
		return float64(v), true
	case uint16:
		return float64(v), true
	case uint32:
		return float64(v), true
	case uint64:
		return float64(v), true
	case int8:
		return float64(v), true
	case int16:
		return float64(v), true
	case int32:
		return float64(v), true
	case int64:
		return float64(v), true
	case float32:
		return float64(v), true
	case float64:
		return float64(v), true
	}
	return 0, false
}
