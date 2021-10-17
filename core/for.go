package core

type For struct {
	Initialization Node
	Condition Node
	PostCondition Node
	Block Block
}

func (f For) String() string {
	return "for " + f.Initialization.String() + "; " +
		f.Condition.String() + "; " +
		f.PostCondition.String() + " " +
		f.Block.String()
}
