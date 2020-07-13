package rbxmk

import (
	"fmt"

	"github.com/robloxapi/rbxfile"
	"github.com/yuin/gopher-lua"
)

type World struct {
	l     *lua.LState
	types map[string]Type
}

func NewWorld(l *lua.LState) *World {
	return &World{
		l:     l,
		types: map[string]Type{},
	}
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
	if t.Metatable != nil {
		mt = w.l.CreateTable(0, len(t.Metatable)+3)
		for name, method := range t.Metatable {
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
				return method(State{World: w, L: l}, v)
			}))
		}
		if t.Metatable["__tostring"] == nil {
			mt.RawSetString("__tostring", w.l.NewFunction(func(l *lua.LState) int {
				l.Push(lua.LString(t.Name))
				return 1
			}))
		}
		if t.Metatable["__index"] != nil || t.Metatable["__newindex"] != nil {
			goto finish
		}
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
			name := l.CheckString(2)
			member, ok := t.Members[name]
			if !ok {
				l.RaiseError("%q is not a valid member of %s", name, t.Name)
				return 0
			}
			if member.Method {
				// Push as method.
				l.Push(l.NewFunction(func(l *lua.LState) int {
					// TODO: validate that s.L.Get(1) matches v, or at least has
					// the expected type.
					return member.Get(State{World: w, L: l}, v)
				}))
				return 1
			} else {
				return member.Get(State{World: w, L: l}, v)
			}
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
			name := l.CheckString(2)
			member, ok := t.Members[name]
			if !ok {
				l.RaiseError("%q is not a valid member of %s", name, t.Name)
				return 0
			}
			if member.Method || member.Set == nil {
				l.RaiseError("%s cannot be assigned to", name)
			}
			member.Set(State{World: w, L: l}, v)
			return 0
		}))
	}
finish:
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
		globals := w.l.Get(lua.GlobalsIndex)
		w.l.SetField(globals, t.Name, ctors)
	}
}

// State returns the underlying Lua state.
func (w *World) State() *lua.LState {
	return w.l
}

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

func (w *World) Serialize(s State, t string, v Value) (sv rbxfile.Value, err error) {
	typ := w.types[t]
	if typ.Name == "" {
		return nil, fmt.Errorf("unknown type %q", t)
	}
	if typ.Serialize == nil {
		return nil, fmt.Errorf("cannot serialize type %q", t)
	}
	return typ.Serialize(State{World: w, L: w.l}, v)
}

func (w *World) Deserialize(s State, t string, sv rbxfile.Value) (v Value, err error) {
	typ := w.types[t]
	if typ.Name == "" {
		return nil, fmt.Errorf("unknown type %q", t)
	}
	if typ.Serialize == nil {
		return nil, fmt.Errorf("cannot deserialize type %q", t)
	}
	return typ.Deserialize(State{World: w, L: w.l}, sv)
}

func (w *World) PushValue(t string, v Value) int {
	return State{World: w, L: w.l}.PushValue(t, v)
}

func (w *World) PullValue(n int, t string) Value {
	return State{World: w, L: w.l}.PullValue(n, t)
}

func (w *World) PullOptValue(n int, t ...string) Value {
	return State{World: w, L: w.l}.PullOptValue(n, t...)
}
