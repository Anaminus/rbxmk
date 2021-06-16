# Summary
Describes an event class member.

# Description
The **EventDesc** type is a table that describes a event member of a
class. It has the following fields:

Field      | Type                             | Description
-----------|----------------------------------|------------
MemberType | [string](##)                     | Indicates the type of member. Always "Event".
Name       | [string](##)                     | The name of the member.
Parameters | {[ParameterDesc][ParameterDesc]} | The parameters of the event.
Security   | [string](##)                     | The security context required to get the member.
Tags       | {[string](##)}                   | The tags set for the member.
