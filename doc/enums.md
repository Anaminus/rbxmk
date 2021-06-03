# Enums
This document contains a reference to the enums available to rbxmk scripts.
References to them are available under
[rbxmk.Enum](libraries.md#user-content-rbxmkenum).


<table>
<thead><tr><th>Table of Contents</th></tr></thead>
<tbody><tr><td>

1. [DescActionElement][DescActionElement]
2. [DescActionType][DescActionType]

</td></tr></tbody>
</table>

This document only describes the enums defined by rbxmk. Roblox enums can be
specified by [descriptors](README.md#user-content-descriptors).

## DescActionElement
[DescActionElement]: #user-content-descactionelement

The **DescActionElement** enum indicates the type of descriptor element to which
a [DescAction][DescAction] applies. It has the following items:

Name     | Value | Description
---------|------:|------------
Class    |     0 | Applies to a [ClassDesc][ClassDesc].
Property |     1 | Applies to a [PropertyDesc][PropertyDesc].
Function |     2 | Applies to a [FunctionDesc][FunctionDesc].
Event    |     3 | Applies to a [EventDesc][EventDesc].
Callback |     4 | Applies to a [CallbackDesc][CallbackDesc].
Enum     |     5 | Applies to a [EnumDesc][EnumDesc].
EnumItem |     6 | Applies to a [EnumItemDesc][EnumItemDesc].

[DescAction]: types.md#user-content-descaction
[CallbackDesc]: types.md#user-content-callbackdesc
[ClassDesc]: types.md#user-content-classdesc
[EnumDesc]: types.md#user-content-enumdesc
[EnumItemDesc]: types.md#user-content-enumitemdesc
[EventDesc]: types.md#user-content-eventdesc
[FunctionDesc]: types.md#user-content-functiondesc
[PropertyDesc]: types.md#user-content-propertydesc

## DescActionType
[DescActionType]: #user-content-descactiontype

The **DescActionType** enum indicates the type of transformation performed by a
[DescAction][DescAction]. It has the following items:

Name   | Value | Description
-------|------:|------------
Remove |    -1 | Removes the referred element.
Change |     0 | Changes the referred element.
Add    |     1 | Adds the referred element.
