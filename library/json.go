package library

import (
	"bytes"
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/formats"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/robloxapi/types"
	"github.com/wI2L/jsondiff"
)

func init() { register(JSON) }

var JSON = rbxmk.Library{
	Name:     "json",
	Import:   []string{"json"},
	Priority: 10,
	Open:     openJSON,
	Dump:     dumpJSON,
	Types: []func() rbxmk.Reflector{
		reflect.JsonPatch,
		reflect.JsonValue,
		reflect.String,
	},
}

func openJSON(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 4)
	lib.RawSetString("diff", s.WrapFunc(jsonDiff))
	lib.RawSetString("fromString", s.WrapFunc(jsonFromString))
	lib.RawSetString("patch", s.WrapFunc(jsonPatch))
	lib.RawSetString("string", s.WrapFunc(jsonString))
	return lib
}

func jsonBytes(b []byte) (v types.Value, err error) {
	format := formats.JSON()
	return format.Decode(rtypes.Global{}, rtypes.FormatSelector{}, bytes.NewReader(b))
}

func jsonPatchBytes(b []byte) (v types.Value, err error) {
	format := formats.JSONPatch()
	return format.Decode(rtypes.Global{}, rtypes.FormatSelector{}, bytes.NewReader(b))
}

func jsonToBytes(v types.Value) (b []byte, err error) {
	format := formats.JSON()
	var w bytes.Buffer
	if err := format.Encode(rtypes.Global{}, rtypes.FormatSelector{
		Options: rtypes.Dictionary{"Indent": types.String("")},
	}, &w, v); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func jsonDiff(s rbxmk.State) int {
	p := s.Pull(1, rtypes.T_JsonValue)
	n := s.Pull(2, rtypes.T_JsonValue)
	patch, err := jsondiff.Compare(rtypes.EncodeJSON(p), rtypes.EncodeJSON(n))
	if err != nil {
		return s.RaiseError("compare: %s", err)
	}
	lpatch := make(rtypes.JsonPatch, len(patch))
	for i, op := range patch {
		lpatch[i] = rtypes.JsonOperation{
			Op:    rtypes.JsonOp(op.Type),
			From:  string(op.From),
			Path:  string(op.Path),
			Value: rtypes.JsonValue{Value: rtypes.DecodeJSON(op.Value)},
		}
	}
	return s.Push(lpatch)
}

func jsonFromString(s rbxmk.State) int {
	format := formats.JSON()
	r := strings.NewReader(string(s.Pull(1, rtypes.T_String).(types.String)))
	v, err := format.Decode(s.Global, rtypes.FormatSelector{}, r)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return s.Push(v)
}

func jsonPatch(s rbxmk.State) int {
	v := s.Pull(1, rtypes.T_JsonValue)
	p := s.Pull(2, rtypes.T_JsonPatch).(rtypes.JsonPatch)

	vf := formats.JSON()
	var vw bytes.Buffer
	if err := vf.Encode(rtypes.Global{}, rtypes.FormatSelector{}, &vw, v); err != nil {
		return s.RaiseError("encode value: %s", err)
	}

	pf := formats.JSONPatch()
	var pw bytes.Buffer
	if err := pf.Encode(rtypes.Global{}, rtypes.FormatSelector{}, &pw, p); err != nil {
		return s.RaiseError("encode patch: %s", err)
	}

	patch, err := jsonpatch.DecodePatch(pw.Bytes())
	if err != nil {
		return s.RaiseError("decode patch: %s", err)
	}

	b, err := patch.Apply(vw.Bytes())
	if err != nil {
		return s.RaiseError("apply patch: %s", err)
	}

	value, err := jsonBytes(b)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return s.Push(value)
}

func jsonString(s rbxmk.State) int {
	format := formats.JSON()
	value := s.PullEncodedFormat(1, format)

	var selector rtypes.FormatSelector
	indent := s.PullOpt(2, nil, rtypes.T_String)
	if indent != nil {
		selector.Options = rtypes.Dictionary{"Indent": indent.(types.String)}
	}

	var w bytes.Buffer
	if err := format.Encode(s.Global, selector, &w, value); err != nil {
		return s.RaiseError("%s", err)
	}
	return s.Push(types.BinaryString(bytes.TrimSuffix(w.Bytes(), []byte{'\n'})))
}

func dumpJSON(s rbxmk.State) dump.Library {
	return dump.Library{
		Struct: dump.Struct{
			Fields: dump.Fields{
				"diff": dump.Function{
					Parameters: dump.Parameters{
						{Name: "prev", Type: dt.Prim(rtypes.T_JsonValue)},
						{Name: "next", Type: dt.Prim(rtypes.T_JsonValue)},
					},
					Returns: dump.Parameters{
						{Type: dt.Optional(dt.Array(dt.Prim(rtypes.T_JsonPatch)))},
					},
					CanError:    true,
					Summary:     "Libraries/json:Fields/diff/Summary",
					Description: "Libraries/json:Fields/diff/Description",
				},
				"fromString": dump.Function{
					Parameters: dump.Parameters{
						{Name: "string", Type: dt.Prim(rtypes.T_String)},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_JsonValue)},
					},
					CanError:    true,
					Summary:     "Libraries/json:Fields/fromString/Summary",
					Description: "Libraries/json:Fields/fromString/Description",
				},
				"patch": dump.Function{
					Parameters: dump.Parameters{
						{Name: "value", Type: dt.Prim(rtypes.T_JsonValue)},
						{Name: "patch", Type: dt.Prim(rtypes.T_JsonPatch)},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_JsonValue)},
					},
					CanError:    true,
					Summary:     "Libraries/json:Fields/patch/Summary",
					Description: "Libraries/json:Fields/patch/Description",
				},
				"string": dump.Function{
					Parameters: dump.Parameters{
						{Name: "value", Type: dt.Prim(rtypes.T_JsonValue)},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_String)},
					},
					CanError:    true,
					Summary:     "Libraries/json:Fields/string/Summary",
					Description: "Libraries/json:Fields/string/Description",
				},
			},
			Summary:     "Libraries/json:Summary",
			Description: "Libraries/json:Description",
		},
	}
}
