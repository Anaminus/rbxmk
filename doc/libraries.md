# Libraries
This document contains a reference to the libraries available to rbxmk scripts.

<table>
<thead><tr><th>Table of Contents</th></tr></thead>
<tbody><tr><td>

1. [Base][base]
2. [rbxmk][rbxmk]
	1. [rbxmk.Enum][rbxmk.Enum]
	2. [rbxmk.decodeFormat][rbxmk.decodeFormat]
	3. [rbxmk.encodeFormat][rbxmk.encodeFormat]
	4. [rbxmk.formatCanDecode][rbxmk.formatCanDecode]
	5. [rbxmk.globalAttrConfig][rbxmk.globalAttrConfig]
	6. [rbxmk.globalDesc][rbxmk.globalDesc]
	7. [rbxmk.loadFile][rbxmk.loadFile]
	8. [rbxmk.loadString][rbxmk.loadString]
	9. [rbxmk.runFile][rbxmk.runFile]
	10. [rbxmk.runString][rbxmk.runString]
3. [Roblox][roblox]
4. [clipboard][clipboard]
	1. [clipboard.read][clipboard.read]
	2. [clipboard.write][clipboard.write]
5. [fs][fs]
	1. [fs.dir][fs.dir]
	1. [fs.mkdir][fs.mkdir]
	2. [fs.read][fs.read]
	2. [fs.remove][fs.remove]
	2. [fs.rename][fs.rename]
	3. [fs.stat][fs.stat]
	4. [fs.write][fs.write]
6. [http][http]
	1. [http.request][http.request]
7. [math][math]
	1. [math.clamp][math.clamp]
	2. [math.log][math.log]
	3. [math.round][math.round]
	4. [math.sign][math.sign]
8. [os][os]
	1. [os.getenv][os.getenv]
9. [path][path]
	1. [path.clean][path.clean]
	2. [path.expand][path.expand]
	3. [path.join][path.join]
	4. [path.split][path.split]
10. [rbxassetid][rbxassetid]
	1. [rbxassetid.read][rbxassetid.read]
	2. [rbxassetid.write][rbxassetid.write]
11. [string][string]
	1. [string.split][string.split]
12. [sym][sym]
13. [table][table]
	1. [table.clear][table.clear]
	2. [table.create][table.create]
	3. [table.find][table.find]
	4. [table.move][table.move]
	5. [table.pack][table.pack]
	6. [table.unpack][table.unpack]
14. [types][types]
	1. [types.none][types.none]
	2. [types.some][types.some]

</td></tr></tbody>
</table>

