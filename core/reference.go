package core

type Reference struct {
	Value Node
}

func (r Reference) String() string {
	return "&" + r.Value.String()
}
