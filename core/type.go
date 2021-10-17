package core

type Type struct {
	Name string
	Type Node
}

func (t Type) String() string {
	return "type " + t.Name + " " + t.Type.String()
}

