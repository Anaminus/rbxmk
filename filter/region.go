package filter

import (
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/types"
)

func init() {
	Filters.Register(
		rbxmk.Filter{Name: "region", Func: Region},
	)
}

func Region(f rbxmk.FilterArgs, opt *rbxmk.Options, arguments []interface{}) (results []interface{}, err error) {
	output := arguments[0].(rbxmk.Data)
	region := arguments[1].(string)
	input := arguments[2].(interface{})
	f.ProcessedArgs()

	var drill rbxmk.Data
	switch output := output.(type) {
	case types.Instance:
		switch output.ClassName {
		case "Script", "LocalScript", "ModuleScript":
			source := types.Property{Properties: output.Properties, Name: "Source"}
			if drill, _, err = source.Drill(opt, []string{region}); err != nil {
				if _, ok := err.(types.RegionError); ok {
					return []interface{}{output}, nil
				}
				return nil, err
			}
		default:
			return nil, rbxmk.NewDataTypeError(output)
		}
	case types.Property, *types.Stringlike:
		if drill, _, err = output.Drill(opt, []string{region}); err != nil {
			return nil, err
		}
	default:
		return nil, rbxmk.NewDataTypeError(output)
	}

	var indata rbxmk.Data
	switch v := input.(type) {
	case rbxmk.Data:
		indata = v
	default:
		if indata = types.NewStringlike(v); indata == nil {
			return nil, fmt.Errorf("unexpected input type")
		}
	}

	output, err = indata.Merge(opt, output, drill)
	return []interface{}{output}, err
}
