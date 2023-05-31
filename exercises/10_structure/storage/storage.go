package storage

import types "github.com/valeriatisch/golang-workshop/exercises/10_structure/types"

// Storage defines the interface for storing and retrieving user data
type Storage interface {
	GetAllUsers() []types.User
	GetUserByID(id int) (*types.User, error)
	GetUsersByAge(age int) []types.User
	AddUser(user *types.User) error
	DeleteUser(id int) error
}
