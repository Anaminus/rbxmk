package formats

import (
	"encoding/json"
	"io"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func decodeJSON(u interface{}) (v types.Value) {
	switch u := u.(type) {
	case nil:
		return rtypes.Nil
	case bool:
		return types.Bool(u)
	case float64:
		return types.Double(u)
	case string:
		return types.String(u)
	case []interface{}:
		a := make(rtypes.Array, len(u))
		for i, u := range u {
			a[i] = decodeJSON(u)
		}
		return a
	case map[string]interface{}:
		a := make(rtypes.Dictionary, len(u))
		for k, u := range u {
			a[k] = decodeJSON(u)
		}
		return a
	}
	return rtypes.Nil
}

func encodeJSON(v types.Value) (u interface{}) {
	//WARN: Must not receive cyclic tables. The Array and Dictionary type
	//reflectors already validate this, but such values could still be produced
	//internally.
	switch v := v.(type) {
	case rtypes.NilType:
		return nil
	case types.Bool:
		return bool(v)
	case types.Double:
		return float64(v)
	case types.String:
		return string(v)
	case rtypes.Array:
		a := make([]interface{}, len(v))
		for i, v := range v {
			a[i] = encodeJSON(v)
		}
		return a
	case rtypes.Dictionary:
		a := make(map[string]interface{}, len(v))
		for k, v := range v {
			a[k] = encodeJSON(v)
		}
		return a
	}
	return nil
}

func init() { register(JSON) }
func JSON() rbxmk.Format {
	return rbxmk.Format{
		Name:       "json",
		MediaTypes: []string{"application/json", "text/plain"},
		CanDecode: func(f rbxmk.FormatOptions, typeName string) bool {
			switch typeName {
			case "nil", "bool", "double", "string", "Array", "Dictionary":
				return true
			}
			return false
		},
		Decode: func(f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			var u interface{}
			j := json.NewDecoder(r)
			if err := j.Decode(&u); err != nil {
				return nil, err
			}
			return decodeJSON(u), nil
		},
		Encode: func(f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			j := json.NewEncoder(w)
			j.SetIndent("", "\t")
			j.SetEscapeHTML(false)
			return j.Encode(encodeJSON(v))
		},
	}
}
