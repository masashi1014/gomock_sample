package repository

import (
	"math/rand"
	"time"
)

type User struct {
	ID int64
	Name string
	IsRegistered bool
}

type UserRepository interface {
	Create() (*User, error)
	GetUserByID(id int64) *User
}

func NewUserRepository() UserRepository{
	return &userRepository{}
}

type userRepository struct{}

func (re *userRepository) Create() (*User, error){
	user := &User{
		ID: 1,
		Name: "DummyUser",
		IsRegistered: true,
	}

	return user, nil
}

func (re *userRepository) GetUserByID(id int64) *User{
	user := &User{
		ID: id,
		Name: "DummyUser",
	}

	// IsRegister is random
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		user.IsRegistered = false
	} else {
		user.IsRegistered = true
	}

	return user
}
