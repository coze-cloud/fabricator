package core

type ForRange struct {
	Declaration Declaration
	Block Block
}

func (f ForRange) String() string {
	return "for " + f.Declaration.listNames() + " := range " +
		f.Declaration.Value.String() + " " + f.Block.String()
}



