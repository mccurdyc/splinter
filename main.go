package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "./tests/1/main.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
	}

	var v visitor
	ast.Walk(v, node)
}

type visitor struct{}

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch d := n.(type) {
	case *ast.ReturnStmt:
		for _, res := range d.Results {
			fmt.Printf("%+v\n", res)
		}
	}
	return v
}
