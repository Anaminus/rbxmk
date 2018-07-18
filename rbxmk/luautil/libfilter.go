package luautil

import (
	"github.com/yuin/gopher-lua"
)

type LibFilter struct {
	Name     string
	OpenFunc lua.LGFunction
	Filter   map[lua.LValue]bool
}

func GetFilteredStdLib() []LibFilter {
	return []LibFilter{
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
			// lua.LString("getmetatable"):   true,
			// lua.LString("load"):           true,
			// lua.LString("loadfile"):       true,
			// lua.LString("loadstring"):     true,
			// lua.LString("module"):         true,
			// lua.LString("rawequal"):       true,
			// lua.LString("rawget"):         true,
			// lua.LString("rawset"):         true,
			// lua.LString("require"):        true,
			// lua.LString("setfenv"):        true,
			// lua.LString("setmetatable"):   true,
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
}

func OpenFilteredLibs(l *lua.LState, libs []LibFilter, upvalues ...lua.LValue) {
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
