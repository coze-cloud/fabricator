package core

type Pointer struct {
	Value Node
}

func (p Pointer) String() string {
	return "*" + p.Value.String()
}
