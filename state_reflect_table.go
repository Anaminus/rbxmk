package rbxmk

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/robloxapi/types"
)

// PushToTable reflects v according to its type as registered with s.World, then
// sets the result to table[field]. The type must be single-value. Does nothing
// if v is nil.
func (s State) PushToTable(table *lua.LTable, field lua.LValue, v types.Value) {
	if err := s.World.PushToTable(table, field, v); err != nil {
		s.RaiseError("field %s: %s", field, err.Error())
	}
}

// PullFromTable gets a value from table[field], and reflects a value from it to
// type t registered with s.World.
func (s State) PullFromTable(table *lua.LTable, field lua.LValue, t string) (v types.Value) {
	v, err := s.World.PullFromTable(table, field, t)
	if err != nil {
		s.RaiseError("field %s: %s", field, err.Error())
		return nil
	}
	return v
}

// PullFromTableOpt gets a value from table[field], and reflects a value from it
// to type t registered with s.World. If the value is nil, d is returned
// instead.
func (s State) PullFromTableOpt(table *lua.LTable, field lua.LValue, d types.Value, t string) (v types.Value) {
	v, err := s.World.PullFromTableOpt(table, field, d, t)
	if err != nil {
		s.RaiseError("field %s: %s", field, err.Error())
		return nil
	}
	return v
}

// PullAnyFromTable gets a value from table[field], and reflects a value from it
// according to the first successful type from t registered with s.World.
func (s State) PullAnyFromTable(table *lua.LTable, field lua.LValue, t ...string) (v types.Value) {
	v, err := s.World.PullAnyFromTable(table, field, t...)
	if err != nil {
		s.RaiseError("field %s: %s", field, err.Error())
		return nil
	}
	return v
}

// PullAnyFromTableOpt gets a value from table[field], and reflects a value from
// it according to the first successful type from t registered with s.World. If
// the field is nil, then d is returned instead.
func (s State) PullAnyFromTableOpt(table *lua.LTable, field lua.LValue, d types.Value, t ...string) (v types.Value) {
	v, err := s.World.PullAnyFromTableOpt(table, field, d, t...)
	if err != nil {
		s.RaiseError("field %s: %s", field, err.Error())
		return nil
	}
	return v
}

// PushToArray is like PushToTable, but receives an int as the index of the
// table.
func (s State) PushToArray(table *lua.LTable, index int, v types.Value) {
	if err := s.World.PushToArray(table, index, v); err != nil {
		s.RaiseError("index %d: %s", index, err.Error())
	}
}

// PullFromArray is like PullFromTable, but receives an int as the index of the
// table.
func (s State) PullFromArray(table *lua.LTable, index int, t string) (v types.Value) {
	v, err := s.World.PullFromArray(table, index, t)
	if err != nil {
		s.RaiseError("index %d: %s", index, err.Error())
		return nil
	}
	return v
}

// PullFromArrayOpt is like PullFromTableOpt, but receives an int as the index
// of the table.
func (s State) PullFromArrayOpt(table *lua.LTable, index int, d types.Value, t string) (v types.Value) {
	v, err := s.World.PullFromArrayOpt(table, index, d, t)
	if err != nil {
		s.RaiseError("index %d: %s", index, err.Error())
		return nil
	}
	return v
}

// PullAnyFromArray is like PullAnyFromTable, but receives an int as the index
// of the table.
func (s State) PullAnyFromArray(table *lua.LTable, index int, t ...string) (v types.Value) {
	v, err := s.World.PullAnyFromArray(table, index, t...)
	if err != nil {
		s.RaiseError("index %d: %s", index, err.Error())
		return nil
	}
	return v
}

// PullAnyFromArrayOpt is like PullAnyFromTableOpt, but receives an int as the
// index of the table.
func (s State) PullAnyFromArrayOpt(table *lua.LTable, index int, d types.Value, t ...string) (v types.Value) {
	v, err := s.World.PullAnyFromArrayOpt(table, index, v)
	if err != nil {
		s.RaiseError("index %d: %s", index, err.Error())
		return nil
	}
	return v
}

// PushToDictionary is like PushToTable, but receives a string as the key of the
// table.
func (s State) PushToDictionary(table *lua.LTable, key string, v types.Value) {
	if err := s.World.PushToDictionary(table, key, v); err != nil {
		s.RaiseError("key %s: %s", key, err.Error())
	}
}

// PullFromDictionary is like PullFromTable, but receives a string as the key of
// the table.
func (s State) PullFromDictionary(table *lua.LTable, key string, t string) (v types.Value) {
	v, err := s.World.PullFromDictionary(table, key, t)
	if err != nil {
		s.RaiseError("key %s: %s", key, err.Error())
		return nil
	}
	return v
}

// PullFromDictionaryOpt is like PullFromTableOpt, but receives a string as the
// key of the table.
func (s State) PullFromDictionaryOpt(table *lua.LTable, key string, d types.Value, t string) (v types.Value) {
	v, err := s.World.PullFromDictionaryOpt(table, key, d, t)
	if err != nil {
		s.RaiseError("key %s: %s", key, err.Error())
		return nil
	}
	return v
}

// PullAnyFromDictionary is like PullAnyFromTable, but receives a string as the
// key of the table.
func (s State) PullAnyFromDictionary(table *lua.LTable, key string, t ...string) (v types.Value) {
	v, err := s.World.PullAnyFromDictionary(table, key, t...)
	if err != nil {
		s.RaiseError("key %s: %s", key, err.Error())
		return nil
	}
	return v
}

// PullAnyFromDictionaryOpt is like PullAnyFromTableOpt, but receives a string
// as the key of the table.
func (s State) PullAnyFromDictionaryOpt(table *lua.LTable, key string, d types.Value, t ...string) (v types.Value) {
	v, err := s.World.PullAnyFromDictionaryOpt(table, key, v)
	if err != nil {
		s.RaiseError("key %s: %s", key, err.Error())
		return nil
	}
	return v
}
