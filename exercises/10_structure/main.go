package main

import (
	"log"
	"net/http"

	api "github.com/valeriatisch/golang-workshop/exercises/10_structure/api"
	storage "github.com/valeriatisch/golang-workshop/exercises/10_structure/storage"
	types "github.com/valeriatisch/golang-workshop/exercises/10_structure/types"
)

func main() {
	// Create an instance of the memory storage
	storage := &storage.MemoryStorage{}

	// Add some initial users
	storage.AddUser(&types.User{ID: 1, Username: "user1", Email: "user1@example.com", Age: 25})
	storage.AddUser(&types.User{ID: 2, Username: "user2", Email: "user2@example.com", Age: 30})
	storage.AddUser(&types.User{ID: 3, Username: "user3", Email: "user3@example.com", Age: 25})

	// Create a new server
	server := api.NewServer(storage)

	// Register the API endpoints
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

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
