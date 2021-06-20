# Summary
The Lua 5.1 standard library, abridged.

# Description
The **base** library is loaded directly into the global environment. It contains
the following items from the [Lua 5.1 standard
library](https://www.lua.org/manual/5.1/manual.html#5):

$FIELDS

# Fields

## \_G
### Summary
The global environment.

### Description
The **\_G** field holds the global environment (`_G._G == _G`). Setting \_G to a
different value does not affect the content of the environment.

## \_VERSION
### Summary
The current version of Lua.

### Description
The **\_VERSION** field is the current version of the Lua interpreter. The
current value is "Lua 5.1".

## assert
### Summary
Tests a condition.

### Description
The **assert** function throws an error if *v* evaluates to false. Otherwise,
each argument (including *v* and *message*) is returned. *message* is used as
the error if the assertion fails.

## error
### Summary
Throws an error.

### Description
The **error** function throws an error by terminating the last protected
function that was called, using *message* as the error.

*level* specifies the stack level from which error position information will be
acquired. A *level* of 0 causes no position information to be included.

## getmetatable
### Summary
Gets the metatable of a value.

### Description
The **getmetatable** function gets the metatable of *v*. Returns nil if *v* does
not have a metatable. Otherwise, if the metatable has the field "\__metatable",
then its value is returned. Otherwise, the metatable itself is returned.

## ipairs
### Summary
Returns an iterator over an array.

### Description
The **ipairs** function iterates over the sequential integers keys of *t*. The
values returned are suitable to be used directly with a generic for loop:

```lua
for index, value in ipairs(t) do ... end
```

## next
### Summary
Gets the next field in a table.

### Description
The **next** function allows all fields of a table to be traversed. *t* is the
table, and *index* is the index in the table. next returns the next index and
its associated value.

If *index* is nil, then the first index and its value are returned. When called
with the last index, nil is returned.

The order in which fields are returned is undefined. This order is not changed
after modifying existing fields, removing fields, or making no changes. How this
order changes, when a value is assigned to a non-extant field, is undefined.

## pairs
### Summary
Returns an iterator over a table.

### Description
The **pairs** function iterates over all fields of *t*. The values returned are
suitable to be used directly with a generic for loop:

```lua
for key, value in pairs(t) do ... end
```

Specifically, *next* is the next function, *t* is the received *t*, and *start*
is nil.

## pcall
### Summary
Calls a function in protected mode.

### Description
The **pcall** function calls *f* in protected mode. Any error thrown within *f*
is caught and returned, instead of propagated. The remaining arguments are
passed to *f*.

*ok* will be true if the call succeeded without errors. In this case, the
remaining values are those returned by *f*. If *ok* is false, then the next
value is the error value.

## print
### Summary
Prints values to standard output.

### Description
The **print** function receives each argument, converts them to strings, and
writes them to standard output.

## select
### Index
#### Summary
Returns arguments after a given index.

#### Description
The **select** function returns all arguments after the argument indicated by
*index*.

```lua
select(1, "A", "B", "C") --> "A", "B", "C"
select(2, "A", "B", "C") --> "B", "C"
select(3, "A", "B", "C") --> "C"
select(4, "A", "B", "C") -->
```

If *index* is less than 0, then indexing starts from the end of the arguments.
An error is thrown if *index* is 0.

### Count
#### Summary
Returns the number of arguments.

#### Description
The **select** function with a *count* of "#", returns the number of additional
arguments passed.

## setmetatable
### Summary
Sets the metatable of a value.

### Description
The **setmetatable** function sets the metatable of *v* to *metatable*. If
*metatable* is nil, then the metatable is removed. An error is thrown if the
current metatable has a "\__metatable" field.

## tonumber
#### Summary
Converts a value to a number.

#### Description
The **tonumber** function attempts to convert *x* into a number. Returns a
number if *x* is a number or a string that can be parsed as a number. Returns
nil otherwise.

*base* specifies the base to interpret the numeral. It may be any integer
between 2 and 36, inclusive. In bases above 10, the letters A to Z
(case-insensitive) represent 10 to 35, respectively. In base 10, the number can
have a decimal part, as well as an optional exponent part. In other bases, only
unsigned integers are accepted.

## tostring
### Summary
Converts a value to a string.

### Description
The **tostring** function converts *v* into a string in a reasonable format. If
*v* has a metatable with the "\__tostring" field, then this field is called with
*v* as its argument, and its result is returned.

## type
### Summary
Returns the Lua type of a value.

### Description
The **type** function returns the Lua type of *v* as a string. Possible values
are "nil", "boolean", "number", "string", "table", "function, "thread", and
"userdata".

To get the internal type of userdata values, use the typeof function from the
roblox library.

## unpack
### Summary
Returns the elements in an array.

### Description
The **unpack** function returns each element in *list* from *i* up to and
including *j*. By default, *i* is 1, and *j* is the length of *list*.

## xpcall
### Summary
Calls a function in protected mode with an error handler.

### Description
The **xpcall** function is like pcall, except that it allows a custom error
handler to be specified. *msgh* is the error handler, and is called within the
erroring stack frame. *msgh* receives the current error message, and returns the
new error message to be returned by xpcall. The remaining arguments to xpcall
are passed to *f*.
