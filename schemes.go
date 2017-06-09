package main

type InputScheme struct {
	Handler InputSchemeHandler
}

// InputSchemeHandler is used to retrieve a Source from a location. ref is the
// first value of node.Reference after it has been parsed by the protocol
// detector, and excludes the scheme ("scheme://") portion of the string, if
// it was given. Returns the retrieved Source, as well as node.Reference after
// it has been processed.
//
// If InputSchemeHandler used a Format to retrieve the source, then it ensures
// that node.Format is set to the Format's extension.
type InputSchemeHandler func(opt *Options, node *InputNode, ref string) (ext string, src *Source, err error)

type OutputScheme struct {
	Handler   OutputSchemeHandler // Get current state of output source from location (if needed)
	Finalizer OutputFinalizer     // Write final source to location
}

// OutputSchemeHandler is used to retrieve a Source from a location. ref is
// the first value of node.Reference after it has been parsed by the protocol
// detector, and excludes the scheme ("scheme://") portion of the string, if
// it was given. Returns the retrieved Source, as well as node.Reference after
// it has been processed.
//
// If retrieving the current state of the location is not applicable, then an
// empty or nil Source may be returned.
type OutputSchemeHandler func(opt *Options, node *OutputNode, ref string) (ext string, src *Source, err error)

// OutputFinalizer is used to write a modified Source to a location. ref is
// the first value of node.Reference after it has been parsed by the protocol
// detector, and excludes the scheme ("scheme://") portion of the string, if
// it was given.
type OutputFinalizer func(opt *Options, node *OutputNode, ref, ext string, outsrc *Source) (err error)

// func init() {
// 	RegisterInputScheme("file", HandleFileInputScheme)
// 	RegisterInputScheme("http", HandleHTTPInputScheme)
// 	RegisterInputScheme("https", HandleHTTPInputScheme)

// 	RegisterOutputScheme("file", HandleFileOutputScheme)
// }

func IsAlnum(s string) bool {
	for _, r := range s {
		if (r >= '0' && r <= '9') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= 'a' && r <= 'z') ||
			r == '_' {
			continue
		}
		return false
	}
	return true
}
func IsDigit(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			continue
		}
		return false
	}
	return true
}
