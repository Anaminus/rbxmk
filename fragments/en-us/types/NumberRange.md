# Summary
A range of numbers.

# Description
The **NumberRange** type represents a range of numbers between a minimum and
maximum.

# Constructors
## new
### Single
#### Summary
Returns a range with a single value.

#### Description
The **new** constructor returns a NumberRange where *value* is both the minimum
and maximum.

### Range
#### Summary
Returns a range with between two values.

#### Description
The **new** constructor returns a NumberRange where *minimum* is the minimum
value and *maximum* is the maximum value. Throws an error if *minimum* is
greater than *maximum*.

# Properties
## Max
### Summary
The maximum value.

### Description
The **Max** field returns the maximum possible value in the range.

## Min
### Summary
The minimum value.

### Description
The **Min** field returns the minimum possible value in the range.

# Operators
## Eq
### Summary
Returns whether two NumberRange values are equal.

### Description
The **equal** operator returns true if both operands are NumberRange, and each
corresponding component is equal.
