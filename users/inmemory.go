package users

import "errors"

type UserStorage interface {
	Add(user User) error
	FetchAll() []User
}

type storage struct {
	users []User
}

func NewUserStorage() UserStorage {
	return storage{
		users: []User{},
	}
}

func (s storage) Add(user User) error {
	for _, u := range s.users {
		if u.GetName() == user.GetName() {
			return errors.New("user \"%s\" already exists")
		}
	}
	s.users = append(s.users, user)
	return nil
}

func (s storage) FetchAll() []User {
	return s.users
}
