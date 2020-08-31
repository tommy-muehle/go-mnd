package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	mnd "github.com/tommy-muehle/go-mnd"
)

func main() { singlechecker.Main(mnd.Analyzer) }
