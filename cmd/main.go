package main

import (
	"github.com/budougumi0617/importcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(importcheck.Analyzer)
}
