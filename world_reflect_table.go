package rbxmk

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/robloxapi/types"
)

// PushToTable reflects v according to its type as registered with s.World, then
// sets the result to table[field]. The type must be single-value. Does nothing
// if v is nil.
func (w *World) PushToTable(table *lua.LTable, field lua.LValue, v types.Value) (err error) {
	if v == nil {
		return
	}
	push, err := w.PusherOf(v.Type())
	if err != nil {
		return err
	}
	lv, err := push(w.Context(), v)
	if err != nil {
		return fmt.Errorf("field %s: %w", field, err)
	}
	table.RawSet(field, lv)
	return nil
}

// PullFromTable gets a value from table[field], and reflects a value from it to
// type t registered with s.World.
func (w *World) PullFromTable(table *lua.LTable, field lua.LValue, t string) (v types.Value, err error) {
	pull, err := w.PullerOf(t)
	if err != nil {
		return nil, err
	}
	if v, err = pull(w.Context(), table.RawGet(field)); err != nil {
		return nil, fmt.Errorf("field %s: %w", field, err)
	}
	return v, nil
}

// PullFromTableOpt gets a value from table[field], and reflects a value from it
// to type t registered with s.World. If the value is nil, d is returned
// instead.
func (w *World) PullFromTableOpt(table *lua.LTable, field lua.LValue, d types.Value, t string) (v types.Value, err error) {
	pull, err := w.PullerOf(t)
	if err != nil {
		return nil, err
	}
	lv := table.RawGet(field)
	if lv == lua.LNil {
		return d, nil
	}
	if v, err = pull(w.Context(), lv); err != nil {
		return nil, fmt.Errorf("field %s: %w", field, err)
	}
	return v, nil
}

// PullAnyFromTable gets a value from table[field], and reflects a value from it
// according to the first successful type from t registered with s.World.
func (w *World) PullAnyFromTable(table *lua.LTable, field lua.LValue, t ...string) (v types.Value, err error) {
	lv := table.RawGet(field)
	for _, t := range t {
		pull, err := w.PullerOf(t)
		if err != nil {
			return nil, err
		}
		if v, err := pull(w.Context(), lv); err == nil {
			return v, nil
		}
	}
	return nil, fmt.Errorf("field %s: %s expected, got %s", field, listTypes(t), w.Typeof(lv))
}

// PullAnyFromTableOpt gets a value from table[field], and reflects a value from
// it according to the first successful type from t registered with s.World. If
// the field is nil, then d is returned instead.
func (w *World) PullAnyFromTableOpt(table *lua.LTable, field lua.LValue, d types.Value, t ...string) (v types.Value, err error) {
	lv := table.RawGet(field)
	if lv == lua.LNil {
		return d, nil
	}
	for _, t := range t {
		pull, err := w.PullerOf(t)
		if err != nil {
			return nil, err
		}
		if v, err := pull(w.Context(), lv); err == nil {
			return v, nil
		}
	}
	return nil, fmt.Errorf("field %s: %s expected, got %s", field, listTypes(t), w.Typeof(lv))
}

// PushToArray is like PushToTable, but receives an int as the index of the
// table.
func (w *World) PushToArray(table *lua.LTable, index int, v types.Value) (err error) {
	return w.PushToTable(table, lua.LNumber(index), v)
}

// PullFromArray is like PullFromTable, but receives an int as the index of the
// table.
func (w *World) PullFromArray(table *lua.LTable, index int, t string) (v types.Value, err error) {
	return w.PullFromTable(table, lua.LNumber(index), t)
}

// PullFromArrayOpt is like PullFromTableOpt, but receives an int as the index
// of the table.
func (w *World) PullFromArrayOpt(table *lua.LTable, index int, d types.Value, t string) (v types.Value, err error) {
	return w.PullFromTableOpt(table, lua.LNumber(index), d, t)
}

// PullAnyFromArray is like PullAnyFromTable, but receives an int as the index
// of the table.
func (w *World) PullAnyFromArray(table *lua.LTable, index int, t ...string) (v types.Value, err error) {
	return w.PullAnyFromTable(table, lua.LNumber(index), t...)
}

// PullAnyFromArrayOpt is like PullAnyFromTableOpt, but receives an int as the
// index of the table.
func (w *World) PullAnyFromArrayOpt(table *lua.LTable, index int, d types.Value, t ...string) (v types.Value, err error) {
	return w.PullAnyFromTableOpt(table, lua.LNumber(index), v, t...)
}

// PushToDictionary is like PushToTable, but receives a string as the key of the
// table.
func (w *World) PushToDictionary(table *lua.LTable, key string, v types.Value) (err error) {
	return w.PushToTable(table, lua.LString(key), v)
}

// PullFromDictionary is like PullFromTable, but receives a string as the key of
// the table.
func (w *World) PullFromDictionary(table *lua.LTable, key string, t string) (v types.Value, err error) {
	return w.PullFromTable(table, lua.LString(key), t)
}

// PullFromDictionaryOpt is like PullFromTableOpt, but receives a string as the
// key of the table.
func (w *World) PullFromDictionaryOpt(table *lua.LTable, key string, d types.Value, t string) (v types.Value, err error) {
	return w.PullFromTableOpt(table, lua.LString(key), d, t)
}

// PullAnyFromDictionary is like PullAnyFromTable, but receives a string as the
// key of the table.
func (w *World) PullAnyFromDictionary(table *lua.LTable, key string, t ...string) (v types.Value, err error) {
	return w.PullAnyFromTable(table, lua.LString(key), t...)
}

// PullAnyFromDictionaryOpt is like PullAnyFromTableOpt, but receives a string
// as the key of the table.
func (w *World) PullAnyFromDictionaryOpt(table *lua.LTable, key string, d types.Value, t ...string) (v types.Value, err error) {
	return w.PullAnyFromTableOpt(table, lua.LString(key), v, t...)
}
