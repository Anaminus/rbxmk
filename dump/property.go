package dump

import "github.com/anaminus/rbxmk/dump/dt"

// Properties maps a name to a Property.
type Properties map[string]Property

// Resolve implements Node.
func (p Properties) Resolve(path ...string) any {
	if len(path) == 0 {
		return p
	}
	if v, ok := p[path[0]]; ok {
		return resolveValue(path[1:], v)
	}
	return nil
}

// Property describes the API of a property.
type Property struct {
	// ValueType is the type of the property's value.
	ValueType dt.Type
	// ReadOnly indicates whether the property can be written to.
	ReadOnly bool `json:",omitempty"`

	// Summary is a fragment reference pointing to a short summary of the
	// property.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the property.
	Description string `json:",omitempty"`

	// Hidden sets whether the property should be hidden from the public API.
	Hidden bool `json:",omitempty"`
}

const V_Property = "Property"

func (v Property) v() {}

func (v Property) Kind() string { return V_Property }

// Type implements Value by returning v.ValueType.
func (v Property) Type() dt.Type {
	return v.ValueType
}

func (v Property) Index(path []string, name string) ([]string, Value) { return nil, nil }

func (v Property) Indices() []string { return nil }
