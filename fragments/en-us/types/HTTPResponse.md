# Summary
Contains the response of an HTTP request.

# Description
The **HTTPResponse** type is a table that contains the response of a request. It
has the following fields:

Field         | Type                       | Description
--------------|----------------------------|------------
Success       | [bool](##)                 | Whether the request succeeded. True if StatusCode between 200 and 299.
StatusCode    | [int](##)                  | The HTTP status code of the response.
StatusMessage | [string](##)               | A readable message corresponding to the StatusCode.
Headers       | [HTTPHeaders][HTTPHeaders] | A set of response headers.
Cookies       | [Cookies][Cookies]         | Cookies parsed from the Set-Cookie header.
Body          | [any](##)?                 | The decoded body of the response.
