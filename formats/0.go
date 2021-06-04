package formats

import (
	"fmt"
	"strconv"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

// registry contains registered Formats.
var registry []func() rbxmk.Format

// register registers a Format to be returned by All.
func register(f func() rbxmk.Format) {
	registry = append(registry, f)
}

// All returns a list of Formats defined in the package.
func All() []func() rbxmk.Format {
	a := make([]func() rbxmk.Format, len(registry))
	copy(a, registry)
	return a

}

// cannotEncode returns an error indicating that v cannot be encoded.
func cannotEncode(v interface{}) error {
	if v, ok := v.(types.Value); ok {
		return fmt.Errorf("cannot encode %s", v.Type())
	}
	return fmt.Errorf("cannot encode %T", v)
}

// stringOf returns field as a string. Returns an empty string if the field does
// not exist or cannot be converted to a string.
func stringOf(f rbxmk.FormatOptions, field string) (v string, ok bool) {
	if f == nil {
		return "", false
	}
	switch v := f.ValueOf(field).(type) {
	case types.Stringlike:
		return v.Stringlike(), true
	case types.Numberlike:
		return strconv.FormatFloat(v.Numberlike(), 'g', 0, 64), true
	case types.Intlike:
		return strconv.FormatInt(v.Intlike(), 10), true
	default:
		return "", false
	}
}

// numberOf returns field as a number. Returns 0 if the field does not exist or
// cannot be converted to a number.
func numberOf(f rbxmk.FormatOptions, field string) (v float64, ok bool) {
	if f == nil {
		return 0, false
	}
	switch v := f.ValueOf(field).(type) {
	case types.Numberlike:
		return v.Numberlike(), true
	case types.Intlike:
		return float64(v.Intlike()), true
	case types.Stringlike:
		n, _ := strconv.ParseFloat(v.Stringlike(), 64)
		return n, true
	default:
		return 0, false
	}
}

// intOf returns field as a string. Returns 0 if the field does not exist or
// cannot be converted to an integer.
func intOf(f rbxmk.FormatOptions, field string) (v int64, ok bool) {
	if f == nil {
		return 0, false
	}
	switch v := f.ValueOf(field).(type) {
	case types.Intlike:
		return v.Intlike(), true
	case types.Numberlike:
		return int64(v.Numberlike()), true
	case types.Stringlike:
		n, _ := strconv.ParseInt(v.Stringlike(), 0, 64)
		return n, true
	default:
		return 0, false
	}
}

// boolOf returns field as a bool. Returns false if the field does not exist, is
// nil, or is the boolean false. Returns true otherwise.
func boolOf(f rbxmk.FormatOptions, field string) (v bool, ok bool) {
	if f == nil {
		return false, false
	}
	switch v := f.ValueOf(field).(type) {
	case nil, rtypes.NilType:
		return false, true
	case types.Bool:
		return v == true, true
	default:
		return true, false
	}
}
