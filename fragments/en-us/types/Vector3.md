# Summary
A three-dimensional Euclidean vector.

# Description
The **Vector3** type represents a three-dimensional Euclidean vector.

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
## Magnitude
### Summary
The length of the vector.

### Description
The **Magnitude** field returns the length of the vector.

## Unit
### Summary
The direction of the vector.

### Description
The **Unit** field returns a vector with the same direction, but a length of 1.

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

## Z
### Summary
The Z coordinate.

### Description
The **Z** field returns the Z coordinate of the vector.

# Methods
## Cross
### Summary
Returns the cross product of two vectors.

### Description
The **Cross** method returns the cross product of the vector and *op*.

## Dot
### Summary
Returns the dot product of two vectors.

### Description
The **Dot** method returns the dot product of the vector and *op*.

## FuzzyEq
### Summary
Returns whether two vectors are approximately equal.

### Description
The **FuzzyEq** method returns whether the distance between the vector and *op*
is less than or equal to *epsilon*.

## Lerp
### Summary
Linearly interpolates between two vectors.

### Description
The **Lerp** method returns a vector linearly interpolated from the vector to
*goal* according to *alpha*, which has an interval of [0, 1].

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
### Vector3
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
### Vector3
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
Returns whether two Vector3 values are equal.

### Description
The **equal** operator returns true if both operands are Vector3, and each
corresponding component is equal.

## Unm
### Summary
The negation of the vector.

### Description
The **unm** operator returns a vector where each component is negated.
