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
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/anaminus/rbxmk/sfs"
	"github.com/robloxapi/types"
)

// udptr holds an untracked reference to a userdata.
type udptr struct {
	p           uintptr
	resurrected bool
}

// World contains the entire state of a Lua environment, including a Lua state,
// and registered Reflectors, Formats, and Sources.
type World struct {
	l          *lua.LState
	fileStack  []FileEntry
	rootdir    string
	libraries  map[string]Library
	reflectors map[string]Reflector
	formats    map[string]Format
	enums      *rtypes.Enums
	enumDumps  map[string]dump.Enum

	rtypes.Global

	tmponce sync.Once
	tmpdir  string

	Client *Client
	FS     sfs.FS

	udmut    sync.Mutex
	userdata map[interface{}]*udptr
}

// NewWorld returns a World initialized with the given Lua state.
func NewWorld(l *lua.LState) *World {
	return &World{
		l:      l,
		Client: NewClient(nil),
	}
}

// UserDataOf returns the userdata value associated with v. If there is no such
// userdata, then a new one is created, with the metatable set to the type
// corresponding to t.
//
// v must be comparable.
func (w *World) UserDataOf(v types.Value, t string) *lua.LUserData {
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
	//
	// This method is safe for the register-based calling convention, since the
	// produced userdata is never a function argument while it is converted
	// between an unsafe pointer.

	w.udmut.Lock()
	defer w.udmut.Unlock()

	if p, ok := w.userdata[v]; ok {
		u := (*lua.LUserData)(unsafe.Pointer(p.p))
		// GC may have finalized u during this time, so tell the finalizer that
		// u has been resurrected.
		p.resurrected = true
		return u
	}

	mt := w.l.GetTypeMetatable(t)
	if mt == lua.LNil {
		panic("expected metatable for type " + t)
	}
	u := w.LuaState().NewUserData(v)
	w.l.SetMetatable(u, mt)

	if w.userdata == nil {
		w.userdata = map[interface{}]*udptr{}
	}
	w.userdata[v] = &udptr{p: uintptr(unsafe.Pointer(u))}
	runtime.SetFinalizer(u, w.finalize)
	return u
}

// finalize is the finalizer for a userdata cached by the world.
func (w *World) finalize(u *lua.LUserData) {
	w.udmut.Lock()
	defer w.udmut.Unlock()
	v := u.Value()
	if p := w.userdata[v]; p.resurrected {
		// u was resurrected while the GC finalized it; reset the finalizer.
		p.resurrected = false
		runtime.SetFinalizer(u, w.finalize)
		return
	}
	delete(w.userdata, v)
}

// UserDataCacheLen return the number of userdata values in the cache.
func (w *World) UserDataCacheLen() int {
	w.udmut.Lock()
	defer w.udmut.Unlock()
	return len(w.userdata)
}

// Library represents a Lua library.
type Library struct {
	// Name is a name that identifies the library.
	Name string
	// ImportedAs is the name that the library is imported as. Empty indicates
	// that the contents of the library are merged into the global environment.
	ImportedAs string
	// Priority indicates the order in which the library is loaded in relation
	// to other libraries.
	Priority int
	// Open returns a table with the contents of the library.
	Open func(s State) *lua.LTable
	// Dump returns a description of the library's API.
	Dump func(s State) dump.Library
	// Types returns a list of type reflector expected by the library. Before
	// opening the library, each reflector is registered.
	Types []func() Reflector
}

// Libraries is a list of Library values that can be sorted by Priority, then
// Name.
type Libraries []Library

func (l Libraries) Len() int      { return len(l) }
func (l Libraries) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l Libraries) Less(i, j int) bool {
	if l[i].Priority == l[j].Priority {
		return l[i].Name < l[j].Name
	}
	return l[i].Priority < l[j].Priority
}

