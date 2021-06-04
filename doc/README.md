# rbxmk reference
This document contains a complete reference of the rbxmk API, and provides
details on how rbxmk works.

<table>
<thead><tr><th>Table of Contents</th></tr></thead>
<tbody><tr><td>

1. [Documents][documents]
2. [Command line][command-line]
	1. [Run command][run-command]
	2. [Help command][help-command]
	3. [Version command][version-command]
	4. [Interactive command][interactive-command]
	5. [Download asset command][download-asset-command]
	6. [Upload asset command][upload-asset-command]
	7. [Dump command][dump-command]
3. [Instances][instances]
4. [Attributes][attributes]
5. [Descriptors][descriptors]
	1. [Descriptor types][descriptor-types]
	2. [Diffing and Patching][diffing-and-patching]
6. [Value inheritance][value-inheritance]
	1. [Indexing][indexing]
	2. [Raw member][raw-member]
	3. [Global field][global-field]
7. [Explicit primitives][explicit-primitives]
8. [File access limits][file-access-limits]

</td></tr></tbody>
</table>

This reference uses [Luau][luau] type annotation syntax to describe the API of
an element. Some liberties are taken for patterns not supported by the Luau
syntax. For example, `...` indicates variable parameters.

[luau]: https://roblox.github.io/luau/

## Documents
[documents]: #user-content-documents

This reference is divided into a number of documents.

Document                     | Description
-----------------------------|------------
[README.md](README.md)       | Describes abstract concepts present in rbxmk (this document).
[formats.md](formats.md)     | Lists formats used to encode and decode data.
[libraries.md](libraries.md) | Lists the libraries provided by the rbxmk environment.
[types.md](types.md)         | Lists the data types present throughout the rbxmk environment.
[enums.md](enums.md)         | Lists the enums defined within the rbxmk environment.

## Command line
[command-line]: #user-content-command-line

rbxmk is a single executable, to be run within a [command-line interface][cli].

```bash
rbxmk COMMAND [ OPTIONS... ]
```

The rbxmk command receives the name of a subcommand followed by a number of
options, which depend on the subcommand. The following subcommands are provided:

Subcommand                                       | Description
-------------------------------------------------|------------
[`rbxmk run`][run-command]                       | Executes a Lua script.
[`rbxmk help`][help-command]                     | Displays help for rbxmk.
[`rbxmk version`][version-command]               | Displays the version of rbxmk.
[`rbxmk i`][interactive-command]                 | Enters interactive mode.
[`rbxmk download-asset`][download-asset-command] | Downloads a Roblox asset.
[`rbxmk upload-asset`][upload-asset-command]     | Uploads a Roblox asset.
[`rbxmk dump`][dump-command]                     | Dumps the rbxmk Lua API.

[cli]: https://en.wikipedia.org/wiki/Command-line_interface

### Run command
[run-command]: #user-content-run-command

```bash
rbxmk run FILE [ VALUE... ]
```

The **run** command receives a path to a file to be executed as a Lua script.

```bash
rbxmk run script.lua
```

If `-` is given, then the script will be read from stdin instead.

```bash
echo 'print("hello world!")' | rbxmk run -
```

The remaining arguments are Lua values to be passed to the file. Numbers, bools,
and nil are parsed into their respective types in Lua, and any other value is
interpreted as a string.

```bash
rbxmk run script.lua true 3.14159 hello!
```

Within the script, these arguments can be received from the `...` operator:

```lua
local arg1, arg2, arg3 = ...
```

For more information about the Lua environment provided by rbxmk, refer to the
[Libraries section](libraries.md).

### Help command
[help-command]: #user-content-help-command

```bash
rbxmk help [ COMMAND ]
```

The **help** command displays information about a subcommand. If no subcommand
is specified, information about using rbxmk is displayed.

### Version command
[version-command]: #user-content-version-command

```bash
rbxmk version
```

