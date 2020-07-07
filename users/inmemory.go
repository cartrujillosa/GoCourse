package users

import (
	"errors"
	"fmt"
)

type UserStorage interface {
	Add(user User) error
	FetchAll() []User
}

type storage struct {
	users []User
}

func NewUserStorage() UserStorage {
	// TODO add singleton pattern (sync.Once)
	return &storage{
		users: []User{},
	}
}

func (s *storage) Add(user User) error {
	for _, u := range s.users {
		if u.GetName() == user.GetName() {
			return errors.New(fmt.Sprintf("user \"%s\" already exists", user.GetName())) // TODO: send to client
		}
	}
	s.users = append(s.users, user) // TODO: use mutex to avoid race condition
	return nil
}

func (s *storage) FetchAll() []User {
	return s.users
}
