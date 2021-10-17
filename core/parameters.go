package core

type Parameters []Parameter

func (p Parameters) String() string {
	return p.listParameters()
}

func (p Parameters) listParameters() string{
	if len(p) == 0 {
		return ""
	}

	parameterList := p[0].String()
	for _, parameter := range p[1:] {
		parameterList = parameterList + ", " + parameter.String()
	}
	return parameterList
}

