# Summary
Specifies the options of an asset request.

# Description
The **RBXAssetOptions** type is a table that specifies the options of a request
to an asset on the Roblox website. It has the following fields:

Field          | Type                             | Description
---------------|----------------------------------|------------
AssetID        | [int64](##)                      | The ID of the asset to request.
Cookies        | [Cookies][Cookies]?              | Optional cookies to send with requests, usually used for authentication.
Format         | [FormatSelector][FormatSelector] | The format used to encode or decode an asset.
Body           | [any](##)?                       | The body of an asset, to be encoded by the specified format.
