package lcheck_test

import (
	"testing"

	"github.com/budougumi0617/lcheck/lcheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, lcheck.Analyzer, "a")
}