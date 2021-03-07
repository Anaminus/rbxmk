package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(Rect) }
func Rect() Reflector {
	return Reflector{
		Name:     "Rect",
		PushTo:   rbxmk.PushTypeTo("Rect"),
		PullFrom: rbxmk.PullTypeFrom("Rect"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Rect").(types.Rect)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Rect").(types.Rect)
				op := s.Pull(2, "Rect").(types.Rect)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: map[string]Member{
			"Min": {
				Get: func(s State, v types.Value) int {
					return s.Push(v.(types.Rect).Min)
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("Vector2")} },
			},
			"Max": {
				Get: func(s State, v types.Value) int {
					return s.Push(v.(types.Rect).Max)
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("Vector2")} },
			},
			"Width": {
				Get: func(s State, v types.Value) int {
					return s.Push(types.Float(v.(types.Rect).Width()))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float")} },
			},
			"Height": {
				Get: func(s State, v types.Value) int {
					return s.Push(types.Float(v.(types.Rect).Height()))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float")} },
			},
		},
		Constructors: Constructors{
			"new": {
				Func: func(s State) int {
					var v types.Rect
					switch s.Count() {
					case 2:
						v.Min = s.Pull(1, "Vector2").(types.Vector2)
						v.Max = s.Pull(2, "Vector2").(types.Vector2)
					case 4:
						v.Min.X = float32(s.Pull(1, "float").(types.Float))
						v.Min.Y = float32(s.Pull(2, "float").(types.Float))
						v.Max.Y = float32(s.Pull(3, "float").(types.Float))
						v.Max.Y = float32(s.Pull(4, "float").(types.Float))
					default:
						return s.RaiseError("expected 2 or 4 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{
						{
							Parameters: dump.Parameters{
								{Name: "min", Type: dt.Prim("Vector2")},
								{Name: "max", Type: dt.Prim("Vector2")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("Rect")},
							},
						},
						{
							Parameters: dump.Parameters{
								{Name: "minX", Type: dt.Prim("float")},
								{Name: "minY", Type: dt.Prim("float")},
								{Name: "maxX", Type: dt.Prim("float")},
								{Name: "maxY", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("Rect")},
							},
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
