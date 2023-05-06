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
	x := 10
	var xPtr *int
	xPtr = &x
	fmt.Println("x:", xPtr)
	fmt.Println(*xPtr)

	// Pointers with functions
	// Pass by value
	incrementValue(x)
	fmt.Println(x)

	// Pass by reference
	incrementPtr(xPtr)
	fmt.Println(x)

	/* Make
	 * - The make function is used to allocate memory for a slice, map, or channel.
	 */

	slice := make([]int, 0, 5)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, 1, 3, 5, 6, 7)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, 6)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, 1, 3, 5, 6, 7)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// Length of a slice is the number of elements it contains.
	// Capacity of a slice is the maximum number of elements it can hold before it needs to be resized.
	// Capacity of a slice can be larger than its length, which means it can hold additional elements without needing to be resized.

	/* New
	 * - The new function allocates memory for a zero value of a type and returns a pointer to it.
	 */
	slice2 := new([]int)
	fmt.Println(slice2)

	number := new(int)
	fmt.Println(*number)
	*number = 10
	fmt.Println(*number)

	// Unnecessary complex
	var p *[]int = new([]int)
	*p = append(*p, 2)
	fmt.Println(*p)

	// Idiomatic way
	v := make([]int, 0, 10)
	fmt.Println(v)
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
