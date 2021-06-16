# Summary
Represents a pending HTTP request.

# Description
The **HTTPRequest** type represents a pending HTTP request. It has the following
members:

$MEMBERS

# Methods
## Cancel
### Summary
Cancels the pending request.

### Description
The **Cancel** method cancels the pending request.

## Resolve
### Summary
Returns the response to the request.

### Description
The **Resolve** method blocks until the request has finished, and returns the
response. Throws an error if a problem occurred while resolving the request.
