package core

type Struct struct {
	Block Block
}

func (s Struct) String() string {
	return "struct " + s.Block.String()
}

