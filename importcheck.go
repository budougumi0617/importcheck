package importcheck

import (
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

// Analyzer provides static analysis for layered architecture.
var Analyzer = &analysis.Analyzer{
	Name: "importcheck",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

// Doc writes description this analyzer.
const Doc = "importcheck is ..."

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, i := range f.Imports {
			if isRepositoryPkg(i.Path.Value) {
				pass.Reportf(i.Pos(), "%s must not include %s", f.Name.Name, i.Path.Value)
			}
		}
	}

	return nil, nil
}

func isRepositoryPkg(s string) bool {
	return strings.HasSuffix(strings.Trim(s, "\""), "repository")
}
