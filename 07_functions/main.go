package main

import (
	"fmt"
	"os"
)

type Person struct {
	Name string
}

// Methods
// Methods are functions that are attached to a particular type.
func (p Person) breath() {
	fmt.Println(p.Name, "is breathing.")
}

// Function
func breath(p Person) {
	fmt.Println(p.Name, "is breathing.")
}

func main() {
	// Defer an action
	// Defer is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup.
	defer fmt.Println("This is the first line in our main func.")
	fmt.Println("This is the second line in our main func.")

	// open and close a file file.txt
	f, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Variadic parameters allows you to pass an arbitrary number of parameters to a function
	println(sum(1, 3, 4, 5, 6))
	println(sum(3))

	// Methods
	p := Person{Name: "John"}
	p.breath()

	// Anonymous functions
	// An anonymous function is a function without a name.
	func() {
		x := 0
		x += 1
		fmt.Println(x)
	}()

	// with an argument
	func(x int) {
		x += 1
		fmt.Println(x)
	}(0)

	// with a return value
	y := func(x int) int {
		x += 1
		return x
	}(0)
	println(y)

	// Func expressions
	// A func expression is a function without a name that is assigned to a variable.
	// The function can be called using the variable name.
	// The function can be passed as an argument to another function or returned from another function.
	fn := func() {
		fmt.Println("I am a func expression.")
	}
	fn()

	anotherFn := func(x int, s string) string {
		return fmt.Sprintf("%d %s", x, s)
	}
	callbackFunc(anotherFn)

	fmt.Println(returnsMultipleValues())

	// Returning a function
	x := returnFuncExpression()
	x()

	z := returnFuncExpression()
	z()

	// Recursion
	println(factorial(3))
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return factorial(x-1) * x
}

// Variadic parameters function
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Callbacks
// Function taking a function as an argument
func callbackFunc(f func(int, string) string) {
	f(1, "Hello")
}

// Function with multiple return values
func returnsMultipleValues() (int, string, bool) {
	return 1, "some string here", true
}

// Returning a function
func returnFunc() func() {
	return func() {
		println("I am a returned anonymous function.")
	}
}

// TODO: Please define a function which will return a func expression.
func returnFuncExpression() func() {
	fn := func() {
		fmt.Println("Whatever")
	}
	return fn
}
