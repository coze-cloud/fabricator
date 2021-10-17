package core

type Block []Node

func (b Block) String() string {
	return "{" + b.listNodes()  + "}"
}

func (b Block) listNodes() string {
	if len(b) == 0 {
		return ""
	}

	nodeList := b[0].String()
	for _, node := range b[1:] {
		nodeList = nodeList + "\n" + node.String()
	}
	return "\n" + nodeList + "\n"
}

