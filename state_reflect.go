package rbxmk

import (
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

// TypeOf returns the result of World.Typeof with the given argument.
func (s State) Typeof(n int) string {
	return s.World.Typeof(s.L.Get(n))
}

// Push reflects v according to its type as registered with s.World, then pushes
// the results to s.L.
func (s State) Push(v types.Value) int {
	lv, err := s.World.Push(v)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	s.L.Push(lv)
	return 1
}

// Pull gets from s.L the values starting from n, and reflects a value from them
// according to type t registered with s.World.
func (s State) Pull(n int, t string) (v types.Value) {
	v, err := s.World.Pull(s.CheckAny(n), t)
	if err != nil {
		s.ArgError(n, err.Error())
		return nil
	}
	return v
}

// PullOpt gets from s.L the value at n, and reflects a value from it according
// to type t registered with s.World. If the value is nil, d is returned
// instead.
func (s State) PullOpt(n int, d types.Value, t string) (v types.Value) {
	v, err := s.World.PullOpt(s.L.Get(n), d, t)
	if err != nil {
		s.ArgError(n, err.Error())
		return nil
	}
	return v
}

// PullAnyOf gets from s.L the values starting from n, and reflects a value from
// them according to any of the types in t registered with s.World. Returns the
// first successful reflection among the types in t. If no types succeeded, then
// a type error is thrown.
func (s State) PullAnyOf(n int, t ...string) (v types.Value) {
	v, err := s.World.PullAnyOf(s.CheckAny(n), t...)
	if err != nil {
		s.TypeError(n, listTypes(t), s.Typeof(n))
		return nil
	}
	return v
}

// PullAnyOfOpt gets from s.L the values starting from n, and reflects a value
// from them according to any of the types in t registered with s.World. Returns
// the first successful reflection among the types in t. If no types succeeded,
// then nil is returned.
func (s State) PullAnyOfOpt(n int, d types.Value, t ...string) (v types.Value) {
	return s.World.PullAnyOfOpt(s.CheckAny(n), d, t...)
}

// PushArrayOf pushes an rtypes.Array, ensuring that each element is reflected
// according to t.
func (s State) PushArrayOf(v rtypes.Array, t string) int {
	lv, err := s.World.PushArrayOf(v, t)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	s.L.Push(lv)
	return 1
}

// PullArrayOf pulls an rtypes.Array from n, ensuring that each element is
// reflected according to t.
func (s State) PullArrayOf(n int, t string) (v rtypes.Array) {
	v, err := s.World.PullArrayOf(s.CheckAny(n), t)
	if err != nil {
		s.ArgError(n, err.Error())
		return nil
	}
	return v
}

// PushDictionaryOf pushes an rtypes.Dictionary, ensuring that each field is
// reflected according to t.
func (s State) PushDictionaryOf(v rtypes.Dictionary, t string) int {
	lv, err := s.World.PushDictionaryOf(v, t)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	s.L.Push(lv)
	return 1
}

// PullDictionaryOf pulls an rtypes.Dictionary from n, ensuring that each
// field is reflected according to t.
func (s State) PullDictionaryOf(n int, t string) (v rtypes.Dictionary) {
	v, err := s.World.PullDictionaryOf(s.CheckAny(n), t)
	if err != nil {
		s.ArgError(n, err.Error())
		return nil
	}
	return v
}

// PullEncoded pulls a value to be encoded according to a FormatSelector. The
// referred format is determined, then Format.EncodeTypes is used to pull the
// value from n. If fs.Format is empty, or if EncodeTypes is empty, then the
// value is pulled as a Variant.
func (s State) PullEncoded(n int, fs rtypes.FormatSelector) (v types.Value) {
	v, err := s.World.PullEncoded(s.CheckAny(n), fs)
	if err != nil {
		s.ArgError(n, err.Error())
		return nil
	}
	return v
}

// PullEncodedFormat pulls a value to be encoded according to a Format.
// Format.EncodeTypes is used to pull the value from n. If EncodeTypes is empty,
// then the value is pulled as a Variant.
func (s State) PullEncodedFormat(n int, f Format) (v types.Value) {
	v, err := s.World.PullEncodedFormat(s.CheckAny(n), f)
	if err != nil {
		s.ArgError(n, err.Error())
		return nil
	}
	return v
}
