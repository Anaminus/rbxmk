package rtypes

import (
	"sort"

	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/rbxdump/diff"
	"github.com/robloxapi/types"
)

type RootDesc struct {
	*rbxdump.Root
	EnumTypes *Enums
}

func (*RootDesc) Type() string {
	return "RootDesc"
}

func (d *RootDesc) String() string {
	return "RootDesc"
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

type ClassDesc struct {
	*rbxdump.Class
}

func (ClassDesc) Type() string {
	return "ClassDesc"
}

func (d ClassDesc) String() string {
	return "ClassDesc"
}

type PropertyDesc struct {
	*rbxdump.Property
}

func (PropertyDesc) Type() string {
	return "PropertyDesc"
}

func (d PropertyDesc) String() string {
	return "PropertyDesc"
}

type FunctionDesc struct {
	*rbxdump.Function
}

func (FunctionDesc) Type() string {
	return "FunctionDesc"
}

func (d FunctionDesc) String() string {
	return "FunctionDesc"
}

type EventDesc struct {
	*rbxdump.Event
}

func (EventDesc) Type() string {
	return "EventDesc"
}

func (d EventDesc) String() string {
	return "EventDesc"
}

type CallbackDesc struct {
	*rbxdump.Callback
}

func (CallbackDesc) Type() string {
	return "CallbackDesc"
}

func (d CallbackDesc) String() string {
	return "CallbackDesc"
}

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

type ParameterDesc struct {
	rbxdump.Parameter
}

func (ParameterDesc) Type() string {
	return "ParameterDesc"
}

func (d ParameterDesc) String() string {
	return "ParameterDesc"
}

type TypeDesc struct {
	Embedded rbxdump.Type
}

func (TypeDesc) Type() string {
	return "TypeDesc"
}

func (d TypeDesc) String() string {
	return "TypeDesc"
}

type EnumDesc struct {
	*rbxdump.Enum
}

func (EnumDesc) Type() string {
	return "EnumDesc"
}

func (d EnumDesc) String() string {
	return "EnumDesc"
}

type EnumItemDesc struct {
	*rbxdump.EnumItem
}

func (EnumItemDesc) Type() string {
	return "EnumItemDesc"
}

func (d EnumItemDesc) String() string {
	return "EnumItemDesc"
}

type DescActions []*DescAction

func (DescActions) Type() string {
	return "DescActions"
}

type DescAction struct {
	diff.Action
}

func (DescAction) Type() string {
	return "DescAction"
}

func (a DescAction) String() string {
	return a.Action.String()
}
