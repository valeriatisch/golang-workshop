package main

import (
	"fmt"
	"math/rand"
)

func main() {
	eveCh := make(chan int)
	oddCh := make(chan int)
	quitCh := make(chan int)

	go send(eveCh, oddCh, quitCh)

	receive(eveCh, oddCh, quitCh)
}

func send(eveCh, oddCh, quitCh chan<- int) {
	for i := 0; i < 10; i++ {
		randNr := rand.Intn(100)
		if randNr%2 == 0 {
			eveCh <- randNr
		} else {
			oddCh <- randNr
		}
	}
	close(eveCh)
	close(oddCh)
	quitCh <- 0
}

func receive(eveCh, oddCh, quitCh <-chan int) {
	// Select statement
	for {
		select {
		case v, ok := <-eveCh:
			fmt.Println("The number", v, "is even", ok)
		case v, ok := <-oddCh:
			fmt.Println("The number", v, "is odd", ok)
		case v, ok := <-quitCh:
			fmt.Println("Quitting...", v, ok)
			return
		}
	}
}