// Open opens lib, then merges the result into the world's global table using
// the ImportedAs field as the name. If the name is already present and is a
// table, then the table of lib is merged into it, preferring the values of lib.
// If the name is an empty string, then lib is merged into the global table
// itself.
//
// The library is registered with the world under the Name of the library. An
// error is returned if the merged value is not a table, or if the name was
// already registered.
func (w *World) Open(lib Library) error {
	if _, ok := w.libraries[lib.Name]; ok {
		return fmt.Errorf("library %q already registered", lib.Name)
	}
	if w.libraries == nil {
		w.libraries = map[string]Library{}
	}
	w.libraries[lib.Name] = lib

	if lib.Types != nil {
		for _, r := range lib.Types {
			w.RegisterReflector(r())
		}
	}
	src := lib.Open(w.State())
	return w.MergeTables(w.l.G.Global, src, lib.ImportedAs)
}

// Library returns the Library registered with the given name. If the name is
// not registered, then Library.Name will be an empty string.
func (w *World) Library(name string) Library {
	return w.libraries[name]
}

// Libraries returns a list of registered libraries.
func (w *World) Libraries() Libraries {
	libraries := make(Libraries, len(w.libraries))
	for _, library := range w.libraries {
		libraries = append(libraries, library)
	}
	sort.Sort(libraries)
	return libraries
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
	// Validate properties.
	for _, property := range r.Properties {
		if property.Get == nil {
			panic("property must define Get field")
		}
	}
	// Validate symbols.
	for _, symbol := range r.Symbols {
		if symbol.Get == nil {
			panic("symbol must define Get field")
		}
	}
	// Validate methods.
	for _, method := range r.Methods {
		if method.Func == nil {
			panic("method must define Func field")
		}
	}

	mt = w.l.CreateTable(0, 8)

	// Unconditional fields.
	mt.RawSetString("__type", lua.LString(r.Name))
	mt.RawSetString("__metatable", lua.LString("the metatable is locked"))
	mt.RawSetString("__tostring", w.l.NewFunction(func(l *lua.LState) int {
		l.Push(lua.LString(r.Name))
		return 1
	}))

	var customIndex Metamethod
	var customNewindex Metamethod
	if len(r.Metatable) > 0 {
		// Set each defined metamethod, overriding predefined values.
		for name, method := range r.Metatable {
			m := method
			mt.RawSetString(name, w.WrapOperator(m))
		}
		// If available, remember index and newindex for member indexing.
		customIndex = r.Metatable["__index"]
		customNewindex = r.Metatable["__newindex"]
	}

	// Setup member getting and setting.
	switch {
	case len(r.Properties)+len(r.Methods) > 0 && len(r.Symbols) > 0:
		// Indexed by both string and symbol.
		mt.RawSetString("__index", w.WrapOperator(func(s State) int {
			v, err := r.PullFrom(s.Context(), s.CheckAny(1))
			if err != nil {
				return s.ArgError(1, err.Error())
			}
			switch index := s.PullAnyOfOpt(2, nil, rtypes.T_String, rtypes.T_Symbol).(type) {
			case types.String:
				if method, ok := r.Methods[string(index)]; ok {
					s.L.Push(s.WrapMethod(func(s State) int {
						v, err := r.PullFrom(s.Context(), s.CheckAny(1))
						if err != nil {
							return s.ArgError(1, err.Error())
						}
						return method.Func(s, v)
					}))
					return 1
				}
				if property, ok := r.Properties[string(index)]; ok {
					return property.Get(s, v)
				}
				if customIndex != nil {
					return customIndex(s)
				}
				return s.RaiseError("%q is not a valid member of %s", index, r.Name)
			case rtypes.Symbol:
				if property, ok := r.Symbols[index]; ok {
					return property.Get(s, v)
				}
				if customIndex != nil {
					return customIndex(s)
				}
				return s.RaiseError("symbol %s is not a valid member of %s", index.Name, r.Name)
			default:
				if customIndex != nil {
					return customIndex(s)
				}
				return s.RaiseError("string or symbol expected for member index of %s, got %s", r.Name, s.Typeof(2))
			}
		}))
		mt.RawSetString("__newindex", w.WrapOperator(func(s State) int {
			v, err := r.PullFrom(s.Context(), s.CheckAny(1))
			if err != nil {
				s.ArgError(1, err.Error())
			}
			switch index := s.PullAnyOfOpt(2, nil, rtypes.T_String, rtypes.T_Symbol).(type) {
			case types.String:
				if _, ok := r.Methods[string(index)]; ok {
					return s.RaiseError("%s of %s cannot be assigned to", index, r.Name)
				}
				if property, ok := r.Properties[string(index)]; ok {
					if property.Set == nil {
						return s.RaiseError("%s of %s cannot be assigned to", index, r.Name)
					}
					property.Set(s, v)
					return 0
				}
				if customNewindex != nil {
					return customNewindex(s)
				}
				return s.RaiseError("%q is not a valid member of %s", index, r.Name)
			case rtypes.Symbol:
				if property, ok := r.Symbols[index]; ok {
					if property.Set == nil {
						return s.RaiseError("symbol %s of %s cannot be assigned to", index.Name, r.Name)
					}
					property.Set(s, v)
					return 0
				}
				if customNewindex != nil {
					return customNewindex(s)
				}
				return s.RaiseError("symbol %s is not a valid member of %s", index.Name, r.Name)
			default:
				if customNewindex != nil {
					return customNewindex(s)
				}
				return s.RaiseError("string or symbol expected for member index of %s, got %s", r.Name, s.Typeof(2))
			}
		}))
	case len(r.Properties)+len(r.Methods) > 0:
		// Indexed only by string.
		mt.RawSetString("__index", w.WrapOperator(func(s State) int {
			v, err := r.PullFrom(s.Context(), s.CheckAny(1))
			if err != nil {
				return s.ArgError(1, err.Error())
			}
			if name, ok := s.L.Get(2).(lua.LString); ok {
				if method, ok := r.Methods[string(name)]; ok {
					s.L.Push(s.WrapMethod(func(s State) int {
						v, err := r.PullFrom(s.Context(), s.CheckAny(1))
						if err != nil {
							return s.ArgError(1, err.Error())
						}
						return method.Func(s, v)
					}))
					return 1
				}
				if property, ok := r.Properties[string(name)]; ok {
					return property.Get(s, v)
				}
				if customIndex != nil {
					return customIndex(s)
				}
				return s.RaiseError("%q is not a valid member of %s", name, r.Name)
			}
			if customIndex != nil {
				return customIndex(s)
			}
			return s.RaiseError("string expected for member name of %s, got %s", r.Name, s.Typeof(2))
		}))
		mt.RawSetString("__newindex", w.WrapOperator(func(s State) int {
			v, err := r.PullFrom(s.Context(), s.CheckAny(1))
			if err != nil {
				s.ArgError(1, err.Error())
			}
			if name, ok := s.L.Get(2).(lua.LString); ok {
				if _, ok := r.Methods[string(name)]; ok {
					return s.RaiseError("%s of %s cannot be assigned to", name, r.Name)
				}
				if property, ok := r.Properties[string(name)]; ok {
					if property.Set == nil {
						return s.RaiseError("%s of %s cannot be assigned to", name, r.Name)
					}
					property.Set(s, v)
					return 0
				}
				if customNewindex != nil {
					return customNewindex(s)
				}
				return s.RaiseError("%q is not a valid member of %s", name, r.Name)
			}
			if customNewindex != nil {
				return customNewindex(s)
			}
			return s.RaiseError("string expected for member name of %s, got %s", r.Name, s.Typeof(2))
		}))
	case len(r.Symbols) > 0:
		// Indexed only by symbol.
		mt.RawSetString("__index", w.WrapOperator(func(s State) int {
			v, err := r.PullFrom(s.Context(), s.CheckAny(1))
			if err != nil {
				return s.ArgError(1, err.Error())
			}
			if symbol, ok := s.PullOpt(2, nil, "Symbol").(rtypes.Symbol); ok {
				if property, ok := r.Symbols[symbol]; ok {
					return property.Get(s, v)
				}
				if customIndex != nil {
					return customIndex(s)
				}
				return s.RaiseError("symbol %s is not a valid member of %s", symbol.Name, r.Name)
			}
			if customIndex != nil {
				return customIndex(s)
			}
			return s.RaiseError("symbol expected for member index of %s, got %s", r.Name, s.Typeof(2))
		}))
		mt.RawSetString("__newindex", w.WrapOperator(func(s State) int {
			v, err := r.PullFrom(s.Context(), s.CheckAny(1))
			if err != nil {
				s.ArgError(1, err.Error())
			}
			if symbol, ok := s.PullOpt(2, nil, "Symbol").(rtypes.Symbol); ok {
				if property, ok := r.Symbols[symbol]; ok {
					if property.Set == nil {
						return s.RaiseError("symbol %s of %s cannot be assigned to", symbol.Name, r.Name)
					}
					property.Set(s, v)
					return 0
				}
				if customNewindex != nil {
					return customNewindex(s)
				}
				return s.RaiseError("symbol %s is not a valid member of %s", symbol.Name, r.Name)
			}
			if customNewindex != nil {
				return customNewindex(s)
			}
			return s.RaiseError("symbol expected for member index of %s, got %s", r.Name, s.Typeof(2))
		}))
	default:
		// Not indexed.
		mt.RawSetString("__index", w.WrapOperator(func(s State) int {
			if customIndex != nil {
				return customIndex(s)
			}
			return s.RaiseError("attempt to index %s with %q", r.Name, s.Typeof(2))
		}))
		mt.RawSetString("__newindex", w.WrapOperator(func(s State) int {
			if customNewindex != nil {
				return customNewindex(s)
			}
			return s.RaiseError("attempt to index %s with %q", r.Name, s.Typeof(2))
		}))
	}

	return mt
}

