package core

type FuncCall struct {
	Reference Node
	Parameters Parameters
}

func (f FuncCall) String() string {
	return f.Reference.String() + "(" + f.Parameters.String() + ")"
}

