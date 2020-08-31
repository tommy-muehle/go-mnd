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
		var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {}()
	}

	wg.Wait()
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
