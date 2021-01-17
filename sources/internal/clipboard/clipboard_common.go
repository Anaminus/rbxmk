package clipboard

// Format associates a media type with content.
type Format struct {
	// Name is a media type that indicates how to interpret Content.
	Name string
	// Content is the data to be set to the clipboard.
	Content []byte
}
