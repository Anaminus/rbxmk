# Documentation
This document provides details on how rbxmk works. For a basic overview, see
[USAGE.md](USAGE.md).

<table>
<thead><tr><th>Table of Contents</th></tr></thead>
<tbody><tr><td><ol>
	<li><a href="#user-content-standard-library">Standard library</a></li>
	<li><a href="#user-content-rbxmk-library">rbxmk library</a></li>
	<li><a href="#user-content-resolve-chain">Resolve chain</a><ol>
		<li><a href="#user-content-input-procedure">Input procedure</a></li>
		<li><a href="#user-content-output-procedure">Output procedure</a></li>
		<li><a href="#user-content-deletion-procedure">Deletion procedure</a></li>
	</ol></li>
	<li><a href="#user-content-reference">Reference</a></li>
	<li><a href="#user-content-data-types">Data types</a><ol>
		<li><a href="#user-content-merging-overview">Merging overview</a></li>
		<li><a href="#user-content-instances-type">Instances type</a></li>
		<li><a href="#user-content-instance-type">Instance type</a></li>
		<li><a href="#user-content-properties-type">Properties type</a></li>
		<li><a href="#user-content-property-type">Property type</a></li>
		<li><a href="#user-content-value-type">Value type</a></li>
		<li><a href="#user-content-stringlike-type">Stringlike type</a></li>
		<li><a href="#user-content-region-type">Region type</a><ol>
			<li><a href="#user-content-region-drilling">Region drilling</a></li>
		</ol></li>
		<li><a href="#user-content-delete-type">Delete type</a></li>
	</ol></li>
	<li><a href="#user-content-formats">Formats</a><ol>
		<li><a href="#user-content-roblox-formats">Roblox formats</a></li>
		<li><a href="#user-content-lua-formats">Lua formats</a></li>
		<li><a href="#user-content-text-formats">Text formats</a></li>
		<li><a href="#user-content-property-formats">Property formats</a></li>
	</ol></li>
	<li><a href="#user-content-schemes">Schemes</a><ol>
		<li><a href="#user-content-file-scheme">File scheme</a></li>
		<li><a href="#user-content-http-scheme">HTTP scheme</a></li>
		<li><a href="#user-content-roblox-asset-id-scheme">Roblox asset ID scheme</a></li>
		<li><a href="#user-content-generate-scheme">Generate scheme</a><ol>
			<li><a href="#user-content-instance-syntax">Instance syntax</a><ol>
				<li><a href="#user-content-instance-item">Instance item</a></li>
				<li><a href="#user-content-property-item">Property item</a></li>
				<li><a href="#user-content-meta-property-item">Meta property item</a></li>
			</ol></li>
			<li><a href="#user-content-property-syntax">Property syntax</a></li>
			<li><a href="#user-content-value-syntax">Value syntax</a></li>
			<li><a href="#user-content-example">Example</a></li>
		</ol></li>
	</ol></li>
	<li><a href="#user-content-filters">Filters</a><ol>
		<li><a href="#user-content-preprocess-filter">Preprocess filter</a></li>
	</ol></li>
</ol></td></tr></tbody>
</table>


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

## rbxmk library
The `rbxmk` library contains the following functions:

