# Summary
Upload an asset.

# Arguments

	[ FLAGS ] [ -id INT ] PATH

# Description
Uploads an asset to the roblox website.

The `-id` flag specifies the ID of the asset to upload. If not specified, then a
new asset will be created, and the ID of the asset will be returned.

The first non-flag argument is the path to a file to read from, which is
required. If the path is "-", then the file will be read from standard input.

Each cookie flag appends to the list of cookies that will be sent with the
request. Such flags can be specified any number of times.

# Flags
## id
The ID of the asset to download (required).

## format
The format to encode the asset as.

## file-format
The format to decode the file as. Defaults to -format.
