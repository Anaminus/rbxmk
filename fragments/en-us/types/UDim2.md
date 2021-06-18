# Summary
A UDim on two dimensions.

# Description
The **UDim2** type represents a UDim on two dimensions.

# Constructors
## fromOffset
### Summary
Returns a UDim2 from offset components.

### Description
The **fromOffset** constructor returns a UDim2 that sets the offset components
of each dimension.

## fromScale
### Summary
Returns a UDim2 from scale components.

### Description
The **fromScale** constructor returns a UDim2 that sets the scale components of
each dimension.

## new
### Components
#### Summary
Returns a UDim2 from components.

#### Description
The **new** constructor returns a UDim2 composed from the components of each
UDim.

### UDim
#### Summary
Returns a UDim2 from UDims.

#### Description
The **new** constructor returns a UDim2 composed from two UDim values.

# Properties
## Height
### Summary
The height of the UDim2.

### Description
The **Height** field returns the Y dimension of the UDim2.

## Width
### Summary
The width of the UDim2.

### Description
The **Width** field returns the X dimension of the UDim2.

## X
### Summary
The X dimension.

### Description
The **X** field returns the X dimension of the UDim2.

## Y
### Summary
The Y dimension.

### Description
The **Y** field returns the Y dimension of the UDim2.

# Methods
## Lerp
### Summary
Lineearly interpolates between two UDim2 values.

### Description
The **Lerp** method returns a UDim2 linearly interpolated from the UDim2 to
*goal* according to *alpha*, which has an interval of [0, 1].

# Operators
## Add
### Summary
The sum of two UDim2 values.

### Description
The **add** operator returns a UDim2 where each corresponding component of the
two operands are summed.

## Sub
### Summary
The difference between two UDim values.

### Description
The **sub** operator returns a UDim2 where each corresponding component of the
two operands are subtracted.

## Eq
### Summary
Returns whether two UDim2 values are equal.

### Description
The **equal** operator returns true if both operands are UDim2, and each
corresponding component is equal.

## Unm
### Summary
The negation of the UDim2.

### Description
The **unm** operator returns a UDim2 where each component is negated.
