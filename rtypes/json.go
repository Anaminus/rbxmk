package rtypes

import (
	"encoding/json"

	"github.com/robloxapi/types"
)

func DecodeJSON(u interface{}) (v types.Value) {
	switch u := u.(type) {
	case nil:
		return Nil
	case bool:
		return types.Bool(u)
	case float64:
		return types.Double(u)
	case string:
		return types.String(u)
	case []interface{}:
		a := make(Array, len(u))
		for i, u := range u {
			a[i] = DecodeJSON(u)
		}
		return a
	case map[string]interface{}:
		a := make(Dictionary, len(u))
		for k, u := range u {
			a[k] = DecodeJSON(u)
		}
		return a
	}
	return Nil
}

func EncodeJSON(v types.Value) (u interface{}) {
	//WARN: Must not receive cyclic tables. The Array and Dictionary type
	//reflectors already validate this, but such values could still be produced
	//internally.
	switch v := v.(type) {
	case NilType:
		return nil
	case types.Bool:
		return bool(v)
	case types.Double:
		return float64(v)
	case types.String:
		return string(v)
	case Array:
		a := make([]interface{}, len(v))
		for i, v := range v {
			a[i] = EncodeJSON(v)
		}
		return a
	case Dictionary:
		a := make(map[string]interface{}, len(v))
		for k, v := range v {
			a[k] = EncodeJSON(v)
		}
		return a
	}
	return nil
}

const T_JsonValue = "JsonValue"

type JsonValue struct {
	types.Value
}

func (JsonValue) Type() string {
	return T_JsonValue
}

func (v *JsonValue) UnmarshalJSON(b []byte) error {
	var u any
	if err := json.Unmarshal(b, &u); err != nil {
		return err
	}
	v.Value = DecodeJSON(u)
	return nil
}

func (v JsonValue) MarshalJSON() (b []byte, err error) {
	u := EncodeJSON(v.Value)
	return json.Marshal(u)
}

// JsonOp indicates a type of JsonOperation.
type JsonOp string

const (
	JsonOpTest    JsonOp = "test"
	JsonOpRemove  JsonOp = "remove"
	JsonOpAdd     JsonOp = "add"
	JsonOpReplace JsonOp = "replace"
	JsonOpMove    JsonOp = "move"
	JsonOpCopy    JsonOp = "copy"
)

func (o JsonOp) Valid() bool {
	switch o {
	case
		JsonOpTest,
		JsonOpRemove,
		JsonOpAdd,
		JsonOpReplace,
		JsonOpMove,
		JsonOpCopy:
		return true
	default:
		return false
	}
}

const T_JsonOperation = "JsonOperation"

// JsonOperation describes a single operation that transforms a JSON structure.
type JsonOperation struct {
	Op    JsonOp    `json:"op"`
	Path  string    `json:"path"`
	From  string    `json:"from,omitempty"`  // Only move and copy
	Value JsonValue `json:"value,omitempty"` // Only test, add, replace
}

// Type returns a string identifying the type of the value.
func (JsonOperation) Type() string {
	return T_JsonOperation
}

const T_JsonPatch = "JsonPatch"

// JsonPatch is a list of operations that transforms a JSON structure.
type JsonPatch []JsonOperation

// Type returns a string identifying the type of the value.
func (JsonPatch) Type() string {
	return T_JsonPatch
}
