[DRAFT]

# Documentation
This document provides details on how rbxmk works. For a basic overview, see
[USAGE.md](USAGE.md).

# Environment

## Standard library
The following items from the standard library are included:

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

## Roblox environment
The rbxmk Lua environment includes an environment similar to the Roblox Lua API.
The following items from the are included.

- `typeof`
- Roblox types (e.g. Instance, CFrame, Vector3, etc)

## `rbxmk` library
The `rbxmk` library contains functions related to the rbxmk engine.

Name           | Description
---------------|------------
`load`         | Run a script.
`encodeformat` | Serialize data into bytes.
`decodeformat` | Deserialize data from bytes.
`readsource`   | Read bytes from an external source.
`writesource`  | Write bytes to an external source.

### `load(path: string, args: ...any): (results ...any)`
### `encodeformat(format: string, value: any): (bytes BinaryString)`
### `decodeformat(format: string, bytes: BinaryString): (value any)`
### `readsource(source: string, args: ...any): (bytes: BinaryString)`
### `writesource(source: string, byte: BinaryString, args: ...any)`

## `os` library
The `os` library is an extension to the standard library. The following
additional functions are included:

Name     | Description
---------|------------
`split`  | Splits a file path into its components.
`join`   | Joins a number of file paths together.
`expand` | Expands predefined file path variables.
`getenv` | Gets an environment variable.
`dir`    | Gets a list of files in a directory.

### `split(path: string, components: ...string): ...string`
### `join(paths: ...string) string`
### `expand(path: string): string`
### `getenv(name: string?): string | Array<string>`
The `getenv` function returns the value of the *name* environment variable. If
*name* is not specified, then a list of environment variables is returned.

### `dir(path: string): Array<File>`

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

In rbxmk, properties have no types by default, because there are no descriptors.
For example, when a property is set to a Lua number, it is always converted into
a `double`. In the absence of extra type metadata, the user needs some way to
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

# Sources
A **source** is an external location from which raw data can be read from and
written to.

A source can be accessed at a low level through the `rbxmk.readsource` and
`rbxmk.writesource` functions.

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

A format can be accessed at a low level through the `rbxmk.encodeformat` and
`rbxmk.decodeformat` functions.

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
`rbxl`  |
`rbxm`  |
`rbxlx` |
`rbxmx` |
