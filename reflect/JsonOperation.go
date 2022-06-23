package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(JsonOperation) }
func JsonOperation() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_JsonOperation,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			op, ok := v.(rtypes.JsonOperation)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_JsonOperation, Got: v.Type()}
			}
			table := c.CreateTable(0, 4)
			if !op.Op.Valid() {
				return nil, fmt.Errorf("invalid operation %q", op.Op)
			}
			if err := c.PushToDictionary(table, "op", types.String(op.Op)); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "path", types.String(op.Path)); err != nil {
				return nil, err
			}
			switch op.Op {
			case rtypes.JsonOpTest, rtypes.JsonOpAdd, rtypes.JsonOpReplace:
				if op.Value.Value != nil {
					if err := c.PushToDictionary(table, "value", op.Value.Value); err != nil {
						return nil, err
					}
				}
			case rtypes.JsonOpMove, rtypes.JsonOpCopy:
				if op.From != "" {
					if err := c.PushToDictionary(table, "from", types.String(op.From)); err != nil {
						return nil, err
					}
				}
			}
			return table, nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_Table, Got: lv.Type().String()}
			}
			var jop rtypes.JsonOperation
			op, err := c.PullFromDictionary(table, "op", rtypes.T_String)
			if err != nil {
				return nil, err
			}
			path, err := c.PullFromDictionary(table, "path", rtypes.T_String)
			if err != nil {
				return nil, err
			}
			jop.Op = rtypes.JsonOp(op.(types.String))
			if !jop.Op.Valid() {
				return nil, fmt.Errorf("invalid operation %q", jop.Op)
			}
			jop.Path = string(path.(types.String))
			switch jop.Op {
			case rtypes.JsonOpTest, rtypes.JsonOpAdd, rtypes.JsonOpReplace:
				value, err := c.PullFromDictionaryOpt(table, "value", nil, rtypes.T_JsonValue)
				if err != nil {
					return nil, err
				}
				if value != nil {
					jop.Value = rtypes.JsonValue{Value: value}
				}
			case rtypes.JsonOpMove, rtypes.JsonOpCopy:
				from, err := c.PullFromDictionaryOpt(table, "value", nil, rtypes.T_String)
				if err != nil {
					return nil, err
				}
				if from, ok := from.(types.String); ok {
					jop.From = string(from)
				}
			}
			return jop, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.JsonOperation:
				*p = v.(rtypes.JsonOperation)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"op":    dt.Prim(rtypes.T_String),
					"path":  dt.Prim(rtypes.T_String),
					"from":  dt.Optional{T: dt.Prim(rtypes.T_String)},
					"value": dt.Optional{T: dt.Prim(rtypes.T_JsonValue)},
				},
				Summary:     "Types/JsonOperation:Summary",
				Description: "Types/JsonOperation:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			JsonValue,
			String,
		},
	}
}
