// The sfs package providing a secure file system, with access limitations.
package sfs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// Flags alter how access checks are made.
type Flags int

const (
	_    Flags = (1 << iota) / 2
	Root       // Path is accessible if it is a root directory.
)

// FS contains wrappers for common file system functions that also check whether
// a given path can be accessed. The zero value of an FS is ready for use.
//
// A path is accessible if it is a descendant of a root in FS. A root is not
// accessible unless the Root flag is specified.
type FS struct {
	mtx   sync.RWMutex
	roots []string
}

// access returns whether the given path can be accessed.
func (fs *FS) access(path string, flags Flags) bool {
	//TODO: Implement.
	return true
}

// AddRoot adds path as a root. Returns an error if the path could not be
// converted to an absolute path. Does nothing if the path is already a root.
func (fs *FS) AddRoot(path string) error {
	fs.mtx.Lock()
	defer fs.mtx.Unlock()

	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	for _, root := range fs.roots {
		if path == root {
			return nil
		}
	}
	fs.roots = append(fs.roots, path)
	return nil
}

// RemoveRoot removes path as a root. Returns an error if the path could not be
// converted to an absolute path. Does nothing if the path is not a root.
func (fs *FS) RemoveRoot(path string) error {
	fs.mtx.Lock()
	defer fs.mtx.Unlock()

	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	for i, root := range fs.roots {
		if path == root {
			fs.roots[i] = fs.roots[len(fs.roots)-1]
			fs.roots = fs.roots[:len(fs.roots)-1]
			return nil
		}
	}
	return nil
}

// accessible returns a common error if path cannot be accessed.
func (fs *FS) accessible(path string, flags Flags) error {
	if fs.access(path, flags) {
		return nil
	}
	return fmt.Errorf("%s: path not accessible", path)
}

// Accessible returns an error if path cannot be accessed.
func (fs *FS) Accessible(path string, flags Flags) error {
	fs.mtx.RLock()
	defer fs.mtx.RUnlock()
	return fs.accessible(path, flags)
}

// Create wraps os.Create, returning an error if name cannot be accessed.
func (fs *FS) Create(name string) (*os.File, error) {
	fs.mtx.RLock()
	defer fs.mtx.RUnlock()
	if err := fs.accessible(name, 0); err != nil {
		return nil, err
	}
	return os.Create(name)
}

// Create wraps os.Mkdir, returning an error if name cannot be accessed.
func (fs *FS) Mkdir(name string, perm os.FileMode) error {
	fs.mtx.RLock()
	defer fs.mtx.RUnlock()
	if err := fs.accessible(name, 0); err != nil {
		return err
	}
	return os.Mkdir(name, perm)
}

// Create wraps os.MkdirAll, returning an error if name cannot be accessed.
func (fs *FS) MkdirAll(name string, perm os.FileMode) error {
	fs.mtx.RLock()
	defer fs.mtx.RUnlock()
	if err := fs.accessible(name, 0); err != nil {
		return err
	}
	return os.MkdirAll(name, perm)
}

// Create wraps os.Open, returning an error if name cannot be accessed.
func (fs *FS) Open(name string) (*os.File, error) {
	fs.mtx.RLock()
	defer fs.mtx.RUnlock()
	if err := fs.accessible(name, 0); err != nil {
		return nil, err
	}
	return os.Open(name)
}

// Create wraps ioutil.ReadDir, returning an error if dirname cannot be
// accessed. dirname is allowed to be a root.
func (fs *FS) ReadDir(dirname string) ([]os.FileInfo, error) {
	fs.mtx.RLock()
	defer fs.mtx.RUnlock()
	if err := fs.accessible(dirname, Root); err != nil {
		return nil, err
	}
	return ioutil.ReadDir(dirname)
}

// Create wraps os.Remove, returning an error if name cannot be accessed.
func (fs *FS) Remove(name string) error {
	fs.mtx.RLock()
	defer fs.mtx.RUnlock()
	if err := fs.accessible(name, 0); err != nil {
		return err
	}
	return os.Remove(name)
}

// Create wraps os.RemoveAll, returning an error if name cannot be accessed.
func (fs *FS) RemoveAll(path string) error {
	fs.mtx.RLock()
	defer fs.mtx.RUnlock()
	if err := fs.accessible(path, 0); err != nil {
		return err
	}
	return os.RemoveAll(path)
}

// Create wraps os.Rename, returning an error if oldpath or newpath cannot be
// accessed.
func (fs *FS) Rename(oldpath, newpath string) error {
	fs.mtx.RLock()
	defer fs.mtx.RUnlock()
	if err := fs.accessible(oldpath, 0); err != nil {
		return err
	}
	if err := fs.accessible(newpath, 0); err != nil {
		return err
	}
	return os.Rename(oldpath, newpath)
}

// Create wraps os.Stat, returning an error if name cannot be accessed.
func (fs *FS) Stat(name string) (os.FileInfo, error) {
	fs.mtx.RLock()
	defer fs.mtx.RUnlock()
	if err := fs.accessible(name, 0); err != nil {
		return nil, err
	}
	return os.Stat(name)
}
