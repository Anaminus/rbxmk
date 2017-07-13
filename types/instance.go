package types

import (
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
)

type Instance struct {
	*rbxfile.Instance
}

func (indata Instance) Type() string {
	if indata.Instance == nil {
		return "Instance<nil>"
	}
	return "Instance<" + indata.ClassName + ">"
}

func (indata Instance) Drill(opt rbxmk.Options, inref []string) (outdata rbxmk.Data, outref []string, err error) {
	if len(inref) == 0 {
		err = rbxmk.EOD
		return indata, inref, err
	}

	if indata.Instance == nil {
		return indata, inref, fmt.Errorf("Instance cannot be nil")
	}

	ref := inref[0]
	if ref == "" {
		return indata, inref, fmt.Errorf("property not specified")
	}
	if ref == "*" {
		// Select all properties.
		return Properties(indata.Properties), inref[1:], nil
	}

	// TODO: API?

	return Property{Properties: indata.Properties, Name: ref}, inref[1:], nil
}

func (indata Instance) Merge(opt rbxmk.Options, rootdata, drilldata rbxmk.Data) (outdata rbxmk.Data, err error) {
	switch drilldata := drilldata.(type) {
	case *Instances:
		*drilldata = append(*drilldata, indata.Instance)
		return rootdata, nil

	case Instance:
		drilldata.AddChild(indata.Instance)
		return rootdata, nil

	case Property:
		if typeOfProperty(opt.Config.API, indata.ClassName, drilldata.Name) == rbxfile.TypeReference ||
			drilldata.Properties[drilldata.Name].Type() != rbxfile.TypeReference {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("property must be a Reference"))
		}
		drilldata.Properties[drilldata.Name] = rbxfile.ValueReference{Instance: indata.Instance}
		return rootdata, nil

	case nil:
		return indata, nil
	}
	return nil, rbxmk.NewMergeError(indata, drilldata, nil)
}
