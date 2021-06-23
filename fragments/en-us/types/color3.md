# Summary
A color in RGB space.

# Description
The **Color3** type represents a color in RGB space.

# Constructors
## fromHSV
### Summary
Returns a Color3 from hue, saturation, and value.

### Description
The **fromHSV** constructor returns a Color3 from the given hue, saturation, and
value.

## fromRGB
### Summary
Returns a Color3 from 8-bit components.

### Description
The **fromRGB** constructor returns a Color3 from the given red, green, and blue
components, each having an interval of [0, 255].

## new
### Zero
#### Summary
Returns the zero Color3.

#### Description
The **new** constructor returns the zero value for Color3, with each component
being zero, resulting in the color black.

### Components
#### Summary
Returns a Color3 from components.

#### Description
The **new** constructor returns a Color3 from the given red, green, and blue
components, each having an interval of [0, 1].

# Properties
## B
### Summary
The blue component.

### Description
The **B** field returns the blue component of the color.

## G
### Summary
The green component.

### Description
The **G** field returns the green component of the color.

## R
### Summary
The red component.

### Description
The **R** field returns the red component of the color.

# Methods
## Lerp
### Summary
Linearly interpolates between two colors.

### Description
The **Lerp** method returns a Color3 linearly interpolated from the color to
*goal* according to *alpha*, which has an interval of [0, 1].

## ToHSV
### Summary
Converts to hue, saturation, and value.

### Description
The **ToHSV** method returns the hue, saturation, and value approximated from
the color.

# Operators
## Eq
### Summary
Returns whether two Color3 values are equal.

### Description
The **equal** operator returns true if both operands are Color3, and each
corresponding component is equal.
