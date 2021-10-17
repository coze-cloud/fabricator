package core

type Assignment struct {
	Names []string
	Value Node
}

func (d Assignment) String() string {
	return d.listNames() + " = " + d.Value.String()
}

func (d Assignment) listNames() string {
	nameList := d.Names[0]
	for _, name := range d.Names[1:] {
		nameList = nameList + ", " + name
	}
	return nameList
}
