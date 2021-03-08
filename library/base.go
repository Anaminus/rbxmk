package library

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
)

func init() { register(Base, -1) }

var Base = rbxmk.Library{Name: "", Open: openBase, Dump: dumpBase}

func openBase(s rbxmk.State) *lua.LTable {
	openFilteredLibs(s.L, filteredStdLib)
	return nil
}

type libFilter struct {
	Name     string
	OpenFunc lua.LGFunction
	Filter   map[lua.LValue]bool
}

var filteredStdLib = []libFilter{
	{lua.BaseLibName, lua.OpenBase, map[lua.LValue]bool{
		lua.LString("_G"):       true,
		lua.LString("_VERSION"): true,
		lua.LString("assert"):   true,
		lua.LString("error"):    true,
		lua.LString("ipairs"):   true,
		lua.LString("next"):     true,
		lua.LString("pairs"):    true,
		lua.LString("pcall"):    true,
		lua.LString("print"):    true,
		lua.LString("select"):   true,
		lua.LString("tonumber"): true,
		lua.LString("tostring"): true,
		lua.LString("type"):     true,
		lua.LString("unpack"):   true,
		lua.LString("xpcall"):   true,
		// lua.LString("collectgarbage"): true,
		// lua.LString("dofile"):         true,
		// lua.LString("getfenv"):        true,
		lua.LString("getmetatable"): true,
		// lua.LString("load"):           true,
		// lua.LString("loadfile"):       true,
		// lua.LString("loadstring"):     true,
		// lua.LString("module"):         true,
		// lua.LString("rawequal"):       true,
		// lua.LString("rawget"):         true,
		// lua.LString("rawset"):         true,
		// lua.LString("require"):        true,
		// lua.LString("setfenv"):        true,
		lua.LString("setmetatable"): true,
	}},
	// {lua.CoroutineLibName, lua.OpenCoroutine, map[lua.LValue]bool{
	// 	lua.LString("create"):  true,
	// 	lua.LString("resume"):  true,
	// 	lua.LString("running"): true,
	// 	lua.LString("status"):  true,
	// 	lua.LString("wrap"):    true,
	// 	lua.LString("yield"):   true,
	// }},
	// {lua.DebugLibName, lua.OpenDebug, map[lua.LValue]bool{
	// 	lua.LString("debug"):        true,
	// 	lua.LString("getfenv"):      true,
	// 	lua.LString("gethook"):      true,
	// 	lua.LString("getinfo"):      true,
	// 	lua.LString("getlocal"):     true,
	// 	lua.LString("getmetatable"): true,
	// 	lua.LString("getregistry"):  true,
	// 	lua.LString("getupvalue"):   true,
	// 	lua.LString("setfenv"):      true,
	// 	lua.LString("sethook"):      true,
	// 	lua.LString("setlocal"):     true,
	// 	lua.LString("setmetatable"): true,
	// 	lua.LString("setupvalue"):   true,
	// 	lua.LString("traceback"):    true,
	// }},
	// {lua.IoLibName, lua.OpenIo, map[lua.LValue]bool{
	// 	lua.LString("close"):   true,
	// 	lua.LString("flush"):   true,
	// 	lua.LString("input"):   true,
	// 	lua.LString("lines"):   true,
	// 	lua.LString("open"):    true,
	// 	lua.LString("output"):  true,
	// 	lua.LString("popen"):   true,
	// 	lua.LString("read"):    true,
	// 	lua.LString("stderr"):  true,
	// 	lua.LString("stdin"):   true,
	// 	lua.LString("stdout"):  true,
	// 	lua.LString("tmpfile"): true,
	// 	lua.LString("type"):    true,
	// 	lua.LString("write"):   true,
	// }},
	{lua.MathLibName, lua.OpenMath, map[lua.LValue]bool{
		lua.LString("abs"):        true,
		lua.LString("acos"):       true,
		lua.LString("asin"):       true,
		lua.LString("atan"):       true,
		lua.LString("atan2"):      true,
		lua.LString("ceil"):       true,
		lua.LString("cos"):        true,
		lua.LString("cosh"):       true,
		lua.LString("deg"):        true,
		lua.LString("exp"):        true,
		lua.LString("floor"):      true,
		lua.LString("fmod"):       true,
		lua.LString("frexp"):      true,
		lua.LString("huge"):       true,
		lua.LString("ldexp"):      true,
		lua.LString("log"):        true,
		lua.LString("log10"):      true,
		lua.LString("max"):        true,
		lua.LString("min"):        true,
		lua.LString("modf"):       true,
		lua.LString("pi"):         true,
		lua.LString("pow"):        true,
		lua.LString("rad"):        true,
		lua.LString("random"):     true,
		lua.LString("randomseed"): true,
		lua.LString("sin"):        true,
		lua.LString("sinh"):       true,
		lua.LString("sqrt"):       true,
		lua.LString("tan"):        true,
		lua.LString("tanh"):       true,
	}},
	{lua.OsLibName, lua.OpenOs, map[lua.LValue]bool{
		lua.LString("clock"):    true,
		lua.LString("date"):     true,
		lua.LString("difftime"): true,
		lua.LString("time"):     true,
		// lua.LString("execute"):   true,
		// lua.LString("exit"):      true,
		// lua.LString("getenv"):    true,
		// lua.LString("remove"):    true,
		// lua.LString("rename"):    true,
		// lua.LString("setlocale"): true,
		// lua.LString("tmpname"):   true,
	}},
	// {lua.LoadLibName, lua.OpenPackage, map[lua.LValue]bool{
	// 	lua.LString("cpath"):   true,
	// 	lua.LString("loaded"):  true,
	// 	lua.LString("loaders"): true,
	// 	lua.LString("loadlib"): true,
	// 	lua.LString("path"):    true,
	// 	lua.LString("preload"): true,
	// 	lua.LString("seeall"):  true,
	// }},
	{lua.StringLibName, lua.OpenString, map[lua.LValue]bool{
		lua.LString("byte"):    true,
		lua.LString("char"):    true,
		lua.LString("find"):    true,
		lua.LString("format"):  true,
		lua.LString("gmatch"):  true,
		lua.LString("gsub"):    true,
		lua.LString("len"):     true,
		lua.LString("lower"):   true,
		lua.LString("match"):   true,
		lua.LString("rep"):     true,
		lua.LString("reverse"): true,
		lua.LString("sub"):     true,
		lua.LString("upper"):   true,
		// lua.LString("dump"): true,
	}},
	{lua.TabLibName, lua.OpenTable, map[lua.LValue]bool{
		lua.LString("concat"): true,
		lua.LString("insert"): true,
		lua.LString("maxn"):   true,
		lua.LString("remove"): true,
		lua.LString("sort"):   true,
	}},
	// {lua.ChannelLibName, lua.OpenChannel, map[lua.LValue]bool{
	// 	lua.LString("make"):   true,
	// 	lua.LString("select"): true,
	// }},
}