// RegisterReflector registers a reflector. If the reflector produces a
// metatable, then it is added as a type metatable to the world's state. Panics
// if the reflector is already registered.
func (w *World) RegisterReflector(r Reflector) {
	if _, ok := w.reflectors[r.Name]; ok {
		return
	}
	if w.reflectors == nil {
		w.reflectors = map[string]Reflector{}
	}
	w.reflectors[r.Name] = r

	if mt := w.createTypeMetatable(r); mt != nil {
		w.l.SetField(w.l.Get(lua.RegistryIndex), r.Name, mt)
	}

	if n := len(r.Constructors); n > 0 {
		ctors := w.l.CreateTable(0, n)
		for name, ctor := range r.Constructors {
			if c := ctor.Func; c != nil {
				ctors.RawSetString(name, w.WrapFunc(func(s State) int {
					return c(s)
				}))
			}
		}
		w.l.G.Global.RawSetString(r.Name, ctors)
	}

	for name, def := range r.Enums {
		enum := def()
		items := make([]rtypes.NewItem, 0, len(enum.Items))
		for name, item := range enum.Items {
			items = append(items, rtypes.NewItem{
				Name:  name,
				Value: item.Value,
			})
		}
		sort.Slice(items, func(i, j int) bool {
			if items[i].Value == items[j].Value {
				return items[i].Name < items[j].Name
			}
			return items[i].Value < items[j].Value
		})
		w.RegisterEnum(rtypes.NewEnum(name, items...), &enum)
	}

	if r.Environment != nil {
		r.Environment(w.State())
	}

	for _, t := range r.Types {
		w.RegisterReflector(t())
	}
}

