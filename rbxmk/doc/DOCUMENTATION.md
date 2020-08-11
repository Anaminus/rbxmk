[DRAFT]

# Documentation
This document provides details on how rbxmk works. For a basic overview, see
[USAGE.md](USAGE.md).

This document uses Luau type annotation syntax to describe the API of an
element. Some liberties are taken for APIs not currently supported by the Luau
syntax. For example, `...` indicates variable parameters. Additionally, the
following types are predefined for documentation purposes:

```luau
-- A list of values of type T.
type Array<T> = {[number]: T}

-- A table mapping a name to a value of type T.
type Dictionary<T> = {[string]: T}
```

# Environment
The Lua environment provided by rbxmk is packaged as a number of libraries. Some
libraries are loaded under a specific name, while others are loaded directly
into the global environment.

Library   | Description
----------|------------
(base)    | The Lua 5.1 standard library, abridged.
rbxmk     | An interface to the rbxmk engine, and the rbxmk environment.
(roblox)  | An environment matching the Roblox Lua API.
os        | Extensions to the standard os library.
sym       | Symbols for accessing instance metadata.
types     | Fallbacks for constructing certain types.
(sources) | An assortment of libraries for interfacing with the various external sources.

## Base library
The following items from the Lua 5.1 standard library are included:

- `_G`
- `_VERSION`
- `assert`
- `error`
- `getmetatable`
- `ipairs`
- `next`
- `pairs`
- `pcall`
- `print`
- `select`
- `setmetatable`
- `tonumber`
- `tostring`
- `type`
- `unpack`
- `xpcall`
- Entire `math` library
- Entire `string` library except `string.dump`
- Entire `table` library
- `os.clock`
- `os.date`
- `os.difftime`
- `os.time`

## `rbxmk` library
The `rbxmk` library contains functions related to the rbxmk engine.

Name           | Description
---------------|------------
`decodeFormat` | Deserialize data from bytes.
`desc`         | Gets or sets the global descriptor.
`diffDesc`     | Get the differences between two descriptors.
`encodeFormat` | Serialize data into bytes.
`loadFile`     | Loads the content of a file as a function.
`loadString`   | Loads a string as a function.
`meta`         | Set metadata of an Instance.
`newDesc`      | Create a new descriptor.
`patchDesc`    | Transform a descriptor by applying differences.
`readSource`   | Read bytes from an external source.
`runFile`      | Runs a file as a Lua chunk.
`runString`    | Runs a string as a Lua chunk.
`writeSource`  | Write bytes to an external source.

