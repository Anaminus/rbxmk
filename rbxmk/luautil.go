package main

import (
	"github.com/yuin/gopher-lua"
	"strconv"
)

type TArgs struct {
	l *lua.LState
	*lua.LTable
}

func typeOf(l *lua.LState, v lua.LValue) string {
	if v.Type() == lua.LTUserData {
		if s, ok := l.CallMeta(v, "__type").(lua.LString); ok {
			return string(s)
		}
	}
	return v.Type().String()
}

func GetArgs(l *lua.LState, index int) TArgs {
	tb := l.Get(index)
	if l.GetTop() != index || tb.Type() != lua.LTTable {
		l.RaiseError("function must have 1 table argument")
	} else if l.GetMetatable(tb) != lua.LNil {
		l.RaiseError("table argument cannot have metatable")
	}
	return TArgs{l: l, LTable: tb.(*lua.LTable)}
}

func (t TArgs) ErrorField(name string, expected, got string) {
	if got == "" {
		t.l.RaiseError("bad value at field %q: %s expected", name, expected)
	} else {
		t.l.RaiseError("bad value at field %q: %s expected, got %s", name, expected, got)
	}
}

func (t TArgs) ErrorIndex(index int, expected, got string) {
	if got == "" {
		t.l.RaiseError("bad value at index #%d: %s expected", index, expected)
	} else {
		t.l.RaiseError("bad value at index #%d: %s expected, got %s", index, expected, got)
	}
}

func (t TArgs) TypeOfField(name string) string {
	return typeOf(t.l, t.RawGetString(name))
}

func (t TArgs) TypeOfIndex(index int) string {
	return typeOf(t.l, t.RawGetInt(index))
}

func (t TArgs) FieldString(name string, opt bool) (s string, ok bool) {
	lv := t.RawGetString(name)
	if typ := typeOf(t.l, lv); typ != "string" {
		if opt && typ == "nil" {
			return "", false
		}
		t.ErrorField(name, "string", typ)
	}
	return string(lv.(lua.LString)), true
}

func (t TArgs) IndexString(index int, opt bool) string {
	lv := t.RawGetInt(index)
	if typ := typeOf(t.l, lv); typ != "string" {
		if opt && typ == "nil" {
			return ""
		}
		t.ErrorIndex(index, "string", typ)
	}
	return string(lv.(lua.LString))
}

func (t TArgs) FieldTyped(name string, valueType string, opt bool) (v interface{}) {
	lv := t.RawGetString(name)
	if typ := typeOf(t.l, lv); typ != valueType {
		if opt && typ == "nil" {
			return nil
		}
		t.ErrorField(name, valueType, typ)
	}
	uv, _ := lv.(*lua.LUserData)
	return uv.Value
}

func (t TArgs) IndexTyped(index int, valueType string, opt bool) (v interface{}) {
	lv := t.RawGetInt(index)
	if typ := typeOf(t.l, lv); typ != valueType {
		if opt && typ == "nil" {
			return nil
		}
		t.ErrorIndex(index, valueType, typ)
	}
	uv, _ := lv.(*lua.LUserData)
	return uv.Value
}

func (t TArgs) IndexValue(index int) interface{} {
	return GetLuaValue(t.RawGetInt(index))
}

func (t TArgs) FieldValue(name string) interface{} {
	return GetLuaValue(t.RawGetString(name))
}

// PushAsArgs takes the indices of the table and pushes them to the stack,
// removing the table afterwards.
func (t TArgs) PushAsArgs() {
	t.l.Pop(1)
	nt := t.Len()
	for i := 1; i <= nt; i++ {
		t.l.Push(t.RawGetInt(i))
	}
}

func GetLuaValue(lv lua.LValue) interface{} {
	switch v := lv.(type) {
	case lua.LBool:
		return bool(v)
	case lua.LNumber:
		return float64(v)
	case lua.LString:
		return string(v)
	case *lua.LTable:
		return v
	case *lua.LFunction:
		return v
	case *lua.LState:
		return v
	case *lua.LUserData:
		return v.Value
	}
	return nil
}

// Stand-in for nil that can be used to cause a table entry to be removed.
var ForceNil = new(lua.LUserData)

func ParseLuaValue(s string, forceNil bool) lua.LValue {
	number, err := strconv.ParseFloat(s, 64)
	switch {
	case err == nil:
		return lua.LNumber(number)
	case s == "true":
		return lua.LTrue
	case s == "false":
		return lua.LFalse
	case s == "nil":
		if forceNil {
			return ForceNil
		}
		return lua.LNil
	default:
		return lua.LString(s)
	}
}

func CheckStringVar(s string) bool {
	if !('A' <= s[0] && s[0] <= 'Z' ||
		'a' <= s[0] && s[0] <= 'z' ||
		s[0] == '_') {
		return false
	}
	for i := 1; i < len(s); i++ {
		if !('0' <= s[i] && s[i] <= '9' ||
			'A' <= s[i] && s[i] <= 'Z' ||
			'a' <= s[i] && s[i] <= 'z' ||
			s[i] == '_') {
			return false
		}
	}
	return true
}