// Reflector returns the Reflector registered with the given name. If the name
// is not registered, then Reflector.Name will be an empty string.
func (w *World) Reflector(name string) Reflector {
	return w.reflectors[name]
}

// MustReflector returns the Reflector registered with the given name. If the
// name is not registered, then MustReflector panics.
func (w *World) MustReflector(name string) Reflector {
	rfl, ok := w.reflectors[name]
	if !ok {
		panic("unregistered type " + name)
	}
	return rfl
}

// PusherOf returns the PushTo field for the Reflector registered as name.
// Returns an error if the name is not registered, or if the reflector does not
// define PushTo.
func (w *World) PusherOf(name string) (p Pusher, err error) {
	rfl, ok := w.reflectors[name]
	if !ok {
		return nil, fmt.Errorf("unknown type %q", name)
	}
	if rfl.PushTo == nil {
		return nil, fmt.Errorf("cannot cast type %s to Lua", name)
	}
	return rfl.PushTo, nil
}

// PullerOf returns the PullFrom field for the Reflector registered as name.
// Returns an error if the name is not registered, or if the reflector does not
// define PullFrom.
func (w *World) PullerOf(name string) (p Puller, err error) {
	rfl, ok := w.reflectors[name]
	if !ok {
		return nil, fmt.Errorf("unknown type %q", name)
	}
	if rfl.PullFrom == nil {
		return nil, fmt.Errorf("cannot cast type %s from Lua", name)
	}
	return rfl.PullFrom, nil
}

