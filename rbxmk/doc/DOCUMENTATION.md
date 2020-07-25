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
with Lua numbers, as well as other numeric exprims. Stringlike exprims can be
concatenated with Lua strings, as well as other stringlike exprims. The result
of such operations is always a value of the underlying primitive. For example,
`types.int(9) + 1` returns 10 as a Lua number.

Because of the limitations of metatables, exprims can be correctly compared only
with other exprims of the same type.

Additionally, an exprim can be called as function, which returns the underlying
primitive.

	local i = types.int(42)
	print(i() == 42)

## Source libraries
Formats and sources can be accessed at a low level through functions in the
`rbxmk` library. A source library wraps this workflow of encoding/decoding a
format and reading from or writing to a source into a single convenient package.

### `file` library
The `file` library handles the `file` source.

Name    | Description
--------|------------
`read`  | Reads data from a file in a certain format.
`write` | Writes data to a file in a certain format.

# Formats
# Sources
