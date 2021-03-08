package rbxmk

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"unsafe"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/sfs"
	"github.com/robloxapi/types"
)

// World contains the entire state of a Lua environment, including a Lua state,
// and registered Reflectors, Formats, and Sources.
type World struct {
	l          *lua.LState
	fileStack  []FileEntry
	rootdir    string
	reflectors map[string]Reflector
	formats    map[string]Format

	Global

	tmponce sync.Once
	tmpdir  string

	Client *Client
	FS     sfs.FS

	udmut    sync.Mutex
	userdata map[interface{}]uintptr
}

// NewWorld returns a World initialized with the given Lua state.
func NewWorld(l *lua.LState) *World {
	return &World{
		l:      l,
		Client: NewClient(nil),
	}
}

// cacheUserData causes userdata created by the world to be cached per value.
const cacheUserdata = true

// UserDataOf returns the userdata value associated with v. If there is no such
// userdata, then a new one is created, with the metatable set to the type
// corresponding to t. The Value field of the userdata must never be modified.
func (w *World) UserDataOf(v types.Value, t string) *lua.LUserData {
	if !cacheUserdata {
		// Fallback in case it turns out that pointer fiddling fails
		// catastrophically for some reason.
		u := w.State().NewUserData()
		u.Value = v
		w.l.SetMetatable(u, w.l.GetTypeMetatable(t))
		return u
	}

	// Normally, a new userdata will be created every single time a value needs
	// to be pushed. This is fine for most cases; __eq will take care of most
	// comparison checks. One problem is that such userdata cannot be properly
	// used as a table key, because the table doesn't know when two userdata
	// refer to the same underlying value.
	//
	// To fix this, a value must consistently map to the same userdata. This
	// could be done by caching the association of the value to its userdata in
	// a map. The problem is that this map will accumulate garbage userdata that
	// has no references other than the map itself. Finalizers can't be used
	// because the map still has that single reference.
	//
	// This problem is resolved by storing a uintptr pointing to the userdata
	// instead. By eliminating the strong reference to the userdata, a finalizer
	// can be used to remove the association when the userdata has no more
	// references.

	w.udmut.Lock()
	defer w.udmut.Unlock()

	if p, ok := w.userdata[v]; ok {
		u := (*lua.LUserData)(unsafe.Pointer(p))
		return u
	}

	u := w.State().NewUserData()
	u.Value = v
	w.l.SetMetatable(u, w.l.GetTypeMetatable(t))

	if w.userdata == nil {
		w.userdata = map[interface{}]uintptr{}
	}
	w.userdata[v] = uintptr(unsafe.Pointer(u))
	runtime.SetFinalizer(u, func(u *lua.LUserData) {
		w.udmut.Lock()
		defer w.udmut.Unlock()
		delete(w.userdata, u.Value)
	})
	return u
}

// UserDataCacheLen return the number of userdata values in the cache.
func (w *World) UserDataCacheLen() int {
	w.udmut.Lock()
	defer w.udmut.Unlock()
	return len(w.userdata)
}

// Library represents a Lua library.
type Library struct {
	// Name is the default name of the library.
	Name string
	// Open returns a table with the contents of the library.
	Open func(s State) *lua.LTable
	// Dump returns a description of the library's API.
	Dump func(s State) dump.Library
}

// Open opens lib according to OpenAs by using the Name of the library.
func (w *World) Open(lib Library) error {
	return w.OpenAs(lib.Name, lib)
}

// OpenAs opens lib, then merges the result into the world's global table using
// name. If the name is already present and is a table, then the table of lib is
// merged into it, preferring the values of lib. If name is an empty string,
// then lib is merged into the global table itself. An error is returned if the
// merged value is not a table.
func (w *World) OpenAs(name string, lib Library) error {
	src := lib.Open(State{World: w, L: w.l})
	// TODO: Is lua.LState.G.Global safe to use?
	return w.MergeTables(w.l.G.Global, src, name)
}

