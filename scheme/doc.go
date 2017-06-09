/*

Package scheme implements several input and output schemes for the rbxmk
command. These schemes can be registered to a rbxmk.Schemes with the Register
function.

Input schemes

The following input schemes are registered:

	file://<filename>

		Retrieve data from <filename> in the file system. The scheme portion
		("file://") is optional.

Output schemes

The following output schemes are registered:

	file://<filename>

		Write data to <filename> in the file system. The scheme portion
		("file://") is optional.

*/
package scheme
