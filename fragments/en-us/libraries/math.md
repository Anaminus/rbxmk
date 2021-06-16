# Summary
Extensions to the standard math library.

# Description
The **math** library is an extension to the standard library that includes the
same additions to [Roblox's math
library](https://developer.roblox.com/en-us/api-reference/lua-docs/math):

$FIELDS

# Fields
## clamp
### Summary
Returns a number clamped between a minimum and maximum.

### Description
The **clamp** function returns *x* clamped so that it is not less than *min* or
greater than *max*.

## log
### Summary
Includes optional base argument.

### Description
The **log** function returns the logarithm of *x* in *base*. The default for
*base* is `e`, returning the natural logarithm of *x*.

## round
### Summary
Rounds a number to the nearest integer.

### Description
The **round** function returns *x* rounded to the nearest integer. The function
rounds half away from zero.

## sign
### Summary
Returns the sign of a number.

### Description
The **sign** function returns the sign of *x*: `1` if *x* is greater than zero,
`-1` of *x* is less than zero, and `0` if *x* equals zero.
