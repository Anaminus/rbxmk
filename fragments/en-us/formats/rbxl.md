# Summary
Encodes binary place data.

# Description
The **rbxl** format encodes Instances in the Roblox binary place format.

Direction | Type                   | Description
----------|------------------------|------------
Decode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [DataModel][DataModel] | A DataModel instance.
Encode    | [Instance][Instance]   | A single instance, interpreted as a child to a DataModel.
Encode    | Objects                | A list of Instances, interpreted as children to a DataModel.

This format has the following options:

Field    | Type                                | Default       | Description
---------|-------------------------------------|---------------|------------
Desc     | [RootDesc][RootDesc] \| bool \| nil | `nil`         | Sets the descriptor to be used when encoding or decoding. If `false`, then no descriptor is used. Otherwise, the descriptor of the root instance is used **with all descendants**, falling back to [globalDesc][rbxmk.globalDesc].
DescMode | string                              | `"NonStrict"` | Determines how deviations from the descriptor are handled. `"NonStrict"` causes deviations to be ignored. `"Strict"` causes an error to be thrown for the first deviation. `"Preserve"` tries to retain as much information as possible, usually by behaving as if no descriptor is set.
