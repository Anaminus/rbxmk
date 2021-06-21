# Summary
Encodes comma-separated values.

# Description
The **csv** format decodes comma-separated values into a two-dimensional array.

Direction | Type  | Description
----------|-------|------------
Decode    | Array | An array of arrays of strings.
Encode    | Array | An array of arrays of strings.

CSV data decodes into a two-dimensional array of strings. For example,

	A,B,C
	D,E,F
	G,H,I

decodes into

	{
		{"A", "B", "C"),
		{"D", "E", "F"),
		{"G", "H", "I"),
	}

When encoding, each field must be string-like, but cannot be an Instance.

When decoding, each record must have the same number of fields. When encoding,
records do not need to have the same number of fields.

This format has no options.
