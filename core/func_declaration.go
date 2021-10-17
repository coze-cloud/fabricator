package core

type FuncDeclaration struct {
	Name string
	Parameters Parameters
	ReturnTypes []string
}

func (f FuncDeclaration) String() string {
	return f.Name + "(" + f.Parameters.String() + ")" + f.listReturnTypes()
}

func (f FuncDeclaration) listReturnTypes() string {
	if len(f.ReturnTypes) == 0 {
		return ""
	}
	if len(f.ReturnTypes) == 1 {
		return " " + f.ReturnTypes[0]
	}

	returnTypeList := f.ReturnTypes[0]
	for _, returnType := range f.ReturnTypes[1:] {
		returnTypeList = returnTypeList + ", " + returnType
	}
	return " (" + returnTypeList + ")"
}

