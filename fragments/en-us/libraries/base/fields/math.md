# Summary
Common mathematical functions.

# Description
The **math** library contains common mathematical functions and variables.

# Fields
## abs
### Summary
Returns the absolute value of *x*.

### Description
The **abs** function returns the absolute value of *x*.

## acos
### Summary
Returns the arc cosine of *x*.

### Description
The **acos** function returns the arc cosine of *x*, in radians.

## asin
### Summary
Returns the arc sine of *x*.

### Description
The **asin** function returns the arc sine of *x*, in radians.

## atan
### Summary
Returns the arc tangent of *x*.

### Description
The **atan** function returns the arc tangent of *x*, in radians.

## atan2
### Summary
Returns the arc tangent of *y*/*x*.

### Description
The *atan2* function returns the arc tangent of *y* divided by *x*, in radians.
The signs of both arguments are used to find the quadrant of the result. The
case of *x* being 0 is also handled.

## ceil
### Summary
Returns the smallest integer >= *x*.

### Description
The **ceil** function returns the smallest integer less than or equal to *x*.

## cos
### Summary
Returns the cosine of *x*.

### Description
The **cos** function returns the cosine of *x*, which is assumed to be in
radians.

## cosh
### Summary
Returns the hyperbolic cosine of *x*.

### Description
The **cosh** function returns the hyperbolic cosine of *x*.

## deg
### Summary
Returns angle *x* in degrees.

### Description
The **deg** function returns angle *x* (radians) in degrees.

## exp
### Summary
Returns e^*x*.

### Description
The **exp** function returns e^*x*, where e is Euler's number.

## floor
### Summary
Returns the largest integer <= *x*.

### Description
The **floor** function returns the largest integer less than or equal to *x*.

## fmod
### Summary
Returns the remainder of *x*/*y*, rounding towards 0.

### Description
The **fmod** function returns the remainder of *x* divided by *y* with the
quotient rounded towards 0.

## frexp
### Summary
Returns such that `x = m2^e`.

### Description
The **frexp** function returns *m* and *e* such that `x = m2^e`. *m* will be in
the range [0.5, 1), or 0 when *x* is 0.

## huge
### Summary
The value >= any other numerical value.

### Description
The **huge** value is greater than or equal to any other numerical value
(infinity).

## ldexp
### Summary
Returns `m2^e`.

### Description
The **ldexp** function returns `m2^e`.

## max
### Summary
Returns the maximum value.

### Description
The **max** function returns the largest value among the given arguments.

## min
### Summary
Returns the minimum value.

### Description
The **min** function returns the smallest value among the given arguments.

## modf
### Summary
Returns the integral and fractional part of *x*.

### Description
The **modf** function returns the integral and fractional part of *x*.

## pi
### Summary
The value of π.

### Description
The **pi** value is the value of the mathematical constant pi.

## pow
### Summary
Returns `x^y`.

### Description
The **pow** function returns `x^y`.

## rad
### Summary
Returns angle *x* in radians.

### Description
The **rad** function returns angle *x* (degrees) in radians.

## random
### Real
#### Summary
Returns a random number in the range [0, 1).

#### Description
The **random** function returns a uniform pseudo-random real number in the range
[0, 1).

### Range
#### Summary
Returns a random number in the range [1, *m*].

#### Description
The **random** function returns a uniform pseudo-random integer in the range [1,
*m*].

### Interval
#### Summary
Returns a random number in the range [*m*, *n*].

#### Description
The **random** function returns a uniform pseudo-random integer in the range
[*m*, *n*].

## randomseed
### Summary
Sets *x* as the seed for the random number generator.

### Description
The **randomseed** function sets *x* as the seed for the random number
generator.

## sin
### Summary
Returns the sine of *x*.

### Description
The **sin** function returns the sine of *x*, which is assumed to be in radians.

## sinh
### Summary
Returns the hyperbolic sine of *x*.

### Description
The **sinh** function returns the hyperbolic sine of *x*.

## sqrt
### Summary
Returns the square root of *x*.

### Description
The **sqrt** function returns the square root of *x*.

## tan
### Summary
Returns the tangent of *x*.

### Description
The **tan** function returns the tangent of *x*, which is assumed to be in
radians.

## tanh
### Summary
Returns the hyperbolic tangent of *x*.

### Description
The **tanh** function returns the hyperbolic tangent of *x*.
