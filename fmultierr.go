package fmultierr

import (
	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
)

const doc = "fmultierr is found 'multierr.Errors'"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "fmultierr",
	Doc:  doc,
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	checkFuncObj := analysisutil.LookupFromImports(
		pass.Pkg.Imports(),
		"go.uber.org/multierr",
		"Errors",
	)

	for ident := range pass.TypesInfo.Uses {
		if checkFuncObj == pass.TypesInfo.ObjectOf(ident) {
			pass.Reportf(ident.Pos(),"multierr.Errors found")
		}
	}
	return nil, nil
}
