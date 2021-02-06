# Libraries
This document contains a reference to the libraries available to rbxmk scripts.

<table>
<thead><tr><th>Table of Contents</th></tr></thead>
<tbody><tr><td>

1. [Base][base]
2. [rbxmk][rbxmk]
	1. [rbxmk.cookiesFrom][rbxmk.cookiesFrom]
	2. [rbxmk.decodeFormat][rbxmk.decodeFormat]
	3. [rbxmk.diffDesc][rbxmk.diffDesc]
	4. [rbxmk.encodeFormat][rbxmk.encodeFormat]
	5. [rbxmk.formatCanDecode][rbxmk.formatCanDecode]
	6. [rbxmk.globalAttrConfig][rbxmk.globalAttrConfig]
	7. [rbxmk.globalDesc][rbxmk.globalDesc]
	8. [rbxmk.loadFile][rbxmk.loadFile]
	9. [rbxmk.loadString][rbxmk.loadString]
	10. [rbxmk.newCookie][rbxmk.newCookie]
	11. [rbxmk.newDesc][rbxmk.newDesc]
	12. [rbxmk.patchDesc][rbxmk.patchDesc]
	13. [rbxmk.runFile][rbxmk.runFile]
	14. [rbxmk.runString][rbxmk.runString]
3. [Roblox][roblox]
4. [math][math]
	1. [math.clamp][math.clamp]
	2. [math.log][math.log]
	3. [math.round][math.round]
	4. [math.sign][math.sign]
5. [os][os]
	1. [os.dir][os.dir]
	2. [os.expand][os.expand]
	3. [os.getenv][os.getenv]
	4. [os.join][os.join]
	5. [os.split][os.split]
	6. [os.stat][os.stat]
6. [string][string]
	1. [string.split][string.split]
7. [table][table]
	1. [table.clear][table.clear]
	2. [table.create][table.create]
	3. [table.find][table.find]
	4. [table.move][table.move]
	5. [table.pack][table.pack]
	6. [table.unpack][table.unpack]
8. [sym][sym]
9. [types][types]

</td></tr></tbody>
</table>

