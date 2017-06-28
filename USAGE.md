# Usage details

This document provides details on how `rbxmk` works, for regular usage.

<table>
<thead><tr><th>Table of Contents</th></tr></thead>
<tbody><tr><td><ol>
	<li><a href="#user-content-lua-environment">Lua environment</a></li>
	<li><a href="#user-content-lua-functions-table">Lua functions table</a></li>
	<li><a href="#user-content-resolve-chain">Resolve chain</a><ol>
		<li><a href="#user-content-input-procedure">Input procedure</a></li>
		<li><a href="#user-content-output-procedure">Output procedure</a></li>
	</ol></li>
	<li><a href="#user-content-reference">Reference</a></li>
	<li><a href="#user-content-data">Data</a><ol>
		<li><a href="#user-content-format-types-table">Format types table</a></li>
	</ol></li>
	<li><a href="#user-content-schemes">Schemes</a><ol>
		<li><a href="#user-content-file">file</a></li>
		<li><a href="#user-content-httphttps">http/https</a></li>
		<li><a href="#user-content-generate">generate</a><ol>
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
	<li><a href="#user-content-formats">Formats</a><ol>
		<li><a href="#user-content-roblox-formats">Roblox formats</a></li>
		<li><a href="#user-content-lua-formats">Lua formats</a></li>
		<li><a href="#user-content-text-formats">Text formats</a></li>
		<li><a href="#user-content-property-formats">Property formats</a></li>
	</ol></li>
</ol></td></tr></tbody>
</table>

## Lua environment

