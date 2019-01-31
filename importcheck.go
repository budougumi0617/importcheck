package importcheck

import (
	"strings"

	"golang.org/x/tools/go/analysis"
)

// Analyzer provides static analysis for layered architecture.
var Analyzer = &analysis.Analyzer{
	Name: "importcheck",
	Doc:  Doc,
	Run:  run,
}

// Doc writes description this analyzer.
const Doc = "importcheck confirms clean architecture."

func run(pass *analysis.Pass) (interface{}, error) {
	if pass.Pkg.Name() != "handler" {
		return nil, nil
	}

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
