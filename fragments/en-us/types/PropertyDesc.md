# Summary
Describes a property class member.

# Description
The **PropertyDesc** type is a table that describes a property member of a
class. It has the following fields:

Field         | Type                 | Description
--------------|----------------------|------------
CanLoad       | [string](##)         | Whether the property is deserialized.
CanSave       | [string](##)         | Whether the property is serialized.
MemberType    | [string](##)         | Indicates the type of member. Always "Property".
Name          | [string](##)         | The name of the member.
ReadSecurity  | [string](##)         | The security context required to get the member.
Tags          | {[string](##)}       | The tags set for the member.
ValueType     | [TypeDesc][TypeDesc] | The type of the value of the property.
WriteSecurity | [string](##)         | The security context required to set the member.
