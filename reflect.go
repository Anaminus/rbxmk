package rbxmk

import (
	"strings"

	"github.com/robloxapi/rbxfile"
	"github.com/yuin/gopher-lua"
)

// Value is a value.
type Value interface{}

// Copier is a Value that can be copied.
type Copier interface {
	Value
	Copy() Value
}

// TValue is a value annotated with a type.
type TValue struct {
	Type string
	Value
}

type Type struct {
	// Name is the name of the type.
	Name string

	// Count indicates the number of Lua values that the type can reflect to and
	// from. A Count of 0 is the same as 1. Less than 0 indicates a variable
	// amount.
	Count int

	// ReflectTo converts v to a number of Lua values. l must be used only for
	// the conversion of values as needed. If err is nil, then lvs must have a
	// length of 1 or greater.
	ReflectTo func(s State, t Type, v Value) (lvs []lua.LValue, err error)

	// ReflectFrom converts a Lua value to v. l must be used only for the
	// conversion of values as needed. lvs must have a length of 1 or greater.
	ReflectFrom func(s State, t Type, lvs ...lua.LValue) (v Value, err error)

	// Serialize serializes a value.
	Serialize func(s State, v Value) (sv rbxfile.Value, err error)

	// Deserialize deserializes a value.
	Deserialize func(s State, sv rbxfile.Value) (v Value, err error)

	// Metatable defines the metamethods of a custom type. If Metatable is
	// non-nil, then a metatable is constructed and registered as a type
	// metatable under Name.
	Metatable Metatable

	// Members defines the members of a custom type. If the __index and
	// __newindex metamethods are not defined by Metatable, then Members defines
	// them according to the given members.
	Members Members

	// Constructors defines functions that construct the type. If non-nil, a
	// table containing each constructor is created and set as a global
	// referenced by Name.
	Constructors Constructors
}

// ValueCount returns the normalized number of Lua values that the type reflects
// between. Less than 0 means the amount is variable.
func (t Type) ValueCount() int {
	switch {
	case t.Count == 0, t.Count == 1:
		return 1
	case t.Count < 0:
		return -1
	}
	return t.Count
}

// Metatable defines the metamethods of a custom type.
type Metatable map[string]Metamethod

// Metamethod is called when a metamethod is invoked. v is the receiver, which
// was successfully reflected from the first argument to the metamethod.
type Metamethod func(s State, v Value) int

// Members is a set of members keyed by name.
type Members map[string]Member

// Member defines a member of a custom type. There are several kinds of members:
//
//     - Property: Method is false, Get must be defined, Set is optionally
//       defined. If Set is not defined, then the member is marked as read-only.
//     - Method: Method is true, Get must be defined, Set is ignored. Get is the
//       method itself.
type Member struct {
	// Get gets the value of a member from v and pushes it onto l. The index is
	// the 2nd argument in s.L. If Method is true, the function is pushed as a
	// method. The first argument will be the same value as v.
	Get func(s State, v Value) int
	// Set gets a value from l and sets it to a member of v. The index and value
	// are the 2nd and 3rd arguments in s.L.
	Set func(s State, v Value)
	// Method marks the member as being a method.
	Method bool
}

// Constructors is a set of constructor functions keyed by name.
type Constructors map[string]Constructor

// Constructor creates a new value of a Type. The function can receive arguments
// from s.L, and must push a new value to s.L.
type Constructor func(s State) int

// State contains references to an environment surrounding a value.
type State struct {
	*World

	L *lua.LState

	// Cycle is used to mark a table as having been traversed. This is non-nil
	// only for types that can contain other types.
	Cycle *Cycle
}

// Count returns the number of arguments in the stack frame.
func (s State) Count() int {
	return s.L.GetTop()
}

// Push reflects v according to type t registered with s.World, then pushes the
// results to s.L.
func (s State) Push(t string, v Value) int {
	typ := s.Type(t)
	if typ.Name == "" {
		panic("unregistered type " + t)
	}
	lvs, err := typ.ReflectTo(s, typ, v)
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	for _, lv := range lvs {
		s.L.Push(lv)
	}
	return len(lvs)
}

// Pull gets from s.L the values starting from n, and reflects a value from them
// according to type t registered with s.World.
func (s State) Pull(n int, t string) Value {
	typ := s.Type(t)
	var v Value
	var err error
	if typ.Count < 0 {
		lvs := make([]lua.LValue, 0, 4)
		for i := n; i <= s.L.GetTop(); i++ {
			lvs = append(lvs, s.L.Get(i))
		}
		v, err = typ.ReflectFrom(s, typ, lvs...)
	} else if typ.Count > 1 {
		lvs := make([]lua.LValue, 0, 4)
		for i := n; i <= typ.Count; i++ {
			lvs = append(lvs, s.L.CheckAny(i))
		}
		v, err = typ.ReflectFrom(s, typ, lvs...)
	} else {
		v, err = typ.ReflectFrom(s, typ, s.L.CheckAny(n))
	}
	if err != nil {
		s.L.ArgError(n, err.Error())
		return nil
	}
	return v
}

// PullOpt gets from s.L the value at n, and reflects a value from it according
// to type t registered with s.World. If the value is nil, d is returned
// instead.
func (s State) PullOpt(n int, t string, d Value) Value {
	typ := s.Type(t)
	if typ.Count < 0 {
		panic("PullOpt cannot pull variable types")
	} else if typ.Count > 1 {
		panic("PullOpt cannot pull multi-value types")
	}
	lv := s.L.Get(n)
	if lv == lua.LNil {
		return d
	}
	v, err := typ.ReflectFrom(s, typ, lv)
	if err != nil {
		s.L.ArgError(n, err.Error())
		return d
	}
	return v
}