// SetterOf returns the SetTo field for the Reflector registered as name.
// Returns an error if the name is not registered, or if the reflector does not
// define SetTo.
func (w *World) SetterOf(name string) (s Setter, err error) {
	rfl, ok := w.reflectors[name]
	if !ok {
		return nil, fmt.Errorf("unknown type %q", name)
	}
	if rfl.SetTo == nil {
		return nil, fmt.Errorf("cannot set type %s", name)
	}
	return rfl.SetTo, nil
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
	if f.Types != nil {
		for _, r := range f.Types {
			w.RegisterReflector(r())
		}
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

// RegisterEnums registers a number of Enum values. Panics if multiple different
// enums are registered with the same name.
func (w *World) RegisterEnums(enums ...*rtypes.Enum) {
	if w.enums == nil {
		w.enums = rtypes.NewEnums()
	}
	for _, enum := range enums {
		if e := w.enums.Enum(enum.Name()); e != nil && e != enum {
			panic("enum " + enum.Name() + " already registered")
		}
	}
	w.enums.Include(enums...)
}

// RegisterEnum registers a single Enum value with an optional dump description.
// Panics if multiple different enums are registered with the same name.
func (w *World) RegisterEnum(enum *rtypes.Enum, desc *dump.Enum) {
	if w.enums == nil {
		w.enums = rtypes.NewEnums()
	}
	if e := w.enums.Enum(enum.Name()); e != nil && e != enum {
		panic("enum " + enum.Name() + " already registered")
	}
	w.enums.Include(enum)
	if desc == nil {
		return
	}
	if w.enumDumps == nil {
		w.enumDumps = map[string]dump.Enum{}
	}
	w.enumDumps[enum.Name()] = *desc
}

// Enum returns the Enum registered with the given name. If the name is not
// registered, then nil is returned.
func (w *World) Enum(name string) *rtypes.Enum {
	if w.enums == nil {
		return nil
	}
	return w.enums.Enum(name)
}

// MustEnum returns the Enum registered with the given name. If the name is not
// registered, then MustEnum panics.
func (w *World) MustEnum(name string) *rtypes.Enum {
	if w.enums != nil {
		if enum := w.enums.Enum(name); enum != nil {
			return enum
		}
	}
	panic("unregistered enum " + name)
}

// Enums returns a list of registered enums.
func (w *World) Enums() *rtypes.Enums {
	if w.enums == nil {
		w.enums = rtypes.NewEnums()
	}
	return w.enums
}

// EnumDump returns a description associated with the enum registered under
// name. Returns nil if the name is not registered, or the enum does not have a
// description.
func (w *World) EnumDump(name string) *dump.Enum {
	enum, ok := w.enumDumps[name]
	if !ok {
		return nil
	}
	return &enum
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
func (w *World) Expand(path string) (s string, err error) {
	s = os.Expand(path, func(v string) string {
		switch v {
		case "script_name", "sn":
			if entry, ok := w.PeekFile(); ok {
				if entry.Path == "" {
					return ""
				}
				var path string
				path, err = filepath.Abs(entry.Path)
				if err != nil {
					err = fmt.Errorf("expand %s: %w", v, err)
				}
				return filepath.Base(path)
			}
		case "script_directory", "script_dir", "sd":
			if entry, ok := w.PeekFile(); ok {
				if entry.Path == "" {
					var dir string
					dir, err = os.Getwd()
					return dir
				}
				var path string
				path, err = filepath.Abs(entry.Path)
				if err != nil {
					err = fmt.Errorf("expand %s: %w", v, err)
				}
				return filepath.Dir(path)
			}
		case "root_script_directory", "root_script_dir", "rsd":
			rootdir := w.RootDir()
			if rootdir == "" {
				rootdir, err = os.Getwd()
				if err != nil {
					err = fmt.Errorf("expand %s: %w", v, err)
				}
			}
			return rootdir
		case "working_directory", "working_dir", "wd":
			var wd string
			wd, err = os.Getwd()
			if err != nil {
				err = fmt.Errorf("expand %s: %w", v, err)
			}
			return wd
		case "temp_directory", "temp_dir", "tmp":
			t := w.TempDir()
			if t == "" {
				err = fmt.Errorf("expand %s: could not find temporary directory", v)
			}
			return t
		}
		err = fmt.Errorf("unknown variable %q", v)
		return ""
	})
	return s, err
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

// LuaState returns the underlying Lua state.
func (w *World) LuaState() *lua.LState {
	return w.l
}

// State returns a State derived from the World.
func (w *World) State() State {
	return State{World: w, L: w.l}
}

// Context returns a Context derived from the World.
func (w *World) Context() Context {
	return Context{World: w, l: w.l}
}

// WrapFunc wraps a function that receives a State into a Lua function.
func (w *World) WrapFunc(f func(State) int) *lua.LFunction {
	return w.l.NewFunction(func(l *lua.LState) int {
		return f(State{World: w, L: l})
	})
}

// WrapMethod is like WrapFunc, but marks the state as being a method.
func (w *World) WrapMethod(f func(State) int) *lua.LFunction {
	return w.l.NewFunction(func(l *lua.LState) int {
		return f(State{World: w, L: l, FrameType: MethodFrame})
	})
}

// WrapOperator is like WrapFunc, but marks the state as being an operator.
func (w *World) WrapOperator(f func(State) int) *lua.LFunction {
	return w.l.NewFunction(func(l *lua.LState) int {
		return f(State{World: w, L: l, FrameType: OperatorFrame})
	})
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
	var fi fs.FileInfo
	var err error
	if len(w.fileStack) == 0 {
		fi, err = os.Stat(fileName)
	} else {
		fi, err = w.FS.Stat(fileName)
	}
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

// SetEnumGlobal sets the "Enum" global to the generated enum types from
// w.Global.Desc. Does nothing if Desc is nil, the Enums type is not registered,
// or if the Enums reflector returns an error.
func (w *World) SetEnumGlobal() {
	if w.Desc == nil {
		return
	}
	rfl := w.Reflector(rtypes.T_Enums)
	if rfl.Name == "" {
		return
	}
	w.Desc.GenerateEnumTypes()
	state := w.State()
	enums, err := rfl.PushTo(state.Context(), w.Desc.EnumTypes)
	if err != nil {
		return
	}
	state.L.SetGlobal("Enum", enums)
}
