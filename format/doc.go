/*

Package format implements several file formats for the rbxmk command. These
formats can be registered to a rbxmk.Formats with the Register function.

This package also contains several reusable functions that implement
rbxmk.Drill and rbxmk.OutputMerger.

Formats

The following formats are registered:

	rbxl, rbxlx, rbxm, rbxmx

		Roblox Place and Model files. Reads and writes Instance data.

	lua

		Lua sourc file. Reads and writes Value data of the ProtectedString
		value type.

*/
package format
