[DRAFT]

# Usage
This document provides an overview of how rbxmk is used. For a detailed
reference, see [DOCUMENTATION.md](DOCUMENTATION.md).

## Command options

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

Remaining arguments are Lua values to be passed to the file. Numbers, bools, and
nil are parsed into their respective types in Lua, and any other value is
interpreted as a string.

```bash
rbxmk script.lua true 3.14159 hello!
```

Within the script, these arguments can be received from the `...` operator:

```lua
local arg1, arg2, arg3 = ...
```

## Lua environment
[Lua](https://lua.org) scripts are used to perform actions. Scripts are run in a
modified environment, where most of the standard functions and libraries are
available. A complete list of included items can be found in
[DOCUMENTATION.md](DOCUMENTATION.md#user-content-standard-library).

Also included is an environment similar to the [Roblox Lua API][apiref],
allowing the user to construct and manipulate objects with a familiar API.

Several additional libraries are included to aid the user's workflow:

- `rbxmk`: General auxiliary functions.
- `os`: Additional operating system functions.
- Various source libraries that provide interfaces to external sources.
<!-- TODO
- file: An interface to the file system.
- http: An interface to resources via HTTP.
- asset: An interface to assets in the Roblox ecosystem.
- clipboard: An interface to the operating system's clipboard.
- studio: An interface to active Roblox Studio sessions.
-->

[apiref]: https://developer.roblox.com/en-us/api-reference

## Workflow
Operating on objects within rbxmk is very similar to operating on objects within
Roblox. Instances can be created and modified. Properties can be retrieved and
updated. Values can be read from and written to external sources.

```lua
local folder = Instance.new("Folder")
folder.Name = "Modules"

local module = Instance.new("ModuleScript")
module.Name = "Hello"
module.Source = [[return "Hello world!"]]
module.Parent = folder
```

### Instances
A major difference between Roblox and rbxmk is what an instance represents. In
Roblox, an instance is a live object that acts and reacts. In rbxmk, an instance
represents *data*, and only data.

Consider the RBXL file format. Files of this format contain information used to
reconstruct the instances that make up a place or model. Such files are static:
they contain only data, but are difficult to manipulate in place. Instances in
rbxmk are like this, except that they are also interactive: the user can freely
modify data and move it around.

Because of this, there are several differences between the Roblox API and the
rbxmk API. For one, any kind of class can be created. Instances are just data,
including the class name. The ClassName property can even be assigned to.

```lua
local foobar = Instance.new("Foobar")
foobar.ClassName = "FizzBuzz" -- allowed
```

Instances also have no defined properties. A value of any type can be assigned
to any property. Likewise, properties that are not explicitly assigned
effectively do not exist.

```lua
local part = Instance.new("Part")
part.Foobar = 42 -- allowed
print(part.Position) --> nil
```

That said, even though it is possible for rbxmk to create arbitrary classes with
arbitrary properties, this does not mean such instances will be interpreted in
any meaningful way when sent over to Roblox. If the user wants to use such data
in Roblox, it is up to them to make sure the data is correct according to
Roblox's API.

Instances in rbxmk implement a subset of Roblox's Instance members:

- `.ClassName`
- `.Name`
- `.Parent`
- `:ClearAllChildren()`
- `:Clone()`
- `:Destroy()`
- `:FindFirstAncestor()`
- `:FindFirstAncestorOfClass()`
- `:FindFirstChild()`
- `:FindFirstChildOfClass()`
- `:GetChildren()`
- `:GetDescendants()`
- `:GetFullName()`
- `:IsAncestorOf()`
- `:IsDescendantOf()`

### DataModel
rbxmk has no singular game tree that contains all objects. Instead, any number
of **DataModel** objects can be created with the `DataModel.new()` function. The
resulting value is an Instance like any other, except that the class name cannot
be assigned to.

```lua
local game = DataModel.new()
foobar.Parent = game
game.ClassName = "Foobar" -- not allowed
```

Additionally, the properties on a DataModel object are interpreted as metadata,
which may be needed in certain cases.

```lua
game.ExplicitAutoJoints = true
```

For convenience, DataModels also define the `GetService` method. This creates a
child instance of the given class, and marks the instance as a service. If such
an instance already exists in the DataModel, then that instance will be returned
instead.

```lua
local workspace = game:GetService("Workspace")
print(workspace == game:GetService("Workspace")) --> true
```

### Data types
Also included in the rbxmk environment are constructors for the various Roblox
types, such as `Vector3`, `CFrame`, and `UDim2`.

```lua
local part = Instance.new("Part")
part.Position = Vector3.new(10, 20, 30)
```

Certain types have an ambiguous representation in Roblox Lua. In rbxmk, this is
remedied by defining constructors for such types. These return userdata values
that are used exclusively for assigning to properties to ensure the correct type
is encoded or decoded. When getting such a value, the original value of the base
type is returned instead of a userdata.

The following numeric type constructors are available:
- `float.new(number)`
- `int.new(number)`
- `int64.new(number)`
- `token.new(number)` (enums)

Normal Lua numbers are interpreted as the `double` type.

The following string type constructors are available:
- `BinaryString.new(string)`
- `ProtectedString.new(string)`
- `Content.new(string)`
- `SharedString.new(string)`

Normal Lua strings are interpreted as the `string` type.

```lua
local value = Instance.new("BinaryValue")
value.Name = "BinaryValue"
value.Value = BinaryValue.new("\0\1\2\3")
print(value.Value) --> "\0\1\2\3" (as a Lua string)
```

<!-- TODO
For now, Enums are not included. However, properties can still be assigned using
the `token.new()` constructor with the numeric value of the enum item.

```lua
part.TopSurface = token.new(0)
part.BottomSurface = token.new(0)
```
-->

## Formats
A principle concept in rbxmk is **formats**. A format describes how a Lua value
can be encoded between a raw sequence of bytes. For example, in rbxmk, the
"rbxl" format can encode a tree of instances into a format compatible with
Roblox's RBXL file format.

Formats can be accessed at a low level through the `rbxmk.encodeformat` and
`rbxmk.decodeformat` functions.

```lua
local folder = Instance.new("Folder")
folder.Name = "Modules"

local data = rbxmk.encodeformat("rbxmx", folder) --> string containing XML model data
local folderCopy = rbxmk.decodeformat("rbxmx", data) --> Folder instance
print(folderCopy.Name) --> Folder
```

## Sources
Another principle concept is **sources**. A source is an external location from
which a stream of bytes can be read or written to. For example, the "file"
source allows encoded data to be read from or written to the file system.

Sources can be accessed at a low level through the `rbxmk.readsource` and
`rbxmk.writesource` functions.

```lua
local folder = Instance.new("Folder")
folder.Name = "Modules"

local data = rbxmk.encodeformat("rbxmx", folder)
rbxmk.writesource("file", data, "folder.rbxmx") -- data written to folder.rbxmx

local dataCopy = rbxmk.readsource("file", "folder.rbxmx") -- dataCopy read from folder.rbxmx
local folderCopy = rbxmk.decodeformat("rbxmx", dataCopy)
print(folderCopy.Name) --> Folder
```

## Libraries
rbxmk defines several additional libraries.

### rbxmk
The **rbxmk** library contains general auxiliary functions.

The `rbxmk.load` function is used to run another script from within the current
script.

```lua
local result = rbxmk.load("subscript.lua", true, 3.14159, "hello!")
```

### os
The **os** library merges into the standard os library by including several
additional functions.

The `os.split` function can be used split a file path into its components.

```lua
print(os.split("projects/project/model/script.lua", "dir", "stem", "ext"))
--> "projects/project/model", "script", ".lua"
```

The `os.join` function can be used to correctly join two file paths together.

```lua
print(os.join("projects/project/model", "script.lua"))
--> "projects/project/model/script.lua"
```

The `os.expand` function expands particular variables within a path to their
full representation.

```lua
print(os.expand("$script_name")) --> "build.lua"
print(os.expand("$script_directory")) --> "projects/project"
```

The `os.getenv` function gets an environment variable.

```lua
print(os.getenv("HOME")) --> "/home/user"
```

The `os.dir` function returns the files within a directory.

```lua
for _, file in ipairs(os.dir("projects/project")) do
	print(file.name)
	print(file.isdir)
	print(file.size)
	print(file.modtime)
end
```

### file
The **file** library provides a convenient interface to the file system.

The `file.read` function reads from a file.

```lua
local game = file.read("place.rbxl") -- "rbxl" format is inferred from extension
local game = file.read("place.rbxl", "rbxlx") -- force "rbxlx" format
```

The `file.write` function writes to a file.

```lua
file.write("place.rbxl", game)
file.write("place.rbxl", "rbxlx", game)
```
