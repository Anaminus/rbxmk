package reflect

import (
	"math/rand"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(BrickColor) }
func BrickColor() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "BrickColor",
		PushTo:   rbxmk.PushTypeTo("BrickColor"),
		PullFrom: rbxmk.PullTypeFrom("BrickColor"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "BrickColor").(types.BrickColor)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "BrickColor").(types.BrickColor)
				op := s.Pull(2, "BrickColor").(types.BrickColor)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.String(v.(types.BrickColor).Name()))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("string"), ReadOnly: true} },
			},
			"Number": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int(v.(types.BrickColor).Number()))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("int"), ReadOnly: true} },
			},
			"R": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.BrickColor).R()))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"G": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.BrickColor).G()))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"B": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.BrickColor).B()))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"Color": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.BrickColor).Color())
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("Color3"), ReadOnly: true} },
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.BrickColor
					switch s.Count() {
					case 1:
						switch arg := s.PullAnyOf(1, "int", "string", "Color3").(type) {
						case types.Int:
							v = types.NewBrickColor(int(arg))
						case types.String:
							v = types.NewBrickColorFromName(string(arg))
						case types.Color3:
							v = types.NewBrickColorFromColor3(arg)
						default:
							return s.ReflectorError(1)
						}
					case 3:
						v = types.NewBrickColorFromColor(
							float64(s.Pull(1, "float").(types.Float)),
							float64(s.Pull(2, "float").(types.Float)),
							float64(s.Pull(3, "float").(types.Float)),
						)
					default:
						return s.RaiseError("expected 1 or 3 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{
						{
							Parameters: dump.Parameters{
								{Name: "value", Type: dt.Prim("int")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("BrickColor")},
							},
						},
						{
							Parameters: dump.Parameters{
								{Name: "r", Type: dt.Prim("float")},
								{Name: "g", Type: dt.Prim("float")},
								{Name: "b", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("BrickColor")},
							},
						},
						{
							Parameters: dump.Parameters{
								{Name: "name", Type: dt.Prim("string")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("BrickColor")},
							},
						},
						{
							Parameters: dump.Parameters{
								{Name: "color", Type: dt.Prim("Color3")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("BrickColor")},
							},
						},
					}
				},
			},
			"palette": {
				Func: func(s rbxmk.State) int {
					index := int(s.Pull(1, "int").(types.Int))
					return s.Push(types.NewBrickColorFromPalette(index))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Parameters: dump.Parameters{
							{Name: "index", Type: dt.Prim("int")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("BrickColor")},
						},
					}}
				},
			},
			"random": {
				Func: func(s rbxmk.State) int {
					index := rand.Intn(types.BrickColorIndexSize)
					return s.Push(types.NewBrickColorFromIndex(index))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Returns: dump.Parameters{
							{Type: dt.Prim("BrickColor")},
						},
					}}
				},
			},
			"White": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("White"))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Returns: dump.Parameters{
							{Type: dt.Prim("BrickColor")},
						},
					}}
				},
			},
			"Gray": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("Medium stone grey"))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Returns: dump.Parameters{
							{Type: dt.Prim("BrickColor")},
						},
					}}
				},
			},
			"DarkGray": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("Dark stone grey"))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Returns: dump.Parameters{
							{Type: dt.Prim("BrickColor")},
						},
					}}
				},
			},
			"Black": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("Black"))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Returns: dump.Parameters{
							{Type: dt.Prim("BrickColor")},
						},
					}}
				},
			},
			"Red": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("Bright red"))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Returns: dump.Parameters{
							{Type: dt.Prim("BrickColor")},
						},
					}}
				},
			},
			"Yellow": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("Bright yellow"))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Returns: dump.Parameters{
							{Type: dt.Prim("BrickColor")},
						},
					}}
				},
			},
			"Green": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("Dark green"))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Returns: dump.Parameters{
							{Type: dt.Prim("BrickColor")},
						},
					}}
				},
			},
			"Blue": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewBrickColorFromName("Bright blue"))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Returns: dump.Parameters{
							{Type: dt.Prim("BrickColor")},
						},
					}}
				},
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
