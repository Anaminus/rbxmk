package rbxmk

import (
	"fmt"
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

// Typeof returns the type of the given Lua value. If it is a userdata, Typeof
// attempts to get the type according to the value's metatable. Panics if v is
// nil (not if nil Lua value).
func (w *World) Typeof(v lua.LValue) string {
	if v == nil {
		panic("value expected")
	}
	u, ok := v.(*lua.LUserData)
	if !ok {
		return v.Type().String()
	}
	t, ok := w.l.GetMetaField(u, "__type").(lua.LString)
	if !ok {
		return u.Type().String()
	}
	return string(t)
}

// Push reflects v to lv.
func (w *World) Push(v types.Value) (lv lua.LValue, err error) {
	push, err := w.PusherOf(v.Type())
	if err != nil {
		return nil, err
	}
	return push(w.Context(), v)
}

// Pull reflects lv to v using registered type t.
func (w *World) Pull(lv lua.LValue, t string) (v types.Value, err error) {
	pull, err := w.PullerOf(t)
	if err != nil {
		return nil, err
	}
	return pull(w.Context(), lv)
}

// PullOpt reflects lv to v using registered type t. If lv is nil, then d is
// returned instead.
func (w *World) PullOpt(lv lua.LValue, d types.Value, t string) (v types.Value, err error) {
	if lv == nil || lv.Type() == lua.LTNil {
		return d, nil
	}
	return w.Pull(lv, t)
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

// PullAnyOf reflects lv to the first successful type from t. Returns an error
// if none of the types were successful.
func (w *World) PullAnyOf(lv lua.LValue, t ...string) (v types.Value, err error) {
	for _, t := range t {
		pull, err := w.PullerOf(t)
		if err != nil {
			return nil, err
		}
		if v, err := pull(w.Context(), lv); err == nil {
			return v, nil
		}
	}

	return nil, TypeError{Want: listTypes(t), Got: w.Typeof(lv)}
}

// PullAnyOfOpt reflects lv to the first successful type from t. Returns d if
// none of the types were successful.
func (w *World) PullAnyOfOpt(lv lua.LValue, d types.Value, t ...string) (v types.Value) {
	for _, t := range t {
		pull, err := w.PullerOf(t)
		if err != nil {
			return d
		}
		if v, err := pull(w.Context(), lv); err == nil {
			return v
		}
	}
	return d
}

// PushArrayOf reflect v to lv, ensuring that each element is reflected
// according to t.
func (w *World) PushArrayOf(v rtypes.Array, t string) (lv *lua.LTable, err error) {
	s := w.Context()
	if s.CycleGuard() {
		defer s.CycleClear()
	}
	if s.CycleMark(&v) {
		return nil, fmt.Errorf("arrays cannot be cyclic")
	}
	push, err := w.PusherOf(t)
	if err != nil {
		return nil, err
	}
	table := w.l.CreateTable(len(v), 0)
	for _, v := range v {
		lv, err := push(s, v)
		if err != nil {
			return nil, err
		}
		table.Append(lv)
	}
	return table, nil
}

// PullArrayOf reflects lv to v, ensuring that lv is a table, and that each
// element is reflected according to t.
func (w *World) PullArrayOf(lv lua.LValue, t string) (v rtypes.Array, err error) {
	pull, err := w.PullerOf(t)
	if err != nil {
		return nil, err
	}
	s := w.Context()
	if s.CycleGuard() {
		defer s.CycleClear()
	}
	table, ok := lv.(*lua.LTable)
	if !ok {
		return nil, TypeError{Want: "table", Got: lv.Type().String()}
	}
	if s.CycleMark(table) {
		return nil, fmt.Errorf("tables cannot be cyclic")
	}
	l := table.Len()
	array := make(rtypes.Array, l)
	for i := 1; i <= l; i++ {
		var err error
		if array[i-1], err = pull(s, table.RawGetInt(i)); err != nil {
			return nil, err
		}
	}
	return array, nil
}

// PushDictionaryOf reflect v to lv, ensuring that each field is reflected
// according to t.
func (w *World) PushDictionaryOf(v rtypes.Dictionary, t string) (lv *lua.LTable, err error) {
	s := w.Context()
	if s.CycleGuard() {
		defer s.CycleClear()
	}
	if s.CycleMark(&v) {
		return nil, fmt.Errorf("dictionaries cannot be cyclic")
	}
	push, err := w.PusherOf(t)
	if err != nil {
		return nil, err
	}
	table := w.l.CreateTable(0, len(v))
	for k, v := range v {
		lv, err := push(s, v)
		if err != nil {
			return nil, err
		}
		table.RawSetString(k, lv)
	}
	return table, nil
}

// PullDictionaryOf reflects lv to v, ensuring that lv is a table, and that each
// string field is reflected according to t.
func (w *World) PullDictionaryOf(lv lua.LValue, t string) (v rtypes.Dictionary, err error) {
	pull, err := w.PullerOf(t)
	if err != nil {
		return nil, err
	}
	s := w.Context()
	if s.CycleGuard() {
		defer s.CycleClear()
	}
	table, ok := lv.(*lua.LTable)
	if !ok {
		return nil, TypeError{Want: "table", Got: lv.Type().String()}
	}
	if s.CycleMark(table) {
		return nil, fmt.Errorf("tables cannot be cyclic")
	}
	dict := make(rtypes.Dictionary)
	err = table.ForEach(func(k, lv lua.LValue) error {
		v, err := pull(s, lv)
		if err != nil {
			return err
		}
		dict[k.String()] = v
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dict, nil
}

// PullEncoded pulls a value to be encoded according to a FormatSelector. The
// referred format is acquired, then Format.EncodeTypes is used to reflect the
// value from lv. If fs.Format is empty, or if EncodeTypes is empty, then the
// value is pulled as a Variant.
func (w *World) PullEncoded(lv lua.LValue, fs rtypes.FormatSelector) (v types.Value, err error) {
	if fs.Format == "" {
		return w.Pull(lv, "Variant")
	}
	format := w.Format(fs.Format)
	if format.Name == "" {
		return nil, fmt.Errorf("unknown format %q", fs.Format)
	}
	return w.PullEncodedFormat(lv, w.Format(fs.Format))
}

// PullEncodedFormat pulls a value to be encoded according to f.
// Format.EncodeTypes is used to reflect the value from lv. If EncodeTypes is
// empty, then the value is pulled as a Variant.
func (w *World) PullEncodedFormat(lv lua.LValue, f Format) (v types.Value, err error) {
	if len(f.EncodeTypes) == 0 {
		return w.Pull(lv, "Variant")
	}
	return w.PullAnyOf(lv, f.EncodeTypes...)
}
