package core

type Parameter struct {
	Name string
	Type Node
}

func (p Parameter) String() string {
	if len(p.Name) == 0 && p.Type != nil {
		return p.Type.String()
	}
	if len(p.Name) != 0 && p.Type == nil {
		return p.Name
	}
	return p.Name + " " + p.Type.String()
}

