package core

type Package struct {
	Name string
	Nodes []Node
}

func (p Package) String() string {
	return "package " + p.Name + "\n" +
		p.listNodes()
}

func (p Package) listNodes() string {
	nodeList := p.Nodes[0].String()
	for _, node := range p.Nodes[1:] {
		nodeList = nodeList + "\n" + node.String()
	}
	return nodeList
}

