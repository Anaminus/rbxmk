//go:build !windows

package clipboard

// Clear removes all data from the clipboard.
func Clear() error {
	return NoDataError{notImplemented: true}
}

// Read gets data from the clipboard. If multiple clipboard formats are
// supported, Read selects the first format that matches one of the given
// media types.
//
// Each argument is a media type (e.g. "text/plain").
//
// If an error is returned, then f will be less than 0. If no data was found,
// then the error will contain NoDataError. If no formats were given, then f
// will be less than 0, and err will be nil.
func Read(formats ...string) (f int, b []byte, err error) {
	return 0, nil, NoDataError{notImplemented: true}
}

// Write sets data to the clipboard. If multiple formats are supported, then
// each given format is written according to the specified media type.
// Otherwise, which format is selected is implementation-defined.
//
// If no formats are given, then the clipboard is cleared with no other action.
func Write(formats []Format) (err error) {
	return NoDataError{notImplemented: true}
}
