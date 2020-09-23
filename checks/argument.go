package checks

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"

	config "github.com/tommy-muehle/go-mnd/config"
)

const ArgumentCheck = "argument"

// Known excludes for the argument check.
var argumentExcludes = map[string]string{
	// package: function
	"time": "Date", // https://golang.org/pkg/time/#Date
}

// constantDefinitions is used to save lines (by number) which contain a constant definition.
var constantDefinitions = map[int]bool{}

type ArgumentAnalyzer struct {
	config *config.Config
	pass   *analysis.Pass
}

func NewArgumentAnalyzer(pass *analysis.Pass, config *config.Config) *ArgumentAnalyzer {
	return &ArgumentAnalyzer{
		pass:   pass,
		config: config,
	}
}

func (a *ArgumentAnalyzer) NodeFilter() []ast.Node {
	return []ast.Node{
		(*ast.GenDecl)(nil),
		(*ast.CallExpr)(nil),
	}
}

func (a *ArgumentAnalyzer) Check(n ast.Node) {
	switch expr := n.(type) {
	case *ast.CallExpr:
		a.checkCallExpr(expr)
	case *ast.GenDecl:
		if expr.Tok.String() == "const" {
			constantDefinitions[a.pass.Fset.Position(expr.TokPos).Line] = true
		}
	}
}

func (a *ArgumentAnalyzer) checkCallExpr(expr *ast.CallExpr) {
	if ok := constantDefinitions[a.pass.Fset.Position(expr.Pos()).Line]; ok {
		return
	}

	// Don't check if package and function combination is excluded
	if s, ok := expr.Fun.(*ast.SelectorExpr); ok && a.isExcluded(s) {
		return
	}

	for i, arg := range expr.Args {
		switch x := arg.(type) {
		case *ast.BasicLit:
			if !a.isMagicNumber(x) {
				continue
			}
			// If it's a magic number and has no previous element, report it
			if i == 0 {
				a.pass.Reportf(x.Pos(), reportMsg, x.Value, ArgumentCheck)
			} else {
				// Otherwise check the previous element type
				switch expr.Args[i-1].(type) {
				case *ast.ChanType:
					// When it's not a simple buffered channel, report it
					if a.isMagicNumber(x) {
						a.pass.Reportf(x.Pos(), reportMsg, x.Value, ArgumentCheck)
					}
				}
			}
		case *ast.BinaryExpr:
			a.checkBinaryExpr(x)
		}
	}
}

func (a *ArgumentAnalyzer) isExcluded(expr *ast.SelectorExpr) bool {
	var p string

	switch x := expr.X.(type) {
	case *ast.Ident:
		p = x.Name
	}

	if v, ok := argumentExcludes[p]; ok && v == expr.Sel.Name {
		return true
	}

	return false
}

func (a *ArgumentAnalyzer) checkBinaryExpr(expr *ast.BinaryExpr) {
	switch x := expr.X.(type) {
	case *ast.BasicLit:
		if a.isMagicNumber(x) {
			a.pass.Reportf(x.Pos(), reportMsg, x.Value, ArgumentCheck)
		}
	}

	switch y := expr.Y.(type) {
	case *ast.BasicLit:
		if a.isMagicNumber(y) {
			a.pass.Reportf(y.Pos(), reportMsg, y.Value, ArgumentCheck)
		}
	}
}

func (a *ArgumentAnalyzer) isMagicNumber(l *ast.BasicLit) bool {
	return (l.Kind == token.FLOAT || l.Kind == token.INT) && !a.config.IsIgnoredNumber(l.Value)
}
