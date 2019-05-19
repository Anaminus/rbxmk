package filter

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/types"
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

func Minify(f rbxmk.FilterArgs, opt *rbxmk.Options, arguments []interface{}) (results []interface{}, err error) {
	value := arguments[0].(interface{})
	f.ProcessedArgs()
	out, err := types.AsString(minifyStringCallback("minify"), value)
	if err != nil {
		return nil, err
	}
	return []interface{}{out}, nil
}

func Unminify(f rbxmk.FilterArgs, opt *rbxmk.Options, arguments []interface{}) (results []interface{}, err error) {
	value := arguments[0].(interface{})
	f.ProcessedArgs()
	out, err := types.AsString(minifyStringCallback("unminify"), value)
	if err != nil {
		return nil, err
	}
	return []interface{}{out}, nil
}

func minifyStringCallback(method string) types.AsStringCallback {
	return func(s *types.Stringlike) error {
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
}
