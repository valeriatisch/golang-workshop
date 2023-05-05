package main

import (
	"fmt"
)

func main() {
	// This is a bidirectional channel
	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	// There are unidirectional channels too
	// This is a send-only channel
	// ch_send := make(chan<- int, 2)
	// fmt.Printf("%T\n", ch_send)

	// This is a receive-only channel
	// ch_rec := make(<-chan int, 2)
	// fmt.Printf("%T\n", ch_rec)

	// The function takes the channel as send-only
	// go foo(ch, 5)
	// The function takes the channel as receive-only
	// It also blocks until the channel receives a value
	// bar(ch)

	// Ranges
	// - A range is a loop that iterates over a channel until it is closed.
	// - A range will block by default if the channel is empty unless it is closed.
	/*
		go func() {
			fmt.Println("Putting values on the channel...")
			for i := 0; i< 10; i++ {
				ch <- i
			}
			close(ch)
		}()

		v, ok := <-ch
		fmt.Println(v, ok)

		fmt.Println("Now reading values from the channel...")
		for v := range ch {
			fmt.Println(v)
		}
		v, ok = <-ch
		fmt.Println(v, ok)
	*/
}

func foo(ch chan<- int, x int) {
	ch <- x * 5
}

func bar(ch <-chan int) {
	fmt.Println(<-ch)
}