// MargeTables merges src into dst according to name. If name is empty, then
// each key in src is set in dst. If dst[name] is a table, then each key in src
// is set in that table. If dst[name] is nil, then it is set directly to src.
// Does nothing if src or dst is nil. Returns an error if the tables could not
// be merged.
func (w *World) MergeTables(dst, src *lua.LTable, name string) error {
	if src == nil || dst == nil {
		return nil
	}
	if name == "" {
		src.ForEach(func(k, v lua.LValue) error {
			dst.RawSet(k, v)
			return nil
		})
		return nil
	}
	switch u := dst.RawGetString(name).(type) {
	case *lua.LTable:
		src.ForEach(func(k, v lua.LValue) error {
			u.RawSet(k, v)
			return nil
		})
	case *lua.LNilType:
		dst.RawSetString(name, src)
	default:
		return fmt.Errorf("cannot merge %s into %s", name, u.Type().String())
	}
	return nil
}

// createTypeMetatable constructs a metatable from the given Reflector. If
// Members and Exprim is set, then the Value field will be injected if it does
// not already exist.
func (w *World) createTypeMetatable(r Reflector) (mt *lua.LTable) {
	if r.Metatable == nil && r.Members == nil && r.Flags&Exprim == 0 {
		// No metatable.
		return nil
	}

	if r.Flags&Exprim != 0 {
		// Inject Value field, if possible.
		if r.Members == nil {
			r.Members = make(map[string]Member, 1)
		}
		if _, ok := r.Members["Value"]; !ok {
			r.Members["Value"] = Member{
				Get: func(s State, v types.Value) int {
					if v, ok := v.(types.Aliaser); ok {
						// Push underlying type.
						return s.Push(v.Alias())
					}
					// Fallback to current value.
					return s.Push(v)
				},
			}
		}
	}
	if r.Members != nil {
		// Validate members.
		for _, member := range r.Members {
			if member.Get == nil {
				panic("member must define Get function")
			}
		}
	}

	mt = w.l.CreateTable(0, 8)

	// Unconditional fields.
	mt.RawSetString("__type", lua.LString(r.Name))
	mt.RawSetString("__metatable", lua.LString("the metatable is locked"))

	if r.Flags&Exprim != 0 {
		// Show type and value, if possible.
		mt.RawSetString("__tostring", w.l.NewFunction(func(l *lua.LState) int {
			if u, ok := l.Get(1).(*lua.LUserData); ok {
				if v, ok := u.Value.(types.Stringer); ok {
					l.Push(lua.LString(r.Name + ": " + v.String()))
					return 1
				}
			}
			l.Push(lua.LString(r.Name))
			return 1
		}))
	} else {
		// Just show type.
		mt.RawSetString("__tostring", w.l.NewFunction(func(l *lua.LState) int {
			l.Push(lua.LString(r.Name))
			return 1
		}))
	}

	var index Metamethod
	var newindex Metamethod
	if r.Metatable != nil {
		// Set each defined metamethod, overriding predefined values.
		for name, method := range r.Metatable {
			m := method
			mt.RawSetString(name, w.WrapFunc(m))
		}
		// If available, remember index and newindex for member indexing.
		index = r.Metatable["__index"]
		newindex = r.Metatable["__newindex"]
	}

	if r.Members != nil {
		// Setup member getting and setting.
		mt.RawSetString("__index", w.l.NewFunction(func(l *lua.LState) int {
			u := l.CheckUserData(1)
			if u.Metatable != mt {
				TypeError(l, 1, r.Name)
				return 0
			}
			v, ok := u.Value.(types.Value)
			if !ok {
				TypeError(l, 1, r.Name)
				return 0
			}
			idx := l.Get(2)
			name, ok := idx.(lua.LString)
			if ok {
				member, ok := r.Members[string(name)]
				if !ok {
					goto customIndex
				}
				if member.Method {
					// Push as method.
					l.Push(l.NewFunction(func(l *lua.LState) int {
						// TODO: validate that s.L.Get(1) matches v, or at least has
						// the expected type.
						return member.Get(State{World: w, L: l}, v)
					}))
					return 1
				}
				return member.Get(State{World: w, L: l}, v)
			}
		customIndex:
			if index != nil {
				// Fallback to custom index.
				return index(State{World: w, L: l})
			}
			if ok {
				l.RaiseError("%q is not a valid member of %s", name, r.Name)
			} else {
				l.ArgError(2, "string expected, got "+idx.Type().String())
			}
			return 0
		}))
		mt.RawSetString("__newindex", w.l.NewFunction(func(l *lua.LState) int {
			u := l.CheckUserData(1)
			if u.Metatable != mt {
				TypeError(l, 1, r.Name)
				return 0
			}
			v, ok := u.Value.(types.Value)
			if !ok {
				TypeError(l, 1, r.Name)
				return 0
			}
			idx := l.Get(2)
			name, ok := idx.(lua.LString)
			if ok {
				member, ok := r.Members[string(name)]
				if !ok {
					goto customNewindex
				}
				if member.Method || member.Set == nil {
					l.RaiseError("%s cannot be assigned to", name)
				}
				member.Set(State{World: w, L: l}, v)
				return 0
			}
		customNewindex:
			if newindex != nil {
				// Fallback to custom newindex.
				return newindex(State{World: w, L: l})
			}
			if ok {
				l.RaiseError("%q is not a valid member of %s", name, r.Name)
			} else {
				l.ArgError(2, "string expected, got "+idx.Type().String())
			}
			return 0
		}))
	}

	return mt
}

