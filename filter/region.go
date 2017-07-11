package filter

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/format"
	"github.com/robloxapi/rbxfile"
)

func init() {
	Filters.Register(
		rbxmk.Filter{Name: "region", Func: Region},
	)
}

func Region(f rbxmk.FilterArgs, opt rbxmk.Options, arguments []interface{}) (results []interface{}, err error) {
	output := arguments[0].(rbxmk.Data)
	region := arguments[1].(string)
	input := arguments[2].(rbxmk.Data)
	f.ProcessedArgs()

	var drill rbxmk.Data
	switch inst := output.(type) {
	case *rbxfile.Instance:
		switch inst.ClassName {
		case "Script", "LocalScript", "ModuleScript":
			source := format.Property{Properties: inst.Properties, Name: "Source"}
			if drill, _, err = format.DrillRegion(opt, source, []string{region}); err != nil {
				if _, ok := err.(format.RegionError); ok {
					return []interface{}{output}, nil
				}
				return nil, err
			}
		default:
			return nil, rbxmk.NewDataTypeError(output)
		}
	default:
		if drill, _, err = format.DrillRegion(opt, output, []string{region}); err != nil {
			return nil, err
		}
	}

	switch v := input.(type) {
	case bool:
		input = rbxfile.ValueBool(v)
	case float64:
		input = rbxfile.ValueDouble(v)
	case string:
		input = rbxfile.ValueString(v)
	}

	output, err = format.MergeTable(opt, output, drill, input)
	return []interface{}{output}, err
}
