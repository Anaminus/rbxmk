# Summary
A position and rotation.

# Description
The **CFrame** type is a matrix representing a combined position and
orientation.

# Constructors
## Angles
### Summary
Returns a CFrame composed from angles.

### Description
The **Angles** constructor returns a CFrame located at the origin, oriented
according to the angles (*rx*, *ry*, *rz*), in radians. Rotations are ordered Z,
Y, X.

## fromAxisAngle
### Summary
Returns a CFrame composed from an axis and angle.

### Description
The **fromAxisAngle** constructor returns a CFrame located at the origin,
rotated *rotation* radians around *axis*.

## fromEulerAnglesXYZ
### Summary
Returns a CFrame composed from Euler angles.

### Description
The **fromEulerAnglesXYZ** constructor returns a CFrame located at the origin,
oriented according to the angles (*rx*, *ry*, *rz*), in radians. Rotations are
ordered Z, Y, X.

## fromEulerAnglesYXZ
### Summary
Returns a CFrame composed from Euler angles ordered Y, X, Z.

### Description
The **fromEulerAnglesYXZ** constructor returns a CFrame located at the origin,
oriented according to the angles (*rx*, *ry*, *rz*), in radians. Rotations are
ordered Y, X, Z.

## fromMatrix
### Summary
Returns a CFrame from a matrix.

### Description
The **fromMatrix** constructor returns a CFrame located at *position*, rotated
according to the following rotation matrix:

```
[vx.X, vy.X, vz.X]
[vx.Y, vy.Y, vz.Y]
[vx.Z, vy.Z, vz.Z]
```

If *vz* is the zero vector, then *vz* is calculated as the unit of the cross
product of vx and vy.

## fromOrientation
### Summary
Returns a CFrame composed from orientations.

### Description
The **fromOrientation** constructor returns a CFrame located at the origin,
oriented according to the angles (*rx*, *ry*, *rz*), in radians. Rotations are
ordered Y, X, Z.

## lookAt
### Summary
Returns a CFrame that points towards a location.

### Description
The **lookAt** constructor returns a CFrame located at *position*, facing
towards *lookAt*, with the local upward direction determined by *up*.

## new
### Identity
#### Summary
Returns the default CFrame.

#### Description
The **new** constructor returns a CFrame with the default orientation located at
the origin.

### Position
#### Summary
Returns a CFrame with a position.

#### Description
The **new** constructor returns a CFrame with the default orientation located at
*position*.

### LookAt
#### Summary
Returns a CFrame pointing towards a location while tilted upward.

#### Description
The **new** constructor returns a CFrame located at *position*, facing towards
*lookAt*. The local upward direction is (0, 1, 0).

### Position components
#### Summary
Returns a CFrame with a position from coordinates.

#### Description
The **new** constructor returns a CFrame located at (*x*, *y*, *z*), with the
default orientation.

### Quaternion
#### Summary
Returns a CFrame from a quaternion.

#### Description
The **new** constructor returns a CFrame located at (*x*, *y*, *z*), oriented
according to the quaternion (*qx*, *qy*, *qz*, *qw*).

### Components
#### Summary
Returns a CFrame from raw components.

#### Description
The **new** constructor returns a CFrame located at (*x*, *y*, *z*), oriented
according to the following rotation matrix:

```
[r00, r01, r02]
[r10, r11, r12]
[r20, r21, r22]
```

# Properties
## LookVector
### Summary
The forward direction.

### Description
The **LookVector** field returns the forward-direction, or the negation of the
third column of the rotation matrix.

## P
### Summary
The position of the CFrame.

### Description
The **P** field returns the position of the CFrame.

## Position
### Summary
The position of the CFrame.

### Description
The **Position** field returns the position of the CFrame.

## RightVector
### Summary
The right direction.

### Description
The **RightVector** field returns the right-direction, or first column of the
rotation matrix.

## UpVector
### Summary
The up direction.

