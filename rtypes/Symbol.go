package rtypes

type Symbol struct {
	Name string
}

func (Symbol) Type() string {
	return "Symbol"
}

func (s Symbol) String() string {
	return "Symbol<" + s.Name + ">"
}