[Lua](https://lua.org/) scripts are used to perform actions in rbxmk. The
environment provided by rbxmk is packaged as a set of libraries. Some libraries
are loaded under a specific name, while others are loaded directly into the
global environment:

Library                  | Description
-------------------------|------------
[(base)][base]           | The Lua 5.1 standard library, abridged.
[rbxmk][rbxmk]           | An interface to the rbxmk engine, and the rbxmk environment.
[(roblox)][roblox]       | An environment emulating the Roblox Lua API.
[math][math]             | Extensions to the standard math library.
[os][os]                 | Extensions to the standard os library.
[string][string]         | Extensions to the standard string library.
[table][table]           | Extensions to the standard table library.
[sym][sym]               | Symbols for accessing instance metadata.
[types][types]           | Fallbacks for constructing certain types.

Additionally, each source provides a library that enables access to the source.
See the [Sources document](sources.md) for more information about each source
library.

The `_RBXMK_VERSION` global variable is defined as a string containing the
current version of rbxmk, formatted according to [semantic
versioning](https://semver.org/).

## Base
[base]: #user-content-base

The **base** library is loaded directly into the global environment. It contains
the following items from the [Lua 5.1 standard
library](https://www.lua.org/manual/5.1/manual.html#5):

- [`_G`](https://www.lua.org/manual/5.1/manual.html#pdf-_G)
- [`_VERSION`](https://www.lua.org/manual/5.1/manual.html#pdf-_VERSION)
- [`assert`](https://www.lua.org/manual/5.1/manual.html#pdf-assert)
- [`error`](https://www.lua.org/manual/5.1/manual.html#pdf-error)
- [`getmetatable`](https://www.lua.org/manual/5.1/manual.html#pdf-getmetatable)
- [`ipairs`](https://www.lua.org/manual/5.1/manual.html#pdf-ipairs)
- [`next`](https://www.lua.org/manual/5.1/manual.html#pdf-next)
- [`pairs`](https://www.lua.org/manual/5.1/manual.html#pdf-pairs)
- [`pcall`](https://www.lua.org/manual/5.1/manual.html#pdf-pcall)
- [`print`](https://www.lua.org/manual/5.1/manual.html#pdf-print)
- [`select`](https://www.lua.org/manual/5.1/manual.html#pdf-select)
- [`setmetatable`](https://www.lua.org/manual/5.1/manual.html#pdf-setmetatable)
- [`tonumber`](https://www.lua.org/manual/5.1/manual.html#pdf-tonumber)
- [`tostring`](https://www.lua.org/manual/5.1/manual.html#pdf-tostring)
- [`type`](https://www.lua.org/manual/5.1/manual.html#pdf-type)
- [`unpack`](https://www.lua.org/manual/5.1/manual.html#pdf-unpack)
- [`xpcall`](https://www.lua.org/manual/5.1/manual.html#pdf-xpcall)
- [`math` library](https://www.lua.org/manual/5.1/manual.html#5.6)
- [`string` library](https://www.lua.org/manual/5.1/manual.html#5.4), except `string.dump`
- [`table` library](https://www.lua.org/manual/5.1/manual.html#5.5)
- [`os.clock`](https://www.lua.org/manual/5.1/manual.html#pdf-os.clock)
- [`os.date`](https://www.lua.org/manual/5.1/manual.html#pdf-os.date)
- [`os.difftime`](https://www.lua.org/manual/5.1/manual.html#pdf-os.difftime)
- [`os.time`](https://www.lua.org/manual/5.1/manual.html#pdf-os.time)

## rbxmk
[rbxmk]: #user-content-rbxmk

The **rbxmk** library contains functions related to the rbxmk engine.

Name                                       | Kind     | Description
-------------------------------------------|----------|------------
[cookiesFrom][rbxmk.cookiesFrom]           | function | Retrieve cookies from a known location.
[decodeFormat][rbxmk.decodeFormat]         | function | Deserialize data from bytes.
[diffDesc][rbxmk.diffDesc]                 | function | Get the differences between two descriptors.
[encodeFormat][rbxmk.encodeFormat]         | function | Serialize data into bytes.
[formatCanDecode][rbxmk.formatCanDecode]   | function | Check whether a format decodes into a type.
[globalAttrConfig][rbxmk.globalAttrConfig] | field    | Get or set the global AttrConfig.
[globalDesc][rbxmk.globalDesc]             | field    | Get or set the global descriptor.
[loadFile][rbxmk.loadFile]                 | function | Load the content of a file as a function.
[loadString][rbxmk.loadString]             | function | Load a string as a function.
[newCookie][rbxmk.newCookie]               | function | Create a new cookie.
[newDesc][rbxmk.newDesc]                   | function | Create a new descriptor.
[patchDesc][rbxmk.patchDesc]               | function | Transform a descriptor by applying differences.
[runFile][rbxmk.runFile]                   | function | Run a file as a Lua chunk.
[runString][rbxmk.runString]               | function | Run a string as a Lua chunk.

### rbxmk.cookiesFrom
[rbxmk.cookiesFrom]: #user-content-rbxmkcookiesfrom
<code>rbxmk.cookiesFrom(location: string): (cookies: [Cookies][Cookies]?)</code>

The **cookiesFrom** function retrieves cookies from a known location. *location*
is case-insensitive.

The following locations are implemented:

Location | Description
---------|------------
`studio` | Returns the cookies used for authentication when logging into Roblox Studio.

Returns nil if no cookies could be retrieved from the location. Throws an error
if an unknown location is given.

### rbxmk.decodeFormat
[rbxmk.decodeFormat]: #user-content-rbxmkdecodeformat
<code>rbxmk.decodeFormat(format: [string](##), bytes: [BinaryString](##)): (value: [any](##))</code>

The **decodeFormat** function decodes *bytes* into a value according to
*format*. The exact details of each format are described in the
[Formats](formats.md) documents.

decodeFormat will throw an error if the format does not exist, or the format has
no decoder defined.

### rbxmk.diffDesc
[rbxmk.diffDesc]: #user-content-rbxmkdiffdesc
<code>rbxmk.diffDesc(prev: [RootDesc][RootDesc]?, next: [RootDesc][RootDesc]?): (diff: {[DescAction][DescAction]})</code>

The **diffDesc** function compares two root descriptors and returns the
differences between them. A nil value for *prev* or *next* is treated the same
as an empty descriptor. The result is a list of actions that describe how to
transform *prev* into *next*.

### rbxmk.encodeFormat
[rbxmk.encodeFormat]: #user-content-rbxmkencodeformat
<code>rbxmk.encodeFormat(format: [string](##), value: [any](##)): (bytes: [BinaryString](##))</code>

The **encodeFormat** function encodes *value* into a sequence of bytes according
to *format*. The exact details of each format are described in the
[Formats](formats.md) document.

encodeFormat will throw an error if the format does not exist, or the format has
no encoder defined.

### rbxmk.formatCanDecode
[rbxmk.formatCanDecode]: #user-content-rbxmkformatcandecode
<code>rbxmk.formatCanDecode(format: [string](##), type: [string](##)): [boolean](##)</code>

The **formatCanDecode** function returns whether *format* decodes into *type*.

formatCanDecode will throw an error if the format does not exist, or the format
does not define types it decodes into.

### rbxmk.globalAttrConfig
[rbxmk.globalAttrConfig]: #user-content-rbxmkglobalattrconfig
<code>rbxmk.globalAttrConfig: [AttrConfig][AttrConfig]?</code>

The **globalAttrConfig** field gets or sets the global AttrConfig. Most items
that utilize an AttrConfig will fallback to the global AttrConfig when possible.

See the [Value inheritance](README.md#user-content-value-inheritance) section
for details on how this field is inherited by [Instances][Instance].

### rbxmk.globalDesc
[rbxmk.globalDesc]: #user-content-rbxmkglobaldesc
<code>rbxmk.globalDesc: [RootDesc][RootDesc]?</code>

The **globalDesc** field gets or sets the global root descriptor. Most items
that utilize a root descriptor will fallback to the global descriptor when
possible.

See the [Value inheritance](README.md#user-content-value-inheritance) section
for details on how this field is inherited by [Instances][Instance].

### rbxmk.loadFile
[rbxmk.loadFile]: #user-content-rbxmkloadfile
<code>rbxmk.loadFile(path: [string](##)): (func: [function](##))</code>

The **loadFile** function loads the content of a file as a Lua function. *path*
is the path to the file.

The function runs in the context of the calling script.

### rbxmk.loadString
[rbxmk.loadString]: #user-content-rbxmkloadstring
<code>rbxmk.loadString(source: [string](##)): (func: [function](##))</code>

The **loadString** function loads the a string as a Lua function. *source* is
the string to load.

The function runs in the context of the calling script.

### rbxmk.newCookie
[rbxmk.newCookie]: #user-content-rbxmknewcookie
<code>rbxmk.newCookie(name: [string](##), value: [string](##)): [Cookie][Cookie]</code>

The **newCookie** function creates a new cookie object.

### rbxmk.newDesc
[rbxmk.newDesc]: #user-content-rbxmknewdesc
<code>rbxmk.newDesc(name: [string](##)): [Descriptor](##)</code>

The **newDesc** function creates a new descriptor object.

newDesc returns a value of whose type corresponds to the given name. The
following types may be constructed:

- [RootDesc][RootDesc]
- [ClassDesc][ClassDesc]
- [PropertyDesc][PropertyDesc]
- [FunctionDesc][FunctionDesc]
- [EventDesc][EventDesc]
- [CallbackDesc][CallbackDesc]
- [ParameterDesc][ParameterDesc]
- [TypeDesc][TypeDesc]
- [EnumDesc][EnumDesc]
- [EnumItemDesc][EnumItemDesc]

TypeDesc values are immutable. To set the fields, they can be passed as extra
arguments to newDesc:

```lua
-- Sets .Category and .Name, respectively.
local typeDesc = rbxmk.newDesc("TypeDesc", "Category", "Name")
```

ParameterDesc values are also immutable. To set the fields, they can be passed
as extra arguments to newDesc:

```lua
-- Sets .Type, .Name, and .Default, respectively.
-- No default value
local paramDesc = rbxmk.newDesc("ParameterDesc", typeDesc, "paramName")
-- Default value
local paramDesc = rbxmk.newDesc("ParameterDesc", typeDesc, "paramName", "ParamDefault")
```

### rbxmk.patchDesc
[rbxmk.patchDesc]: #user-content-rbxmkpatchdesc
<code>rbxmk.patchDesc(desc: [RootDesc][RootDesc], actions: {[DescAction][DescAction]})</code>

The **patchDesc** function transforms a root descriptor according to a list of
actions. Each action in the list is applied in order. Actions that are
incompatible are ignored.

### rbxmk.runFile
[rbxmk.runFile]: #user-content-rbxmkrunfile
<code>rbxmk.runFile(path: [string](##), args: ...[any](##)): (results: ...[any](##))</code>

The **runFile** function runs the content of a file as a Lua script. *path* is
the path to the file. *args* are passed into the script as arguments. Returns
the values returned by the script.

The script runs in the context of the referred file. Files cannot be run
recursively; if a file is already running as a script, then runFile will throw
an error.

### rbxmk.runString
[rbxmk.runString]: #user-content-rbxmkrunstring
<code>rbxmk.runString(source: [string](##), args: ...[any](##)): (results: ...[any](##))</code>

The **runString** function runs a string as a Lua script. *source* is the string
to run. *args* are passed into the script as arguments. Returns the values
returned by the script.

The script runs in the context of the calling script.

## Roblox
[roblox]: #user-content-roblox

The Roblox library contains an environment similar to the Roblox Lua API. It is
included directly into the global environment.

The **typeof** function is included to get the type of a userdata. In addition
to the usual Roblox types, `typeof` will work for various types specific to
rbxmk.

Included are constructors for the following types:

- Axes
- BrickColor
- CFrame
- Color3
- ColorSequence
- ColorSequenceKeypoint
- Faces
- [Instance][Instance]
- NumberRange
- NumberSequence
- NumberSequenceKeypoint
- PhysicalProperties
- Ray
- Rect
- Region3
- Region3int16
- UDim
- UDim2
- Vector2
- Vector2int16
- Vector3
- Vector3int16

Each of these types has an implementation that matches that of Roblox. The
[DevHub](https://developer.roblox.com/en-us/api-reference/data-types) has more
information about the API of such types.

Additionally, the [`DataModel.new`][DataModel] constructor creates a special
Instance of the DataModel class, to be used to contain instances in a game tree.

## math
[math]: #user-content-math

The **math** library is an extension to the standard library that includes the
same additions to [Roblox's math
library](https://developer.roblox.com/en-us/api-reference/lua-docs/math):

Name                | Description
--------------------|------------
[clamp][math.clamp] | Returns a number clamped between a minimum and maximum.
[log][math.log]     | Includes optional base argument.
[round][math.round] | Rounds a number to the nearest integer.
[sign][math.sign]   | Returns the sign of a number.

### math.clamp
[math.clamp]: #user-content-mathclamp
<code>math.clamp(x: [number](##), min: [number](##), max: [number](##)): [number](##)</code>

The **clamp** function returns *x* clamped so that it is not less than *min* or
greater than *max*.

### math.log
[math.log]: #user-content-mathlog
<code>math.log(x: [number](##), base: [number](##)?): [number](##)</code>

The **log** function returns the logarithm of *x* in *base*. The default for
*base* is `e`, returning the natural logarithm of *x*.

### math.round
[math.round]: #user-content-mathround
<code>math.round(x: [number](##)): [number](##)</code>

The **round** function returns *x* rounded to the nearest integer. The function
rounds half away from zero.

### math.sign
[math.sign]: #user-content-mathsign
<code>math.sign(x: [number](##)): [number](##)</code>

The **sign** function returns the sign of *x*: `1` if *x* is greater than zero,
`-1` of *x* is less than zero, and `0` if *x* equals zero.

## os
[os]: #user-content-os

The **os** library is an extension to the standard library. The following
additional functions are included:

Name                | Description
--------------------|------------
[dir][os.dir]       | Gets a list of files in a directory.
[expand][os.expand] | Expands predefined file path variables.
[getenv][os.getenv] | Gets an environment variable.
[join][os.join]     | Joins a number of file paths together.
[split][os.split]   | Splits a file path into its components.
[stat][os.stat]     | Gets metadata about a file.

### os.dir
[os.dir]: #user-content-osdir
<code>os.dir(path: [string](##)): {[File](##)}</code>

The **dir** function returns a list of files in the given directory. Each file
is a table with the same fields as returned by [os.stat][os.stat].

dir throws an error if the file does not exist.

### os.expand
[os.expand]: #user-content-osexpand
<code>os.expand(path: [string](##)): [string](##)</code>

The **expand** function scans *path* for certain variables of the form `$var` or
`${var}` an expands them. The following variables are expanded:

Variable                                    | Description
--------------------------------------------|------------
`$script_name`, `$sn`                       | The base name of the currently running script.
`$script_directory`, `$script_dir`, `$sd`   | The directory of the currently running script.
`$working_directory`, `$working_dir`, `$wd` | The current working directory.
`$temp_directory`, `$temp_dir`, `$tmp`      | The directory for temporary files.

### os.getenv
[os.getenv]: #user-content-osgetenv
<code>os.getenv(name: [string](##)?): [string](##) \| {[string](##)}</code>

The **getenv** function returns the value of the *name* environment variable. If
*name* is not specified, then a list of environment variables is returned.

### os.join
[os.join]: #user-content-osjoin
<code>os.join(paths: ...[string](##)): [string](##)</code>

The **join** function joins each *path* element into a single path, separating
them using the operating system's path separator. This also cleans up the path.

### os.split
[os.split]: #user-content-ossplit
<code>os.split(path: [string](##), components: ...[string](##)): ...[string](##)</code>

The **split** function returns the components of a file path.

Component | `project/scripts/main.script.lua` | Description
----------|-----------------------------------|------------
`base`    | `main.script.lua`                 | The file name; the last element of the path.
`dir`     | `project/scripts`                 | The directory; all but the last element of the path.
`ext`     | `.lua`                            | The extension; the suffix starting at the last dot of the last element of the path.
`fext`    | `.script.lua`                     | The format extension, as determined by registered formats.
`fstem`   | `main`                            | The base without the format extension.
`stem`    | `main.script`                     | The base without the extension.

A format extension depends on the available formats. See [Formats](formats.md)
for more information.

### os.stat
[os.stat]: #user-content-osstat
<code>os.stat(path: [string](##)): [File](##)</code>

The **stat** function gets metadata of the given file. Returns a table with the
following fields:

Field   | Type    | Description
--------|---------|------------
Name    | string  | The base name of the file.
IsDir   | boolean | Whether the file is a directory.
Size    | number  | The size of the file, in bytes.
ModTime | number  | The modification time of the file, in Unix time.

stat throws an error if the file does not exist.

## string
[string]: #user-content-string

The **string** library is an extension to the standard library that includes the
same additions to [Roblox's string
library](https://developer.roblox.com/en-us/api-reference/lua-docs/string):

Name                  | Description
----------------------|------------
[split][string.split] | Splits a string into a list of substrings.

### string.split
[string.split]: #user-content-stringsplit
<code>string.split(s: [string](##), sep: [string](##)?): {[string](##)}</code>

The **split** function splits *s* into substrings separated by *sep*.

If *sep* is nil, or if *sep* is not nil but not in *s*, then split returns a
table with *s* as its only element.

If *sep* is empty, then *s* is split after each UTF-8 sequence.

**Note**: Roblox's implementation splits per byte, while this implementation
splits per UTF-8 character.

## table
[table]: #user-content-table

The **table** library is an extension to the standard library that includes the
same additions to [Roblox's table
library](https://developer.roblox.com/en-us/api-reference/lua-docs/table):

Name                   | Description
-----------------------|------------
[clear][table.clear]   | Removes all entries from a table.
[create][table.create] | Creates a new table with a preallocated capacity.
[find][table.find]     | Find the index of a value in a table.
[move][table.move]     | Copies the entries in a table.
[pack][table.pack]     | Packs arguments into a table.
[unpack][table.unpack] | Unpacks a table into arguments.

### table.clear
[table.clear]: #user-content-tableclear
<code>table.clear(t: [table](##)?)</code>

The **clear** function removes all the entries from *t*.

### table.create
[table.create]: #user-content-tablecreate
<code>table.create(cap: [number](##), value: [any](##)?): [table](##)</code>

The **create** function returns a table with the array part allocated with a
capacity of *cap*. Each entry in the array is optionally filled with *value*.

### table.find
[table.find]: #user-content-tablefind
<code>table.find(t: [table](##), value: [any](##), init: [number](##)?): number?</code>

The **find** function returns the index in *t* of the first occurrence of
*value*, or nil if *value* was not found. Starts at index *init*, or 1 if
unspecified.

### table.move
[table.move]: #user-content-tablemove
<code>table.move(a1: [table](##), f: [number](##), e: [number](##), t: [number](##), a2: [table](##)?): table</code>

The **move** function copies elements from *a1* to *a2*, performing the
equivalent to the multiple assignment

	a2[t], ... = a1[f], ..., a1[e]

The default for *a2* is *a1*. The destination range can overlap the source
range. Returns *a2*.

### table.pack
[table.pack]: #user-content-tablepack
<code>table.pack(...[any](##)?): table</code>

The **pack** function returns a table with each argument stored at keys 1, 2,
etc. Also sets field "n" to the number of arguments. Note that the resulting
table may not be a sequence.

### table.unpack
[table.unpack]: #user-content-tableunpack
<code>table.unpack(list: [table](##), i: [number](##)?, j: [number](##)?): ...[any](##)</code>

Returns the elements from *list*, equivalent to

	list[i], list[i+1], ..., list[j]

By default, *i* is 1 and *j* is the length of *list*.

## sym
[sym]: #user-content-sym

The **sym** library contains **Symbol** values. A symbol is a unique identifier
that can be used to access certain metadata fields of an [Instance][Instance].

An instance can be indexed with a symbol to get a metadata value in the same way
it can be indexed with a string to get a property value:

```lua
local instance = Instance.new("Workspace")
instance[sym.IsService] = true
print(instance[sym.IsService]) --> true
```

The following symbols are defined:

Symbol                                            | Description
--------------------------------------------------|------------
[`sym.AttrConfig`][Instance.sym.AttrConfig]       | Gets the inherited [AttrConfig][AttrConfig] of an instance.
[`sym.Desc`][Instance.sym.Desc]                   | Gets the inherited [descriptor][RootDesc] of an instance.
[`sym.IsService`][Instance.sym.IsService]         | Determines whether an instance is a service.
[`sym.RawAttrConfig`][Instance.sym.RawAttrConfig] | Accesses the direct [AttrConfig][AttrConfig] of an instance.
[`sym.RawDesc`][Instance.sym.RawDesc]             | Accesses the direct [descriptor][RootDesc] of an instance.
[`sym.Reference`][Instance.sym.Reference]         | Determines the value used to identify the instance.

## types
[types]: #user-content-types

The **types** library contains functions for constructing explicit primitives.
The name of a function corresponds directly to the type. See [Explicit
primitives](README.md#user-content-explicit-primitives) for more information.

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

[AttrConfig]: types.md#user-content-attrconfig
[CallbackDesc]: types.md#user-content-callbackdesc
[ClassDesc]: types.md#user-content-classdesc
[Cookie]: types.md#user-content-cookie
[Cookies]: types.md#user-content-cookies
[DataModel]: types.md#user-content-datamodel
[DescAction]: types.md#user-content-descaction
[EnumDesc]: types.md#user-content-enumdesc
[EnumItemDesc]: types.md#user-content-enumitemdesc
[EventDesc]: types.md#user-content-eventdesc
[FunctionDesc]: types.md#user-content-functiondesc
[Instance.sym.AttrConfig]: types.md#user-content-instancesymattrconfig
[Instance.sym.Desc]: types.md#user-content-instancesymdesc
[Instance.sym.IsService]: types.md#user-content-instancesymisservice
[Instance.sym.RawAttrConfig]: types.md#user-content-instancesymrawattrconfig
[Instance.sym.RawDesc]: types.md#user-content-instancesymrawdesc
[Instance.sym.Reference]: types.md#user-content-instancesymreference
[Instance]: types.md#user-content-instance
[ParameterDesc]: types.md#user-content-parameterdesc
[PropertyDesc]: types.md#user-content-propertydesc
[RootDesc]: types.md#user-content-rootdesc
[TypeDesc]: types.md#user-content-typedesc
