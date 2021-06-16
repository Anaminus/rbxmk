# Summary
Configures instance attributes.

# Description
The **AttrConfig** type configures how an instance encodes and decodes
attributes.

$MEMBERS

# Constructors
## new
### Summary
Creates a new AttrConfig.

### Description
The **new** constructor creates a new AttrConfig. *property* sets the
[Property][AttrConfig.Property] field, defaulting to an empty string.

# Properties
## Property
### Summary
The property that attributes are applied to.

### Description
The **Property** property determines which property of an [Instance][Instance]
attributes are applied to. If an empty string, instances will default to
"AttributesSerialize".