Name                                          | Description
----------------------------------------------|------------
[configure](#user-content-configure-function) | Configure the behavior of rbxmk.
[delete](#user-content-delete-function)       | Delete an output node.
[filename](#user-content-filename-function)   | Get parts of a file name.
[filter](#user-content-filter-function)       | Transform nodes.
[getenv](#user-content-getenv-function)       | Get the value of an environment variable.
[input](#user-content-input-function)         | Create an input node.
[load](#user-content-load-function)           | Load and execute a script.
[loadapi](#user-content-loadapi-function)     | Load an API file.
[map](#user-content-map-function)             | Map one or more inputs to one or more outputs.
[output](#user-content-output-function)       | Create an output node.
[path](#user-content-path-function)           | Join file paths.
[printf](#user-content-printf-function)       | Print a formatted string to stdout.
[readdir](#user-content-readdir-function)     | List the files in a directory.
[sprintf](#user-content-sprintf-function)     | Return a formatted string.
[type](#user-content-type-function)           | Return the type of a value as a string.

### `configure` function
`rbxmk.configure{...}`

`configure` changes the behavior of rbxmk. Each named argument specifies an
option to change. The following options are available:

#### Configure options table
Name      | Type   | Default      | Description
----------|--------|--------------|------------
`api`     | api    | nil          | Sets the default API value to be used by all applicable functions. Several functions have an `api` argument, which can be used to override the default API for that call.
`define`  | table  | -            | Defines a number of variables to be used by the [preprocessor](#user-content-preprocess-filter). Each key-value pair in the given table is merged with the existing set of variables. Keys must be strings that are identifiers, and values can be either bools, numbers, or strings.<br>Note that variables defined by `configure` cannot override variables defined by the [`--define`](USAGE.md#user-content-command-options) command option.
`host`    | string | "roblox.com" | The website used for interacting with Roblox web APIs.
`rbxauth` | table  | -            | Specifies a number of Roblox user accounts to authenticate. [(More)](#user-content-config-rbxauth)
`undef`   | table  | -            | Undefines a number of [preprocessor](#user-content-preprocess-filter) variables. The given table is a list of variable names to be undefined.<br>If both `define` and `undef` configs are given, then `undef` is applied first.

#### Config `rbxauth`
The `rbxauth` option is a table specifying a number of Roblox user accounts to
authenticate. Each entry is a table with the following fields (of which each are
optional unless otherwise specified):

Field    | Type          | Description
---------|---------------|------------
`type`   | string        | Specifies the type of identifier. Can be either "Username", "Email", "PhoneNumber", or "UserID".
`ident`  | string/number | The identifier itself. In the case of UserID, the value is interpreted as an integer. Otherwise, the value is interpreted as a string.
`prompt` | boolean       | If true, then the user will be prompted to login. An unspecified type or identifier will be prompted for as necessary.
`file`   | string        | Specifies the path to a file, which contains cookies representing an authenticated session. The file must be formatted as a list of "Set-Cookie" HTTP headers. This overrides the `prompt` option.
`logout` | boolean       | If true, then the user of the specified credentials will be logged out. In this case, the "type" and "ident" fields must be specified.

### `delete` function
`rbxmk.delete{...output}`

For each output node received, `delete` removes the data pointed to by the node.

### `filename` function
`string = rbxmk.filename{string, ...string}`

`filename` returns parts of a file path. The first argument is the path. Each
remaining argument is a string specifying a part of the path to return. Each
return value corresponds to each specified part.

The following part values are available:

Option  | `project/scripts/main.script.lua` | Description
--------|-----------------------------------|------------
`dir`   | `project/scripts`                 | The directory; all but the last element of the path.
`base`  | `main.script.lua`                 | The file name; the last element of the path.
`ext`   | `.lua`                            | The extension; the suffix starting at the last dot of the last element of the path.
`stem`  | `main.script`                     | The base without the extension.
`fext`  | `.script.lua`                     | The format extension, as determined by `schemes.GuessFileExtension`.
`fstem` | `main`                            | The base without the format extension.

### `filter` function
`... = rbxmk.filter{api=nil, string, ...}`

`filter` transforms nodes. The first argument specifies the filter name.
Subsequent arguments and return values depend on the filter. See
[Filters](#user-content-filters) for a list of filters and their arguments.

The optional `api` argument specifies an API value to enhance the handling of
instances and properties. Specifying a non-nil API overrides the default API.
The API is used only by applicable filters.

### `getenv` function
`string = rbxmk.getenv{string}`

`getenv` returns the value of an environment variable of the given name. Returns
nil if the variable is not present.

### `input` function
`node = rbxmk.input{format=string, api=nil, user=nil, ...string}`

`input` creates an input node. The arguments specify the
[Reference](#user-content-reference) to the input.

The optional `format` argument forces the file format, if needed. This can be
either a format name or format extension. That is, the leading dot character
(`.`) is optional.

The optional `api` argument specifies an API value to enhance the handling of
instances and properties. Specifying a non-nil API overrides the default API.

The optional `user` argument specifies user credentials when authentication is
required (when downloading from a website, for example). The argument can be one
of the following types:

Type   | Description
-------|------------
table  | Requires the `type` and `ident` fields defined in the [rbxauth config](#user-content-config-rbxauth).
string | A user name identifier. The type is set as "Username".
number | A user ID identifier. The type is set as "UserID".

### `load` function
`... = rbxmk.load{string, ...}`

`load` executes a file as a script. The first argument is the file name.
Subsequent arguments are passed to the script (which may be received with the
`...` operator). `load` returns any values returned by the script.

### `loadapi` function
`api = rbxmk.loadapi{string}`

`loadapi` receives a path to a file containing an API dump, returning a value
representing the API. This value can be passed to certain functions to enhance
how instances and properties are handled.

### `map` function
`rbxmk.map{...node}`

`map` maps one or more input nodes to one or more output nodes. Either kind of
node may be passed to `map`.

Nodes are mapped in the order they are received. That is, inputs are gathered in
one list, and outputs are gathered in another. Then each node in the input list
is mapped to each node in the output list, in order. For example:

```lua
A, B = rbxmk.input{...}, rbxmk.input{...}
X, Y = rbxmk.output{...}, rbxmk.output{...}
rbxmk.map{A, X, B, Y}
-- 1: A -> X
-- 2: A -> Y
-- 3: B -> X
-- 4: B -> Y
```

### `output` function
```lua
node = rbxmk.output{format=string, api=nil, user=nil, ...string}
```

`output` creates an output node. The arguments specify the
[Reference](#user-content-reference) to the output.

The optional `format` argument forces the file format, if needed. This can be
either a format name or format extension. That is, the leading dot character
(`.`) is optional.

The optional `api` argument specifies an API value to enhance the handling of
instances and properties. Specifying a non-nil API overrides the default API.

The optional `user` argument specifies user credentials when authentication is
required (when uploading to a website, for example). This argument behaves the
same as with the [`input`](#user-content-input-function) function.

### `path` function
`string = rbxmk.path{...}`

`path` receives a number of strings, and joins them together into a single file
path, adding separators as necessary.

A string may contain variables, which begin with a `$` followed by a sequence of
letters, digits, and underscores. Variables will be expanded into their final
values before the string is joined.

The following variables are available:

Variable                                 | Description
-----------------------------------------|------------
`script_directory`, `script_dir`, `sd`   | Expands to the directory of the script currently running.
`script_name`, `sn`                      | Expands to the base name of the script currently running.
`working_directory`, `working_dir`, `wd` | Expands to the current working directory.
`temp_directory`, `temp_dir`, `tmp`      | Expants to the directory for temporary files.

Any other variable returns an empty string. An empty string will also be
returned if a path could not be located.

### `printf` function
`rbxmk.printf{string, ...}`

`printf` receives a number of values, formats them according to the first
argument, and writes the result to standard output. This function follows the
same rules as Go's [fmt.Printf](https://golang.org/pkg/fmt/#Printf).

### `readdir` function
`table = rbxmk.readdir{string}`

`readdir` receives a directory path and returns a list of files within the
directory. Each file is represented by a table with the following fields:

Field     | Description
----------|------------
`name`    | The base name of the file.
`isdir`   | Whether the file is a directory.
`size`    | The size of the file in bytes.
`modtime` | The modification time, in Unix time.

### `sprintf` function
`string = rbxmk.sprintf{string, ...}`

`sprintf` receives a number of values, formats them according to the first
argument, and returns the resulting string. This function follows the same rules
as Go's [fmt.Sprintf](https://golang.org/pkg/fmt/#Sprintf).

### `type` function
`string = rbxmk.type{value}`

`type` returns the type of the given value as a string. In addition to the
regular Lua types, the following types are available:

- `input`: an input node.
- `output`: an output node.

## Resolve chain
When creating or mapping an input or output node, rbxmk has a procedure that
chains together predefined components in order to resolve the node. In its most
basic form, the procedure looks like this:

```
Scheme -> Format -> Drills
```

1. Scheme retrieves a file containing raw data.
2. Format turns the file into data of a known type.
3. Drills select data within the data.

The exact procedures for inputs and outputs differ, but still follow this
sequence overall.

### Input procedure
This procedure is used when creating an input node.

1. Determine the Scheme.
	- Uses the first string of the [Reference](#user-content-reference).
2. Use Scheme to get a file from a specified location.
3. Determine the Format.
	- Provided by the `format` argument, or guessed, depending on the Scheme.
4. Decode the file using Format.
	- Format returns Data of a known type.
5. Drill into Data using Data's Drill function.
	- Data has a Drill function that drills into itself, returning another Data
	  which, in turn, may also have Drill function.

### Output procedure
This procedure is used when mapping an input node to an output node.

1. Determine the Scheme.
	- Uses the first string of the [Reference](#user-content-reference).
2. Use Scheme to get current state of the output from a specified location; may
   be skipped if not applicable to the Scheme.
	1. Use Scheme to get a file from a specified location.
	2. Determine the Format.
		- Provided by the `format` argument, or guessed, depending on the
		  Scheme.
	3. Decode the file using Format.
		- Format returns Data of a known type.
	4. Drill into Data using Data's Drill function.
		- Data has a Drill function that drills into itself, returning another
		  Data which, in turn, may also have Drill function.
3. Merge input Data into decoded Data using Format.
	- Input Data has a Merge function that knows how to merge the input into an
	  output of another type.
4. Encode resulting Data into a file using Format.
5. Use Scheme to write the file to the specified location.

### Deletion procedure
The `rbxmk.delete` function works exactly like mapping an input to an output,
where the input is a special `Delete` type. This type merges into other types by
deleting content instead of adding to or replacing it.

## Reference
A Reference is a list of strings used to specify a piece of Data. It is passed
into an input or output resolve chain. Each step of the chain processes a number
of strings, then passes the remaining strings to the next step.

The first string in the reference is a URI pointing to a location that usually
contains a file. This URI is resolved depending on the scheme (e.g. `file://`,
`http://`, etc). If the scheme part of the URI is omitted, the the URI is
assumed to be of the `file` scheme.

While most schemes deal primarily with the first string, a scheme may process
any number of strings.

After the format has been decoded into a Data, the remaining Reference is
continuously passed to the Data's Drill; the drill returns another Data, which
itself has a drill. Each drill works like a step in the chain, processing a
number of strings, and returning the remaining strings to the next drill.

This continues until and EOD (end-of-drill) marker is received. This happens
when there are no more Reference strings to process, or when the Data simply
cannot be drilled into.

Here's an example. Each unnamed argument to the `rbxmk.input` function is a
string of a Reference.

```
rbxmk.input{"file://place.rbxl", "Workspace.Model.Script", "Source"}
```

1. From the first string, the scheme is determined to be `file`.
2. The `file` scheme reads the `place.rbxl` from the file system, guessing the
   format as `rbxl` based on the file's extension.
3. Using the `rbxl` format, the contents of the file are decoded into some kind
   of Data (in this case, a list of Instances).
4. The next string is used by the Instance list drill, which selects a
   descendant instance within the list of instances. In this case, the "Script"
   instance is selected, which is of the Instance type.
5. The string after that is used by the Instance's drill, which selects a
   property within an instance. In this case, the value of the script's "Source"
   property is selected, which is of the Property type.
6. There are no more strings to process, so the Property's drill returns EOD,
   signalling the end of the drill.

## Data types
**Data** is a value with a type known by rbxmk. Each type has a Drill and Merge
function. A Drill function selects a component of the Data. A Merge function
detimines how, as an input, the Data will be merged into another Data, which is
an output.

This section lists and describes the types that rbxmk recognizes by default.

Type                                        | Description
--------------------------------------------|------------
[Instances](#user-content-instances-type)   | A list of instances.
[Instance](#user-content-instance-type)     | A single instance.
[Properties](#user-content-properties-type) | A table of property names mapped to values.
[Property](#user-content-property-type)     | A single property name mapped to a value.
[Value](#user-content-value-type)           | A single value with a certain type.
[Stringlike](#user-content-stringlike-type) | A string-like value.
[Region](#user-content-region-type)         | A subsection of a string-like value.
[Delete](#user-content-delete-type)         | A special type that removes values it is merged with.

### Merging overview
The following table provides an overview of which types can be merged.

| <sub>in</sub>╲<sup>out</sup> | Instances  | Instance | Properties | Property | Value | Stringlike | Region | Delete |
|------------------------------|------------|----------|------------|----------|-------|------------|--------|--------|
| Instances                    | YES        | YES      | NO         | NO       | NO    | NO         | NO     | NO     |
| Instance                     | YES        | YES      | NO         | COND     | NO    | NO         | NO     | NO     |
| Properties                   | YES        | YES      | YES        | YES      | NO    | NO         | COND   | NO     |
| Property                     | YES        | YES      | YES        | YES      | YES   | YES        | YES    | NO     |
| Value                        | NO         | NO       | NO         | YES      | YES   | COND       | COND   | NO     |
| Stringlike                   | NO         | NO       | NO         | YES      | YES   | YES        | YES    | NO     |
| Region                       | COND       | COND     | COND       | YES      | YES   | YES        | YES    | NO     |
| Delete                       |  YES       | YES      | YES        | YES      | YES   | YES        | YES    | NO     |

- `NO`: The types canot be merged.
- `YES`: The types can be merged.
- `COND`: The types can be merged under certain conditions.

### `Instances` type
`Instances` represents a list of Roblox instances.

#### Drilling
Drilling into an `Instances` type selects an `Instance` type.

A single Reference string is processed. This string contains a number of
identifiers, separated by `.` characters. Each successive identifier refers to a
child instance in the instance of the previous identifier.

An identifier can be a sequence of letters, numbers, and underscores, which
doesn't start with a number. This is used to match the Name property of an
instance. An identifier may also be a positive integer number, which selects the
*n*th child instance, starting at 0.

For example, `Workspace.0.Part` selects an instance named `Part`, which is the
first child of an instance named `Workspace`, which is in the list.

#### Merging
As an input, `Instances` can be merged into the following types:

Output type | Result
------------|-------
Instances   | Each input instance is append to the output list.
Instance    | Each input instance is added as a child to the output instance.

### `Instance` type
`Instance` represents a single Roblox instance.

#### Drilling
Drilling into an `Instance` type can select either a `Property` or `Properties`
type.

A single Reference string is processed. This string can be the name of a
property in the instance, which selects the property as a `Property` type. The
string may also be the literal `*`, which selects all properties in the instance
as a `Properties` type.

#### Merging
As an input, `Instance` can be merged into the following types:

Output type | Result
------------|-------
Instances   | The instance is appended to the list.
Instance    | The input instance is added as a child to the output instance.
Property    | If the property value is of the Reference type, then the reference is set to the instance.

### `Properties` type
`Properties` represents a table of property name mapped to Roblox values.

#### Drilling
Drilling into a `Properties` type selects a `Property` type.

A single reference string is processed. This string can be the name of a
property, which selects the property in the table.

#### Merging
As an input, `Properties` can be merged into the following types:

Output type | Result
------------|-------
Instances   | Each property is set to each instance, only if the output property is nil, or the value types match.
Instance    | Each property is set the instance, only if the output property is nil, or the value types match.
Properties  | Each input property is set to the output table, only if the output property is nil, or the value types match.
Property    | Uses the output name to select an input property, and sets the result to the output property.
Region      | The region must derive from a property. Uses the output name to select an input property (must be string-like), and sets it to the region.

### `Property` type
`Property` represents a property name mapped to a Roblox value.

#### Drilling
Drilling into a `Property` type selects a `Region` type, but only of the
property value is string-like.

The [Region drilling](#user-content-region-drilling) section describes how to
drill to select a Region.

#### Merging
As an input, `Property` can be merged into the following types:

Output type | Result
------------|-------
Instances   | The property is set to each instance, only if the output property is nil, or the value types match.
Instance    | The property is set to the instance, only if the output property is nil, or the value types match.
Properties  | The property is set to the table, only if the output property is nil, or the value types match.
Property    | The input property is set to the output property, only if the output property is nil, or the value types match.
Value       | If the property type matches the value type, then the property replaces the value.
Stringlike  | If the property is string-like, then it replaces the output.
Region      | If the property is string-like, then it is set to the region.

### `Value` type
`Value` represents a Roblox value with a Roblox type (e.g. bool, string,
Vector3, CFrame, etc).

#### Drilling
Drilling into a `Value` type selects a `Region` type, but only of the value is
string-like.

The [Region drilling](#user-content-region-drilling) section describes how to
drill to select a Region.

#### Merging
As an input, `Value` can be merged into the following types:

Output type | Result
------------|-------
Property    | If the value type matches the property type, then the value is assigned to the property.
Value       | If the value types match, the input values replaces the output value.
Stringlike  | If the value is string-like, then it replaces the output.
Region      | If the value is string-like, then it is set to the region.

### `Stringlike` type
`Stringlike` represents any value that is string-like.

#### Drilling
Drilling into a `Stringlike` type selects a `Region` type..

The [Region drilling](#user-content-region-drilling) section describes how to
drill to select a Region.

#### Merging
As an input, `Stringlike` can be merged into the following types:

Output type | Result
------------|-------
Property    | If the property is string-like or nil, then the string is assigned to the property.
Value       | If the value is string-like, then the string replaces the value.
Stringlike  | The input string replaces the output string.
Region      | The string is set to the region.

### `Region` type
`Region` represents a slice of a string-like value.

#### Drilling
The `Region` type cannot be drilled into.

#### Merging
As an input, `Region` can be merged depending on where the region originates. If
it originates from a Property, then it is merged in the same way as a
[Property](#user-content-merging-3). Otherwise, it is merged in the same way as
a [Stringlike](#user-content-merging-5).

#### Region drilling
Several types can be drilled into to select a Region. The values of these types
must be string-like. This section describes how to drill, and what is selected.

When drilling, a single Reference string is processed. This string contains a
number of identifiers, separated by `.` characters. If the string ends with a
`+` character, then the Region will be in Append mode.

Each successive identifier refers to a subregion in the region of the previous
identifier. An identifier can be a sequence of letters and numbers. This matches
the name of a region.

A **region** can be defined as a slice of a string, delimited by **tags**. A tag
has specific markers, and is designed to be contained within a Lua comment,
which means it shares a similar syntax:

- `--@Name`: A start tag.
- `--@/Name`: An end tag.
- `--[[@Name]]`: An inline start tag.
- `--[[@/Name]]`: An inline end tag.

A region can have sub-regions, which is indicated by opening a region with a
start tag before the previous start tag is closed.

```lua
--@Region
--@SubRegion
--@/SubRegion
--@/Region
```

Closing a region will also close all subregions.

```lua
--@Region
--@SubRegion
--@/Region
```

An unmatched end tag is interpreted as a region of size 0.

```lua
--@Region
--@/EmptySubRegion
--@/Region
```

A region may have any content between the tags. When a Region type is merged
into as an output, this content is replaced, and the tags are removed. If the
Region is in Append mode, then the input is appended to the content, and the
tags are preserved.

Normal tags and inline tags select different parts of the region.

- Normal start tag: The selection begins after the first newline after the tag.
- Normal end tag: The selection ends before the tag.
- Inline start tag: The selection before just after the tag.
- Inline end tag: The selection ends just before the tag.

Note: for regions selected by normal tags, it is good practice for merged
content to end with a newline.

### `Delete` type
`Delete` represents the absence of a value; when merged as an input, it removes
the output value.

#### Drilling
The `Delete` type cannot be drilled into.

#### Merging
As an input, `Delete` can be merged into the following types:

Output type | Result
------------|-------
Instances   | Removes all instances from the list.
Instance    | Sets the parent of the instance to nil.
Properties  | Removes the all properties from the table.
Property    | Removes the property from the table.
Value       | Replaces the value with nil.
Stringlike  | Replaces the string with an empty string.
Region      | Replaces the content of the region with nothing.

## Formats
A Format describes how to decode raw data into a known [data type](#user-content-data-types),
and how to encode data types into raw data.

When encoding or decoding, a Format may receive extra information from a
[Scheme](#user-content-schemes), which it can use for context.

This section lists and describes the formats that rbxmk includes by default.

Format Name                                      | Decoded type             | Encodable types
-------------------------------------------------|--------------------------|--------
[rbxl](#user-content-roblox-formats)             | `Instances`              | `Instances`, `Instance`
[rbxlx](#user-content-roblox-formats)            | `Instances`              | `Instances`, `Instance`
[rbxm](#user-content-roblox-formats)             | `Instances`              | `Instances`, `Instance`
[rbxmx](#user-content-roblox-formats)            | `Instances`              | `Instances`, `Instance`
[lua](#user-content-lua-formats)                 | `Value<ProtectedString>` | `Instances`*, `Instance`**, `Stringlike`
[script.lua](#user-content-lua-formats)          | `Instance<Script>`       | `Instances`*, `Instance`**, `Stringlike`
[localscript.lua](#user-content-lua-formats)     | `Instance<LocalScript>`  | `Instances`*, `Instance`**, `Stringlike`
[modulescript.lua](#user-content-lua-formats)    | `Instance<ModuleScript>` | `Instances`*, `Instance`**, `Stringlike`
[txt](#user-content-text-formats)                | `Value<String>`          | `Stringlike`
[bin](#user-content-text-formats)                | `Value<BinaryString>`    | `Stringlike`
[properties.json](#user-content-property-formats)| `Properties`             | `Instances`* `Instance`, `Properties`, `Property`
[properties.xml](#user-content-property-formats) | `Properties`             | `Instances`* `Instance`, `Properties`, `Property`

- `*`: Selects the first instance in the list.
- `**`: Instance must be script-like (Script, LocalScript, ModuleScript)
- `Stringlike`: Includes any type that can be converted to a Stringlike

### Roblox formats
These formats describe official Roblox file formats. Each format contains a list
of Roblox instances.

Name    | Description
--------|------------
`rbxl`  | A Roblox place file.
`rbxlx` | A Roblox place file in XML format.
`rbxm`  | A Roblox model file.
`rbxmx` | A Roblox model file in XML format.

### Lua formats
These formats describe Lua script files.

When using the `file` scheme, the formats that decode into script-like instances
use the base of the file name to set the Name property of the script.

Name               | Description
-------------------|------------
`lua`              | A Lua script unassociated with any type of script instance.
`script.lua`       | A Lua script decoded as a Script instance.
`localscript.lua`  | A Lua script decoded as a LocalScript instance.
`modulescript.lua` | A Lua script decoded as a ModuleScript instance.

### Text formats
These formats describe generic text or data.

Name  | Description
------|------------
`txt` | A file containing text.
`bin` | A file containing binary data.

### Property formats
These formats describe mappings of property names to values.

Name              | Description
------------------|------------
`properties.json` | A file containing properties in JSON format.
`properties.xml`  | A file containing properties in XML format.

## Schemes
A Scheme describes how to retrieve raw data from a resource.

This section lists and describes the schemes that rbxmk includes by default.

Scheme                                    | Type          | Description
------------------------------------------|---------------|------------
[file](#user-content-file-scheme)         | input, output | Accesses a file on the file system.
[http, https](#user-content-http-scheme)  | input, output | Accesses the web from a URL.
[generate](#user-content-generate-scheme) | input         | Generates a value from scratch.

### File scheme
The `file` scheme is used to refer to files on the operating system. It is
defined for both inputs and outputs.

The syntax is simply a path to a file on the operating system.

```
file://C:/Users/user/projects/project/file.rbxl
file:///home/user/projects/project/file.rbxl
```

Because the file scheme is the default, the scheme portion may be omitted from
the reference.

```
C:/Users/user/projects/project/file.rbxl
/home/user/projects/project/file.rbxl
```

The Format of the selected file, if not provided, is determined by the file
extension. The extension name is the same as the name of the Format.

When encoding or decoding, the file name is passed to the format as context.

### HTTP scheme
The `http` and `https` schemes retrieve files using the HTTP protocol. It is
defined for both inputs and outputs.

The syntax is a standard URL.

```
http://www.example.com/path/to/file?etc
```

Drilling into an output is disabled for this scheme, because it may not be
possible to receive data from the output location.

This scheme cannot determine the Format automatically, so it must be provided
explicitly.

#### Roblox asset ID scheme
The `rbxassetid` scheme downloads and uploads asset files from the Roblox
website. It is defined for both inputs and outputs.

The syntax is the ID of the asset:

```
rbxassetid://1818
```

An input node with this scheme will download the specified asset from the
website. Authentication is required for private assets, but not for public
assets.

An output node with this scheme will update the specified asset on the website.
The output node can be drilled into, in which case the asset is first
downloaded. Note that this method cannot be used to create new assets.

When uploading to or downloading from the website, the `user` argument of an
input or output node can be used when authentication is required. The argument
refers to a user configured with the [`rbxauth`](#user-content-config-rbxauth)
configuration option. If no user has been specified or configured, then the user
will be prompted to login when the node is mapped. This session will persist
until the program ends or the user is explicitly logged out. See the
[input](#user-content-input-function) function for details on the `user`
argument.

This scheme cannot determine the Format automatically, so it must be provided
explicitly. The format should correspond to the type of the target asset.

### Generate scheme
The `generate` scheme generates Data from scratch. It is defined for only for
inputs.

The syntax is a word indicating the type of Data to be generated.

Reference                                              | Data type
-------------------------------------------------------|----------
[`generate://Instance`](#user-content-instance-syntax) | `Instances`
[`generate://Property`](#user-content-property-syntax) | `Properties`
[`generate://Value`](#user-content-value-syntax)       | `Value`

This scheme processes the next string in the Reference. This string describes
the data, and has a specific syntax for each type of data. In general,
whitespace is ignored.

#### Instance syntax
Specifying `Instance` generates a list of instances. Each instance is described
by a class name, followed by curly brackets enclosing the content of the
instance. Each instance is separated by a semi-colon.

```
ClassName{}; ClassName{}
```

The content of an instance consists of a number of items, each being one of 3
types:

- Instance
- Property
- Meta property

Each item is separated by a semi-colon.

##### Instance item
An instance item describes a child instance of the current instance. It has the
same syntax as an instance.

```
Instance{ChildInstance{}; ChildInstance{}}
```

##### Property item
A property item describes a property of the current instance. It has the
following syntax:

```
PropertyName : PropertyType = PropertyValue
```

If an API has been specified, then the type can be omitted. The API will be used
to discover the type automatically.

```
PropertyName = PropertyValue
```

A property value consists of a comma-separated list of bools, numbers, and
strings. Each element in the list specifies a component of the value. For
example, 1 string describes a String value, or 3 numbers describe a Vector3
value:

```
Part{Name:string = "Part"; Size:Vector3 = 4,1,2}
```

##### Meta property item
A meta property item describes an internal value of the current instance. It
consists of name followed by a value enclosed in parentheses.

```
MetaProp(value)
```

There are two meta properties defined:

```
Instance{
	IsService(true);
	Reference("RBX0123456789ABCDEF");
}
```

`IsService` is a bool that determines whether the instance is meant to be loaded
as a service.

`Reference` is a string that may be used by properties of type "Reference".
Properties cannot refer to instances outside of the generated content.

```
ObjectValue{
	Name:string = "Value";
	Value:Reference = "part";
}
Part{
	Reference(part);
	Name:string = "Part";
	Size:Vector3 = 4,1,2;
};
```

#### Property syntax
Specifying `Property` generates a set of properties, where a property name is
mapped to a Value. It shares the same syntax as a property item in the Instance
syntax:

```
PropertyName : PropertyType = PropertyValue
```

Each property is separated by a semi-colon. The property type is not optional.

#### Value syntax
Specifying `Value` generates a single Value. It shares a similar syntax as
property items in the Instance syntax:

```
PropertyType = PropertyValue
```

#### Example
**Instance**
```
Workspace{IsService(true);
	Name: string = "Workspace";
	Model{
		Name: string           = "Model";
		PrimaryPart: Reference = "primary";
		Part{Reference(primary);
			Name: string        = "Part";
			Anchored: bool      = true;
			Transparency: float = 0.5;
			Position: Vector3   = 0,10,0;
			Size: Vector3       = 4,1,2;
		};
	};
};
ServerScriptStorage{IsService(true);
	Name: string = "ServerScriptService";
	Script{
		Name: string            = "Hello";
		Source: ProtectedString = "print(\"Hello world!\")";
	};
};
```

**Property**
```
Name: string        = "Part";
Anchored: bool      = true;
Transparency: float = 0.5;
Position: Vector3   = 0,10,0;
Size: Vector3       = 4,1,2;
```

**Value**
```
Vector3 = 4,1,2
```

In cases where a name is required, either an identifier or a quoted string can
be used. An identifier is a sequence of letters, digits, and underscores, which
doesn't begin with a digit. A string is delimited with either single or double
quotes, and may use a backslash to escape characters. Whitespace is preserved
within a string.

## Filters
A Filter is simply a function that receives arguments and returns values. In
general, they are a way to specify procedures with more specific behaviors,
without polluting the global environment. Usually, a filter receives inputs or
outputs and transforms them in some way, hence the name "filter".

This section lists and describes the filters that rbxmk includes by default.

Name                                          | Description
----------------------------------------------|------------
[minify](#user-content-minify-filter)         | Shrink the content of a Lua script.
[preprocess](#user-content-preprocess-filter) | Run the preprocessor on some text.
[region](#user-content-region-filter)         | Replace regions of a Lua script.
[unminify](#user-content-unminify-filter)     | Expand the content of a Lua script.

### Minify filter
`data = rbxmk.filter{"minify", data}`

`minify` uses [lua-minify](https://github.com/stravant/lua-minify) to minify a
Lua script. It receives a single Data, and returns the modified result.

Data can be one of the following types:

Type           | Description
---------------|------------
`Instances`    | Each script-like instance is modified.
`Instance`     | If the instance is script-like, then its Source property is modified, if it exists.
`Value`        | Only if the value is string-like.
`string` (Lua) | Modifies the string.

### Preprocess filter
`data = rbxmk.filter{"preprocess", data}`

`preprocess` runs a preprocessor on a string-like value. It receives a single
Data, and returns the modified result.

`preprocess` receives the same Data types as the
[`minify`](#user-content-minify-filter) filter.

Within a string, a preprocessor can be signaled with the following syntax: when
the first character of a Lua comment is a `#` character, the rest of the comment
is treated as Lua. This includes long comments.

The long comment form has several specific features:

- `--[[#explist]]`: Evalutes a Lua expression list and outputs each result as a
  string. Note that the results are concatenated directly. Nil values output
  nothing.
- `--[[#chunk]]`: Evaluates the chunk and outputs nothing.
- `--[[#'#']]`: Outputs a literal `#` character to the output file.

Each successive preprocessor creates a continuous chunk of Lua code. Any
non-preprocessor portion of the input file is outputted directly if it is
reached in the Lua code. For example, with the following input file:

```text
--#if condition then
Hello
--#else
Goodbye
--#end
```

If `condition` is true, then the content of the output file will be "Hello".
Otherwise, it will be "Goodbye".

Any valid Lua syntax may be used. This example preprocesses a Lua file to unroll
a loop:

```lua
--#for i = 1, 4 do
print(--[[#i]])
--#end
```

This evaluates to the following file:

```lua
print(1)
print(2)
print(3)
print(4)
```

The Lua code runs in its own isolated environment, which exists only for the
duration of the preprocess filter call. The environment contains the same
[standard library](#user-content-standard-library) as rbxmk scripts (the rbxmk
library is not included).

Also included are the variables specified by the
[`--define`](USAGE.md#user-content-command-options) command option and the
[`rbxmk.configure`](#user-content-configure-function) function. These variables
cannot be modified by the preprocessor. Note that variables defined with the
command option take precedence over variables defined by `configure`.

For example, the `condition` variable can be defined via the shell:

```shell
rbxmk -f script.lua --define condition:true
```

Or from a script:

```lua
rbxmk.configure{define={condition=true}}
```

Finally included is the `_put` function. Any arguments passed to this function
are converted to strings and written to the output file. Nil values output
nothing.

### Region filter
`data = rbxmk.filter{"region", data, string, data}`

`region` works like the Region drill; it searches for a region of text within a
script, and replaces it.

The first argument is the value to modify. This can be a script-like instance or
any string-like value. The second argument is the name of the region to select,
which has the same rules a Reference string given to the Region drill. The third
argument is any value capable of being merged into a region.

`region` returns the first argument after it has been modified. If the region
could not be found, then the first argument is returned unchanged.

### Unminify filter
`data = rbxmk.filter{"unminify", data}`

`unminify` uses [lua-minify](https://github.com/stravant/lua-minify) to unminify
a Lua script. It receives a single Data, and returns the modified result.

`unminify` receives the same Data types as the
[`minify`](#user-content-minify-filter) filter.
