# Summary
Provides functions for string manipulation.

# Description
The **string** library provides functions for string manipulation.

All strings have a metatable with the \__index field set to this library. This
allows functions to be called as methods on strings:

```lua
s:byte(...)
-- Equivalent to:
string.byte(s, ...)
```

# Fields
## byte
### Summary
Gets the bytes of a string subset.

### Description
The **byte** function returns the numerical byte codes of characters within *s*
between *i* and *j*.

## char
### Summary
Converts bytes to a string.

### Description
The **char** function converts each given numeric byte code to a character, then
returns the concatenation of each character.

## find
### Summary
Finds the indices of a pattern.

### Description
The **find** function searches for the first match of *pattern* in *s*. If
found, then find returns the indices of *s* where the match starts and ends. If
*pattern* contains captures, they are also returned after *start* and *end*. If
no match is found, then nil is returned.

*init* specifies where the search begins. If *plain* is true, then pattern
matching is disabled, and *pattern* is searched for as a simple substring.

## format
### Summary
Formats values as a string.

### Description
The **format** function formats each remaining argument according to *format*,
and returns the result.

*format* follows the same rules for printing as [Go's fmt
package](https://golang.org/pkg/fmt/#hdr-Printing), of which Lua's string.format
is mostly a subset.

## gmatch
### Summary
Iterates over the captures of a pattern.

### Description
The **gmatch** function returns an iterator function. Each time this function is
called, it returns the next captures from *pattern* for *s*. If *pattern*
contains no captures, then the entire match is returned.

## gsub
### Summary
Replaces patterns.

### Description
The **gsub** function returns *s* for which each occurrence of *pattern* is
replaced according to *repl*. If *n* is specified, then only the first *n*
occurrences are replaced.

If *repl* is a string, then it is used for replacement. A sequence within *repl*
of the form `%n`, where `n` is between 1 and 9, stands for the value of the
`n`-th captured substring. The sequence `%0` stands for the entire match, and
`%%` stands for a literal `%` character.

If *repl* is a table, then it is queried for each match, using the first capture
as the key. If *pattern* has no captures, then the entire match is used as the
key. The value is used for replacement.

If *repl* is a function, then it is called each time a match occurs. Each
capture is passed as an argument, or the entire match if no captures are
specified.

If a replacement value is a string or number, then it is used as the replacement
by converting to a string. If it is false or nil, then the match is not
replaced.

## len
### Summary
Returns the length of a string.

### Description
The **len** function returns the byte length of the string, equivalent to the
`#` operator.

## lower
### Summary
Converts uppercase letters to lowercase.

### Description
The **lower** function returns *s* with all uppercase letters converted to
lowercase, and any other character left unchanged.

## match
### Summary
Returns the captures from a pattern.

### Description
The **match** function searches for the first match of *pattern* in *s*. If
found, then each capture within *pattern* is returned. Otherwise, nil is
returned. If *pattern* has no captures, then the whole match is returned
instead. *init* specifies where to start the search.

## rep
### Summary
Repeats a string.

### Description
The **rep** function returns the concatenation of *n* copies of *s*.

## reverse
### Summary
Reverses a string.

### Description
The **reverse** function returns *s* with bytes in reversed order.

## sub
### Summary
Returns a subset of a string.

### Description
The **sub** function returns the substring of *s* starting from *i* up to and
including *j*.

## upper
### Summary
Converts lowercase letters to uppercase.

### Description
The **upper** function returns *s* with all lowercase letters converted to
uppercase, and any other character left unchanged.
