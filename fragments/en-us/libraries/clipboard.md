# Summary
An interface to the clipboard.

# Description
The **clipboard** library provides an interface to the operating system's
clipboard.

$FIELDS

**The clipboard is currently available only on Windows. Other operating systems
return no data.**

# Fields
## read
### Summary
Gets data from the clipboard in one of a number of formats.

### Description
The **read** function gets a value from the clipboard according to one of the
given [formats](formats.md).

Each format has a number of associated [media
types](https://en.wikipedia.org/wiki/Media_type). Each format is traversed in
order, and each media type within a format is traversed in order. The data that
matches the first media type found in the clipboard is selected. This data is
decoded by the format corresponding to the matched media type, and the result is
returned.

Throws an error if *value* could not be decoded from the format, or if data
could not be retrieved from the clipboard. If no data was found, then nil is
returned.

## write
### Summary
Sets data to the clipboard in a number of formats.

### Description
The **write** function sets *value* to the clipboard according to the given
[formats](formats.md).

Each format has a number of associated [media
types](https://en.wikipedia.org/wiki/Media_type). For each format, the data is
encoded in the format, which is then sent to the clipboard for each of the
format's media type. Data is not sent for a media type if that media type was
already sent.

If no formats are given, then the clipboard is cleared with no further action.

Throws an error if *value* could not be encoded in a format, or if data could
not be sent to the clipboard.
