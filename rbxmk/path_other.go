//+build !shortpaths

package main

// shortenPath transforms the given path so that it is relative to the working
// directory. Returns the original path if that fails.
func shortenPath(filename string) string {
	return filename
}
