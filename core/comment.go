package core

type Comment struct {
	Value string
}

func (c Comment) String() string {
	return "// " + c.Value
}
