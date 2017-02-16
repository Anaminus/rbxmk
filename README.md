# rbxmk

`rbxmk` is a command-line tool for manipulating Roblox files.

The general workflow is that **inputs** are specified, transformed somehow,
then mapped to **outputs**.

## Installation

1. [Install Go](https://golang.org/doc/install)
2. [Install Git](http://git-scm.com/downloads)
3. Using a shell with Git (such as Git Bash), run the following command:

```
go get -u github.com/anaminus/rbxmk
```

If you installed Go correctly, this will install rbxmk to `$GOPATH/bin`,
which will allow you run it directly from a shell.

## Command options

Options are grouped together as "nodes". Certain flags delimit nodes. For
example, the `-i` flag delimits an input node, and also specifies a reference
for that node. The `-o` flag delimits an output node, also defining a
reference. All flags specified before a delimiting flag are counted as being
apart of the node. All flags after a delimiter will be apart of the next node.

Several flags, like `--id`, specify information for the node they are apart
of.

Other flags, like `--options`, are global; they do not belong to any
particular node, and may be specified anywhere.

In general, any flag may be specified multiple times. Flags are read from left
to right. If the option requires a single value, then only the last flag will
be counted.

Global Flags     | Description
-----------------|------------
`-h`, `--help`   | Display a help message.
`--options FILE` | Specify a file containing options.
`--api FILE`     | Specify an API dump file, which may be used by file formats for more correct encoding.

Delimiter Flags          | Description
-------------------------|------------
`-i REF`, `--input REF`  | Define the reference of the current node. Delimits an input node.
`-o REF`, `--output REF` | Define the reference of the current node. Delimits an output node.

Node Flags        | Description
------------------|------------
`--id STRING`     | Force ID of the current node.
`--map MAPPING`   | Map input nodes to output nodes.
`--format STRING` | Force the file format of the current node.

### Options file

The `--options` flag loads options from a file. It can be specified any number
of times.

Any options (including `--options`) may be specified in the file. The options
are applied immediately, so the order of the options matters. `--options` can
be treated as expanding into the options within the file.

Option files cannot be recursive; a file may only be loaded once.

#### File syntax

Each line specifies one option. Leading and trailing whitespace is ignored, as
well as empty lines. Lines beginning with a `#` character are ignored.

Within a line, the name of the option occurs up to the first whitespace. After
any whitespace, the remaining text is the option argument.

```
api path/to/file
options path/to/suboptions

# comment
i inputReference
o outputReference
```

## Node IDs

IDs are used to refer to nodes for further processing, such as mapping. IDs
are case-sensitive, and may only contain letters, digits, and `_`.

IDs can be assigned manually by the user with the `-id` flag. For example,

```
-id first -i input
```

The first input node now has an ID of `first`.

If a node is not assigned an ID manually, it is automatically given a
numerical name. Unassigned nodes are given the next available number for the
type of node. For example, the first unassigned input node will have an ID of
`0`. The next input node will have an ID of `1`, and so on.

## Mapping

Inputs must be mapped to outputs somehow. By default, when no mapping is
defined, each input is mapped to each output. This can be modified with the
`-map` flag.

To start, let's define some nodes:

```
-id A -i refA
-id B -i refB
-id C -i refC
-id X -o refX
-id Y -o refY
-id Z -o refZ
```

So there are 3 inputs: A, B, and C, then 3 outputs: X, Y, and Z.

The value of the `-map` flag has a certain syntax. In general, there are
*left* and *right* sides of the mapping, which are separated by a `:`
character. The left side selects inputs, while the right side selects outputs.
For example, input A can be mapped to output X like so:

```
-map A:X    Map A to X
                A -> X
```

On either side, multiple nodes can be selected with a `,` character:

```
-map A,B,C:X    Map A, B, and C to X
                    A -> X
                    B -> X
                    C -> X
-map A:X,Y,Z    Map A to X, Y, and Z
                    A -> X
                    A -> Y
                    A -> Z
-map A,B:X,Y    Map A and B to X and Y
                    A -> X
                    A -> Y
                    B -> X
                    B -> Y
```

The `*` character will select all nodes for the current side.

```
-map *:X    Map each input to X
                A -> X
                B -> X
                C -> X
-map A:*    Map A to each output
                A -> X
                A -> Y
                A -> Z
```

The mapping of a node may be negated by using a `-` character before the node.
This causes an input to no longer be mapped to an output.

```
-map "-A:X"     Negate mapping of A to X
-map "A:-X"     Negate mapping of A to X (same result)
```

Operators may be combined to create more complex mappings.

```
-map "*,-B:X"   Map each input, except B, to X
                    A -> X
                    C -> X
-map "-*:X"     Negate mapping of each input to X (same as "*:-X")
```

As stated before, when no mappings are specified, the default is that each
input is mapped to each output. This is equivalent to `*:*`:

```
-map *:*    Map each input to each output
                A -> X
                B -> X
                C -> X
                A -> Y
                B -> Y
                C -> Y
                A -> Z
                B -> Z
                C -> Z
```

When a `-map` flag is given inside a node, that node can be used for the
mapping. The side which is the same as the node type may be omitted. In this
case, the value of the `-map` specifies the opposite side of the mapping.

```
-id A -map X -i refA         Map input A to output X
-id X -map "*,-B" -o refX    Map each input, except B, to X
```

Note that an input can be mapped to an output only once. For the mapping
`*,-B:X`, each input is mapped to X, and then B is no longer mapped to X. This
overrides any previous mappings for the inputs and outputs involved.

```
-map "*:X,Y"    Each input is mapped to X and Y
-map "-B:X"     B is no longer mapped to X, but is still mapped to Y
-map "B:X,-Y"   B is now mapped to X, but no longer mapped to Y
-map "-B:*"     B is no longer mapped to any output
```

## References

References are the way to specify where data will be read from for input
nodes, and where data will be written to for output nodes.

A reference is used to select a **file**, which is retrieved from some
location. This selection may be further refined by selecting **data** within
the file, or even further by selecting data within the data.

Files have particular formats. Several formats are supported for either
inputs, outputs, or both. The format of a file can usually be determined
automatically. For example, by the file extension. The format of a file
referred to by a node may also be forced directly, with the `-format` flag.

When data is retrieved from a file, it is formatted in a generic way. There
are three types of data:

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

Drilling is indicated by a `:` character, called the **drill separator**,
which delimits parts of a reference string.


### Syntax

References have a syntax. It begins with a URI-like scheme:

```
scheme://
```

The next part of the reference is used to locate a file. The syntax of the
next part depends on the scheme. Schemes may be defined for either inputs,
outputs, or both.

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

The format of the selected file, if not forced, is determined by the file extension.

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
