package library

import (
	"math"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
)

func init() { register(Math, 10) }

var Math = rbxmk.Library{Name: "math", Open: openMath, Dump: dumpMath}

func openMath(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 4)
	lib.RawSetString("clamp", s.WrapFunc(mathClamp))
	lib.RawSetString("log", s.WrapFunc(mathLog))
	lib.RawSetString("round", s.WrapFunc(mathRound))
	lib.RawSetString("sign", s.WrapFunc(mathSign))
	return lib
}

func mathClamp(s rbxmk.State) int {
	x := s.CheckNumber(1)
	min := s.CheckNumber(2)
	max := s.CheckNumber(3)
	if min > max {
		s.L.RaiseError("max must be greater than min")
	}
	if x < min {
		x = min
	} else if x > max {
		x = max
	}
	s.L.Push(x)
	return 1
}

func mathLog(s rbxmk.State) int {
	x := s.CheckNumber(1)
	if s.L.Get(2) == lua.LNil {
		s.L.Push(lua.LNumber(math.Log(float64(x))))
		return 1
	}
	var res float64
	switch base := s.CheckNumber(2); base {
	case 2:
		res = math.Log2(float64(x))
	case 10:
		res = math.Log10(float64(x))
	default:
		res = math.Log(float64(x)) / math.Log(float64(base))
	}
	s.L.Push(lua.LNumber(res))
	return 1
}

func mathRound(s rbxmk.State) int {
	// Half away from zero.
	s.L.Push(lua.LNumber(math.Round(float64(s.CheckNumber(1)))))
	return 1
}

func mathSign(s rbxmk.State) int {
	x := s.CheckNumber(1)
	if x > 0 {
		s.L.Push(lua.LNumber(1))
	} else if x < 0 {
		s.L.Push(lua.LNumber(-1))
	} else {
		s.L.Push(lua.LNumber(0))
	}
	return 1
}

func dumpMath(s rbxmk.State) dump.Library {
	return dump.Library{
		Struct: dump.Struct{
			Fields: dump.Fields{
				"clamp": dump.Function{
					Parameters: dump.Parameters{
						{Name: "x", Type: dt.Prim("number")},
						{Name: "min", Type: dt.Prim("number")},
						{Name: "max", Type: dt.Prim("number")},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("number")},
					},
					CanError:    true,
					Summary:     "libraries/math:Fields/clamp/Summary",
					Description: "libraries/math:Fields/clamp/Description",
				},
				"log": dump.Function{
					Parameters: dump.Parameters{
						{Name: "x", Type: dt.Prim("number")},
						{Name: "base", Type: dt.Optional{T: dt.Prim("number")}, Default: "𝑒"},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("number")},
					},
					Summary:     "libraries/math:Fields/log/Summary",
					Description: "libraries/math:Fields/log/Description",
				},
				"round": dump.Function{
					Parameters: dump.Parameters{
						{Name: "x", Type: dt.Prim("number")},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("number")},
					},
					Summary:     "libraries/math:Fields/round/Summary",
					Description: "libraries/math:Fields/round/Description",
				},
				"sign": dump.Function{
					Parameters: dump.Parameters{
						{Name: "x", Type: dt.Prim("number")},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("number")},
					},
					Summary:     "libraries/math:Fields/sign/Summary",
					Description: "libraries/math:Fields/sign/Description",
				},
			},
			Summary:     "libraries/math:Summary",
			Description: "libraries/math:Description",
		},
	}
}
