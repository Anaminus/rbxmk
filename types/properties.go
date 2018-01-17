package types

import (
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/config"
	"github.com/robloxapi/rbxfile"
)

type Properties map[string]rbxfile.Value

func (indata Properties) Type() string {
	return "Properties"
}

func (indata Properties) Drill(opt rbxmk.Options, inref []string) (outdata rbxmk.Data, outref []string, err error) {
	if len(inref) == 0 {
		err = rbxmk.EOD
		return indata, inref, err
	}

	if _, exists := indata[inref[0]]; !exists {
		return indata, inref, fmt.Errorf("property %q not present in instance", inref[0])
	}
	return Property{Properties: indata, Name: inref[0]}, inref[1:], nil
}

func (indata Properties) Merge(opt rbxmk.Options, rootdata, drilldata rbxmk.Data) (outdata rbxmk.Data, err error) {
	switch drilldata := drilldata.(type) {
	case *Instances:
		for _, inst := range *drilldata {
			for name, value := range indata {
				if propertyIsOfType(config.API(opt), inst, name, value.Type()) {
					inst.Properties[name] = value
				}
			}
		}
		return rootdata, nil

	case Instance:
		for name, value := range indata {
			if propertyIsOfType(config.API(opt), drilldata.Instance, name, value.Type()) {
				drilldata.Properties[name] = value
			}
		}
		return rootdata, nil

	case Properties:
		for name, value := range indata {
			if v, _ := drilldata[name]; v == nil || v.Type() == value.Type() {
				drilldata[name] = value
			}
		}
		return rootdata, nil

	case Property:
		value, ok := indata[drilldata.Name]
		if !ok {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("\"%s\" not found in properties", drilldata.Name))
		}
		drilldata.Properties[drilldata.Name] = value
		return rootdata, nil

	case *Region:
		if drilldata.Property == nil {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("Region must have Property"))
		}
		value, ok := indata[drilldata.Property.Name]
		if !ok {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("\"%s\" not found in properties", drilldata.Property.Name))
		}
		s := NewStringlike(value)
		if s == nil {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("value of \"%s\" must be stringlike", drilldata.Property.Name))
		}
		drilldata.Set(s.Bytes)
		return rootdata, nil

	case nil:
		return indata, nil
	}
	return nil, rbxmk.NewMergeError(indata, drilldata, nil)
}
