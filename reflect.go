package rbxmk

import (
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

// Reflector defines reflection behavior for a type. It defines how to convert a
// types.Value between a Lua value, and behaviors when the type is a userdata.
// It also defines functions for constructing values of the type. A Reflector
// can be registered with a World.
type Reflector struct {
	// Name is the name of the type.
	Name string

	Flags ReflectorFlags

	// Count indicates the number of Lua values that the type can reflect to and
	// from. A Count of 0 is the same as 1. Less than 0 indicates a variable
	// amount.
	Count int

	// PushTo converts v to a number of Lua values. l must be used only for the
	// conversion of values as needed. If err is nil, then lvs must have a
	// length of 1 or greater.
	PushTo func(s State, v types.Value) (lvs []lua.LValue, err error)

	// PullFrom converts a Lua value to v. l must be used only for the
	// conversion of values as needed. lvs must have a length of 1 or greater.
	PullFrom func(s State, lvs ...lua.LValue) (v types.Value, err error)

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

	// Environment is called after the type is registered to provide additional
	// setup. env is the table representing the base library.
	Environment func(s State, env *lua.LTable)

	// ConvertFrom receives an arbitrary value and attempts to convert it to the
	// reflector's type. Returns nil if the value could not be converted.
	ConvertFrom func(v types.Value) types.Value
}

type ReflectorFlags uint8

const (
	_      ReflectorFlags = (1 << iota) / 2
	Exprim                // Whether the type is an explicit primitive.
)

// ValueCount returns the normalized number of Lua values that the type reflects
// between. Less than 0 means the amount is variable.
func (r Reflector) ValueCount() int {
	switch {
	case r.Count == 0, r.Count == 1:
		return 1
	case r.Count < 0:
		return -1
	}
	return r.Count
}

// Metatable defines the metamethods of a custom type.
type Metatable map[string]Metamethod

// Metamethod is called when a metamethod is invoked.
type Metamethod func(s State) int

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
	Get func(s State, v types.Value) int
	// Set gets a value from l and sets it to a member of v. The index and value
	// are the 2nd and 3rd arguments in s.L.
	Set func(s State, v types.Value)
	// Method marks the member as being a method.
	Method bool
}

// Constructors is a set of constructor functions keyed by name.
type Constructors map[string]Constructor

// Constructor creates a new value of a Reflector. The function can receive
// arguments from s.L, and must push a new value to s.L.
type Constructor func(s State) int

// State contains references to an environment surrounding a value.
type State struct {
	*World

	L *lua.LState

	// cycle is used to mark a table as having been traversed. This is non-nil
	// only for types that can contain other types.
	cycle map[interface{}]struct{}
}

// CycleGuard begins a guard against reference cycles when reflecting with the
// state. Returns false if a guard was already set up for the state. If true is
// returned, the guard must be cleared via CycleClear. For example:
//
//     if s.CycleGuard() {
//         defer s.CycleClear()
//     }
//
func (s *State) CycleGuard() bool {
	if s.cycle == nil {
		s.cycle = make(map[interface{}]struct{}, 4)
		return true
	}
	return false
}

// CycleClear clears the cycle guard on the state. Panics if the state has no
// guard.
func (s *State) CycleClear() {
	if s.cycle == nil {
		panic("state has no cycle guard")
	}
	s.cycle = nil
}

// CycleMark marks t as visited, and returns whether t was already visited.
// Panics if the state has no guard.
func (s State) CycleMark(t interface{}) bool {
	if s.cycle == nil {
		panic("attempt to mark reference without cycle guard")
	}
	_, ok := s.cycle[t]
	if !ok {
		s.cycle[t] = struct{}{}
	}
	return ok
}

// Count returns the number of arguments in the stack frame.
func (s State) Count() int {
	return s.L.GetTop()
}

// Push reflects v according to its type as registered with s.World, then pushes
// the results to s.L.
func (s State) Push(v types.Value) int {
	rfl := s.Reflector(v.Type())
	if rfl.Name == "" {
		panic("unregistered type " + v.Type())
	}
	lvs, err := rfl.PushTo(s, v)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	for _, lv := range lvs {
		s.L.Push(lv)
	}
	return len(lvs)
}

