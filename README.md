# rbxmk

`rbxmk` is a command-line tool for manipulating Roblox files.

The general workflow is that **inputs** are specified, transformed somehow,
then mapped to **outputs**.

[Lua](https://lua.org) scripts are used to perform actions. Scripts are run in
a stripped-down environment, with only a small set of functions available.

## Installation

1. [Install Go](https://golang.org/doc/install)
2. [Install Git](http://git-scm.com/downloads)
3. Using a shell with Git (such as Git Bash), run the following command:

```
go get -u github.com/anaminus/rbxmk
```

If you installed Go correctly, this will install rbxmk to `$GOPATH/bin`,
which will allow you run it directly from a shell.

This document uses POSIX-style flags (`-f`, `--flag`), although windows-style
flags (`/f`, `/flag`) are possible when rbxmk is compiled for Windows. If you
are compiling for Windows, you may choose to force POSIX-style flags with the
`forceposix` build tag:

```
go get -u -tags forceposix github.com/anaminus/rbxmk
```

For more information, see the [go-flags](https://godoc.org/github.com/jessevdk/go-flags) package.

## Command options

Since almost all actions are done in Lua, there are only a few command
options:

```
rbxmk [OPTIONS...] [- | SCRIPT]
```

Options          | Description
-----------------|------------
`-h`, `--help`   | Display a help message.
`-f`, `--file`   | Run a script from a file. This may be given any number of times.

Passing `-` runs a script from stdin. Otherwise, the first non-option
(`SCRIPT`) will run as a script. Scripts run in the order they are given. `-`
or `SCRIPT` can only be given after all other options.

## Lua environment

None of the regular base functions and libraries are loaded. Instead, a small
set of functions are made available. Each of these functions accept a table as
their only argument.

Lua has the following bit of syntax sugar to support this: if a constructed
table is the only argument to a function, then the function's parentheses may
be omitted:

```lua
func({arg1, arg2})
func{arg1, arg2}
-- Both calls are equivalent
```

Henceforth, an "argument" will refer to a value within this kind of table.

Because of how tables work, there can be two kinds of arguments: named and
unnamed.

```lua
-- Unnamed arguments
func{value, value}

-- Named arguments
func{arg1=value, arg2=value}
```

### Functions

The following functions are available:

Name                             | Description
---------------------------------|------------
[`input`](#user-content-input)   | Create an input node.
[`output`](#user-content-output) | Create an output node.
[`map`](#user-content-map)       | Map one or more inputs to one or more outputs.
[`filter`](#user-content-filter) | Transform nodes.
[`load`](#user-content-load)     | Load and execute a script.
[`type`](#user-content-type)     | Return the type of a value as a string.
[`error`](#user-content-error)   | Create an error node.
[`exit`](#user-content-exit)     | Force the program to exit.

#### input

`node = input{format=string, ...string}`

`input` creates an input node. The arguments specify the
[reference](#user-content-references) to the input. The `format` argument
forces the file format, if needed.

#### output

```lua
node = output{format=string, ...string}
```

`output` creates an output node. The arguments specify the
[reference](#user-content-references) to the output. The `format` argument
forces the file format, if needed.

#### map

`node = map{...node}`

`map` maps one or more input nodes to one or more output nodes. Any kind of
node may be passed to `map` (error nodes will be ignored). If an error occurs,
`map` returns an error node.

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

#### filter

`... = filter{string, ...}`

`filter` transforms nodes. The first argument specifies the filter name.
Subsequent arguments and return values depend on the filter. See
[Filters](#user-content-filters) for a list of filters and their arguments.

#### load

`... = load{string, ...}`

`load` executes a file as a script. The first argument is the file name.
Subsequent arguments are passed to the script (which may be received with the
`...` operator). If an error occurs, `load` returns an error node. Otherwise,
`load` returns any values returned by the script.

#### type

`string = type{value}`

`type` returns the type of the given value as a string. In addition to the
regular Lua types, the following types are available:

- `input`: an input node.
- `output`: an output node.
- `error`: an error node.

#### error

`node = error{string}`

`error` creates an error node, with the first argument as the error message.

#### exit

`exit{node}`

`exit` forces the program to exit. An optional error node can be given, which
will be passed to the program.

## References

References are the way to specify where data will be read from for input
nodes, and where data will be written to for output nodes.

A reference is used to select a **file**, which is retrieved from some
location. This selection may be further refined by selecting **data** within
the file, or even further by selecting data within the data.

Files have particular formats. Several formats are supported for either
inputs, outputs, or both. The format of a file can usually be determined
automatically (e.g. the file extension). The format of a file referred to by a
node may also be forced directly, with the `format` argument.

When data is selected from a file, it is formatted in a generic way. There are
three types of data:

- **Instances**: These are the usual Roblox instances, consisting of a
  ClassName, a set of properties, and a list of child Instances.
- **Properties**: These are like the properties in an Instance. They consist
  of a name, which is mapped to a value.
- **Values**: These are like the values of a property. They consist of a type,
  and the actual value.

In order to refine a selection, the concept of **drilling** is introduced.
Depending on the type of data, drilling can be used to further specify the
selected data. For example, if the data is an Instance, then you can drill
down into the Instance to select a property.

In terms of input, drilling specifies the data that is retrieved. For output,
drilling determines where data will go.

The [`input`](#user-content-input) and [`output`](#user-content-output)
functions specify that their unnamed arguments make up the returned node's
reference. Each successive argument drills down into the data of the previous
argument.

### Syntax

The first argument of a reference indicates a file location, and has a
specific syntax. It begins with a URI-like scheme:

```
scheme://
```

The rest of the argument is used to locate a file. The syntax of this part
depends on the scheme. Schemes may be defined for either inputs, outputs, or
both.

If no scheme is specified, then the reference is assumed to be of the
`file://` scheme.

### file://

The `file` scheme is used to refer to files on the operating system. It is
defined for both inputs and outputs.

The syntax is simply a path to a file on the operating system.

```
file://C:/Users/user/projects/project/file.rbxl
file:///home/user/projects/project/file.rbxl
```

Because the `file` scheme is the default, the scheme portion may be omitted
from the reference.

```
C:/Users/user/projects/project/file.rbxl
/home/user/projects/project/file.rbxl
```

The format of the selected file, if not forced, is determined by the file
extension.

### http://, https://

The `http/https` scheme retrieves files using the HTTP protocol.

The syntax is a standard URL.

```
http://www.example.com/path/to/file
```

Drilling into an output is disabled for this scheme, because it may not be
possible to receive data from the output location.

The format of the selected file, if not forced, is determined by the MIME type
given by the response.

## Formats

Extension | Data       | Description
----------|------------|------------
rbxl      | Instances  | Roblox place
rbxm      | Instances  | Roblox model
rbxlx     | Instances  | Roblox place in XML
rbxmx     | Instances  | Roblox model in XML
json      | Properties | Set of properties in JSON format
xml       | Properties | Set of properties in XML format
lua       | Value      | Lua source file
txt       | Value      | Normal text file
bin       | Value      | binary file