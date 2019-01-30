# importcheck
Analyzer for validate package relationships.

# Demo
```bash
$ GOPATH=`pwd`/testdata:$GOPATH go run ./cmd/main.go ./testdata/src/...
${GOPATH}/github.com/budougumi0617/importcheck/testdata/src/handler/h.go:7:2: handler must not include "repository"
exit status 3
```
