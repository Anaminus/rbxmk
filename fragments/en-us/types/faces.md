# Summary
A set of active directions.

# Description
The **Faces** type represents a set of directions, on three orthogonal axes,
that are considered active. Corresponds to the [Faces][Faces-roblox] type in
Roblox.

[Faces-roblox]: https://developer.roblox.com/en-us/api-reference/datatype/Faces


# Constructors
## fromComponents
### Summary
Creates a new Faces with components set directly from each argument.

### Description
The **fromComponents** constructor returns a new Faces value, with each argument
setting the corresponding component.

## new
### Summary
Creates a new Faces with components set from a number of values.

### Description
The **new** constructor returns a new Faces value. Each valid argument sets a
component of the value, depending on the type:

- EnumItem:
	- Enum name is `"Axis"`, and item name is one of the following:
		- `"X"`: sets Right and Left.
		- `"Y"`: sets Top and Bottom.
		- `"Z"`: sets Back and Front.
	- Enum name is `"NormalId"`, and item name is `"Right"`, `"Top"`, `"Back"`,
	  `"Left"`, `"Bottom"`, or `"Front"`.
- number: value is one of the following:
	- `0`: sets Right.
	- `1`: sets Top.
	- `2`: sets Back.
	- `3`: sets Left.
	- `4`: sets Bottom.
	- `5`: sets Front.
- string: value is `"Right"`, `"Top"`, `"Back"`, `"Left"`, `"Bottom"`, or
  `"Front"`.

Other values will be ignored.

# Properties
## Back
### Summary
Whether the back face is active.

### Description
The **Back** field returns whether the back face is active.

## Bottom
### Summary
Whether the bottom face is active.

### Description
The **Bottom** field returns whether the bottom face is active.

## Front
### Summary
Whether the front face is active.

### Description
The **Front** field returns whether the front face is active.

## Left
### Summary
Whether the left face is active.

### Description
The **Left** field returns whether the left face is active.

## Right
### Summary
Whether the right face is active.

### Description
The **Right** field returns whether the right face is active.

## Top
### Summary
Whether the top face is active.

### Description
The **Top** field returns whether the top face is active.

# Operators
## Eq
### Summary
Returns whether two Faces values are equal.

### Description
The **equal** operator returns true if both operands are Faces, and each
corresponding component is equal.


