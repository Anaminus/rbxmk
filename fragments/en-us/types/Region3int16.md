# Summary
An axis-aligned rectangular cuboid with 16-bit integer precision.

# Description
The **Region3** type represents an axis-aligned three-dimensional rectangular
cuboid with a lower and upper boundary, with each component having 16-bit
integer precision.

# Constructors
## new
### Summary
Returns Region3 composed by two Vector3s.

### Description
The **new** constructor returns a Region3 composed by two Vector3int16 values,
where *min* is the lower bound of the region, and *max* is the upper bound.

# Properties
## Max
### Summary
The upper bounds of the region.

### Description
The **Max** field returns the upper bounds of the region.

## Min
### Summary
The lower bounds of the region.

### Description
The **Min** field returns the lower bounds of the region.

# Operators
## Eq
### Summary
Returns whether two Region3int16 values are equal.

### Description
The **equal** operator returns true if both operands are Region3int16, and each
corresponding component is equal.
