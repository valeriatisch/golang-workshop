package main

import (
	"fmt"
)

func main() {
	// declare a variable & assign a value to it with :=

	// re-assign the variable

	// can't assign a value to a variable with a different type
	// x = "hello"
	// x = false

	// declare a variable and assign a value with var

	// declare a variable with a type

	// declare multiple variables at once

	// Shadowing

	// Zero value: default value for a variable
	// int: 0
	// float: 0.0
	// string: ""
	// bool: false
	// nil: functions, pointers, slices etc.

	// String formatting

	// ASCII

	// A string is immutable, you can assign a new value but can't change the bytes

	// Different values for the same variable
	fmt.Printf("%b \n", 42)         // binary
	fmt.Printf("%c \n", 42)         // character
	fmt.Printf("%d \n", 42)         // decimal
	fmt.Printf("%.2f \n", 43.44635) // floats
	fmt.Printf("%o \n", 42)         // octal
	fmt.Printf("%q \n", 42)         // quoted string
	fmt.Printf("%#x \n", 42)        // hex
	fmt.Printf("%#X \n", 42)        // hex
	fmt.Printf("%U \n", 42)         // unicode
	fmt.Printf("%v \n", 42)         // normal value

	// Creating your own type

	// Conversion

	// Logical operators
	var fls bool = false
	var tr bool = true
	fmt.Println(fls && tr) // AND
	fmt.Println(fls || tr) // OR
	fmt.Println(!fls)      // NOT

	// Byte is alias for uint8

	// Rune is alias for int32

	// Constants
	// Typed

	// define multiple constants

	// Untyped

	// Iota

}

// Can't do this in outer scope/package-level:
// x := 1
// but this is ok:
var globalVar = 2
