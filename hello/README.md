## Running Go Code

```sh
# build and run without outputting .exe
go run .
```

alternatively...

```sh
# build and outpout as helloworld.exe
go build -o helloworld
```

The binary can be shipped and executed on both linux/windows machines.

### Dependency management

Manage packags using `go.mod`

```sh
# this adds the package in your go.mod automatically too
go get rsc.io/quote@latest
```

This gets triggered whenever you run go CLI commands like `go build`, `go run`, `go mod tidy` (this one cleans up your go.mod)

### Importing external packages

Use the Go package registry to download external packages

```sh
# https://pkg.go.dev/rsc.io/quote/v4
import "rsc.io/quote"
```

### Go compilation and running code process

Go compiler compiles source code `.go` file into a binary executable (using Go `gc` compiler). It performs type checking, optimizations, and dependency linking.
