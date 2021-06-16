# Summary
A color from a predefined collection.

# Description
The **BrickColor** type represents a color from a predefined collection.

# Constructors
## Black
### Summary
Returns a black color.

### Description
The **Black** constructor returns the "Black" brick color.

## Blue
### Summary
Returns a blue color.

### Description
The **Blue** constructor returns the "Bright blue" brick color.

## DarkGray
### Summary
Returns a dark gray color.

### Description
The **DarkGray** constructor returns the "Dark stone grey" brick color.

## Gray
### Summary
Returns a gray color.

### Description
The **Gray** constructor returns the "Medium stone grey" brick color.

## Green
### Summary
Returns a green color.

### Description
The **Green** constructor returns the "Dark green" brick color.

## Red
### Summary
Returns a red color.

### Description
The **Red** constructor returns the "Bright red" brick color.

## White
### Summary
Returns a white color.

### Description
The **White** constructor returns the "White" brick color.

## Yellow
### Summary
Returns a yellow color.

### Description
The **Yellow** constructor returns the "Bright yellow" brick color.

## new
### Number
#### Summary
Creates a BrickColor from a numeric value.

#### Description
The **new** constructor returns the BrickColor for which the Number field is equal
to *value*.

### Components
#### Summary
Creates a BrickColor from color components.

#### Description
The **new** constructor returns the BrickColor that is nearest to the given
color components.

### Name
#### Summary
Creates a BrickColor from a name.

#### Description
The **new** constructor returns the BrickColor for which the Name field is equal
to *name*.

### Color
#### Summary
Creates a BrickColor from a Color3.

#### Description
The **new** constructor returns the BrickColor that is nearest to *color*.

## palette
### Summary
Returns a BrickColor from a predefined list of colors.

### Description
The **palette** constructors returns the BrickColor corresponding to *index* in
a predefined list of 128 colors. An error is thrown if *index* is not between 0
and 127.

## random
### Summary
Returns a random BrickColor.

### Description
The **random** constructor returns a randomly selected BrickColor.

# Properties
## B
### Summary
The blue component.

### Description
The **B** field is the blue component of the BrickColor.

## Color
### Summary
The Color3 corresponding to the BrickColor.

### Description
The **Color** field is the Color3 value corresponding to the BrickColor.

## G
### Summary
The green component.

### Description
The **G** field is the green component of the BrickColor.

## Name
### Summary
The name of the BrickColor.

### Description
The **Name** field is a name that describes the BrickColor.

## Number
### Summary
The numeric value of the BrickColor.

### Description
The **Number** field is a numeric value that uniquely identifies the BrickColor.

## R
### Summary
The red component.

### Description
The **R** field is the red component of the BrickColor.

# Operators
## Eq
### Summary
Returns whether two BrickColor values are equal.

### Description
The **equal** operator returns true if both operands are BrickColor, and their
Number fields are equal.
