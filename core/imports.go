package core

import "strings"

type Imports []Import

func (i Imports) String() string {
	return "import (\n" + listNameAndPath(i...) + "\n)"
}

func listNameAndPath(imports ...Import) string {
	importList := ""
	for _, imp := range imports {
		if len(imp.Name) > 0 {
			importList = importList + "\n" + imp.Name + " \"" + imp.Path + "\""
			continue
		}
		importList = importList + "\n\"" + imp.Path + "\""
	}
	return strings.TrimPrefix(importList, "\n")
}

