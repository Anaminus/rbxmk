package dt

// KindGroup is a Type that ensures the inner type is grouped unambiguously.
type KindGroup struct {
	Type
}

const K_Group = "group"

func Group(t Type) Type { return Type{Kind: KindGroup{Type: t}} }

func (t KindGroup) k()           {}
func (t KindGroup) Kind() string { return K_Group }
func (t KindGroup) String() string {
	return "(" + t.Type.String() + ")"
}
