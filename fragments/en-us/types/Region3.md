# Summary
An axis-aligned rectangular cuboid.

# Description
The **Region3** type represents an axis-aligned three-dimensional rectangular
cuboid with a lower and upper boundary.

# Constructors
## new
### Summary
Returns Region3 composed by two Vector3s.

### Description
The **new** constructor returns a Region3 composed by two Vector3 values, where
*min* is the lower bound of the region, and *max* is the upper bound.

# Properties
## CFrame
### Summary
The center of the region.

### Description
The **CFrame** field returns a CFrame with its Position located at the center of
the region.

## Size
### Summary
The size of the region.

### Description
The **Size** returns the size of the region.

# Methods
## ExpandToGrid
### Summary
Expands the region to align to a grid.

### Description
The **ExpandToGrid** method returns the region expanded so that is lower and
upper bounds align to *resolution*. If *resolution* is 0, the region is returned
unchanged.

# Operators
## Eq
### Summary
Returns whether two Region3 values are equal.

### Description
The **equal** operator returns true if both operands are Region3, and each
corresponding component is equal.
