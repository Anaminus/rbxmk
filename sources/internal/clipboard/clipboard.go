// +build !windows

package clipboard

import "errors"

var notImplemented = errors.New("clipboard not implemented")

// Clear removes all data from the clipboard.
func Clear() error {
	return notImplemented
}

// Read gets data from the clipboard. If multiple clipboard formats are
// supported, Read selects the first format that matches one of the given
// media types.
//
// Each argument is a media type (e.g. "text/plain").
func Read(formats ...string) (b []byte, err error) {
	return nil, notImplemented
}

// Write sets data to the clipboard. If multiple formats are supported, then
// each given format is written according to the specified media type.
// Otherwise, which format is selected is implementation-defined.
func Write(formats []Format) (err error) {
	return notImplemented
}
