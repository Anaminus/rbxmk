# Summary
Specifies how an HTTP request is made.

# Description
The **HTTPOptions** type is a table that specifies how an HTTP request is made.
It has the following fields:

Field          | Type                              | Description
---------------|-----------------------------------|------------
URL            | [string](##)                      | The URL to make to request to.
Method         | [string](##)?                     | The HTTP method. Defaults to GET.
RequestFormat  | [FormatSelector][FormatSelector]? | The format used to encode the request body.
ResponseFormat | [FormatSelector][FormatSelector]? | The format used to decode the response body.
Headers        | [HTTPHeaders][HTTPHeaders]?       | The HTTP headers to include with the request.
Cookies        | [Cookies][Cookies]?               | Cookies to append to the Cookie header.
Body           | [any](##)?                        | The body of the request, to be encoded by the specified format.

If RequestFormat is unspecified, then no request body is sent. If ResponseFormat
is unspecified, then no response body is returned.

Use of the Cookies field ensures that cookies sent with the request are
well-formed, and is preferred over setting the Cookie header directly.
