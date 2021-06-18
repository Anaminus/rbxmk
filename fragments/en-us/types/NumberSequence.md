# Summary
An interpolated sequence of numbers.

# Description
The **NumberSequence** type represents an interpolated sequence of numbers.

# Constructors
## new
### Single
Returns a sequence with a single number.

#### Summary
The **new** constructor returns a NumberSequence with a single number.

#### Description
### Range
#### Summary
Returns a sequence with two numbers.

#### Description
The **new** constructor returns a NumberSequence that interpolates between
*value0* and *value1*.

### Keypoints
#### Summary
Returns a sequence composed of the given keypoints.

#### Description
The **new** constructor returns a NumberSequence composed from *keypoints*. Each
keypoint must be ordered ascending by their Time field. The first keypoint must
have a Time of 0, and the last keypoint must have a Time of 1.

# Properties
## Keypoints
### Summary
The keypoints of the sequence.

### Description
The **Keypoints** field returns the NumberSequenceKeypoints of the sequence.

# Operators
## Eq
### Summary
Returns whether two NumberSequence values are equal.

### Description
The **equal** operator returns true if both operands are NumberSequence, and
each corresponding keypoint is equivalent.
