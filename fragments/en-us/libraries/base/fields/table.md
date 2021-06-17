# Summary
Provides functions for manipulating tables.

# Description
The **table** library provides functions for manipulating tables.

Most functions assume that the table is an array, where "length" refers to the
result of the `#` operator.

# Fields
## concat
### Summary
Concatenates each element as a string.

### Description
The **concat** function concatenates each element within the table from *i* to
*j*. Each element must be a string or a number. *sep* is concatenated between
each element. An empty string is returned if *i* is greater than *j*.

## insert
### Insert
#### Summary
Inserts an element.

#### Description
The **insert** function inserts *value* into *t* at *index*. Each element after
*index* is shifted upward to make room.

### Append
#### Summary
Appends an element.

#### Description
The **insert** function appends *value* to the end of *t*, such that the index
is the length *t* plus one.

## maxn
### Summary
Returns the largest positive numerical index.

### Description
The **maxn** function returns the largest positive numerical index in *t*, or 0
if none can be found.

## remove
### Summary
Removes an element.

### Description
The **remove** function removes the element at *index* from *t*. Each element
after *index* is shifted downward to close the gap. Returns the removed element.

## sort
### Summary
Sorts elements.

### Description
The **sort** function sorts the elements in *t* in-place from index 1 to the
length of *t*. If *comp* is specified, then it must return true when *a* is less
than *b*. If *comp* is unspecified, then the `<` operator is used instead.

The sorting algorithm is not stable. Elements considered equal may have their
relative positions changed.
