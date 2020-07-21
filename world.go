package rbxmk

import (
	"fmt"
	"os"
	"strings"

	"github.com/yuin/gopher-lua"
)

type World struct {
	l         *lua.LState
	fileStack []FileInfo
	types     map[string]Type
	formats   map[string]Format
	sources   map[string]Source
}

func NewWorld(l *lua.LState) *World {
	return &World{
		l:     l,
		types: map[string]Type{},
	}
}

func (w *World) Open(lib func(s State)) {
	lib(State{World: w, L: w.l})
}

// Type returns the Type registered with the given name. If the name is not
// registered, then Type.Name will be an empty string.
func (w *World) Type(name string) Type {
	return w.types[name]
}

// RegisterType registers a type. Panics if the type is already registered.
func (w *World) RegisterType(t Type) {
	if _, ok := w.types[t.Name]; ok {
		panic("type " + t.Name + " already registered")
	}
	if w.types == nil {
		w.types = map[string]Type{}
	}
	w.types[t.Name] = t

	var mt *lua.LTable
	var index Metamethod
	var newindex Metamethod
	if t.Metatable != nil {
		mt = w.l.CreateTable(0, len(t.Metatable)+3)
		for name, method := range t.Metatable {
			m := method
			mt.RawSetString(name, w.l.NewFunction(func(l *lua.LState) int {
				u := l.CheckUserData(1)
				if u.Metatable != mt {
					TypeError(l, 1, t.Name)
					return 0
				}
				v, ok := u.Value.(Value)
				if !ok {
					TypeError(l, 1, t.Name)
					return 0
				}
				return m(State{World: w, L: l}, v)
			}))
		}
		mt.RawSetString("__type", lua.LString(t.Name))
		if t.Metatable["__tostring"] == nil {
			mt.RawSetString("__tostring", w.l.NewFunction(func(l *lua.LState) int {
				l.Push(lua.LString(t.Name))
				return 1
			}))
		}
		if t.Metatable["__metatable"] == nil {
			mt.RawSetString("__metatable", lua.LString("the metatable is locked"))
		}
		index = t.Metatable["__index"]
		newindex = t.Metatable["__newindex"]
	}
	if t.Members != nil {
		for _, member := range t.Members {
			if member.Get == nil {
				panic("member must define Get function")
			}
		}
		if mt == nil {
			mt = w.l.CreateTable(0, 2)
		}
		mt.RawSetString("__index", w.l.NewFunction(func(l *lua.LState) int {
			u := l.CheckUserData(1)
			if u.Metatable != mt {
				TypeError(l, 1, t.Name)
				return 0
			}
			v, ok := u.Value.(Value)
			if !ok {
				TypeError(l, 1, t.Name)
				return 0
			}
			idx := l.Get(2)
			name, ok := idx.(lua.LString)
			if ok {
				member, ok := t.Members[string(name)]
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
				return index(State{World: w, L: l}, v)
			}
			if ok {
				l.RaiseError("%q is not a valid member of %s", name, t.Name)
			} else {
				l.ArgError(2, "string expected, got "+idx.Type().String())
			}
			return 0
		}))
		mt.RawSetString("__newindex", w.l.NewFunction(func(l *lua.LState) int {
			u := l.CheckUserData(1)
			if u.Metatable != mt {
				TypeError(l, 1, t.Name)
				return 0
			}
			v, ok := u.Value.(Value)
			if !ok {
				TypeError(l, 1, t.Name)
				return 0
			}
			idx := l.Get(2)
			name, ok := idx.(lua.LString)
			if ok {
				member, ok := t.Members[string(name)]
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
				return newindex(State{World: w, L: l}, v)
			}
			if ok {
				l.RaiseError("%q is not a valid member of %s", name, t.Name)
			} else {
				l.ArgError(2, "string expected, got "+idx.Type().String())
			}
			return 0
		}))
	}
	if mt != nil {
		w.l.SetField(w.l.Get(lua.RegistryIndex), t.Name, mt)
	}
	if t.Constructors != nil {
		ctors := w.l.CreateTable(0, len(t.Constructors))
		for name, ctor := range t.Constructors {
			ctors.RawSetString(name, w.l.NewFunction(func(l *lua.LState) int {
				return ctor(State{World: w, L: w.l})
			}))
		}
		w.l.SetGlobal(t.Name, ctors)
	}
	if t.Environment != nil {
		t.Environment(State{World: w, L: w.l})
	}
}

// Format returns the Format registered with the given name. If the name is not
// registered, then Format.Name will be an empty string.
func (w *World) Format(name string) Format {
	return w.formats[strings.TrimPrefix(name, ".")]
}

