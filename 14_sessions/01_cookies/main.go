package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHome)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, req *http.Request) {

}

// Goal:
// 1. Create a cookie with a unique ID for each user & store it in a map (as key)
// 2. Create a cookie that counts how often the user visits the page and store the count in a map (as value)
// 3. Create a route /user which will display the UUID and the count for the user
