[DRAFT]

# Documentation
This document provides details on how rbxmk works. For a basic overview, see
[USAGE.md](USAGE.md).

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
types     | Fallbacks for constructing certain types.
(sources) | An assortment of libraries for interfacing with the various external sources.

## Base library
The following items from the Lua 5.1 standard library are included:

- `_G`
- `_VERSION`
- `assert`
- `error`
- `ipairs`
- `next`
- `pairs`
- `pcall`
- `print`
- `select`
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

### `rbxmk.meta(inst: Instance, name: string, value: any?) any?`
The `meta` function gets or sets metadata on an instance. `meta` has two
signatures; passing two arguments gets a value, while passing three arguments
sets a value:

- Get: `meta(inst: Instance, name: string): (value: any)`
- Set: `meta(inst: Instance, name: string, value: any)`

The following metadata values are possible:

#### Desc
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

#### IsService
IsService indicates whether the instance is a service, such as Workspace or
Lighting. This is used by some formats to determine how to encode and decode the
instance.

#### RawDesc
RawDesc is similar to Desc, except that it considers only the direct descriptor
of the current instance.

Getting RawDesc will return a RootDesc if the instance has a descriptor
assigned, false if the descriptor is blocked, or nil if no descriptor is
assigned.

Setting RawDesc behaves the same as setting Desc.

#### Reference
Reference is a string used to refer to the instance from within a DataModel.
Certain formats use this to encode a reference to an instance. For example, the
RBXMX format will generate random UUIDs for its references (e.g.
"RBX8B658F72923F487FAE2F7437482EF16D").

### `rbxmk.newDesc(name: string): Descriptor`
The `newDesc` function creates a new descriptor object.

*name* may be one of the following:

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

`newDesc` returns a value of whose type corresponds to the given name.

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

## `types` library
The `types` library contains functions for constructing explicit primitives. The
name of a function corresponds directly to the type.

Type              | Primitive
------------------|----------
`BinaryString`    | string
`Content`         | string
`ProtectedString` | string
`SharedString`    | string
`float`           | number
`int`             | number
`int64`           | number
`token`           | number

The properties of instances in Roblox have a number of different types. Many of
these types can be expressed in Lua through constructors. Examples of such are
`CFrame`, `Vector3`, `UDim2`, and so on. These types correspond to internal data
types within the Roblox engine. The Lua representation of, say, a `CFrame`, is a
userdata with accessible fields.

Some Roblox types are represented with a simple Lua primitive, such as a number
or string. For example, the Roblox types `int`, `int64`, `float`, and `double`
all map to Lua's `number` type. When setting a property, the engine is able to
reflect this Lua `number` back to the correct Roblox type, because the property
has extra metadata called a **descriptor** that includes the property's type.

In rbxmk, when no descriptors are specified, properties have no types. For
example, when a property is set to a Lua number, it is always converted into a
`double`. In the absence of extra type information, the user needs some way to
set specific Roblox types.

This problem is solved with "explicit primitives", or **exprims**. An exprim is
a userdata representation of an otherwise ambiguous type. This userdata carries
type metadata along with a given value, allowing the value to be mapped to the
correct Roblox type when it is set as a property.

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
which also has no exprim.

An exprim userdata supports operations with its underlying primitive type, as
well as other similar exprims. Numeric exprims can be added, multiplied, etc,
with Lua numbers, as well as other numeric exprims. String-like exprims can be
concatenated with Lua strings, as well as other string-like exprims. The result
of such operations is always a value of the underlying primitive. For example,
`types.int(9) + 1` returns 10 as a Lua number.

Because of the limitations of metatables, exprims can be correctly compared only
with other exprims of the same type.

Additionally, an exprim can be called as function, which returns the underlying
primitive.

	local i = types.int(42)
	print(i() == 42)

Exprims are meant to be a niche feature for use in the absence of descriptors.
The user will find descriptors to be a much more convenient solution.

# Descriptors

# Sources
A **source** is an external location from which raw data can be read from and
written to.

A source can be accessed at a low level through the `rbxmk.readSource` and
`rbxmk.writeSource` functions.

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
into a value.

A format can be accessed at a low level through the `rbxmk.encodeFormat` and
`rbxmk.decodeFormat` functions.

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
