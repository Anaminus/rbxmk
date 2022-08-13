package dt

import (
	"sort"
	"strings"
)

// KindStruct is a Type that indicates a table with a number of named fields.
type KindStruct map[string]Type

const K_Struct = "struct"

func Struct(k KindStruct) Type { return Type{Kind: k} }

func (t KindStruct) k()           {}
func (t KindStruct) Kind() string { return K_Struct }
func (t KindStruct) String() string {
	f := make([]string, 0, 16)
	var variadic Type
	for k, v := range t {
		if k == "..." {
			w := v
			variadic = w
			continue
		}
		f = append(f, k)
	}
	var s strings.Builder
	s.WriteByte('{')
	sort.Strings(f)
	for i, k := range f {
		if i > 0 {
			s.WriteString(", ")
		}
		v := t[k]
		s.WriteString(k)
		s.WriteString(": ")
		s.WriteString(v.String())
	}
	if variadic.Kind != nil {
		if len(f) > 0 {
			s.WriteString(", ")
		}
		s.WriteString("...: ")
		s.WriteString(variadic.String())
	}
	s.WriteByte('}')
	return s.String()
}
