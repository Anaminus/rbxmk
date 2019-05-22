# Usage
This document provides an overview of how rbxmk is used. For a detailed
reference, see [DOCUMENTATION.md](DOCUMENTATION.md).

<table>
<thead><tr><th>Table of Contents</th></tr></thead>
<tbody><tr><td><ol>
	<li><a href="#user-content-command-options">Command options</a></li>
	<li><a href="#user-content-lua-environment">Lua environment</a></li>
	<li><a href="#user-content-workflow">Workflow</a></li>
	<li><a href="#user-content-references">References</a></li>
	<li><a href="#user-content-schemes">Schemes</a></li>
	<li><a href="#user-content-formats">Formats</a></li>
	<li><a href="#user-content-data-types">Data types</a></li>
	<li><a href="#user-content-functions">Functions</a><ol>
		<li><a href="#user-content-filters">Filters</a></li>
		<li><a href="#user-content-path">Path</a></li>
		<li><a href="#user-content-loading">Loading</a></li>
		<li><a href="#user-content-deletion">Deletion</a></li>
		<li><a href="#user-content-other-functions">Other functions</a></li>
	</ol></li>
</ol></td></tr></tbody>
</table>

## Command options
Since almost all actions are done in Lua, the `rbxmk` command has only a few
options:

```
rbxmk [ -h ] [ -a VALUE ] [ -d NAME:VALUE ] [ FILE ]
```

