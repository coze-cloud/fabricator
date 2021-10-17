package core

type Func struct {
	Receiver Receiver
	Name string
	Parameters Parameters
	ReturnTypes []string
	Block Block
}

func (f Func) String() string {
	if len(f.Name) == 0 {
		return "func(" + f.Parameters.listParameters()  + ")" + f.listReturnTypes() + " " + f.Block.String()
	}
	return "func " + f.Receiver.String() + f.Name + "(" + f.Parameters.listParameters()  + ")" +
		f.listReturnTypes() + " " + f.Block.String()
}

func (f Func) listReturnTypes() string {
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
