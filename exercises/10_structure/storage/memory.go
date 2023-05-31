package storage

import (
	"errors"

	types "github.com/valeriatisch/golang-workshop/exercises/10_structure/types"
)

// MemoryStorage is an in-memory storage implementation
type MemoryStorage struct {
	users []types.User
}

// GetAllUsers returns all users
func (ms *MemoryStorage) GetAllUsers() []types.User {
	return ms.users
}

// GetUserByID retrieves a user by ID
func (ms *MemoryStorage) GetUserByID(id int) (*types.User, error) {
	for _, user := range ms.users {
		if user.ID == id {
			return &user, nil
		}
	}
	// return others 0 and error
	return nil, errors.New("User not found")
}

// GetUsersByAge retrieves users by age
func (ms *MemoryStorage) GetUsersByAge(age int) []types.User {
	var matchingUsers []types.User
	for _, user := range ms.users {
		if user.Age == age {
			matchingUsers = append(matchingUsers, user)
		}
	}
	return matchingUsers
}

// AddUser adds a new user to the storage
func (ms *MemoryStorage) AddUser(user *types.User) error {
	user.ID = len(ms.users) + 1 // Generate a unique ID
	ms.users = append(ms.users, *user)
	return nil
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
	return nil // User not found
}
