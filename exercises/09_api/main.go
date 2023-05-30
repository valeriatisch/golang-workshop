package main

import (
	"log"
	"net/http"

	api "github.com/valeriatisch/golang-workshop/exercises/09_api/api"
)

func main() {
	// TODO: Create an instance of the memory storage
	storage := &api.MemoryStorage{}

	// TODO: Add some initial users to the storage
	storage.AddUser(&api.User{ID: 1, Username: "user1", Email: "user1@example.com", Age: 25})
	storage.AddUser(&api.User{ID: 2, Username: "user2", Email: "user2@example.com", Age: 30})
	storage.AddUser(&api.User{ID: 3, Username: "user3", Email: "user3@example.com", Age: 25})

	// TODO: Create a new server
	server := api.NewServer(storage)

	// TODO: Register the API endpoints
	http.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			server.HandleGetUsers(w, req)
		case http.MethodPost:
			server.HandlePostUser(w, req)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			server.HandleGetUserByID(w, req)
		case http.MethodDelete:
			server.HandleDeleteUser(w, req)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// TODO: Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
