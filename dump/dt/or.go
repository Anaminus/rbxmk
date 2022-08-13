package dt

import "strings"

// KindOr is a Type that indicates a union of two or more types.
type KindOr []Type

const K_Or = "or"

func Or(t ...Type) Type { return Type{Kind: KindOr(t)} }

func (t KindOr) k()           {}
func (t KindOr) Kind() string { return K_Or }
func (t KindOr) String() string {
	var s strings.Builder
	prim := true
	for _, v := range t {
		if _, ok := v.Kind.(KindPrim); !ok {
			prim = false
			break
		}
	}
	if prim {
		for i, v := range t {
			if i > 0 {
				s.WriteString(" | ")
			}
			s.WriteString(v.String())
		}
	} else {
		for i, v := range t {
			if i > 0 {
				s.WriteByte('|')
			}
			s.WriteString(v.String())
		}
	}
	return s.String()
}
