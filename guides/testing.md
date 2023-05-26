# Testing, Benchmarking & Examples in Go 🔬

This guide will help you understand how to test your Go applications, including unit testing, creating examples in documentation, and benchmarking for performance optimization. 🧪

Go has a built-in testing package named `testing`. This package provides a framework for writing unit tests and benchmark tests.

## Writing a Unit Test ✍️

Let's consider a function `Sum` in a file named `mathutil.go`:

```go
func Sum(x int, y int) int {
    return x + y
}
```

To test the `Sum` function, create a new file `mathutil_test.go` in the same package.
The file **must** end with `_test.go` to be recognized as a test file.

```go
func TestSum(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}
```

In the test function (`TestSum`), the argument is a pointer to `testing.T`, which provides methods for reporting test failures and logging additional information.
If the test fails, the `t.Errorf` is called to log the error message.

To run the test, use the following commands:
```shell
go test
```

### Table-Driven Tests 📋

When you need to test a function with various values at once, you can use structs to define the test cases.

Here's an example:

```go
func TestSum(t *testing.T) {
    testCases := []struct {
        x    int
        y    int
        sum  int
    }{
        {1, 1, 2},
        {1, 2, 3},
        {2, 2, 4},
        {5, 5, 10},
    }

    for _, tc := range testCases {
        total := Sum(tc.x, tc.y)
        if total != tc.sum {
            t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", tc.x, tc.y, total, tc.sum)
        }
    }
}
```

## Documenting Example Tests 📝

Go encourages developers to write examples as a form of documentation.
Examples are both tests and documentation.
Users can see how to use the function, and the `go test` tool can verify that the function works as expected.

Here's an example:

```go
func ExampleSum() {
    total := Sum(4, 5)
    fmt.Println(total)
    // Output: 9
}
```

The special comment `// Output: ` is used to indicate the expected output.
The test passes if the function's output matches the expected output. ✅<br>
If you run `godoc -http=:6060` and navigate to the function, you will see the example in the documentation.

## Benchmarking ⏱️

Benchmarking measures the performance of your code.
It involves running a piece of code many times and measuring how long it takes. ⏰<br>
To write a benchmark test, the function name must start with `Benchmark`, take a pointer to `testing.B`, and call the `b.N` within a loop.

Here is an example:

```go
func BenchmarkSum(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Sum(4, 5)
    }
}
```

Run the benchmark test:

```shell
go test -bench=.
```

The output shows the number of iterations and the used time per operation.

To get memory allocation statistics too, use:

```shell
go test -bench=. -benchmem
```

The output shows the number of allocations and the used memory per operation.

## Code Coverage 📊

Code coverage measures how much of your code is covered by tests.
High code coverage usually means fewer bugs.
Go has built-in support for coverage analysis.

### Creating a Coverage Profile

To create and print out your coverage profile:
```shell
go test -cover
```

If you want to create a coverage profile file, use:

```shell
go test -coverprofile=cover.out
```

You can name the file however you want.

This creates a file named `cover.out` in your directory, which contains the coverage profile.
It can be analyzed and give you a more detailed look at what functions and statements are maybe not covered yet.

### Visualizing Coverage 🌈

To generate a pretty HTML report, use:

```shell
go tool cover -html=cover.out
```

This will open up a webpage in your browser displaying a detailed, color-coded overview of your codebase. 