The **version** command displays the version of the rbxmk command. The result is
a string formatted according to [semantic versioning](https://semver.org/).

### Interactive command
[interactive-command]: #user-content-interactive-command

```bash
rbxmk i
```

The *i* command enters interactive mode. Each prompt executes a chunk of Lua
code.

If a prompt begins with `=`, then the comma-separated list of expressions that
follow are evaluated and printed to standard output.

The environment contains the `os.exit` function. When called, interactive mode
is terminated, and the program exits.

Within supported terminals, the following shortcuts are available:

Shortcut                  | Description
--------------------------|------------
`Ctrl+A`, `Home`          | Move cursor to beginning of line.
`Ctrl+E`, `End`           | Move cursor to end of line
`Ctrl+B`, `Left`          | Move cursor one character left.
`Ctrl+F`, `Right`         | Move cursor one character right.
`Ctrl+Left`, `Alt+B`      | Move cursor to previous word.
`Ctrl+Right`, `Alt+F`     | Move cursor to next word
`Ctrl+D`, `Del`           | If line is not empty, delete character under cursor.
`Ctrl+D`                  | If line is empty, end of file.
`Ctrl+C`                  | Reset input (create new empty prompt).
`Ctrl+L`                  | Clear screen (line is unmodified).
`Ctrl+T`                  | Transpose previous character with current character.
`Ctrl+H`, `BackSpace`     | Delete character before cursor.
`Ctrl+W`, `Alt+BackSpace` | Delete word leading up to cursor.
`Alt+D`                   | Delete word following cursor.
`Ctrl+K`                  | Delete from cursor to end of line.
`Ctrl+U`                  | Delete from start of line to cursor.
`Ctrl+P`, `Up`            | Previous match from history.
`Ctrl+N`, `Down`          | Next match from history.
`Ctrl+R`                  | Reverse Search history (`Ctrl+S` forward, `Ctrl+G` cancel).
`Ctrl+Y`                  | Paste from Yank buffer (`Alt+Y` to paste next yank instead).
`Tab`                     | Next completion.
`Shift+Tab`               | (after `Tab`) Previous completion.

### Download asset command
[download-asset-command]: #user-content-download-asset-command

```bash
rbxmk download-asset [ FLAGS ] -id INT [ PATH ]
```

The **download-asset** command downloads an asset from the roblox website.

The `-id` flag, which is required, specifies the ID of the asset to download.

The first non-flag argument is the path to a file to write to. If not specified,
then the file will be written to standard output.

### Upload asset command
[upload-asset-command]: #user-content-upload-asset-command

```bash
rbxmk upload-asset [ FLAGS ] [ -id INT ] PATH
```

Uploads an asset to the roblox website.

The `-id` flag specifies the ID of the asset to upload.

The first non-flag argument is the path to a file to read from, which is
required. If the path is `-`, then the file will be read from standard input.

### Dump command
[dump-command]: #user-content-dump-command

```bash
rbxmk dump FORMAT
```

Dumps the API of the rbxmk Lua environment. The following formats are supported:

Format   | Description
---------|------------
json     | General JSON format.
json-min | Minified JSON format.
selene   | [Selene][selene] TOML format.

[selene]: https://kampfkarren.github.io/selene/

## Instances
[instances]: #user-content-instances

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
the correctness of instances is to use [descriptors][descriptors].

Another notable difference is that children cannot be indexed by name directly.
Without a descriptor, *all* properties are considered valid, so no there would
be no room to interpret an index as a child.

More generally, child indexing has been proven to have poor forward
compatibility. New properties are added to the API all the time, and every such
change has the potential to cause a script to break, because the script expected
a child and got a new property instead.

rbxmk moves past this problem by simply not implementing child indexing.
Instead, the [Instance.Descend][Instance.Descend] method is introduced to
provide a convenient and safe alternative.

## Attributes
[Attributes]: #user-content-attributes

Instances in Roblox and rbxmk have **attributes**, which are similar to custom
properties.

Roblox serializes all attributes into a single property in a binary format. In
rbxmk, this format is implemented by the [rbxattr
format](formats.md#user-content-rbxattr).

rbxmk provides the same API as Roblox for manipulating attributes:

- [Instance.GetAttribute](types.md#user-content-instancegetattribute)
- [Instance.GetAttributes](types.md#user-content-instancegetattributes)
- [Instance.SetAttribute](types.md#user-content-instancesetattribute)

Additionally, rbxmk provides the
[SetAttributes](types.md#user-content-instancesetattributes) method for setting
all the attributes of an instance more efficiently.

In order to maintain rbxmk's theme of forward-compatibility, rbxmk provides the
[AttrConfig][AttrConfig] type to configure how attributes are applied to an
instance. AttrConfigs are inherited, the behavior of which is described in the
[Value inheritance][value-inheritance] section.

## Descriptors
[descriptors]: #user-content-descriptors

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

The primary descriptor type is the [**RootDesc**][RootDesc]. This contains a
complete description of the classes and enums of an entire API.

An [Instance][Instance] can have a RootDesc assigned to it. This state is
inherited by any descendant instances. See [sym.Desc][Instance.sym.Desc] for
more information.

Additionally, the [rbxmk.globalDesc][rbxmk.globalDesc] field may be used to
apply a RootDesc globally. When globalDesc is set, any instance that wouldn't
otherwise inherit a descriptor will use this global descriptor.

When an instance has a descriptor, several behaviors are enforced:

- When the global descriptor is set,
  [Instance.new](types.md#user-content-instancenew) errors if the given class
  name does not exist (Instance.new can also receive a descriptor).
- A property will throw an error if it does not exist for the class.
- Getting an uninitialized property will throw an error.
- Getting a property that currently has an incorrect type will throw an error.
- Setting a property to a value of the incorrect type will throw an error.
- A property of the "Class" type category will throw an error if the assigned
  value is not an instance of the expected class.
- The value assigned to a property of the "Enum" type category will be coerced
  into a token. The value can be an enum item of the expected enum, or a number
  or string of the correct value.
- The class of an instance created from
  [DataModel.GetService](types.md#user-content-datamodelgetservice) must have
  the "Service" tag.

## Descriptor types
[descriptor-types]: #user-content-descriptor-types

Descriptors are first-class values like any other, and can be modified on the
fly. There are a number of descriptor types, each with their own fields.

Type                           | Description
-------------------------------|------------
[RootDesc][RootDesc]           | Describes an entire API.
[ClassDesc][ClassDesc]         | Describes a class.
[PropertyDesc][PropertyDesc]   | Describes a property member.
[FunctionDesc][FunctionDesc]   | Describes a function member.
[EventDesc][EventDesc]         | Describes an event member.
[CallbackDesc][CallbackDesc]   | Describes a callback member.
[ParameterDesc][ParameterDesc] | Describes a parameter of a function, event, or callback. Immutable.
[TypeDesc][TypeDesc]           | Describes a type. Immutable.
[EnumDesc][EnumDesc]           | Describes an enum.
[EnumItemDesc][EnumItemDesc]   | Describes an enum item.

## Diffing and Patching
[diffing-and-patching]: #user-content-diffing-and-patching

Descriptors can be compared and patched with the
[rbxmk.diffDesc](libraries.md#user-content-rbxmkdiffdesc) and
[rbxmk.patchDesc](libraries.md#user-content-rbxmkpatchdesc) functions. diffDesc
returns a list of [**DescActions**](types.md#user-content-descaction), which
describe how to transform the first descriptor into the second. patchDesc can
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

## Value inheritance
[value-inheritance]: #user-content-value-inheritance

Certain symbol fields on [Instances][Instance] have an inheritance behavior.

Member                                    | Principal type           | Raw member                                      | Global field
------------------------------------------|--------------------------|-------------------------------------------------|-------------
[sym.AttrConfig][Instance.sym.AttrConfig] | [AttrConfig][AttrConfig] | [sym.RawAttrConfig][Instance.sym.RawAttrConfig] | [rbxmk.globalAttrConfig][rbxmk.globalAttrConfig]
[sym.Desc][Instance.sym.Desc]             | [RootDesc][RootDesc]     | [sym.RawDesc][Instance.sym.RawDesc]             | [rbxmk.globalDesc][rbxmk.globalDesc]

The following sections describe the aspects of this behavior for each member.

### Indexing
[indexing]: #user-content-indexing

The member has a **principal type**, which indicates the type of the main value
assigned to the member.

Getting the member will return either a value of the principal type, or nil. If
the instance has no value, then each ancestor of the instance is searched until
a value is found. If none are still found, then the global value is returned. If
there is no global value, then nil is finally returned.

When setting the member, the value can be of the principal type, false, or nil.
Setting to the member sets the value only for the current instance.

- Setting to a value of the principal type will set the value directly for the
  current instance, which may be inherited.
- Setting to nil will cause the instance to have no direct value, and the value
  will be inherited instead.
- Setting to false will "block" inheritance, forcing the instance to have no
  value. This behaves sort of like a value that is empty; the instance wont
  inherit any other values, and this blocked state can be inherited.

### Raw member
[raw-member]: #user-content-raw-member

The member has a corresponding **raw member**, which gets the value directly.
Getting the raw member will return the value if the instance has a value
assigned, false if the member is blocked, or nil if no value is assigned.
Setting the raw member behaves the same as setting the corresponding member.

### Global field
[global-field]: #user-content-global-field

The member has a corresponding **global field** in the [rbxmk
library](libraries.md#user-content-rbxmk), which sets a global value to be
applied to any instance that would otherwise inherit nothing.

## Explicit primitives
[explicit-primitives]: #user-content-explicit-primitives

The properties of instances in Roblox have a number of different types. Many of
these types can be expressed in Lua through constructors. Examples of such are
CFrame, Vector3, UDim2, and so on. These types correspond to internal data types
within the Roblox engine. The Lua representation of, say, a CFrame, is a
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

This problem can be solved with "explicit primitives", or **exprims**. An exprim
is a userdata representation of an otherwise ambiguous type. This userdata
carries type metadata along with a given value, allowing the value to be mapped
to the correct Roblox type when it is set as a property.

The [types library](libraries.md#user-content-types) contains a constructor
function for each exprim type.

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

An exprim userdata has no fields or operators other than the "Value" field,
which returns the underlying primitive value:

	v.Value = types.int64(v.Value.Value + 1)

Exprims are meant to be short-lived, and shouldn't really be used beyond getting
or setting a property in the absence of [descriptors][descriptors]. When
possible, descriptors should be utilized instead.

## File access limits
[file-access-limits]: #user-content-file-access-limits

To reduce the impact of malicious scripts, rbxmk limits a script's access to the
file system. An environment specifies a number of **root** directories. Only
file paths within a root can be accessed. A root path itself cannot be accessed,
except for moving files into ([fs.rename][fs.rename]), or getting the contents
of ([fs.dir][fs.dir]).

The following directories are marked as roots:
- The working directory (<code>[os.expand][os.expand]("$wd")</code>).
- The directory of the first running script file (<code>[os.expand][os.expand]("$rsd")</code>).
- The temporary directory provided by rbxmk (<code>[os.expand][os.expand]("$tmp")</code>).

[AttrConfig]: types.md#user-content-attrconfig
[CallbackDesc]: types.md#user-content-callbackdesc
[ClassDesc]: types.md#user-content-classdesc
[EnumDesc]: types.md#user-content-enumdesc
[EnumItemDesc]: types.md#user-content-enumitemdesc
[EventDesc]: types.md#user-content-eventdesc
[fs.dir]: libraries.md#user-content-fsdir
[fs.rename]: libraries.md#user-content-fsrename
[FunctionDesc]: types.md#user-content-functiondesc
[Instance.Descend]: libraries.md#user-content-instancedescend
[Instance.sym.AttrConfig]: types.md#user-content-instancesymattrconfig
[Instance.sym.Desc]: types.md#user-content-instancesymdesc
[Instance.sym.RawAttrConfig]: types.md#user-content-instancesymrawattrconfig
[Instance.sym.RawDesc]: types.md#user-content-instancesymrawdesc
[Instance]: types.md#user-content-instance
[os.expand]: libraries.md#user-content-osexpand
[ParameterDesc]: types.md#user-content-parameterdesc
[PropertyDesc]: types.md#user-content-propertydesc
[rbxmk.globalAttrConfig]: libraries.md#user-content-rbxmkglobalattrconfig
[rbxmk.globalDesc]: libraries.md#user-content-rbxmkglobaldesc
[RootDesc]: types.md#user-content-rootdesc
[TypeDesc]: types.md#user-content-typedesc
