package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: Test your handlers, benchmark your storage functions, create some example tests for the docs

func TestGetUsers(t *testing.T) {
	// Create a new server with a memory storage
	storage := &MemoryStorage{}
	server := NewServer(storage)

	// Add some test users
	storage.AddUser(&User{ID: 1, Username: "user1", Email: "user1@example.com", Age: 25})
	storage.AddUser(&User{ID: 2, Username: "user2", Email: "user2@example.com", Age: 30})

	// Create a new request
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the request
	server.HandleGetUsers(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	// Decode the response body
	var users []User
	err = json.Unmarshal(recorder.Body.Bytes(), &users)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check the number of users
	expectedCount := 2
	if len(users) != expectedCount {
		t.Errorf("Expected %d users, got %d", expectedCount, len(users))
	}
}

func TestGetUserByID(t *testing.T) {
	// Create a new server with a memory storage
	storage := &MemoryStorage{}
	server := NewServer(storage)

	// Add some test users
	storage.AddUser(&User{ID: 1, Username: "user1", Email: "user1@example.com", Age: 25})
	storage.AddUser(&User{ID: 2, Username: "user2", Email: "user2@example.com", Age: 30})

	// Create a new request
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the request
	server.HandleGetUserByID(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	// Decode the response body
	var user User
	err = json.Unmarshal(recorder.Body.Bytes(), &user)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check the user details
	expectedUser := User{ID: 1, Username: "user1", Email: "user1@example.com", Age: 25}
	if user != expectedUser {
		t.Errorf("Expected user %+v, got %+v", expectedUser, user)
	}
}

func TestCreateUser(t *testing.T) {
	// Create a new server with a memory storage
	storage := &MemoryStorage{}
	server := NewServer(storage)

	// Create a new user to add
	user := User{
		ID:       1,
		Username: "user1",
		Email:    "user1@example.com",
		Age:      25,
	}

	// Encode the user as JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to encode user: %v", err)
	}

	// Create a new request with the user JSON as the request body
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the request
	server.HandlePostUser(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, recorder.Code)
	}

	// Decode the response body
	var createdUser User
	err = json.Unmarshal(recorder.Body.Bytes(), &createdUser)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check the created user details
	if createdUser != user {
		t.Errorf("Expected created user %+v, got %+v", user, createdUser)
	}
}

func TestDeleteUser(t *testing.T) {
	// Create a new server with a memory storage
	storage := &MemoryStorage{}
	server := NewServer(storage)

	// Add some test users
	storage.AddUser(&User{ID: 1, Username: "user1", Email: "user1@example.com", Age: 25})
	storage.AddUser(&User{ID: 2, Username: "user2", Email: "user2@example.com", Age: 30})

	// Create a new request
	req, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the request
	server.HandleDeleteUser(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, recorder.Code)
	}

	// Check that the user has been deleted
	_, err = storage.GetUserByID(1)
	if err == nil {
		t.Errorf("Expected user to be deleted, but user still exists")
	}
}

func TestMemoryStorage_GetAllUsers(t *testing.T) {
	storage := &MemoryStorage{
		users: []User{
			{ID: 1, Username: "user1", Email: "user1@example.com", Age: 20},
			{ID: 2, Username: "user2", Email: "user2@example.com", Age: 30},
		},
	}

	users := storage.GetAllUsers()
	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
}

func TestMemoryStorage_GetUserByID(t *testing.T) {
	storage := &MemoryStorage{
		users: []User{
			{ID: 1, Username: "user1", Email: "user1@example.com", Age: 20},
			{ID: 2, Username: "user2", Email: "user2@example.com", Age: 30},
		},
	}

	// Test existing user
	user, err := storage.GetUserByID(1)
	if err != nil {
		t.Errorf("Error occurred: %s", err.Error())
	}
	if user == nil {
		t.Error("Expected user not to be nil")
	}

	// Test non-existing user
	user, err = storage.GetUserByID(3)
	if err == nil {
		t.Error("Expected error to be returned for non-existing user")
	}
	if user != nil {
		t.Error("Expected user to be nil")
	}
}

func TestMemoryStorage_GetUsersByAge(t *testing.T) {
	storage := &MemoryStorage{
		users: []User{
			{ID: 1, Username: "user1", Email: "user1@example.com", Age: 20},
			{ID: 2, Username: "user2", Email: "user2@example.com", Age: 30},
			{ID: 3, Username: "user3", Email: "user3@example.com", Age: 20},
		},
	}

	users := storage.GetUsersByAge(20)
	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
}

