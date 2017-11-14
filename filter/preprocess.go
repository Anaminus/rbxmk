package filter

import (
	"bytes"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/luautil"
	"github.com/anaminus/rbxmk/types"
	"github.com/yuin/gopher-lua"
	"io"
	"strings"
)

func init() {
	Filters.Register(
		rbxmk.Filter{Name: "preprocess", Func: Preprocess},
	)
}

func Preprocess(f rbxmk.FilterArgs, opt rbxmk.Options, arguments []interface{}) (results []interface{}, err error) {
	value := arguments[0].(interface{})
	f.ProcessedArgs()
	out, err := ProcessStringlikeInterface(preprocessStringCallback(opt.Config.PreprocessorEnvs), value)
	if err != nil {
		return nil, err
	}
	return []interface{}{out}, nil
}

const putFuncName = "_put"

func preprocessStringCallback(envs []*lua.LTable) ProcessStringlikeCallback {
	return func(s *types.Stringlike) error {
		l := lua.NewState(lua.Options{SkipOpenLibs: true})
		// Parse preprocessors into Lua source.
		source := parsePreprocessors(s.Bytes, l.Load)

		// Load standard library and readonly environment.
		luautil.OpenFilteredLibs(l, luautil.GetFilteredStdLib())
		{
			varEnv := l.NewTable()
			for _, env := range envs {
				env.ForEach(func(k, v lua.LValue) {
					varEnv.RawSet(k, v)
				})
			}

			mt := l.CreateTable(0, 2)
			mt.RawSetString("__index", varEnv)
			mt.RawSetString("__metatable", lua.LString("The metatable is locked"))
			globals := l.Get(lua.GlobalsIndex).(*lua.LTable)
			l.SetMetatable(globals, mt)
		}

		// Add Put function.
		output := make([]string, 0, 64)
		l.SetGlobal(putFuncName, l.NewFunction(func(l *lua.LState) int {
			top := l.GetTop()
			for i := 1; i <= top; i++ {
				arg := l.Get(i)
				if arg == lua.LNil {
					continue
				}
				output = append(output, l.ToStringMeta(arg).String())
			}
			return 0
		}))
		fn, err := l.Load(bytes.NewReader(source), "<preprocess>")
		if err != nil {
			return err
		}
		l.Push(fn)
		if err := l.PCall(0, 0, nil); err != nil {
			return err
		}

		s.Bytes = []byte(strings.Join(output, ""))
		return nil
	}
}

func wrapText(pieces [][]byte, c, eq, text []byte) [][]byte {
	if len(text) > 0 {
		// "put[====[text]====]"
		pieces = append(pieces, []byte(putFuncName)) // "put"
		pieces = append(pieces, c[0:2])              // "[="
		if len(eq) > 0 {
			pieces = append(pieces, eq) // "="...
		}
		pieces = append(pieces, c[0:1]) // "["
		pieces = append(pieces, text)   // text
		pieces = append(pieces, c[2:3]) // "]"
		if len(eq) > 0 {
			pieces = append(pieces, eq) // "="...
		}
		pieces = append(pieces, c[1:3]) // "=]"
	}
	return pieces
}

func parsePreprocessors(input []byte, checkexp func(io.Reader, string) (*lua.LFunction, error)) (source []byte) {
	pieces := make([][]byte, 0)
	c := []byte("[=]()return ")
	eq, eqi := 0, 0
	h := 0
	for i := 0; i < len(input); i++ {
		switch b := input[i]; b {
		case '=':
			if i > 0 && input[i-1] != '=' {
				eq = 0
				eqi = i
			}
			eq++
		case '#':
			if i > 0 && input[i-1] != '\n' {
				continue
			}
			j := i
			for ; j < len(input); j++ {
				if input[j] == '\n' {
					j++
					break
				}
			}

			text := input[h:i]
			chunk := input[i+1 : j] // exclude '#'
			pieces = wrapText(pieces, c, input[eqi:eqi+eq], text)
			pieces = append(pieces, c[11:12]) // " "
			pieces = append(pieces, chunk)    // chunk
			pieces = append(pieces, c[11:12]) // " "
			h, i = j, j-1
			eq, eqi = 0, 0
		case '$':
			if i < len(input)-1 && input[i+1] != '(' {
				continue
			}
			j := i + 2
			n := 1
		loop:
			for ; j < len(input); j++ {
				switch b := input[j]; b {
				case '(':
					n++
				case ')':
					n--
					if n == 0 {
						j++
						break loop
					}
				}
			}
			if n > 0 {
				// Unmatched bracket.
				continue
			}

			// input[i:j] = $(...)
			// input[i+2:j-1] = ...
			text := input[h:i]
			chunk := input[i+2 : j-1]
			pieces = wrapText(pieces, c, input[eqi:eqi+eq], text)
			pieces = append(pieces, c[11:12]) // " "
			if _, err := checkexp(io.MultiReader(
				bytes.NewReader(c[5:]),
				bytes.NewReader(chunk),
			), "<exp>"); err == nil {
				// Write as expression list.
				pieces = append(pieces, []byte(putFuncName)) // "put"
				pieces = append(pieces, c[3:4])              // "("
				pieces = append(pieces, chunk)               // chunk
				pieces = append(pieces, c[4:5])              // ")"
			} else {
				// Write directly.
				pieces = append(pieces, chunk) // chunk
			}
			pieces = append(pieces, c[11:12]) // " "
			h, i = j, j-1
			eq, eqi = 0, 0
		}
	}
	if text := input[h:]; len(text) > 0 {
		pieces = wrapText(pieces, c, input[eqi:eqi+eq], text)
	}

	source = bytes.Join(pieces, nil)
	return
}
