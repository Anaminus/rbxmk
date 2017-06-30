package format

import (
	"encoding/json"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	rbxfilejson "github.com/robloxapi/rbxfile/json"
	"io"
)

func init() {
	register(rbxmk.Format{
		Name: "JSON Properties",
		Ext:  "properties.json",
		Codec: func(opt rbxmk.Options, ctx interface{}) rbxmk.FormatCodec {
			return &JSONCodec{}
		},
		InputDrills:  []rbxmk.Drill{DrillProperty},
		OutputDrills: []rbxmk.Drill{DrillProperty},
		Merger:       MergeProperties,
	})
}

type jsonPropList map[string]jsonProp
type jsonProp struct {
	typ   string      `json:"type"`
	value interface{} `json:"value"`
}

type JSONCodec struct{}

func (c JSONCodec) Decode(r io.Reader, data *rbxmk.Data) (err error) {
	jprops := jsonPropList{}
	if err := json.NewDecoder(r).Decode(&jprops); err != nil {
		return err
	}
	props := make(map[string]rbxfile.Value, len(jprops))
	for name, prop := range jprops {
		value := rbxfilejson.ValueFromJSONInterface(rbxfile.TypeFromString(prop.typ), prop.value)
		if value == nil {
			continue
		}
		props[name] = value
	}
	*data = props
	return nil
}

func (c JSONCodec) Encode(w io.Writer, data rbxmk.Data) (err error) {
	var instance *rbxfile.Instance
	switch v := data.(type) {
	case []*rbxfile.Instance:
		if len(v) > 0 {
			instance = v[0]
		}
	case *rbxfile.Instance:
		instance = v
	case Property:
		data = map[string]rbxfile.Value{v.Name: v.Properties[v.Name]}
	}
	if instance != nil {
		data = instance.Properties
	}

	props, ok := data.(map[string]rbxfile.Value)
	if !ok {
		return NewDataTypeError(data)
	}
	jprops := make(map[string]jsonProp, len(props))
	for name, value := range props {
		jprops[name] = jsonProp{
			typ:   value.Type().String(),
			value: rbxfilejson.ValueToJSONInterface(value, nil),
		}
	}

	je := json.NewEncoder(w)
	je.SetIndent("", "\t")
	return je.Encode(&jprops)
}
