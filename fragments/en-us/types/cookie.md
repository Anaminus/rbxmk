# Summary
An HTTP cookie.

# Description
The **Cookie** type contains information about an HTTP cookie. It has the
following members:

$MEMBERS

For security reasons, the value of the cookie cannot be accessed.

Cookie is immutable. A Cookie can be created with the [Cookie.new][Cookie.new]
constructor. Additionally, Cookies can be fetched from known locations with the
[Cookie.from][Cookie.from] function.

# Constructors
## from
### Summary
Gets cookies from a known location.

### Description
The **from** constructor retrieves cookies from a known location. *location* is
case-insensitive.

The following locations are implemented:

Location | Description
---------|------------
`studio` | Returns the cookies used for authentication when logging into Roblox Studio.

Returns nil if no cookies could be retrieved from the location. Throws an error
if an unknown location is given.

## new
### Summary
Creates a new cookie.

### Description
The **new** constructor creates a new cookie object.

# Properties
## Name
### Summary
The name of the cookie.

### Description
The **Name** field is the name of the cookie.

# Operators
## Eq
### Summary
Returns whether two Cookie values are equal.

### Description
The **equal** operator returns true if both operands are Cookie values, and
their Name properties are equal, and their internal values are equal.
