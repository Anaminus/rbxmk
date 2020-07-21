package rbxmk

// Stringlike is any value that can be converted to a sequence of bytes.
type Stringlike interface {
	Stringlike() (v []byte, ok bool)
}

// Floatlike is any value that can be converted to a floating-point number.
type Floatlike interface {
	Floatlike() (v float64, ok bool)
}

// Intlike is any value that can be converted to an integer.
type Intlike interface {
	Intlike() (v int64, ok bool)
}
