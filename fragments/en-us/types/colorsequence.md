# Summary
An interpolated sequence of colors.

# Description
The **ColorSequence** type represents an interpolated sequence of colors.

# Constructors
## new
### Single
Returns a sequence with a single color.

#### Summary
The **new** constructor returns a ColorSequence with a single color.

#### Description
### Range
#### Summary
Returns a sequence with two colors.

#### Description
The **new** constructor returns a ColorSequence that interpolates between
*color0* and *color1*.

### Keypoints
#### Summary
Returns a sequence composed of the given keypoints.

#### Description
The **new** constructor returns a ColorSequence composed from *keypoints*. Each
keypoint must be ordered ascending by their Time field. The first keypoint must
have a Time of 0, and the last keypoint must have a Time of 1.

# Properties
## Keypoints
### Summary
The keypoints of the sequence.

### Description
The **Keypoints** field returns the ColorSequenceKeypoints of the sequence.

# Operators
## Eq
### Summary
Returns whether two ColorSequence values are equal.

### Description
The **equal** operator returns true if both operands are ColorSequence, and each
corresponding keypoint is equivalent.
