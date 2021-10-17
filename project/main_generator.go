package project

import "github.com/cozy-hosting/fabricator/core"

type MainGenerator struct {

}

func NewMainGenerator() *MainGenerator {
	return &MainGenerator{}
}

func (m MainGenerator) String() string {
	return core.Package{
		Name: "main",
		Nodes: []core.Node{
			core.Statement{},
			core.Import{
				Path: "fmt",
			},
			core.Statement{},
			core.Type{
				Name: "Test",
				Type: core.Struct{},
			},
			core.Statement{},
			core.Func{
				Name: "main",
				Block: core.Block{
					core.InlineDeclaration{
						Names: []string{"test"},
						Value: core.Reference{
							Value: core.StructCall{
								Name: "Test",
							},
						},
					},
					core.FuncCall{
						Reference: core.Statement{
							Value: "fmt.Println",
						},
						Parameters: core.Parameters{
							core.Parameter{
								Name: "test",
							},
						},
					},
				},
			},
		},
	}.String()
}
