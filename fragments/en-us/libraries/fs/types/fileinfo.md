# Summary
The metadata of a file.

# Description
The **FileInfo** type is a table that contains metadata about a file. It has the
following fields:

Field   | Type    | Description
--------|---------|------------
Name    | string  | The base name of the file.
IsDir   | boolean | Whether the file is a directory.
Size    | number  | The size of the file, in bytes.
ModTime | number  | The modification time of the file, in Unix time.
