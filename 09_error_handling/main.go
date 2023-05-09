package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

// Define error
var SqrtError = errors.New("Square root of negative number is not allowed.")

// Function that returns an error
func userJustFarted() error {
	return errors.New("User just farted")
}

func main() {

	file, err := os.Open("somefile.txt")
	defer fmt.Println("A deferred statement")
	if err != nil {
		// Fatal calls os.Exit(1) after printing the error and therefore deferred functions will not run
		// log.Fatalln("Error:", err)
		// Panic allows deferred functions to run
		log.Panicln(err)
	}
	defer file.Close()

	goFrom1To10()

	// We don't care
	//_ = userJustFarted()

	result, err := sqrt(16)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)

	result, err = sqrt(-1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError
	}
	return math.Sqrt(x), nil
}

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
