package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

var (
	src = `
	package main
	
	func main() {
		var y int
		_ = 5 * y * 6
	}
	`
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	if err := ast.Print(fset, f); err != nil {
		panic(err)
	}
}
