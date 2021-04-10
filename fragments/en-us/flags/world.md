# allow-insecure-paths
Disable path restrictions, allowing scripts to access any path in the file
system.

# debug
Display stack traces when an error occurs.

# libraries
A comma-separated list of libraries to include or exclude. A name prefixed with
`+` or nothing is included, and a name prefixed with `-` is excluded. If the
name is `*`, then the state is applied to all libraries. States are applied in
order.

# include-root
Mark a path as an accessible root directory. May be specified any number of
times.
