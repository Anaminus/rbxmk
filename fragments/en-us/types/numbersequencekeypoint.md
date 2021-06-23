# Summary
A keypoint of a NumberSequence.

# Description
The **NumberSequenceKeypoint** type represents a single keypoint of a
NumberSequence.

# Constructors
## new
### Components
#### Summary
Creates a new keypoint.

#### Description
The **new** constructor returns a new NumberSequenceKeypoint, *time* sets the
Time field, and *value* sets the Value field.

### Envelope
#### Summary
Creates a new keypoint with an envelope.

#### Description
The **new** constructor returns a new NumberSequenceKeypoint, *time* sets the
Time field, *value* sets the Value field, and *envelope* sets the Envelope
field.

# Properties
## Envelope
### Summary
The amount of variance allowed from the Value.

### Description
The **Envelope** field returns the amount of variance allowed from the value.

## Time
### Summary
The time location along the sequence.

### Description
The **Time** field returns the temporal location of the keypoint in the
sequence. Has an interval of [0, 1].

## Value
### Summary
The base value.

### Description
The **Value** field returns the base value of the keypoint.

# Operators
## Eq
### Summary
Returns whether two NumberSequenceKeypoint values are equal.

### Description
The **equal** operator returns true if both operands are NumberSequenceKeypoint,
and each corresponding component is equal.
