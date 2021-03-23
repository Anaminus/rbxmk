# Summary
Download an asset.

# Arguments
	[ FLAGS ] -id INT [ PATH ]

# Description
The **download-asset** command downloads an asset from the roblox website.

The `-id` flag, which is required, specifies the ID of the asset to download.

The first non-flag argument is the path to a file to write to. If not specified,
then the file will be written to standard output.

Each cookie flag appends to the list of cookies that will be sent with the
request. Such flags can be specified any number of times.

# Flags
## file-format
The format to encode the file as. Defaults to -format.

## format
The format to decode the asset as.

## id
The ID of the asset to download (required).
