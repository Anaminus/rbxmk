package dt

// KindDictionary is a Type that indicates a table where each element maps a
// string to a value.
type KindDictionary struct {
	Type
}

const K_Dictionary = "dictionary"

func Dictionary(v Type) Type { return Type{Kind: KindDictionary{Type: v}} }

func (t KindDictionary) k()           {}
func (t KindDictionary) Kind() string { return K_Dictionary }
func (t KindDictionary) String() string {
	return "{[string]: " + t.Type.String() + "}"
}
