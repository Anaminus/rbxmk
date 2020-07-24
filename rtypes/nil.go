package rtypes

type NilType struct{}

var Nil NilType

func (NilType) Type() string {
	return "nil"
}

func (NilType) String() string {
	return "nil"
}
