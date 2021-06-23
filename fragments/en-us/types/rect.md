# Summary
An axis-aligned rectangle.

# Description
The **Rect** type represents an axis-aligned two-dimensional rectangle with a
lower and upper boundary.

# Constructors
## new
### Vector2
#### Summary
Returns a Rect composed by two Vector2s.

#### Description
The **new** constructor returns a Rect composed by two Vector2s. *min* is the
lower bounds, and *max* is the upper bounds.

### Components
#### Summary
Returns a Rect composed by boundary components. *minX* and *minY* compose the
lower bounds, and *maxX* and *maxY* compose the upper bounds.

#### Description
The **new** constructor returns a Rect composed by boundary components.

# Properties
## Height
### Summary
The height of the rectangle.

### Description
The **Height** field returns the height of the rectangle, or the difference
between the upper and lower bounds on the Y axis.

## Max
### Summary
The upper bounds of the rectangle.

### Description
The **Max** field returns the upper bounds of the rectangle.

## Min
### Summary
The lower bounds of the rectangle.

### Description
The **Min** field returns the lower bounds of the rectangle.

## Width
### Summary
The width of the rectangle.

### Description
The **Width** field returns the width of the rectangle, or the difference
between the upper and lower bounds on the X axis.

# Operators
## Eq
### Summary
Returns whether two Rect values are equal.

### Description
The **equal** operator returns true if both operands are Rect, and each
corresponding component is equal.
