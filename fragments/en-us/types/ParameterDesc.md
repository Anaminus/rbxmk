# Summary
Describes a parameter.

# Description
The **ParameterDesc** type describes a parameter of a function, event, or
callback member. It has the following members:

Field   | Type                 | Description
--------|----------------------|------------
Type    | [TypeDesc][TypeDesc] | The type of the parameter.
Name    | [string](##)         | The name of the parameter.
Default | [string](##)?        | Describes the default value of the parameter. If nil, then the parameter has no default value.
