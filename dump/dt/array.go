package dt

// KindArray is a Type that indicates an array of elements of some type.
type KindArray struct {
	Type
}

const K_Array = "array"

func Array(t Type) Type { return Type{Kind: KindArray{Type: t}} }

func (t KindArray) k()           {}
func (t KindArray) Kind() string { return K_Array }
func (t KindArray) String() string {
	return "{" + t.Type.String() + "}"
}
