package rbxmk

import (
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

// FrameType indicates the kind of frame for a State.
type FrameType uint8

const (
	// Frame is a regular function.
	Function FrameType = iota
	// Frame is a method; exclude first argument.
	Method
	// Frame is an operator, avoid displaying arguments.
	Operator
)

// State contains references to an environment surrounding a value.
type State struct {
	*World

	L *lua.LState

	// FrameType provides a hint to how errors should be produced.
	FrameType FrameType

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
	rfl := s.MustReflector(v.Type())
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
	rfl := s.MustReflector(t)
	var v types.Value
	var err error
	if rfl.Count < 0 {
		lvs := make([]lua.LValue, 0, 4)
		for i := n; i <= s.Count(); i++ {
			lvs = append(lvs, s.L.Get(i))
		}
		v, err = rfl.PullFrom(s, lvs...)
	} else if rfl.Count > 1 {
		lvs := make([]lua.LValue, 0, 4)
		for i := n; i <= rfl.Count; i++ {
			lvs = append(lvs, s.CheckAny(i))
		}
		v, err = rfl.PullFrom(s, lvs...)
	} else {
		v, err = rfl.PullFrom(s, s.CheckAny(n))
	}
	if err != nil {
		s.ArgError(n, err.Error())
		return nil
	}
	return v
}

// PullOpt gets from s.L the value at n, and reflects a value from it according
// to type t registered with s.World. If the value is nil, d is returned
// instead.
func (s State) PullOpt(n int, t string, d types.Value) types.Value {
	rfl := s.MustReflector(t)
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
		s.ArgError(n, err.Error())
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

// ReflectorError raises an error indicating that a reflector pushed or pulled
// an unexpected type. Under normal circumstances, this error should be
// unreachable.
func (s State) ReflectorError(n int) int {
	return s.ArgError(n, "unreachable error: reflector mismatch")
}

// PullAnyOf gets from s.L the values starting from n, and reflects a value from
// them according to any of the types in t registered with s.World. Returns the
// first successful reflection among the types in t. If no types succeeded, then
// a type error is thrown.
func (s State) PullAnyOf(n int, t ...string) types.Value {
	if n > s.Count() {
		// Every type must reflect at least one value, so no values is an
		// immediate error.
		s.ArgError(n, "value expected")
		return nil
	}
	// Find the maximum count among the given types. 0 is treated the same as 1.
	// <0 indicates an arbitrary number of values.
	max := 1
	ts := make([]Reflector, 0, 4)
	for _, t := range t {
		rfl := s.MustReflector(t)
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
		v := s.CheckAny(n)
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
				for i := n; i <= s.Count(); i++ {
					lvs = append(lvs, s.L.Get(i))
				}
				v, err = t.PullFrom(s, lvs...)
			} else if t.Count > 1 {
				// Append up to type count.
				for i := n; i <= t.Count; i++ {
					lvs = append(lvs, s.CheckAny(i))
				}
				v, err = t.PullFrom(s, lvs...)
			} else {
				// Append single value.
				v, err = t.PullFrom(s, s.CheckAny(n))
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
				lvs = append(lvs, s.CheckAny(i))
			}
			v, err := t.PullFrom(s, lvs...)
			if err != nil {
				continue
			}
			return v
		}
	}
	s.TypeError(n, listTypes(t), "")
	return nil
}

