# Summary
Describes an API.

# Description
The **RootDesc** type describes an entire API. It has the following members:

$MEMBERS

# Constructors
## new
### Summary
Creates a RootDesc.

### Description
The **new** constructor creates a new RootDesc.

# Methods
## Class
### Summary
Gets data about a class.

### Description
The **Class** method returns the class of the API corresponding to the name
*class*, or nil if the class does not exist.

## ClassTag
### Summary
Gets whether a tag is set for a class.

### Description
Returns whether *tag* is set for the class corresponding to the name *class*.
Returns nil if the class does not exist.

## Classes
### Summary
Gets a list of class names.

### Description
The **Classes** method returns a list of names of all the classes of the API.

## Diff
### Summary
Gets the difference between two descriptors.

### Description
The **Diff** method compares the root descriptor with another and returns the
differences between them. A nil value for *next* is treated the same as an empty
descriptor. The result is a list of actions that describe how to transform the
descriptor into *next*.

## Enum
### Summary
Gets data about an enum.

### Description
The **Enum** method returns the enum of the API corresponding to the name
*enum*, or nil if the enum does not exist.

## EnumItem
### Summary
Gets data about an enum item.

### Description
The **EnumItem** method returns the enum item of the API corresponding to the
enum name *enum* and item name *item*, or nil if the enum or item does not
exist.

## EnumItemTag
### Summary
Gets whether a tag is set for an enum item.

### Description
Returns whether *tag* is set for the enum item corresponding to the name *item*
of the enum named *enum*. Returns nil if the enum or item does not exist.

## EnumItems
### Summary
Gets a list of item names of an enum.

### Description
The **Classes** method returns a list of names of all the items of the enum
corresponding to the name *enum*. Returns nil if the enum does not exist.

## EnumTag
### Summary
Gets whether a tag is set for an enum.

### Description
Returns whether *tag* is set for the enum corresponding to the name *enum*.
Returns nil if the enum does not exist.

## EnumTypes
### Summary
Gets enum values generated from the descriptor.

### Description
The **EnumTypes** method returns a set of enum values generated from the current
state of the RootDesc. These enums are associated with the RootDesc, and may be
used by certain properties, so it is important to generate them before operating
on such properties. Additionally, EnumTypes should be called after modifying
enum and enum item descriptors, to regenerate the enum values.

The API of the resulting enums matches that of Roblox's Enums type. A common
pattern is to assign the result of EnumTypes to the "Enum" variable so that it
matches Roblox's API:

```lua
Enum = rootDesc:EnumTypes()
print(Enum.NormalId.Front)
```

## Enums
### Summary
Gets a list of enum names.

### Description
The **Enums** method returns a list of names of all the enums of the API.

## Member
### Summary
Gets data about a class member.

### Description
The **Member** method returns the class member of the API corresponding to the
class name *enum* and member name *item*, or nil if the class or member does not
exist.

## MemberTag
### Summary
Gets whether a tag is set for a class member.

### Description
Returns whether *tag* is set for the class member corresponding to the name
*member* of the class named *class*. Returns nil if the class or member does not
exist.

## Members
### Summary
Gets a list of member names of a class.

### Description
Returns whether *tag* is set for the class member corresponding to the name
*member* of the class named *class*. Returns nil if the class or member does not
exist.

## Patch
### Summary
Transforms the descriptor.

### Description
The **Patch** method transforms the root descriptor according to a list of
actions. Each action in the list is applied in order. Actions that are
incompatible are ignored.

## SetClass
### Summary
Sets data about a class.

### Description
The **SetClass** method sets the fields of the class corresponding to the name
*class*. If the class already exists, then only non-nil fields are set. Fields
with the incorrect type are ignored. If *desc* is nil, then the class is
removed.

Returns false if *desc* is nil and no class exists. Returns true otherwise.

## SetEnum
### Summary
Sets data about an enum.

### Description
The **SetEnum** method sets the fields of the enum corresponding to the name
*enum*. If the enum already exists, then only non-nil fields are set. Fields
with the incorrect type are ignored. If *desc* is nil, then the enum is removed.

Returns false if *desc* is nil and no enum exists. Returns true otherwise.

## SetEnumItem
### Summary
Sets data about an enum item.

### Description
The **SetEnumItem** method sets the fields of the enum item corresponding to the
name *item* of the enum named *enum*. If the item already exists, then only
non-nil fields are set. Fields with the incorrect type are ignored. If *desc* is
nil, then the enum is removed.

Returns nil if the enum does not exist. Returns false if *desc* is nil and no
item exists. Returns true otherwise.

## SetMember
### Summary
Sets data about a class member.

### Description
The **SetMember** method sets the fields of the member corresponding to the name
*member* of the class named *class*. If the member already exists, then only
non-nil fields are set. Fields with the incorrect type are ignored. If *desc* is
nil, then the member is removed.

Returns nil if the class does not exist. Returns false if *desc* is nil and no
member exists. Returns true otherwise.
