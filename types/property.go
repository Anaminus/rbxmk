package types

import (
	"fmt"
	"github.com/anaminus/rbxmk"
)

type Property struct {
	Properties Properties
	Name       string
}

func (indata Property) Type() string {
	return "Property"
}

func (indata Property) Drill(opt rbxmk.Options, inref []string) (outdata rbxmk.Data, outref []string, err error) {
	if s := NewStringlike(indata.Properties[indata.Name]); s != nil {
		if outdata, outref, err = s.Drill(opt, inref); err != nil {
			return indata, inref, err
		}
		region := outdata.(*Region)
		region.Property = &indata
		return outdata, outref, nil
	}
	return indata, inref, rbxmk.EOD
}

func (indata Property) Merge(opt rbxmk.Options, rootdata, drilldata rbxmk.Data) (outdata rbxmk.Data, err error) {
	value := indata.Properties[indata.Name]
	if value == nil {
		return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("input property \"%s\" cannot be nil", indata.Name))
	}
	switch drilldata := drilldata.(type) {
	case *Instances:
		for _, inst := range *drilldata {
			if propertyIsOfType(opt.Config.API, inst, indata.Name, value.Type()) {
				inst.Properties[indata.Name] = value
			}
		}
		return rootdata, nil

	case Instance:
		if propertyIsOfType(opt.Config.API, drilldata.Instance, indata.Name, value.Type()) {
			drilldata.Properties[indata.Name] = value
		}
		return rootdata, nil

	case Properties:
		if v, _ := drilldata[indata.Name]; v == nil || v.Type() == value.Type() {
			drilldata[indata.Name] = value
		}
		return rootdata, nil

	case Property:
		if v, _ := drilldata.Properties[drilldata.Name]; v == nil || v.Type() == value.Type() {
			drilldata.Properties[drilldata.Name] = value
		} else {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf(
				"input property \"%s\" cannot be assigned to output property \"%s\": expected %s, got %s",
				indata.Name, drilldata.Name, v.Type(), value.Type(),
			))
		}
		return rootdata, nil

	case *Region:
		s := NewStringlike(value)
		if s == nil {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("input property \"%s\" must be stringlike", indata.Name))
		}
		drilldata.Set(s.Bytes)
		if drilldata.Property == nil {
			return drilldata.Value, nil
		}
		return rootdata, nil

	case *Stringlike:
		if !drilldata.SetFrom(value) {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("input property \"%s\" must be stringlike", indata.Name))
		}
		return drilldata, nil

	case Value:
		if drilldata.Value.Type() != value.Type() {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("expected input type %s, got %s", drilldata.Type(), value.Type()))
		}
		return Value{value}, nil

	case nil:
		return Value{indata.Properties[indata.Name]}, nil
	}
	return nil, rbxmk.NewMergeError(indata, drilldata, nil)
}
