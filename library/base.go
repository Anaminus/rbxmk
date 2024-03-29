package library

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
)

func init() { register(Base) }

var Base = rbxmk.Library{
	Name:     "base",
	Priority: -1,
	Open:     openBase,
	Dump:     dumpBase,
}

func openBase(s rbxmk.State) *lua.LTable {
	openFilteredLibs(s, filteredStdLib)
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
		lua.LString("abs"):   true,
		lua.LString("acos"):  true,
		lua.LString("asin"):  true,
		lua.LString("atan"):  true,
		lua.LString("atan2"): true,
		lua.LString("ceil"):  true,
		lua.LString("cos"):   true,
		lua.LString("cosh"):  true,
		lua.LString("deg"):   true,
		lua.LString("exp"):   true,
		lua.LString("floor"): true,
		lua.LString("fmod"):  true,
		lua.LString("frexp"): true,
		lua.LString("huge"):  true,
		lua.LString("ldexp"): true,
		// lua.LString("log"):        true,
		// lua.LString("log10"):      true,
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

func openFilteredLibs(s rbxmk.State, libs []libFilter, upvalues ...lua.LValue) {
	for _, lib := range libs {
		s.L.Push(s.L.NewClosure(lib.OpenFunc, upvalues...))
		// LState.OpenLibs passes the library name as an argument for whatever
		// reason.
		s.L.Push(lua.LString(lib.Name))

		if lib.Filter == nil {
			s.L.Call(1, 0)
			continue
		}
		s.L.Call(1, 1)
		table := s.L.CheckTable(1)
		s.L.Pop(1)
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
				"_G": dump.Property{
					ValueType:   dt.Prim(rtypes.T_LuaTable),
					ReadOnly:    true,
					Summary:     "Libraries/base:Fields/_G/Summary",
					Description: "Libraries/base:Fields/_G/Description",
				},
				"_VERSION": dump.Property{
					ValueType:   dt.Prim(rtypes.T_LuaString),
					ReadOnly:    true,
					Summary:     "Libraries/base:Fields/_VERSION/Summary",
					Description: "Libraries/base:Fields/_VERSION/Description",
				},
				"assert": dump.Function{
					Parameters: dump.Parameters{
						{Name: "v", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
						{Name: "message", Type: dt.Optional(dt.Prim(rtypes.T_LuaString)), Default: `"assertion failed!"`},
						{Name: "...", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
					},
					Returns: dump.Parameters{
						{Name: "...", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
					},
					CanError:    true,
					Summary:     "Libraries/base:Fields/assert/Summary",
					Description: "Libraries/base:Fields/assert/Description",
				},
				"error": dump.Function{
					Parameters: dump.Parameters{
						{Name: "message", Type: dt.Prim(rtypes.T_Any)},
						{Name: "level", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger)), Default: `1`},
					},
					CanError:    true,
					Summary:     "Libraries/base:Fields/error/Summary",
					Description: "Libraries/base:Fields/error/Description",
				},
				"ipairs": dump.Function{
					Parameters: dump.Parameters{
						{Name: "t", Type: dt.Prim(rtypes.T_LuaTable)},
					},
					Returns: dump.Parameters{
						{Name: "iterator", Type: dt.Prim(rtypes.T_LuaFunction)},
						{Name: "t", Type: dt.Prim(rtypes.T_LuaTable)},
						{Name: "start", Type: dt.Prim(rtypes.T_LuaInteger)},
					},
					Summary:     "Libraries/base:Fields/ipairs/Summary",
					Description: "Libraries/base:Fields/ipairs/Description",
				},
				"next": dump.Function{
					Parameters: dump.Parameters{
						{Name: "t", Type: dt.Prim(rtypes.T_LuaTable)},
						{Name: "index", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
					},
					Returns: dump.Parameters{
						{Name: "index", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
						{Name: "value", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
					},
					Summary:     "Libraries/base:Fields/next/Summary",
					Description: "Libraries/base:Fields/next/Description",
				},
				"pairs": dump.Function{
					Parameters: dump.Parameters{
						{Name: "t", Type: dt.Prim(rtypes.T_LuaTable)},
					},
					Returns: dump.Parameters{
						{Name: "next", Type: dt.Prim(rtypes.T_LuaFunction)},
						{Name: "t", Type: dt.Prim(rtypes.T_LuaTable)},
						{Name: "start", Type: dt.Prim(rtypes.T_LuaNil)},
					},
					Summary:     "Libraries/base:Fields/pairs/Summary",
					Description: "Libraries/base:Fields/pairs/Description",
				},
				"pcall": dump.Function{
					Parameters: dump.Parameters{
						{Name: "f", Type: dt.Prim(rtypes.T_LuaFunction)},
						{Name: "...", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
					},
					Returns: dump.Parameters{
						{Name: "ok", Type: dt.Prim(rtypes.T_LuaBoolean)},
						{Name: "...", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
					},
					Summary:     "Libraries/base:Fields/pcall/Summary",
					Description: "Libraries/base:Fields/pcall/Description",
				},
				"print": dump.Function{
					Parameters: dump.Parameters{
						{Name: "...", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
					},
					Summary:     "Libraries/base:Fields/print/Summary",
					Description: "Libraries/base:Fields/print/Description",
				},
				"select": dump.MultiFunction{
					{
						Parameters: dump.Parameters{
							{Name: "index", Type: dt.Prim(rtypes.T_LuaInteger)},
							{Name: "...", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
						},
						Returns: dump.Parameters{
							{Name: "...", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
						},
						CanError:    true,
						Summary:     "Libraries/base:Fields/select/Index/Summary",
						Description: "Libraries/base:Fields/select/Index/Description",
					},
					{
						Parameters: dump.Parameters{
							{Name: "count", Type: dt.Prim(rtypes.T_LuaString), Enums: dt.Enums{`"#"`}},
							{Name: "...", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(rtypes.T_LuaInteger)},
						},
						CanError:    true,
						Summary:     "Libraries/base:Fields/select/Count/Summary",
						Description: "Libraries/base:Fields/select/Count/Description",
					},
				},
				"tonumber": dump.Function{
					Parameters: dump.Parameters{
						{Name: "x", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
						{Name: "base", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger)), Default: `10`},
					},
					Returns: dump.Parameters{
						{Type: dt.Optional(dt.Prim(rtypes.T_LuaNumber))},
					},
					Summary:     "Libraries/base:Fields/tonumber/Summary",
					Description: "Libraries/base:Fields/tonumber/Description",
				},
				"tostring": dump.Function{
					Parameters: dump.Parameters{
						{Name: "v", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_LuaString)},
					},
					Summary:     "Libraries/base:Fields/tostring/Summary",
					Description: "Libraries/base:Fields/tostring/Description",
				},
				"type": dump.Function{
					Parameters: dump.Parameters{
						{Name: "v", Type: dt.Prim(rtypes.T_Any)},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_LuaString)},
					},
					Summary:     "Libraries/base:Fields/type/Summary",
					Description: "Libraries/base:Fields/type/Description",
				},
				"unpack": dump.Function{
					Parameters: dump.Parameters{
						{Name: "list", Type: dt.Prim(rtypes.T_LuaTable)},
						{Name: "i", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger))},
						{Name: "j", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger))},
					},
					Returns: dump.Parameters{
						{Name: "...", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
					},
					Summary:     "Libraries/base:Fields/unpack/Summary",
					Description: "Libraries/base:Fields/unpack/Description",
				},
				"xpcall": dump.Function{
					Parameters: dump.Parameters{
						{Name: "f", Type: dt.Prim(rtypes.T_LuaFunction)},
						{Name: "msgh", Type: dt.Function(dt.KindFunction{
							Parameters: []dt.Parameter{{Name: "err", Type: dt.Prim(rtypes.T_Any)}},
							Returns:    []dt.Parameter{{Type: dt.Prim(rtypes.T_Any)}},
						})},
						{Name: "...", Type: dt.Prim(rtypes.T_Any)},
					},
					Returns: dump.Parameters{
						{Name: "ok", Type: dt.Prim(rtypes.T_LuaBoolean)},
						{Name: "...", Type: dt.Optional(dt.Prim(rtypes.T_Any))},
					},
					Summary:     "Libraries/base:Fields/xpcall/Summary",
					Description: "Libraries/base:Fields/xpcall/Description",
				},
				"getmetatable": dump.Function{
					Parameters: dump.Parameters{
						{Name: "v", Type: dt.Prim(rtypes.T_Any)},
					},
					Returns: dump.Parameters{
						{Type: dt.Optional(dt.Prim(rtypes.T_LuaTable))},
					},
					Summary:     "Libraries/base:Fields/getmetatable/Summary",
					Description: "Libraries/base:Fields/getmetatable/Description",
				},
				"setmetatable": dump.Function{
					Parameters: dump.Parameters{
						{Name: "v", Type: dt.Prim(rtypes.T_LuaTable)},
						{Name: "metatable", Type: dt.Optional(dt.Prim(rtypes.T_LuaTable))},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_LuaTable)},
					},
					CanError:    true,
					Summary:     "Libraries/base:Fields/setmetatable/Summary",
					Description: "Libraries/base:Fields/setmetatable/Description",
				},
				"math": dump.Struct{
					Fields: dump.Fields{
						"abs": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/abs/Summary",
							Description: "Libraries/base/Fields/math:Fields/abs/Description",
						},
						"acos": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/acos/Summary",
							Description: "Libraries/base/Fields/math:Fields/acos/Description",
						},
						"asin": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/asin/Summary",
							Description: "Libraries/base/Fields/math:Fields/asin/Description",
						},
						"atan": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/atan/Summary",
							Description: "Libraries/base/Fields/math:Fields/atan/Description",
						},
						"atan2": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
								{Name: "y", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/atan2/Summary",
							Description: "Libraries/base/Fields/math:Fields/atan2/Description",
						},
						"ceil": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaInteger)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/ceil/Summary",
							Description: "Libraries/base/Fields/math:Fields/ceil/Description",
						},
						"cos": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/cos/Summary",
							Description: "Libraries/base/Fields/math:Fields/cos/Description",
						},
						"cosh": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/cosh/Summary",
							Description: "Libraries/base/Fields/math:Fields/cosh/Description",
						},
						"deg": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/deg/Summary",
							Description: "Libraries/base/Fields/math:Fields/deg/Description",
						},
						"exp": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/exp/Summary",
							Description: "Libraries/base/Fields/math:Fields/exp/Description",
						},
						"floor": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaInteger)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/floor/Summary",
							Description: "Libraries/base/Fields/math:Fields/floor/Description",
						},
						"fmod": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
								{Name: "y", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/fmod/Summary",
							Description: "Libraries/base/Fields/math:Fields/fmod/Description",
						},
						"frexp": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Name: "m", Type: dt.Prim(rtypes.T_LuaNumber)},
								{Name: "e", Type: dt.Prim(rtypes.T_LuaInteger)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/frexp/Summary",
							Description: "Libraries/base/Fields/math:Fields/frexp/Description",
						},
						"huge": dump.Property{
							ValueType:   dt.Prim(rtypes.T_LuaNumber),
							ReadOnly:    true,
							Summary:     "Libraries/base/Fields/math:Fields/huge/Summary",
							Description: "Libraries/base/Fields/math:Fields/huge/Description",
						},
						"ldexp": dump.Function{
							Parameters: dump.Parameters{
								{Name: "m", Type: dt.Prim(rtypes.T_LuaNumber)},
								{Name: "e", Type: dt.Prim(rtypes.T_LuaInteger)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/ldexp/Summary",
							Description: "Libraries/base/Fields/math:Fields/ldexp/Description",
						},
						"max": dump.Function{
							Parameters: dump.Parameters{
								{Name: "...", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/max/Summary",
							Description: "Libraries/base/Fields/math:Fields/max/Description",
						},
						"min": dump.Function{
							Parameters: dump.Parameters{
								{Name: "...", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/min/Summary",
							Description: "Libraries/base/Fields/math:Fields/min/Description",
						},
						"modf": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaInteger)},
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/modf/Summary",
							Description: "Libraries/base/Fields/math:Fields/modf/Description",
						},
						"pi": dump.Property{
							ValueType:   dt.Prim(rtypes.T_LuaNumber),
							ReadOnly:    true,
							Summary:     "Libraries/base/Fields/math:Fields/pi/Summary",
							Description: "Libraries/base/Fields/math:Fields/pi/Description",
						},
						"pow": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/pow/Summary",
							Description: "Libraries/base/Fields/math:Fields/pow/Description",
						},
						"rad": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/rad/Summary",
							Description: "Libraries/base/Fields/math:Fields/rad/Description",
						},
						"random": dump.MultiFunction{
							{
								Returns: dump.Parameters{
									{Type: dt.Prim(rtypes.T_LuaNumber)},
								},
								Summary:     "Libraries/base/Fields/math:Fields/random/Real/Summary",
								Description: "Libraries/base/Fields/math:Fields/random/Real/Description",
							},
							{
								Parameters: dump.Parameters{
									{Name: "m", Type: dt.Prim(rtypes.T_LuaInteger)},
								},
								Returns: dump.Parameters{
									{Type: dt.Prim(rtypes.T_LuaNumber)},
								},
								Summary:     "Libraries/base/Fields/math:Fields/random/Range/Summary",
								Description: "Libraries/base/Fields/math:Fields/random/Range/Description",
							},
							{
								Parameters: dump.Parameters{
									{Name: "m", Type: dt.Prim(rtypes.T_LuaInteger)},
									{Name: "n", Type: dt.Prim(rtypes.T_LuaInteger)},
								},
								Returns: dump.Parameters{
									{Type: dt.Prim(rtypes.T_LuaNumber)},
								},
								CanError:    true,
								Summary:     "Libraries/base/Fields/math:Fields/random/Interval/Summary",
								Description: "Libraries/base/Fields/math:Fields/random/Interval/Description",
							},
						},
						"randomseed": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/randomseed/Summary",
							Description: "Libraries/base/Fields/math:Fields/randomseed/Description",
						},
						"sin": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/sin/Summary",
							Description: "Libraries/base/Fields/math:Fields/sin/Description",
						},
						"sinh": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/sinh/Summary",
							Description: "Libraries/base/Fields/math:Fields/sinh/Description",
						},
						"sqrt": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/sqrt/Summary",
							Description: "Libraries/base/Fields/math:Fields/sqrt/Description",
						},
						"tan": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/tan/Summary",
							Description: "Libraries/base/Fields/math:Fields/tan/Description",
						},
						"tanh": dump.Function{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/math:Fields/tanh/Summary",
							Description: "Libraries/base/Fields/math:Fields/tanh/Description",
						},
					},
					Summary:     "Libraries/base/Fields/math:Summary",
					Description: "Libraries/base/Fields/math:Description",
				},
				"os": dump.Struct{
					Fields: dump.Fields{
						"clock": dump.Function{
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/os:Fields/clock/Summary",
							Description: "Libraries/base/Fields/os:Fields/clock/Description",
						},
						"date": dump.MultiFunction{
							{
								Returns: dump.Parameters{
									{Type: dt.Prim(rtypes.T_LuaString)},
								},
								Summary:     "Libraries/base/Fields/os:Fields/date/Current/Summary",
								Description: "Libraries/base/Fields/os:Fields/date/Current/Description",
							},
							{
								Parameters: dump.Parameters{
									{Name: "format", Type: dt.Prim(rtypes.T_LuaString), Enums: dt.Enums{`"*t"`, `!*t`}},
									{Name: "time", Type: dt.Optional(dt.Prim(rtypes.T_LuaNumber))},
								},
								Returns: dump.Parameters{
									{Type: dt.Struct(dt.KindStruct{
										"year":  dt.Prim(rtypes.T_LuaInteger),
										"month": dt.Prim(rtypes.T_LuaInteger),
										"day":   dt.Prim(rtypes.T_LuaInteger),
										"hour":  dt.Optional(dt.Prim(rtypes.T_LuaInteger)),
										"min":   dt.Optional(dt.Prim(rtypes.T_LuaInteger)),
										"sec":   dt.Optional(dt.Prim(rtypes.T_LuaInteger)),
										"wday":  dt.Optional(dt.Prim(rtypes.T_LuaInteger)),
										"yday":  dt.Optional(dt.Prim(rtypes.T_LuaInteger)),
										"isdst": dt.Optional(dt.Prim(rtypes.T_LuaBoolean)),
									})},
								},
								Summary:     "Libraries/base/Fields/os:Fields/date/Tabular/Summary",
								Description: "Libraries/base/Fields/os:Fields/date/Tabular/Description",
							},
							{
								Parameters: dump.Parameters{
									{Name: "format", Type: dt.Prim(rtypes.T_LuaString)},
									{Name: "time", Type: dt.Optional(dt.Prim(rtypes.T_LuaNumber))},
								},
								Returns: dump.Parameters{
									{Type: dt.Prim(rtypes.T_LuaString)},
								},
								Summary:     "Libraries/base/Fields/os:Fields/date/Formatted/Summary",
								Description: "Libraries/base/Fields/os:Fields/date/Formatted/Description",
							},
						},
						"difftime": dump.Function{
							Parameters: dump.Parameters{
								{Name: "t2", Type: dt.Prim(rtypes.T_LuaNumber)},
								{Name: "t1", Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/os:Fields/difftime/Summary",
							Description: "Libraries/base/Fields/os:Fields/difftime/Description",
						},
						"time": dump.Function{
							Parameters: dump.Parameters{
								{Name: "t", Type: dt.Optional(dt.Struct(dt.KindStruct{
									"year":  dt.Prim(rtypes.T_LuaInteger),
									"month": dt.Prim(rtypes.T_LuaInteger),
									"day":   dt.Prim(rtypes.T_LuaInteger),
									"hour":  dt.Optional(dt.Prim(rtypes.T_LuaInteger)),
									"min":   dt.Optional(dt.Prim(rtypes.T_LuaInteger)),
									"sec":   dt.Optional(dt.Prim(rtypes.T_LuaInteger)),
									"isdst": dt.Optional(dt.Prim(rtypes.T_LuaBoolean)),
								}))},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaNumber)},
							},
							Summary:     "Libraries/base/Fields/os:Fields/time/Summary",
							Description: "Libraries/base/Fields/os:Fields/time/Description",
						},
					},
					Summary:     "Libraries/base/Fields/os:Summary",
					Description: "Libraries/base/Fields/os:Description",
				},
				"string": dump.Struct{
					Fields: dump.Fields{
						"byte": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim(rtypes.T_LuaString)},
								{Name: "i", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger)), Default: `1`},
								{Name: "j", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger)), Default: `i`},
							},
							Returns: dump.Parameters{
								{Name: "...", Type: dt.Prim(rtypes.T_LuaInteger)},
							},
							Summary:     "Libraries/base/Fields/string:Fields/byte/Summary",
							Description: "Libraries/base/Fields/string:Fields/byte/Description",
						},
						"char": dump.Function{
							Parameters: dump.Parameters{
								{Name: "...", Type: dt.Prim(rtypes.T_LuaInteger)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaString)},
							},
							CanError:    true,
							Summary:     "Libraries/base/Fields/string:Fields/char/Summary",
							Description: "Libraries/base/Fields/string:Fields/char/Description",
						},
						"find": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim(rtypes.T_LuaString)},
								{Name: "pattern", Type: dt.Prim(rtypes.T_LuaString)},
								{Name: "init", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger)), Default: `1`},
								{Name: "plain", Type: dt.Optional(dt.Prim(rtypes.T_LuaBoolean)), Default: `false`},
							},
							Returns: dump.Parameters{
								{Name: "start", Type: dt.Optional(dt.Prim(rtypes.T_LuaNumber))},
								{Name: "end", Type: dt.Optional(dt.Prim(rtypes.T_LuaNumber))},
								{Name: "...", Type: dt.Optional(dt.Prim(rtypes.T_LuaString))},
							},
							CanError:    true,
							Summary:     "Libraries/base/Fields/string:Fields/find/Summary",
							Description: "Libraries/base/Fields/string:Fields/find/Description",
						},
						"format": dump.Function{
							Parameters: dump.Parameters{
								{Name: "format", Type: dt.Prim(rtypes.T_LuaString)},
								{Name: "...", Type: dt.Prim(rtypes.T_Any)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaString)},
							},
							CanError:    true,
							Summary:     "Libraries/base/Fields/string:Fields/format/Summary",
							Description: "Libraries/base/Fields/string:Fields/format/Description",
						},
						"gmatch": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim(rtypes.T_LuaString)},
								{Name: "pattern", Type: dt.Prim(rtypes.T_LuaString)},
							},
							Returns: dump.Parameters{
								{Type: dt.Function(dt.KindFunction{
									Returns: dump.Parameters{
										{Name: "...", Type: dt.Prim(rtypes.T_LuaString)},
									},
								})},
							},
							CanError:    true,
							Summary:     "Libraries/base/Fields/string:Fields/gmatch/Summary",
							Description: "Libraries/base/Fields/string:Fields/gmatch/Description",
						},
						"gsub": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim(rtypes.T_LuaString)},
								{Name: "pattern", Type: dt.Prim(rtypes.T_LuaString)},
								{Name: "repl", Type: dt.Or(
									dt.Prim(rtypes.T_LuaString),
									dt.Map(
										dt.Prim(rtypes.T_LuaString),
										dt.Or(dt.Prim(rtypes.T_LuaString), dt.Prim(rtypes.T_LuaNumber), dt.Prim("false")),
									),
									dt.Function(dt.KindFunction{
										Parameters: dump.Parameters{
											{Name: "...", Type: dt.Prim(rtypes.T_LuaString)},
										},
										Returns: dump.Parameters{
											{Type: dt.Or(dt.Prim(rtypes.T_LuaString), dt.Prim(rtypes.T_LuaNumber), dt.Prim("false"), dt.Prim(rtypes.T_LuaNil))},
										},
									}),
								)},
								{Name: "n", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger))},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaString)},
								{Type: dt.Prim(rtypes.T_LuaInteger)},
							},
							CanError:    true,
							Summary:     "Libraries/base/Fields/string:Fields/gsub/Summary",
							Description: "Libraries/base/Fields/string:Fields/gsub/Description",
						},
						"len": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim(rtypes.T_LuaString)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaInteger)},
							},
							Summary:     "Libraries/base/Fields/string:Fields/len/Summary",
							Description: "Libraries/base/Fields/string:Fields/len/Description",
						},
						"lower": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim(rtypes.T_LuaString)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaString)},
							},
							Summary:     "Libraries/base/Fields/string:Fields/lower/Summary",
							Description: "Libraries/base/Fields/string:Fields/lower/Description",
						},
						"match": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim(rtypes.T_LuaString)},
								{Name: "pattern", Type: dt.Prim(rtypes.T_LuaString)},
								{Name: "init", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger)), Default: `1`},
							},
							Returns: dump.Parameters{
								{Name: "...", Type: dt.Optional(dt.Prim(rtypes.T_LuaString))},
							},
							CanError:    true,
							Summary:     "Libraries/base/Fields/string:Fields/match/Summary",
							Description: "Libraries/base/Fields/string:Fields/match/Description",
						},
						"rep": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim(rtypes.T_LuaString)},
								{Name: "n", Type: dt.Prim(rtypes.T_LuaInteger)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaString)},
							},
							Summary:     "Libraries/base/Fields/string:Fields/rep/Summary",
							Description: "Libraries/base/Fields/string:Fields/rep/Description",
						},
						"reverse": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim(rtypes.T_LuaString)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaString)},
							},
							Summary:     "Libraries/base/Fields/string:Fields/reverse/Summary",
							Description: "Libraries/base/Fields/string:Fields/reverse/Description",
						},
						"sub": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim(rtypes.T_LuaString)},
								{Name: "i", Type: dt.Prim(rtypes.T_LuaInteger)},
								{Name: "j", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger)), Default: `-1`},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaString)},
							},
							Summary:     "Libraries/base/Fields/string:Fields/sub/Summary",
							Description: "Libraries/base/Fields/string:Fields/sub/Description",
						},
						"upper": dump.Function{
							Parameters: dump.Parameters{
								{Name: "s", Type: dt.Prim(rtypes.T_LuaString)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaString)},
							},
							Summary:     "Libraries/base/Fields/string:Fields/upper/Summary",
							Description: "Libraries/base/Fields/string:Fields/upper/Description",
						},
					},
					Summary:     "Libraries/base/Fields/string:Summary",
					Description: "Libraries/base/Fields/string:Description",
				},
				"table": dump.Struct{
					Fields: dump.Fields{
						"concat": dump.Function{
							Parameters: dump.Parameters{
								{Name: "t", Type: dt.Array(dt.Or(dt.Prim(rtypes.T_LuaString), dt.Prim(rtypes.T_LuaNumber)))},
								{Name: "sep", Type: dt.Optional(dt.Prim(rtypes.T_LuaString)), Default: `""`},
								{Name: "i", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger)), Default: `1`},
								{Name: "j", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger)), Default: `#t`},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaString)},
							},
							CanError:    true,
							Summary:     "Libraries/base/Fields/table:Fields/concat/Summary",
							Description: "Libraries/base/Fields/table:Fields/concat/Description",
						},
						"insert": dump.MultiFunction{
							{
								Parameters: dump.Parameters{
									{Name: "t", Type: dt.Prim(rtypes.T_LuaTable)},
									{Name: "index", Type: dt.Prim(rtypes.T_LuaInteger)},
									{Name: "value", Type: dt.Prim(rtypes.T_Any)},
								},
								Summary:     "Libraries/base/Fields/table:Fields/insert/Insert/Summary",
								Description: "Libraries/base/Fields/table:Fields/insert/Insert/Description",
							},
							{
								Parameters: dump.Parameters{
									{Name: "t", Type: dt.Prim(rtypes.T_LuaTable)},
									{Name: "value", Type: dt.Prim(rtypes.T_Any)},
								},
								Summary:     "Libraries/base/Fields/table:Fields/insert/Append/Summary",
								Description: "Libraries/base/Fields/table:Fields/insert/Append/Description",
							},
						},
						"maxn": dump.Function{
							Parameters: dump.Parameters{
								{Name: "t", Type: dt.Prim(rtypes.T_LuaTable)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_LuaInteger)},
							},
							Summary:     "Libraries/base/Fields/table:Fields/maxn/Summary",
							Description: "Libraries/base/Fields/table:Fields/maxn/Description",
						},
						"remove": dump.Function{
							Parameters: dump.Parameters{
								{Name: "t", Type: dt.Prim(rtypes.T_LuaTable)},
								{Name: "index", Type: dt.Optional(dt.Prim(rtypes.T_LuaInteger)), Default: `#t`},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_Any)},
							},
							Summary:     "Libraries/base/Fields/table:Fields/remove/Summary",
							Description: "Libraries/base/Fields/table:Fields/remove/Description",
						},
						"sort": dump.Function{
							Parameters: dump.Parameters{
								{Name: "t", Type: dt.Prim(rtypes.T_LuaTable)},
								{Name: "comp", Type: dt.Optional(dt.Function(dt.KindFunction{
									Parameters: dump.Parameters{
										{Name: "a", Type: dt.Prim(rtypes.T_Any)},
										{Name: "b", Type: dt.Prim(rtypes.T_Any)},
									},
									Returns: dump.Parameters{
										{Type: dt.Prim(rtypes.T_LuaBoolean)},
									},
								}))},
							},
							CanError:    true,
							Summary:     "Libraries/base/Fields/table:Fields/sort/Summary",
							Description: "Libraries/base/Fields/table:Fields/sort/Description",
						},
					},
					Summary:     "Libraries/base/Fields/table:Summary",
					Description: "Libraries/base/Fields/table:Description",
				},
			},
			Summary:     "Libraries/base:Summary",
			Description: "Libraries/base:Description",
		},
	}
}
