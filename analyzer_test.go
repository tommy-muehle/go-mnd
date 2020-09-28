package magic_numbers

import (
	"flag"
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

func TestCanExcludeNumbers(t *testing.T) {
	testdata := analysistest.TestData()

	options := flag.NewFlagSet("", flag.ExitOnError)
	options.String("checks", "assign,argument", "")
	options.String("excludes", "", "")
	options.String("ignored-numbers", "1000,1234_567_890,3.14159264", "")

	analyzer := Analyzer
	analyzer.Flags = *options

	analysistest.Run(t, testdata, analyzer, "ignored")
}
