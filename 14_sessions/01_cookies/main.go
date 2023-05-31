package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

var users map[string]int

func main() {
	users = make(map[string]int)

	http.HandleFunc("/", handleHome)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, req *http.Request) {

	uuidCookie, err := req.Cookie("uuid")
	// check if user already visited website
	if err == http.ErrNoCookie {
		// create UUID Cookie
		uid := uuid.New().String() // new user
		uuidCookie = &http.Cookie{
			Name:  "uuid",
			Value: uid,
		}
		http.SetCookie(w, uuidCookie)

		users[uid] = 0 // add user to map

		fmt.Fprintln(w, "Your new Cookie has been created:", uuidCookie.Value)
	}

	countCookie, err := req.Cookie("countCookie") // retrieve countCookie
	if err == http.ErrNoCookie {
		// create Count Cookie
		countCookie = &http.Cookie{
			Name:  "countCookie",
			Value: "0",
			Path: "/users",
		}
		http.SetCookie(w, countCookie)
		return
	}

	cookieValue, _ := strconv.Atoi(countCookie.Value) // convert str to int
	cookieValue ++ // increase counter
	users[uuidCookie.Value] = cookieValue // set Counter
	countCookie.Value = strconv.Itoa(cookieValue)
	http.SetCookie(w, countCookie)
	
	fmt.Fprintln(w, "Your cookieCounter is:", cookieValue)
	// cookie, err := req.Cookie("my-cookie")
	// if err == http.ErrNoCookie {
	// 	cookie = &http.Cookie{
	// 		Name:  "my-cookie",
	// 		Value: "first value",
	// 	}
	// 	http.SetCookie(w, cookie)
	// 	return
	// }
	// cookie = &http.Cookie{
	// 	Name:  "my-cookie",
	// 	Value: "second value",
	// }
	// http.SetCookie(w, cookie)
}

// Goal:
// 1. Create a cookie with a unique ID for each user & store it in a map (as key)
// 2. Create a cookie that counts how often the user visits the page and store the count in a map (as value)
// 3. Create a route /user which will display the UUID and the count for the user
