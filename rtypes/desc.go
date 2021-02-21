package rtypes

import (
	"sort"

	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/rbxdump/diff"
	"github.com/robloxapi/types"
)

// RootDesc wraps a rbxdump.Root to implement types.Value.
type RootDesc struct {
	*rbxdump.Root
	EnumTypes *Enums
}

// Type returns a string identifying the type of the value.
func (*RootDesc) Type() string {
	return "RootDesc"
}

// String returns a string representation of the value.
func (d *RootDesc) String() string {
	return "RootDesc"
}

// Member gets a member descriptor from a class, or any class it inherits from.
func (d *RootDesc) Member(class, name string) (member rbxdump.Member) {
	classDesc := d.Classes[class]
	for classDesc != nil {
		if member = classDesc.Members[name]; member != nil {
			return member
		}
		classDesc = d.Classes[classDesc.Superclass]
	}
	return nil
}

// Property gets a property descriptor from a class, or any class it inherits
// from.
func (d *RootDesc) Property(class, name string) *rbxdump.Property {
	classDesc := d.Classes[class]
	for classDesc != nil {
		if member := classDesc.Members[name]; member != nil {
			if prop, ok := member.(*rbxdump.Property); ok {
				return prop
			}
			return nil
		}
		classDesc = d.Classes[classDesc.Superclass]
	}
	return nil
}

// GenerateEnumTypes sets EnumTypes to a collection of enum values generated
// from the root's enum descriptors.
func (d *RootDesc) GenerateEnumTypes() {
	enums := make([]*Enum, 0, len(d.Enums))
	for name, enumDesc := range d.Enums {
		itemDescs := make([]*rbxdump.EnumItem, 0, len(enumDesc.Items))
		for _, itemDesc := range enumDesc.Items {
			itemDescs = append(itemDescs, itemDesc)
		}
		sort.Slice(itemDescs, func(i, j int) bool {
			if itemDescs[i].Index == itemDescs[j].Index {
				return itemDescs[i].Value < itemDescs[j].Value
			}
			return itemDescs[i].Index < itemDescs[j].Index
		})
		items := make([]NewItem, len(itemDescs))
		for i, itemDesc := range itemDescs {
			items[i] = NewItem{
				Name:  itemDesc.Name,
				Value: itemDesc.Value,
			}
		}
		enums = append(enums, NewEnum(name, items...))
	}
	d.EnumTypes = NewEnums(enums...)
}

// Of returns the root descriptor of an instance. If inst is nil, r is returned.
func (d *RootDesc) Of(inst *Instance) *RootDesc {
	if inst != nil {
		if desc := inst.Desc(); desc != nil {
			return desc
		}
	}
	return d
}

// ClassDesc wraps a rbxdump.Class to implement types.Value.
type ClassDesc struct {
	*rbxdump.Class
}

// Type returns a string identifying the type of the value.
func (ClassDesc) Type() string {
	return "ClassDesc"
}

// String returns a string representation of the value.
func (d ClassDesc) String() string {
	return "ClassDesc"
}

// PropertyDesc wraps a rbxdump.Property to implement types.Value.
type PropertyDesc struct {
	*rbxdump.Property
}

// Type returns a string identifying the type of the value.
func (PropertyDesc) Type() string {
	return "PropertyDesc"
}

// String returns a string representation of the value.
func (d PropertyDesc) String() string {
	return "PropertyDesc"
}

// FunctionDesc wraps a rbxdump.Function to implement types.Value.
type FunctionDesc struct {
	*rbxdump.Function
}

// Type returns a string identifying the type of the value.
func (FunctionDesc) Type() string {
	return "FunctionDesc"
}

// String returns a string representation of the value.
func (d FunctionDesc) String() string {
	return "FunctionDesc"
}

// EventDesc wraps a rbxdump.Event to implement types.Value.
type EventDesc struct {
	*rbxdump.Event
}

// Type returns a string identifying the type of the value.
func (EventDesc) Type() string {
	return "EventDesc"
}

// String returns a string representation of the value.
func (d EventDesc) String() string {
	return "EventDesc"
}

// CallbackDesc wraps a rbxdump.Callback to implement types.Value.
type CallbackDesc struct {
	*rbxdump.Callback
}

// Type returns a string identifying the type of the value.
func (CallbackDesc) Type() string {
	return "CallbackDesc"
}

// String returns a string representation of the value.
func (d CallbackDesc) String() string {
	return "CallbackDesc"
}

// NewMemberDesc returns a rbxdump.Member wrapped in the corresponding member
// descriptor.
func NewMemberDesc(member rbxdump.Member) types.Value {
	switch member := member.(type) {
	case *rbxdump.Property:
		return PropertyDesc{Property: member}
	case *rbxdump.Function:
		return FunctionDesc{Function: member}
	case *rbxdump.Event:
		return EventDesc{Event: member}
	case *rbxdump.Callback:
		return CallbackDesc{Callback: member}
	}
	return nil
}

// ParameterDesc wraps a rbxdump.Parameter to implement types.Value.
type ParameterDesc struct {
	rbxdump.Parameter
}

// Type returns a string identifying the type of the value.
func (ParameterDesc) Type() string {
	return "ParameterDesc"
}

// String returns a string representation of the value.
func (d ParameterDesc) String() string {
	return "ParameterDesc"
}

// TypeDesc wraps a rbxdump.Type to implement types.Value.
type TypeDesc struct {
	Embedded rbxdump.Type
}

// Type returns a string identifying the type of the value.
func (TypeDesc) Type() string {
	return "TypeDesc"
}

// String returns a string representation of the value.
func (d TypeDesc) String() string {
	return "TypeDesc"
}

// EnumDesc wraps a rbxdump.Enum to implement types.Value.
type EnumDesc struct {
	*rbxdump.Enum
}

// Type returns a string identifying the type of the value.
func (EnumDesc) Type() string {
	return "EnumDesc"
}

// String returns a string representation of the value.
func (d EnumDesc) String() string {
	return "EnumDesc"
}

// EnumItemDesc wraps a rbxdump.EnumItem to implement types.Value.
type EnumItemDesc struct {
	*rbxdump.EnumItem
}

// Type returns a string identifying the type of the value.
func (EnumItemDesc) Type() string {
	return "EnumItemDesc"
}

// String returns a string representation of the value.
func (d EnumItemDesc) String() string {
	return "EnumItemDesc"
}

// DescActions is a list of DescAction values that implements types.Value.
type DescActions []*DescAction

// Type returns a string identifying the type of the value.
func (DescActions) Type() string {
	return "DescActions"
}

// DescAction wraps a diff.Action to implement types.Value.
type DescAction struct {
	diff.Action
}

// Type returns a string identifying the type of the value.
func (DescAction) Type() string {
	return "DescAction"
}

// String returns a string representation of the value.
func (a DescAction) String() string {
	return a.Action.String()
}