### `rbxmk.decodeFormat(format: string, bytes: BinaryString): (value: any)`
The `decodeFormat` function decodes *bytes* into a value according to *format*.
The exact details of each format are described in the
[Formats](#user-content-formats) section.

### `rbxmk.desc: RootDesc`
The `desc` field gets or sets the global root descriptor. Most elements that
utilize a root descriptor will fallback to the global descriptor when possible.

### `rbxmk.diffDesc(prev: RootDesc?, next: RootDesc?): (diff: Array<DescAction>)`
The `diffDesc` function compares two descriptors and returns the differences
between the two. A nil value for *prev* or *next* is treated the same as an
empty descriptor. The result is a list of actions that describes how to
transform *prev* into *next*.

### `rbxmk.encodeFormat(format: string, value: any): (bytes: BinaryString)`
The `encodeFormat` function encodes *value* into a sequence of bytes according
to *format*. The exact details of each format are described in the
[Formats](#user-content-formats) section.

### `rbxmk.loadFile(path: string): (func: function)`
The `loadFile` function loads the content of a file as a Lua function. *path* is
the path to the file.

The function runs in the context of the calling script.

### `rbxmk.loadString(source: string): (func: function)`
The `loadString` function loads the a string as a Lua function. *source* is the
string to load.

The function runs in the context of the calling script.

### `rbxmk.newDesc(name: string): Descriptor`
The `newDesc` function creates a new descriptor object.

`newDesc` returns a value of whose type corresponds to the given name. *name*
may be one of the following:

Name          | Returned type
--------------|--------------
`"Root"`      | RootDesc
`"Class"`     | ClassDesc
`"Property"`  | PropertyDesc
`"Function"`  | FunctionDesc
`"Event"`     | EventDesc
`"Callback"`  | CallbackDesc
`"Parameter"` | ParameterDesc
`"Type"`      | TypeDesc
`"Enum"`      | EnumDesc
`"EnumItem"`  | EnumItemDesc

TypeDesc values are immutable. To set the fields, they can be passed as extra
arguments to newDesc:

```lua
-- Sets .Category and .Name, respectively.
local typeDesc = rbxmk.newDesc("Type", "Category", "Name")
```

ParameterDesc values are immutable. To set the fields, they can be passed as
extra arguments to newDesc:

```lua
-- Sets .Type, .Name, and .Default, respectively.
-- No default value
local paramDesc = rbxmk.newDesc("Parameter", typeDesc, "paramName")
-- Default value
local paramDesc = rbxmk.newDesc("Parameter", typeDesc, "paramName", "ParamDefault")
```

### `rbxmk.patchDesc(desc: RootDesc, actions: Array<DescAction>)`
The `patchDesc` function transforms a descriptor according to a list of actions.
Each action in the list is applied in order. Actions that do not apply are
ignored.

### `rbxmk.readSource(source: string, args: ...any): (bytes: BinaryString)`
The `readSource` function reads a sequence of bytes from an external source
indicated by *source*. *args* depends on the source. The exact details of each
source are described in the [Sources](#user-content-sources) section.

### `rbxmk.runFile(path: string, args: ...any): (results: ...any)`
The `runFile` function runs the content of a file as a Lua script. *path* is the
path to the file. *args* are passed into the script as arguments. Returns the
values returned by the script.

The script runs in the context of the referred file. Files cannot be run
recursively; if a file is already running as a script, then `runFile` will throw
an error.

### `rbxmk.runString(source: string, args: ...any): (results: ...any)`
The `runString` function runs a string as a Lua script. *source* is the string
to run. *args* are passed into the script as arguments. Returns the values
returned by the script.

The script runs in the context of the calling script.

### `rbxmk.writeSource(source: string, bytes: BinaryString, args: ...any)`
The `writeSource` function writes a sequence of bytes to an external source
indicated by *source*. *args* depends on the source. The exact details of each
source are described in the [Sources](#user-content-sources) section.

## Roblox library
The Roblox library includes an environment similar to the Roblox Lua API. This
is included directly into the global environment.

The `typeof` function is included to get the type of a userdata. In addition to
the usual Roblox types, `typeof` will work for various types specific to rbxmk.

The following type constructors are included:

- `Axes`
- `BrickColor`
- `CFrame`
- `Color3`
- `ColorSequence`
- `ColorSequenceKeypoint`
- `Faces`
- `Instance`
- `NumberRange`
- `NumberSequence`
- `NumberSequenceKeypoint`
- `PhysicalProperties`
- `Ray`
- `Rect`
- `Region3`
- `Region3int16`
- `UDim`
- `UDim2`
- `Vector2`
- `Vector2int16`
- `Vector3`
- `Vector3int16`

Additionally, the `DataModel.new` constructor creates a special Instance of the
DataModel class, to be used to contain instances in a game tree. Unlike a normal
Instance, the ClassName property cannot be modified, and other properties
determine metadata used by certain formats.

## `os` library
The `os` library is an extension to the standard library. The following
additional functions are included:

Name     | Description
---------|------------
`dir`    | Gets a list of files in a directory.
`expand` | Expands predefined file path variables.
`getenv` | Gets an environment variable.
`join`   | Joins a number of file paths together.
`split`  | Splits a file path into its components.

### `os.dir(path: string): Array<File>`
The `dir` function returns a list of files in the given directory.

Each file is a table with the following fields:

Field   | Type    | Description
--------|---------|------------
Name    | string  | The base name of the file.
IsDir   | boolean | Whether the file is a directory.
Size    | number  | The size of the file, in bytes.
ModTime | number  | The modification time of the file, in Unix time.

### `os.expand(path: string): string`
The `expand` function scans *path* for certain variables of the form `$var` or
`${var}` an expands them. The following variables are expanded:

Variable                                    | Description
--------------------------------------------|------------
`$script_name`, `$sn`                       | The base name of the currently running script.
`$script_directory`, `$script_dir`, `$sd`   | The directory of the currently running script.
`$working_directory`, `$working_dir`, `$wd` | The current working directory.
`$temp_directory`, `$temp_dir`, `$tmp`      | The directory for temporary files.

### `os.getenv(name: string?): string | Array<string>`
The `getenv` function returns the value of the *name* environment variable. If
*name* is not specified, then a list of environment variables is returned.

### `os.join(paths: ...string) string`
The `join` function joins each *path* element into a single path, separating
them using the operating system's path separator. This also cleans up the path.

### `os.split(path: string, components: ...string): ...string`
The `split` function returns the components of a file path.

Component | `project/scripts/main.script.lua` | Description
----------|-----------------------------------|------------
`base`    | `main.script.lua`                 | The file name; the last element of the path.
`dir`     | `project/scripts`                 | The directory; all but the last element of the path.
`ext`     | `.lua`                            | The extension; the suffix starting at the last dot of the last element of the path.
`fext`    | `.script.lua`                     | The format extension, as determined by registered formats.
`fstem`   | `main`                            | The base without the format extension.
`stem`    | `main.script`                     | The base without the extension.

A format extension depends on the available formats. See
[Formats](#user-content-formats) for more information.

## `sym` library
The `sym` library contains **Symbol** values. A symbol is a unique identifier
that can be used to access certain metadata fields of an Instance.

An instance can be indexed with a symbol to get a metadata value in the same way
it can be indexed with a string to get a property value:

```lua
instance = Instance.new("Workspace")
instance[sym.IsService] = true
print(instance[sym.IsService]) --> true
```

The following symbols are defined:

Symbol      | Description
------------|------------
`Desc`      | Gets the inherited descriptor of an instance.
`IsService` | Determines whether an instance is a service.
`RawDesc`   | Accesses the direct director of an instance.
`Reference` | Determines the value used to identify the instance.

#### `Instance[sym.Desc]: RootDesc | nil`
Desc is the descriptor being used by the instance. Descriptors are inherited; if
the instance has no descriptor, then each ancestor of the instance is searched
until a descriptor is found. If none are still found, then the global descriptor
is returned. If there is no global descriptor, then nil is returned.

Getting Desc will return either a RootDesc, or nil.

When setting Desc, the value can be a RootDesc, false, or nil. Setting to Desc
sets the descriptor only for the current instance.

- Setting to a RootDesc will set the descriptor directly for the current
  instance, which may be inherited.
- Setting to nil will cause the instance to have no direct descriptor, and the
  descriptor will be inherited.
- Setting to false will "block", forcing the instance to have no descriptor.
  This behaves sort of like a RootDesc that is empty; there is no descriptor,
  but this state will not inherit, and can be inherited.

#### `Instance[sym.IsService]: bool`
IsService indicates whether the instance is a service, such as Workspace or
Lighting. This is used by some formats to determine how to encode and decode the
instance.

#### `Instance[sym.RawDesc]: RootDesc | bool | nil`
RawDesc is similar to Desc, except that it considers only the direct descriptor
of the current instance.

Getting RawDesc will return a RootDesc if the instance has a descriptor
assigned, false if the descriptor is blocked, or nil if no descriptor is
assigned.

Setting RawDesc behaves the same as setting Desc.

#### `Instance[sym.Reference]: string`
Reference is a string used to refer to the instance from within a DataModel.
Certain formats use this to encode a reference to an instance. For example, the
RBXMX format will generate random UUIDs for its references (e.g.
"RBX8B658F72923F487FAE2F7437482EF16D").

## `types` library
The `types` library contains functions for constructing explicit primitives. The
name of a function corresponds directly to the type. See [Explicit
primitives](#user-content-explicit-primitives) for more information.

Type              | Primitive
------------------|----------
`BinaryString`    | string
`Color3uint8`     | Color3
`Content`         | string
`float`           | number
`int64`           | number
`int`             | number
`ProtectedString` | string
`SharedString`    | string
`token`           | number

# Instances
A major difference between Roblox and rbxmk is what an instance represents. In
Roblox, an instance is a live object that acts and reacts. In rbxmk, an instance
represents *data*, and only data.

Consider the RBXL file format. Files of this format contain information used to
reconstruct the instances that make up a place or model. Such files are static:
they contain only data, but are difficult to manipulate in-place. Instances in
rbxmk are like this, except that they are also interactive: the user can freely
modify data and move it around.

Because of this, there are several differences between the Roblox API and the
rbxmk API. By default, any kind of class can be created. Instances are just
data, including the class name. The ClassName property can even be assigned to.

```lua
local foobar = Instance.new("Foobar")
foobar.ClassName = "FizzBuzz" -- allowed
```

Instances also have no defined properties by default. A value of any type can be
assigned to any property. Likewise, properties that are not explicitly assigned
effectively do not exist.

```lua
local part = Instance.new("Part")
part.Foobar = 42 -- allowed
print(part.Position) --> nil
```

That said, even though it is possible for rbxmk to create arbitrary classes with
arbitrary properties, this does not mean such instances will be interpreted in
any meaningful way when sent over to Roblox. The most convenient way to enforce
the correctness of instances is to use [Descriptors](#user-content-descriptors).

# Descriptors
By default, rbxmk has no knowledge of the classes, members, and enums that are
defined by Roblox. As such, instances can be of any class, properties can be of
any type, and there are no constant enum values. By not explicitly requiring
such information, rbxmk can remain relatively forward-compatible with future
updates to Roblox. It also allows the user to construct values outside the
constraints of the Roblox API.

However, most of the time, the user will be using rbxmk to manipulate values
specifically for Roblox. It's often less convenient for the user to specify type
information manually; the API is slightly off from that of Roblox, therefore
being less familiar. It would be great if this type information could be defined
and enforced automatically.

The solution to this is **descriptors**. A descriptor contains information about
what classes exist, the properties that exist on each class, what enums are
defined, and so on.

The primary descriptor type is the **RootDesc**. This contains a complete
description of the classes and enums of an entire API.

Each instance can have a RootDesc assigned to it. This state is inherited by any
descendant instances. See [`rbxmk.meta`](#user-content-TODO) for more
information on how to assign descriptors to an instance.

Additionally, the `rbxmk.desc` field may be used to apply a RootDesc globally.
When `rbxmk.desc` is set, any instance that wouldn't otherwise inherit a
descriptor will use this global descriptor.

When an instance has a descriptor, several behaviors are enforced:

- When the global descriptor is set, `Instance.new` errors if the given class
  name does not exist (`Instance.new` can also receive a descriptor).
- A property will throw an error if it does not exist for the class.
- Getting an uninitialized property will throw an error.
- Getting a property that currently has an incorrect type will throw an error.
- Setting a property to a value of the incorrect type will throw an error.
- A property of the "Class" type category will throw an error if the assigned
  value is not an instance of the expected class.
- The value assigned to a property of the "Enum" type category will be coerced
  into a token. The value can be an enum item of the expected enum, or a number
  or string of the correct value.
- The class of an instance created from `DataModel:GetService()` must have the
  "Service" tag.

## Types
Descriptors are first-class values like any other, and can be modified on the
fly. There are a number of descriptor types, each with their own fields. See
[`rbxmk.newDesc`](#user-content-TODO) for creating descriptors.

Type          | Description
--------------|------------
RootDesc      | Describes an entire API.
ClassDesc     | Describes a class.
PropertyDesc  | Describes a property member.
FunctionDesc  | Describes a function member.
EventDesc     | Describes a event member.
CallbackDesc  | Describes a callback member.
ParameterDesc | Describes a parameter of a function, event, or callback. Immutable.
TypeDesc      | Describes a type. Immutable.
EnumDesc      | Describes an enum.
EnumItemDesc  | Describes an enum item.

### RootDesc
RootDesc describes an entire API. It has the following members:

Member      | Type
------------|-----
Class       | method
Classes     | method
AddClass    | method
RemoveClass | method
Enum        | method
Enums       | method
AddEnum     | method
RemoveEnum  | method
EnumTypes   | method

#### `RootDesc:Class(name: string): ClassDesc`
Class returns a ClassDesc for the given name, or nil if no such class exists.

#### `RootDesc:Classes(): Array<ClassDesc>`
Classes returns a list of all the classes defined in the RootDesc.

#### `RootDesc:AddClass(class: ClassDesc): bool`
AddClass adds a new class to the RootDesc, returning whether the class was added
successfully. The class will fail to be added if a class of the same name
already exists.

#### `RootDesc:RemoveClass(name: string): bool`
RemoveClass removes a class from the RootDesc, returning whether the class was
removed successfully. False will be returned if a class of the given name does
not exist.

#### `RootDesc:Enum(name: string): EnumDesc`
Enum returns an EnumDesc for the given name, or nil if no such enum exists.

#### `RootDesc:Enums(): Array<EnumDesc>`
Enums returns a list of all the enums defined in the RootDesc.

#### `RootDesc:AddEnum(enum: EnumDesc): bool`
AddEnum adds a new enum to the RootDesc, returning whether the enum was added
successfully. The enum will fail to be added if an enum of the same name already
exists.

#### `RootDesc:RemoveEnum(name: string): bool`
RemoveEnum removes an enum from the RootDesc, returning whether the enum was
removed successfully. False will be returned if an enum of the given name does
not exist.

#### `RootDesc:EnumTypes(): Enums`
EnumTypes returns a set of enum values generated from the current state of the
RootDesc. These enums are associated with the RootDesc, and may be used by
certain properties, so it is important to generate them before operating on such
properties. Additionally, EnumTypes should be called after modifying enum and
enum item descriptors, to regenerate the enum values.

The API of the resulting enums matches that of Roblox's Enums type. A common
pattern is to assign the result of EnumTypes to the "Enum" variable, which
matches Roblox's API:

```lua
Enum = rootDesc:EnumTypes()
print(Enum.NormalId.Front)
```

### ClassDesc
ClassDesc describes a class. It has the following members:

Member         | Type
---------------|-----
Name           | string
Superclass     | string
MemoryCategory | string
Member         | method
Members        | method
AddMember      | method
RemoveMember   | method
Tag            | method
Tags           | method
SetTag         | method
UnsetTag       | method

#### `ClassDesc.Name: string`
Name is the name of the class.

#### `ClassDesc.Superclass: string`
Superclass is the name of the class from which the current class inherits.

#### `ClassDesc.MemoryCategory: string`
MemoryCategory describes the category of the class.

#### `ClassDesc:Member(name: string): MemberDesc`
Member returns a MemberDesc for the given name, or nil of no such member exists.

MemberDesc is any one of the PropertyDesc, FunctionDesc, EventDesc, or
CallbackDesc types.

#### `ClassDesc:Members(): Array<MemberDesc>`
Members returns a list of all the members of the class.

#### `ClassDesc:AddMember(member: MemberDesc): bool`
AddMember adds a new member to the ClassDesc, returning whether the member was
added successfully. The member will fail to be added if a member of the same
name already exists.

#### `ClassDesc:RemoveMember(name: string): bool`
RemoveMember removes a member from the ClassDesc, returning whether the member
was removed successfully. False will be returned if a member of the given name
does not exist.

#### `ClassDesc:Tag(name: string): bool`
Tag returns whether a tag of the given name is set on the descriptor.

#### `ClassDesc:Tags(): Array<string>`
Tags returns a list of tags that are set on the descriptor.

#### `ClassDesc:SetTag(tags: ...string)`
SetTags sets the given tags on the descriptor.

#### `ClassDesc:UnsetTag(tags: ...string)`
SetTags unsets the given tags on the descriptor.

### PropertyDesc
PropertyDesc describes a property member of a class. It has the following
members:

Member        | Type
--------------|-----
Name          | string
ValueType     | TypeDesc
ReadSecurity  | string
WriteSecurity | string
CanLoad       | bool
CanSave       | bool
Tag           | method
Tags          | method
SetTag        | method
UnsetTag      | method

#### `PropertyDesc.Name: string`
Name is the name of the member.

#### `PropertyDesc.ValueType: TypeDesc`
ValueType is the value type of the property.

#### `PropertyDesc.ReadSecurity: string`
ReadSecurity indicates the security context required to get the property.

#### `PropertyDesc.WriteSecurity: string`
WriteSecurity indicates the security context required to set the property.

#### `PropertyDesc.CanLoad: bool`
CanLoad indicates whether the property is deserialized when decoding from a file.

#### `PropertyDesc.CanSave: bool`
CanLoad indicates whether the property is serialized when encoding to a file.

#### `PropertyDesc:Tag(name: string): bool`
Tag returns whether a tag of the given name is set on the descriptor.

#### `PropertyDesc:Tags(): Array<string>`
Tags returns a list of tags that are set on the descriptor.

#### `PropertyDesc:SetTag(tags: ...string)`
SetTags sets the given tags on the descriptor.

#### `PropertyDesc:UnsetTag(tags: ...string)`
SetTags unsets the given tags on the descriptor.

### FunctionDesc
FunctionDesc describes a function member of a class. It has the following
members:

Member        | Type
--------------|-----
Name          | string
Parameters    | method
SetParameters | method
ReturnType    | TypeDesc
Security      | string
Tag           | method
Tags          | method
SetTag        | method
UnsetTag      | method

#### `FunctionDesc.Name: string`
Name is the name of the member.

#### `FunctionDesc:Parameters(): Array<ParameterDesc>`
Parameters returns a list of parameters of the function.

#### `FunctionDesc:SetParameters(params: Array<ParameterDesc>)`
SetParameters sets the parameters of the function.

#### `FunctionDesc.ReturnType: TypeDesc`
ReturnType is the type returned by the function.

#### `FunctionDesc.Security: string`
Security indicates the security content required to index the member.

#### `FunctionDesc:Tag(name: string): bool`
Tag returns whether a tag of the given name is set on the descriptor.

#### `FunctionDesc:Tags(): Array<string>`
Tags returns a list of tags that are set on the descriptor.

#### `FunctionDesc:SetTag(tags: ...string)`
SetTags sets the given tags on the descriptor.

#### `FunctionDesc:UnsetTag(tags: ...string)`
SetTags unsets the given tags on the descriptor.

### EventDesc
EventDesc describes an event member of a class. It has the following members:

Member        | Type
--------------|-----
Name          | string
Parameters    | method
SetParameters | method
Security      | string
Tag           | method
Tags          | method
SetTag        | method
UnsetTag      | method

#### `EventDesc.Name: string`
Name is the name of the member.

#### `EventDesc:Parameters(): Array<ParameterDesc>`
Parameters returns a list of parameters of the event.

#### `EventDesc:SetParameters(params: Array<ParameterDesc>)`
SetParameters sets the parameters of the event.

#### `EventDesc.Security: string`
Security indicates the security content required to index the member.

#### `EventDesc:Tag(name: string): bool`
Tag returns whether a tag of the given name is set on the descriptor.

#### `EventDesc:Tags(): Array<string>`
Tags returns a list of tags that are set on the descriptor.

#### `EventDesc:SetTag(tags: ...string)`
SetTags sets the given tags on the descriptor.

#### `EventDesc:UnsetTag(tags: ...string)`
SetTags unsets the given tags on the descriptor.

### CallbackDesc
CallbackDesc describes a callback member of a class. It has the following
members:

Member        | Type
--------------|-----
Name          | string
Parameters    | method
SetParameters | method
ReturnType    | TypeDesc
Security      | string
Tag           | method
Tags          | method
SetTag        | method
UnsetTag      | method

#### `CallbackDesc.Name: string`
Name is the name of the member.

#### `CallbackDesc:Parameters(): Array<ParameterDesc>`
Parameters returns a list of parameters of the callback.

#### `CallbackDesc:SetParameters(params: Array<ParameterDesc>)`
SetParameters sets the parameters of the callback.

#### `CallbackDesc.ReturnType: TypeDesc`
ReturnType is the type returned by the callback.

#### `CallbackDesc.Security: string`
Security indicates the security content required to set the member.

#### `CallbackDesc:Tag(name: string): bool`
Tag returns whether a tag of the given name is set on the descriptor.

#### `CallbackDesc:Tags(): Array<string>`
Tags returns a list of tags that are set on the descriptor.

#### `CallbackDesc:SetTag(tags: ...string)`
SetTags sets the given tags on the descriptor.

#### `CallbackDesc:UnsetTag(tags: ...string)`
SetTags unsets the given tags on the descriptor.

### ParameterDesc
ParameterDesc describes a parameter of a function, event or callback member. It
has the following immutable members:

Member  | Type
--------|-----
Type    | TypeDesc
Name    | string
Default | string?

ParameterDesc is immutable. A new value with different fields can be created
with rbxmk.newDesc.

#### `ParameterDesc.Type: TypeDesc`
Type is the type of the parameter.

#### `ParameterDesc.Name: string`
Name is a name describing the parameter.

#### `ParameterDesc.Default: string?`
Default is a string describing the default value of the parameter. May also be
nil, indicating that the parameter has no default value.

### TypeDesc
TypeDesc describes a value type. It has the following immutable members:

Member   | Type
---------|-----
Category | string
Name     | string

TypeDesc is immutable. A new value with different fields can be created with
rbxmk.newDesc.

#### `TypeDesc.Category: string`
Category is the category of the type. Certain categories are treated specially:

- `Class`: Name is the name of a class. A value of the type is expected to be an
  Instance of the class.
- `Enum`: Name is the name of an enum. A value of the type is expected to be an
  enum item of the enum.

#### `TypeDesc.Name: string`
Name is the name of the type.

### EnumDesc
EnumDesc describes an enum. It has the following members:

Member     | Type
-----------|-----
Name       | string
Item       | method
Items      | method
AddItem    | method
RemoveItem | method
Tag        | method
Tags       | method
SetTag     | method
UnsetTag   | method

#### `EnumDesc.Name: string`
Name is the name of the enum.

#### `EnumDesc:Item(name: string): EnumItemDesc`
Item returns an EnumItemDesc for the given name, or nil of no such item exists.

#### `EnumDesc:Items(): Array<EnumItemDesc>`
Items returns a list of all the items of the enum.

#### `EnumDesc:AddItem(item: EnumItemDesc): bool`
AddItem adds a new item to the EnumDesc, returning whether the item was added
successfully. The item will fail to be added if an item of the same name already
exists.

#### `EnumDesc:RemoveItem(name: string): bool`
RemoveItem removes an item from the EnumDesc, returning whether the item was
removed successfully. False will be returned if an item of the given name does
not exist.

#### `EnumDesc:Tag(name: string): bool`
Tag returns whether a tag of the given name is set on the descriptor.

#### `EnumDesc:Tags(): Array<string>`
Tags returns a list of tags that are set on the descriptor.

#### `EnumDesc:SetTag(tags: ...string)`
SetTags sets the given tags on the descriptor.

#### `EnumDesc:UnsetTag(tags: ...string)`
SetTags unsets the given tags on the descriptor.

### EnumItemDesc
EnumDesc describes an enum item. It has the following members:

Member   | Type
---------|-----
Name     | string
Value    | int
Index    | int
Tag      | method
Tags     | method
SetTag   | method
UnsetTag | method

#### `EnumItemDesc.Name: string`
Name is the name of the enum item.

#### `EnumItemDesc.Value: int`
Value is the numeric value of the enum item.

#### `EnumItemDesc.Index: int`
Index is an integer that hints the order of the enum item.

#### `EnumItemDesc:Tag(name: string): bool`
Tag returns whether a tag of the given name is set on the descriptor.

#### `EnumItemDesc:Tags(): Array<string>`
Tags returns a list of tags that are set on the descriptor.

#### `EnumItemDesc:SetTag(tags: ...string)`
SetTags sets the given tags on the descriptor.

#### `EnumItemDesc:UnsetTag(tags: ...string)`
SetTags unsets the given tags on the descriptor.

## Diffing and Patching
Descriptors can be compared and patched with the `rbxmk.diffDesc` and
`rbxmk.patchDesc` functions. `diffDesc` returns a list of **DescActions**, which
describe how to transform the first descriptor into the second. `patchDesc` can
used to apply this transformation.

```lua
-- List differences.
local diff = rbxmk.diffDesc(prevDesc, nextDesc)
-- Transform prev into next.
rbxmk.patchDesc(prevDesc, diff)
```

Patching is used primarily to augment some pregenerated descriptor with elements
that aren't present. For example, Roblox's API dump does not include the
`BinaryStringValue.Value` member. This can be patched with an action that adds
the member, allowing it to be used as normal.

More generally, patching allows any version of an API dump to be used, while
maintaining a separate, constant set of additional API elements.

### DescAction
DescAction describes a single action that transforms a descriptor.

Currently, DescAction has no members. However, converting a DescAction to a
string will display the content of the action in a human-readable format.

Actions may also be created directly with the
[`desc-patch.json`](#user-content-TODO) format.

# Explicit primitives
The properties of instances in Roblox have a number of different types. Many of
these types can be expressed in Lua through constructors. Examples of such are
`CFrame`, `Vector3`, `UDim2`, and so on. These types correspond to internal data
types within the Roblox engine. The Lua representation of, say, a `CFrame`, is a
userdata with accessible fields.

Some Roblox types are represented with a simple Lua primitive, such as a number
or string. For example, the Roblox types `int`, `int64`, `float`, and `double`
all map to Lua's `number` type. When setting a property, the engine is able to
reflect this Lua `number` back to the correct Roblox type, because the property
has a descriptor that includes the property's type.

In rbxmk, when an instance has a descriptor, it is able to make this conversion
as expected. However, the user may not always have access to descriptors. When
no descriptors are specified, properties have no types. For example, when a
property is set to a Lua number, it is always converted into a `double`. In the
absence of extra type information, the user needs some way to set specific
Roblox types.

This problem is solved with "explicit primitives", or **exprims**. An exprim is
a userdata representation of an otherwise ambiguous type. This userdata carries
type metadata along with a given value, allowing the value to be mapped to the
correct Roblox type when it is set as a property.

The `types` library contains functions for constructing exprim types. See the
[types library](#user-content-types-library) section for a list of exprims.

	-- Problem
	local v = Instance.new("IntValue")
	v.Name = "Value"
	v.Value = 42 -- type is `double`; not correct for IntValue.

	-- Solution
	local v = Instance.new("IntValue")
	v.Name = "Value"
	v.Value = types.int64(42) -- Type is int64.

The default Roblox type that maps to Lua strings is `string`. As such, `string`
has no exprim. Likewise, the default type that maps to Lua numbers is `double`,
so it also has no exprim.

An exprim userdata has no fields or operators other than the `Value` field,
which returns the underlying primitive value. Exprims are meant to be
short-lived, and shouldn't really be used beyond getting or setting a property
in the absence of descriptors. When possible, descriptors should be utilized
instead.

# Sources
A **source** is an external location from which raw data can be read from and
written to. A source can be accessed at a low level through the
`rbxmk.readSource` and `rbxmk.writeSource` functions.

A source usually has a corresponding library that provides convenient access for
common cases.

## `file` source
The `file` source provides access to the file system.

### Reading
### Writing
### `file` library
The `file` library handles the `file` source.

Name    | Description
--------|------------
`read`  | Reads data from a file in a certain format.
`write` | Writes data to a file in a certain format.

# Formats
A **format** is capable of encoding a value to raw bytes, or decoding raw bytes
into a value. A format can be accessed at a low level through the
`rbxmk.encodeFormat` and `rbxmk.decodeFormat` functions.

The name of a format corresponds to the extension of a file name. For example,
the `lua` format corresponds to the `.lua` file extension. When determining a
format from a file extension, format names are greedy; if a file extension is
`.server.lua`, this will select the `server.lua` format before the `lua` format.
For convenience, a format name may have an optional leading `.` separator.

## String formats
Several string formats are defined for encoding string-like values.

Format | Name   | Description
-------|--------|------------
`txt`  | Text   | Encodes string-like values to UTF-8 text.
`bin`  | Binary | Encodes string-like values to raw bytes.

## Lua formats
Several formats are defined for decoding Lua files into script instances.

Format             | Description
-------------------|------------
`modulescript.lua` | Decodes into a ModuleScript instance.
`script.lua`       | Decodes into a Script instance.
`localscript.lua`  | Decodes into a LocalScript instance.
`lua`              | Alias for `modulescript.lua`.
`server.lua`       | Alias for `script.lua`.
`client.lua`       | Alias for `localscript.lua`.

## Roblox formats
Several formats are defined for serializing instances.

Format  | Description
--------|------------
`rbxl`  | The Roblox binary place format.
`rbxm`  | The Roblox binary model format.
`rbxlx` | The Roblox XML place format.
`rbxmx` | The Roblox XML model format.

## Descriptor formats
Several formats are defined for encoding descriptors.

Format            | Description
------------------|------------
`desc.json`       | Descriptors in JSON format. More commonly known as an "API dump".
`desc-patch.json` | Actions that describe changes to descriptors, in JSON format.