// RegisterReflector registers a reflector. If the reflector produces a
// metatable, then it is added as a type metatable to the world's state. Panics
// if the reflector is already registered.
func (w *World) RegisterReflector(r Reflector) {
	if _, ok := w.reflectors[r.Name]; ok {
		panic("reflector " + r.Name + " already registered")
	}
	if w.reflectors == nil {
		w.reflectors = map[string]Reflector{}
	}
	w.reflectors[r.Name] = r

	if mt := w.createTypeMetatable(r); mt != nil {
		w.l.SetField(w.l.Get(lua.RegistryIndex), r.Name, mt)
	}
}

// Reflector returns the Reflector registered with the given name. If the name
// is not registered, then Reflector.Name will be an empty string.
func (w *World) Reflector(name string) Reflector {
	return w.reflectors[name]
}

// Reflectors returns a list of reflectors that have all of the given flags set.
func (w *World) Reflectors(flags ReflectorFlags) []Reflector {
	ts := []Reflector{}
	for _, t := range w.reflectors {
		if t.Flags&flags == flags {
			ts = append(ts, t)
		}
	}
	sort.Slice(ts, func(i, j int) bool {
		return ts[i].Name < ts[j].Name
	})
	return ts
}

// ApplyReflector applies a reflector to a table by setting contructors and
// initializing the reflector's environment.
func (w *World) ApplyReflector(r Reflector, t *lua.LTable) {
	if r.Constructors != nil {
		ctors := w.l.CreateTable(0, len(r.Constructors))
		for name, ctor := range r.Constructors {
			if c := ctor.Func; c != nil {
				ctors.RawSetString(name, w.l.NewFunction(func(l *lua.LState) int {
					return c(State{World: w, L: w.l})
				}))
			}
		}
		t.RawSetString(r.Name, ctors)
	}
	if r.Environment != nil {
		r.Environment(State{World: w, L: w.l}, t)
	}
}

// Typeof returns the type of the given Lua value. If it is a userdata, Typeof
// attempts to get the type according to the value's metatable. Panics if v is
// nil (not if nil Lua value).
func (w *World) Typeof(v lua.LValue) string {
	if v == nil {
		panic("value expected")
	}
	u, ok := v.(*lua.LUserData)
	if !ok {
		return v.Type().String()
	}
	t, ok := w.l.GetMetaField(u, "__type").(lua.LString)
	if !ok {
		return u.Type().String()
	}
	return string(t)
}

// RegisterFormat registers a format. Panics if the format is already
// registered.
func (w *World) RegisterFormat(f Format) {
	f.Name = strings.TrimPrefix(f.Name, ".")
	if _, ok := w.formats[f.Name]; ok {
		panic("format " + f.Name + " already registered")
	}
	if _, ok := f.Options["Format"]; ok {
		panic("format " + f.Name + " specifies reserved \"Format\" option")
	}
	if w.formats == nil {
		w.formats = map[string]Format{}
	}
	w.formats[f.Name] = f
}

// Format returns the Format registered with the given name. If the name is not
// registered, then Format.Name will be an empty string.
func (w *World) Format(name string) Format {
	return w.formats[strings.TrimPrefix(name, ".")]
}

// Formats returns a list of registered formats.
func (w *World) Formats() []Format {
	formats := []Format{}
	for _, format := range w.formats {
		formats = append(formats, format)
	}
	sort.Slice(formats, func(i, j int) bool {
		return formats[i].Name < formats[j].Name
	})
	return formats
}