// PushToTable reflects v according to its type as registered with s.World, then
// sets the result to table[field]. The type must be single-value. Does nothing
// if v is nil.
func (s State) PushToTable(table *lua.LTable, field lua.LValue, v types.Value) {
	if v == nil {
		return
	}
	rfl := s.MustReflector(v.Type())
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
	rfl := s.MustReflector(t)
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
	rfl := s.MustReflector(t)
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
	rfl := s.MustReflector(t)
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
	rfl := s.MustReflector(t)
	lv := s.CheckAny(n)
	if s.CycleGuard() {
		defer s.CycleClear()
	}
	table, ok := lv.(*lua.LTable)
	if !ok {
		s.ArgError(n, TypeError("table", lv.Type().String()).Error())
		return nil
	}
	if s.CycleMark(table) {
		s.ArgError(n, "tables cannot be cyclic")
		return nil
	}
	l := table.Len()
	array := make(rtypes.Array, l)
	for i := 1; i <= l; i++ {
		var err error
		if array[i-1], err = rfl.PullFrom(s, table.RawGetInt(i)); err != nil {
			s.ArgError(n, err.Error())
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
	rfl := s.MustReflector(t)
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
	rfl := s.MustReflector(t)
	lv := s.CheckAny(n)
	if s.CycleGuard() {
		defer s.CycleClear()
	}
	table, ok := lv.(*lua.LTable)
	if !ok {
		s.ArgError(n, TypeError("table", lv.Type().String()).Error())
		return nil
	}
	if s.CycleMark(table) {
		s.ArgError(n, "tables cannot be cyclic")
		return nil
	}
	dict := make(rtypes.Dictionary)
	err := table.ForEach(func(k, lv lua.LValue) error {
		v, err := rfl.PullFrom(s, lv)
		if err != nil {
			return err
		}
		dict[k.String()] = v
		return nil
	})
	if err != nil {
		s.ArgError(n, err.Error())
		return nil
	}
	return dict
}

// RaiseError is a shortcut for LState.RaiseError that returns 0.
func (s State) RaiseError(format string, args ...interface{}) int {
	s.L.RaiseError(format, args...)
	return 0
}

// ArgError raises an argument error depending on the state's frame type.
func (s State) ArgError(n int, msg string) int {
	switch s.FrameType {
	case Method:
		if n <= 1 {
			s.RaiseError("bad method receiver: %s", msg)
		} else {
			s.L.ArgError(n-1, msg)
		}
	case Operator:
		s.RaiseError(msg)
	default:
		s.L.ArgError(n, msg)
	}
	return 0
}

// TypeError raises an argument type error depending on the state's frame type.
func (s State) TypeError(n int, want, got string) int {
	err := TypeError(want, got)
	switch s.FrameType {
	case Method:
		if n <= 1 {
			s.RaiseError("bad method receiver: %s", err)
		} else {
			s.L.ArgError(n-1, err.Error())
		}
	case Operator:
		s.RaiseError("%s", err.Error())
	default:
		s.L.ArgError(n, err.Error())
	}
	return 0
}

// CheckAny returns the nth argument, which can be any type as long as the
// argument exists.
func (s State) CheckAny(n int) lua.LValue {
	if n > s.Count() {
		s.ArgError(n, "value expected")
		return nil
	}
	return s.L.Get(n)
}

// CheckBool returns the nth argument, expecting a boolean.
func (s State) CheckBool(n int) bool {
	v := s.L.Get(n)
	if lv, ok := v.(lua.LBool); ok {
		return bool(lv)
	}
	s.TypeError(n, lua.LTBool.String(), v.Type().String())
	return false
}

// CheckInt returns the nth argument as an int, expecting a number.
func (s State) CheckInt(n int) int {
	v := s.L.Get(n)
	if lv, ok := v.(lua.LNumber); ok {
		return int(lv)
	}
	s.TypeError(n, lua.LTNumber.String(), v.Type().String())
	return 0
}

// CheckInt64 returns the nth argument as an int64, expecting a number.
func (s State) CheckInt64(n int) int64 {
	v := s.L.Get(n)
	if lv, ok := v.(lua.LNumber); ok {
		return int64(lv)
	}
	s.TypeError(n, lua.LTNumber.String(), v.Type().String())
	return 0
}

// CheckNumber returns the nth argument, expecting a number.
func (s State) CheckNumber(n int) lua.LNumber {
	v := s.L.Get(n)
	if lv, ok := v.(lua.LNumber); ok {
		return lv
	}
	s.TypeError(n, lua.LTNumber.String(), v.Type().String())
	return 0
}

// CheckString returns the nth argument, expecting a string. Unlike
// LState.CheckString, it does not try to convert non-string values into a
// string.
func (s State) CheckString(n int) string {
	v := s.L.Get(n)
	if lv, ok := v.(lua.LString); ok {
		return string(lv)
	}
	s.TypeError(n, lua.LTString.String(), v.Type().String())
	return ""
}

// CheckTable returns the nth argument, expecting a table.
func (s State) CheckTable(n int) *lua.LTable {
	v := s.L.Get(n)
	if lv, ok := v.(*lua.LTable); ok {
		return lv
	}
	s.TypeError(n, lua.LTTable.String(), v.Type().String())
	return nil
}

// CheckFunction returns the nth argument, expecting a function.
func (s State) CheckFunction(n int) *lua.LFunction {
	v := s.L.Get(n)
	if lv, ok := v.(*lua.LFunction); ok {
		return lv
	}
	s.TypeError(n, lua.LTFunction.String(), v.Type().String())
	return nil
}

// CheckUserData returns the nth argument, expecting a userdata.
func (s State) CheckUserData(n int) *lua.LUserData {
	v := s.L.Get(n)
	if lv, ok := v.(*lua.LUserData); ok {
		return lv
	}
	s.TypeError(n, lua.LTUserData.String(), v.Type().String())
	return nil
}

// CheckThread returns the nth argument, expecting a thread.
func (s State) CheckThread(n int) *lua.LState {
	v := s.L.Get(n)
	if lv, ok := v.(*lua.LState); ok {
		return lv
	}
	s.TypeError(n, lua.LTThread.String(), v.Type().String())
	return nil
}

// OptBool returns the nth argument as a bool, or d if the argument is nil.
func (s State) OptBool(n int, d bool) bool {
	v := s.L.Get(n)
	if v == lua.LNil {
		return d
	}
	if lv, ok := v.(lua.LBool); ok {
		return bool(lv)
	}
	s.TypeError(n, lua.LTBool.String(), v.Type().String())
	return false
}

// OptInt returns the nth argument as an int, or d if the argument is nil.
func (s State) OptInt(n int, d int) int {
	v := s.L.Get(n)
	if v == lua.LNil {
		return d
	}
	if lv, ok := v.(lua.LNumber); ok {
		return int(lv)
	}
	s.TypeError(n, lua.LTNumber.String(), v.Type().String())
	return 0
}

// OptInt64 returns the nth argument as an int64, or d if the argument is nil.
func (s State) OptInt64(n int, d int64) int64 {
	v := s.L.Get(n)
	if v == lua.LNil {
		return d
	}
	if lv, ok := v.(lua.LNumber); ok {
		return int64(lv)
	}
	s.TypeError(n, lua.LTNumber.String(), v.Type().String())
	return 0
}

// OptNumber returns the nth argument as a number, or d if the argument is nil.
func (s State) OptNumber(n int, d lua.LNumber) lua.LNumber {
	v := s.L.Get(n)
	if v == lua.LNil {
		return d
	}
	if lv, ok := v.(lua.LNumber); ok {
		return lv
	}
	s.TypeError(n, lua.LTNumber.String(), v.Type().String())
	return 0
}

// OptString returns the nth argument as a string, or d if the argument is nil.
func (s State) OptString(n int, d string) string {
	v := s.L.Get(n)
	if v == lua.LNil {
		return d
	}
	if lv, ok := v.(lua.LString); ok {
		return string(lv)
	}
	s.TypeError(n, lua.LTString.String(), v.Type().String())
	return ""
}

// OptTable returns the nth argument as a table, or d if the argument is nil.
func (s State) OptTable(n int, d *lua.LTable) *lua.LTable {
	v := s.L.Get(n)
	if v == lua.LNil {
		return d
	}
	if lv, ok := v.(*lua.LTable); ok {
		return lv
	}
	s.TypeError(n, lua.LTTable.String(), v.Type().String())
	return nil
}

// OptFunction returns the nth argument as a function, or d if the argument is
// nil.
func (s State) OptFunction(n int, d *lua.LFunction) *lua.LFunction {
	v := s.L.Get(n)
	if v == lua.LNil {
		return d
	}
	if lv, ok := v.(*lua.LFunction); ok {
		return lv
	}
	s.TypeError(n, lua.LTFunction.String(), v.Type().String())
	return nil
}

// OptUserData returns the nth argument as a userdata, or d if the argument is
// nil.
func (s State) OptUserData(n int, d *lua.LUserData) *lua.LUserData {
	v := s.L.Get(n)
	if v == lua.LNil {
		return d
	}
	if lv, ok := v.(*lua.LUserData); ok {
		return lv
	}
	s.TypeError(n, lua.LTUserData.String(), v.Type().String())
	return nil
}
