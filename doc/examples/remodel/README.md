# rbxmk vs Remodel
[Remodel][remodel] is a tool similar to rbxmk. It also runs Lua scripts, and has
its own Lua API. This document lists the Lua APIs provided by Remodel, and their
equivalents in rbxmk.

[remodel]: https://github.com/rojo-rbx/remodel

## Descriptors
Remodel has API descriptors built-in, whereas descriptors in rbxmk must be
specified by the user.

To fetch the latest descriptors provided by Roblox, the `--desc-latest` flag can
be passed. Additionally, the [dump-patch](../dump-patch) file can be applied
afterwards to further improve compatibility. For example:

```bash
rbxmk run --desc-latest --desc-patch dump.desc-patch.json script.lua
```

## Instances
Most methods for navigating instance trees are available in both Remodel and
rbxmk.

Roblox | Remodel | rbxmk
-------|---------|------
`Instance.new`            | The *parent* argument is not supported. | *parent* is supported. Has third optional argument to override the instance's descriptor.
`Instance.Name`           | Equivalent | Equivalent
`Instance.ClassName`      | Equivalent | Can be written, but no transformations are made to properties.
`Instance.Parent`         | Equivalent | Equivalent
`Instance.Destroy`        | Equivalent | Equivalent
`Instance.Clone`          | Equivalent | Equivalent
`Instance.GetChildren`    | Equivalent | Equivalent
`Instance.GetDescendants` | Equivalent | Equivalent
`Instance.FindFirstChild` | *recursive* argument is not supported. | *recursive* is supported.
`DataModel.GetService`    | Inserts service if it does not exist. | Inserts service if it does not exist.

## Types
rbxmk supports all property types available in Roblox, which includes all types
supported by Remodel. In rbxmk, while a descriptor is set, value type
conversions are performed automatically when getting or setting a property.
Without a descriptor, the [types][types] library can be used to specify a type
that would otherwise be ambiguous (for example, differentiating a number between
an int, float or double).

[types]: https://github.com/Anaminus/rbxmk/blob/imperative/doc/libraries.md#user-content-types

## Authentication
Remodel can be passed the `--auth` flag to receive a cookie for authentication.
APIs that require authentication automatically use this cookie when it is
available.

In rbxmk, cookies must handled within the script through the [Cookie][Cookie]
type. APIs requiring authentication have an option to receive these cookies.

rbxmk can attempt to retrieve cookies from the user's Studio session using the
[Cookie.from][Cookie.from] function:

```lua
-- Returns a list of cookies.
local cookies = Cookie.from("studio")
```

Cookies can also be created from scratch with the [Cookie.new][Cookie.new]
function. This can be combined with the [os.getenv][os.getenv] function to
emulate Remodel's behavior with the REMODEL_AUTH environment variable:

```lua
local authCookie = Cookie.new(".ROBLOSECURITY", os.getenv("REMODEL_AUTH") or "")
```

Or, the cookie value can be received as an argument to the script, which is
similar to the `--auth` flag:

```lua
-- From first argument.
local authCookie = Cookie.new(".ROBLOSECURITY", (...))
```

[Cookie]: https://github.com/Anaminus/rbxmk/blob/imperative/doc/types.md#user-content-cookie
[Cookie.from]: https://github.com/Anaminus/rbxmk/blob/imperative/doc/types.md#user-content-cookiefrom
[Cookie.new]: https://github.com/Anaminus/rbxmk/blob/imperative/doc/types.md#user-content-cookienew
[os.getenv]: https://github.com/Anaminus/rbxmk/blob/imperative/doc/libraries.md#user-content-osgetenv

## Library functions
The following section lists each Remodel API and its equivalent in rbxmk.

### remodel.readPlaceFile
remodel:
```lua
local game = remodel.readPlaceFile("file.rbxl")
local game = remodel.readPlaceFile("file.rbxlx")
```

rbxmk:
```lua
local game = fs.read("file.rbxl") -- Implicit format from extension.
local game = fs.read("file.rbxlx") -- Implicit format from extension.
local game = fs.read("file.rbxl", "rbxl") -- Explicit format.
local game = fs.read("file.rbxlx", "rbxlx") -- Explicit format.
```

### remodel.readModelFile
remodel:
```lua
local model = remodel.readPlaceFile("file.rbxm")
local model = remodel.readPlaceFile("file.rbxmx")
```

rbxmk:
```lua
-- Returns a DataModel with top-level instances as children.
local model = fs.read("file.rbxm"):GetChildren() -- Implicit format from extension.
local model = fs.read("file.rbxmx"):GetChildren() -- Implicit format from extension.
local model = fs.read("file.rbxm", "rbxm"):GetChildren() -- Explicit format.
local model = fs.read("file.rbxmx", "rbxmx"):GetChildren() -- Explicit format.
```

### remodel.readPlaceAsset
remodel:
```lua
local game = remodel.readPlaceAsset("123456789")
```

rbxmk:
```lua
-- No authentication.
local game = rbxassetid.read({AssetId=123456789, Format="rbxl"})

-- Using REMODEL_AUTH environment variable for authentication.
local auth = {Cookie.new(".ROBLOSECURITY", os.getenv("REMODEL_AUTH") or "")}
local game = rbxassetid.read({AssetId=123456789, Format="rbxl", Cookies=auth})
```

### remodel.readModelAsset
remodel:
```lua
local model = remodel.readModelAsset("123456789")
```

