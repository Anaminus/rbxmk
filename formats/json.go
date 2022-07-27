package formats

import (
	"encoding/json"
	"io"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

const F_JSON = "json"

func init() { register(JSON) }
func JSON() rbxmk.Format {
	return rbxmk.Format{
		Name:        F_JSON,
		EncodeTypes: []string{rtypes.T_JsonValue},
		MediaTypes:  []string{"application/json", "text/plain"},
		Options: map[string][]string{
			"Indent": {rtypes.T_String},
		},
		CanDecode: func(g rtypes.Global, f rbxmk.FormatOptions, typeName string) bool {
			switch typeName {
			case rtypes.T_Nil, rtypes.T_Bool, rtypes.T_Double, rtypes.T_String, rtypes.T_Array, rtypes.T_Dictionary:
				return true
			}
			return false
		},
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			var u interface{}
			j := json.NewDecoder(r)
			if err := j.Decode(&u); err != nil {
				return nil, err
			}
			return rtypes.DecodeJSON(u), nil
		},
		Encode: func(g rtypes.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			j := json.NewEncoder(w)
			j.SetIndent("", "\t")
			if v, ok := stringOf(f, "Indent"); ok {
				j.SetIndent("", v)
			}
			j.SetEscapeHTML(false)
			return j.Encode(rtypes.EncodeJSON(v))
		},
		Dump: func() dump.Format {
			return dump.Format{
				Options: dump.FormatOptions{
					"Indent": dump.FormatOption{
						Type:        dt.Prim(rtypes.T_String),
						Default:     `"\t"`,
						Description: "Formats/options/json:Indent",
					},
				},
				Summary:     "Formats/json:Summary",
				Description: "Formats/json:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			reflect.Array,
			reflect.Bool,
			reflect.Dictionary,
			reflect.Double,
			reflect.Nil,
			reflect.String,
		},
	}
}

const F_JSONPatch = "patch.json"

func init() { register(JSONPatch) }
func JSONPatch() rbxmk.Format {
	return rbxmk.Format{
		Name:        F_JSONPatch,
		EncodeTypes: []string{rtypes.T_JsonPatch},
		MediaTypes:  []string{"application/json", "text/plain"},
		Options: map[string][]string{
			"Indent": {rtypes.T_String},
		},
		CanDecode: func(g rtypes.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == rtypes.T_JsonPatch
		},
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			var patch rtypes.JsonPatch
			j := json.NewDecoder(r)
			if err := j.Decode(&patch); err != nil {
				return nil, err
			}
			return patch, nil
		},
		Encode: func(g rtypes.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			j := json.NewEncoder(w)
			j.SetIndent("", "\t")
			if v, ok := stringOf(f, "Indent"); ok {
				j.SetIndent("", v)
			}
			j.SetEscapeHTML(false)
			patch := v.(rtypes.JsonPatch)
			return j.Encode(patch)
		},
		Dump: func() dump.Format {
			return dump.Format{
				Options: dump.FormatOptions{
					"Indent": dump.FormatOption{
						Type:        dt.Prim(rtypes.T_String),
						Default:     `"\t"`,
						Description: "Formats/options/json:Indent",
					},
				},
				Summary:     "Formats/patch.json:Summary",
				Description: "Formats/patch.json:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			reflect.JsonPatch,
		},
	}
}
