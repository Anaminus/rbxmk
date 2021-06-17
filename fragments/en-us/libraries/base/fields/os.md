# Summary
Functions related to the operating system.

# Description
The **os** library contains functions related to the operating system.

# Fields
## clock
### Summary
Returns the CPU time.

### Description
The **clock** function returns, approximately, the seconds of CPU time used by
the program.

## date
### Current
#### Summary
Returns the current time.

#### Description
The **date** function returns a string-representation of the current time.

### Tabular
#### Summary
Returns the time as a table.

#### Description
The **date** function returns the time as a table with a number of fields for
each component:

Field | Notes
------|------
year  | 4 digits
month | 1 - 12
day   | 1 - 31
hour  | 0 - 23
min   | 0 - 59
sec   | 0 - 61
wday  | Day of the week, 1 - 7 starting on Sunday.
yday  | Day of the year, 1 - 366.
isdst | Whether Daylight Saving Time is active.

If *time* is specified, then it specifies the time to be formatted. Otherwise,
the current time is used. See os.time for a description of this value.

### Formatted
#### Summary
Returns the time as a formatted string.

#### Description
The **date** function returns the formatted according to *format*, which has the
same rules as the C function strftime.

If *time* is specified, then it specifies the time to be formatted. Otherwise,
the current time is used. See os.time for a description of this value.

## difftime
### Summary
Returns the difference between two times.

### Description
The **difftime** function returns the number of seconds between *t2* and *t1*.
On most systems, this value is exactly equal to `t2 - t1`.

## time
### Summary
Returns a numeric time.

### Description
The **time** function returns a numeric time corresponding to the fields of *t*.
If *t* is unspecified, then the current time is returned.

On most systems, the returned value is the relative to the UNIX epoch. On
others, the meaning is unspecified, and may only be passed to os.date and
os.difftime.
