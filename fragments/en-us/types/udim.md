# Summary
A dimension with a dynamic and constant component.

# Description
The **UDim** type represents one dimension with a dynamic and constant
component.

# Constructors
## new
### Summary
Returns a new UDim.

### Description
The **new** constructor returns a new UDim. *scale* sets the Scale component,
and *offset* sets the Offset component.

# Properties
## Offset
### Summary
The constant component.

### Description
The **Offset** field returns the constant component of the UDim.

## Scale
### Summary
The dynamic component.

### Description
The **Scale** field returns the dynamic component of the UDim.

# Operators
## Add
### Summary
The sum of two UDim values.

### Description
The **add** operator returns a UDim where each corresponding component of the
two operands are summed.

## Sub
### Summary
The difference between two UDim values.

### Description
The **sub** operator returns a UDim where each corresponding component of the
two operands are subtracted.

## Eq
### Summary
Returns whether two UDim values are equal.

### Description
The **equal** operator returns true if both operands are UDim, and each
corresponding component is equal.

## Unm
### Summary
The negation of the UDim.

### Description
The **unm** operator returns a UDim where each component is negated.
