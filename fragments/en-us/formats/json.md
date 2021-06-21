# Summary
Encodes JSON data.

# Description
The **json** format is defined for encoding general data in JSON format.

Direction | Type       | Description
----------|------------|------------
Decode    | nil        | A JSON null.
Decode    | boolean    | A JSON boolean.
Decode    | number     | The nearest representation of a JSON number.
Decode    | string     | A JSON string.
Decode    | Array      | An JSON array.
Decode    | Dictionary | A JSON object.
Encode    | nil        | A Lua nil.
Encode    | boolean    | A Lua boolean.
Encode    | number     | A Lua number.
Encode    | string     | A Lua string.
Encode    | Array      | An array-like table, having a non-zero length.
Encode    | Dictionary | A dictionary-like table, having a length of zero.

Other value types are encoded as null.

This format has the following options:

Field  | Type   | Default | Description
-------|--------|---------|------------
Indent | string | `"\t"`  | Determines the indentation of encoded content. If an empty string, then the content is minified.
