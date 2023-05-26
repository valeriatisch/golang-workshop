package main

import (
	"net"
)

// When you're using net to create a TCP server, you're working at a lower level,
// where you have to manually handle the reading and writing of data to the TCP connection.
// This means you have more control, but you also have to do more work.

func main() {
	// Listen on TCP port 8080

	for {
		// Wait for a connection.

		// Handle connections in a new goroutine.
	}
}

func handleRequest(conn net.Conn) {

	// Make a buffer to hold incoming data

	// Convert the incoming message to a string, removing the newline character

	// Send a response back to the person contacting us.
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
