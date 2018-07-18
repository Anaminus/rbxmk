package types

import (
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
)

type Delete struct{}

func (indata Delete) Type() string {
	return "Delete"
}

func (indata Delete) Drill(opt *rbxmk.Options, inref []string) (outdata rbxmk.Data, outref []string, err error) {
	return indata, inref, rbxmk.EOD
}

func (indata Delete) Merge(opt *rbxmk.Options, rootdata, drilldata rbxmk.Data) (outdata rbxmk.Data, err error) {
	switch drilldata := drilldata.(type) {
	case *Instances:
		*drilldata = (*drilldata)[:0]
		return rootdata, nil

	case Instance:
		drilldata.SetParent(nil)
		return rootdata, nil

	case Properties:
		for k := range drilldata {
			delete(drilldata, k)
		}
		return rootdata, nil

	case Property:
		delete(drilldata.Properties, drilldata.Name)
		return rootdata, nil

	case *Region:
		drilldata.Set(nil)
		if drilldata.Property == nil {
			return drilldata.Value, nil
		}
		return rootdata, nil

	case *Stringlike:
		drilldata.ValueType = rbxfile.TypeInvalid
		drilldata.Bytes = nil
		return drilldata, nil

	case Value:
		return nil, nil

	case nil:
		return nil, nil
	}
	return nil, rbxmk.NewMergeError(indata, drilldata, nil)
}