// Pull gets from s.L the values starting from n, and reflects a value from them
// according to type t registered with s.World.
func (s State) Pull(n int, t string) types.Value {
	rfl := s.Reflector(t)
	var v types.Value
	var err error
	if rfl.Count < 0 {
		lvs := make([]lua.LValue, 0, 4)
		for i := n; i <= s.L.GetTop(); i++ {
			lvs = append(lvs, s.L.Get(i))
		}
		v, err = rfl.PullFrom(s, lvs...)
	} else if rfl.Count > 1 {
		lvs := make([]lua.LValue, 0, 4)
		for i := n; i <= rfl.Count; i++ {
			lvs = append(lvs, s.L.CheckAny(i))
		}
		v, err = rfl.PullFrom(s, lvs...)
	} else {
		v, err = rfl.PullFrom(s, s.L.CheckAny(n))
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
func (s State) PullOpt(n int, t string, d types.Value) types.Value {
	rfl := s.Reflector(t)
	if rfl.Count < 0 {
		panic("PullOpt cannot pull variable types")
	} else if rfl.Count > 1 {
		panic("PullOpt cannot pull multi-value types")
	}
	lv := s.L.Get(n)
	if lv == lua.LNil {
		return d
	}
	v, err := rfl.PullFrom(s, lv)
	if err != nil {
		s.L.ArgError(n, err.Error())
		return d
	}
	return v
}

// listTypes returns each type listed in a natural sentence.
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
func (s State) PullAnyOf(n int, t ...string) types.Value {
	if n > s.L.GetTop() {
		// Every type must reflect at least one value, so no values is an
		// immediate error.
		s.L.ArgError(n, "value expected")
		return nil
	}
	// Find the maximum count among the given types. 0 is treated the same as 1.
	// <0 indicates an arbitrary number of values.
	max := 1
	ts := make([]Reflector, 0, 4)
	for _, t := range t {
		rfl := s.Reflector(t)
		ts = append(ts, rfl)
		if rfl.Count > 1 {
			max = rfl.Count
		} else if rfl.Count < 0 {
			max = -1
			break
		}
	}
	switch max {
	case 1: // All types have 1 value.
		v := s.L.CheckAny(n)
		for _, t := range ts {
			if v, err := t.PullFrom(s, v); err == nil {
				return v
			}
		}
	case -1: // At least one type has arbitrary values.
		lvs := make([]lua.LValue, 0, 4)
		for _, t := range ts {
			lvs = lvs[:0]
			var v types.Value
			var err error
			if t.Count < 0 {
				// Append all values.
				for i := n; i <= s.L.GetTop(); i++ {
					lvs = append(lvs, s.L.Get(i))
				}
				v, err = t.PullFrom(s, lvs...)
			} else if t.Count > 1 {
				// Append up to type count.
				for i := n; i <= t.Count; i++ {
					lvs = append(lvs, s.L.CheckAny(i))
				}
				v, err = t.PullFrom(s, lvs...)
			} else {
				// Append single value.
				v, err = t.PullFrom(s, s.L.CheckAny(n))
			}
			if err != nil {
				continue
			}
			return v
		}
	default: // Constant maximum.
		lvs := make([]lua.LValue, 0, 4)
		for _, t := range ts {
			lvs = lvs[:0]
			n := t.Count
			if n == 0 {
				n = 1
			}
			for i := n; i <= t.Count; i++ {
				lvs = append(lvs, s.L.CheckAny(i))
			}
			v, err := t.PullFrom(s, lvs...)
			if err != nil {
				continue
			}
			return v
		}
	}
	TypeError(s.L, n, listTypes(t))
	return nil
}

// PushToTable reflects v according to its type as registered with s.World, then
// sets the result to table[field]. The type must be single-value. Does nothing
// if v is nil.
func (s State) PushToTable(table *lua.LTable, field lua.LValue, v types.Value) {
	if v == nil {
		return
	}
	rfl := s.Reflector(v.Type())
	if rfl.Name == "" {
		panic("unregistered type " + v.Type())
	}
	if rfl.Count < 0 {
		panic("PushToTable cannot push variable types")
	} else if rfl.Count > 1 {
		panic("PushToTable cannot push multi-value types")
	}
	lvs, err := rfl.PushTo(s, v)
	if err != nil {
		s.RaiseError("field %s: %s", field, err.Error())
		return
	}
	table.RawSet(field, lvs[0])
}

// PullFromTable gets a value from table[field], and reflects a value from it to
// type t registered with s.World.
func (s State) PullFromTable(table *lua.LTable, field lua.LValue, t string) types.Value {
	rfl := s.Reflector(t)
	if rfl.Count < 0 {
		panic("PullFromTable cannot push variable types")
	} else if rfl.Count > 1 {
		panic("PullFromTable cannot push multi-value types")
	}
	v, err := rfl.PullFrom(s, table.RawGet(field))
	if err != nil {
		s.RaiseError("field %s: %s", field, err.Error())
		return nil
	}
	return v
}

// PullFromTableOpt gets a value from table[field], and reflects a value from it
// to type t registered with s.World. If the value is nil, d is returned
// instead.
func (s State) PullFromTableOpt(table *lua.LTable, field lua.LValue, t string, d types.Value) types.Value {
	rfl := s.Reflector(t)
	if rfl.Count < 0 {
		panic("PullFromTableOpt cannot pull variable types")
	} else if rfl.Count > 1 {
		panic("PullFromTableOpt cannot pull multi-value types")
	}
	lv := table.RawGet(field)
	if lv == lua.LNil {
		return d
	}
	v, err := rfl.PullFrom(s, lv)
	if err != nil {
		s.RaiseError("field %s: %s", field, err.Error())
		return d
	}
	return v
}

// PushArrayOf pushes an rtypes.Array, ensuring that each element is reflected
// according to t.
func (s State) PushArrayOf(t string, v rtypes.Array) int {
	if s.CycleGuard() {
		defer s.CycleClear()
	}
	if s.CycleMark(&v) {
		return s.RaiseError("arrays cannot be cyclic")
	}
	rfl := s.Reflector(t)
	table := s.L.CreateTable(len(v), 0)
	for i, v := range v {
		lv, err := rfl.PushTo(s, v)
		if err != nil {
			return s.RaiseError("%s", err)
		}
		table.RawSetInt(i+1, lv[0])
	}
	s.L.Push(table)
	return 1
}

// PullArrayOf pulls an rtypes.Array from n, ensuring that each element is
// reflected according to t.
func (s State) PullArrayOf(n int, t string) rtypes.Array {
	rfl := s.Reflector(t)
	lv := s.L.CheckAny(n)
	if s.CycleGuard() {
		defer s.CycleClear()
	}
	table, ok := lv.(*lua.LTable)
	if !ok {
		s.L.ArgError(n, TypeError(nil, 0, "table").Error())
		return nil
	}
	if s.CycleMark(table) {
		s.L.ArgError(n, "tables cannot be cyclic")
		return nil
	}
	l := table.Len()
	array := make(rtypes.Array, l)
	for i := 1; i <= l; i++ {
		var err error
		if array[i-1], err = rfl.PullFrom(s, table.RawGetInt(i)); err != nil {
			s.L.ArgError(n, err.Error())
			return nil
		}
	}
	return array
}

func (s State) PushDictionaryOf(n int, t string, v rtypes.Dictionary) int {
	if s.CycleGuard() {
		defer s.CycleClear()
	}
	if s.CycleMark(&v) {
		return s.RaiseError("dictionaries cannot be cyclic")
	}
	rfl := s.Reflector(t)
	table := s.L.CreateTable(0, len(v))
	for k, v := range v {
		lv, err := rfl.PushTo(s, v)
		if err != nil {
			return s.RaiseError("%s", err)
		}
		table.RawSetString(k, lv[0])
	}
	s.L.Push(table)
	return 1
}

func (s State) PullDictionaryOf(n int, t string) rtypes.Dictionary {
	rfl := s.Reflector(t)
	lv := s.L.CheckAny(n)
	if s.CycleGuard() {
		defer s.CycleClear()
	}
	table, ok := lv.(*lua.LTable)
	if !ok {
		s.L.ArgError(n, TypeError(nil, 0, "table").Error())
		return nil
	}
	if s.CycleMark(table) {
		s.L.ArgError(n, "tables cannot be cyclic")
		return nil
	}
	dict := make(rtypes.Dictionary)
	var err error
	table.ForEach(func(k, lv lua.LValue) {
		if err != nil {
			return
		}
		var v types.Value
		if v, err = rfl.PullFrom(s, lv); err == nil {
			dict[k.String()] = v
		}
	})
	if err != nil {
		s.L.ArgError(n, err.Error())
		return nil
	}
	return dict
}

// RaiseError is a shortcut for LState.RaiseError that returns 0.
func (s State) RaiseError(format string, args ...interface{}) int {
	s.L.RaiseError(format, args...)
	return 0
}

// CheckString is like lua.LState.CheckString, except that it does not try to
// convert non-string values into a string.
func (s State) CheckString(n int) string {
	v := s.L.Get(n)
	if lv, ok := v.(lua.LString); ok {
		return string(lv)
	}
	s.L.TypeError(n, lua.LTString)
	return ""
}

// PushTypeTo returns a Reflector.PushTo that converts v to a userdata set with
// a type metatable registered as type t.
func PushTypeTo(t string) func(s State, v types.Value) (lvs []lua.LValue, err error) {
	return func(s State, v types.Value) (lvs []lua.LValue, err error) {
		u := s.UserDataOf(v, t)
		return append(lvs, u), nil
	}
}

// PullTypeFrom returns a Reflector.PullFrom that converts v from a userdata set
// with a type metatable registered as type t.
func PullTypeFrom(t string) func(s State, lvs ...lua.LValue) (v types.Value, err error) {
	return func(s State, lvs ...lua.LValue) (v types.Value, err error) {
		u, ok := lvs[0].(*lua.LUserData)
		if !ok {
			return nil, TypeError(nil, 0, t)
		}
		if u.Metatable != s.L.GetTypeMetatable(t) {
			return nil, TypeError(nil, 0, t)
		}
		if v, ok = u.Value.(types.Value); !ok {
			return nil, TypeError(nil, 0, t)
		}
		return v, nil
	}
}

// typeError is an error where a type was received where another was expected.
type typeError struct {
	expected string
	got      string
}

// Error implements the error interface.
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
