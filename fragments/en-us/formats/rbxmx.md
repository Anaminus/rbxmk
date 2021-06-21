# Summary
Encodes XML model data.

# Description
The **rbxmx** format encodes Instances in the Roblox XML model format.

Direction | Type                   | Description
----------|------------------------|------------
Decode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [Instance][Instance]   | A single instance, interpreted as a child to a DataModel.
Encode    | Objects                | A list of Instances, interpreted as children to a DataModel.

This format has the same options as the [rbxl][rbxl] format.
