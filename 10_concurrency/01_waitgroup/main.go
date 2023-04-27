package main

import (
	"fmt"
	"runtime"
	"time"
)

// There is always an init() running before main()

func main() {
	// Computer specs
	fmt.Println("\nOS:\t\t", runtime.GOOS)
	fmt.Println("Arch:\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	start := time.Now()

	// A goroutine is a lightweight thread managed by the Go runtime.
	// A goroutine is capable of running concurrently with other functions.
	// Launch a goroutine

	// It takes some time to launch a goroutine, during which the main() function continues to run.
	// The main() function might finish before the goroutine does.
	// To avoid this, we can use the WaitGroup.
	// Add(int): Add increases the WaitGroup counter by the given delta.
	// Done(): Done decrements the WaitGroup counter by 1.
	// Wait(): Wait blocks until the WaitGroup counter is zero.

	// Goroutines in for-loop

	fmt.Println("\nThat took", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "John"
}

func fetchUserLikes(userName string) int {
	time.Sleep(time.Millisecond * 150) // cannot be faster than 150ms overall
	return 10
}

func fetchUserMatch(userName string) string {
	time.Sleep(time.Millisecond * 100)
	return "Jane"
}
