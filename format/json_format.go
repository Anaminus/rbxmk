package format

import (
	"encoding/json"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	rbxfilejson "github.com/robloxapi/rbxfile/json"
	"io"
)

func init() {
	register(rbxmk.FormatInfo{
		Name:           "JSON",
		Ext:            "json",
		Init:           func(_ *rbxmk.Options) rbxmk.Format { return &JSONFormat{} },
		InputDrills:    nil,
		OutputDrills:   nil,
		OutputResolver: ResolveOutputSource,
	})
}

type jsonPropList map[string]jsonProp
type jsonProp struct {
	typ   string      `json:"type"`
	value interface{} `json:"value"`
}

type JSONFormat struct{}

func (f JSONFormat) Decode(r io.Reader) (src *rbxmk.Source, err error) {
	props := jsonPropList{}
	if err := json.NewDecoder(r).Decode(&props); err != nil {
		return nil, err
	}
	src = &rbxmk.Source{Properties: make(map[string]rbxfile.Value, len(props))}
	for name, prop := range props {
		value := rbxfilejson.ValueFromJSONInterface(rbxfile.TypeFromString(prop.typ), prop.value)
		if value == nil {
			continue
		}
		src.Properties[name] = value
	}
	return src, nil
}

func (f JSONFormat) CanEncode(src *rbxmk.Source) bool {
	return len(src.Instances) == 0 && len(src.Values) == 0
}

func (f JSONFormat) Encode(w io.Writer, src *rbxmk.Source) (err error) {
	props := make(map[string]jsonProp, len(src.Properties))
	for name, value := range src.Properties {
		props[name] = jsonProp{
			typ:   value.Type().String(),
			value: rbxfilejson.ValueToJSONInterface(value, nil),
		}
	}

	je := json.NewEncoder(w)
	je.SetIndent("", "\t")
	return je.Encode(&props)
}
