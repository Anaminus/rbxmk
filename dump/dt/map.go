package dt

// KindMap is a Type that indicates a table where each element maps a key to a
// value.
type KindMap struct {
	K Type
	V Type
}

const K_Map = "map"

func Map(k, v Type) Type { return Type{Kind: KindMap{K: k, V: v}} }

func (t KindMap) k()           {}
func (t KindMap) Kind() string { return K_Map }
func (t KindMap) String() string {
	return "{[" + t.K.String() + "]: " + t.V.String() + "}"
}
