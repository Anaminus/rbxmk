# Summary
Encodes XML place data.

# Description
The **rbxlx** format encodes Instances in the Roblox XML place format.

Direction | Type                   | Description
----------|------------------------|------------
Decode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [Instance][Instance]   | A single instance, interpreted as a child to a DataModel.
Encode    | Objects                | A list of Instances, interpreted as children to a DataModel.

This format has the same options as the [rbxl][rbxl] format.