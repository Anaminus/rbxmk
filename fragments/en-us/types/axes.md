# Summary
A set of active axes.

# Description
The **Axes** type represents a set of orthogonal coordinate axes that are
considered active. Corresponds to the [Axes][Axes-roblox] type in Roblox.

[Axes-roblox]: https://developer.roblox.com/en-us/api-reference/datatype/Axes

# Constructors
## fromComponents
### Summary
Creates a new Axes with components set directly from each argument.

### Description
The **fromComponents** constructor returns a new Axes value, with each
argument setting the corresponding component.

## new
### Summary
Creates a new Axes with components set from a number of values.

### Description
The **new** constructor returns a new Axes value. Each valid argument sets a
component of the value, depending on the type:

- EnumItem:
	- Enum name is `"Axis"`, and item name is `"X"`, `"Y"`, or `"Z"`.
	- Enum name is `"NormalId"`, and item name is one of the following:
		- `"Right"`, `"Left"`: sets X.
		- `"Top"`, `"Bottom"`: sets Y.
		- `"Back"`, `"Front"`: sets Z.
- number: value is one of the following:
	- `0`: sets X.
	- `1`: sets Y.
	- `2`: sets Z.
- string: value is `"X"`, `"Y"`, or `"Z"`.

Other values will be ignored.

# Properties
## Back
### Summary
Corresponds to the Z axes.

### Description
The **Back** field returns whether the Z axis is active.

## Bottom
### Summary
Corresponds to the Y axes.

### Description
The **Bottom** field returns whether the Y axis is active.

## Front
### Summary
Corresponds to the Z axes.

### Description
The **Front** field returns whether the Z axis is active.

## Left
### Summary
Corresponds to the X axes.

### Description
The **Left** field returns whether the X axis is active.

## Right
### Summary
Corresponds to the X axes.

### Description
The **Right** field returns whether the X axis is active.

## Top
### Summary
Corresponds to the Y axes.

### Description
The **Top** field returns whether the Y axis is active.

## X
### Summary
Whether the X axis is active.

### Description
The **X** field returns whether the X axis is active.

## Y
### Summary
Whether the Y axis is active.

### Description
The **Y** field returns whether the Y axis is active.

## Z
### Summary
Whether the Z axis is active.

### Description
The **Z** field returns whether the Z axis is active.

# Operators
## Eq
### Summary
Returns whether two Axes values are equal.

### Description
The **equal** operator returns true if both operands are Axes, and each
corresponding component is equal.
