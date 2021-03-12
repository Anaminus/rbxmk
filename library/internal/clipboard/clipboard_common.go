package clipboard

import "errors"

// Format associates a media type with content.
type Format struct {
	// Name is a media type that indicates how to interpret Content.
	Name string
	// Content is the data to be set to the clipboard.
	Content []byte
}

// NoData indicates that no data was found for any of a number of given formats.
type NoDataError struct {
	notImplemented bool
}

func (err NoDataError) Error() string {
	if err.notImplemented {
		return "clipboard not implemented"
	}
	return "no clipboard data for given formats"
}

// NotImplemented indicates that the clipboard has no data because it is not
// implemented for the platform.
func (err NoDataError) NotImplemented() bool {
	return err.notImplemented
}

// IsNoData returns whether an error indicates that no data was found.
func IsNoData(err error) bool {
	var e NoDataError
	return errors.As(err, &e)
}
