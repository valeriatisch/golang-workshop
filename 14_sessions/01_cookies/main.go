package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/setGeneral", handleSetCookie)
	http.HandleFunc("/", handleMyCookie)
	http.HandleFunc("/set", handleSetMyCookie)
	http.HandleFunc("/read", handleReadMyCookie)

	http.ListenAndServe(":8080", nil)
}

func handleSetCookie(w http.ResponseWriter, req *http.Request) {
	// Set a "general" cookie with some value

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
}

func handleSetMyCookie(w http.ResponseWriter, req *http.Request) {
	// Set the "my-cookie" cookie with an initial value of 1

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
}

func handleMyCookie(w http.ResponseWriter, req *http.Request) {
	// Check if the cookie exists

	// If the cookie doesn't exist, redirect the user to the "/set" path

	// If the cookie exists, extract the its value

	// Update the cookie with the new count

	// Display the visit count to the user
	
}

func handleReadMyCookie(w http.ResponseWriter, req *http.Request) {
	// Read and display the value of the "my-cookie" cookie
	c, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #1:", c)
	}

	// Read and display the value of the "general" cookie
	c1, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #2:", c1)
	}
}
