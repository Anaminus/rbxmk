package filter

import (
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/types"
	"github.com/robloxapi/rbxfile"
)

// ProcessStringlikeCallback receives and modifies a Stringlike value.
type ProcessStringlikeCallback func(s *types.Stringlike) error

// ProcessStringlikeInterface receives an arbitrary value, and converts it
// into a types.Stringlike, if possible. The Stringlike is processed by a
// callback, and the result is applied back to the original location of the
// value. Returns the result as a rbxmk.Data.
//
// The following types are handled:
//
//     - types.Instance (Script, LocalScript, or ModuleScript; modifies the Source property)
//     - *rbxfile.Instances (modifies each instance)
//     - types.Property (any string-like property)
//     - types.Value (any string-like value)
//     - *types.Stringlike
//     - string (returns as a Stringlike)
//     - []byte (returns as a Stringlike)
func ProcessStringlikeInterface(cb ProcessStringlikeCallback, v interface{}) (out rbxmk.Data, err error) {
	switch v := v.(type) {
	case rbxmk.Data:
		switch v := v.(type) {
		case *types.Instances:
			for _, inst := range *v {
				if err := processStringlikeInstance(cb, inst, false); err != nil {
					return nil, err
				}
			}
			return v, nil
		case types.Instance:
			if err := processStringlikeInstance(cb, v.Instance, true); err != nil {
				return nil, err
			}
			return v, nil
		case types.Property:
			value, err := processStringlikeValue(cb, types.Value{v.Properties[v.Name]})
			if err != nil {
				return nil, err
			}
			v.Properties[v.Name] = value.Value
			return v, nil
		case types.Value:
			return processStringlikeValue(cb, v)
		case *types.Stringlike:
			if err := cb(v); err != nil {
				return nil, err
			}
			return v, nil
		default:
			return nil, rbxmk.NewDataTypeError(v)
		}
	case string, []byte:
		s := types.NewStringlike(v)
		if err := cb(s); err != nil {
			return nil, err
		}
		return s, nil
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("unexpected type")
}

func processStringlikeInstance(cb ProcessStringlikeCallback, inst *rbxfile.Instance, fail bool) (err error) {
	switch inst.ClassName {
	case "Script", "LocalScript", "ModuleScript":
		if source, ok := inst.Properties["Source"]; ok {
			value, _ := processStringlikeValue(cb, types.Value{source})
			inst.Properties["Source"] = value.Value
		}
		return nil
	}
	if fail {
		return fmt.Errorf("instance must be script-like")
	}
	return nil
}

func processStringlikeValue(cb ProcessStringlikeCallback, value types.Value) (out types.Value, err error) {
	if s := types.NewStringlike(value); s != nil {
		if err := cb(s); err != nil {
			return out, err
		}
		return s.GetValue(true), nil
	}
	return out, fmt.Errorf("value must be string-like")
}