func TestMemoryStorage_AddUser(t *testing.T) {
	storage := &MemoryStorage{
		users: []User{
			{ID: 1, Username: "user1", Email: "user1@example.com", Age: 20},
			{ID: 2, Username: "user2", Email: "user2@example.com", Age: 30},
		},
	}

	user := &User{ID: 3, Username: "user3", Email: "user3@example.com", Age: 25}
	storage.AddUser(user)

	users := storage.GetAllUsers()
	if len(users) != 3 {
		t.Errorf("Expected 3 users, got %d", len(users))
	}
}

func TestMemoryStorage_DeleteUser(t *testing.T) {
	storage := &MemoryStorage{
		users: []User{
			{ID: 1, Username: "user1", Email: "user1@example.com", Age: 20},
			{ID: 2, Username: "user2", Email: "user2@example.com", Age: 30},
		},
	}

	err := storage.DeleteUser(1)
	if err != nil {
		t.Errorf("Error occurred: %s", err.Error())
	}

	users := storage.GetAllUsers()
	if len(users) != 1 {
		t.Errorf("Expected 1 user, got %d", len(users))
	}
}

func BenchmarkMemoryStorage_GetAllUsers(b *testing.B) {
	storage := &MemoryStorage{
		users: []User{
			// Add sample users here
		},
	}

	for i := 0; i < b.N; i++ {
		_ = storage.GetAllUsers()
	}
}

func BenchmarkMemoryStorage_GetUserByID(b *testing.B) {
	storage := &MemoryStorage{
		users: []User{
			// Add sample users here
		},
	}

	for i := 0; i < b.N; i++ {
		_, _ = storage.GetUserByID(1)
	}
}

func BenchmarkMemoryStorage_GetUsersByAge(b *testing.B) {
	storage := &MemoryStorage{
		users: []User{
			// Add sample users here
		},
	}

	for i := 0; i < b.N; i++ {
		_ = storage.GetUsersByAge(20)
	}
}

func BenchmarkMemoryStorage_AddUser(b *testing.B) {
	storage := &MemoryStorage{
		users: []User{
			// Add sample users here
		},
	}

	user := &User{
		// Create a sample user here
	}

	for i := 0; i < b.N; i++ {
		storage.AddUser(user)
	}
}

func BenchmarkMemoryStorage_DeleteUser(b *testing.B) {
	storage := &MemoryStorage{
		users: []User{
			// Add sample users here
		},
	}

	for i := 0; i < b.N; i++ {
		_ = storage.DeleteUser(1)
	}
}

func ExampleMemoryStorage_GetAllUsers() {
	storage := &MemoryStorage{
		users: []User{
			{ID: 1, Username: "user1", Email: "user1@example.com", Age: 25},
			{ID: 2, Username: "user2", Email: "user2@example.com", Age: 30},
		},
	}

	users := storage.GetAllUsers()
	for _, user := range users {
		fmt.Println(user)
	}
	// Output:
	// {1 user1 user1@example.com 25}
	// {2 user2 user2@example.com 30}
}

func ExampleMemoryStorage_GetUserByID() {
	storage := &MemoryStorage{
		users: []User{
			{ID: 1, Username: "user1", Email: "user1@example.com", Age: 25},
			{ID: 2, Username: "user2", Email: "user2@example.com", Age: 30},
		},
	}

	user, err := storage.GetUserByID(1)
	if err != nil {
		fmt.Println("User not found")
		return
	}

	fmt.Println(user)
	// Output: &{1 user1 user1@example.com 25}
}

func ExampleMemoryStorage_GetUsersByAge() {
	storage := &MemoryStorage{
		users: []User{
			{ID: 1, Username: "user1", Email: "user1@example.com", Age: 25},
			{ID: 2, Username: "user2", Email: "user2@example.com", Age: 30},
		},
	}

	users := storage.GetUsersByAge(30)
	for _, user := range users {
		fmt.Println(user)
	}
	// Output: {2 user2 user2@example.com 30}
}

func ExampleMemoryStorage_AddUser() {
	storage := &MemoryStorage{
		users: []User{
			{ID: 1, Username: "user1", Email: "user1@example.com", Age: 25},
		},
	}

	user := &User{
		ID:       2,
		Username: "user2",
		Email:    "user2@example.com",
		Age:      30,
	}

	storage.AddUser(user)

	user, err := storage.GetUserByID(2)
	if err != nil {
		fmt.Println("User not found")
		return
	}
	fmt.Println(user)
	// Output: &{2 user2 user2@example.com 30}
}

func ExampleMemoryStorage_DeleteUser() {
	storage := &MemoryStorage{
		users: []User{
			{ID: 1, Username: "user1", Email: "user1@example.com", Age: 25},
			{ID: 2, Username: "user2", Email: "user2@example.com", Age: 30},
		},
	}

	err := storage.DeleteUser(1)
	if err != nil {
		fmt.Println("User not found")
		return
	}
	// Output:
}
