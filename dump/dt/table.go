package dt

import (
	"sort"
	"strings"
)

// KindTable is a Type that indicates a table with both a map part and a struct
// part.
type KindTable struct {
	Key    Type
	Value  Type
	Fields KindStruct
}

const K_Table = "table"

func Table(k KindTable) Type { return Type{Kind: k} }

func (t KindTable) k()           {}
func (t KindTable) Kind() string { return K_Table }
func (t KindTable) String() string {
	f := make([]string, 0, 16)
	for k := range t.Fields {
		f = append(f, k)
	}
	sort.Strings(f)
	var s strings.Builder
	s.WriteString("{[")
	s.WriteString(t.Key.String())
	s.WriteString("]: ")
	s.WriteString(t.Value.String())
	for _, k := range f {
		s.WriteString(", ")
		v := t.Fields[k]
		s.WriteString(k)
		s.WriteString(": ")
		s.WriteString(v.String())
	}
	s.WriteByte('}')
	return s.String()
}
