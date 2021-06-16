# Summary
A collection of DescAction fields.

# Description
The **DescFields** type maps is a collection of DescAction fields. Each element
is a pair that maps a field name to a value.

The following value types are allowed:

Type          | Notes
--------------|------
bool          |
number        |
string        |
TypeDesc      | Table must contain "Category" and "Name" fields.
ParameterDesc | Field name must be "Parameters".
{string}      | Field name must be "Tags".
