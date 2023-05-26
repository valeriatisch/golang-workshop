# Documentation in Go 🎓

## Official Go Documentation 📚

The official Go documentation is an excellent resource for understanding the language, standard library, and built-in packages. It can be found at [https://go.dev/doc/](https://go.dev/doc/).

### Key Sections 🔑

1. **[Packages](https://pkg.go.dev/):** Here you can find all the standard library and third-party packages. Inside of a documentation page, go to Index to find all exported functions and types of a package. 📦

2. **[Language Specification](https://go.dev/ref/spec):** Defines the syntax, semantics, and rules of the language. 📜

3. **[Effective Go](https://go.dev/doc/effective_go):** Guidelines and best practices for writing clear, idiomatic, and efficient Go code. ✅

4. **[Blog](https://go.dev/blog/):** The Go team maintains an official blog where they share updates, announcements, and interesting in-depth articles. 📰

5. **[Go by Example](https://gobyexample.com/):** Collection of annotated code snippets. 📝

## Reading Documentation 📖

### `go doc` 👩‍⚕️

To view the documentation for a package, use the following command:

```shell
go doc <package/path>
```

For example:

```shell
go doc net/http
go doc http # shorthand for net/http
```

To view the documentation for a specific function, type, or method, append the name after the package path:

```shell
go doc <package/path>.<FunctionName>
go doc <package/path>.<TypeName>.<MethodName>
```

For example:

```shell
go doc fmt.Println
go doc fmt.State.Write
```

#### Flags 🚩

- `-u` displays unexported symbols
- `-all` displays the whole documentation of a package

### `godoc` 🧑‍⚕️

- provides a web-based interface to browse and search the documentation

Install `godoc` first:

```shell
go install golang.org/x/tools/cmd/godoc@latest
```

To start the `godoc` server, run the following command:

```shell
godoc -http=:6060
```

Open your web browser and go to `http://localhost:6060` to access the documentation.

## Generating Documentation 👩‍💻

### Locally 🏠

#### Step 1: Initialize your Go module

Navigate to your project directory and initialize a new module by running the following command:

```shell
go mod init <module-path>
```

That project should contain your packages you want to document.

#### Step 2: Write GoDoc-friendly Comments ✍️

To generate the documentation, Go uses the comments written directly above the declarations of types, functions, etc. Start with a sentence that describes the package. This will be the package's summary. In general, comments must begin with the name of the item they are describing. For example, if you are documenting a function called `Sum`, your comment should start with "Sum ..." and be a full sentence.

Here's an example:

```go
// Package mypackage provides functions to manipulate strings.
package mypackage

// MyFunction is an example function.
// It does nothing particularly useful.
func MyFunction() { }
```

#### Step 3: Generate the Documentation 🚀

Run the following command inside your module (meaning in the same directory the `go.mod` file lies):

```
godoc -http=:6060
```

This command starts a local server at port

 6060 which serves the GoDoc generated documentation. Open a web browser and go to `localhost:6060/pkg/module-path/package` to see your package's documentation.

### Publishing 🌍

1. Navigate to your module's repository.
2. Tag your module with a new version number. Commit your changes and create a new tag:

   ```
   $ git commit -m "mymodule: changes for v0.1.0"
   $ git tag v0.1.0
   ```

3. Push the new tag to your repository with `git push origin v0.1.0`.
4. Make the module available by updating Go's index of modules. Set the GOPROXY environment variable to a Go proxy, and run the `go list` command to do so:

   ```
   $ GOPROXY=proxy.golang.org go list -m example.com/mymodule@v0.1.0
   ```

5. Other developers can now import your module and use `go get` to resolve it as a dependency, specifying the version if necessary:

   ```
   $ go get example.com/mymodule@v0.1.0
   ```

Remember: once published, don't change a tagged version of a module, as this will cause a security error. Instead, publish a new version.

Refer to the [official publishing guide](https://go.dev/doc/modules/publishing) for more information. 🌟
