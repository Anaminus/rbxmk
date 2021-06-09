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
		Name: "DescFields",
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			fields, ok := v.(rtypes.DescFields)
			if !ok {
				return nil, rbxmk.TypeError{Want: "DescFields", Got: v.Type()}
			}
			table := s.L.CreateTable(0, len(fields))
			for k, v := range fields {
				if lv := pushDescField(s, v); lv != nil {
					table.RawSetString(k, lv)
				}
			}
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lvs[0].Type().String()}
			}
			fields := rbxdump.Fields{}
			err = table.ForEach(func(k, v lua.LValue) error {
				key, ok := k.(lua.LString)
				if !ok {
					return nil
				}
				name := string(key)
				fields[name], err = pullDescField(s, name, v)
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
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying:  dt.Dictionary{V: dt.Prim("any")},
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

func pushDescField(s rbxmk.State, v interface{}) lua.LValue {
	switch v := v.(type) {
	case bool:
		return lua.LBool(v)
	case int:
		return lua.LNumber(v)
	case string:
		return lua.LString(v)
	case rbxdump.Tags:
		a := s.L.CreateTable(len(v), 0)
		for _, v := range v {
			a.Append(lua.LString(v))
		}
		return a
	case []string:
		a := s.L.CreateTable(len(v), 0)
		for _, v := range v {
			a.Append(lua.LString(v))
		}
		return a
	case rbxdump.Type:
		lvs, _ := s.PushTo(rtypes.TypeDesc{Embedded: v})
		return lvs[0]
	case []rbxdump.Parameter:
		a := s.L.CreateTable(len(v), 0)
		for _, v := range v {
			lvs, _ := s.PushTo(rtypes.ParameterDesc{Parameter: v})
			a.Append(lvs[0])
		}
		return a
	}
	return nil
}

func pullDescField(s rbxmk.State, k string, v lua.LValue) (interface{}, error) {
	switch v := v.(type) {
	case lua.LBool:
		return bool(v), nil
	case lua.LNumber:
		return int(v), nil
	case lua.LString:
		return string(v), nil
	case *lua.LTable:
		if v.RawGetString("Category") != lua.LNil && v.RawGetString("Name") != lua.LNil {
			t, err := s.PullFrom("TypeDesc", v)
			if err != nil {
				return nil, fmt.Errorf("field %q: %w", k, err)
			}
			return t.(rtypes.TypeDesc).Embedded, nil
		}
		switch k {
		case "Parameters":
			a := make([]rbxdump.Parameter, v.Len())
			for i := 1; i <= len(a); i++ {
				a[i-1] = s.PullFromTable(v, lua.LNumber(i), "ParameterDesc").(rtypes.ParameterDesc).Parameter
			}
			return a, nil
		case "Tags":
			a := make(rbxdump.Tags, v.Len())
			for i := 1; i <= len(a); i++ {
				a[i-1] = string(s.PullFromTable(v, lua.LNumber(i), "string").(types.String))
			}
			return a, nil
		}
	}
	return nil, fmt.Errorf("field %s: unexpected type %s", k, v.Type())
}