Option           | Value        | Description
-----------------|--------------|------------
`-h`, `--help`   |              | Display a help message.
`-a`, `--arg`    | `VALUE`      | Specify an argument to be passed to the script.
`-d`, `--define` | `NAME:VALUE` | Define a variable to be used by the [preprocessor](DOCUMENTATION.md#user-content-preprocess-filter).

rbxmk receives `FILE`, which is the file path of a script to be executed.

```shell
rbxmk script.lua
```

If `FILE` is not specified, then the script is read from stdin.

```shell
echo print("hello world") | rbxmk
```

The `-a` option is used to pass arguments to the script. It can be given more
than once to provide multiple values.

```shell
rbxmk script.lua -a hello -a world -a 123
```

A script can read these arguments with the `...` operator.

```lua
Arg1, Arg2, Arg3 = ... -- "hello", "world", 123
AllArgs = {...}
```

The `-d` option can also be given more than once. For this option, `NAME` must
be a string of digits, letters, and underscores, which doesn't start with a
digit. Note that variables defined by this command option will take precedence
over variables defined elsewhere.

```shell
rbxmk script.lua --define DEBUG:true --define LOGLEVEL:3
```

For either option, `VALUE` is parsed into a Lua value. This may be a number,
"true" or "false" for bool, or "nil" for nil. Any other value is read as a
string.

## Lua environment
[Lua](https://lua.org) scripts are used to perform actions. Scripts are run in a
modified environment, where most of the standard functions and libraries are
available. A complete list of included items can be found in
[DOCUMENTATION.md](DOCUMENTATION.md#user-content-standard-library).

Also included is the `rbxmk` library. Each function in this library accepts a
table as its only argument. Lua has the following bit of syntax sugar to support
this: if a constructed table is the only argument to a function, then the
function's parentheses may be omitted:

```lua
func({arg1, arg2})
func{arg1, arg2}
-- Both calls are equivalent
```

Henceforth, an "argument", in the context of these functions, will refer to a
value within this kind of table. Unnamed arguments are referred to in sequence
("first", "second", etc), and named arguments are referred to by name.

Because of how tables work, there can be two kinds of arguments: named and
unnamed.

```lua
-- Unnamed arguments
func{value, value}

-- Named arguments
func{arg1=value, arg2=value}
```

## Workflow
The primary functions used in every script are the `rbxmk.input`,
`rbxmk.output`, and `rbxmk.map` functions.

The `input` and `output` functions each refer to some piece of **data**,
returning a representation of the data.

```lua
input_data = rbxmk.input{reference...}
output_data = rbxmk.output{reference...}
```

Data is retrieved from some location referred to by the arguments of each
function. In the case of `output`, the returned value is a reference that merely
points to the data, which can be resolved at a later time.

The `map` function receives a number of inputs and a number of outputs, then
maps each input to each output.

```lua
rbxmk.map{input_data, output_data}
```

At this point, the output reference is resolved, then merged with the input
data. Finally, the results are written back to the location referred to by the
output reference.

*A less friendly, more detailed description of how this process works is
available in [DOCUMENTATION.md](DOCUMENTATION.md#resolve-chain).*

## References
A **reference** is a list of strings that specify a location to read data from
(`rbxmk.input`), or to write data to (`rbxmk.output`). Usually, each successive
string specifies a piece of data within the data referred to by the previous
string.

For example, the first string can refer to the location of a file in the file
system. Depending on the format of the file, the next string in the reference
will **drill** down into the file, selecting a piece of data within it.

```lua
input_source = rbxmk.input{"Documents/Roblox/place.rbxl", "Workspace.Model.Part.Script", "Source"}
```

This example selects the "place.rbxl" file, then drills into this `rbxl` file to
select a Script instance, then drills into the Script to select its "Source"
property. The `input_source` variable now contains the Source property's value,
which is a string.

The same procedure applies to outputs, except that they refer to the location
where data will be written to. For example, we can select another script in
another file, and write `input_source` to there.

```lua
output_source = rbxmk.output{"another_place.rbxlx", "ServerScriptService.ModuleScript", "Source"}
rbxmk.map{input_source, output_source}
```

In this example, mapping `input_source` to the output causes the file to be
rewritten to include the source at the specified location. The following steps
describe what happens:

1. Read the content of "another_place.rbxlx", represented as a tree of
   instances.
2. Drill into the tree to the specified location (the Source property of a
   ModuleScript).
3. Merge `input_source` into the location (merging a string into a string simply
   overwrites one with the other).
4. Write the entire tree (with the modifications) back to the file.

## Schemes
The first string of a Reference is always a URI. It begins with a scheme
(`scheme-name://`), which specifies the type of resource being identified. The
remainder of the URI depends on the scheme. Here are a few of them:

- `file://`: The URI is a path to a file on the file system.
- `http://`: The URI is regular URL referring to a web resource.
- `generate://`: The URI specifies the type of value to be generated. The next
  string in the reference defines the value itself.

As this list suggests, while most schemes deal primarily with the first string
of the reference, some schemes may process more than one. Some schemes may also
be defined only for inputs or outputs. For example, the "generate" scheme cannot
be used as an output.

Since the file scheme is the most commonly used, the scheme portion of its URI
can be omitted.

```lua
-- These two are equivalent.
place = rbxmk.input{"Documents/Roblox/place.rbxl"}
place = rbxmk.input{"file://Documents/Roblox/place.rbxl"}
```

A complete list of schemes and how they work can be found in
[DOCUMENTATION.md](DOCUMENTATION.md#user-content-schemes).

## Formats
Once the scheme has been resolved, we are left with the raw data retrieved from
the resource. This isn't very useful on its own, so the raw data is processed
further by using a **format**. A format defines how to decode raw data into
**data** of a known type that is easier to handle. It also defines how to encode
data.

Some schemes are able to guess the format, others require it to be specified
explicitly. Some don't require a format at all.

The file scheme uses the name of a file to determine the format by matching a
format name to the "format extension" of the file. This is almost like a regular
file extension, except more than the last dot (`.`) character can be matched.

For example, `Workspace.properties.json` has the `.json` extension, but also the
`.properties.json` *format* extension, which corresponds to the
`properties.json` format.

The `rbxmk.input` and `rbxmk.output` functions have a named argument called
`format`. When specified with a format name or extension (the leading dot is
optional), it will override whatever format is guessed by the scheme.

```lua
-- Read a regular Lua file as a ModuleScript Lua file.
script = rbxmk.input{format="modulescript.lua", "file.lua"}

-- The format extension also works.
script = rbxmk.input{format=".modulescript.lua", "file.lua"}
```
```lua
rbxmk.map{
	-- The http scheme cannot guess the format, so it must be provided by the
	-- user.
	rbxmk.output{format="rbxl", "https://www.roblox.com/Data/Upload.aspx?assetid=1818"},
	rbxmk.input{"crossroads.rbxl"},
}
```
```lua
-- The generate scheme does not require a format, because no raw data needs to
-- be decoded.
props = rbxmk.input{
	format="properties.json", -- Unnecessary.
	"generate://Property",
	"Size:Vector3=4,1,2;Anchored=true",
}
```

Here is a selection from the many available formats:

- `rbxl`: A Roblox place file.
- `rbxmx`: A Roblox model file in XML format.
- `lua`: An untyped Lua file.
- `script.lua`: A Lua file decoded into a Script instance.
- `properties.json`: A list of property name-value pairs.

A complete list of formats and how they work can be found in
[DOCUMENTATION.md](DOCUMENTATION.md#user-content-formats).

## Data types
After a format has decoded the raw data, it returns typed data that rbxmk knows
how to handle. Here are some examples of the types returned by several formats:

Format            | Data Type    | Description
------------------|--------------|------------
`rbxl`            | `Instances`  | A list of instances.
`lua`             | `Stringlike` | A string.
`properties.json` | `Properties` | A table of property name-value pairs.

The power of a type is that it can be **drilled** into. As previously mentioned,
a `rbxl` file can be drilled into to select a single Instance. What this really
means is that `rbxl` returns an `Instances` type, which returns an `Instance`
type when it is drilled into. This `Instance` type can also be drilled into, to
select a `Property` type. Depending on the value, even this can be drilled
further.

Here's a list of several types and how they can be drilled into:

Input type   | Output type  | Reference
-------------|--------------|----------
`Instances`  | `Instance`   | A dot-separated list of names, with each successive name selecting the child of the previous.
`Instance`   | `Property`   | The name of a property within the instance.
`Instance`   | `Properties` | A `*` character, which selects all properties in the instance.
`Properties` | `Property`   | The name of a property in the table.

Examples:
```lua
-- Instances -> Instance -> Property
rbxmk.input{"place.rbxl", "Workspace.Model.Part", "Anchored"}

-- Instances -> Instance -> Properties -> Property
rbxmk.input{"place.rbxl", "Workspace.Model.Part", "*", "Size"}

-- Properties -> Property
rbxmk.input{"workspace.properties.json", "FilteringEnabled"}
```

A complete list of types and how they work can be found in
[DOCUMENTATION.md](DOCUMENTATION.md#user-content-data-types).

## Functions
There are a handful of other functions beside `rbxmk.input`, `rbxmk.output`, and
`rbxmk.map`.

### Filters
The [`rbxmk.filter`](#user-content-filter-function) function is used to
transform values in some way. The first argument is a string specifying the name
of the filter to use. The remaining arguments, as well as the return values,
depend on the selected filter.

For example, the "minify" filter receives a value and, assuming it's the source
of a Lua script, minifies the content, returning the modified value.

```lua
script = rbxmk.input{"generate://Instance", [[
	Script{
		Name:string = "Script";
		Source:string = "
			for i = 1, 10 do
				print(i)
			end
		";
	}
]]}
script = rbxmk.filter{"minify", script}
-- Value of Source is now "for a=1,10 do print(a)end"
```

A complete list of filters and how they work can be found in
[DOCUMENTATION.md](DOCUMENTATION.md#user-content-filters).

### Path
The [`rbxmk.path`](#user-content-path-function) function is used to join file
paths, adding separators as needed.

```lua
projects = "C:/Users/John/Documents/Roblox"
project_name = "Sword Fights on the Heights"
place_name = "SFOTH.rbxl"
result = rbxmk.path{projects, project_name, place_name}
-- "C:\Users\John\Documents\Roblox\Sword Fights on the Heights\SFOTH.rbxl"
```

`path` also recognizes several variables that will be expanded when they are
included somewhere in an argument:

- `$script_directory`: Expands to the directory of the script currently running.
- `$script_name`: Expands to the base name of the script currently running.
- `$working_directory`: Expands to the current working directory.

```lua
-- Get place file in same directory as running script.
place = rbxmk.input{rbxmk.path{"$script_directory/place.rbxl"}}

-- These are technically the same.
place = rbxmk.input{rbxmk.path{"$working_directory/place.rbxl"}}
place = rbxmk.input{rbxmk.path{"place.rbxl"}}
```

Paths are formatted according to the current platform, so parsing paths directly
should be avoided. Instead, use the file-related functions provided by the rbxmk
library.

### Loading
The [`rbxmk.load`](#user-content-load-function) function allows other scripts to
be run from within a script. The first argument is the path to the script file.
Remaining arguments are passed to the script, which can be received with the
`...` operator. Any values returned by the script are returned by `load`.

```lua
-- template.lua: Make a template.
local template_type = ...
if template_type == "Part" then
	template = rbxmk.input{"generate://Instance", [[
		Part{
			Name:string = "Part";
			Anchored:bool = true;
			Position:Vector3 = 0,0,0;
			Size:Vector3 = 4,1,2;
		}
	]]}
elseif template_type == "Folder" then
	template = rbxmk.input{"generate://Instance", [[
		Folder{Name:string = "Folder"}
	]]}
end
return template
```

```lua
-- Create some templates.
part   = rbxmk.load{"template.lua", "Part"}
folder = rbxmk.load{"template.lua", "Folder"}
```

### Deletion
Sometimes a value needs to be removed rather than replaced or added to. This can
be accomplished with the [`rbxmk.delete`](#user-content-delete-function)
function. This receives one or more outputs, each of which will be removed.

```lua
-- Remove Position property from a Part instance.
rbxmk.delete{rbxmk.output{"place.rbxl", "Workspace.Model.Part", "Position"}}

-- Remove a Part instance.
rbxmk.delete{rbxmk.output{"place.rbxl", "Workspace.Model.Part"}}

-- Remove all instances in the file.
rbxmk.delete{rbxmk.output{"place.rbxl"}}
```

Note that, when deleting an output that points to an entire file, the file
itself will not be removed. It will still exist, and its content will be the
result of the file's format encoding zero values.

### Other functions
A complete list of functions and how they work can be found in
[DOCUMENTATION.md](DOCUMENTATION.md#user-content-rbxmk-library).
