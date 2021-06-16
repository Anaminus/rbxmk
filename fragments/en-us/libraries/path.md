# Summary
Handles file paths.

# Description
The **path** library provides functions that handle file paths. The following
functions are included:

$FIELDS

# Fields
## clean
### Summary
Cleans up a file path.

### Description
The **clean** function returns the shortest path equivalent to *path*.
Separators are replaced with the OS-specific path separator.

## expand
### Summary
Expands predefined file path variables.

### Description
The **expand** function scans *path* for certain variables of the form `$var` or
`${var}` an expands them. The following variables are expanded:

Variable                                             | Description
-----------------------------------------------------|------------
`$script_name`, `$sn`                                | The base name of the currently running script. Empty for stdin.
`$script_directory`, `$script_dir`, `$sd`            | The directory of the currently running script. Empty for stdin.
`$root_script_directory`, `$root_script_dir`, `$rsd` | The directory of the first running script. Empty for stdin.
`$working_directory`, `$working_dir`, `$wd`          | The current working directory.
`$temp_directory`, `$temp_dir`, `$tmp`               | The directory for temporary files.

## join
### Summary
Joins a number of file paths together.

### Description
The **join** function joins each *path* element into a single path, separating
them using the operating system's path separator. This also cleans up the path.

## split
### Summary
Splits a file path into its components.

### Description
The **split** function returns the components of a file path.

Component | `project/scripts/main.script.lua` | Description
----------|-----------------------------------|------------
`base`    | `main.script.lua`                 | The file name; the last element of the path.
`dir`     | `project/scripts`                 | The directory; all but the last element of the path.
`ext`     | `.lua`                            | The extension; the suffix starting at the last dot of the last element of the path.
`fext`    | `.script.lua`                     | The format extension, as determined by registered formats.
`fstem`   | `main`                            | The base without the format extension.
`stem`    | `main.script`                     | The base without the extension.

A format extension depends on the available formats. See [Formats](formats.md)
for more information.