[Lua](https://lua.org/) scripts are used to perform actions in rbxmk. The
environment provided by rbxmk is packaged as a set of libraries. Some libraries
are loaded under a specific name, while others are loaded directly into the
global environment:

Library            | Description
-------------------|------------
[(base)][base]     | The Lua 5.1 standard library, abridged.
[rbxmk][rbxmk]     | An interface to the rbxmk engine, and the rbxmk environment.
[(roblox)][roblox] | An environment emulating the Roblox Lua API.
[math][math]       | Extensions to the standard math library.
[os][os]           | Extensions to the standard os library.
[path][path]       | File path manipulation.
[string][string]   | Extensions to the standard string library.
[table][table]     | Extensions to the standard table library.
[sym][sym]         | Symbols for accessing instance metadata.
[types][types]     | Fallbacks for constructing certain types.

The **\_RBXMK_VERSION** global variable is defined as a string containing the
current version of rbxmk, formatted according to [semantic
versioning](https://semver.org/).

## Base
[base]: #user-content-base

The **base** library is loaded directly into the global environment. It contains
the following items from the [Lua 5.1 standard
library](https://www.lua.org/manual/5.1/manual.html#5):

- [\_G](https://www.lua.org/manual/5.1/manual.html#pdf-_G)
- [\_VERSION](https://www.lua.org/manual/5.1/manual.html#pdf-_VERSION)
- [assert](https://www.lua.org/manual/5.1/manual.html#pdf-assert)
- [error](https://www.lua.org/manual/5.1/manual.html#pdf-error)
- [getmetatable](https://www.lua.org/manual/5.1/manual.html#pdf-getmetatable)
- [ipairs](https://www.lua.org/manual/5.1/manual.html#pdf-ipairs)
- [next](https://www.lua.org/manual/5.1/manual.html#pdf-next)
- [pairs](https://www.lua.org/manual/5.1/manual.html#pdf-pairs)
- [pcall](https://www.lua.org/manual/5.1/manual.html#pdf-pcall)
- [print](https://www.lua.org/manual/5.1/manual.html#pdf-print)
- [select](https://www.lua.org/manual/5.1/manual.html#pdf-select)
- [setmetatable](https://www.lua.org/manual/5.1/manual.html#pdf-setmetatable)
- [tonumber](https://www.lua.org/manual/5.1/manual.html#pdf-tonumber)
- [tostring](https://www.lua.org/manual/5.1/manual.html#pdf-tostring)
- [type](https://www.lua.org/manual/5.1/manual.html#pdf-type)
- [unpack](https://www.lua.org/manual/5.1/manual.html#pdf-unpack)
- [xpcall](https://www.lua.org/manual/5.1/manual.html#pdf-xpcall)
- [math library](https://www.lua.org/manual/5.1/manual.html#5.6)
- [string library](https://www.lua.org/manual/5.1/manual.html#5.4), except string.dump
- [table library](https://www.lua.org/manual/5.1/manual.html#5.5)
- [os.clock](https://www.lua.org/manual/5.1/manual.html#pdf-os.clock)
- [os.date](https://www.lua.org/manual/5.1/manual.html#pdf-os.date)
- [os.difftime](https://www.lua.org/manual/5.1/manual.html#pdf-os.difftime)
- [os.time](https://www.lua.org/manual/5.1/manual.html#pdf-os.time)

## rbxmk
[rbxmk]: #user-content-rbxmk

The **rbxmk** library contains functions related to the rbxmk engine.

Name                                             | Kind     | Description
-------------------------------------------------|----------|------------
[rbxmk.Enum][rbxmk.Enum]                         | Enums    | Collection of rbxmk-defined enums.
[rbxmk.decodeFormat][rbxmk.decodeFormat]         | function | Deserialize data from bytes.
[rbxmk.encodeFormat][rbxmk.encodeFormat]         | function | Serialize data into bytes.
[rbxmk.formatCanDecode][rbxmk.formatCanDecode]   | function | Check whether a format decodes into a type.
[rbxmk.globalAttrConfig][rbxmk.globalAttrConfig] | field    | Get or set the global AttrConfig.
[rbxmk.globalDesc][rbxmk.globalDesc]             | field    | Get or set the global descriptor.
[rbxmk.loadFile][rbxmk.loadFile]                 | function | Load the content of a file as a function.
[rbxmk.loadString][rbxmk.loadString]             | function | Load a string as a function.
[rbxmk.runFile][rbxmk.runFile]                   | function | Run a file as a Lua chunk.
[rbxmk.runString][rbxmk.runString]               | function | Run a string as a Lua chunk.

### rbxmk.Enum
[rbxmk.Enum]: #user-content-rbxmkenum
<code>rbxmk.Enum: [Enums](##)</code>

The **Enum** field is a collection of Enum values defined by rbxmk.

### rbxmk.decodeFormat
[rbxmk.decodeFormat]: #user-content-rbxmkdecodeformat
<code>rbxmk.decodeFormat(format: [string](##), bytes: [BinaryString](##)): (value: [any](##))</code>

The **decodeFormat** function decodes *bytes* into a value according to
*format*. The exact details of each format are described in the
[Formats](formats.md) documents.

decodeFormat will throw an error if the format does not exist, or the format has
no decoder defined.

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
<code>rbxmk.globalDesc: [Desc][Desc]?</code>

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
to the usual Roblox types, typeof will work for various types specific to rbxmk.

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

## clipboard
[clipboard]: #user-content-clipboard

The **clipboard** library provides an interface to the operating system's
clipboard.

Name                               | Description
-----------------------------------|------------
[clipboard.read][clipboard.read]   | Gets data from the clipboard in one of a number of formats.
[clipboard.write][clipboard.write] | Sets data to the clipboard in a number of formats.

**The clipboard is currently available only on Windows. Other operating systems
return no data.**

### clipboard.read
[clipboard.read]: #user-content-clipboardread
<code>clipboard.read(formats: ...[string](##)): (value: [any](##)?)</code>

The **read** function gets a value from the clipboard according to one of the
given [formats](formats.md).

Each format has a number of associated [media
types](https://en.wikipedia.org/wiki/Media_type). Each format is traversed in
order, and each media type within a format is traversed in order. The data that
matches the first media type found in the clipboard is selected. This data is
decoded by the format corresponding to the matched media type, and the result is
returned.

Throws an error if *value* could not be decoded from the format, or if data
could not be retrieved from the clipboard. If no data was found, then nil is
returned.

### clipboard.write
[clipboard.write]: #user-content-clipboardwrite
<code>clipboard.write(value: [any](##), formats: ...[string](##))</code>

The **write** function sets *value* to the clipboard according to the given
[formats](formats.md).

Each format has a number of associated [media
types](https://en.wikipedia.org/wiki/Media_type). For each format, the data is
encoded in the format, which is then sent to the clipboard for each of the
format's media type. Data is not sent for a media type if that media type was
already sent.

If no formats are given, then the clipboard is cleared with no further action.

Throws an error if *value* could not be encoded in a format, or if data could
not be sent to the clipboard.

## fs
[fs]: #user-content-fs

The **fs** library provides an interface to the file system.

Name                   | Description
-----------------------|------------
[fs.dir][fs.dir]       | Gets a list of files in a directory.
[fs.mkdir][fs.mkdir]   | Makes a new directory.
[fs.read][fs.read]     | Reads data from a file in a certain format.
[fs.remove][fs.remove] | Removes a file or directory.
[fs.rename][fs.rename] | Moves a file or directory.
[fs.stat][fs.stat]     | Gets metadata about a file.
[fs.write][fs.write]   | Writes data to a file in a certain format.

### fs.dir
[fs.dir]: #user-content-fsdir
<code>fs.dir(path: [string](##)): {[DirEntry](##)}?</code>

The **dir** function returns a list of files in the given directory. Each file
is a table with the following fields:

Field   | Type    | Description
--------|---------|------------
Name    | string  | The base name of the file.
IsDir   | boolean | Whether the file is a directory.

dir returns nil if the directory does not exist. An error is thrown if a problem
otherwise occurred while reading the directory.

### fs.mkdir
[fs.mkdir]: #user-content-fsmkdir
<code>fs.mkdir(path: [string](##), all: [bool](##)?): [bool](##)</code>

The **mkdir** function creates a directory at *path*. If *all* is true, then
mkdir will create each parent directory as needed. *all* defaults to false.

Returns true if all the directories were created successfully. Returns false if
all of the directories already exist. Throws an error if a problem otherwise
occurred while creating a directory.

### fs.read
[fs.read]: #user-content-fsread
<code>fs.read(path: [string](##), format: [string](##)?): (value: [any](##))</code>

The **read** function reads the content of the file at *path*, and decodes it
into *value* according to the [format](formats.md) matching the file extension
of *path*. If *format* is given, then it will be used instead of the file
extension.

If the format returns an [Instance][Instance], then the Name property will be
set to the "fstem" component of *path* according to
[path.split](libraries.md#user-content-pathsplit).

### fs.remove
[fs.remove]: #user-content-fsremove
<code>fs.remove(path: [string](##), all: [bool](##)?): [bool](##)</code>

The **remove** function removes the file or directory at *path*. If *all* is
true, then removing a directory will also recursively remove all of its
children. *all* defaults to false.

Returns true if every file is removed successfully. Returns false if the file or
directory does not exist. Throws an error if a problem occurred while removing a
file.

### fs.rename
[fs.rename]: #user-content-fsrename
<code>fs.rename(old: [string](##), new: [string](##)): [bool](##)</code>

The **rename** functions moves the file or directory at path *old* to path
*new*. If *new* exists and is not a directory, it is replaced.

Returns true if the file was moved successfully. Returns false if the file or
directory does not exist. Throws an error if a problem otherwise occurred while
moving the file.

### fs.stat
[fs.stat]: #user-content-fsstat
<code>fs.stat(path: [string](##)): [FileInfo](##)?</code>

The **stat** function gets metadata of the given file. Returns a table with the
following fields:

Field   | Type    | Description
--------|---------|------------
Name    | string  | The base name of the file.
IsDir   | boolean | Whether the file is a directory.
Size    | number  | The size of the file, in bytes.
ModTime | number  | The modification time of the file, in Unix time.

stats returns nil if the file does not exist. An error will be thrown if a
problem otherwise occurred while getting the metadata.

stat does not follow symbolic links.

### fs.write
[fs.write]: #user-content-fswrite
<code>fs.write(path: [string](##), value: [any](##), format: [string](##)?)</code>

The **write** function encodes *value* according to the [format](formats.md)
matching the file extension of *path*, and writes the result to the file at
*path*. If *format* is given, then it will be used instead of the file
extension.

## http
[http]: #user-content-http

The **http** library provides an interface to resources on the network via HTTP.

Name                         | Description
-----------------------------|------------
[http.request][http.request] | Begins an HTTP request.

### http.request
[http.request]: #user-content-httprequest
<code>http.request(options: [HttpOptions][HttpOptions]): (req: [HttpRequest][HttpRequest])</code>

The **request** function begins a request with the specified
[options][HttpOptions]. Returns a [request object][HttpRequest] that may be
[resolved][HttpRequest.Resolve] or [canceled][HttpRequest.Cancel]. Throws an
error if the request could not be started.

## math
[math]: #user-content-math

The **math** library is an extension to the standard library that includes the
same additions to [Roblox's math
library](https://developer.roblox.com/en-us/api-reference/lua-docs/math):

Name                     | Description
-------------------------|------------
[math.clamp][math.clamp] | Returns a number clamped between a minimum and maximum.
[math.log][math.log]     | Includes optional base argument.
[math.round][math.round] | Rounds a number to the nearest integer.
[math.sign][math.sign]   | Returns the sign of a number.

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

Name                   | Description
-----------------------|------------
[os.getenv][os.getenv] | Gets an environment variable.

### os.getenv
[os.getenv]: #user-content-osgetenv
<code>os.getenv(name: [string](##)?): [string](##)? \| {\[[string](##)\]: [string](##)}</code>

The **getenv** function returns the value of the *name* environment variable, or
nil if no such value is defined. If *name* is not specified, then a table of
environment variable names mapped to values is returned.

## path
[path]: #user-content-path

The **path** library provides functions that handle file paths. The following
functions are included:

Name                       | Description
---------------------------|------------
[path.clean][path.clean]   | Cleans up a file path.
[path.expand][path.expand] | Expands predefined file path variables.
[path.join][path.join]     | Joins a number of file paths together.
[path.split][path.split]   | Splits a file path into its components.

### path.clean
[path.clean]: #user-content-pathclean
<code>path.clean(path: [string](##)): [string](##)</code>

The **clean** function returns the shortest path equivalent to *path*.
Separators are replaced with the OS-specific path separator.

### path.expand
[path.expand]: #user-content-pathexpand
<code>path.expand(path: [string](##)): [string](##)</code>

The **expand** function scans *path* for certain variables of the form `$var` or
`${var}` an expands them. The following variables are expanded:

Variable                                             | Description
-----------------------------------------------------|------------
`$script_name`, `$sn`                                | The base name of the currently running script. Empty for stdin.
`$script_directory`, `$script_dir`, `$sd`            | The directory of the currently running script. Empty for stdin.
`$root_script_directory`, `$root_script_dir`, `$rsd` | The directory of the first running script. Empty for stdin.
`$working_directory`, `$working_dir`, `$wd`          | The current working directory.
`$temp_directory`, `$temp_dir`, `$tmp`               | The directory for temporary files.

### path.join
[path.join]: #user-content-pathjoin
<code>path.join(paths: ...[string](##)): [string](##)</code>

The **join** function joins each *path* element into a single path, separating
them using the operating system's path separator. This also cleans up the path.

### path.split
[path.split]: #user-content-pathsplit
<code>path.split(path: [string](##), components: ...[string](##)): ...[string](##)</code>

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

## rbxassetid
[rbxassetid]: #user-content-rbxassetid

The **rbxassetid** library provides an interface to assets on the Roblox
website.

Name                                 | Description
-------------------------------------|------------
[rbxassetid.read][rbxassetid.read]   | Reads data from a rbxassetid in a certain format.
[rbxassetid.write][rbxassetid.write] | Writes data to a rbxassetid in a certain format.

### rbxassetid.read
[rbxassetid.read]: #user-content-rbxassetidread
<code>rbxassetid.read(options: [RbxAssetOptions][RbxAssetOptions]): (value: [any](##))</code>

The **read** function downloads an asset according to the given
[options][RbxAssetOptions]. Returns the content of the asset corresponding to
AssetId, decoded according to Format.

Throws an error if a problem occurred while downloading the asset.

### rbxassetid.write
[rbxassetid.write]: #user-content-rbxassetidwrite
<code>rbxassetid.write(options: [RbxAssetOptions][RbxAssetOptions])</code>

The **write** function uploads to an existing asset according to the given
[options][RbxAssetOptions]. The Body is encoded according to Format, then
uploaded to AssetId. AssetId must be the ID of an existing asset.

Throws an error if a problem occurred while uploading the asset.

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

Symbol                                          | Description
------------------------------------------------|------------
[sym.AttrConfig][Instance.sym.AttrConfig]       | Gets the inherited [AttrConfig][AttrConfig] of an instance.
[sym.Desc][Instance.sym.Desc]                   | Gets the inherited [descriptor][Desc] of an instance.
[sym.IsService][Instance.sym.IsService]         | Determines whether an instance is a service.
[sym.Metadata][DataModel.sym.Metadata]          | Gets the metadata of a [DataModel][DataModel].
[sym.Properties][Instance.sym.Properties]       | Gets the properties of an instance.
[sym.RawAttrConfig][Instance.sym.RawAttrConfig] | Accesses the direct [AttrConfig][AttrConfig] of an instance.
[sym.RawDesc][Instance.sym.RawDesc]             | Accesses the direct [descriptor][Desc] of an instance.
[sym.Reference][Instance.sym.Reference]         | Determines the value used to identify the instance.

## table
[table]: #user-content-table

The **table** library is an extension to the standard library that includes the
same additions to [Roblox's table
library](https://developer.roblox.com/en-us/api-reference/lua-docs/table):

Name                         | Description
-----------------------------|------------
[table.clear][table.clear]   | Removes all entries from a table.
[table.create][table.create] | Creates a new table with a preallocated capacity.
[table.find][table.find]     | Find the index of a value in a table.
[table.move][table.move]     | Copies the entries in a table.
[table.pack][table.pack]     | Packs arguments into a table.
[table.unpack][table.unpack] | Unpacks a table into arguments.

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

## types
[types]: #user-content-types

The **types** library contains functions for constructing explicit primitives.
The name of a function corresponds directly to the type. See [Explicit
primitives](README.md#user-content-explicit-primitives) for more information.

Type            | Primitive
----------------|----------
BinaryString    | string
Color3uint8     | Color3
Content         | string
float           | number
int64           | number
int             | number
ProtectedString | string
SharedString    | string
token           | number

The library contains several additional functions:

### types.none
[types.none]: #user-content-typesnone
<code>types.none(type: [string](##)): [Optional](##)</code>

The **none** function returns an empty Optional exprim of the given type.

```lua
model.WorldPivotData = types.none("CFrame") -- type is Optional<CFrame>
print(typeof(model.WorldPivotData.Value)) --> nil
```

### types.some
[types.some]: #user-content-typessome
<code>types.some(value: [any](##)): [Optional](##)</code>

The **some** function returns an Optional exprim with the type of *value*, that
encapsulates *value*.

```lua
local value = CFrame.new(1,2,3)
model.WorldPivotData = types.some(value) -- type is Optional<CFrame>
print(typeof(model.WorldPivotData.Value)) --> CFrame
```

[AttrConfig]: types.md#user-content-attrconfig
[CallbackDesc]: types.md#user-content-callbackdesc
[ClassDesc]: types.md#user-content-classdesc
[DataModel.sym.Metadata]: types.md#user-content-datamodelsymmetadata
[DataModel]: types.md#user-content-datamodel
[DescAction]: types.md#user-content-descaction
[EnumDesc]: types.md#user-content-enumdesc
[EnumItemDesc]: types.md#user-content-enumitemdesc
[EventDesc]: types.md#user-content-eventdesc
[FunctionDesc]: types.md#user-content-functiondesc
[HttpOptions]: types.md#user-content-httpoptions
[HttpRequest]: types.md#user-content-httprequest
[HttpRequest.Resolve]: types.md#user-content-httprequestresolve
[HttpRequest.Cancel]: types.md#user-content-httprequestcancel
[Instance.sym.AttrConfig]: types.md#user-content-instancesymattrconfig
[Instance.sym.Desc]: types.md#user-content-instancesymdesc
[Instance.sym.IsService]: types.md#user-content-instancesymisservice
[Instance.sym.Properties]: types.md#user-content-instancesymproperties
[Instance.sym.RawAttrConfig]: types.md#user-content-instancesymrawattrconfig
[Instance.sym.RawDesc]: types.md#user-content-instancesymrawdesc
[Instance.sym.Reference]: types.md#user-content-instancesymreference
[Instance]: types.md#user-content-instance
[ParameterDesc]: types.md#user-content-parameterdesc
[PropertyDesc]: types.md#user-content-propertydesc
[RbxAssetOptions]: types.md#user-content-rbxassetoptions
[Desc]: types.md#user-content-desc
[TypeDesc]: types.md#user-content-typedesc
