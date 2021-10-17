package core

type Statement struct {
	Value string
}

func (a Statement) String() string {
	return a.Value
}

