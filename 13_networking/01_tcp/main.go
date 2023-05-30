package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// When you're using net to create a TCP server, you're working at a lower level,
// where you have to manually handle the reading and writing of data to the TCP connection.
// This means you have more control, but you also have to do more work.

func main() {
	// Listen on TCP port 8080
	listener, err := net.Listen("tcp", ":8080")
	checkError(err)
	defer listener.Close()

	for {
		// Wait for a connection.
		conn, err := listener.Accept()
		checkError(err)

		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	// Make a buffer to hold incoming data
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	checkError(err)

	// Convert the incoming message to a string, removing the newline character
	clientMsg := strings.TrimSuffix(string(buffer), "\n")
	fmt.Println("Message received:", clientMsg)

	// Send a response back to the person contacting us.
	conn.Write([]byte("ACK. Received your message, client ;)\n"))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
