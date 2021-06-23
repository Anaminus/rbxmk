# Summary
Describes a callback.

# Description
The **CallbackDesc** type is a table that describes a callback member of a
class. It has the following fields:

Field      | Type                             | Description
-----------|----------------------------------|------------
MemberType | [string](##)                     | Indicates the type of member. Always "Callback".
Name       | [string](##)                     | The name of the member.
Parameters | {[ParameterDesc][ParameterDesc]} | The parameters of the callback.
ReturnType | [TypeDesc][TypeDesc]             | The type of the value returned by the callback.
Security   | [string](##)                     | The security context required to set the member.
Tags       | {[string](##)}                   | The tags set for the member.
