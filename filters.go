package main

type Filter interface {
	Signature() (name string, args []interface{})
	Filter(in ...interface{}) (out []*Source)
}

// type FilterDrill struct{}

// func (FilterDrill) Signature() string {
// 	return "drill", []interface{}{&Source{}, ""}
// }

// func (FilterDrill) Filter(in ...interface{}) (out []*Source) {
// 	input := in[0].(*Source)
// 	ref := in[1].(string)
// }
