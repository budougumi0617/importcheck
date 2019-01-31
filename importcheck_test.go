package importcheck_test

import (
	"testing"

	"github.com/budougumi0617/importcheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	// GOPAH is ./testdata in test
	analysistest.Run(t, testdata, importcheck.Analyzer, "handler")
}
