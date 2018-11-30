package main

import (
	"bytes"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rbxmk/luautil"
	"github.com/anaminus/rbxmk/types"
	"github.com/yuin/gopher-lua"
	"io"
	"strings"
)

// Order of preprocessor variable environments.
const (
	PPEnvScript  = iota // Defined via script (rbxmk.configure).
	PPEnvCommand        // Defined via --define option.
	PPEnvLen            // Number of environments.
)

func Preprocess(f rbxmk.FilterArgs, opt *rbxmk.Options, arguments []interface{}) (results []interface{}, err error) {
	value := arguments[0].(interface{})
	f.ProcessedArgs()
	envs, _ := opt.Config["PPEnvs"].([]*lua.LTable)
	out, err := types.ProcessStringlikeInterface(preprocessStringCallback(envs), value)
	if err != nil {
		return nil, err
	}
	return []interface{}{out}, nil
}

const putFuncName = "_put"

func preprocessStringCallback(envs []*lua.LTable) types.ProcessStringlikeCallback {
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
					if v == luautil.ForceNil {
						varEnv.RawSet(k, lua.LNil)
					} else {
						varEnv.RawSet(k, v)
					}
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
		fn, err := l.Load(strings.NewReader(source), "<preprocess>")
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

// Parse a Lua comment. Starts after `--` characters.
func parseComment(b []byte, i int) (j, k int) {
	var c rune
	k = i
	next := func() {
		if k+1 < len(b) {
			k++
			c = rune(b[k])
		} else {
			k = len(b)
			c = -1
		}
	}
	next()
	if c == '[' {
		next()
		eq := 0
		for c == '=' {
			eq++
			next()
		}
		if c != '[' {
			return -1, -1
		}
		next()
		j = k
	loop:
		for {
			if c == -1 {
				return -1, -1
			}
			if c == ']' {
				next()
				for i := 0; i < eq; i++ {
					if c != '=' {
						continue loop
					}
					next()
				}
				if c == ']' {
					next()
					break
				}
			}
			next()
		}
		return j, k
	}
	j = k
	for c != '\n' && c != -1 {
		next()
	}
	if c != -1 {
		k++
	}
	return j, k
}

// Find shortest closing bracket not in the string.
func shortestEnclosingBracket(b []byte) (eq int) {
loop:
	for i := 0; i < len(b); i++ {
		if b[i] == ']' {
			i++
			count := 0
			for ; b[i] == '='; i++ {
				count++
			}
			if b[i] == ']' && count == eq {
				eq++
				goto loop
			}
		}
	}
	return eq
}

func wrapText(builder *strings.Builder, text []byte) {
	if len(text) == 0 {
		return
	}
	// `_put[====[text]====]`
	// `_put("\n"..[====[text]====])`
	builder.WriteString(putFuncName)
	if text[0] == '\n' {
		// Add back newline truncated by Lua string literal. Append it
		// before literal so that line numbers don't get screwed up.
		builder.WriteString(`("\n"..`)
	}
	eq := shortestEnclosingBracket(text)
	builder.WriteByte('[')
	for i := 0; i < eq; i++ {
		builder.WriteByte('=')
	}
	builder.WriteByte('[')
	builder.Write(text)
	builder.WriteByte(']')
	for i := 0; i < eq; i++ {
		builder.WriteByte('=')
	}
	builder.WriteByte(']')
	if text[0] == '\n' {
		builder.WriteByte(')')
	}
}

func parsePreprocessors(input []byte, checkexp func(io.Reader, string) (*lua.LFunction, error)) (source string) {
	var builder strings.Builder
	h := 0

	for i := 0; i < len(input)-2; i++ {
		if input[i] != '-' || input[i+1] != '-' {
			continue
		}
		j, k := parseComment(input, i+1)
		if j < 0 {
			continue
		}
		if input[j] == '#' {
			wrapText(&builder, input[h:i])
			if j-i > 2 {
				// Long comment.
				chunk := input[j+1 : k-(j-i-2)]
				builder.WriteByte(' ')
				if _, err := checkexp(io.MultiReader(
					strings.NewReader("return "),
					bytes.NewReader(chunk),
				), "<exp>"); err == nil {
					// Write as expression list.
					builder.WriteString(putFuncName)
					builder.WriteByte('(')
					builder.Write(chunk)
					builder.WriteByte(')')
				} else {
					// Write directly.
					builder.Write(chunk)
				}
				builder.WriteByte(' ')
			} else {
				// Comment.
				builder.WriteByte(' ')
				builder.Write(input[j+1 : k])
				builder.WriteByte(' ')
			}
		} else {
			continue
		}
		h, i = k, k-1
	}
	wrapText(&builder, input[h:])
	return builder.String()
}
