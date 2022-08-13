package dump

// Root describes an entire API.
type Root struct {
	// Libraries contains libraries defined in the API.
	Libraries Libraries `json:",omitempty"`
	// Types contains types defined by the API.
	Types TypeDefs `json:",omitempty"`
	// Enums contains enums defined by the API.
	Enums Enums `json:",omitempty"`
	// Formats contains formats registered by a world.
	Formats Formats `json:",omitempty"`
	// Program contains the root command created by the program.
	Program Command
}
