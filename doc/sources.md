# Sources
This document contains a reference to the sources available to rbxmk scripts.

<table>
<thead><tr><th>Table of Contents</th></tr></thead>
<tbody><tr><td>

1. [`clipboard` source][clipboard]
2. [`file` source][file]
3. [`http` source][http]
4. [`rbxassetid` source][rbxassetid]

</td></tr></tbody>
</table>

A **source** is an external location from which raw data can be read from and
written to. A source can be accessed in a general way through the
[`rbxmk.readSource`][readSource] and [`rbxmk.writeSource`][writeSource]
functions. This document describes the implementation of readSource and
writeSource for each source.

A source usually has a corresponding library that provides more detailed access.
Such libraries are described in the [Libraries](libraries.md) document.

Source                   | Library                          | Description
-------------------------|----------------------------------|------------
[clipboard][clipboard]   | [clipboard][clipboard-library]   | The OS clipboard.
[file][file]             | [file][file-library]             | The filesystem.
[http][http]             | [http][http-library]             | An HTTP resource.
[rbxassetid][rbxassetid] | [rbxassetid][rbxassetid-library] | A Roblox asset.

[readSource]: libraries.md#user-content-rbxmkreadsource
[writeSource]: libraries.md#user-content-rbxmkwritesource

## `clipboard` source
[clipboard]: #user-content-clipboard-source
[clipboard-library]: libraries.md#user-content-clipboard

The `clipboard` source provides access to the operating system's clipboard.

**The clipboard is currently available only on Windows. Other operating systems
return no data.**

### `readSource`
Each additional argument to [readSource][readSource] is a [format](formats.md)
that describes how interpret data retrieved from the clipboard.

Each format has a number of associated [media
types](https://en.wikipedia.org/wiki/Media_type). Each format is traversed in
order, and each media type within a format is traversed in order. The data that
matches the first media type found in the clipboard is returned.

The formats passed to readSource are used only to select data from the
clipboard. The returned data is still in raw bytes, and it is up to the user to
decoded it with the expected format.

```lua
local bytes = rbxmk.readSource("clipboard", "txt", "bin")
```

### `writeSource`
Each additional argument to [writeSource][writeSource] is a [format](formats.md)
that describes how to format data sent to the clipboard.

Each format has a number of associated [media
types](https://en.wikipedia.org/wiki/Media_type). For each given format, the
bytes are sent to the clipboard for each of the format's media types.

The formats passed to writeSource are used only to select the clipboard formats
to write to. The same bytes will be written to every clipboard format, and it is
up to the user to ensure that the data is correct for each clipboard format. For
more flexible encoding in multiple formats,
[clipboard.write](libraries.md#user-content-clipboardwrite) should be used
instead.

```lua
rbxmk.writeSource("clipboard", bytes, "txt", "bin")
```

## `file` source
[file]: #user-content-file-source
[file-library]: libraries.md#user-content-file

The `file` source provides access to the file system.

### `readSource`
The first additional argument to [readSource][readSource] is the path to the
file to read from.

```lua
local bytes = rbxmk.readSource("file", "path/to/file.ext")
```

### `writeSource`
The first additional argument to [writeSource][writeSource] is the path to the
file to write to.

```lua
rbxmk.writeSource("file", bytes, "path/to/file.ext")
```

## `http` source
[http]: #user-content-http-source
[http-library]: libraries.md#user-content-http

The `http` source provides access to an HTTP client.

### `readSource`
The first additional argument to [readSource][readSource] is an
[HTTPOptions][HTTPOptions]. Several options are ignored:

- Method is ignored. A GET request is always performed.
- RequestFormat is ignored. No body is sent with GET requests.
- ResponseFormat is ignored. The response is always decoded into raw bytes.
- Body is ignored. No body is sent with GET requests.

Returns the raw body of the response. Throws an error if the response status is
not 2XX.

```lua
local bytes = rbxmk.readSource("http", {URL="https://www.example.com/resource"})
```

### `writeSource`
The first additional argument to [writeSource][writeSource] is an
[HTTPOptions][HTTPOptions]. Several options are ignored:

- Method is ignored. A POST request is always performed.
- RequestFormat is ignored. The *bytes* argument of writeSource is used as the
  raw body of the request.
- ResponseFormat is ignored. The response body is discarded.
- Body is ignored. The *bytes* argument of writeSource is used as the raw body
  of the request.

Throws an error if the response status is not 2XX.

```lua
rbxmk.writeSource("http", bytes, {URL="https://www.example.com/resource"})
```

## `rbxassetid` source
[rbxassetid]: #user-content-rbxassetid-source
[rbxassetid-library]: libraries.md#user-content-rbxassetid

The `rbxassetid` source provides access to assets on the Roblox website.

### `readSource`
The first additional argument to [readSource][readSource] is a
[RBXWebOptions][RBXWebOptions]. Several options are ignored:

- Format is ignored. The content is always decoded into raw bytes.
- Body is ignored.

Returns the raw content of the asset. Throws an error if a problem occurred
while downloading the asset.

```lua
local bytes = rbxmk.readSource("rbxassetid", {AssetID=1818})
```

### `writeSource`
The first additional argument to [writeSource][writeSource] is a
[RBXWebOptions][RBXWebOptions]. Several options are ignored:

- Format is ignored. The *bytes* argument of writeSource is used as the raw
  content of the asset.
- Body is ignored. The *bytes* argument of writeSource is used as the raw
  content of the asset.

Throws an error if a problem occurred while uploading the asset.

```lua
rbxmk.writeSource("rbxassetid", bytes, {AssetID=1818})```
```

[HTTPOptions]: types.md#user-content-httpoptions
[RBXWebOptions]: types.md#user-content-rbxweboptions