func openFilteredLibs(l *lua.LState, libs []libFilter, upvalues ...lua.LValue) {
	for _, lib := range libs {
		l.Push(l.NewClosure(lib.OpenFunc, upvalues...))
		// LState.OpenLibs passes the library name as an argument for whatever
		// reason.
		l.Push(lua.LString(lib.Name))

		if lib.Filter == nil {
			l.Call(1, 0)
			continue
		}
		l.Call(1, 1)
		table := l.CheckTable(1)
		l.Pop(1)
		for k, _ := table.Next(lua.LNil); k != lua.LNil; k, _ = table.Next(k) {
			if !lib.Filter[k] {
				table.RawSet(k, lua.LNil)
			}
		}
	}
}

func dumpBase(s rbxmk.State) dump.Library {
	return dump.Library{
		Struct: dump.Struct{
			Fields: dump.Fields{
				"_G":       dump.Property{ValueType: dt.Prim("table")},
				"_VERSION": dump.Property{ValueType: dt.Prim("table")},
				"assert": dump.Function{
					Parameters: dump.Parameters{
						{Name: "v", Type: dt.Optional{T: dt.Prim("any")}},
						{Name: "message", Type: dt.Optional{T: dt.Prim("string")}, Default: `"assertion failed!"`},
					},
					CanError: true,
				},
				"error": dump.Function{
					Parameters: dump.Parameters{
						{Name: "message", Type: dt.Prim("any")},
						{Name: "level", Type: dt.Optional{T: dt.Prim("int")}, Default: `1`},
					},
					CanError: true,
				},
				"ipairs": dump.Function{
					Parameters: dump.Parameters{
						{Name: "t", Type: dt.Prim("table")},
					},
					Returns: dump.Parameters{
						{Name: "iterator", Type: dt.Prim("function")},
						{Name: "t", Type: dt.Prim("table")},
						{Name: "start", Type: dt.Prim("int")},
					},
				},
				"next": dump.Function{
					Parameters: dump.Parameters{
						{Name: "t", Type: dt.Prim("table")},
						{Name: "index", Type: dt.Optional{T: dt.Prim("any")}},
					},
					Returns: dump.Parameters{
						{Name: "index", Type: dt.Optional{T: dt.Prim("any")}},
						{Name: "value", Type: dt.Optional{T: dt.Prim("any")}},
					},
				},
				"pairs": dump.Function{
					Parameters: dump.Parameters{
						{Name: "t", Type: dt.Prim("table")},
					},
					Returns: dump.Parameters{
						{Name: "next", Type: dt.Prim("function")},
						{Name: "t", Type: dt.Prim("table")},
						{Name: "start", Type: dt.Prim("nil")},
					},
				},
				"pcall": dump.Function{
					Parameters: dump.Parameters{
						{Name: "f", Type: dt.Prim("function")},
						{Name: "...", Type: dt.Optional{T: dt.Prim("any")}},
					},
					Returns: dump.Parameters{
						{Name: "ok", Type: dt.Prim("boolean")},
						{Name: "...", Type: dt.Optional{T: dt.Prim("any")}},
					},
				},
				"print": dump.Function{
					Parameters: dump.Parameters{
						{Name: "...", Type: dt.Optional{T: dt.Prim("any")}},
					},
				},
				"select": dump.MultiFunction{
					{
						Parameters: dump.Parameters{
							{Name: "index", Type: dt.Prim("int")},
							{Name: "...", Type: dt.Optional{T: dt.Prim("any")}},
						},
						Returns: dump.Parameters{
							{Name: "...", Type: dt.Optional{T: dt.Prim("any")}},
						},
					},
					{
						Parameters: dump.Parameters{
							{Name: "count", Type: dt.Prim("string"), Enums: dt.Enums{`"#"`}},
							{Name: "...", Type: dt.Optional{T: dt.Prim("any")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("int")},
						},
					},
				},
				"tonumber": dump.MultiFunction{
					{
						Parameters: dump.Parameters{
							{Name: "x", Type: dt.Optional{T: dt.Prim("any")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("number")},
						},
					},
					{
						Parameters: dump.Parameters{
							{Name: "x", Type: dt.Prim("string")},
							{Name: "base", Type: dt.Prim("int")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("number")},
						},
					},
				},
				"tostring": dump.Function{
					Parameters: dump.Parameters{
						{Name: "v", Type: dt.Optional{T: dt.Prim("any")}},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("string")},
					},
				},
				"type": dump.Function{
					Parameters: dump.Parameters{
						{Name: "v", Type: dt.Prim("any")},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("string")},
					},
				},
				"unpack": dump.Function{
					Parameters: dump.Parameters{
						{Name: "list", Type: dt.Prim("table")},
						{Name: "i", Type: dt.Optional{T: dt.Prim("int")}},
						{Name: "j", Type: dt.Optional{T: dt.Prim("int")}},
					},
					Returns: dump.Parameters{
						{Name: "...", Type: dt.Optional{T: dt.Prim("any")}},
					},
				},
				"xpcall": dump.Function{
					Parameters: dump.Parameters{
						{Name: "f", Type: dt.Prim("function")},
						{Name: "msgh", Type: dt.Prim("function")},
						{Name: "...", Type: dt.Optional{T: dt.Prim("any")}},
					},
					Returns: dump.Parameters{
						{Name: "ok", Type: dt.Prim("boolean")},
						{Name: "...", Type: dt.Optional{T: dt.Prim("any")}},
					},
				},
				"getmetatable": dump.Function{
					Parameters: dump.Parameters{
						{Name: "v", Type: dt.Prim("any")},
					},
					Returns: dump.Parameters{
						{Type: dt.Optional{T: dt.Prim("table")}},
					},
				},
				"setmetatable": dump.Function{
					Parameters: dump.Parameters{
						{Name: "v", Type: dt.Prim("any")},
						{Name: "metatable", Type: dt.Optional{T: dt.Prim("table")}},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("table")},
					},
				},
				"math": dump.Struct{
					Fields: dump.Fields{
						"abs": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"acos": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"asin": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"atan": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"atan2": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
								{Name: "y", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"ceil": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("int")},
							},
						},
						"cos": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"cosh": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"deg": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"exp": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"floor": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("int")},
							},
						},
						"fmod": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
								{Name: "y", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"frexp": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Name: "m", Type: dt.Prim("number")},
								{Name: "e", Type: dt.Prim("int")},
							},
						},
						"huge": dump.Property{ValueType: dt.Prim("number")},
						"ldexp": dump.Function{
							Parameters: dump.Parameters{
								{Name: "m", Type: dt.Prim("number")},
								{Name: "e", Type: dt.Prim("int")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"log": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"log10": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"max": dump.Function{
							Parameters: dump.Parameters{
								{Name: "...", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"min": dump.Function{
							Parameters: dump.Parameters{
								{Name: "...", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"modf": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("int")},
								{Type: dt.Prim("number")},
							},
						},
						"pi": dump.Property{ValueType: dt.Prim("number")},
						"pow": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"rad": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"random": dump.MultiFunction{
							{
								Returns: dump.Parameters{
									{Type: dt.Prim("number")},
								},
							},
							{
								Parameters: dump.Parameters{
									{Name: "m", Type: dt.Prim("int")},
								},
								Returns: dump.Parameters{
									{Type: dt.Prim("number")},
								},
							},
							{
								Parameters: dump.Parameters{
									{Name: "m", Type: dt.Prim("int")},
									{Name: "n", Type: dt.Prim("int")},
								},
								Returns: dump.Parameters{
									{Type: dt.Prim("number")},
								},
								CanError: true,
							},
						},
						"randomseed": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
						},
						"sin": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"sinh": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"sqrt": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"tan": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"tanh": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
					},
				},
				"os": dump.Struct{
					Fields: dump.Fields{
						"clock": dump.Function{
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"date": dump.MultiFunction{
							{
								Returns: dump.Parameters{
									{Type: dt.Prim("string")},
								},
							},
							{
								Parameters: dump.Parameters{
									{Name: "format", Type: dt.Prim("string"), Enums: dt.Enums{`"*t"`, `!*t`}},
									{Name: "time", Type: dt.Optional{T: dt.Prim("number")}},
								},
								Returns: dump.Parameters{
									{Type: dt.Optional{T: dt.Struct{
										"year":  dt.Prim("int"),
										"month": dt.Prim("int"),
										"day":   dt.Prim("int"),
										"hour":  dt.Optional{T: dt.Prim("int")},
										"min":   dt.Optional{T: dt.Prim("int")},
										"sec":   dt.Optional{T: dt.Prim("int")},
										"wday":  dt.Optional{T: dt.Prim("int")},
										"yday":  dt.Optional{T: dt.Prim("int")},
										"isdst": dt.Optional{T: dt.Prim("boolean")},
									}}},
								},
							},
							{
								Parameters: dump.Parameters{
									{Name: "format", Type: dt.Prim("string")},
									{Name: "time", Type: dt.Optional{T: dt.Prim("number")}},
								},
								Returns: dump.Parameters{
									{Type: dt.Prim("string")},
								},
							},
						},
						"difftime": dump.Function{
							Parameters: dump.Parameters{
								{Name: "t2", Type: dt.Prim("number")},
								{Name: "t1", Type: dt.Prim("number")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
						"time": dump.Function{
							Parameters: dump.Parameters{
								{Name: "t", Type: dt.Optional{T: dt.Struct{
									"year":  dt.Prim("int"),
									"month": dt.Prim("int"),
									"day":   dt.Prim("int"),
									"hour":  dt.Optional{T: dt.Prim("int")},
									"min":   dt.Optional{T: dt.Prim("int")},
									"sec":   dt.Optional{T: dt.Prim("int")},
									"isdst": dt.Optional{T: dt.Prim("boolean")},
								}}},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("number")},
							},
						},
					},
				},
				"string": dump.Struct{
					Fields: dump.Fields{
						"byte": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim("string")},
								{Name: "i", Type: dt.Optional{T: dt.Prim("int")}, Default: `1`},
								{Name: "j", Type: dt.Optional{T: dt.Prim("int")}, Default: `i`},
							},
							Returns: dump.Parameters{
								{Name: "...", Type: dt.Prim("int")},
							},
						},
						"char": dump.Function{
							Parameters: dump.Parameters{
								{Name: "...", Type: dt.Prim("int")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("string")},
							},
						},
						"find": dump.MultiFunction{
							{
								Parameters: dump.Parameters{
									{Name: "s", Type: dt.Prim("string")},
									{Name: "pattern", Type: dt.Prim("string")},
									{Name: "init", Type: dt.Optional{T: dt.Prim("int")}, Default: `1`},
								},
								Returns: dump.Parameters{
									{Name: "start", Type: dt.Optional{T: dt.Prim("number")}},
									{Name: "end", Type: dt.Optional{T: dt.Prim("number")}},
								},
								CanError: true,
							},
							{
								Parameters: dump.Parameters{
									{Name: "s", Type: dt.Prim("string")},
									{Name: "pattern", Type: dt.Prim("string")},
									{Name: "init", Type: dt.Prim("int")},
									{Name: "plain", Type: dt.Optional{T: dt.Prim("boolean")}, Default: `false`},
								},
								Returns: dump.Parameters{
									{Name: "start", Type: dt.Optional{T: dt.Prim("number")}},
									{Name: "end", Type: dt.Optional{T: dt.Prim("number")}},
								},
								CanError: true,
							},
						},
						"format": dump.Function{
							Parameters: dump.Parameters{
								{Name: "format", Type: dt.Prim("string")},
								{Name: "...", Type: dt.Prim("any")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("string")},
							},
							CanError: true,
						},
						"gmatch": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim("string")},
								{Name: "pattern", Type: dt.Prim("string")},
							},
							Returns: dump.Parameters{
								{Type: dt.Function{
									Returns: dump.Parameters{
										{Name: "...", Type: dt.Prim("string")},
									},
								}},
							},
							CanError: true,
						},
						"gsub": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim("string")},
								{Name: "pattern", Type: dt.Prim("string")},
								{Name: "repl", Type: dt.Or{
									dt.Prim("string"),
									dt.Map{
										K: dt.Prim("string"),
										V: dt.Or{dt.Prim("string"), dt.Prim("number"), dt.Prim("false")},
									},
									dt.Function{
										Parameters: dump.Parameters{
											{Name: "...", Type: dt.Prim("string")},
										},
										Returns: dump.Parameters{
											{Type: dt.Or{dt.Prim("string"), dt.Prim("number"), dt.Prim("false"), dt.Prim("nil")}},
										},
									},
								}},
								{Name: "n", Type: dt.Optional{T: dt.Prim("int")}},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("string")},
								{Type: dt.Prim("int")},
							},
							CanError: true,
						},
						"len": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim("string")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("int")},
							},
						},
						"lower": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim("string")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("string")},
							},
						},
						"match": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim("string")},
								{Name: "pattern", Type: dt.Prim("string")},
								{Name: "init", Type: dt.Optional{T: dt.Prim("int")}, Default: `1`},
							},
							Returns: dump.Parameters{
								{Name: "...", Type: dt.Optional{T: dt.Prim("string")}},
							},
							CanError: true,
						},
						"rep": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim("string")},
								{Name: "n", Type: dt.Prim("int")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("string")},
							},
						},
						"reverse": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim("string")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("string")},
							},
						},
						"sub": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim("string")},
								{Name: "i", Type: dt.Prim("int")},
								{Name: "j", Type: dt.Optional{T: dt.Prim("int")}, Default: `-1`},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("string")},
							},
						},
						"upper": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim("string")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("string")},
							},
						},
					},
				},
				"table": dump.Struct{
					Fields: dump.Fields{
						"concat": dump.Function{
							Parameters: dump.Parameters{
								{Name: "t", Type: dt.Array{T: dt.Or{dt.Prim("string"), dt.Prim("number")}}},
								{Name: "sep", Type: dt.Optional{T: dt.Prim("string")}, Default: `""`},
								{Name: "i", Type: dt.Optional{T: dt.Prim("int")}, Default: `1`},
								{Name: "j", Type: dt.Optional{T: dt.Prim("int")}, Default: `#t`},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("string")},
							},
							CanError: true,
						},
						"insert": dump.MultiFunction{
							{
								Parameters: dump.Parameters{
									{Name: "t", Type: dt.Prim("table")},
									{Name: "index", Type: dt.Prim("int")},
									{Name: "value", Type: dt.Prim("any")},
								},
							},
							{
								Parameters: dump.Parameters{
									{Name: "t", Type: dt.Prim("table")},
									{Name: "value", Type: dt.Prim("any")},
								},
							},
						},
						"maxn": dump.Function{
							Parameters: dump.Parameters{
								{Name: "t", Type: dt.Prim("table")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("int")},
							},
						},
						"remove": dump.Function{
							Parameters: dump.Parameters{
								{Name: "t", Type: dt.Prim("table")},
								{Name: "index", Type: dt.Optional{T: dt.Prim("int")}, Default: `#t`},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("any")},
							},
						},
						"sort": dump.Function{
							Parameters: dump.Parameters{
								{Name: "t", Type: dt.Prim("table")},
								{Name: "comp", Type: dt.Optional{T: dt.Function{
									Parameters: dump.Parameters{
										{Name: "a", Type: dt.Prim("any")},
										{Name: "b", Type: dt.Prim("any")},
									},
									Returns: dump.Parameters{
										{Type: dt.Prim("boolean")},
									},
								}}},
							},
							CanError: true,
						},
					},
				},
			},
		},
	}
}
