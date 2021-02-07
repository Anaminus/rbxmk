// Modified from golang.org/x/dep/internal/fs.go

// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sfs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// HasFilepathPrefix will determine if "path" starts with "prefix" from
// the point of view of a filesystem.
//
// Unlike filepath.HasPrefix, this function is path-aware, meaning that
// it knows that two directories /foo and /foobar are not the same
// thing, and therefore HasFilepathPrefix("/foobar", "/foo") will return
// false.
//
// This function also handles the case where the involved filesystems
// are case-insensitive, meaning /foo/bar and /Foo/Bar correspond to the
// same file. In that situation HasFilepathPrefix("/Foo/Bar", "/foo")
// will return true. The implementation is *not* OS-specific, so a FAT32
// filesystem mounted on Linux will be handled correctly.
func HasFilepathPrefix(path, prefix string) (bool, error) {
	// this function is more convoluted then ideal due to need for special
	// handling of volume name/drive letter on Windows. vnPath and vnPrefix
	// are first compared, and then used to initialize initial values of p and
	// d which will be appended to for incremental checks using
	// IsCaseSensitiveFilesystem and then equality.

	// no need to check IsCaseSensitiveFilesystem because VolumeName return
	// empty string on all non-Windows machines
	vnPath := strings.ToLower(filepath.VolumeName(path))
	vnPrefix := strings.ToLower(filepath.VolumeName(prefix))
	if vnPath != vnPrefix {
		return false, nil
	}

	// Because filepath.Join("c:","dir") returns "c:dir", we have to manually
	// add path separator to drive letters. Also, we need to set the path root
	// on *nix systems, since filepath.Join("", "dir") returns a relative path.
	vnPath += string(os.PathSeparator)
	vnPrefix += string(os.PathSeparator)

	var dn string

	if isDir, err := IsDir(path); err != nil {
		return false, fmt.Errorf("failed to check filepath prefix: %w", err)
	} else if isDir {
		dn = path
	} else {
		dn = filepath.Dir(path)
	}

	dn = filepath.Clean(dn)
	prefix = filepath.Clean(prefix)

	// [1:] in the lines below eliminates empty string on *nix and volume name on Windows
	dirs := strings.Split(dn, string(os.PathSeparator))[1:]
	prefixes := strings.Split(prefix, string(os.PathSeparator))[1:]

	if len(prefixes) > len(dirs) {
		return false, nil
	}

	// d,p are initialized with "/" on *nix and volume name on Windows
	d := vnPath
	p := vnPrefix

	var caseSensitive bool
	for i := range prefixes {
		// need to test each component of the path for
		// case-sensitiveness because on Unix we could have
		// something like ext4 filesystem mounted on FAT
		// mountpoint, mounted on ext4 filesystem, i.e. the
		// problematic filesystem is not the last one.
		cs, err := IsCaseSensitiveFilesystem(filepath.Join(d, dirs[i]))
		if err == nil {
			caseSensitive = cs
		}
		// If case-sensitivity check fails, reuse previous state. A failing path
		// is assumed to inherit the latest known state.
		if caseSensitive {
			d = filepath.Join(d, dirs[i])
			p = filepath.Join(p, prefixes[i])
		} else {
			d = filepath.Join(d, strings.ToLower(dirs[i]))
			p = filepath.Join(p, strings.ToLower(prefixes[i]))
		}

		if p != d {
			return false, nil
		}
	}

	return true, nil
}

