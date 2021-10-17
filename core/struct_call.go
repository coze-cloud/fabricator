package core

import "strings"

type StructCall struct {
	Name string
	Parameters Parameters
}

func (c StructCall) String() string {
	if len(c.Parameters) >= 3 {
		parameters := strings.ReplaceAll(c.listParameters(), ", ", ",\n") + ","
		return c.Name + "{\n" + parameters + "\n}"
	}
	return c.Name + "{" +  c.listParameters() + "}"
}

func (c StructCall) listParameters() string{
	if len(c.Parameters) == 0 {
		return ""
	}

	parameterList := c.Parameters[0].Name + ": " + c.Parameters[0].Type.String()
	for _, parameter := range c.Parameters[1:] {
		parameterList = parameterList + ", " + parameter.Name + ": " + parameter.Type.String()
	}
	return parameterList
}
