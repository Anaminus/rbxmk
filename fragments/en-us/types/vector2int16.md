# Summary
A two-dimensional Euclidean vector with 16-bit integer precision.

# Description
The **Vector2int16** type represents a two-dimensional Euclidean vector with
16-bit integer precision.

# Constructors
## new
### Zero
#### Summary
Returns the zero vector.

#### Description
The **new** constructor returns the origin vector, where each coordinate is 0.

### Components
#### Summary
Returns a vector composed by each given coordinate.

#### Description
The **new** constructor returns a vector composed by each given component.

# Properties
## X
### Summary
The X coordinate.

### Description
The **X** field returns the X coordinate of the vector.

## Y
### Summary
The Y coordinate.

### Description
The **Y** field returns the Y coordinate of the vector.

# Operators
## Add
### Summary
The sum of two vectors.

### Description
The **add** operator returns a vector where each corresponding component of the
two operands are summed.

## Sub
### Summary
The difference between two vectors.

### Description
The **sub** operator returns a vector where each corresponding component of the
two operands are subtracted.

## Mul
### Vector2int16
#### Summary
Multiplies each corresponding component.

#### Description
The **mul** operator returns a vector where each corresponding component of the
two operands are multiplied.

### Number
#### Summary
Multiplies each component.

#### Description
The **mul** operator returns a vector where each component of the first operand
is multiplied by the second operand.

## Div
### Vector2int16
#### Summary
Divides each corresponding component.

#### Description
The **div** operator returns a vector where each corresponding component of the
two operands are divided.

### Number
#### Summary
Divides each component.

#### Description
The **div** operator returns a vector where each component of the first operand
is divided by the second operand.

## Eq
### Summary
Returns whether two Vector2int16 values are equal.

### Description
The **equal** operator returns true if both operands are Vector2int16, and each
corresponding component is equal.

## Unm
### Summary
The negation of the vector.

### Description
The **unm** operator returns a vector where each component is negated.
