# Formats
This document contains a reference to the formats available to rbxmk scripts.

<table>
<thead><tr><th>Table of Contents</th></tr></thead>
<tbody><tr><td>

1. [bin][bin]
2. [csv][csv]
3. [l10n.csv][l10n.csv]
4. [json][json]
5. [desc.json][desc.json]
6. [desc-patch.json][desc-patch.json]
7. [lua][lua]
8. [client.lua][client.lua]
9. [localscript.lua][localscript.lua]
10. [modulescript.lua][modulescript.lua]
11. [script.lua][script.lua]
12. [server.lua][server.lua]
13. [luau][luau]
14. [client.luau][client.luau]
15. [localscript.luau][localscript.luau]
16. [modulescript.luau][modulescript.luau]
17. [script.luau][script.luau]
18. [server.luau][server.luau]
19. [rbxattr][rbxattr]
20. [rbxl][rbxl]
21. [rbxlx][rbxlx]
22. [rbxm][rbxm]
23. [rbxmx][rbxmx]
24. [txt][txt]

</td></tr></tbody>
</table>

A **format** is capable of encoding a value to raw bytes, or decoding raw bytes
into a value. A format can be accessed at a low level through the
[rbxmk.encodeFormat](libraries.md#user-content-rbxmkencodeformat) and
[rbxmk.decodeFormat](libraries.md#user-content-rbxmkdecodeformat) functions.

The name of a format corresponds to the extension of a file name. For example,
the `lua` format corresponds to the `.lua` file extension. When determining a
format from a file extension, format names are greedy; if a file extension is
`.server.lua`, this will select the `server.lua` format before the `lua` format.
For convenience, in places where a format name is received, the name may have an
optional leading `.` character.

APIs that require a format receive a [FormatSelector][FormatSelector]. A
selector is usually a string, which specifies the name of the format to use. A
selector can also be a table, where the "Format" field specifies the name, and a
number of additional options can be specified. If a format has any such options,
this document will describe them.

A format can decode into a number of certain types, and encode a number of
certain types. A format may also have no definition for either decoding or
encoding at all.

A format that can encode a **Stringable** type accepts any type that can be
converted to a string. Additionally, an [Instance][Instance] will be accepted as
a Stringable when it has a particular [ClassName][Instance.ClassName], with a
selected property that has a [Stringlike][Stringlike] value. In this case, the
property is encoded.

ClassName         | Property
------------------|---------
LocalizationTable | Contents
LocalScript       | Source
ModuleScript      | Source
Script            | Source

A format that can encode a **Numberable** type accepts any type that can be
converted to a floating-point number. An **Intable** is similar, converting to
an integer instead.

## bin
[bin]: #user-content-bin

The **bin** format encodes string-like values with the assurance that the bytes
will be interpreted exactly as-is.

Direction | Type         | Description
----------|--------------|------------
Decode    | BinaryString | Raw binary data.
Encode    | Stringable   | Any string-like value.

This format has no options.

## csv
[csv]: #user-content-csv

The **csv** format decodes comma-separated values into a two-dimensional array.

Direction | Type  | Description
----------|-------|------------
Decode    | Array | An array of arrays of strings.
Encode    | Array | An array of arrays of strings.

CSV data decodes into a two-dimensional array of strings. For example,

	A,B,C
	D,E,F
	G,H,I

decodes into

	{
		{"A", "B", "C"),
		{"D", "E", "F"),
		{"G", "H", "I"),
	}

When encoding, each field must be string-like, but cannot be an Instance.

When decoding, each record must have the same number of fields. When encoding,
records do not need to have the same number of fields.

This format has no options.

## l10n.csv
[l10n.csv]: #user-content-l10ncsv

The **l10n.csv** format decodes comma-separated localization data into a
LocalizationTable instance, where the data is assigned to the Contents property.

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A LocalizationTable with a Contents property.
Encode    | Stringable           | Any string-like value.

Decoded data is a string in JSON format with the following structure:

	[
		{
			"key":      "string",
			"context":  "string",
			"examples": "string",
			"source":   "string",
			"values": {
				"locale": "string",
				...
			},
		},
		...
	]

Well-formed data has certain constraints, which are described in the
[LocalizationTable
page](https://developer.roblox.com/en-us/api-reference/class/LocalizationTable)
of the DevHub. rbxmk applies these same constraints when encoding and decoding.
To avoid data loss, they are applied more strictly. Rather than discarding data,
any conflict that arises will throw an error that describes the conflict in
detail.

This format has no options.

## json
[json]: #user-content-json

The **json** format is defined for encoding general data in JSON format.

Direction | Type       | Description
----------|------------|------------
Decode    | nil        | A JSON null.
Decode    | boolean    | A JSON boolean.
Decode    | number     | The nearest representation of a JSON number.
Decode    | string     | A JSON string.
Decode    | Array      | An JSON array.
Decode    | Dictionary | A JSON object.
Encode    | nil        | A Lua nil.
Encode    | boolean    | A Lua boolean.
Encode    | number     | A Lua number.
Encode    | string     | A Lua string.
Encode    | Array      | An array-like table, having a non-zero length.
Encode    | Dictionary | A dictionary-like table, having a length of zero.

Other value types are encoded as null.

This format has the following options:

Field  | Type   | Default | Description
-------|--------|---------|------------
Indent | string | `"\t"`  | Determines the indentation of encoded content. If an empty string, then the content is minified.

## desc.json
[desc.json]: #user-content-descjson

The **desc.json** format encodes a root descriptor file, more commonly known as
an "API dump", in JSON format.

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [RootDesc][RootDesc] | A root descriptor.
Encode    | [RootDesc][RootDesc] | A root descriptor.

This format has no options.

## desc-patch.json
[desc-patch.json]: #user-content-desc-patchjson

The **desc-patch.json** format encodes actions that transform descriptors, in
JSON format.

Direction | Type        | Description
----------|-------------|------------
Decode    | DescActions | A list of [DescAction][DescAction] values.
Encode    | DescActions | A list of [DescAction][DescAction] values.

This format has no options.

## lua
[lua]: #user-content-lua

The **lua** format is an alias for [modulescript.lua][modulescript.lua].

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A ModuleScript with a Source property.
Encode    | Stringable           | Any string-like value.

This format has no options.

## client.lua
[client.lua]: #user-content-clientlua

The **client.lua** format is an alias for [localscript.lua][localscript.lua].

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A LocalScript with a Source property.
Encode    | Stringable           | Any string-like value.

This format has no options.

## localscript.lua
[localscript.lua]: #user-content-localscriptlua

The **localscript.lua** format is a shortcut for decoding Lua code into a
LocalScript instance, where the content is assigned to the Source property.

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A LocalScript with a Source property.
Encode    | Stringable           | Any string-like value.

This format has no options.

## modulescript.lua
[modulescript.lua]: #user-content-modulescriptlua

The **modulescript.lua** format is a shortcut for decoding Lua code into a
ModuleScript instance, where the content is assigned to the Source property.

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A ModuleScript with a Source property.
Encode    | Stringable           | Any string-like value.

This format has no options.

## script.lua
[script.lua]: #user-content-scriptlua

The **script.lua** format is a shortcut for decoding Lua code into a
Script instance, where the content is assigned to the Source property.

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A Script with a Source property.
Encode    | Stringable           | Any string-like value.

This format has no options.

## server.lua
[server.lua]: #user-content-serverlua

The **server.lua** format is an alias for [script.lua][script.lua].

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A Script with a Source property.
Encode    | Stringable           | Any string-like value.

This format has no options.

## luau
[luau]: #user-content-luau

The **luau** format is an alias for [modulescript.lua][modulescript.lua].

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A ModuleScript with a Source property.
Encode    | Stringable           | Any string-like value.

This format has no options.

## client.luau
[client.luau]: #user-content-clientluau

The **client.luau** format is an alias for [localscript.lua][localscript.lua].

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A LocalScript with a Source property.
Encode    | Stringable           | Any string-like value.

This format has no options.

## localscript.luau
[localscript.luau]: #user-content-localscriptluau

The **localscript.luau** format is an alias for
[localscript.lua][localscript.lua].

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A LocalScript with a Source property.
Encode    | Stringable           | Any string-like value.

This format has no options.

## modulescript.luau
[modulescript.luau]: #user-content-modulescriptluau

The **modulescript.luau** format is an alias for
[modulescript.lua][modulescript.lua].

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A ModuleScript with a Source property.
Encode    | Stringable           | Any string-like value.

This format has no options.

## script.luau
[script.luau]: #user-content-scriptluau

The **script.luau** format is an alias for [script.lua][script.lua].

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A Script with a Source property.
Encode    | Stringable           | Any string-like value.

This format has no options.

## server.luau
[server.luau]: #user-content-serverluau

The **server.luau** format is an alias for [script.lua][script.lua].

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A Script with a Source property.
Encode    | Stringable           | Any string-like value.

This format has no options.

## rbxattr
[rbxattr]: #user-content-rbxattr

The **rbxattr** format is defined for serializing instance attributes, encoding
a Dictionary of attribute values.

Direction | Type       | Description
----------|------------|------------
Decode    | Dictionary | A dictionary of attribute names mapped to values.
Encode    | Dictionary | A dictionary of attribute names mapped to values.

The following value types are encoded and decoded:
- string
- bool
- float
- double
- UDim
- UDim2
- BrickColor
- Color3
- Vector2
- Vector3
- NumberSequence
- ColorSequence
- NumberRange
- Rect

Additionally, any Stringable value is encoded as a string, and any Numberable
value is encoded as a double.

This format has no options.

## rbxl
[rbxl]: #user-content-rbxl

The **rbxl** format encodes Instances in the Roblox binary place format.

Direction | Type                   | Description
----------|------------------------|------------
Decode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [Instance][Instance]   | A single instance, interpreted as a child to a DataModel.
Encode    | Objects                | A list of Instances, interpreted as children to a DataModel.

This format has the following options:

Field    | Type                                | Default       | Description
---------|-------------------------------------|---------------|------------
Desc     | [RootDesc][RootDesc] \| bool \| nil | `nil`         | Sets the descriptor to be used when encoding or decoding. If `false`, then no descriptor is used. Otherwise, the descriptor of the root instance is used **with all descendants**, falling back to [globalDesc][rbxmk.globalDesc].
DescMode | string                              | `"NonStrict"` | Determines how deviations from the descriptor are handled. `"NonStrict"` causes deviations to be ignored. `"Strict"` causes an error to be thrown for the first deviation. `"Preserve"` tries to retain as much information as possible, usually by behaving as if no descriptor is set.

## rbxlx
[rbxlx]: #user-content-rbxlx

The **rbxlx** format encodes Instances in the Roblox XML place format.

Direction | Type                   | Description
----------|------------------------|------------
Decode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [Instance][Instance]   | A single instance, interpreted as a child to a DataModel.
Encode    | Objects                | A list of Instances, interpreted as children to a DataModel.

This format has the same options as the [rbxl][rbxl] format.

## rbxm
[rbxm]: #user-content-rbxm

The **rbxm** format encodes Instances in the Roblox binary model format.

Direction | Type                   | Description
----------|------------------------|------------
Decode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [Instance][Instance]   | A single instance, interpreted as a child to a DataModel.
Encode    | Objects                | A list of Instances, interpreted as children to a DataModel.

This format has the same options as the [rbxl][rbxl] format.

## rbxmx
[rbxmx]: #user-content-rbxmx

The **rbxmx** format encodes Instances in the Roblox XML model format.

Direction | Type                   | Description
----------|------------------------|------------
Decode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [Instance][Instance]   | A single instance, interpreted as a child to a DataModel.
Encode    | Objects                | A list of Instances, interpreted as children to a DataModel.

This format has the same options as the [rbxl][rbxl] format.

## txt format
[txt]: #user-content-txt-format

The **txt** format encodes UTF-8 text.

Direction | Type       | Description
----------|------------|------------
Decode    | string     | UTF-8 text.
Encode    | Stringable | Any string-like value.

This format has no options.

[DataModel]: types.md#user-content-datamodel
[DescAction]: types.md#user-content-descaction
[FormatSelector]: types.md#user-content-formatselector
[Instance.ClassName]: types.md#user-content-instanceclassname
[Instance]: types.md#user-content-instance
[Intlike]: types.md#user-content-intlike
[Numberlike]: types.md#user-content-numberlike
[rbxmk.globalDesc]: libraries.md#user-content-rbxmkglobaldesc
[RootDesc]: types.md#user-content-rootdesc
[Stringlike]: types.md#user-content-stringlike
