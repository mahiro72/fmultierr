package fmultierr_test

import (
	"testing"

	"github.com/mahiro72/fmultierr"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)

	t.Run("multierrのErrors関数が使われていた場合, 検出する",func(t *testing.T) {
		analysistest.Run(t, testdata, fmultierr.Analyzer, "a")
	})

	t.Run("multierrが別名でimportされ、Errors関数が使われていた場合, 検出する",func(t *testing.T) {
		analysistest.Run(t, testdata, fmultierr.Analyzer, "b")
	})

	t.Run("今回の検出対象と別のmultierrパッケージがErrors関数をもっていた場合、検出しない",func(t *testing.T) {
		analysistest.Run(t, testdata, fmultierr.Analyzer, "c")
	})
}
