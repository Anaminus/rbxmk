# Sources
This document contains a reference to the sources available to rbxmk scripts.

<table>
<thead><tr><th>Table of Contents</th></tr></thead>
<tbody><tr><td>

1. [clipboard][clipboard]
	1. [clipboard.read][clipboard.read]
	2. [clipboard.write][clipboard.write]
2. [fs][fs]
	1. [fs.read][fs.read]
	2. [fs.write][fs.write]
3. [http][http]
	1. [http.request][http.request]
4. [rbxassetid][rbxassetid]
	1. [rbxassetid.read][rbxassetid.read]
	2. [rbxassetid.write][rbxassetid.write]

</td></tr></tbody>
</table>

A **source** is an interface to an external location outside of the rbxmk
environment. Each source has a corresponding library that provides access to the
source.

## clipboard
[clipboard]: #user-content-clipboard

The `clipboard` source provides an interface to the operating system's
clipboard.

Name                               | Description
-----------------------------------|------------
[clipboard.read][clipboard.read]   | Gets data from the clipboard in one of a number of formats.
[clipboard.write][clipboard.write] | Sets data to the clipboard in a number of formats.

**The clipboard is currently available only on Windows. Other operating systems
return no data.**

#### clipboard.read
[clipboard.read]: #user-content-clipboardread
<code>clipboard.read(formats: ...[string](##)): (value: [any](##))</code>

The `read` function gets a value from the clipboard according to one of the
given [formats](formats.md).

Each format has a number of associated [media
types](https://en.wikipedia.org/wiki/Media_type). Each format is traversed in
order, and each media type within a format is traversed in order. The data that
matches the first media type found in the clipboard is selected. This data is
decoded by the format corresponding to the matched media type, and the result is
returned.

Throws an error if *value* could not be decoded from the format, or if data
could not be retrieved from the clipboard.

#### clipboard.write
[clipboard.write]: #user-content-clipboardwrite
<code>clipboard.write(value: [any](##), formats: ...[string](##))</code>

The `write` function sets *value* to the clipboard according to the given
[formats](formats.md).

Each format has a number of associated [media
types](https://en.wikipedia.org/wiki/Media_type). For each format, the data is
encoded in the format, which is then sent to the clipboard for each of the
format's media type. Data is not sent for a media type if that media type was
already sent.

Throws an error if *value* could not be encoded in a format, or if data could
not be sent to the clipboard.

## fs
[fs]: #user-content-fs

The `fs` source provides an interface to the file system.

Name                 | Description
---------------------|------------
[fs.read][fs.read]   | Reads data from a file in a certain format.
[fs.write][fs.write] | Writes data to a file in a certain format.

#### fs.read
[fs.read]: #user-content-fsread
<code>fs.read(path: [string](##), format: [string](##)?): (value: [any](##))</code>

The `read` function reads the content of the file at *path*, and decodes it into
*value* according to the [format](formats.md) matching the file extension of
*path*. If *format* is given, then it will be used instead of the file
extension.

If the format returns an [Instance][Instance], then the Name property will be
set to the "fstem" component of *path* according to
[os.split](libraries.md#user-content-ossplit).

#### fs.write
[fs.write]: #user-content-fswrite
<code>fs.write(path: [string](##), value: [any](##), format: [string](##)?)</code>

The `write` function encodes *value* according to the [format](formats.md)
matching the file extension of *path*, and writes the result to the file at
*path*. If *format* is given, then it will be used instead of the file
extension.

## http
[http]: #user-content-http

The `http` source provides an interface to resources on the network via HTTP.

Name                         | Description
-----------------------------|------------
[http.request][http.request] | Begins an HTTP request.

#### http.request
[http.request]: #user-content-httprequest
<code>http.request(options: [HTTPOptions][HTTPOptions]): (req: [HTTPRequest][HTTPRequest])</code>

The `request` function begins a request with the specified
[options][HTTPOptions]. Returns a [request object][HTTPRequest] that may be
resolved or canceled. Throws an error if the request could not be started.

## rbxassetid
[rbxassetid]: #user-content-rbxassetid

The `rbxassetid` source provides an interface to assets on the Roblox website.

Name                                 | Description
-------------------------------------|------------
[rbxassetid.read][rbxassetid.read]   | Reads data from a rbxassetid in a certain format.
[rbxassetid.write][rbxassetid.write] | Writes data to a rbxassetid in a certain format.

#### rbxassetid.read
[rbxassetid.read]: #user-content-rbxassetidread
<code>rbxassetid.read(options: [RBXAssetOptions][RBXAssetOptions]): (value: [any](##))</code>

The `read` function downloads an asset according to the given
[options][RBXAssetOptions]. Returns the content of the asset corresponding to
AssetID, decoded according to Format.

Throws an error if a problem occurred while downloading the asset.

#### rbxassetid.write
[rbxassetid.write]: #user-content-rbxassetidwrite
<code>rbxassetid.write(options: [RBXAssetOptions][RBXAssetOptions])</code>

The `write` function uploads to an existing asset according to the given
[options][RBXAssetOptions]. The Body is encoding according to Format, then
uploaded to AssetID. AssetID must be the ID of an existing asset.

Throws an error if a problem occurred while uploading the asset.

[HTTPOptions]: types.md#user-content-httpoptions
[HTTPRequest]: types.md#user-content-httprequest
[Instance]: types.md#user-content-instance
[RBXAssetOptions]: types.md#user-content-rbxassetoptions
