package rtypes

import (
	"sort"

	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/rbxdump/diff"
)

const T_Desc = "Desc"

// Desc wraps a rbxdump.Root to implement types.Value.
type Desc struct {
	*rbxdump.Root
	EnumTypes *Enums
}

// Type returns a string identifying the type of the value.
func (*Desc) Type() string {
	return T_Desc
}

// String returns a string representation of the value.
func (d *Desc) String() string {
	return "Desc"
}

func (d *Desc) Copy() *Desc {
	c := &Desc{Root: d.Root.Copy()}
	if d.EnumTypes != nil {
		c.GenerateEnumTypes()
	}
	return c
}

// ClassIsA returns whether class is a subclass of superclass. Returns false if
// d is nil, or if class or superclass are not valid classes.
func (d *Desc) ClassIsA(class, superclass string) bool {
	if d == nil {
		return false
	}
	classDesc := d.Classes[class]
	for classDesc != nil {
		if classDesc.Name == superclass {
			return true
		}
		classDesc = d.Classes[classDesc.Superclass]
	}
	return false
}

// Class returns the class descriptor from a class name. Returns nil if d is
// nil, or the class was not found.
func (d *Desc) Class(name string) (class *rbxdump.Class) {
	if d == nil {
		return nil
	}
	return d.Classes[name]
}

// Enum returns the enum descriptor from an enum name. Returns nil if d is nil,
// or the enum was not found.
func (d *Desc) Enum(name string) (enum *rbxdump.Enum) {
	if d == nil {
		return nil
	}
	return d.Enums[name]
}

// EnumItem returns the enum item descriptor from an enum and item name. Returns
// nil if d is nil, or the enum or item was not found.
func (d *Desc) EnumItem(enum, name string) (item *rbxdump.EnumItem) {
	if d == nil {
		return nil
	}
	e := d.Enums[enum]
	if e == nil {
		return nil
	}
	return e.Items[name]
}

// Member gets a member descriptor from a class, or any class it inherits from.
// Returns nil if d is nil, or if the class or member was not found.
func (d *Desc) Member(class, name string) (member rbxdump.Member) {
	if d == nil {
		return nil
	}
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
// from. Returns nil if d is nil, or if the class or member was not found.
func (d *Desc) Property(class, name string) *rbxdump.Property {
	if d == nil {
		return nil
	}
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
func (d *Desc) GenerateEnumTypes() {
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
func (d *Desc) Of(inst *Instance) *Desc {
	if inst != nil {
		if desc := inst.Desc(); desc != nil {
			return desc
		}
	}
	return d
}

const T_ClassDesc = "ClassDesc"

// ClassDesc wraps a rbxdump.Class to implement types.Value.
type ClassDesc rbxdump.Class

// Type returns a string identifying the type of the value.
func (ClassDesc) Type() string {
	return T_ClassDesc
}

const T_MemberDesc = "MemberDesc"

// MemberDesc wraps a rbxdump.Member to implement types.Value.
type MemberDesc struct {
	rbxdump.Member
}

// Type returns a string identifying the type of the value.
func (MemberDesc) Type() string {
	return T_MemberDesc
}

const T_PropertyDesc = "PropertyDesc"

// PropertyDesc wraps a rbxdump.Property to implement types.Value.
type PropertyDesc rbxdump.Property

// Type returns a string identifying the type of the value.
func (PropertyDesc) Type() string {
	return T_PropertyDesc
}

const T_FunctionDesc = "FunctionDesc"

// FunctionDesc wraps a rbxdump.Function to implement types.Value.
type FunctionDesc rbxdump.Function

// Type returns a string identifying the type of the value.
func (FunctionDesc) Type() string {
	return T_FunctionDesc
}

const T_EventDesc = "EventDesc"

// EventDesc wraps a rbxdump.Event to implement types.Value.
type EventDesc rbxdump.Event

// Type returns a string identifying the type of the value.
func (EventDesc) Type() string {
	return T_EventDesc
}

const T_CallbackDesc = "CallbackDesc"

// CallbackDesc wraps a rbxdump.Callback to implement types.Value.
type CallbackDesc rbxdump.Callback

// Type returns a string identifying the type of the value.
func (CallbackDesc) Type() string {
	return T_CallbackDesc
}

const T_ParameterDesc = "ParameterDesc"

// ParameterDesc wraps a rbxdump.Parameter to implement types.Value.
type ParameterDesc struct {
	rbxdump.Parameter
}

// Type returns a string identifying the type of the value.
func (ParameterDesc) Type() string {
	return T_ParameterDesc
}

const T_TypeDesc = "TypeDesc"

// TypeDesc wraps a rbxdump.Type to implement types.Value.
type TypeDesc struct {
	Embedded rbxdump.Type
}

// Type returns a string identifying the type of the value.
func (TypeDesc) Type() string {
	return T_TypeDesc
}

const T_EnumDesc = "EnumDesc"

// EnumDesc wraps a rbxdump.Enum to implement types.Value.
type EnumDesc rbxdump.Enum

// Type returns a string identifying the type of the value.
func (EnumDesc) Type() string {
	return T_EnumDesc
}

const T_EnumItemDesc = "EnumItemDesc"

// EnumItemDesc wraps a rbxdump.EnumItem to implement types.Value.
type EnumItemDesc rbxdump.EnumItem

// Type returns a string identifying the type of the value.
func (EnumItemDesc) Type() string {
	return T_EnumItemDesc
}

const T_DescActions = "DescActions"

// DescActions is a list of DescAction values that implements types.Value.
type DescActions []*DescAction

// Type returns a string identifying the type of the value.
func (DescActions) Type() string {
	return T_DescActions
}

const T_DescAction = "DescAction"

// DescAction wraps a diff.Action to implement types.Value.
type DescAction struct {
	diff.Action
}

// Type returns a string identifying the type of the value.
func (*DescAction) Type() string {
	return T_DescAction
}

// String returns a string representation of the value.
func (a *DescAction) String() string {
	return a.Action.String()
}

const T_DescFields = "DescFields"

// DescFields wraps a rbxdump.Fields to implement types.Value.
type DescFields rbxdump.Fields

// Type returns a string identifying the type of the value.
func (DescFields) Type() string {
	return T_DescFields
}
