package lib

import (
	"errors"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:email"`
}

type UserRepository interface {
	GetAll() []User
	Create(User) error
	Validate(User) error
}
type InMemoryUserRepository struct {
	DataStore []User
}

func (repo *InMemoryUserRepository) GetAll() []User {
	return repo.DataStore
}
func (repo *InMemoryUserRepository) Create(user User) error {
	err := repo.Validate(user)
	if err != nil {
		return err
	}
	repo.DataStore = append(repo.DataStore, user)
	return nil
}
func (repo *InMemoryUserRepository) Validate(user User) error {
	for _, u := range repo.DataStore {
		if u.Email == user.Email {
			return errors.New("The Email is already exists")
		}
	}
	return nil
}
func NewInMemoryUserRepo() *InMemoryUserRepository {
	return &InMemoryUserRepository{DataStore: []User{}}
}
