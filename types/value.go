package types

import (
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
)

type Value struct {
	rbxfile.Value
}

func (indata Value) Type() string {
	if indata.Value == nil {
		return "Value<nil>"
	}
	return "Value<" + indata.Value.Type().String() + ">"
}

func (indata Value) Drill(opt *rbxmk.Options, inref []string) (outdata rbxmk.Data, outref []string, err error) {
	if s := NewStringlike(indata); s != nil {
		return s.Drill(opt, inref)
	}
	return indata, inref, rbxmk.EOD
}

func (indata Value) Merge(opt *rbxmk.Options, rootdata, drilldata rbxmk.Data) (outdata rbxmk.Data, err error) {
	switch drilldata := drilldata.(type) {
	case Property:
		if v := drilldata.Properties[drilldata.Name]; v != nil && indata.Value.Type() != v.Type() {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("expected input type %s, got %s", v.Type(), indata.Type()))
		}
		drilldata.Properties[drilldata.Name] = indata.Value
		return rootdata, nil

	case Value:
		if drilldata.Value != nil && indata.Value.Type() != drilldata.Value.Type() {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("expected input type %s, got %s", drilldata.Type(), indata.Type()))
		}
		return indata, nil

	case nil:
		return indata, nil
	}
	s := NewStringlike(indata)
	if s == nil {
		return nil, rbxmk.NewMergeError(indata, drilldata, nil)
	}
	return s.Merge(opt, rootdata, drilldata)
}
