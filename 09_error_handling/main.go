package main

import (
	"fmt"
)

// Define error

// Function that returns an error

// Recover

func main() {

	// Fatal calls os.Exit(1) after printing the error and therefore deferred functions will not run

	// Panic allows deferred functions to run

	goFrom1To10()

	// Recover allows us to exit the program gracefully

}

// Panic and recover in function
func goFrom1To10() {
	for i := 0; i <= 10; i++ {
		// Here we recover and continue with the loop because it only affects the function in current iteration.
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("panic occured: ", r)
				}
			}()

			if i == 2 {
				panic("got 2")
			}
			fmt.Println(i)
		}()
	}
}
