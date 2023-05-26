package api

import (
	"encoding/json"
	"net/http"
)

// TODO: Create a User struct with some fields, add json tags

// TODO: Create a Storage interface which define storing and retrieving user data (e.g. GetAllUsers, AddUser, etc.)

// TODO: Create a MemoryStorage struct which will implement the Storage interface

// TODO: Implement all functions from the Storage interface for the MemoryStorage struct

// GetUserByID retrieves a user by ID

// GetUsersByAge retrieves users by age

// AddUser adds a new user to the storage

// DeleteUser removes a user from the storage by ID

// TODO: Create a Server struct with a Storage field

// TODO: Create a NewServer func that takes a storage and creates a new instance of the Server and returns

// TODO: Implement some Handle functions for the Server struct (e.g., HandleGetUsers, HandlePostUser, etc.)

// respondWithJSON sends a JSON response with the specified status code
// You can use this function to send responses from your Handle functions
// The data parameter expects a struct or slice which will be encoded as JSON
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
