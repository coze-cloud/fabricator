package core

type Declaration struct {
	Names []string
	Type string
	Value Node
}

func (d Declaration) String() string {
	if len(d.Type) > 0 {
		return "var " + d.listNames() + " " + d.Type
	}
	return "var " + d.listNames() + " = " + d.Value.String()
}

func (d Declaration) listNames() string {
	nameList := d.Names[0]
	for _, name := range d.Names[1:] {
		nameList = nameList + ", " + name
	}
	return nameList
}