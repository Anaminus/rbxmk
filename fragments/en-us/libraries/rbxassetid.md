# Summary
An interface to Roblox assets.

# Description
The **rbxassetid** library provides an interface to assets on the Roblox
website.

$FIELDS

# Fields
## read
### Summary
Reads data from a rbxassetid in a certain format.

### Description
The **read** function downloads an asset according to the given
[options][RBXAssetOptions]. Returns the content of the asset corresponding to
AssetID, decoded according to Format.

Throws an error if a problem occurred while downloading the asset.

## write
### Summary
Writes data to a rbxassetid in a certain format.

### Description
The **write** function uploads to an existing asset according to the given
[options][RBXAssetOptions]. The Body is encoded according to Format, then
uploaded to AssetID. AssetID must be the ID of an existing asset.

Throws an error if a problem occurred while uploading the asset.
