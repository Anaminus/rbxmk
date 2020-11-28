# Documentation
This document contains a complete reference of the rbxmk API, and provides
details on how rbxmk works.

<table>
<thead><tr><th>Table of Contents</th></tr></thead>
<tbody><tr><td>

1. [Command line][command-line]
2. [Environment][environment]
	1. [Base library][base-lib]
	2. [`rbxmk` library][rbxmk-lib]
	3. [Roblox library][roblox-lib]
	4. [`os` library][os-lib]
	5. [`sym` library][sym-lib]
	6. [`types` library][types-lib]
3. [Instances][instances]
	1. [Instance][Instance]
		1. [DataModel][DataModel]
4. [Descriptors][descriptors]
	1. [Descriptor types][descriptor-types]
	2. [Diffing and Patching][diffing-and-patching]
5. [Explicit primitives][explicit-primitives]
6. [Sources][sources]
	1. [`file` source][file-source]
	2. [`http` source][http-source]
7. [Formats][formats]
	1. [String formats][string-formats]
	3. [Lua formats][lua-formats]
	4. [Roblox formats][roblox-formats]
	5. [Descriptor formats][descriptor-formats]

</td></tr></tbody>
</table>

This document uses [Luau][luau] type annotation syntax to describe the API of an
element. Some liberties are taken for patterns not supported by the Luau syntax.
For example, `...` indicates variable parameters.

[luau]: https://roblox.github.io/luau/

# Command line
[command-line]: #user-content-command-line

```bash
rbxmk [ FILE ] [ ...VALUE ]
```

The rbxmk command receives a path to a file to be executed as a Lua script.

```bash
rbxmk script.lua
```

If `-` is given, then the script will be read from stdin instead.

```bash
echo 'print("hello world!")' | rbxmk -
```

The remaining arguments are Lua values to be passed to the file. Numbers, bools,
and nil are parsed into their respective types in Lua, and any other value is
interpreted as a string.

```bash
rbxmk script.lua true 3.14159 hello!
```

Within the script, these arguments can be received from the `...` operator:

```lua
local arg1, arg2, arg3 = ...
```

# Environment
[environment]: #user-content-environment

[Lua][lua] scripts are used to perform actions in rbxmk. The environment
provided by rbxmk is packaged as a set of libraries. Some libraries are loaded
under a specific name, while others are loaded directly into the global
environment:

Library                | Description
-----------------------|------------
[(base)][base-lib]     | The Lua 5.1 standard library, abridged.
[rbxmk ][rbxmk-lib]    | An interface to the rbxmk engine, and the rbxmk environment.
[(roblox)][roblox-lib] | An environment emulating the Roblox Lua API.
[os][os-lib]           | Extensions to the standard os library.
[sym][sym-lib]         | Symbols for accessing instance metadata.
[types][types-lib]     | Fallbacks for constructing certain types.
[(sources)][sources]   | An assortment of libraries for interfacing with various external sources.

Additionally, the `_RBXMK_VERSION` global variable is defined as a string
containing the current version of rbxmk, formatted according to [semantic
versioning][semver].

[lua]: https://lua.org/
[semver]: https://semver.org/

## Base library
[base-lib]: #user-content-base-library

The following items from the [Lua 5.1 standard library][luastdlib] are included:

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

[luastdlib]: https://www.lua.org/manual/5.1/manual.html#5

## `rbxmk` library
[rbxmk-lib]: #user-content-rbxmk-library

The **rbxmk** library contains functions related to the rbxmk engine.

Name                               | Description
-----------------------------------|------------
[decodeFormat][rbxmk.decodeFormat] | Deserialize data from bytes.
[diffDesc][rbxmk.diffDesc]         | Get the differences between two descriptors.
[encodeFormat][rbxmk.encodeFormat] | Serialize data into bytes.
[globalDesc][rbxmk.globalDesc]     | Get or set the global descriptor.
[loadFile][rbxmk.loadFile]         | Load the content of a file as a function.
[loadString][rbxmk.loadString]     | Load a string as a function.
[newDesc][rbxmk.newDesc]           | Create a new descriptor.
[patchDesc][rbxmk.patchDesc]       | Transform a descriptor by applying differences.
[readSource][rbxmk.readSource]     | Read bytes from an external source.
[runFile][rbxmk.runFile]           | Run a file as a Lua chunk.
[runString][rbxmk.runString]       | Run a string as a Lua chunk.
[writeSource][rbxmk.writeSource]   | Write bytes to an external source.

