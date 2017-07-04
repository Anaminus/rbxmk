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

func Minify(f rbxmk.FilterArgs, opt rbxmk.Options, arguments []interface{}) (results []interface{}) {
	data := arguments[0].(rbxmk.Data)
	f.ProcessedArgs()
	return []interface{}{doMinify("minify", data)}
}

func Unminify(f rbxmk.FilterArgs, opt rbxmk.Options, arguments []interface{}) (results []interface{}) {
	data := arguments[0].(rbxmk.Data)
	f.ProcessedArgs()
	return []interface{}{doMinify("unminify", data)}
}

func doMinify(method string, data rbxmk.Data) rbxmk.Data {
	switch v := data.(type) {
	case *[]*rbxfile.Instance:
		for _, inst := range *v {
			doMinifyInstance(method, inst, false)
		}
	case *rbxfile.Instance:
		doMinifyInstance(method, v, true)
	case rbxfile.Value:
		data = doMinifyValue(method, v)
	case string:
		data = doMinifyString(method, v)
	case nil:
	default:
		panic(rbxmk.NewDataTypeError(data))
	}
	return data
}

func doMinifyInstance(method string, inst *rbxfile.Instance, fail bool) {
	switch inst.ClassName {
	case "Script", "LocalScript", "ModuleScript":
		if source, ok := inst.Properties["Source"]; ok {
			inst.Properties["Source"] = doMinifyValue(method, source)
		}
		return
	}
	if fail {
		panic(fmt.Errorf("instance must be script-like"))
	}
}

func doMinifyValue(method string, value rbxfile.Value) rbxfile.Value {
	switch v := value.(type) {
	case rbxfile.ValueString:
		return rbxfile.ValueString(doMinifyString(method, string(v)))
	case rbxfile.ValueBinaryString:
		return rbxfile.ValueBinaryString(doMinifyString(method, string(v)))
	case rbxfile.ValueProtectedString:
		return rbxfile.ValueProtectedString(doMinifyString(method, string(v)))
	}
	panic(fmt.Errorf("value must be string-like"))
}

func doMinifyString(method string, s string) string {
	var l *lua.LState
	{
		l = lua.NewState()
		src := minify_lua()
		b, _ := ioutil.ReadAll(src)
		src.Close()
		fn, err := l.LoadString(string(b))
		if err != nil {
			panic(err)
		}
		l.Push(fn)
	}
	l.Push(lua.LString(method))
	l.Push(lua.LString(s))
	if err := l.PCall(2, 1, nil); err != nil {
		panic(err)
	}
	return string(l.Get(-1).(lua.LString))
}
