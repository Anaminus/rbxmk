# Summary
Describes a function class member.

# Description
The **FunctionDesc** type is a table that describes a function member of a
class. It has the following fields:

Field      | Type                             | Description
-----------|----------------------------------|------------
MemberType | [string](##)                     | Indicates the type of member. Always "Function".
Name       | [string](##)                     | The name of the member.
ReturnType | [TypeDesc][TypeDesc]             | The type of the value returned by the function.
Security   | [string](##)                     | The security context required to set the member.
Parameters | {[ParameterDesc][ParameterDesc]} | The parameters of the function.
Tags       | {[string](##)}                   | The tags set for the member.
