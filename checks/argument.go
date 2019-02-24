package checks

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

const ArgumentCheck = "argument"

type ArgumentAnalyzer struct {
	pass *analysis.Pass
}

func NewArgumentAnalyzer(pass *analysis.Pass) *ArgumentAnalyzer {
	return &ArgumentAnalyzer{
		pass: pass,
	}
}

func (a *ArgumentAnalyzer) NodeFilter() []ast.Node {
	return []ast.Node{
		(*ast.CallExpr)(nil),
	}
}

func (a *ArgumentAnalyzer) Check(n ast.Node) {
	expr, ok := n.(*ast.CallExpr)
	if !ok {
		return
	}

	for _, arg := range expr.Args {
		switch x := arg.(type) {
		case *ast.BasicLit:
			if isMagicNumber(x) {
				a.pass.Reportf(x.Pos(), reportMsg, x.Value, ArgumentCheck)
			}
		case *ast.BinaryExpr:
			a.checkBinaryExpr(x)
		}
	}
}

func (a *ArgumentAnalyzer) checkBinaryExpr(expr *ast.BinaryExpr) {
	switch x := expr.X.(type) {
	case *ast.BasicLit:
		if isMagicNumber(x) {
			a.pass.Reportf(x.Pos(), reportMsg, x.Value, ArgumentCheck)
		}
	}

	switch y := expr.Y.(type) {
	case *ast.BasicLit:
		if isMagicNumber(y) {
			a.pass.Reportf(y.Pos(), reportMsg, y.Value, ArgumentCheck)
		}
	}
}