### Description
The **UpVector** field returns the up-direction, or second column of the rotation
matrix.

## X
### Summary
The X component of the Position.

### Description
The **X** field returns the X component of the Position.

## XVector
### Summary
The first row of the rotation matrix.

### Description
The **XVector** field returns the first row of the rotation matrix.

## Y
### Summary
The Y component of the Position.

### Description
The **Y** field returns the Y component of the Position.

## YVector
### Summary
The second row of the rotation matrix.

### Description
The **YVector** field returns the second row of the rotation matrix.

## Z
### Summary
The Z component of the Position.

### Description
The **Z** field returns the Z component of the Position.

## ZVector
### Summary
The third row of the rotation matrix.

### Description
The **ZVector** field returns the third row of the rotation matrix.

# Methods
## GetComponents
### Summary
Returns the raw components.

### Description
The **GetComponents** method returns the components of the CFrame's position and
rotation matrix.

## Inverse
### Summary
Returns the inverse CFrame.

### Description
The **Inverse** method returns the inverse of the CFrame.

## Lerp
### Summary
Linearly interpolates between two CFrames.

### Description
The **Lerp** method returns a CFrame linearly interpolated from the CFrame to
*goal* according to *alpha*, which has an interval of [0, 1].

## PointToObjectSpace
### Summary
Transforms a Vector3 to local space.

### Description
The **PointToObjectSpace** method returns a Vector3 transformed from world to
local space of the CFrame.

## PointToWorldSpace
### Summary
Transforms a Vector3 to world space.

### Description
The **PointToWorldSpace** method returns a Vector3 transformed from local to
world space of the CFrame.

## ToAxisAngle
### Summary
Returns the orientation as an axis and angle.

### Description
The **ToAxisAngle** method returns the orientation of the CFrame as an angle, in
radians, rotated around an axis.

## ToEulerAnglesXYZ
### Summary
Returns the orientation as Euler angles.

### Description
The **ToEulerAnglesXYZ** method returns the approximate angles of the CFrame's
orientation, in radians, if ordered Z, Y, X.

## ToEulerAnglesYXZ
### Summary
Returns the orientation as Euler angles order Y, X, Z.

### Description
The **ToEulerAnglesYXZ** method returns the approximate angles of the CFrame's
orientation, in radians, if ordered Y, X, Z.

## ToObjectSpace
### Summary
Transforms to local space.

### Description
The **ToObjectSpace** method returns the CFrame transformed from world to local
space of *cf*.

## ToOrientation
### Summary
Returns the orientation.

### Description
The **ToOrientation** method returns the approximate angles of the CFrame's
orientation, in radians, if ordered Y, X, Z.

## ToWorldSpace
### Summary
Transforms to world space.

### Description
The **ToWorldSpace** method returns the CFrame transformed from local to world
space of *cf*.

## VectorToObjectSpace
### Summary
Returns a Vector3 rotated to local space.

### Description
The **VectorToObjectSpace** method returns a Vector3 rotated from world to local
space of the CFrame.

## VectorToWorldSpace
### Summary
Returns a Vector3 rotated to world space.

### Description
The **VectorToWorldSpace** method returns a Vector3 rotated from local to world
space of the CFrame.

# Operators
## Add
### Summary
Translates by adding.

### Description
The **add** operator returns the CFrame translated in world space by the
operand.

## Sub
### Summary
Translates by subtracting.

### Description
The **sub** operator returns the CFrame translated in world space by the
negation of the operand.

## Mul
### CFrame
#### Summary
Composes two CFrames.

#### Description
The **mul** operator returns the composition of two CFrames.

### Vector3
#### Summary
Transforms to world space.

#### Description
The **mul** operator returns the operand transformed from local to world space
of the CFrame.

## Eq
### Summary
Returns whether two CFrame values are equal.

### Description
The **equal** operator returns true if both operands are CFrame, and each
corresponding component is equal.
