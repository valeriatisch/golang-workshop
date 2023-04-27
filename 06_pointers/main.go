package main

import "fmt"

type Player struct {
	matches int
}

func main() {
	/**
	 * Memory allocation & reference
	 * - A pointer is a variable that stores the memory address of another variable.
	 * - A pointer is declared by using the * operator before the type of the variable it points to.
	 * - A pointer is created by using the & operator before the variable name.
	 * - A pointer is dereferenced by using the * operator before the pointer name.
	 */

	// Pointers with functions
	// Pass by value

	// Pass by reference

	// Pointers with structs

	/* Make
	 * - The make function is used to allocate memory for a slice, map, or channel.
	 */

	// Length of a slice is the number of elements it contains.
	// Capacity of a slice is the maximum number of elements it can hold before it needs to be resized.
	// Capacity of a slice can be larger than its length, which means it can hold additional elements without needing to be resized.

	/* New
	 * - The new function allocates memory for a zero value of a type and returns a pointer to it.
	 */

	// Unnecessary complex

	// Idiomatic way

}

func incrementValue(x int) {
	fmt.Println("Took:", x)
	x += 1
	fmt.Println("Incremented:", x)
}

func incrementPtr(x *int) {
	fmt.Println("Took:", *x)
	*x += 1
	fmt.Println("Incremented:", *x)
}