// Ext returns the extension of filename that most closely matches the name of a
// registered format. Returns an empty string if no format was found.
func (w *World) Ext(filename string) (ext string) {
	i := len(filename) - 1
	for ; i >= 0 && !os.IsPathSeparator(filename[i]); i-- {
	}
	filename = filename[i+1:]
	for {
		for i := 0; i < len(filename); i++ {
			if filename[i] == '.' {
				filename = filename[i+1:]
				goto check
			}
		}
		return ""
	check:
		if w.Format(filename).Name != "" {
			return filename
		}
	}
}

// Expand expands a string containing predefined variables.
func (w *World) Expand(path string) string {
	return os.Expand(path, func(v string) string {
		switch v {
		case "script_name", "sn":
			if entry, ok := w.PeekFile(); ok {
				if entry.Path == "" {
					return ""
				}
				path, _ := filepath.Abs(entry.Path)
				return filepath.Base(path)
			}
		case "script_directory", "script_dir", "sd":
			if entry, ok := w.PeekFile(); ok {
				if entry.Path == "" {
					return ""
				}
				path, _ := filepath.Abs(entry.Path)
				return filepath.Dir(path)
			}
		case "root_script_directory", "root_script_dir", "rsd":
			return w.RootDir()
		case "working_directory", "working_dir", "wd":
			wd, _ := os.Getwd()
			return wd
		case "temp_directory", "temp_dir", "tmp":
			return w.TempDir()
		}
		return ""
	})
}

// Split returns the components of a file path.
func (w *World) Split(path string, components ...string) ([]string, error) {
	parts := make([]string, len(components))
	for i, comp := range components {
		var result string
		switch comp {
		case "dir":
			result = filepath.Dir(path)
		case "base":
			result = filepath.Base(path)
		case "ext":
			result = filepath.Ext(path)
		case "stem":
			result = filepath.Base(path)
			result = result[:len(result)-len(filepath.Ext(path))]
		case "fext":
			result = w.Ext(path)
			if result != "" && result != "." {
				result = "." + result
			}
		case "fstem":
			ext := w.Ext(path)
			if ext != "" && ext != "." {
				ext = "." + ext
			}
			result = filepath.Base(path)
			result = result[:len(result)-len(ext)]
		default:
			return nil, fmt.Errorf("unknown argument %q", comp)
		}
		parts[i] = result
	}
	return parts, nil
}

// State returns the underlying Lua state.
func (w *World) State() *lua.LState {
	return w.l
}

// WrapFunc wraps a function that receives a State into a Lua function.
func (w *World) WrapFunc(f func(State) int) *lua.LFunction {
	return w.l.NewFunction(func(l *lua.LState) int {
		return f(State{World: w, L: l})
	})
}

// PushTo reflects v to lvs using registered type t.
func (w *World) PushTo(t string, v types.Value) (lvs []lua.LValue, err error) {
	rfl := w.reflectors[t]
	if rfl.Name == "" {
		return nil, fmt.Errorf("unknown type %q", t)
	}
	if rfl.PushTo == nil {
		return nil, fmt.Errorf("cannot cast type %q to Lua", t)
	}
	return rfl.PushTo(State{World: w, L: w.l}, v)
}

// PullFrom reflects lvs to v using registered type t.
func (w *World) PullFrom(t string, lvs ...lua.LValue) (v types.Value, err error) {
	rfl := w.reflectors[t]
	if rfl.Name == "" {
		return nil, fmt.Errorf("unknown type %q", t)
	}
	if rfl.PullFrom == nil {
		return nil, fmt.Errorf("cannot cast type %q from Lua", t)
	}
	return rfl.PullFrom(State{World: w, L: w.l}, lvs...)
}

// Push reflects v according to its type as registered, then pushes the results
// to the world's state.
func (w *World) Push(v types.Value) int {
	return State{World: w, L: w.l}.Push(v)
}

// Pull gets from the world's Lua state the values starting from n, and reflects
// a value from them according to registered type t.
func (w *World) Pull(n int, t string) types.Value {
	return State{World: w, L: w.l}.Pull(n, t)
}

// PullOpt gets from the world's Lua state the value at n, and reflects a value
// from it according to registered type t. If the value is nil, d is returned
// instead.
func (w *World) PullOpt(n int, t string, d types.Value) types.Value {
	return State{World: w, L: w.l}.PullOpt(n, t, d)
}

