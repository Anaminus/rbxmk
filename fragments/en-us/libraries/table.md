# Summary
Extensions to the standard table library.

# Description
The **table** library is an extension to the standard library that includes the
same additions to [Roblox's table
library](https://developer.roblox.com/en-us/api-reference/lua-docs/table):

$FIELDS

# Fields
## clear
### Summary
Removes all entries from a table.

### Description
The **clear** function removes all the entries from *t*.

## create
### Summary
Creates a new table with a preallocated capacity.

### Description
The **create** function returns a table with the array part allocated with a
capacity of *cap*. Each entry in the array is optionally filled with *value*.

## find
### Summary
Find the index of a value in a table.

### Description
The **find** function returns the index in *t* of the first occurrence of
*value*, or nil if *value* was not found. Starts at index *init*, or 1 if
unspecified.

## move
### Summary
Copies the entries in a table.

### Description
The **move** function copies elements from *a1* to *a2*, performing the
equivalent to the multiple assignment

	a2[t], ... = a1[f], ..., a1[e]

The default for *a2* is *a1*. The destination range can overlap the source
range. Returns *a2*.

## pack
### Summary
Packs arguments into a table.

### Description
The **pack** function returns a table with each argument stored at keys 1, 2,
etc. Also sets field "n" to the number of arguments. Note that the resulting
table may not be a sequence.

## unpack
### Summary
Unpacks a table into arguments.

### Description
Returns the elements from *list*, equivalent to

	list[i], list[i+1], ..., list[j]

By default, *i* is 1 and *j* is the length of *list*.
