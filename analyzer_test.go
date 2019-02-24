package magic_numbers

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestCaseChecks(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "case")
}

func TestArgumentChecks(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "argument")
}

func TestAssignChecks(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "assign")
}

func TestConditionChecks(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "condition")
}

func TestOperationChecks(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "operation")
}

func TestReturnChecks(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "return")
}
