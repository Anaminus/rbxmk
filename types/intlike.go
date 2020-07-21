package types

// Intlike implements rbxmk.Intlike for a number of types.
type Intlike struct {
	Value interface{}
}

func (f Intlike) Intlike() (v int64, ok bool) {
	switch v := f.Value.(type) {
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
	}
	return 0, false
}
