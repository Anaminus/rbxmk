package rbxmk

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk/dump"
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

	// Properties defines the properties of a custom type. If the __index and
	// __newindex metamethods are not defined by Metatable, then Properties
	// defines them according to the given properties. In case of name
	// conflicts, methods are prioritized over properties.
	Properties Properties

	// Symbols defines the symbols of a custom type. If the __index and
	// __newindex metamethods are not defined by Metatable, then Symbols defines
	// them according to the given properties.
	Symbols Symbols

	// Methods defines the methods of a custom type. If the __index and
	// __newindex metamethods are not defined by Metatable, then Methods defines
	// them according to the given members. In case of name conflicts, methods
	// are prioritized over properties.
	Methods Methods

	// Constructors defines functions that construct the type. If non-nil, a
	// table containing each constructor is created and set as a global
	// referenced by Name.
	Constructors Constructors

	// Environment is called after the type is registered to provide additional
	// setup.
	Environment func(s State)

	// ConvertFrom receives an arbitrary value and attempts to convert it to the
	// reflector's type. Returns nil if the value could not be converted.
	ConvertFrom func(v types.Value) types.Value

	// Types is a list of additional type reflectors that this reflector depends
	// on.
	Types []func() Reflector

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
	for name, property := range r.Properties {
		if property.Dump == nil {
			continue
		}
		if _, ok := def.Properties[name]; !ok {
			if def.Properties == nil {
				def.Properties = dump.Properties{}
			}
			def.Properties[name] = property.Dump()
		}
	}
	for symbol, property := range r.Symbols {
		if property.Dump == nil {
			continue
		}
		if _, ok := def.Symbols[symbol.Name]; !ok {
			if def.Symbols == nil {
				def.Symbols = dump.Symbols{}
			}
			def.Symbols[symbol.Name] = property.Dump()
		}
	}
	for name, method := range r.Methods {
		if method.Dump == nil {
			continue
		}
		if _, ok := def.Methods[name]; !ok {
			if def.Methods == nil {
				def.Methods = dump.Methods{}
			}
			def.Methods[name] = method.Dump()
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

// Properties is a set of properties keyed by name.
type Properties map[string]Property

// Symbols is a set of properties keyed by symbol.
type Symbols map[rtypes.Symbol]Property

// Property defines a property of a custom type.
type Property struct {
	// Get gets the value of a member from v and pushes it onto s. The index is
	// the 2nd argument in s.
	Get func(s State, v types.Value) int
	// Set gets a value from s and sets it to a member of v. The index and value
	// are the 2nd and 3rd arguments in s. Set is optional, if nil, the property
	// will be treated as read-only.
	Set func(s State, v types.Value)
	// Dump returns a description of the member's API.
	Dump func() dump.Property
}

// Methods is a set of methods keyed by name.
type Methods map[string]Method

// Method defines a member of a custom type.
type Method struct {
	// Func is the body of the method. The first argument will be the same value
	// as v.
	Func func(s State, v types.Value) int
	// Dump returns a description of the member's API.
	Dump func() dump.Function
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
// a type metatable registered as type t. Each push always produces a new
// userdata. This results in better performance, but makes the value unsuitable
// as a table key.
func PushTypeTo(t string) func(s State, v types.Value) (lvs []lua.LValue, err error) {
	return func(s State, v types.Value) (lvs []lua.LValue, err error) {
		u := s.L.NewUserData(v)
		s.L.SetMetatable(u, s.L.GetTypeMetatable(t))
		return append(lvs, u), nil
	}
}

// PushPtrTypeTo returns a Reflector.PushTo that converts v to a userdata set
// with a type metatable registered as type t. The same value will push the same
// userdata, making the value usable as a table key.
func PushPtrTypeTo(t string) func(s State, v types.Value) (lvs []lua.LValue, err error) {
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
			return nil, TypeError{Want: t, Got: lvs[0].Type().String()}
		}
		if u.Metatable != s.L.GetTypeMetatable(t) {
			return nil, TypeError{Want: t, Got: lvs[0].Type().String()}
		}
		if v, ok = u.Value().(types.Value); !ok {
			return nil, TypeError{Want: t, Got: lvs[0].Type().String()}
		}
		return v, nil
	}
}

// typeError is an error where a type was received where another was expected.
type TypeError struct {
	Want string
	Got  string
}

// Error implements the error interface.
func (err TypeError) Error() string {
	if err.Got == "" {
		return err.Want + " expected"
	}
	return err.Want + " expected, got " + err.Got
}
