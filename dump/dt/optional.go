package dt

// KindOptional is a Type that indicates a type of T or nil (shorthand for T | nil).
type KindOptional struct {
	Type
}

const K_Optional = "optional"

func Optional(t Type) Type { return Type{Kind: KindOptional{Type: t}} }

func (t KindOptional) k()           {}
func (t KindOptional) Kind() string { return K_Optional }
func (t KindOptional) String() string {
	return t.Type.String() + "?"
}
