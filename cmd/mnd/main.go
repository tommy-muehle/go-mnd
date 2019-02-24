package main

import (
	"github.com/tommy-muehle/go-mnd"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(magic_numbers.Analyzer) }
