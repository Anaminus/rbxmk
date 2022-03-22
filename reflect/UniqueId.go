package reflect

import (
	"sync"
	"time"

	"math/rand"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

var uid struct {
	mut  sync.Mutex
	base rtypes.UniqueId
}

func updateUID() {
	const epoch = 1609459200 // 2021-01-01 00:00:00
	uid.base.Time = uint32(time.Now().Unix() - epoch)
	uid.base.Random = types.Int64(rand.Int63())
}

func nextUID() rtypes.UniqueId {
	uid.mut.Lock()
	defer uid.mut.Unlock()
	uid.base.Index++
	if uid.base.Index == 0 {
		updateUID()
	}
	return uid.base
}

func init() {
	register(UniqueId)

	uid.mut.Lock()
	defer uid.mut.Unlock()
	uid.base.Index = 0x10000
	updateUID()
}

func UniqueId() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "UniqueId",
		PushTo:   rbxmk.PushTypeTo("UniqueId"),
		PullFrom: rbxmk.PullTypeFrom("UniqueId"),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.UniqueId:
				*p = v.(rtypes.UniqueId)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "UniqueId").(rtypes.UniqueId)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "UniqueId").(rtypes.UniqueId)
				op := s.Pull(2, "UniqueId").(rtypes.UniqueId)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Random": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(rtypes.UniqueId).Random)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("int64"),
						ReadOnly:    true,
						Summary:     "Types/UniqueId:Properties/Random/Summary",
						Description: "Types/UniqueId:Properties/Random/Description",
					}
				},
			},
			"Time": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int64(v.(rtypes.UniqueId).Time))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("int64"),
						ReadOnly:    true,
						Summary:     "Types/UniqueId:Properties/Time/Summary",
						Description: "Types/UniqueId:Properties/Time/Description",
					}
				},
			},
			"Index": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int64(v.(rtypes.UniqueId).Index))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("int64"),
						ReadOnly:    true,
						Summary:     "Types/UniqueId:Properties/Index/Summary",
						Description: "Types/UniqueId:Properties/Index/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v rtypes.UniqueId
					switch s.Count() {
					case 0:
						v = nextUID()
					case 3:
						v.Random = s.Pull(1, "int64").(types.Int64)
						v.Time = uint32(s.Pull(2, "int64").(types.Int64))
						v.Index = uint32(s.Pull(2, "int64").(types.Int64))
					default:
						return s.RaiseError("expected 0 or 3 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{},
							Returns: dump.Parameters{
								{Type: dt.Prim("UniqueId")},
							},
							Summary:     "Types/UniqueId:Constructors/new/Generated/Summary",
							Description: "Types/UniqueId:Constructors/new/Generated/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "random", Type: dt.Prim("int64")},
								{Name: "time", Type: dt.Prim("int64")},
								{Name: "index", Type: dt.Prim("int64")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("UniqueId")},
							},
							Summary:     "Types/UniqueId:Constructors/new/Components/Summary",
							Description: "Types/UniqueId:Constructors/new/Components/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/UniqueId:Operators/Eq/Summary",
						Description: "Types/UniqueId:Operators/Eq/Description",
					},
				},
				Summary:     "Types/UniqueId:Summary",
				Description: "Types/UniqueId:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Int64,
		},
	}
}
