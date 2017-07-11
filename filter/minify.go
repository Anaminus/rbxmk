package filter

import (
	"fmt"
	"github.com/anaminus/rbxmk"
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
	data := arguments[0].(rbxmk.Data)
	f.ProcessedArgs()
	v, err := doMinify("minify", data)
	return []interface{}{v}, err
}

func Unminify(f rbxmk.FilterArgs, opt rbxmk.Options, arguments []interface{}) (results []interface{}, err error) {
	data := arguments[0].(rbxmk.Data)
	f.ProcessedArgs()
	v, err := doMinify("unminify", data)
	return []interface{}{v}, err
}

func doMinify(method string, data rbxmk.Data) (out rbxmk.Data, err error) {
	switch v := data.(type) {
	case *[]*rbxfile.Instance:
		for _, inst := range *v {
			if err := doMinifyInstance(method, inst, false); err != nil {
				return nil, err
			}
		}
	case *rbxfile.Instance:
		err = doMinifyInstance(method, v, true)
	case rbxfile.Value:
		out, err = doMinifyValue(method, v)
	case string:
		out, err = doMinifyString(method, v)
	case nil:
	default:
		return nil, rbxmk.NewDataTypeError(data)
	}
	return out, err
}

func doMinifyInstance(method string, inst *rbxfile.Instance, fail bool) (err error) {
	switch inst.ClassName {
	case "Script", "LocalScript", "ModuleScript":
		if source, ok := inst.Properties["Source"]; ok {
			inst.Properties["Source"], err = doMinifyValue(method, source)
		}
		return nil
	}
	if fail {
		return fmt.Errorf("instance must be script-like")
	}
	return nil
}

func doMinifyValue(method string, value rbxfile.Value) (rbxfile.Value, error) {
	switch v := value.(type) {
	case rbxfile.ValueString:
		s, err := doMinifyString(method, string(v))
		return rbxfile.ValueString(s), err
	case rbxfile.ValueBinaryString:
		s, err := doMinifyString(method, string(v))
		return rbxfile.ValueBinaryString(s), err
	case rbxfile.ValueProtectedString:
		s, err := doMinifyString(method, string(v))
		return rbxfile.ValueProtectedString(s), err
	}
	return nil, fmt.Errorf("value must be string-like")
}

func doMinifyString(method string, s string) (string, error) {
	var l *lua.LState
	{
		l = lua.NewState()
		src := minify_lua()
		b, _ := ioutil.ReadAll(src)
		src.Close()
		fn, err := l.LoadString(string(b))
		if err != nil {
			return "", err
		}
		l.Push(fn)
	}
	l.Push(lua.LString(method))
	l.Push(lua.LString(s))
	if err := l.PCall(2, 1, nil); err != nil {
		return "", err
	}
	return string(l.Get(-1).(lua.LString)), nil
}
