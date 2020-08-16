package rtypes

// Symbol is a unique identifier used for accessing members.
type Symbol struct {
	Name string
}

// Type returns a string identifying the type of the value.
func (Symbol) Type() string {
	return "Symbol"
}

// String returns a string representation of the value.
func (s Symbol) String() string {
	return "Symbol<" + s.Name + ">"
}
