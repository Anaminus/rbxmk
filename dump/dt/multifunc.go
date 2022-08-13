package dt

// KindMultiFunctionType is a Type that indicates a function with multiple
// signatures.
type KindMultiFunctionType struct{}

const K_Functions = "functions"

func Functions() Type { return Type{Kind: KindMultiFunctionType{}} }

func (KindMultiFunctionType) k()           {}
func (KindMultiFunctionType) Kind() string { return K_Functions }
func (KindMultiFunctionType) String() string {
	return "(...) -> (...)"
}