// RegisterFormat registers a format. Panics if the format is already
// registered.
func (w *World) RegisterFormat(f Format) {
	f.Name = strings.TrimPrefix(f.Name, ".")
	if _, ok := w.formats[f.Name]; ok {
		panic("format " + f.Name + " already registered")
	}
	if w.formats == nil {
		w.formats = map[string]Format{}
	}
	w.formats[f.Name] = f
}

// Source returns the Source registered with the given name. If the name is not
// registered, then Source.Name will be an empty string.
func (w *World) Source(name string) Source {
	return w.sources[name]
}

// RegisterSource registers a source. Panics if the source is already
// registered.
func (w *World) RegisterSource(s Source) {
	if _, ok := w.sources[s.Name]; ok {
		panic("source " + s.Name + " already registered")
	}
	if w.sources == nil {
		w.sources = map[string]Source{}
	}
	w.sources[s.Name] = s
}

// State returns the underlying Lua state.
func (w *World) State() *lua.LState {
	return w.l
}

func (w *World) WrapFunc(f func(State) int) *lua.LFunction {
	return w.l.NewFunction(func(l *lua.LState) int {
		return f(State{World: w, L: l})
	})
}

// ReflectTo reflects v to lvs using registered type t.
func (w *World) ReflectTo(t string, v Value) (lvs []lua.LValue, err error) {
	typ := w.types[t]
	if typ.Name == "" {
		return nil, fmt.Errorf("unknown type %q", t)
	}
	if typ.ReflectTo == nil {
		return nil, fmt.Errorf("cannot cast type %q to Lua", t)
	}
	return typ.ReflectTo(State{World: w, L: w.l}, typ, v)
}

// ReflectFrom reflects lvs to v using registered type t.
func (w *World) ReflectFrom(t string, lvs ...lua.LValue) (v Value, err error) {
	typ := w.types[t]
	if typ.Name == "" {
		return nil, fmt.Errorf("unknown type %q", t)
	}
	if typ.ReflectFrom == nil {
		return nil, fmt.Errorf("cannot cast type %q from Lua", t)
	}
	return typ.ReflectFrom(State{World: w, L: w.l}, typ, lvs...)
}

// Push reflects v according to registered type t, then pushes the results to
// the world's state.
func (w *World) Push(t string, v Value) int {
	return State{World: w, L: w.l}.Push(t, v)
}

// Pull gets from the world's Lua state the values starting from n, and reflects
// a value from them according to registered type t.
func (w *World) Pull(n int, t string) Value {
	return State{World: w, L: w.l}.Pull(n, t)
}

// PullOpt gets from the world's Lua state the value at n, and reflects a value
// from it according to registered type t. If the value is nil, d is returned
// instead.
func (w *World) PullOpt(n int, t string, d Value) Value {
	return State{World: w, L: w.l}.PullOpt(n, t, d)
}

// PullAnyOf gets from the world's Lua state the values starting from n, and
// reflects a value from them according to any of the registered types in t.
// Returns the first successful reflection among the types in t. If no types
// succeeded, then a type error is thrown.
func (w *World) PullAnyOf(n int, t ...string) Value {
	return State{World: w, L: w.l}.PullAnyOf(n, t...)
}

type FileInfo struct {
	Path string
	os.FileInfo
}

// PushFile marks a file as the currently running file.
func (w *World) PushFile(fi FileInfo) error {
	for _, f := range w.fileStack {
		if os.SameFile(fi.FileInfo, f.FileInfo) {
			return fmt.Errorf("\"%s\" is already running", fi.Path)
		}
	}
	w.fileStack = append(w.fileStack, fi)
	return nil
}

// PopFile unmarks the currently running file.
func (w *World) PopFile() {
	if len(w.fileStack) > 0 {
		w.fileStack[len(w.fileStack)-1] = FileInfo{}
		w.fileStack = w.fileStack[:len(w.fileStack)-1]
	}
}

// PeekFile returns the info of the currently running file. Returns false if
// there is no running file.
func (w *World) PeekFile() (fi FileInfo, ok bool) {
	if len(w.fileStack) == 0 {
		return
	}
	fi = w.fileStack[len(w.fileStack)-1]
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
	fi, err := os.Stat(fileName)
	if err != nil {
		return err
	}
	if err = w.PushFile(FileInfo{fileName, fi}); err != nil {
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

type File interface {
	Name() string
	Stat() (os.FileInfo, error)
	Read([]byte) (int, error)
	Write([]byte) (int, error)
}

// DoFile executes the contents of file f as Lua. args is the number of
// arguments currently on the stack that should be passed in. The file is marked
// as actively running, and is unmarked when the file returns.
func (w *World) DoFileHandle(f File, args int) error {
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	if err = w.PushFile(FileInfo{f.Name(), fi}); err != nil {
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
