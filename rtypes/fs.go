package rtypes

import "io/fs"

const T_FileInfo = "FileInfo"

// FileInfo describes the metadata of a file.
type FileInfo struct {
	fs.FileInfo
}

// Type returns a string identifying the type of the value.
func (FileInfo) Type() string {
	return T_FileInfo
}

const T_DirEntry = "DirEntry"

// DirEntry is an entry read from a directory.
type DirEntry struct {
	fs.DirEntry
}

// Type returns a string identifying the type of the value.
func (DirEntry) Type() string {
	return T_DirEntry
}
