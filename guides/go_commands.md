[Go's official Command Documentation](https://go.dev/doc/cmd)

# Important Go Commands

| command              | description                                                                                   |
| -------------------- | --------------------------------------------------------------------------------------------- |
| `go help`            | Provides documentation and help for Go commands                                               |
| `go version`         | Shows currently installed Go version                                                          |
| `go env `            | Displays the current Go environment variables and their values                                |
| `go fmt file.go`     | Formats your Go source code according to the standard Go formatting guidelines                |
| `go build file.go`   | Compiles the source code of your package and produces an executable binary                    |
| `go run file.go`     | Compiles and runs your Go program in a sigle step                                             |
| `go vet file.go`     | Analyses your Go code for potential errors and suspicious constructs                          |
| `go test`            | Executes tests from files ending with *_test.go*                                              |
| `go clean`           | Removes build artefacts, such as executables and cached objects                               |
| `go get package`     | Fetches and installs third-party packages and its dependencies from remote                    |
| `go install package` | Builds and installs local packages or the current one (if none is specified)                  |
| `go fix`             | Applies updates and fixes to your Go source code to make it compatible with newer Go versions |
| `go list`            | Lists Go packages and provides information about them                                         |
| `go doc`             | Displays the documentation for the specified package, function, or other identifier           |
| `go mod`             | Manages Go modules                                                                            |
| `go tool`            | Provides access to various Go tools and utilities                                             |
| `go generate`        | Executes code generation tasks specified within your Go source files using special comments   |

## Examples
### `go help`
Get help for a specific Go command, e.g. `build`:
```bash
go help build
```
Get help for a specific Go topic, e.g. modules:
```bash
go help modules
```
Get help for a sub-command, e.g. `mod tidy`:
```bash
go mod help tidy
```

### `go build`
Detect data race conditions:
```
go build -race file.go
```

### `go get`
Install a specific version of a module:
```bash
go get <module-path>@<version>
```
e.g., the sampler package:
```bash
go get rsc.io/sampler@v1.3.1
```
Upgrade all direct and indirect dependencies to their latest minor or patch releases, including pre-releases:
```bash
go get -u -t -all ./...
```

### `go list`
List all metadata for each package in a json format:
```bash
go list -json ./...
```
List all packages in the standard library:
```bash
go list std
```
List all current modules and its dependencies:
```bash
go list -m all ./...
```
List all packages imported by the current package using the format flag:
```bash
go list -f '{{ join .Imports "\n"}}' ./...
```
List available versions of the sampler module:
```bash
go list -m -versions rsc.io/sampler
```

### `go mod`
Initialise a new module:
```bash
go mod init <module_path>
```
If you are working within a GitHub repository, it's a good practice to use the GitHub repository path in the module declaration.
If the project has a domain, use *example.com/module_path*. Otherwise use a placeholder for the domain.

Remove unused dependencies from the `go.mod` file:
```bash
go mod tidy
```
Create vendor directory and copy all dependencies listed in the `go.mod` file into it:
```bash
go mod vendor
```
### `go tool`
Profile and analyse the performance of your Go applications:
```bash
go tool pprof <profile-output-file>
```
Generate and analyse execution traces for your Go applications:
```bash
go tool trace <trace-output-file>
```
