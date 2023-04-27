package main

// Mantra: Do not communicate by sharing memory, share memory by communicating.

func main() {
	// This is a bidirectional channel

	// There are unidirectional channels too
	// This is a send-only channel

	// This is a receive-only channel

	// The function takes the channel as send-only

	// The function takes the channel as receive-only
	// It also blocks until the channel receives a value

	// Ranges
	// - A range is a loop that iterates over a channel until it is closed.
	// - A range will block by default if the channel is empty unless it is closed.

}

// Function that takes the channel as send-only

// Function that takes the channel as receive-only
