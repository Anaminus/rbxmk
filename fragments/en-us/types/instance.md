# Summary
Represents the content of a Roblox instance.

# Description
The **Instance** type represents the content of a Roblox instance. It provides a
similar API to that of Roblox. In addition to getting and setting properties,
instances have the following members defined:

$MEMBERS

See the [Instances section](README.md#user-content-instances) for details on the
implementation of Instances.

# Constructors
## new
### Summary
Creates a new Instance.

### Description
The **new** constructor returns a new Instance of the given class. *className*
sets the [ClassName][Instance.ClassName] property of the instance. If *parent*
is specified, it sets the [Parent][Instance.Parent] property.

If *desc* is specified, then it sets the [sym.Desc][Instance.sym.Desc] member.
Additionally, new will throw an error if the class does not exist. If no
descriptor is specified, then any class name will be accepted.

If *className* is "DataModel", then a [DataModel][DataModel] value is returned.
In this case, new will throw an error if *parent* is not nil.

# Properties
## ClassName
### Summary
The class of the instance.

### Description
The **ClassName** property gets or sets the class of the instance.

Unlike in Roblox, ClassName can be modified.

## Name
### Summary
A name identifying the instance.

### Description
The **Name** property gets or sets a name identifying the instance.

## Parent
### Summary
The parent instance.

### Description
The **Parent** property gets or sets the parent of the instance, which may be
nil.

# Symbols
## AttrConfig
### Summary
The AttrConfig used by the instance.

### Description
The **AttrConfig** symbol is the [AttrConfig][AttrConfig] being used by the
instance. AttrConfig is inherited, the behavior of which is described in the
[Value inheritance](README.md#user-content-value-inheritance) section.

## Desc
### Summary
The descriptor used by the instance.

### Description
The **Desc** symbol is the descriptor being used by the instance. Desc is
inherited, the behavior of which is described in the [Value
inheritance](README.md#user-content-value-inheritance) section.

## IsService
### Summary
Whether the instance is a service.

### Description
The **IsService** symbol indicates whether the instance is a service, such as
Workspace or Lighting. This is used by some formats to determine how to encode
and decode the instance.

## Properties
### Summary
All properties of the instance.

### Description
The **Properties** symbol gets or sets all properties of the instance. Each
entry in the table is a property name mapped to the value of the property.

When getting, properties that would produce an error are ignored.

When setting, properties in the instance that are not in the table are removed.
If any property could not be set, then an error is thrown, and no properties are
set or removed.

## RawAttrConfig
### Summary
The direct AttrConfig of the instance.

### Description
The **RawAttrConfig** symbol is the raw member corresponding to to
[sym.AttrConfig][Instance.sym.AttrConfig]. It is similar to AttrConfig, except
that it considers only the direct value of the current instance. The exact
behavior of RawAttrConfig is described in the [Value
inheritance](README.md#user-content-value-inheritance) section.

## RawDesc
### Summary
The direct descriptor of the instance.

### Description
The **RawDesc** symbol is the raw member corresponding to to
[sym.Desc][Instance.sym.Desc]. It is similar to Desc, except that it considers
only the direct value of the current instance. The exact behavior of RawDesc is
described in the [Value inheritance](README.md#user-content-value-inheritance)
section.

## Reference
### Summary
A reference to the instance.

### Description
The **Reference** symbol is a string used to refer to the instance from within a
[DataModel][DataModel]. Certain formats use this to encode a reference to an
instance. For example, the RBXMX format will generate random UUIDs for its
references (e.g. "RBX8B658F72923F487FAE2F7437482EF16D").

A reference should not be expected to persist when being encoded or decoded.

# Methods
## ClearAllChildren
### Summary
Remove each child of the instance.

### Description
The **ClearAllChildren** method sets the [Parent][Instance.Parent] of each child
of the instance to nil.

Unlike in Roblox, ClearAllChildren does not affect descendants.

## Clone
### Summary
Creates a copy of the instance.

### Description
The **Clone** method returns a copy of the instance.

Unlike in Roblox, Clone does not ignore an instance if its Archivable property
is set to false.

## Descend
### Summary
Gets a descendant by name.

### Description
The **Descend** method returns a descendant of the instance by recursively
searching for each name in succession according to
[FindFirstChild][Instance.FindFirstChild]. Returns nil if a child could not be
found. Throws an error if no arguments are given.

Descend provides a safe alternative to indexing the children of an instance,
which is not implemented by rbxmk.

```lua
local face = game:Descend("Workspace", "Noob", "Head", "face")
```

## Destroy
### Summary
Removes an instance.

### Description
The **Destroy** method sets the [Parent][Instance.Parent] of the instance to
nil.

Unlike in Roblox, the Parent of the instance remains unlocked. Destroy also does
not affect descendants.

## FindFirstAncestor
### Summary
Gets an ancestor by name.

### Description
The **FindFirstAncestor** method returns the first ancestor whose
[Name][Instance.Name] equals *name*, or nil if no such instance was found.

## FindFirstAncestorOfClass
### Summary
Gets an ancestor by class name.

### Description
The **FindFirstAncestorOfClass** method returns the first ancestor of the
instance whose [ClassName][Instance.ClassName] equals *className*, or nil if no
such instance was found.

## FindFirstAncestorWhichIsA
### Summary
Gets an ancestor by inherited class name.

### Description
The **FindFirstAncestorWhichIsA** method returns the first ancestor of the
instance whose [ClassName][Instance.ClassName] inherits *className* according to
the instance's descriptor, or nil if no such instance was found. If the instance
has no descriptor, then the ClassName is compared directly.

## FindFirstChild
### Summary
Gets a child by name.

### Description
The **FindFirstChild** method returns the first child of the instance whose
[Name][Instance.Name] equals *name*, or nil if no such instance was found. If
*recurse* is true, then descendants are also searched, top-down.

## FindFirstChildOfClass
### Summary
Gets a child by class name.

### Description
The **FindFirstChildOfClass** method returns the first child of the instance
whose [ClassName][Instance.ClassName] equals *className*, or nil if no such
instance was found. If *recurse* is true, then descendants are also searched,
top-down.

## FindFirstChildWhichIsA
### Summary
Gets a child by inherited class name.

### Description
The **FindFirstChildWhichIsA** method returns the first child of the instance
whose [ClassName][Instance.ClassName] inherits *className*, or nil if no such
instance was found. If the instance has no descriptor, then the ClassName is
compared directly. If *recurse* is true, then descendants are also searched,
top-down.

## GetAttribute
### Summary
Gets an attribute.

### Description
The **GetAttribute** method returns the value of *attribute*, or nil if the
attribute is not found.

This function uses the instance's [sym.AttrConfig][Instance.sym.AttrConfig] to
select the property to decode from, which is expected to be string-like. An
error is thrown if the data could not be decoded.

See the [rbxattr format](formats.md#user-content-rbxattr) for a list of possible
attribute value types.

The [Attributes](README.md#user-content-attributes) section provides a more
general description of attributes.

## GetAttributes
### Summary
Gets all attributes.

### Description
The **GetAttributes** method returns a dictionary of attribute names mapped to
values.

This function uses the instance's [sym.AttrConfig][Instance.sym.AttrConfig] to
select the property to decode from, which is expected to be string-like. An
error is thrown if the data could not be decoded.

See the [rbxattr format](formats.md#user-content-rbxattr) for a list of possible
attribute value types.

The [Attributes](README.md#user-content-attributes) section provides a more
general description of attributes.

## GetChildren
### Summary
Gets the children.

### Description
The **GetChildren** method returns a list of children of the instance.

## GetDescendants
### Summary
Gets all descendants.

### Description
The **GetDescendants** method returns a list of descendants of the instance.

## GetFullName
### Summary
Gets a name according to the ancestors of the instance.

### Description
The **GetFullName** method returns the concatenation of the
[Name][Instance.Name] of each ancestor of the instance and the instance itself,
separated by `.` characters. If an ancestor is a [DataModel][DataModel], it is
not included.

## IsA
### Summary
Gets whether the class inherits from a given class name.

### Description
The **IsA** method returns whether the [ClassName][Instance.ClassName] inherits
from *className*, according to the instance's descriptor. If the instance has no
descriptor, then IsA returns whether ClassName equals *className*.

## IsAncestorOf
### Summary
Gets whether the instance is an ancestor.

### Description
The **IsAncestorOf** method returns whether the instance of an ancestor of
*descendant*.

## IsDescendantOf
### Summary
Gets whether the instance is an descendant.

### Description
The **IsDescendantOf** method returns whether the instance of a descendant of
*ancestor*.

## SetAttribute
### Summary
Sets an attribute.

### Description
The **SetAttribute** method sets *attribute* to *value*. If *value* is nil, then
the attribute is removed.

This function uses the instance's [sym.AttrConfig][Instance.sym.AttrConfig] to
select the property to decode from, which is expected to be string-like. This
function decodes the serialized attributes, sets the given value, then
re-encodes the attributes. An error is thrown if the data could not be decoded
or encoded.

See the [rbxattr format](formats.md#user-content-rbxattr) for a list of possible
attribute value types.

The [Attributes](README.md#user-content-attributes) section provides a more
general description of attributes.

## SetAttributes
### Summary
Sets all attributes.

### Description
The **SetAttributes** method replaces all attributes with the content of
*attributes*, which contains attribute names mapped to values.

This function uses the instance's [sym.AttrConfig][Instance.sym.AttrConfig] to
select the property to encode to. An error is thrown if the data could not be
encoded.

See the [rbxattr format](formats.md#user-content-rbxattr) for a list of possible
attribute value types.

The [Attributes](README.md#user-content-attributes) section provides a more
general description of attributes.

# Operators
## Index
### Summary
Gets a property.

### Description
The **index** operator gets the value of a property of the instance. If the
instance has a descriptor, then the the properties are enforced by the
descriptor. Otherwise, any property can be any type of value.

## Newindex
### Summary
Sets a property.

### Description
The **newindex** operator sets the value of a property of the instance. If the
instance has a descriptor, then the the properties are enforced by the
descriptor. Otherwise, any property can be set to any value.
