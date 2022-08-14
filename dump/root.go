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
	// Environment recursively maps the content of the Lua environment.
	Environment *EnvRef
}

// Resolve implements Node.
func (r Root) Resolve(path ...string) any {
	if len(path) == 0 {
		return r
	}
	switch name, path := path[0], path[1:]; name {
	case "Libraries":
		return r.Libraries.Resolve(path...)
	case "Types":
		return r.Types.Resolve(path...)
	case "Enums":
		return r.Enums.Resolve(path...)
	case "Formats":
		return r.Formats.Resolve(path...)
	case "Program":
		return r.Program.Resolve(path...)
	}
	return nil
}

// EnvRef represents a value in an environment, which has an associated dump
// object referred to by Path.
type EnvRef struct {
	Path   []string           `json:",omitempty"`
	Fields map[string]*EnvRef `json:",omitempty"`
}
