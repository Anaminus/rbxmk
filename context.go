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
func (c Context) Context() Context {
	return Context{World: c.World, l: c.l}
}

// CycleGuard begins a guard against reference cycles when reflecting with the
// state. Returns false if a guard was already set up for the state. If true is
// returned, the guard must be cleared via CycleClear. For example:
//
//     if c.CycleGuard() {
//         defer c.CycleClear()
//     }
//
func (c *Context) CycleGuard() bool {
	if c.cycle == nil {
		c.cycle = make(map[interface{}]struct{}, 4)
		return true
	}
	return false
}

// CycleClear clears the cycle guard on the state. Panics if the state has no
// guard.
func (c *Context) CycleClear() {
	if c.cycle == nil {
		panic("state has no cycle guard")
	}
	c.cycle = nil
}

// CycleMark marks t as visited, and returns whether t was already visited.
// Panics if the state has no guard.
func (c Context) CycleMark(t interface{}) bool {
	if c.cycle == nil {
		panic("attempt to mark reference without cycle guard")
	}
	_, ok := c.cycle[t]
	if !ok {
		c.cycle[t] = struct{}{}
	}
	return ok
}

// CreateTable returns a new table according to the context's LState.
func (c Context) CreateTable(acap, hcap int) *lua.LTable {
	return c.l.CreateTable(acap, hcap)
}

// NewUserData returns a new userdata according to the context's LState.
func (c Context) NewUserData(value interface{}) *lua.LUserData {
	return c.l.NewUserData(value)
}

// GetMetaField gets the value of the event field from the metatable of obj.
// Returns LNil if obj has no metatable, or the metatable has no such field.
func (c Context) GetMetaField(obj lua.LValue, event string) lua.LValue {
	return c.l.GetMetaField(obj, event)
}

// GetTypeMetatable returns the typ metatable registered with the context's
// LState.
func (c Context) GetTypeMetatable(typ string) lua.LValue {
	return c.l.GetTypeMetatable(typ)
}

// SetMetatable sets the metatable of obj to mt according to the context's
// LState.
func (c Context) SetMetatable(obj, mt lua.LValue) {
	c.l.SetMetatable(obj, mt)
}

// ReflectorError panics, indicating that a reflector pushed or pulled an
// unexpected type. Under normal circumstances, this should be unreachable.
func (c Context) ReflectorError() {
	panic("unreachable error: reflector mismatch")
}
