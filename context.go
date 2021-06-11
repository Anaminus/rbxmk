package rbxmk

import lua "github.com/anaminus/gopher-lua"

// Context provides a context for reflecting values.
type Context struct {
	*World

	L *lua.LState

	// cycle is used to mark a table as having been traversed. This is non-nil
	// only for types that can contain other types.
	cycle map[interface{}]struct{}
}

// Context returns a Context derived from the World.
func (s Context) Context() Context {
	return Context{World: s.World, L: s.L}
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

// ReflectorError raises an error indicating that a reflector pushed or pulled
// an unexpected type. Under normal circumstances, this error should be
// unreachable.
func (s Context) ReflectorError(n int) int {
	panic("unreachable error: reflector mismatch")
}
