package rtypes

import (
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
)

type RootDesc struct {
	*rbxdump.Root
}

func (RootDesc) Type() string {
	return "RootDesc"
}

func (d RootDesc) String() string {
	return "RootDesc"
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
	*rbxdump.Parameter
}

func (ParameterDesc) Type() string {
	return "ParameterDesc"
}

func (d ParameterDesc) String() string {
	return "ParameterDesc"
}

type TypeDesc struct {
	Embedded *rbxdump.Type
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
