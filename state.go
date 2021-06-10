package rbxmk

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

// FrameType indicates the kind of frame for a State.
type FrameType uint8

const (
	// Frame is a regular function.
	FunctionFrame FrameType = iota
	// Frame is a method; exclude first argument.
	MethodFrame
	// Frame is an operator, avoid displaying arguments.
	OperatorFrame
)

// State facilitates the reflection of values to a particular Lua state.
type State struct {
	*World

	L *lua.LState

	// FrameType provides a hint to how errors should be produced.
	FrameType FrameType
}

// Context returns a Context derived from the State.
func (s State) Context() Context {
	return Context{State: s}
}

// Count returns the number of arguments in the stack frame.
func (s State) Count() int {
	return s.L.GetTop()
}

// ReflectorError raises an error indicating that a reflector pushed or pulled
// an unexpected type. Under normal circumstances, this error should be
// unreachable.
func (s State) ReflectorError(n int) int {
	return s.ArgError(n, "unreachable error: reflector mismatch")
}

// PushTuple pushes each value.
func (s State) PushTuple(values ...types.Value) int {
	lvs := make([]lua.LValue, len(values))
	for i, value := range values {
		rfl := s.MustReflector(value.Type())
		lv, err := rfl.PushTo(s.Context(), value)
		if err != nil {
			return s.RaiseError("%s", err)
		}
		lvs[i] = lv
	}
	for _, lv := range lvs {
		s.L.Push(lv)
	}
	return len(lvs)
}

// PullTuple pulls each value starting from n as a Variant.
func (s State) PullTuple(n int) rtypes.Tuple {
	c := s.Count()
	length := c - n + 1
	if length <= 0 {
		return nil
	}
	rfl := s.MustReflector("Variant")
	vs := make(rtypes.Tuple, length)
	for i := n; i <= c; i++ {
		lv := s.L.Get(i)
		v, err := rfl.PullFrom(s.Context(), lv)
		if err != nil {
			s.ArgError(i, err.Error())
			return nil
		}
		vs[i-n] = v
	}
	return vs
}

// RaiseError is a shortcut for LState.RaiseError that returns 0.
func (s State) RaiseError(format string, args ...interface{}) int {
	s.L.RaiseError(format, args...)
	return 0
}

// ArgError raises an argument error depending on the state's frame type.
func (s State) ArgError(n int, msg string, v ...interface{}) int {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	switch s.FrameType {
	case MethodFrame:
		if n <= 1 {
			s.RaiseError("bad method receiver: %s", msg)
		} else {
			s.L.ArgError(n-1, msg)
		}
	case OperatorFrame:
		s.RaiseError(msg)
	default:
		s.L.ArgError(n, msg)
	}
	return 0
}

// TypeError raises an argument type error depending on the state's frame type.
func (s State) TypeError(n int, want, got string) int {
	err := TypeError{Want: want, Got: got}
	switch s.FrameType {
	case MethodFrame:
		if n <= 1 {
			s.RaiseError("bad method receiver: %s", err)
		} else {
			s.L.ArgError(n-1, err.Error())
		}
	case OperatorFrame:
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
