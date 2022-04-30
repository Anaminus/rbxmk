package library

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/reflect"
)

func init() { register(Roblox) }

var Roblox = rbxmk.Library{
	Name:       "roblox",
	ImportedAs: "",
	Priority:   1,
	Open:       openRoblox,
	Dump:       dumpRoblox,
	Types: []func() rbxmk.Reflector{
		reflect.Array,
		reflect.Axes,
		reflect.BinaryString,
		reflect.Bool,
		reflect.BrickColor,
		reflect.CFrame,
		reflect.Color3,
		reflect.Color3uint8,
		reflect.ColorSequence,
		reflect.ColorSequenceKeypoint,
		reflect.Content,
		reflect.Dictionary,
		reflect.Double,
		reflect.Enum,
		reflect.EnumItem,
		reflect.Enums,
		reflect.Faces,
		reflect.Float,
		reflect.Instance,
		reflect.Int,
		reflect.Int64,
		reflect.Number,
		reflect.NumberRange,
		reflect.NumberSequence,
		reflect.NumberSequenceKeypoint,
		reflect.Objects,
		reflect.PhysicalProperties,
		reflect.ProtectedString,
		reflect.Ray,
		reflect.Rect,
		reflect.Region3,
		reflect.Region3int16,
		reflect.SharedString,
		reflect.String,
		reflect.Token,
		reflect.Tuple,
		reflect.UDim,
		reflect.UDim2,
		reflect.UniqueId,
		reflect.Variant,
		reflect.Vector2,
		reflect.Vector2int16,
		reflect.Vector3,
		reflect.Vector3int16,
	},
}

func openRoblox(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 1)
	lib.RawSetString("typeof", s.WrapFunc(robloxTypeof))
	return lib
}

func robloxTypeof(s rbxmk.State) int {
	v := s.CheckAny(1)
	t := s.World.Typeof(v)
	s.L.Push(lua.LString(t))
	return 1
}

func dumpRoblox(s rbxmk.State) dump.Library {
	lib := dump.Library{
		Name: "roblox",
		Struct: dump.Struct{
			Fields: dump.Fields{
				"typeof": dump.Function{
					Parameters: dump.Parameters{
						{Name: "value", Type: dt.Prim("any")},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("string")},
					},
					Summary:     "Libraries/roblox:Fields/typeof/Summary",
					Description: "Libraries/roblox:Fields/typeof/Description",
				},
			},
			Summary:     "Libraries/roblox:Summary",
			Description: "Libraries/roblox:Description",
		},
		Types: dump.TypeDefs{
			"DataModel": dump.TypeDef{
				Underlying: dt.Prim("Instance"),
				Symbols: dump.Properties{
					"Metadata": dump.Property{
						ValueType:   dt.Dictionary{V: dt.Prim("string")},
						Summary:     "Libraries/roblox/Types/DataModel:Symbols/Metadata/Summary",
						Description: "Libraries/roblox/Types/DataModel:Symbols/Metadata/Description",
					},
				},
				Methods: dump.Methods{
					"GetService": dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Instance")},
						},
						Summary:     "Libraries/roblox/Types/DataModel:Methods/GetService/Summary",
						Description: "Libraries/roblox/Types/DataModel:Methods/GetService/Description",
					},
				},
				Summary:     "Libraries/roblox/Types/DataModel:Summary",
				Description: "Libraries/roblox/Types/DataModel:Description",
			},
		},
	}
	return lib
}
