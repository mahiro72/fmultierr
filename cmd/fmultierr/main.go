package main

import (
	"github.com/mahiro72/fmultierr"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(fmultierr.Analyzer) }