[Lua](https://lua.org) scripts are used to perform actions. Scripts are run in
a stripped-down environment; none of the regular base functions and libraries
are loaded. Instead, a set of new functions is made available. Each of these
functions accept a table as their only argument.

Lua has the following bit of syntax sugar to support this: if a constructed
table is the only argument to a function, then the function's parentheses may
be omitted:

```lua
func({arg1, arg2})
func{arg1, arg2}
-- Both calls are equivalent
```

Henceforth, an "argument", in the context of a Lua function, will refer to a
value within this kind of table.

Because of how tables work, there can be two kinds of arguments: named and
unnamed.

```lua
-- Unnamed arguments
func{value, value}

-- Named arguments
func{arg1=value, arg2=value}
```

## Lua functions table

The following functions are available:

Name                                     | Description
-----------------------------------------|------------
[`error`](#user-content-error)           | Throw an error.
[`exit`](#user-content-exit)             | Force the program to exit.
[`filter`](#user-content-filter)         | Transform nodes.
[`getenv`](#user-content-getenv)         | Get the value of an environment variable.
[`globalapi`](#user-content-globalapi)   | Set an API as the default for all functions.
[`input`](#user-content-input)           | Create an input node.
[`load`](#user-content-load)             | Load and execute a script.
[`loadapi`](#user-content-loadapi)       | Load an API file.
[`map`](#user-content-map)               | Map one or more inputs to one or more outputs.
[`output`](#user-content-output)         | Create an output node.
[`pcall`](#user-content-pcall)           | Call a function in protected mode.
[`print`](#user-content-print)           | Print values to stdout.
[`printf`](#user-content-printf)         | Print a formatted string to stdout.
[`sprintf`](#user-content-sprintf)       | Return a formatted string.
[`type`](#user-content-type)             | Return the type of a value as a string.

### input

`node = input{format=string, api=nil, ...string}`

`input` creates an input node. The arguments specify the
[reference](#user-content-references) to the input. The `format` argument
forces the file format, if needed.

The optional `api` argument specifies an API value to enhance the handling of
instances and properties. Specifying a non-nil API overrides the default API.

An input node returned by `input` has two methods:

Name                                           | Description
-----------------------------------------------|------------
[`CheckInstance`](#user-content-checkinstance) | Check if an instance exists.
[`CheckProperty`](#user-content-checkproperty) | Check if a property exists.

#### CheckInstance

`bool = inputNode:CheckInstance{string...}`

`CheckInstance` returns whether an instance exists within the node. It may
also be used to check whether an instance has a property. The arguments to
`CheckInstance` are a reference, which can drill into instances.

#### CheckProperty

`bool = inputNode:CheckProperty{string}`

`CheckProperty` returns whether a property of the given name exists within the
node.

### output

```lua
node = output{format=string, api=nil, ...string}
```

`output` creates an output node. The arguments specify the
[reference](#user-content-references) to the output. The `format` argument
forces the file format, if needed.

The optional `api` argument specifies an API value to enhance the handling of
instances and properties. Specifying a non-nil API overrides the default API.

### map

`map{...node}`

`map` maps one or more input nodes to one or more output nodes. Either kind of
node may be passed to `map`.

Nodes are mapped in the order they are received. That is, inputs are gathered
in one list, and outputs are gathered in another. Then each node in the input
list is mapped to each node in the output list, in order. For example:

```lua
A,B = input{...},input{...}
X,Y = output{...},output{...}
map{A, X, B, Y}
-- 1: A -> X
-- 2: A -> Y
-- 3: B -> X
-- 4: B -> Y
```

### filter

`... = filter{api=nil, string, ...}`

`filter` transforms nodes. The first argument specifies the filter name.
Subsequent arguments and return values depend on the filter. See
[Filters](#user-content-filters) for a list of filters and their arguments.

The optional `api` argument specifies an API value to enhance the handling of
instances and properties. Specifying a non-nil API overrides the default API.
The API is used only by applicable filters.

### load

`... = load{string, ...}`

`load` executes a file as a script. The first argument is the file name.
Subsequent arguments are passed to the script (which may be received with the
`...` operator). `load` returns any values returned by the script.

### type

`string = type{value}`

`type` returns the type of the given value as a string. In addition to the
regular Lua types, the following types are available:

- `input`: an input node.
- `output`: an output node.

### error

`error{string}`

`error` throws an error, with the first argument as the error message.

### exit

`exit{string}`

`exit` forces the program to exit. An optional message can be given, which
will be passed to the program.

### pcall

`pcall{function, ...}`

`pcall` calls a function with the given arguments. If an error occurs, `pcall`
returns false, followed by the error message. If no errors occur, `pcall`
return true, followed by any values returned by the called function.

### getenv

`string = getenv{string}`

`getenv` returns the value of an environment variable of the given name.
Returns nil if the variable is not present.

### print

`print{...}`

`print` receives a number of values and writes them to standard output.
`print` follows the same rules as
[fmt.Println](https://golang.org/pkg/fmt/#Print).

### printf

`printf{string, ...}`

`printf` receives a number of values, formats them according to the first
argument, and writes the result to standard output. `printf` follows the same
rules as [fmt.Printf](https://golang.org/pkg/fmt/#Printf).

### sprintf

`string = sprintf{string, ...}`

`sprintf` receives a number of values, formats them according to the first
argument, and returns the resulting string. `sprintf` follows the same rules
as [fmt.Sprintf](https://golang.org/pkg/fmt/#Sprintf).

### loadapi

`api = loadapi{string}`

`loadapi` receives a path to a file containing an API dump. `loadapi` returns
a value representing the API. This can be passed to certain functions to
enhance how the function handles instances and properties.

### globalapi

`globalapi{api}`

`globalapi` sets the default API value to be used by all applicable functions.
Several functions have an `api` argument, which can be used to override the
default API for that call.

Initially, the default API is nil.

## Resolve chain

When creating or mapping an input or output node, rbxmk has a procedure that
chains together predefined components in order to resolve the node. In its
most basic form, the procedure looks like this:

```
Scheme -> Format -> Drills
```

1. Scheme retrieves a file containing raw data.
2. Format turns the file into data of a known type.
3. Drills select data within the data.

The exact procedures for inputs and outputs differ, but still follow this
sequence overall.

### Input procedure

1. Determine the Scheme.
	- Uses the first string of the [Reference](#user-content-reference).
2. Use Scheme to get a file from a specified location.
3. Determine the Format.
	- Provided by the `format` argument, or guessed, depending on the Scheme.
4. Decode the file using Format.
	- Format returns Data of a known type.
5. Drill into Data using Format's input Drills.
	- Format has a number of drills that can operate on the returned Data.

### Output procedure

1. Determine the Scheme.
	- Uses the first string of the [Reference](#user-content-reference).
2. Use Scheme to get current state of the output from a specified location;
   may be skipped if not applicable to the Scheme.
	1. Use Scheme to get a file from a specified location.
	2. Determine the Format.
		- Provided by the `format` argument, or guessed, depending on the
		  Scheme.
	3. Decode the file using Format.
		- Format returns Data of a known type.
	4. Drill into Data using Format's output Drills.
		- Format has a number of drills that knows how to operate on the type
		  of the returned Data.
3. Merge input Data into decoded Data using Format.
	- Format has a Merge function that knows how to operate on the Data types.
4. Encode resulting Data into a file using Format.
5. Use Scheme to write the file to the specified location.

## Reference

A Reference is a list of strings used to specify a piece of Data. It is passed
into an input or output resolve chain. Each step of the chain processes a
number of strings, then passes the remaining strings to the next step.

The first string in the reference is a URL pointing to a location that
contains a file. This URL is resolved depending on the scheme (e.g. `file://`,
`http://`, etc). If the scheme part (`scheme://`) of the URL is omitted, the
the URL is assumed to be of the `file` scheme.

While schemes deal primarily with the first string, they may process any
number of strings.

After the format has been decoded, the remaining Reference is passed to each
of the format's Drills. Each drill works like a step in the chain, processing
a number of strings, and returning the remaining strings to the next drill.

Here's an example. Each unnamed argument to the `input` function is a string
of a Reference.

```
input{"file://place.rbxl", "Workspace.Model.Script", "Source"}
```

1. From the first string, the scheme is determined to be `file`.
2. The `file` scheme reads the `place.rbxl` from the file system, guessing the
   format as `rbxl` based on the file's extension.
3. Using the `rbxl` format, the contents of the file are decoded into some
   kind of Data (in this case, a list of Instances).
4. The next string is used by the format's first drill, which selects a
   descendant instance within a tree of instances. In this case, the "Script"
   instance is selected.
5. The string after that is used by the format's second drill, which selects a
   property within an instance. In this case, the value of the script's
   "Source" property is selected.

## Data

Data is a value with a known type. This document describes Data types in the
following ways:

- `T`: A value of type T.
- `[]T`: A list of values of type T.
- `[TK]TV`: A map containing values of type TK mapped to values of type TV.

The types known by the default set of Schemes, Formats, and Filters are
described as follows:

Type       | Description
-----------|------------
`Instance` | A Roblox instance.
`Value`    | A value of a Roblox type (e.g. bool, string, float, Vector3, CFrame).
`string`   | A string of text.

Recognized sub-types:

Type              | Implements
------------------|-----------
`Script`          | `Instance`
`LocalScript`     | `Instance`
`ModuleScript`    | `Instance`
`String`          | `Value`
`BinaryString`    | `Value`
`ProtectedString` | `Value`

### Format types table

Each format encodes and decodes between raw data and Data of specific types.
The following table shows the type of Data each format decodes into, and types
of Data that the format can encode.

Format           | Decodes           | Encodes
-----------------|-------------------|--------
rbxl             | `[]Instance`      | `[]Instance`, `Instance`
rbxlx            | `[]Instance`      | `[]Instance`, `Instance`
rbxm             | `[]Instance`      | `[]Instance`, `Instance`
rbxmx            | `[]Instance`      | `[]Instance`, `Instance`
lua              | `ProtectedString` | `[]Instance*`, `Script-like`, `String-like`
script.lua       | `Script`          | `[]Instance*`, `Script-like`, `String-like`
localscript.lua  | `LocalScript`     | `[]Instance*`, `Script-like`, `String-like`
modulescript.lua | `ModuleScript`    | `[]Instance*`, `Script-like`, `String-like`
txt              | `String`          | `String-like`
bin              | `BinaryString`    | `String-like`
json             | `[string]Value`   | `[]Instance*`, `Instance`, `[string]Value`
xml              | `[string]Value`   | `[]Instance*`, `Instance`, `[string]Value`

- `*`: Selects the first Instance in the list.
- `Script-like`: An `Instance` of the class `Script`, `LocalScript`, or
  `ModuleScript`.
- `String-like`: A `Value` of the type `String`, `BinaryString`,
  `ProtectedString`.


## Schemes

### file

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

The format of the selected file, if not provided, is determined by the file
extension.

The `file` scheme passes the file name to the format, which may be used when
decoding or encoding the file.

### http/https

The `http/https` scheme retrieves files using the HTTP protocol. The syntax is
a standard URL.

```
http://www.example.com/path/to/file
```

Drilling into an output is disabled for this scheme, because it may not be
possible to receive data from the output location.

The format of the selected file, if not provided, is determined by the MIME
type given by the response.

MIME              | Format
------------------|-------
(not implemented) | --

### generate

The `generate` scheme generates Data from scratch. The syntax is a word indicating the type of Data to be generated.

Reference             | Data type
----------------------|----------
`generate://Instance` | `[]Instance`
`generate://Property` | `[string]Value`
`generate://Value`    | `Value`

The next string in the Reference describes the data. This string has a
specific syntax for each type of data. In general, whitespace is ignored.

#### Instance syntax

Specifying `Instance` generates a list of instances. Each instance is
described by a class name, followed by curly brackets enclosing the content of
the instance. Each instance is separated by a semi-colon.

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

An instance item describes a child instance of the current instance. It has
the same syntax as an instance.

```
Instance{ChildInstance{}; ChildInstance{}}
```

##### Property item

A property item describes a property of the current instance. It has the
following syntax:

```
PropertyName : PropertyType = PropertyValue
```

If an API has been specified, then the type can be omitted. The API will be
used to discover the type automatically.

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

`IsService` is a bool that determines whether the instance is meant to be
loaded as a service. `Reference` is a string that may be used by properties of
type "Reference". Properties cannot refer to instances outside of the
generated content.

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
mapped to a Value. It shares the same syntax as a property item in the
Instance syntax:

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
be used. An identifier is a sequence of letters, digits, and underscores,
which doesn't begin with a digit. A string is delimited with either single or
double quotes, and may use a backslash to escape characters. Whitespace is
preserved within a string.

## Formats

### Roblox formats

These formats describe official Roblox file formats.

Name  | Extension | Description
------|-----------|------------
RBXL  | `rbxl`    | Roblox Place
RBXLX | `rbxlx`   | Roblox Place XML
RBXM  | `rbxm`    | Roblox Model
RBXMX | `rbxmx`   | Roblox Model XML

### Lua formats

These formats describe Lua script files.

When using the `file` scheme, the formats that decode into script instances
use the base of the file name to set the Name property of the script.

Name             | Extension          | Description
-----------------|--------------------|------------
Lua              | `lua`              | A Lua script unassociated with any type of script instance.
Lua Script       | `script.lua`       | A Lua script decoded as a Script instance.
Lua LocalScript  | `localscript.lua`  | A Lua script decoded as a LocalScript instance.
Lua ModuleScript | `modulescript.lua` | A Lua script decoded as a ModuleScript instance.

### Text formats

These formats describe generic text or data.

Name   | Extension | Description
-------|-----------|------------
Text   | `txt`     | A file containing text.
Binary | `bin`     | A file containing binary data.

### Property formats

These formats describe mappings of property names to values.

Name | Extension | Description
-----|-----------|------------
JSON | `json`    | A file containing properties in JSON format.
XML  | `xml`     | A file containing properties in XML format.



