package format

import (
	"encoding/json"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/types"
	"github.com/robloxapi/rbxfile"
	rbxfilejson "github.com/robloxapi/rbxfile/json"
	"io"
)

func init() {
	Formats.Register(rbxmk.Format{
		Name: "JSON Properties",
		Ext:  "properties.json",
		Codec: func(opt *rbxmk.Options, ctx interface{}) rbxmk.FormatCodec {
			return &JSONCodec{}
		},
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
	props := make(types.Properties, len(jprops))
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
	switch v := data.(type) {
	case *types.Instances:
		if len(*v) > 0 {
			data = types.Properties((*v)[0].Properties)
		}
	case types.Instance:
		data = types.Properties(v.Properties)
	case types.Property:
		data = types.Properties{v.Name: v.Properties[v.Name]}
	case nil:
		data = types.Properties{}
	}

	props, ok := data.(types.Properties)
	if !ok {
		return rbxmk.NewDataTypeError(data)
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
