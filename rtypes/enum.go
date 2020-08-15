package rtypes

import (
	"sort"
)

type Enums struct {
	enums     []*Enum
	enumIndex map[string]*Enum
}

func (Enums) Type() string {
	return "Enums"
}

func (Enums) String() string {
	return "Enums"
}

func (e *Enums) Enum(name string) *Enum {
	return e.enumIndex[name]
}

func (e *Enums) Enums() []*Enum {
	enums := make([]*Enum, len(e.enums))
	copy(enums, e.enums)
	return enums
}

type Enum struct {
	name       string
	items      []*EnumItem
	nameIndex  map[string]*EnumItem
	valueIndex map[int]*EnumItem
}

func (*Enum) Type() string {
	return "Enum"
}

func (e *Enum) String() string {
	return e.name
}

func (e *Enum) Name() string {
	return e.name
}

func (e *Enum) Item(name string) *EnumItem {
	return e.nameIndex[name]
}

func (e *Enum) Value(value int) *EnumItem {
	return e.valueIndex[value]
}

func (e *Enum) Items() []*EnumItem {
	items := make([]*EnumItem, len(e.items))
	for i, item := range e.items {
		items[i] = item
	}
	return items
}

type EnumItem struct {
	enum  *Enum
	name  string
	value int
}

func (*EnumItem) Type() string {
	return "EnumItem"
}

func (e *EnumItem) String() string {
	return "Enum." + e.enum.name + "." + e.name
}

func (e *EnumItem) Enum() *Enum {
	return e.enum
}

func (e *EnumItem) Name() string {
	return e.name
}

func (e *EnumItem) Value() int {
	return e.value
}

type NewItem struct {
	Name  string
	Value int
}

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
