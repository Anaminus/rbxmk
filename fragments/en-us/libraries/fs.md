# Summary
An interface to the file system.

# Description
The **fs** library provides an interface to the file system.

$FIELDS

# Fields
## dir
### Summary
Gets a list of files in a directory.

### Description
The **dir** function returns a list of files in the given directory. Each file
is a table with the following fields:

Field   | Type    | Description
--------|---------|------------
Name    | string  | The base name of the file.
IsDir   | boolean | Whether the file is a directory.

dir returns nil if the directory does not exist. An error is thrown if a problem
otherwise occurred while reading the directory.

## mkdir
### Summary
Makes a new directory.

### Description
The **mkdir** function creates a directory at *path*. If *all* is true, then
mkdir will create each parent directory as needed. *all* defaults to false.

Returns true if all the directories were created successfully. Returns false if
all of the directories already exist. Throws an error if a problem otherwise
occurred while creating a directory.

## read
### Summary
Reads data from a file in a certain format.

### Description
The **read** function reads the content of the file at *path*, and decodes it
into *value* according to the [format](formats.md) matching the file extension
of *path*. If *format* is given, then it will be used instead of the file
extension.

If the format returns an [Instance][Instance], then the Name property will be
set to the "fstem" component of *path* according to
[path.split](libraries.md#user-content-pathsplit).

## remove
### Summary
Removes a file or directory.

### Description
The **remove** function removes the file or directory at *path*. If *all* is
true, then removing a directory will also recursively remove all of its
children. *all* defaults to false.

Returns true if every file is removed successfully. Returns false if the file or
directory does not exist. Throws an error if a problem occurred while removing a
file.

## rename
### Summary
Moves a file or directory.

### Description
The **rename** functions moves the file or directory at path *old* to path
*new*. If *new* exists and is not a directory, it is replaced.

Returns true if the file was moved successfully. Returns false if the file or
directory does not exist. Throws an error if a problem otherwise occurred while
moving the file.

## stat
### Summary
Gets metadata about a file.

### Description
The **stat** function gets metadata of the given file. Returns a table with the
following fields:

Field   | Type    | Description
--------|---------|------------
Name    | string  | The base name of the file.
IsDir   | boolean | Whether the file is a directory.
Size    | number  | The size of the file, in bytes.
ModTime | number  | The modification time of the file, in Unix time.

stats returns nil if the file does not exist. An error will be thrown if a
problem otherwise occurred while getting the metadata.

stat does not follow symbolic links.

## write
### Summary
Writes data to a file in a certain format.

### Description
The **write** function encodes *value* according to the [format](formats.md)
matching the file extension of *path*, and writes the result to the file at
*path*. If *format* is given, then it will be used instead of the file
extension.
