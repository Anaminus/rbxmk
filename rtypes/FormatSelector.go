package rtypes

import "github.com/robloxapi/types"

// FormatSelector selects a format and provides options for configuring the
// format.
type FormatSelector struct {
	Format  string
	Options Dictionary
}

// Type returns a string identifying the type of the value.
func (FormatSelector) Type() string {
	return "FormatSelector"
}

// String returns a string representation of the value.
func (f FormatSelector) String() string {
	return f.Format
}

// ValueOf returns the value of field. Returns nil if the value does not exist.
func (f FormatSelector) ValueOf(field string) types.Value {
	return f.Options[field]
}
