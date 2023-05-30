package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// Create a User struct with some fields, add json tags
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

// Create a Storage interface which define storing and retrieving user data (e.g. GetAllUsers, AddUser, etc.)
type Storage interface {
	GetAllUsers() []User
	GetUserByID(id int) (*User, error)
	GetUsersByAge(age int) []User
	AddUser(user *User)
	DeleteUser(id int) error
}

// Create a MemoryStorage struct which will implement the Storage interface
type MemoryStorage struct {
	users []User
}

// Implement all functions from the Storage interface for the MemoryStorage struct
func (ms *MemoryStorage) GetAllUsers() []User {
	return ms.users
}

// GetUserByID retrieves a user by ID
func (ms *MemoryStorage) GetUserByID(id int) (*User, error) {
	for _, user := range ms.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

// GetUsersByAge retrieves users by age
func (ms *MemoryStorage) GetUsersByAge(age int) []User {
	var matchingUsers []User
	for _, user := range ms.users {
		if user.Age == age {
			matchingUsers = append(matchingUsers, user)
		}
	}
	return matchingUsers
}

// AddUser adds a new user to the storage
func (ms *MemoryStorage) AddUser(user *User) {
	user.ID = len(ms.users) + 1 // Generate a unique ID
	ms.users = append(ms.users, *user)
}

// DeleteUser removes a user from the storage by ID
func (ms *MemoryStorage) DeleteUser(id int) error {
	for i, user := range ms.users {
		if user.ID == id {
			// Remove the user from the slice
			ms.users = append(ms.users[:i], ms.users[i+1:]...)
			return nil
		}
	}
	return errors.New("User not found")  // User not found
}

// Create a Server struct with a Storage field
type Server struct {
	storage Storage
}

// Create a NewServer func that takes a storage and creates a new instance of the Server and returns
func NewServer(storage Storage) *Server {
	return &Server{
		storage: storage,
	}
}

// Implement some Handle functions for the Server struct (e.g., HandleGetUsers, HandlePostUser, etc.) 
// HandleGetUsers retrieves all users
func (s *Server) HandleGetUsers(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ageStr := req.URL.Query().Get("age")
	if ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}
		users := s.storage.GetUsersByAge(age)
		respondWithJSON(w, http.StatusOK, users)
		return
	}

	users := s.storage.GetAllUsers()
	respondWithJSON(w, http.StatusOK, users)
}

// HandleGetUserByID retrieves a user by ID
func (s *Server) HandleGetUserByID(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := req.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := s.storage.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

// HandlePostUser creates a new user
func (s *Server) HandlePostUser(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	s.storage.AddUser(&user)

	respondWithJSON(w, http.StatusCreated, user)
}

// HandleDeleteUser deletes a user by ID
func (s *Server) HandleDeleteUser(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := req.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = s.storage.DeleteUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// respondWithJSON sends a JSON response with the specified status code
// You can use this function to send responses from your Handle functions
// The data parameter expects a struct or slice which will be encoded as JSON
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
