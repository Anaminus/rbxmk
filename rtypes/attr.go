package rtypes

// AttrConfig configures an Instance's attributes API.
type AttrConfig struct {
	// Property is the name of the property to which attributes will be
	// serialized. An empty string defaults to "AttributesSerialize".
	Property string
}

// Type returns a string identifying the type of the value.
func (*AttrConfig) Type() string {
	return "AttrConfig"
}

// String returns a string representation of the value.
func (*AttrConfig) String() string {
	return "Attr"
}
