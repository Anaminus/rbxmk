package rbxmk

import lua "github.com/anaminus/gopher-lua"

// Context provides a context for reflecting values.
type Context struct {
	*World

	l *lua.LState

	// cycle is used to mark a table as having been traversed. This is non-nil
	// only for types that can contain other types.
	cycle map[interface{}]struct{}
}

// Context returns a Context derived from the World.
func (s Context) Context() Context {
	return Context{World: s.World, l: s.l}
}

// CycleGuard begins a guard against reference cycles when reflecting with the
// state. Returns false if a guard was already set up for the state. If true is
// returned, the guard must be cleared via CycleClear. For example:
//
//     if s.CycleGuard() {
//         defer s.CycleClear()
//     }
//
func (s *Context) CycleGuard() bool {
	if s.cycle == nil {
		s.cycle = make(map[interface{}]struct{}, 4)
		return true
	}
	return false
}

// CycleClear clears the cycle guard on the state. Panics if the state has no
// guard.
func (s *Context) CycleClear() {
	if s.cycle == nil {
		panic("state has no cycle guard")
	}
	s.cycle = nil
}

// CycleMark marks t as visited, and returns whether t was already visited.
// Panics if the state has no guard.
func (s Context) CycleMark(t interface{}) bool {
	if s.cycle == nil {
		panic("attempt to mark reference without cycle guard")
	}
	_, ok := s.cycle[t]
	if !ok {
		s.cycle[t] = struct{}{}
	}
	return ok
}

// CreateTable returns a new table according to the context's LState.
func (s Context) CreateTable(acap, hcap int) *lua.LTable {
	return s.l.CreateTable(acap, hcap)
}

// NewUserData returns a new userdata according to the context's LState.
func (s Context) NewUserData(value interface{}) *lua.LUserData {
	return s.l.NewUserData(value)
}

// GetMetaField gets the value of the event field from the metatable of obj.
// Returns LNil if obj has no metatable, or the metatable has no such field.
func (s Context) GetMetaField(obj lua.LValue, event string) lua.LValue {
	return s.l.GetMetaField(obj, event)
}

// GetTypeMetatable returns the typ metatable registered with the context's
// LState.
func (s Context) GetTypeMetatable(typ string) lua.LValue {
	return s.l.GetTypeMetatable(typ)
}

// SetMetatable sets the metatable of obj to mt according to the context's
// LState.
func (s Context) SetMetatable(obj, mt lua.LValue) {
	s.l.SetMetatable(obj, mt)
}

// ReflectorError raises an error indicating that a reflector pushed or pulled
// an unexpected type. Under normal circumstances, this error should be
// unreachable.
func (s Context) ReflectorError(n int) int {
	panic("unreachable error: reflector mismatch")
}
