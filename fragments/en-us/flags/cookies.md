# cookies-from
Append cookies from a known location. See the documentation of
[Cookie.from](type:Cookie.from) for a list of locations. Can be given any number
of times.

# cookies-file
Append cookies from a file. The file is formatted as a number of
[Set-Cookie][Set-Cookie] headers. Can be given any number of times.

[Set-Cookie]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie

# cookie-var
Append a cookie from an environment variable. The content is formatted as a
number of Set-Cookie headers. Can be given any number of times.
