# Summary
A line extending infinitely in one direction.

# Description
The **Ray** type represents a line that extends infinitely in one direction.

# Constructors
## new
### Summary
Creates a new Ray.

### Description
The **new** constructor returns a Ray, where *origin* sets the Origin, and
*direction* sets the Direction.

# Properties
## Direction
### Summary
The direction of the ray.

### Description
The **Direction** field returns the direction of the ray.

## Origin
### Summary
The position of the ray.

### Description
The **Origin** field returns the position of the ray.

# Methods
## ClosestPoint
### Summary
Returns the nearest point on the ray.

### Description
The **ClosestPoint** method returns the position on the ray that is nearest to
*point*.

## Distance
### Summary
Returns distance to the nearest point on the ray.

### Description
The **Distance** method returns the distance between *point* and the point on
the ray nearest to *point*.

# Operators
## Eq
### Summary
Returns whether two Ray values are equal.

### Description
The **equal** operator returns true if both operands are Ray, and each
corresponding component is equal.
