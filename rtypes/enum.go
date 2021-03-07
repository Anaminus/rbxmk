package rtypes

import (
	"sort"

	lua "github.com/anaminus/gopher-lua"
)

// Enums is a collection of Enum values.
type Enums struct {
	enums     []*Enum
	enumIndex map[string]*Enum
}

// Type returns a string identifying the type of the value.
func (*Enums) Type() string {
	return "Enums"
}

// String returns a string representation of the value.
func (*Enums) String() string {
	return "Enums"
}

// Enum returns the Enum corresponding to the given name, or nil if no such enum
// was found.
func (e *Enums) Enum(name string) *Enum {
	return e.enumIndex[name]
}

// Enum returns the enums in the collection as a list.
func (e *Enums) Enums() []*Enum {
	enums := make([]*Enum, len(e.enums))
	copy(enums, e.enums)
	return enums
}

// Enum contains a set of named values.
type Enum struct {
	name       string
	items      []*EnumItem
	nameIndex  map[string]*EnumItem
	valueIndex map[int]*EnumItem
}

// Type returns a string identifying the type of the value.
func (*Enum) Type() string {
	return "Enum"
}

// String returns a string representation of the value.
func (e *Enum) String() string {
	return e.name
}

// Name returns the name of the enum.
func (e *Enum) Name() string {
	return e.name
}

// Item returns the enum item corresponding to the given name, or nil if no such
// item exists.
func (e *Enum) Item(name string) *EnumItem {
	return e.nameIndex[name]
}

// Value returns the enum item corresponding to the given value, or nil if no
// such item exists.
func (e *Enum) Value(value int) *EnumItem {
	return e.valueIndex[value]
}

// Items returns the items of the enum in a list.
func (e *Enum) Items() []*EnumItem {
	items := make([]*EnumItem, len(e.items))
	for i, item := range e.items {
		items[i] = item
	}
	return items
}

// Pull attempts to convert a Lua value to an item of the enum. Returns nil if
// the value could not be converted.
//
// A value is converted if it is a number that matches the value of an item, if
// it is a string that matches the name of an item, or if it is an EnumItem
// userdata that is an item of the enum.
func (e *Enum) Pull(lv lua.LValue) *EnumItem {
	switch lv := lv.(type) {
	case lua.LNumber:
		if item, ok := e.valueIndex[int(lv)]; ok {
			return item
		}
	case lua.LString:
		if item, ok := e.nameIndex[string(lv)]; ok {
			return item
		}
	case *lua.LUserData:
		if item, ok := lv.Value.(*EnumItem); ok {
			if item.Enum() == e {
				return item
			}
		}
	}
	return nil
}

// EnumItem represents one possible value of an enum.
type EnumItem struct {
	enum  *Enum
	name  string
	value int
}

// Type returns a string identifying the type of the value.
func (*EnumItem) Type() string {
	return "EnumItem"
}

// String returns a string representation of the value.
func (e *EnumItem) String() string {
	if e.enum == nil {
		return e.name
	}
	return "Enum." + e.enum.name + "." + e.name
}

// Enum returns the enum to which the item belongs.
func (e *EnumItem) Enum() *Enum {
	return e.enum
}

// Name returns the name of the item.
func (e *EnumItem) Name() string {
	return e.name
}

// Value returns the value of the item.
func (e *EnumItem) Value() int {
	return e.value
}

// NewItem is passed to NewEnum to define an item of the enum.
type NewItem struct {
	Name  string
	Value int
}

// NewEnum defines a an immutable enum.
func NewEnum(name string, items ...NewItem) *Enum {
	enum := Enum{
		name:       name,
		items:      make([]*EnumItem, len(items)),
		nameIndex:  make(map[string]*EnumItem, len(items)),
		valueIndex: make(map[int]*EnumItem, len(items)),
	}
	for i, newItem := range items {
		item := EnumItem{
			enum:  &enum,
			name:  newItem.Name,
			value: newItem.Value,
		}
		enum.items[i] = &item
		enum.nameIndex[item.name] = &item
		enum.valueIndex[item.value] = &item
	}
	return &enum
}

// NewEnums creates an immutable collections of enums.
func NewEnums(enums ...*Enum) *Enums {
	sort.Slice(enums, func(i, j int) bool {
		return enums[i].name < enums[j].name
	})
	es := Enums{
		enums:     enums,
		enumIndex: make(map[string]*Enum, len(enums)),
	}
	for _, enum := range enums {
		es.enumIndex[enum.name] = enum
	}
	return &es
}
