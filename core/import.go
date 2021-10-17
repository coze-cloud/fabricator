package core

type Import struct {
	Name string
	Path string
}



func (i Import) String() string {
	if len(i.Name) != 0 {
		return "import " + i.Name + " \"" + i.Path + "\""
	}
	return "import \"" + i.Path + "\""
}