// EquivalentPaths compares the paths passed to check if they are equivalent.
// It respects the case-sensitivity of the underlying filesysyems.
func EquivalentPaths(p1, p2 string) (bool, error) {
	p1 = filepath.Clean(p1)
	p2 = filepath.Clean(p2)

	fi1, err := os.Stat(p1)
	if err != nil {
		return false, fmt.Errorf("could not check for path equivalence: %w", err)
	}
	fi2, err := os.Stat(p2)
	if err != nil {
		return false, fmt.Errorf("could not check for path equivalence: %w", err)
	}

	p1Filename, p2Filename := "", ""

	if !fi1.IsDir() {
		p1, p1Filename = filepath.Split(p1)
	}
	if !fi2.IsDir() {
		p2, p2Filename = filepath.Split(p2)
	}

	if isPrefix1, err := HasFilepathPrefix(p1, p2); err != nil {
		return false, fmt.Errorf("failed to check for path equivalence: %w", err)
	} else if isPrefix2, err := HasFilepathPrefix(p2, p1); err != nil {
		return false, fmt.Errorf("failed to check for path equivalence: %w", err)
	} else if !isPrefix1 || !isPrefix2 {
		return false, nil
	}

	if p1Filename != "" || p2Filename != "" {
		caseSensitive, err := IsCaseSensitiveFilesystem(filepath.Join(p1, p1Filename))
		if err != nil {
			return false, fmt.Errorf("could not check for filesystem case-sensitivity: %w", err)
		}
		if caseSensitive {
			if p1Filename != p2Filename {
				return false, nil
			}
		} else {
			if !strings.EqualFold(p1Filename, p2Filename) {
				return false, nil
			}
		}
	}

	return true, nil
}

// IsCaseSensitiveFilesystem determines if the filesystem where dir
// exists is case sensitive or not.
//
// CAVEAT: this function works by taking the last component of the given
// path and flipping the case of the first letter for which case
// flipping is a reversible operation (/foo/Bar → /foo/bar), then
// testing for the existence of the new filename. There are two
// possibilities:
//
// 1. The alternate filename does not exist. We can conclude that the
// filesystem is case sensitive.
//
// 2. The filename happens to exist. We have to test if the two files
// are the same file (case insensitive file system) or different ones
// (case sensitive filesystem).
//
// If the input directory is such that the last component is composed
// exclusively of case-less codepoints (e.g.  numbers), this function will
// return false.
func IsCaseSensitiveFilesystem(dir string) (bool, error) {
	alt := filepath.Join(filepath.Dir(dir), genTestFilename(filepath.Base(dir)))

	dInfo, err := os.Stat(dir)
	if err != nil {
		return false, fmt.Errorf("could not determine the case-sensitivity of the filesystem: %w", err)
	}

	aInfo, err := os.Stat(alt)
	if err != nil {
		// If the file doesn't exists, assume we are on a case-sensitive filesystem.
		if os.IsNotExist(err) {
			return true, nil
		}

		return false, fmt.Errorf("could not determine the case-sensitivity of the filesystem: %w", err)
	}

	return !os.SameFile(dInfo, aInfo), nil
}

// genTestFilename returns a string with at most one rune case-flipped.
//
// The transformation is applied only to the first rune that can be
// reversibly case-flipped, meaning:
//
// * A lowercase rune for which it's true that lower(upper(r)) == r
// * An uppercase rune for which it's true that upper(lower(r)) == r
//
// All the other runes are left intact.
func genTestFilename(str string) string {
	flip := true
	return strings.Map(func(r rune) rune {
		if flip {
			if unicode.IsLower(r) {
				u := unicode.ToUpper(r)
				if unicode.ToLower(u) == r {
					r = u
					flip = false
				}
			} else if unicode.IsUpper(r) {
				l := unicode.ToLower(r)
				if unicode.ToUpper(l) == r {
					r = l
					flip = false
				}
			}
		}
		return r
	}, str)
}

// IsDir determines is the path given is a directory or not.
func IsDir(name string) (bool, error) {
	fi, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return fi.IsDir(), nil
}

// IsRegular determines if the path given is a regular file or not.
func IsRegular(name string) (bool, error) {
	fi, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	mode := fi.Mode()
	if mode&os.ModeType != 0 {
		return false, fmt.Errorf("%q is a %v, expected a file", name, mode)
	}
	return true, nil
}

// IsSymlink determines if the given path is a symbolic link.
func IsSymlink(path string) (bool, error) {
	l, err := os.Lstat(path)
	if err != nil {
		return false, err
	}

	return l.Mode()&os.ModeSymlink == os.ModeSymlink, nil
}
