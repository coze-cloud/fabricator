package core

type Interface struct {
	Block Block
}

func (i Interface) String() string {
	return "interface " + i.Block.String()
}