rbxmk:
```lua
-- No authentication.
local model = rbxassetid.read({AssetId=123456789, Format="rbxm"}):GetChildren()

-- Using REMODEL_AUTH environment variable for authentication.
local auth = {Cookie.new(".ROBLOSECURITY", os.getenv("REMODEL_AUTH") or "")}
local model = rbxassetid.read({AssetId=123456789, Format="rbxm", Cookies=auth}):GetChildren()
```

### remodel.writePlaceFile
remodel:
```lua
remodel.writePlaceFile(game, "file.rbxl")
remodel.writePlaceFile(game, "file.rbxlx")
```

rbxmk:
```lua
-- game may be a DataModel, Instance, or array of Instances.
fs.write("file.rbxl", game) -- Implicit format from extension.
fs.write("file.rbxlx", game) -- Implicit format from extension.
fs.write("file.rbxl", game, "rbxl") -- Explicit format.
fs.write("file.rbxlx", game, "rbxlx") -- Explicit format.
```

### remodel.writeModelFile
remodel:
```lua
remodel.writeModelFile(instance, "file.rbxm")
remodel.writeModelFile(instance, "file.rbxmx")
```

rbxmk:
```lua
-- instance may be a DataModel, Instance, or array of Instances.
fs.write("file.rbxm", instance) -- Implicit format from extension.
fs.write("file.rbxmx", instance) -- Implicit format from extension.
fs.write("file.rbxm", instance, "rbxm") -- Explicit format.
fs.write("file.rbxmx", instance, "rbxmx") -- Explicit format.
```

### remodel.writeExistingPlaceAsset
remodel:
```lua
remodel.writeExistingPlaceAsset(game, "123456789")
```

rbxmk:
```lua
-- Using REMODEL_AUTH environment variable for authentication.
local auth = {Cookie.new(".ROBLOSECURITY", os.getenv("REMODEL_AUTH") or "")}
rbxassetid.write({AssetId=123456789, Format="rbxl", Body=game})
```


### remodel.writeExistingModelAsset
remodel:
```lua
remodel.writeExistingModelAsset(instance, "123456789")
```

rbxmk:
```lua
-- Using REMODEL_AUTH environment variable for authentication.
local auth = {Cookie.new(".ROBLOSECURITY", os.getenv("REMODEL_AUTH") or "")}
rbxassetid.write({AssetId=123456789, Format="rbxm", Body=instance})
```

### remodel.getRawProperty
remodel:
```lua
local value = remodel.getRawProperty(instance, "Property")
```

rbxmk:
```lua
-- Get property. May have descriptor validation depending on
-- `instance[sym.Desc]`.
local value = instance.Property

-- Ensure validation is disabled.
instance[sym.RawDesc] = false
local value = instance.Property
```

### remodel.setRawProperty
remodel:
```lua
remodel.setRawProperty(instance, "Property", "String", value)
```

rbxmk:
```lua
-- Set property. May have descriptor validation depending on
-- `instance[sym.Desc]`.
instance.Property = value

-- Ensure validation is disabled.
instance[sym.RawDesc] = false
instance.Property = value
```

### remodel.readFile
remodel:
```lua
local content = remodel.readFile(path)
```

rbxmk:
```lua
-- Read raw content of file as raw bytes.
local content = fs.read(path, "bin")
```

### remodel.readDir
remodel:
```lua
local files = remodel.readDir(path)
```

rbxmk:
```lua
-- fs.dir returns an array of tables with Name and IsDir fields.
local files = fs.dir(path)
for i, file in ipairs(files) do
	files[i] = file.Name
end
```

### remodel.writeFile
remodel:
```lua
remodel.writeFile(path, content)
```

rbxmk:
```lua
-- Write content to file as raw bytes.
fs.write(path, content, "bin")
```

### remodel.createDirAll
remodel:
```lua
remodel.createDirAll(path)
```

rbxmk:
```lua
-- Second argument causes parent directories to be created as needed.
fs.mkdir(path, true)
```

### remodel.isFile
remodel:
```lua
local isFile = remodel.isFile(path)
```

rbxmk:
```lua
-- fs.stat returns metadata of file.
local stat = fs.stat(path)
if stat == nil then
	error("file does not exist")
end
local isFile = not stat.IsDir
```

### remodel.isDir
remodel:
```lua
local isDir = remodel.isDir(path)
```

rbxmk:
```lua
-- fs.stat returns metadata of file.
local stat = fs.stat(path)
if stat == nil then
	error("file does not exist")
end
local isDir = stat.IsDir
```

### json.fromString
remodel:
```lua
local data = json.fromString(source)
```

rbxmk:
```lua
-- Decode string content according to json format.
local data = rbxmk.decodeFormat("json", source)
```

### json.toString
remodel:
```lua
local output = json.toString(data)
```

rbxmk:
```lua
-- Encode value according to json format. Default configuration produces
-- minified JSON.
local output = rbxmk.encodeFormat("json", data)
```

### json.toStringPretty
remodel:
```lua
local output = json.toStringPretty(data)
local output = json.toStringPretty(data, "\t")
```

rbxmk:
```lua
-- Encode value according to json format, with Indent option set to use two
-- spaces.
local output = rbxmk.encodeFormat({Format="json", Indent="  "}, data)
-- Using tabs for indentation.
local output = rbxmk.encodeFormat({Format="json", Indent="\t"}, data)
```
