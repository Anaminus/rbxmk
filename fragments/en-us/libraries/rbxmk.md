# Summary
An interface to the rbxmk engine, and the rbxmk environment.

# Description
The **rbxmk** library contains functions related to the rbxmk engine.

$FIELDS

# Fields
## Enum
### Summary
Collection of rbxmk-defined enums.

### Description
The **Enum** field is a collection of Enum values defined by rbxmk. The
following enums are defined:

$ENUMS

## decodeFormat
### Summary
Deserialize data from bytes.

### Description
The **decodeFormat** function decodes *bytes* into a value according to
*format*. The exact details of each format are described in the
[Formats](formats.md) documents.

decodeFormat will throw an error if the format does not exist, or the format has
no decoder defined.

## encodeFormat
### Summary
Serialize data into bytes.

### Description
The **encodeFormat** function encodes *value* into a sequence of bytes according
to *format*. The exact details of each format are described in the
[Formats](formats.md) document.

encodeFormat will throw an error if the format does not exist, or the format has
no encoder defined.

## formatCanDecode
### Summary
Check whether a format decodes into a type.

### Description
The **formatCanDecode** function returns whether *format* decodes into *type*.

formatCanDecode will throw an error if the format does not exist, or the format
does not define types it decodes into.

## globalAttrConfig
### Summary
Get or set the global AttrConfig.

### Description
The **globalAttrConfig** field gets or sets the global AttrConfig. Most items
that utilize an AttrConfig will fallback to the global AttrConfig when possible.

See the [Value inheritance](README.md#user-content-value-inheritance) section
for details on how this field is inherited by [Instances][Instance].

## globalDesc
### Summary
Get or set the global descriptor.

### Description
The **globalDesc** field gets or sets the global root descriptor. Most items
that utilize a root descriptor will fallback to the global descriptor when
possible.

See the [Value inheritance](README.md#user-content-value-inheritance) section
for details on how this field is inherited by [Instances][Instance].

## loadFile
### Summary
Load the content of a file as a function.

### Description
The **loadFile** function loads the content of a file as a Lua function. *path*
is the path to the file.

The function runs in the context of the calling script.

## loadString
### Summary
Load a string as a function.

### Description
The **loadString** function loads the a string as a Lua function. *source* is
the string to load.

The function runs in the context of the calling script.

## runFile
### Summary
Run a file as a Lua chunk.

### Description
The **runFile** function runs the content of a file as a Lua script. *path* is
the path to the file. *args* are passed into the script as arguments. Returns
the values returned by the script.

The script runs in the context of the referred file. Files cannot be run
recursively; if a file is already running as a script, then runFile will throw
an error.

## runString
### Summary
Run a string as a Lua chunk.

### Description
The **runString** function runs a string as a Lua script. *source* is the string
to run. *args* are passed into the script as arguments. Returns the values
returned by the script.

The script runs in the context of the calling script.
