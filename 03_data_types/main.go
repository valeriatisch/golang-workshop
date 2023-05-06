package main

import (
	"fmt"
)

func main() {
	// declare a variable & assign a value to it with :=
	a := 2
	fmt.Println(a)

	// re-assign the variable
	a = 3

	// can't assign a value to a variable with a different type
	// x := "hello"
	// x = false

	// declare a variable and assign a value with var
	var b int
	b = 4

	// declare a variable with a type
	var c = 4
	fmt.Println(b, c)

	// declare multiple variables at once
	var d, f int = 5, 6
	var (
		g         = 7
		h float64 = 8.5
	)
	fmt.Println(d, f, g, h)

	// Shadowing
	var globalVar = "test"
	fmt.Println(globalVar)

	// Zero value: default value for a variable
	// int: 0
	// float: 0.0
	// string: ""
	// bool: false
	// nil: functions, pointers, slices etc.

	// String formatting
	var str string = `He said: "Hello"`
	fmt.Println(str)

	// A string is immutable, you can assign a new value but can't change the bytes
	str = "new string here"
	// str[0] = "N"

	// Different values for the same variable
	fmt.Printf("%b \n", 42)         // binary
	fmt.Printf("%c \n", 42)         // character
	fmt.Printf("%d \n", 42)         // decimal
	fmt.Printf("%.4f \n", 43.44635) // floats
	fmt.Printf("%o \n", 42)         // octal
	fmt.Printf("%q \n", 42)         // quoted string
	fmt.Printf("%#x \n", 42)        // hex
	fmt.Printf("%#X \n", 42)        // hex
	fmt.Printf("%U \n", 42)         // unicode
	fmt.Printf("%v \n", 42)         // normal value

	// Creating your own type
	type myType int
	var myVar myType = 42
	fmt.Printf("%T \n", myVar)

	// Conversion
	var myVar2 int = int(myVar)
	fmt.Printf("%T \n", myVar2)

	// Logical operators
	var fls bool = false
	var tr bool = true
	fmt.Println(fls && tr) // AND
	fmt.Println(fls || tr) // OR
	fmt.Println(!fls)      // NOT

	// Byte is alias for uint8
	var bt byte = 1
	fmt.Printf("%T \n", bt)

	// Rune is alias for int32
	var rn rune = 'a'
	fmt.Printf("%T \n", rn)

	// Constants
	// Typed
	const aConst int = 42

	// define multiple constants
	// Untyped
	const (
		const1 = 1
		const2 = 2
	)

	// Iota
	type Weekday int
	const (
		monday Weekday = iota
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)
	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}

// Can't do this in outer scope/package-level:
// x := 1
// but this is ok:
var globalVar = 2
