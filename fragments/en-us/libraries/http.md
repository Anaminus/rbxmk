# Summary
An interface to resources via HTTP.

# Description
The **http** library provides an interface to resources on the network via HTTP.

$FIELDS

# Fields
## request
### Summary
Begins an HTTP request.

### Description
The **request** function begins a request with the specified
[options][HTTPOptions]. Returns a [request object][HTTPRequest] that may be
[resolved][HTTPRequest.Resolve] or [canceled][HTTPRequest.Cancel]. Throws an
error if the request could not be started.
