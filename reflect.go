package rbxmk

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk/dump"
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

	// Dump returns an additional description of the API of the reflector's
	// type. Member and constructor APIs should be described by their respective
	// fields.
	Dump func() dump.TypeDef
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

// DumpAll returns a full description of the API of the reflector's type by
// merging the result of Dump, Members, and Constructors.
func (r Reflector) DumpAll() dump.TypeDef {
	var def dump.TypeDef
	if r.Dump != nil {
		def = r.Dump()
	}
	for name, member := range r.Members {
		if member.Dump == nil {
			continue
		}
		switch v := member.Dump().(type) {
		case dump.Property:
			if _, ok := def.Properties[name]; !ok {
				if def.Properties == nil {
					def.Properties = dump.Properties{}
				}
				def.Properties[name] = v
			}
		case dump.Function:
			if _, ok := def.Methods[name]; !ok {
				if def.Methods == nil {
					def.Methods = dump.Methods{}
				}
				def.Methods[name] = v
			}
		}
	}
	for name, ctor := range r.Constructors {
		if ctor.Dump == nil {
			continue
		}
		funcs := append(def.Constructors[name], ctor.Dump()...)
		if len(funcs) > 0 {
			if def.Constructors == nil {
				def.Constructors = dump.Constructors{}
			}
			def.Constructors[name] = funcs
		}
	}
	return def
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
	// Dump returns a description of the member's API.
	Dump func() dump.Value
}

// Constructors is a set of constructor functions keyed by name.
type Constructors map[string]Constructor

// Constructor creates a new value of a Reflector. The function can receive
// arguments from s.L, and must push a new value to s.L.
type Constructor struct {
	Func func(s State) int
	// Dump returns a description of constructor's API. Each function describes
	// one possible signature of the constructor.
	Dump func() dump.MultiFunction
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
