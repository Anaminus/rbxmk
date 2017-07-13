package filter

import (
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/types"
	"github.com/robloxapi/rbxfile"
	"github.com/yuin/gopher-lua"
	"io/ioutil"
)

//go:generate gobake -package=$GOPACKAGE -compress -output minify.lua.go minify.lua

func init() {
	Filters.Register(
		rbxmk.Filter{Name: "minify", Func: Minify},
		rbxmk.Filter{Name: "unminify", Func: Unminify},
	)
}

func Minify(f rbxmk.FilterArgs, opt rbxmk.Options, arguments []interface{}) (results []interface{}, err error) {
	value := arguments[0].(interface{})
	f.ProcessedArgs()
	out, err := doMinify("minify", value)
	if err != nil {
		return nil, err
	}
	return []interface{}{out}, nil
}

func Unminify(f rbxmk.FilterArgs, opt rbxmk.Options, arguments []interface{}) (results []interface{}, err error) {
	value := arguments[0].(interface{})
	f.ProcessedArgs()
	out, err := doMinify("unminify", value)
	if err != nil {
		return nil, err
	}
	return []interface{}{out}, nil
}

func doMinify(method string, v interface{}) (out rbxmk.Data, err error) {
	switch v := v.(type) {
	case rbxmk.Data:
		switch v := v.(type) {
		case *types.Instances:
			for _, inst := range *v {
				if err := doMinifyInstance(method, inst, false); err != nil {
					return nil, err
				}
			}
			return v, nil
		case types.Instance:
			if err := doMinifyInstance(method, v.Instance, true); err != nil {
				return nil, err
			}
			return v, nil
		case types.Property:
			value, err := doMinifyValue(method, types.Value{v.Properties[v.Name]})
			if err != nil {
				return nil, err
			}
			v.Properties[v.Name] = value.Value
			return v, nil
		case types.Value:
			return doMinifyValue(method, v)
		case *types.Stringlike:
			if err := doMinifyString(method, v); err != nil {
				return nil, err
			}
			return v, nil
		default:
			return nil, rbxmk.NewDataTypeError(v)
		}
	case string, []byte:
		s := types.NewStringlike(v)
		if err := doMinifyString(method, s); err != nil {
			return nil, err
		}
		return s, nil
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("unexpected type")
}

func doMinifyInstance(method string, inst *rbxfile.Instance, fail bool) (err error) {
	switch inst.ClassName {
	case "Script", "LocalScript", "ModuleScript":
		if source, ok := inst.Properties["Source"]; ok {
			value, _ := doMinifyValue(method, types.Value{source})
			inst.Properties["Source"] = value.Value
		}
		return nil
	}
	if fail {
		return fmt.Errorf("instance must be script-like")
	}
	return nil
}

func doMinifyValue(method string, value types.Value) (out types.Value, err error) {
	if s := types.NewStringlike(value); s != nil {
		if err := doMinifyString(method, s); err != nil {
			return out, err
		}
		return s.GetValue(true), nil
	}
	return out, fmt.Errorf("value must be string-like")
}

func doMinifyString(method string, s *types.Stringlike) error {
	var l *lua.LState
	{
		l = lua.NewState()
		src := minify_lua()
		b, _ := ioutil.ReadAll(src)
		src.Close()
		fn, err := l.LoadString(string(b))
		if err != nil {
			return err
		}
		l.Push(fn)
	}
	l.Push(lua.LString(method))
	l.Push(lua.LString(s.Bytes))
	if err := l.PCall(2, 1, nil); err != nil {
		return err
	}
	s.Bytes = []byte(l.Get(-1).(lua.LString))
	return nil
}