### rbxmk.decodeFormat
[rbxmk.decodeFormat]: #user-content-rbxmkdecodeformat
<code>rbxmk.decodeFormat(format: [string](##), bytes: [BinaryString](##)): (value: [any](##))</code>

The **decodeFormat** function decodes *bytes* into a value according to
*format*. The exact details of each format are described in the
[Formats][formats] section.

decodeFormat will throw an error if the format does not exist, or the format has
no decoder defined.

### rbxmk.globalDesc
[rbxmk.globalDesc]: #user-content-rbxmkglobaldesc
<code>rbxmk.globalDesc: [RootDesc][RootDesc]?</code>

The **globalDesc** field gets or sets the global root descriptor. Most items
that utilize a root descriptor will fallback to the global descriptor when
possible.

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
[Formats][formats] section.

encodeFormat will throw an error if the format does not exist, or the format has
no encoder defined.

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

### rbxmk.readSource
[rbxmk.readSource]: #user-content-rbxmkreadsource
<code>rbxmk.readSource(source: [string](##), args: ...[any](##)): (bytes: [BinaryString](##))</code>

The **readSource** function reads a sequence of bytes from an external source
indicated by *source*. *args* depends on the source. The exact details of each
source are described in the [Sources][sources] section.

readSource will throw an error if *source* does not exist, or the source cannot
be read from.

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

### rbxmk.writeSource
[rbxmk.writeSource]: #user-content-rbxmkwritesource
<code>rbxmk.writeSource(source: [string](##), bytes: [BinaryString](##), args: ...[any](##))</code>

The **writeSource** function writes a sequence of bytes to an external source
indicated by *source*. *args* depends on the source. The exact details of each
source are described in the [Sources][sources] section.

writeSource will throw an error if *source* does not exist, or cannot be written
to.

## Roblox library
[roblox-lib]: #user-content-roblox-library

The Roblox library includes an environment similar to the Roblox Lua API. This
is included directly into the global environment.

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

## `os` library
[os-lib]: #user-content-os-library

The **os** library is an extension to the standard library. The following
additional functions are included:

Name                | Description
--------------------|------------
[dir][os.dir]       | Gets a list of files in a directory.
[expand][os.expand] | Expands predefined file path variables.
[getenv][os.getenv] | Gets an environment variable.
[join][os.join]     | Joins a number of file paths together.
[split][os.split]   | Splits a file path into its components.

### os.dir
[os.dir]: #user-content-osdir
<code>os.dir(path: [string](##)): {[File](##)}</code>

The **dir** function returns a list of files in the given directory.

Each file is a table with the following fields:

Field   | Type    | Description
--------|---------|------------
Name    | string  | The base name of the file.
IsDir   | boolean | Whether the file is a directory.
Size    | number  | The size of the file, in bytes.
ModTime | number  | The modification time of the file, in Unix time.

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

A format extension depends on the available formats. See [Formats][formats] for
more information.

## `sym` library
[sym-lib]: #user-content-sym-library

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

Symbol                                    | Description
------------------------------------------|------------
[`sym.Desc`][Instance.sym.Desc]           | Gets the inherited descriptor of an instance.
[`sym.IsService`][Instance.sym.IsService] | Determines whether an instance is a service.
[`sym.RawDesc`][Instance.sym.RawDesc]     | Accesses the direct director of an instance.
[`sym.Reference`][Instance.sym.Reference] | Determines the value used to identify the instance.

## `types` library
[types-lib]: #user-content-types-library

The **types** library contains functions for constructing explicit primitives.
The name of a function corresponds directly to the type. See [Explicit
primitives][explicit-primitives] for more information.

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

# Instances
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

## Instance type
[Instance]: #user-content-instance-type

The **Instance** type provides a similar API to that of Roblox. In addition to
getting and setting properties as described previously, instances have the
following members defined:

Member                                                          | Kind
----------------------------------------------------------------|-----
[ClassName][Instance.ClassName]                                 | property
[Name][Instance.Name]                                           | property
[Parent][Instance.Parent]                                       | property
[ClearAllChildren][Instance.ClearAllChildren]                   | method
[Clone][Instance.Clone]                                         | method
[Destroy][Instance.Destroy]                                     | method
[FindFirstAncestor][Instance.FindFirstAncestor]                 | method
[FindFirstAncestorOfClass][Instance.FindFirstAncestorOfClass]   | method
[FindFirstAncestorWhichIsA][Instance.FindFirstAncestorWhichIsA] | method
[FindFirstChild][Instance.FindFirstChild]                       | method
[FindFirstChildOfClass][Instance.FindFirstChildOfClass]         | method
[FindFirstChildWhichIsA][Instance.FindFirstChildWhichIsA]       | method
[GetChildren][Instance.GetChildren]                             | method
[GetDescendants][Instance.GetDescendants]                       | method
[GetFullName][Instance.GetFullName]                             | method
[IsA][Instance.IsA]                                             | method
[IsAncestorOf][Instance.IsAncestorOf]                           | method
[IsDescendantOf][Instance.IsDescendantOf]                       | method
[sym.Desc][Instance.sym.Desc]                                   | symbol
[sym.IsService][Instance.sym.IsService]                         | symbol
[sym.RawDesc][Instance.sym.RawDesc]                             | symbol
[sym.Reference][Instance.sym.Reference]                         | symbol

### Instance.new
[Instance.new]: #user-content-instancenew
<code>Instance.new(className: [string](##), parent: [Instance][Instance]?, desc: [RootDesc][RootDesc]?): [Instance][Instance]</code>

The `Instance.new` constructor returns a new Instance of the given class.
*className* sets the [ClassName][Instance.ClassName] property of the instance.
If *parent* is specified, it sets the [Parent][Instance.Parent] property.

If *desc* is specified, then it sets the [`sym.Desc`][Instance.sym.Desc] member.
Additionally, `Instance.new` will throw an error if the class does not exist, or
has the "NotCreatable" tag. If no descriptor is specified, then any class name
will be accepted.

### Instance.ClassName
[Instance.ClassName]: #user-content-instanceclassname
<code>Instance.ClassName: [string](##)</code>

ClassName gets or sets the class of the instance.

Unlike in Roblox, ClassName can be modified.

### Instance.Name
[Instance.Name]: #user-content-instancename
<code>Instance.Name: [string](##)</code>

Name gets or sets a name identifying the instance.

### Instance.Parent
[Instance.Parent]: #user-content-instanceparent
<code>Instance.Parent: [Instance][Instance]?</code>

Parent gets or sets the parent of the instance, which may be nil.

### Instance.ClearAllChildren
[Instance.ClearAllChildren]: #user-content-instanceclearallchildren
<code>Instance:ClearAllChildren()</code>

ClearAllChildren sets the [Parent][Instance.Parent] of each child of the
instance to nil.

Unlike in Roblox, ClearAllChildren does not affect descendants.

### Instance.Clone
[Instance.Clone]: #user-content-instanceclone
<code>Instance:Clone(): [Instance][Instance]</code>

Clone returns a copy of the instance.

Unlike in Roblox, Clone does not ignore an instance if its Archivable property
is set to false.

### Instance.Destroy
[Instance.Destroy]: #user-content-instancedestroy
<code>Instance:Destroy()</code>

Destroy sets the [Parent][Instance.Parent] of the instance to nil.

Unlike in Roblox, the Parent of the instance remains unlocked. Destroy also does
not affect descendants.

### Instance.FindFirstAncestor
[Instance.FindFirstAncestor]: #user-content-instancefindfirstancestor
<code>Instance:FindFirstAncestor(name: [string](##)): [Instance][Instance]?</code>

FindFirstAncestor returns the first ancestor whose [Name][Instance.Name] equals
*name*, or nil if no such instance was found.

### Instance.FindFirstAncestorOfClass
[Instance.FindFirstAncestorOfClass]: #user-content-instancefindfirstancestorofclass
<code>Instance:FindFirstAncestorOfClass(className: [string](##)): [Instance][Instance]?</code>

FindFirstAncestorOfClass returns the first ancestor of the instance whose
[ClassName][Instance.ClassName] equals *className*, or nil if no such instance
was found.

### Instance.FindFirstAncestorWhichIsA
[Instance.FindFirstAncestorWhichIsA]: #user-content-instancefindfirstancestorwhichisa
<code>Instance:FindFirstAncestorWhichIsA(className: [string](##)): [Instance][Instance]?</code>

FindFirstAncestorWhichIsA returns the first ancestor of the instance whose
[ClassName][Instance.ClassName] inherits *className* according to the instance's
descriptor, or nil if no such instance was found. If the instance has no
descriptor, then the ClassName is compared directly.

### Instance.FindFirstChild
[Instance.FindFirstChild]: #user-content-instancefindfirstchild
<code>Instance:FindFirstChild(name: [string](##), recursive: [bool](##)?): [Instance][Instance]?</code>

FindFirstChild returns the first child of the instance whose
[Name][Instance.Name] equals *name*, or nil if no such instance was found. If
*recurse* is true, then descendants are also searched, top-down.

### Instance.FindFirstChildOfClass
[Instance.FindFirstChildOfClass]: #user-content-instancefindfirstchildofclass
<code>Instance:FindFirstChildOfClass(className: [string](##), recursive: [bool](##)?): [Instance][Instance]?</code>

FindFirstChildOfClass returns the first child of the instance whose
[ClassName][Instance.ClassName] equals *className*, or nil if no such instance
was found. If *recurse* is true, then descendants are also searched, top-down.

### Instance.FindFirstChildWhichIsA
[Instance.FindFirstChildWhichIsA]: #user-content-instancefindfirstchildwhichisa
<code>Instance:FindFirstChildWhichIsA(className: [string](##), recursive: [bool](##)?): [Instance][Instance]?</code>

FindFirstChildWhichIsA returns the first child of the instance whose
[ClassName][Instance.ClassName] inherits *className*, or nil if no such instance
was found. If the instance has no descriptor, then the ClassName is compared
directly. If *recurse* is true, then descendants are also searched, top-down.

### Instance.GetChildren
[Instance.GetChildren]: #user-content-instancegetchildren
<code>Instance:GetChildren(): Objects</code>

GetChildren returns a list of children of the instance.

### Instance.GetDescendants
[Instance.GetDescendants]: #user-content-instancegetdescendants
<code>Instance:GetDescendants(): [Objects](##)</code>

GetDescendants returns a list of descendants of the instance.

### Instance.GetFullName
[Instance.GetFullName]: #user-content-instancegetfullname
<code>Instance:GetFullName(): [string](##)</code>

GetFullName returns the concatenation of the [Name][Instance.Name] of each
ancestor of the instance and the instance itself, separated by `.` characters.
If an ancestor is a [DataModel][DataModel], it is not included.

### Instance.IsA
[Instance.IsA]: #user-content-instanceisa
<code>Instance:IsA(className: [string](##)): [bool](##)</code>

IsA returns whether the [ClassName][Instance.ClassName] inherits from
*className*, according to the instance's descriptor. If the instance has no
descriptor, then IsA returns whether ClassName equals *className*.

### Instance.IsAncestorOf
[Instance.IsAncestorOf]: #user-content-instanceisancestorof
<code>Instance:IsAncestorOf(descendant: [Instance][Instance]): [bool](##)</code>

IsAncestorOf returns whether the instance of an ancestor of *descendant*.

### Instance.IsDescendantOf
[Instance.IsDescendantOf]: #user-content-instanceisdescendantof
<code>Instance:IsDescendantOf(ancestor: [Instance][Instance]): [bool](##)</code>

IsDescendantOf returns whether the instance of a descendant of *ancestor*.

### Instance[sym.Desc]
[Instance.sym.Desc]: #user-content-instancesymdesc
<code>Instance\[sym.Desc\]: [RootDesc][RootDesc] \| [nil](##)</code>

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
  This behaves sort of like a RootDesc that is empty; instance wont inherit any
  other descriptors, and the state can be inherited.

### Instance[sym.IsService]
[Instance.sym.IsService]: #user-content-instancesymisservice
<code>Instance\[sym.IsService\]: [bool](##)</code>

IsService indicates whether the instance is a service, such as Workspace or
Lighting. This is used by some formats to determine how to encode and decode the
instance.

### Instance[sym.RawDesc]
[Instance.sym.RawDesc]: #user-content-instancesymrawdesc
<code>Instance\[sym.RawDesc\]: [RootDesc][RootDesc] \| [bool](##) \| [nil](##)</code>

RawDesc is similar to [`sym.Desc`][Instance.sym.Desc], except that it considers
only the direct descriptor of the current instance.

Getting RawDesc will return a RootDesc if the instance has a descriptor
assigned, false if the descriptor is blocked, or nil if no descriptor is
assigned.

Setting RawDesc behaves the same as setting Desc.

### Instance[sym.Reference]
[Instance.sym.Reference]: #user-content-instancesymreference
<code>Instance\[sym.Reference\]: [string](##)</code>

Reference is a string used to refer to the instance from within a
[DataModel][DataModel]. Certain formats use this to encode a reference to an
instance. For example, the RBXMX format will generate random UUIDs for its
references (e.g. "RBX8B658F72923F487FAE2F7437482EF16D").

### DataModel
[DataModel]: #user-content-datamodel

A **DataModel** is a special case of an [Instance][Instance]. Unlike a normal
Instance, the [ClassName][Instance.ClassName] property of a DataModel cannot be
modified, and the instance has a [GetService][DataModel.GetService] method.
Additionally, other properties are not serialized, and instead determine
metadata used by certain formats (e.g. ExplicitAutoJoints).

### DataModel.new
[DataModel.new]: #user-content-datamodelnew
<code>DataModel.new(desc: [RootDesc][RootDesc]?): [Instance][Instance]</code>

The `DataModel.new` constructor returns a new Instance of the DataModel class.
If *desc* is specified, then it sets the [`sym.Desc`][Instance.sym.Desc] member.

#### DataModel.GetService
[DataModel.GetService]: #user-content-datamodelgetservice
<code>DataModel:GetService(className: [string](##)): [Instance][Instance]</code>

GetService returns the first child of the DataModel whose
[ClassName][Instance.ClassName] equals *className*. If no such child exists,
then a new instance of *className* is created. The [Name][Instance.Name] of the
instance is set to *className*, [`sym.IsService`][Instance.sym.IsService] is set
to true, and [Parent][Instance.Parent] is set to the DataModel.

If the DataModel has a descriptor, then GetService will throw an error if the
created class's descriptor does not have the "Service" tag set.

# Descriptors
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
inherited by any descendant instances. See [`sym.Desc`][Instance.sym.Desc] for
more information.

Additionally, the [`rbxmk.globalDesc`][rbxmk.globalDesc] field may be used to
apply a RootDesc globally. When `globalDesc` is set, any instance that wouldn't
otherwise inherit a descriptor will use this global descriptor.

When an instance has a descriptor, several behaviors are enforced:

- When the global descriptor is set, [`Instance.new`][Instance.new] errors if
  the given class name does not exist (`Instance.new` can also receive a
  descriptor).
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
  [`DataModel.GetService`][DataModel.GetService] must have the "Service" tag.

## Descriptor types
[descriptor-types]: #user-content-descriptor-types

Descriptors are first-class values like any other, and can be modified on the
fly. There are a number of descriptor types, each with their own fields. See
[`rbxmk.newDesc`][rbxmk.newDesc] for creating descriptors.

Type                           | Description
-------------------------------|------------
[RootDesc][RootDesc]           | Describes an entire API.
[ClassDesc][ClassDesc]         | Describes a class.
[PropertyDesc][PropertyDesc]   | Describes a property member.
[FunctionDesc][FunctionDesc]   | Describes a function member.
[EventDesc][EventDesc]         | Describes a event member.
[CallbackDesc][CallbackDesc]   | Describes a callback member.
[ParameterDesc][ParameterDesc] | Describes a parameter of a function, event, or callback. Immutable.
[TypeDesc][TypeDesc]           | Describes a type. Immutable.
[EnumDesc][EnumDesc]           | Describes an enum.
[EnumItemDesc][EnumItemDesc]   | Describes an enum item.

### RootDesc
[RootDesc]: #user-content-rootdesc

RootDesc describes an entire API. It has the following members:

Member                              | Kind
------------------------------------|-----
[Class][RootDesc.Class]             | method
[Classes][RootDesc.Classes]         | method
[AddClass][RootDesc.AddClass]       | method
[RemoveClass][RootDesc.RemoveClass] | method
[Enum][RootDesc.Enum]               | method
[Enums][RootDesc.Enums]             | method
[AddEnum][RootDesc.AddEnum]         | method
[RemoveEnum][RootDesc.RemoveEnum]   | method
[EnumTypes][RootDesc.EnumTypes]     | method

#### RootDesc.Class
[RootDesc.Class]: #user-content-rootdescclass
<code>RootDesc:Class(name: [string](##)): [ClassDesc][ClassDesc]</code>

Class returns the class of the API corresponding to the given name, or nil if no
such class exists.

#### RootDesc.Classes
[RootDesc.Classes]: #user-content-rootdescclasses
<code>RootDesc:Classes(): {[ClassDesc][ClassDesc]}</code>

Classes returns a list of all the classes of the API.

#### RootDesc.AddClass
[RootDesc.AddClass]: #user-content-rootdescaddclass
<code>RootDesc:AddClass(class: [ClassDesc][ClassDesc]): [bool](##)</code>

AddClass adds a new class to the RootDesc, returning whether the class was added
successfully. The class will fail to be added if a class of the same name
already exists.

#### RootDesc.RemoveClass
[RootDesc.RemoveClass]: #user-content-rootdescremoveclass
<code>RootDesc:RemoveClass(name: [string](##)): [bool](##)</code>

RemoveClass removes a class from the RootDesc, returning whether the class was
removed successfully. False will be returned if a class of the given name does
not exist.

#### RootDesc.Enum
[RootDesc.Enum]: #user-content-rootdescenum
<code>RootDesc:Enum(name: [string](##)): [EnumDesc][EnumDesc]</code>

Enum returns an enum of the API corresponding to the given name, or nil if no
such enum exists.

#### RootDesc.Enums
[RootDesc.Enums]: #user-content-rootdescenums
<code>RootDesc:Enums(): {[EnumDesc][EnumDesc]}</code>

Enums returns a list of all the enums of the API.

#### RootDesc.AddEnum
[RootDesc.AddEnum]: #user-content-rootdescaddenum
<code>RootDesc:AddEnum(enum: [EnumDesc][EnumDesc]): [bool](##)</code>

AddEnum adds a new enum to the RootDesc, returning whether the enum was added
successfully. The enum will fail to be added if an enum of the same name already
exists.

#### RootDesc.RemoveEnum
[RootDesc.RemoveEnum]: #user-content-rootdescremoveenum
<code>RootDesc:RemoveEnum(name: [string](##)): [bool](##)</code>

RemoveEnum removes an enum from the RootDesc, returning whether the enum was
removed successfully. False will be returned if an enum of the given name does
not exist.

#### RootDesc.EnumTypes
[RootDesc.EnumTypes]: #user-content-rootdescenumtypes
<code>RootDesc:EnumTypes(): [Enums](##)</code>

EnumTypes returns a set of enum values generated from the current state of the
RootDesc. These enums are associated with the RootDesc, and may be used by
certain properties, so it is important to generate them before operating on such
properties. Additionally, EnumTypes should be called after modifying enum and
enum item descriptors, to regenerate the enum values.

The API of the resulting enums matches that of Roblox's Enums type. A common
pattern is to assign the result of EnumTypes to the "Enum" variable so that it
matches Roblox's API:

```lua
Enum = rootDesc:EnumTypes()
print(Enum.NormalId.Front)
```

### ClassDesc
[ClassDesc]: #user-content-classdesc

ClassDesc describes a class. It has the following members:

Member                                     | Kind
-------------------------------------------|-----
[Name][ClassDesc.Name]                     | field
[Superclass][ClassDesc.Superclass]         | field
[MemoryCategory][ClassDesc.MemoryCategory] | field
[Member][ClassDesc.Member]                 | method
[Members][ClassDesc.Members]               | method
[AddMember][ClassDesc.AddMember]           | method
[RemoveMember][ClassDesc.RemoveMember]     | method
[Tag][ClassDesc.Tag]                       | method
[Tags][ClassDesc.Tags]                     | method
[SetTag][ClassDesc.SetTag]                 | method
[UnsetTag][ClassDesc.UnsetTag]             | method

#### ClassDesc.Name
[ClassDesc.Name]: #user-content-classdescname
<code>ClassDesc.Name: [string](##)</code>

Name is the name of the class.

#### ClassDesc.Superclass
[ClassDesc.Superclass]: #user-content-classdescsuperclass
<code>ClassDesc.Superclass: [string](##)</code>

Superclass is the name of the class from which the current class inherits.

#### ClassDesc.MemoryCategory
[ClassDesc.MemoryCategory]: #user-content-classdescmemorycategory
<code>ClassDesc.MemoryCategory: [string](##)</code>

MemoryCategory describes the category of the class.

#### ClassDesc.Member
[ClassDesc.Member]: #user-content-classdescmember
<code>ClassDesc:Member(name: [string](##)): [MemberDesc](##)</code>

Member returns a member of the class corresponding to the given name, or nil of
no such member exists.

MemberDesc is any one of the [PropertyDesc][PropertyDesc],
[FunctionDesc][FunctionDesc], [EventDesc][EventDesc], or
[CallbackDesc][CallbackDesc] types.

#### ClassDesc.Members
[ClassDesc.Members]: #user-content-classdescmembers
<code>ClassDesc:Members(): {[MemberDesc](##)}</code>

Members returns a list of all the members of the class.

MemberDesc is any one of the [PropertyDesc][PropertyDesc],
[FunctionDesc][FunctionDesc], [EventDesc][EventDesc], or
[CallbackDesc][CallbackDesc] types.

#### ClassDesc.AddMember
[ClassDesc.AddMember]: #user-content-classdescaddmember
<code>ClassDesc:AddMember(member: [MemberDesc](##)): [bool](##)</code>

AddMember adds a new member to the ClassDesc, returning whether the member was
added successfully. The member will fail to be added if a member of the same
name already exists.

MemberDesc is any one of the [PropertyDesc][PropertyDesc],
[FunctionDesc][FunctionDesc], [EventDesc][EventDesc], or
[CallbackDesc][CallbackDesc] types.

#### ClassDesc.RemoveMember
[ClassDesc.RemoveMember]: #user-content-classdescremovemember
<code>ClassDesc:RemoveMember(name: [string](##)): [bool](##)</code>

RemoveMember removes a member from the ClassDesc, returning whether the member
was removed successfully. False will be returned if a member of the given name
does not exist.

#### ClassDesc.Tag
[ClassDesc.Tag]: #user-content-classdesctag
<code>ClassDesc:Tag(name: [string](##)): [bool](##)</code>

Tag returns whether a tag of the given name is set on the descriptor.

#### ClassDesc.Tags
[ClassDesc.Tags]: #user-content-classdesctags
<code>ClassDesc:Tags(): {[string](##)}</code>

Tags returns a list of tags that are set on the descriptor.

#### ClassDesc.SetTag
[ClassDesc.SetTag]: #user-content-classdescsettag
<code>ClassDesc:SetTag(tags: ...[string](##))</code>

SetTags sets the given tags on the descriptor.

#### ClassDesc.UnsetTag
[ClassDesc.UnsetTag]: #user-content-classdescunsettag
<code>ClassDesc:UnsetTag(tags: ...[string](##))</code>

SetTags unsets the given tags on the descriptor.

### PropertyDesc
[PropertyDesc]: #user-content-propertydesc

PropertyDesc describes a property member of a class. It has the following
members:

Member                                      | Kind
--------------------------------------------|-----
[Name][PropertyDesc.Name]                   | field
[ValueType][PropertyDesc.ValueType]         | field
[ReadSecurity][PropertyDesc.ReadSecurity]   | field
[WriteSecurity][PropertyDesc.WriteSecurity] | field
[CanLoad][PropertyDesc.CanLoad]             | field
[CanSave][PropertyDesc.CanSave]             | field
[Tag][PropertyDesc.Tag]                     | method
[Tags][PropertyDesc.Tags]                   | method
[SetTag][PropertyDesc.SetTag]               | method
[UnsetTag][PropertyDesc.UnsetTag]           | method

#### PropertyDesc.Name
[PropertyDesc.Name]: #user-content-propertydescname
<code>PropertyDesc.Name: [string](##)</code>

Name is the name of the member.

#### PropertyDesc.ValueType
[PropertyDesc.ValueType]: #user-content-propertydescvaluetype
<code>PropertyDesc.ValueType: [TypeDesc][TypeDesc]</code>

ValueType is the value type of the property.

#### PropertyDesc.ReadSecurity
[PropertyDesc.ReadSecurity]: #user-content-propertydescreadsecurity
<code>PropertyDesc.ReadSecurity: [string](##)</code>

ReadSecurity indicates the security context required to get the property.

#### PropertyDesc.WriteSecurity
[PropertyDesc.WriteSecurity]: #user-content-propertydescwritesecurity
<code>PropertyDesc.WriteSecurity: [string](##)</code>

WriteSecurity indicates the security context required to set the property.

#### PropertyDesc.CanLoad
[PropertyDesc.CanLoad]: #user-content-propertydesccanload
<code>PropertyDesc.CanLoad: [bool](##)</code>

CanLoad indicates whether the property is deserialized when decoding from a file.

#### PropertyDesc.CanSave
[PropertyDesc.CanSave]: #user-content-propertydesccansave
<code>PropertyDesc.CanSave: [bool](##)</code>

CanLoad indicates whether the property is serialized when encoding to a file.

#### PropertyDesc.Tag
[PropertyDesc.Tag]: #user-content-propertydesctag
<code>PropertyDesc:Tag(name: [string](##)): [bool](##)</code>

Tag returns whether a tag of the given name is set on the descriptor.

#### PropertyDesc.Tags
[PropertyDesc.Tags]: #user-content-propertydesctags
<code>PropertyDesc:Tags(): {string}</code>

Tags returns a list of tags that are set on the descriptor.

#### PropertyDesc.SetTag
[PropertyDesc.SetTag]: #user-content-propertydescsettag
<code>PropertyDesc:SetTag(tags: ...[string](##))</code>

SetTags sets the given tags on the descriptor.

#### PropertyDesc.UnsetTag
[PropertyDesc.UnsetTag]: #user-content-propertydescunsettag
<code>PropertyDesc:UnsetTag(tags: ...[string](##))</code>

SetTags unsets the given tags on the descriptor.

### FunctionDesc
[FunctionDesc]: #user-content-functiondesc

FunctionDesc describes a function member of a class. It has the following
members:

Member                                      | Kind
--------------------------------------------|-----
[Name][FunctionDesc.Name]                   | field
[Parameters][FunctionDesc.Parameters]       | method
[SetParameters][FunctionDesc.SetParameters] | method
[ReturnType][FunctionDesc.ReturnType]       | field
[Security][FunctionDesc.Security]           | field
[Tag][FunctionDesc.Tag]                     | method
[Tags][FunctionDesc.Tags]                   | method
[SetTag][FunctionDesc.SetTag]               | method
[UnsetTag][FunctionDesc.UnsetTag]           | method

#### FunctionDesc.Name
[FunctionDesc.Name]: #user-content-functiondescname
<code>FunctionDesc.Name: [string](##)</code>

Name is the name of the member.

#### FunctionDesc.Parameters
[FunctionDesc.Parameters]: #user-content-functiondescparameters
<code>FunctionDesc:Parameters(): {[ParameterDesc][ParameterDesc]}</code>

Parameters returns a list of parameters of the function.

#### FunctionDesc.SetParameters
[FunctionDesc.SetParameters]: #user-content-functiondescsetparameters
<code>FunctionDesc:SetParameters(params: {[ParameterDesc][ParameterDesc]})</code>

SetParameters sets the parameters of the function.

#### FunctionDesc.ReturnType
[FunctionDesc.ReturnType]: #user-content-functiondescreturntype
<code>FunctionDesc.ReturnType: [TypeDesc][TypeDesc]</code>

ReturnType is the type returned by the function.

#### FunctionDesc.Security
[FunctionDesc.Security]: #user-content-functiondescsecurity
<code>FunctionDesc.Security: [string](##)</code>

Security indicates the security content required to index the member.

#### FunctionDesc.Tag
[FunctionDesc.Tag]: #user-content-functiondesctag
<code>FunctionDesc:Tag(name: [string](##)): [bool](##)</code>

Tag returns whether a tag of the given name is set on the descriptor.

#### FunctionDesc.Tags
[FunctionDesc.Tags]: #user-content-functiondesctags
<code>FunctionDesc:Tags(): {string}</code>

Tags returns a list of tags that are set on the descriptor.

#### FunctionDesc.SetTag
[FunctionDesc.SetTag]: #user-content-functiondescsettag
<code>FunctionDesc:SetTag(tags: ...[string](##))</code>

SetTags sets the given tags on the descriptor.

#### FunctionDesc.UnsetTag
[FunctionDesc.UnsetTag]: #user-content-functiondescunsettag
<code>FunctionDesc:UnsetTag(tags: ...[string](##))</code>

SetTags unsets the given tags on the descriptor.

### EventDesc
[EventDesc]: #user-content-eventdesc

EventDesc describes an event member of a class. It has the following members:

Member                                   | Kind
-----------------------------------------|-----
[Name][EventDesc.Name]                   | field
[Parameters][EventDesc.Parameters]       | method
[SetParameters][EventDesc.SetParameters] | method
[Security][EventDesc.Security]           | field
[Tag][EventDesc.Tag]                     | method
[Tags][EventDesc.Tags]                   | method
[SetTag][EventDesc.SetTag]               | method
[UnsetTag][EventDesc.UnsetTag]           | method

#### EventDesc.Name
[EventDesc.Name]: #user-content-eventdescname
<code>EventDesc.Name: [string](##)</code>

Name is the name of the member.

#### EventDesc.Parameters
[EventDesc.Parameters]: #user-content-eventdescparameters
<code>EventDesc:Parameters(): {[ParameterDesc][ParameterDesc]}</code>

Parameters returns a list of parameters of the event.

#### EventDesc.SetParameters
[EventDesc.SetParameters]: #user-content-eventdescsetparameters
<code>EventDesc:SetParameters(params: {[ParameterDesc][ParameterDesc]})</code>

SetParameters sets the parameters of the event.

#### EventDesc.Security
[EventDesc.Security]: #user-content-eventdescsecurity
<code>EventDesc.Security: [string](##)</code>

Security indicates the security content required to index the member.

#### EventDesc.Tag
[EventDesc.Tag]: #user-content-eventdesctag
<code>EventDesc:Tag(name: [string](##)): [bool](##)</code>

Tag returns whether a tag of the given name is set on the descriptor.

#### EventDesc.Tags
[EventDesc.Tags]: #user-content-eventdesctags
<code>EventDesc:Tags(): {string}</code>

Tags returns a list of tags that are set on the descriptor.

#### EventDesc.SetTag
[EventDesc.SetTag]: #user-content-eventdescsettag
<code>EventDesc:SetTag(tags: ...[string](##))</code>

SetTags sets the given tags on the descriptor.

#### EventDesc.UnsetTag
[EventDesc.UnsetTag]: #user-content-eventdescunsettag
<code>EventDesc:UnsetTag(tags: ...[string](##))</code>

SetTags unsets the given tags on the descriptor.

### CallbackDesc
[CallbackDesc]: #user-content-callbackdesc

CallbackDesc describes a callback member of a class. It has the following
members:

Member                                      | Kind
--------------------------------------------|-----
[Name][CallbackDesc.Name]                   | field
[Parameters][CallbackDesc.Parameters]       | method
[SetParameters][CallbackDesc.SetParameters] | method
[ReturnType][CallbackDesc.ReturnType]       | field
[Security][CallbackDesc.Security]           | field
[Tag][CallbackDesc.Tag]                     | method
[Tags][CallbackDesc.Tags]                   | method
[SetTag][CallbackDesc.SetTag]               | method
[UnsetTag][CallbackDesc.UnsetTag]           | method

#### CallbackDesc.Name
[CallbackDesc.Name]: #user-content-callbackdescname
<code>CallbackDesc.Name: [string](##)</code>

Name is the name of the member.

#### CallbackDesc.Parameters
[CallbackDesc.Parameters]: #user-content-callbackdescparameters
<code>CallbackDesc:Parameters(): {[ParameterDesc][ParameterDesc]}</code>

Parameters returns a list of parameters of the callback.

#### CallbackDesc.SetParameters
[CallbackDesc.SetParameters]: #user-content-callbackdescsetparameters
<code>CallbackDesc:SetParameters(params: {[ParameterDesc][ParameterDesc]})</code>

SetParameters sets the parameters of the callback.

#### CallbackDesc.ReturnType
[CallbackDesc.ReturnType]: #user-content-callbackdescreturntype
<code>CallbackDesc.ReturnType: [TypeDesc][TypeDesc]</code>

ReturnType is the type returned by the callback.

#### CallbackDesc.Security
[CallbackDesc.Security]: #user-content-callbackdescsecurity
<code>CallbackDesc.Security: [string](##)</code>

Security indicates the security content required to set the member.

#### CallbackDesc.Tag
[CallbackDesc.Tag]: #user-content-callbackdesctag
<code>CallbackDesc:Tag(name: [string](##)): [bool](##)</code>

Tag returns whether a tag of the given name is set on the descriptor.

#### CallbackDesc.Tags
[CallbackDesc.Tags]: #user-content-callbackdesctags
<code>CallbackDesc:Tags(): {string}</code>

Tags returns a list of tags that are set on the descriptor.

#### CallbackDesc.SetTag
[CallbackDesc.SetTag]: #user-content-callbackdescsettag
<code>CallbackDesc:SetTag(tags: ...[string](##))</code>

SetTags sets the given tags on the descriptor.

#### CallbackDesc.UnsetTag
[CallbackDesc.UnsetTag]: #user-content-callbackdescunsettag
<code>CallbackDesc:UnsetTag(tags: ...[string](##))</code>

SetTags unsets the given tags on the descriptor.

### ParameterDesc
[ParameterDesc]: #user-content-parameterdesc

ParameterDesc describes a parameter of a function, event, or callback member. It
has the following members:

Member                           | Kind
---------------------------------|-----
[Type][ParameterDesc.Type]       | field
[Name][ParameterDesc.Name]       | field
[Default][ParameterDesc.Default] | field

ParameterDesc is immutable. A new value with different fields can be created
with [`rbxmk.newDesc`][rbxmk.newDesc].

#### ParameterDesc.Type
[ParameterDesc.Type]: #user-content-parameterdesctype
<code>ParameterDesc.Type: [TypeDesc][TypeDesc]</code>

Type is the type of the parameter.

#### ParameterDesc.Name
[ParameterDesc.Name]: #user-content-parameterdescname
<code>ParameterDesc.Name: [string](##)</code>

Name is a name describing the parameter.

#### ParameterDesc.Default
[ParameterDesc.Default]: #user-content-parameterdescdefault
<code>ParameterDesc.Default: [string](##)?</code>

Default is a string describing the default value of the parameter. May also be
nil, indicating that the parameter has no default value.

### TypeDesc
[TypeDesc]: #user-content-typedesc

TypeDesc describes a value type. It has the following members:

Member                        | Kind
------------------------------|-----
[Category][TypeDesc.Category] | field
[Name][TypeDesc.Name]         | field

TypeDesc is immutable. A new value with different fields can be created with
rbxmk.newDesc.

#### TypeDesc.Category
[TypeDesc.Category]: #user-content-typedesccategory
<code>TypeDesc.Category: [string](##)</code>

Category is the category of the type. Certain categories are treated specially:

- **Class**: Name is the name of a class. A value of the type is expected to be
  an Instance of the class.
- **Enum**: Name is the name of an enum. A value of the type is expected to be
  an enum item of the enum.

#### TypeDesc.Name
[TypeDesc.Name]: #user-content-typedescname
<code>TypeDesc.Name: [string](##)</code>

Name is the name of the type.

### EnumDesc
[EnumDesc]: #user-content-enumdesc

EnumDesc describes an enum. It has the following members:

Member                            | Kind
----------------------------------|-----
[Name][EnumDesc.Name]             | field
[Item][EnumDesc.Item]             | method
[Items][EnumDesc.Items]           | method
[AddItem][EnumDesc.AddItem]       | method
[RemoveItem][EnumDesc.RemoveItem] | method
[Tag][EnumDesc.Tag]               | method
[Tags][EnumDesc.Tags]             | method
[SetTag][EnumDesc.SetTag]         | method
[UnsetTag][EnumDesc.UnsetTag]     | method

#### EnumDesc.Name
[EnumDesc.Name]: #user-content-enumdescname
<code>EnumDesc.Name: [string](##)</code>

Name is the name of the enum.

#### EnumDesc.Item
[EnumDesc.Item]: #user-content-enumdescitem
<code>EnumDesc:Item(name: [string](##)): [EnumItemDesc][EnumItemDesc]</code>

Item returns an item of the enum corresponding to given name, or nil of no such
item exists.

#### EnumDesc.Items
[EnumDesc.Items]: #user-content-enumdescitems
<code>EnumDesc:Items(): {[EnumItemDesc][EnumItemDesc]}</code>

Items returns a list of all the items of the enum.

#### EnumDesc.AddItem
[EnumDesc.AddItem]: #user-content-enumdescadditem
<code>EnumDesc:AddItem(item: [EnumItemDesc][EnumItemDesc]): [bool](##)</code>

AddItem adds a new item to the EnumDesc, returning whether the item was added
successfully. The item will fail to be added if an item of the same name already
exists.

#### EnumDesc.RemoveItem
[EnumDesc.RemoveItem]: #user-content-enumdescremoveitem
<code>EnumDesc:RemoveItem(name: [string](##)): [bool](##)</code>

RemoveItem removes an item from the EnumDesc, returning whether the item was
removed successfully. False will be returned if an item of the given name does
not exist.

#### EnumDesc.Tag
[EnumDesc.Tag]: #user-content-enumdesctag
<code>EnumDesc:Tag(name: [string](##)): [bool](##)</code>

Tag returns whether a tag of the given name is set on the descriptor.

#### EnumDesc.Tags
[EnumDesc.Tags]: #user-content-enumdesctags
<code>EnumDesc:Tags(): {string}</code>

Tags returns a list of tags that are set on the descriptor.

#### EnumDesc.SetTag
[EnumDesc.SetTag]: #user-content-enumdescsettag
<code>EnumDesc:SetTag(tags: ...[string](##))</code>

SetTags sets the given tags on the descriptor.

#### EnumDesc.UnsetTag
[EnumDesc.UnsetTag]: #user-content-enumdescunsettag
<code>EnumDesc:UnsetTag(tags: ...[string](##))</code>

SetTags unsets the given tags on the descriptor.

### EnumItemDesc
[EnumItemDesc]: #user-content-enumitemdesc

EnumDesc describes an enum item. It has the following members:

Member                            | Kind
----------------------------------|-----
[Name][EnumItemDesc.Name]         | field
[Value][EnumItemDesc.Value]       | field
[Index][EnumItemDesc.Index]       | field
[Tag][EnumItemDesc.Tag]           | method
[Tags][EnumItemDesc.Tags]         | method
[SetTag][EnumItemDesc.SetTag]     | method
[UnsetTag][EnumItemDesc.UnsetTag] | method

#### EnumItemDesc.Name
[EnumItemDesc.Name]: #user-content-enumitemdescname
<code>EnumItemDesc.Name: [string](##)</code>

Name is the name of the enum item.

#### EnumItemDesc.Value
[EnumItemDesc.Value]: #user-content-enumitemdescvalue
<code>EnumItemDesc.Value: [int](##)</code>

Value is the numeric value of the enum item.

#### EnumItemDesc.Index
[EnumItemDesc.Index]: #user-content-enumitemdescindex
<code>EnumItemDesc.Index: [int](##)</code>

Index is an integer that hints the order of the enum item.

#### EnumItemDesc.Tag
[EnumItemDesc.Tag]: #user-content-enumitemdesctag
<code>EnumItemDesc:Tag(name: [string](##)): [bool](##)</code>

Tag returns whether a tag of the given name is set on the descriptor.

#### EnumItemDesc.Tags
[EnumItemDesc.Tags]: #user-content-enumitemdesctags
<code>EnumItemDesc:Tags(): {string}</code>

Tags returns a list of tags that are set on the descriptor.

#### EnumItemDesc.SetTag
[EnumItemDesc.SetTag]: #user-content-enumitemdescsettag
<code>EnumItemDesc:SetTag(tags: ...[string](##))</code>

SetTags sets the given tags on the descriptor.

#### EnumItemDesc.UnsetTag
[EnumItemDesc.UnsetTag]: #user-content-enumitemdescunsettag
<code>EnumItemDesc:UnsetTag(tags: ...[string](##))</code>

SetTags unsets the given tags on the descriptor.

## Diffing and Patching
[diffing-and-patching]: #user-content-diffing-and-patching

Descriptors can be compared and patched with the
[`rbxmk.diffDesc`][rbxmk.diffDesc] and [`rbxmk.patchDesc`][rbxmk.patchDesc]
functions. `diffDesc` returns a list of [**DescActions**][DescAction], which
describe how to transform the first descriptor into the second. `patchDesc` can
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

### DescAction
[DescAction]: #user-content-descaction

A **DescAction** describes a single action that transforms a descriptor.

Currently, DescAction has no members. However, converting a DescAction to a
string will display the content of the action in a human-readable format.

Actions may also be created directly with the
[`desc-patch.json`][desc-patch.json-fmt] format.

# Explicit primitives
[explicit-primitives]: #user-content-explicit-primitives

The properties of instances in Roblox have a number of different types. Many of
these types can be expressed in Lua through constructors. Examples of such are
`CFrame`, `Vector3`, `UDim2`, and so on. These types correspond to internal data
types within the Roblox engine. The Lua representation of, say, a `CFrame`, is a
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

The [`types` library][types-lib] contains a constructor function for each exprim
type.

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

An exprim userdata has no fields or operators other than the `Value` field,
which returns the underlying primitive value:

	v.Value = types.int64(v.Value.Value + 1)

Exprims are meant to be short-lived, and shouldn't really be used beyond getting
or setting a property in the absence of [descriptors][descriptors]. When
possible, descriptors should be utilized instead.

# Sources
[sources]: #user-content-sources

A **source** is an external location from which raw data can be read from and
written to. A source can be accessed at a low level through the
[`rbxmk.readSource`][rbxmk.readSource] and
[`rbxmk.writeSource`][rbxmk.writeSource] functions.

A source usually has a corresponding library that provides convenient access for
common cases.

## `file` source
[file-source]: #user-content-file-source

The `file` source provides access to the file system.

### `readSource`
[file.readSource]: #user-content-readsource

The first additional argument to [`readSource`][rbxmk.readSource] is the path to
the file to read from.

```lua
local bytes = rbxmk.readSource("file", "path/to/file.ext")
```

### `writeSource`
[file.writeSource]: #user-content-writesource

The first additional argument to [`writeSource`][rbxmk.writeSource] is the path
to the file to write to.

```lua
rbxmk.writeSource("file", bytes, "path/to/file.ext")
```

### `file` library
[file-lib]: #user-content-file-library

The `file` library handles the `file` source.

Name                | Description
--------------------|------------
[read][file.read]   | Reads data from a file in a certain format.
[write][file.write] | Writes data to a file in a certain format.

#### file.read
[file.read]: #user-content-fileread
<code>file.read(path: [string](##), format: [string](##)?): (value: [any](##))</code>

The `read` function reads the content of the file at *path*, and decodes it into
*value* according to the [format][formats] matching the file extension of
*path*. If *format* is given, then it will be used instead of the file
extension.

If the format returns an Instance, then the Name property will be set to the
"fstem" component of *path* according to `os.split`.

#### file.write
[file.write]: #user-content-filewrite
<code>file.write(path: [string](##), value: [any](##), format: [string](##)?)</code>

The `write` function encodes *value* according to the [format][formats] matching
the file extension of *path*, and writes the result to the file at *path*. If
*format* is given, then it will be used instead of the file extension.

## `http` source
[http-source]: #user-content-http-source

The `http` source provides access to an HTTP client.

### `readSource`
[http.readSource]: #user-content-readsource-1

The first additional argument to [`readSource`][rbxmk.readSource] is the URL to
which a GET request will be made. Returns the body of the response. Throws an
error if the response status is not 2XX.

```lua
local bytes = rbxmk.readSource("http", "https://www.example.com/resource")
```

### `writeSource`
[http.readSource]: #user-content-writesource-1

The first additional argument to [`writeSource`][rbxmk.writeSource] is the URL
to which a POST request will be made. The bytes are sent as the body of the
request. Throws an error if the response status is not 2XX.

```lua
rbxmk.writeSource("http", bytes, "https://www.example.com/resource")
```

### `http` library
[http-lib]: #user-content-http-library

The `http` library handles the `http` source.

Name                | Description
--------------------|------------
[read][http.read]   | Reads data from an HTTP URL in a certain format.
[write][http.write] | Writes data to an HTTP URL in a certain format.

#### http.read
[http.read]: #user-content-httpread
<code>http.read(url: [string](##), format: [string](##)?): (value: [any](##))</code>

The `read` function issues a GET request to *url*, and decodes the response body
into *value* according to the [format][formats] matching *format*. Throws an
error if the response status is not 2XX.

#### http.write
[http.write]: #user-content-httpwrite
<code>http.write(url: [string](##), value: [any](##), format: [string](##))</code>

The `write` function encodes *value* according to the [format][formats] matching
*format*, and sends the result in a POST request to *url*. Throws an error if
the response status is not 2XX.

# Formats
[formats]: #user-content-formats

A **format** is capable of encoding a value to raw bytes, or decoding raw bytes
into a value. A format can be accessed at a low level through the
[`rbxmk.encodeFormat`][rbxmk.encodeFormat] and
[`rbxmk.decodeFormat`][rbxmk.decodeFormat] functions.

The name of a format corresponds to the extension of a file name. For example,
the `lua` format corresponds to the `.lua` file extension. When determining a
format from a file extension, format names are greedy; if a file extension is
`.server.lua`, this will select the `server.lua` format before the `lua` format.
For convenience, in places where a format name is received, the name may have an
optional leading `.` character.

A format can decode into a number of certain types, and encode a number of
certain types. A format may also have no definition for either decoding or
encoding at all.

A format that can encode a **Stringlike** type accepts any type that can be
converted to a string. Additionally, an [Instance][Instance] will be accepted as
a Stringlike when its [ClassName][Instance.ClassName] is "Script",
"LocalScript", or "ModuleScript", and the Source property has a Stringlike
value. In this case, the value of the Source property is encoded.

## String formats
[string-formats]: #user-content-string-formats

Several string formats are defined for encoding string-like values.

Format           | Description
-----------------|------------
[`txt`][txt-fmt] | Encodes string-like values to UTF-8 text.
[`bin`][bin-fmt] | Encodes string-like values to raw bytes.

### `txt` format
[txt-fmt]: #user-content-txt-format

The **txt** format encodes UTF-8 text.

Direction | Type       | Description
----------|------------|------------
Decode    | string     | UTF-8 text.
Encode    | Stringlike | Any string-like value.

### `bin` format
[bin-fmt]: #user-content-bin-format

The **bin** format encodes with the assurance that the bytes will be interpreted
exactly as-is.

Direction | Type         | Description
----------|--------------|------------
Decode    | BinaryString | Raw binary data.
Encode    | Stringlike   | Any string-like value.

## Lua formats
[lua-formats]: #user-content-lua-formats

Several formats are defined for decoding Lua files into script instances.

Format                                     | Description
-------------------------------------------|------------
[`modulescript.lua`][modulescript.lua-fmt] | Decodes into a ModuleScript instance.
[`script.lua`][script.lua-fmt]             | Decodes into a Script instance.
[`localscript.lua`][localscript.lua-fmt]   | Decodes into a LocalScript instance.
[`lua`][lua-fmt]                           | Alias for `modulescript.lua`.
[`server.lua`][server.lua-fmt]             | Alias for `script.lua`.
[`client.lua`][client.lua-fmt]             | Alias for `localscript.lua`.

### `modulescript.lua` format
[modulescript.lua-fmt]: #user-content-modulescriptlua-format

The **modulescript.lua** format is a shortcut for decoding Lua code into a
ModuleScript instance, where the content is assigned to the Source property.

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A ModuleScript with a Source property.
Encode    | Stringlike           | Any string-like value.

### `script.lua` format
[script.lua-fmt]: #user-content-scriptlua-format

The **script.lua** format is a shortcut for decoding Lua code into a
Script instance, where the content is assigned to the Source property.

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A Script with a Source property.
Encode    | Stringlike           | Any string-like value.

### `localscript.lua` format
[localscript.lua-fmt]: #user-content-localscriptlua-format

The **localscript.lua** format is a shortcut for decoding Lua code into a
LocalScript instance, where the content is assigned to the Source property.

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A LocalScript with a Source property.
Encode    | Stringlike           | Any string-like value.

### `lua` format
[lua-fmt]: #user-content-lua-format

The **lua** format is an alias for [`modulescript.lua`][modulescript.lua-fmt].

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A ModuleScript with a Source property.
Encode    | Stringlike           | Any string-like value.

### `server.lua` format
[server.lua-fmt]: #user-content-serverlua-format

The **server.lua** format is an alias for [`script.lua`][script.lua-fmt].

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A Script with a Source property.
Encode    | Stringlike           | Any string-like value.

### `client.lua` format
[client.lua-fmt]: #user-content-clientlua-format

The **client.lua** format is an alias for [`localscript.lua`][localscript.lua-fmt].

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [Instance][Instance] | A LocalScript with a Source property.
Encode    | Stringlike           | Any string-like value.

## Roblox formats
[roblox-formats]: #user-content-roblox-formats

Several formats are defined for serializing instances in formats known by
Roblox.

Format  | Description
--------|------------
`rbxl`  | The Roblox binary place format.
`rbxm`  | The Roblox binary model format.
`rbxlx` | The Roblox XML place format.
`rbxmx` | The Roblox XML model format.

Each format can encode and decode the following types:

Direction | Type                   | Description
----------|------------------------|------------
Decode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [Instance][Instance]   | A single instance, interpreted as a child to a DataModel.
Encode    | Objects                | A list of Instances, interpreted as children to a DataModel.

## Descriptor formats
[descriptor-formats]: #user-content-descriptor-formats

Several formats are defined for encoding descriptors.

Format                                   | Description
-----------------------------------------|------------
[`desc.json`][desc.json-fmt]             | Descriptors in JSON format.
[`desc-patch.json`][desc-patch.json-fmt] | Actions that describe changes to descriptors, in JSON format.

### `desc.json` format
[desc.json-fmt]: #user-content-descjson-format

The **desc.json** format encodes a root descriptor file, more commonly known as
an "API dump".

Direction | Type                 | Description
----------|----------------------|------------
Decode    | [RootDesc][RootDesc] | A root descriptor.
Encode    | [RootDesc][RootDesc] | A root descriptor.

### `desc-patch.json` format
[desc-patch.json-fmt]: #user-content-desc-patchjson-format

The **desc-patch.json** format encodes actions that transform descriptors.

Direction | Type        | Description
----------|-------------|------------
Decode    | DescActions | A list of [DescAction][DescAction] values.
Encode    | DescActions | A list of [DescAction][DescAction] values.