// PullAnyOf gets from the world's Lua state the values starting from n, and
// reflects a value from them according to any of the registered types in t.
// Returns the first successful reflection among the types in t. If no types
// succeeded, then a type error is thrown.
func (w *World) PullAnyOf(n int, t ...string) types.Value {
	return State{World: w, L: w.l}.PullAnyOf(n, t...)
}

// FileEntry describes a file, including the full path. An empty Path indicates
// stdin.
type FileEntry struct {
	Path string
	os.FileInfo
}

// PushFile marks a file as the currently running file. Returns an error if the
// file is already running. If the file is the first file pushed, its directory
// is added as a root to w.FS.
func (w *World) PushFile(entry FileEntry) error {
	for _, f := range w.fileStack {
		if os.SameFile(entry.FileInfo, f.FileInfo) {
			return fmt.Errorf("\"%s\" is already running", entry.Path)
		}
	}
	if len(w.fileStack) == 0 {
		// If not stdin.
		if entry.Path != "" {
			// Set RootDir to file at bottom of stack.
			if abs, err := filepath.Abs(entry.Path); err == nil {
				w.rootdir = filepath.Dir(abs)
				w.FS.AddRoot(w.rootdir)
			}
		}
	}
	w.fileStack = append(w.fileStack, entry)
	return nil
}

// PopFile unmarks the currently running file. If the last file on the stack is
// popped, the file's directory is removed as a root from w.FS.
func (w *World) PopFile() {
	if len(w.fileStack) > 0 {
		w.fileStack[len(w.fileStack)-1] = FileEntry{}
		w.fileStack = w.fileStack[:len(w.fileStack)-1]
		if len(w.fileStack) == 0 {
			w.FS.RemoveRoot(w.rootdir)
			w.rootdir = ""
		}
	}
}

// PeekFile returns the info of the currently running file. Returns false if
// there is no running file.
func (w *World) PeekFile() (entry FileEntry, ok bool) {
	if len(w.fileStack) == 0 {
		return
	}
	entry = w.fileStack[len(w.fileStack)-1]
	ok = true
	return
}

// DoString executes string s as Lua. args is the number of arguments currently
// on the stack that should be passed in.
func (w *World) DoString(s, name string, args int) (err error) {
	fn, err := w.l.Load(strings.NewReader(s), name)
	if err != nil {
		return err
	}
	w.l.Insert(fn, -args-1)
	return w.l.PCall(args, lua.MultRet, nil)
}

// DoFile executes the contents of the file at fileName as Lua. args is the
// number of arguments currently on the stack that should be passed in. The file
// is marked as actively running, and is unmarked when the file returns.
func (w *World) DoFile(fileName string, args int) error {
	fi, err := w.FS.Stat(fileName)
	if err != nil {
		return err
	}
	if err = w.PushFile(FileEntry{fileName, fi}); err != nil {
		return err
	}

	fn, err := w.l.LoadFile(fileName)
	if err != nil {
		w.PopFile()
		return err
	}
	w.l.Insert(fn, -args-1)
	err = w.l.PCall(args, lua.MultRet, nil)
	w.PopFile()
	return err
}

// RootDir returns the directory of the first file pushed onto the running file
// stack. Returns an empty string if there are no files on the stack, or the
// absolute path of the file could not be determined.
func (w *World) RootDir() string {
	return w.rootdir
}

// TempDir returns a directory used for temporary files, which is unique per
// world. Returns an empty string if a temporary directory could not be found.
func (w *World) TempDir() string {
	// Create directory lazily.
	w.tmponce.Do(func() {
		if tmp, err := os.MkdirTemp("", "rbxmk_"); err == nil {
			w.tmpdir = tmp
			w.FS.AddRoot(tmp)
		}
	})
	return w.tmpdir
}

// DoFile executes the contents of file f as Lua. args is the number of
// arguments currently on the stack that should be passed in. The file is marked
// as actively running, and is unmarked when the file returns.
func (w *World) DoFileHandle(f fs.File, name string, args int) error {
	if f == nil {
		return fmt.Errorf("expected non-nil file handle")
	}
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	if err = w.PushFile(FileEntry{name, fi}); err != nil {
		return err
	}

	fn, err := w.l.Load(f, fi.Name())
	if err != nil {
		w.PopFile()
		return err
	}
	w.l.Insert(fn, -args-1)
	err = w.l.PCall(args, lua.MultRet, nil)
	w.PopFile()
	return err
}
