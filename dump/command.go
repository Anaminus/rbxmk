package dump

// Commands maps a name to a command.
type Commands map[string]Command

// Resolve implements Node.
func (c Commands) Resolve(path ...string) any {
	if len(path) == 0 {
		return c
	}
	if v, ok := c[path[0]]; ok {
		return v.Resolve(path[1:]...)
	}
	return nil
}

// Command describes a program command.
type Command struct {
	// Aliases lists available aliases for the command.
	Aliases []string `json:",omitempty"`
	// Hidden indicates whether the command is hidden.
	Hidden bool `json:",omitempty"`

	// Arguments is a fragment reference pointing to a definition of the
	// command's arguments.
	Arguments string `json:",omitempty"`
	// Summary is a fragment reference pointing to a short summary of the
	// command.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the command.
	Description string `json:",omitempty"`
	// Deprecated is a fragment reference pointing to a message detailing the
	// deprecation of the command.
	Deprecated string `json:",omitempty"`

	// Flags contains the flags defined on the command.
	Flags Flags `json:",omitempty"`
	// Commands contains subcommands defined on the command.
	Commands Commands `json:",omitempty"`
}

// Resolve implements Node.
func (c Command) Resolve(path ...string) any {
	if len(path) == 0 {
		return c
	}
	switch name, path := path[0], path[1:]; name {
	case "Flags":
		return c.Flags.Resolve(path...)
	case "Commands":
		return c.Commands.Resolve(path...)
	}
	return nil
}

// Flags maps a name to a flag.
type Flags map[string]Flag

// Resolve implements Node.
func (f Flags) Resolve(path ...string) any {
	if len(path) == 0 {
		return f
	}
	if v, ok := f[path[0]]; ok {
		return resolveValue(path[1:], v)
	}
	return nil
}

// Flag describes a command flag.
type Flag struct {
	// Type indicates the value type of the flag.
	Type string
	// Default indicates the default value for the flag.
	Default string `json:",omitempty"`
	// Shorthand indicates a one-letter abbreviation for the flag.
	Shorthand string `json:",omitempty"`
	// Hidden indicates whether the flag is hidden.
	Hidden bool `json:",omitempty"`
	// Whether the flag is inherited by subcommands.
	Persistent bool `json:",omitempty"`
	// Description is a fragment reference pointing to a description of the
	// flag.

	Description string `json:",omitempty"`
	// Deprecated indicates whether the flag is deprecated, and if so, a
	// fragment reference pointing to a message describing the deprecation.
	Deprecated string `json:",omitempty"`
	// ShorthandDeprecated indicates whether the shorthand of the flag is
	// deprecated, and if so, a fragment reference pointing to a message
	// describing the deprecation.
	ShorthandDeprecated string `json:",omitempty"`
}
