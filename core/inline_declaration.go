package core

type InlineDeclaration struct {
	Names []string
	Value Node
}

func (d InlineDeclaration) String() string {
	return d.listNames() + " := " + d.Value.String()
}

func (d InlineDeclaration) listNames() string {
	nameList := d.Names[0]
	for _, name := range d.Names[1:] {
		nameList = nameList + ", " + name
	}
	return nameList
}