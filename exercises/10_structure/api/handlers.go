package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	types "github.com/valeriatisch/golang-workshop/exercises/10_structure/types"
)


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

	var user types.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = s.storage.AddUser(&user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

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
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
