package core

type If struct {
	Condition Condition
	Block Block
}

func (i If) String() string {
	return "if " + i.Condition.String() + " " + i.Block.String()
}