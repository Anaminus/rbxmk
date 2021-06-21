# Summary
Encodes instance attributes.

# Description
The **rbxattr** format is defined for serializing instance attributes, encoding
a Dictionary of attribute values.

Direction | Type       | Description
----------|------------|------------
Decode    | Dictionary | A dictionary of attribute names mapped to values.
Encode    | Dictionary | A dictionary of attribute names mapped to values.

The following value types are encoded and decoded:
- string
- bool
- float
- double
- UDim
- UDim2
- BrickColor
- Color3
- Vector2
- Vector3
- NumberSequence
- ColorSequence
- NumberRange
- Rect

Additionally, any Stringable value is encoded as a string, and any Numberable
value is encoded as a double.

This format has no options.
