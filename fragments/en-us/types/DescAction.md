# Summary
An action that transforms a descriptor.

# Description
The **DescAction** type describes a single action that transforms a descriptor.
It has the following members:

$MEMBERS

A DescAction can be created with the [DescAction.new][DescAction.new]
constructor. It is also returned from the [RootDesc.Diff][RootDesc.Diff] method
and the [desc-patch.json](formats.md#user-content-desc-patchjson) format.

# Constructors
## new
### Summary
Creates a new DescAction.

### Description
The **new** constructor returns a new DescAction initialized with a type and an
element.

# Properties
## Element
### Summary
The type of element.

### Description
The **Element** property is the type of element to which the action applies.

## Primary
### Summary
The name of the primary element.

### Description
The **Primary** property is the name of the primary element. For example, the
class name or enum name.

## Secondary
### Summary
The name of the secondary element.

### Description
The **Secondary** property is the name of the secondary element. For example,
the name of a class member or enum item. An empty string indicates that the
action applies to the primary element.

## Type
### Summary
The type of transformation.

### Description
The **Type** property is the type of transformation performed by the action.

# Methods
## Field
### Summary
Gets a field.

### Description
The **Field** method returns the value of the *name* field, or nil if the action
has no such field.

## Fields
### Summary
Gets all fields.

### Description
The **Fields** method returns a collection of field names mapped to values.

## SetField
### Summary
Sets a field.

### Description
The **SetField** method sets the value of the *name* field to *value*.

## SetFields
### Summary
Sets all fields.

### Description
The **SetFields** method replaces all fields of the action with *fields*.
