package core

type Receiver struct {
	Name string
	Type string
}

func (r Receiver) String() string {
	if len(r.Name) == 0 || len(r.Type) == 0 {
		return ""
	}
	return "(" + r.Name + " " + r.Type + ") "
}