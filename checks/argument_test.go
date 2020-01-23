package checks

import (
	"github.com/stretchr/testify/assert"

	"go/ast"
	"testing"
)

var excludeTests = []struct {
	name string
	in   *ast.SelectorExpr
	out  bool
}{
	{
		"time.Date",
		&ast.SelectorExpr{
			X: &ast.Ident{
				Name: "time",
			},
			Sel: &ast.Ident{
				Name: "Date",
			},
		},
		true,
	},
	{
		"http.StatusText",
		&ast.SelectorExpr{
			X: &ast.Ident{
				Name: "http",
			},
			Sel: &ast.Ident{
				Name: "StatusText",
			},
		},
		false,
	},
}

func Test_isExcluded(t *testing.T) {
	assert := assert.New(t)
	a := NewArgumentAnalyzer(nil, nil)

	for _, tt := range excludeTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(tt.out, a.isExcluded(tt.in))
		})
	}
}
