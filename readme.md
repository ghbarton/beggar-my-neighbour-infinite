To test the solution run `go test solution` in the cli

# Common error:

```# solution
FAIL	solution [setup failed]
solution_test.go:4:2: missing go.sum entry for module providing package github.com/stretchr/testify/assert (imported by solution); to add:
FAIL
go get -t solution

Compilation finished with exit code 1
```

## Fix: 
`go mod tidy`

To test the solution run `go test solution` in the cli
