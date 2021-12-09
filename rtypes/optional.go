package rtypes

import (
	"github.com/robloxapi/types"
)

// Optional is a types.Value that may have some value, or none.
type Optional struct {
	typ   string
	value types.Value
}

func (o Optional) Type() string {
	return "Optional"
}

// Copy returns a copy of the Optional. If the underlying value does not
// implement PropValue, then None is returned.
func (t Optional) Copy() types.PropValue {
	if t.value == nil {
		return t
	}
	if pv, ok := t.value.(types.PropValue); ok {
		t.value = pv.Copy()
	} else {
		t.value = nil
	}
	return t
}

// Some returns an Optional with the given value and value's type.
func Some(value types.Value) Optional {
	if value == nil {
		panic("option value cannot be nil")
	}
	return Optional{
		typ:   value.Type(),
		value: value,
	}
}

// None returns an Optional with type t and no value.
func None(t string) Optional {
	return Optional{
		typ:   t,
		value: nil,
	}
}

// Some sets the option to have the given value and value's type.
func (o *Optional) Some(value types.Value) Optional {
	if value == nil {
		panic("option value cannot be nil")
	}
	o.typ = value.Type()
	o.value = value
	return *o
}

// None sets the option to have type t with no value.
func (o *Optional) None(t string) Optional {
	o.typ = t
	o.value = nil
	return *o
}

// Value returns the value of the option, or nil if the option has no value.
func (o Optional) Value() types.Value {
	return o.value
}

// ValueType returns the the value type of the option.
func (o Optional) ValueType() string {
	return o.typ
}

// Alias returns the underlying value.
func (o Optional) Alias() types.Value {
	if o.value == nil {
		return Nil
	}
	return o.value
}