func listTypes(types []string) string {
	switch len(types) {
	case 0:
		return ""
	case 1:
		return types[0]
	case 2:
		return types[0] + " or " + types[1]
	}
	return strings.Join(types[:len(types)-2], ", ") + ", or " + types[len(types)-1]
}

// PullAnyOf gets from s.L the values starting from n, and reflects a value from
// them according to any of the types in t registered with s.World. Returns the
// first successful reflection among the types in t. If no types succeeded, then
// a type error is thrown.
func (s State) PullAnyOf(n int, t ...string) Value {
	if n > s.L.GetTop() {
		// Every type must reflect at least one value, so no values is an
		// immediate error.
		s.L.ArgError(n, "value expected")
		return nil
	}
	// Find the maximum count among the given types. 0 is treated the same as 1.
	// <0 indicates an arbitrary number of values.
	max := 1
	types := make([]Type, 0, 4)
	for _, t := range t {
		typ := s.Type(t)
		types = append(types, typ)
		if typ.Count > 1 {
			max = typ.Count
		} else if typ.Count < 0 {
			max = -1
			break
		}
	}
	switch max {
	case 1: // All types have 1 value.
		v := s.L.CheckAny(n)
		for _, t := range types {
			if v, err := t.ReflectFrom(s, t, v); err == nil {
				return v
			}
		}
	case -1: // At least one type has arbitrary values.
		lvs := make([]lua.LValue, 0, 4)
		for _, t := range types {
			lvs = lvs[:0]
			var v Value
			var err error
			if t.Count < 0 {
				// Append all values.
				for i := n; i <= s.L.GetTop(); i++ {
					lvs = append(lvs, s.L.Get(i))
				}
				v, err = t.ReflectFrom(s, t, lvs...)
			} else if t.Count > 1 {
				// Append up to type count.
				for i := n; i <= t.Count; i++ {
					lvs = append(lvs, s.L.CheckAny(i))
				}
				v, err = t.ReflectFrom(s, t, lvs...)
			} else {
				// Append single value.
				v, err = t.ReflectFrom(s, t, s.L.CheckAny(n))
			}
			if err != nil {
				continue
			}
			return v
		}
	default: // Constant maximum.
		lvs := make([]lua.LValue, 0, 4)
		for _, t := range types {
			lvs = lvs[:0]
			n := t.Count
			if n == 0 {
				n = 1
			}
			for i := n; i <= t.Count; i++ {
				lvs = append(lvs, s.L.CheckAny(i))
			}
			v, err := t.ReflectFrom(s, t, lvs...)
			if err != nil {
				continue
			}
			return v
		}
	}
	TypeError(s.L, n, listTypes(t))
	return nil
}

// Cycle is used to detect cyclic references by containing values that have
// already been traversed.
type Cycle struct {
	m map[interface{}]struct{}
}

// Has returns whether a value has already been visited.
func (c *Cycle) Has(t interface{}) bool {
	if c == nil {
		return false
	}
	_, ok := c.m[t]
	return ok
}

// Put marks a value has having been visited.
func (c *Cycle) Put(t interface{}) {
	if c == nil {
		return
	}
	if c.m == nil {
		c.m = make(map[interface{}]struct{}, 4)
	}
	c.m[t] = struct{}{}
}

// ReflectTypeTo is a Type.ReflectTo that converts v to a userdata set with a
// type metatable registered as t.Name.
func ReflectTypeTo(s State, t Type, v Value) (lvs []lua.LValue, err error) {
	u := s.L.NewUserData()
	u.Value = v
	s.L.SetMetatable(u, s.L.GetTypeMetatable(t.Name))
	return append(lvs, u), nil
}

// ReflectTypeFrom is a Type.ReflectFrom that converts v from a userdata set
// with a type metatable registered as t.Name.
func ReflectTypeFrom(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
	u, ok := lvs[0].(*lua.LUserData)
	if !ok {
		return nil, TypeError(nil, 0, t.Name)
	}
	if u.Metatable != s.L.GetTypeMetatable(t.Name) {
		return nil, TypeError(nil, 0, t.Name)
	}
	if v, ok = u.Value.(Value); !ok {
		return nil, TypeError(nil, 0, t.Name)
	}
	return v, nil
}

type typeError struct {
	expected string
	got      string
}

func (err typeError) Error() string {
	if err.got == "" {
		return err.expected + " expected"
	}
	return err.expected + " expected, got " + err.got
}

// TypeError raises an argument error indicating that a given type was expected.
func TypeError(l *lua.LState, n int, typ string) (err error) {
	if l != nil {
		err = typeError{expected: typ, got: l.Get(n).Type().String()}
		l.ArgError(n, err.Error())
	} else {
		err = typeError{expected: typ}
	}
	return err
}

// CheckType returns the underlying value if the value at n is a userdata that
// has type metatable corresponding to typ.
func CheckType(l *lua.LState, n int, typ string) interface{} {
	if u, ok := l.Get(n).(*lua.LUserData); ok {
		if u.Metatable == l.GetTypeMetatable(typ) {
			return u.Value
		}
	}
	TypeError(l, n, typ)
	return nil
}

// OptType is like CheckType, but returns nil if the value is nil.
func OptType(l *lua.LState, n int, typ string) interface{} {
	v := l.Get(n)
	if v == lua.LNil {
		return nil
	}
	if u, ok := v.(*lua.LUserData); ok {
		if u.Metatable == l.GetTypeMetatable(typ) {
			return u.Value
		}
	}
	TypeError(l, n, typ)
	return nil
}
