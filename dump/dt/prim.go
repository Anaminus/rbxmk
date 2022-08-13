package dt

// KindPrim is a Type that indicates the name of some defined type.
type KindPrim string

const K_Primitive = "primitive"

func Prim(t string) Type { return Type{Kind: KindPrim(t)} }

func (t KindPrim) k()           {}
func (t KindPrim) Kind() string { return K_Primitive }
func (t KindPrim) String() string {
	return string(t)
}
