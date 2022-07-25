package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
)

func init() { register(DescFields) }
func DescFields() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_DescFields,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			fields, ok := v.(rtypes.DescFields)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_DescFields, Got: v.Type()}
			}
			table := c.CreateTable(0, len(fields))
			for k, v := range fields {
				if lv := pushDescField(c, v); lv != nil {
					table.RawSetString(k, lv)
				}
			}
			return table, nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_Table, Got: lv.Type().String()}
			}
			fields := rbxdump.Fields{}
			err = table.ForEach(func(k, v lua.LValue) error {
				key, ok := k.(lua.LString)
				if !ok {
					return nil
				}
				name := string(key)
				fields[name], err = pullDescField(c, name, v)
				if err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				return nil, err
			}
			return rtypes.DescFields(fields), nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.DescFields:
				*p = v.(rtypes.DescFields)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category:    "rbxmk",
				Underlying:  dt.Dictionary{V: dt.Prim(rtypes.T_Any)},
				Summary:     "Types/DescFields:Summary",
				Description: "Types/DescFields:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			ParameterDesc,
			String,
			TypeDesc,
		},
	}
}

func pushDescField(c rbxmk.Context, v interface{}) lua.LValue {
	switch v := v.(type) {
	case bool:
		return lua.LBool(v)
	case int:
		return lua.LNumber(v)
	case string:
		return lua.LString(v)
	case rbxdump.Tags:
		a := c.CreateTable(len(v), 0)
		for _, v := range v {
			a.Append(lua.LString(v))
		}
		return a
	case []string:
		a := c.CreateTable(len(v), 0)
		for _, v := range v {
			a.Append(lua.LString(v))
		}
		return a
	case rbxdump.Type:
		lv, _ := c.World.Push(rtypes.TypeDesc{Embedded: v})
		return lv
	case []rbxdump.Parameter:
		a := c.CreateTable(len(v), 0)
		for _, v := range v {
			lv, _ := c.World.Push(rtypes.ParameterDesc{Parameter: v})
			a.Append(lv)
		}
		return a
	}
	return nil
}

func pullDescField(c rbxmk.Context, k string, v lua.LValue) (f interface{}, err error) {
	switch v := v.(type) {
	case lua.LBool:
		return bool(v), nil
	case lua.LNumber:
		return int(v), nil
	case lua.LString:
		return string(v), nil
	case *lua.LTable:
		if v.RawGetString("Category") != lua.LNil && v.RawGetString("Name") != lua.LNil {
			t, err := c.World.Pull(v, rtypes.T_TypeDesc)
			if err != nil {
				return nil, fmt.Errorf("field %q: %w", k, err)
			}
			return t.(rtypes.TypeDesc).Embedded, nil
		}
		switch k {
		case "Parameters":
			a := make([]rbxdump.Parameter, v.Len())
			for i := 1; i <= len(a); i++ {
				p, err := c.PullFromArray(v, i, rtypes.T_ParameterDesc)
				if err != nil {
					return nil, fmt.Errorf("field %s[%d]: %w", k, i, err)
				}
				a[i-1] = p.(rtypes.ParameterDesc).Parameter
			}
			return a, nil
		case "Tags":
			a := make(rbxdump.Tags, v.Len())
			for i := 1; i <= len(a); i++ {
				tag, err := c.PullFromArray(v, i, rtypes.T_String)
				if err != nil {
					return nil, fmt.Errorf("field %s[%d]: %w", k, i, err)
				}
				a[i-1] = string(tag.(types.String))
			}
			return a, nil
		}
	}
	return nil, fmt.Errorf("field %s: unexpected type %s", k, v.Type())
}
