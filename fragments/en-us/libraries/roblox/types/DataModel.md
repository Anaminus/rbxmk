# Summary
The root of an instance tree.

# Description
The **DataModel** type is a special case of an [Instance][Instance]. In addition
to the members of Instance, DataModel has the following members:

$MEMBERS

Unlike a normal Instance, the [ClassName][Instance.ClassName] property of a
DataModel cannot be modified. Properties on a DataModel are usually not
serialized.

A DataModel can be created with the [Instance.new][Instance.new] constructor
with "DataModel" as the *className*.

# Symbols
## Metadata
### Summary
Gets metadata of the DataModel.

### Description
The **Metadata** symbol gets or sets the metadata associated with the DataModel.
This metadata is used by certain formats (e.g. ExplicitAutoJoints).

# Methods
## GetService
### Summary
Gets a service instance from the tree.

### Description
The **GetService** method returns the first child of the DataModel whose
[ClassName][Instance.ClassName] equals *className*. If no such child exists,
then a new instance of *className* is created. The [Name][Instance.Name] of the
instance is set to *className*, [sym.IsService][Instance.sym.IsService] is set
to true, and [Parent][Instance.Parent] is set to the DataModel.

If the DataModel has a descriptor, then GetService will throw an error if the
created class's descriptor does not have the "Service" tag set.
