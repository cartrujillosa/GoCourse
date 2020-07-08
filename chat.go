package main

import (
	"errors"
	"fmt"
	"log"
	"net"

	userslib "github.com/cartrujillosa/GoCourse/project/users"
)

type Chat interface {
	SendMessage(from userslib.User, msg string)
	RegisterUser(user userslib.User) error
	Listener() net.Listener
	Broadcast(msg string)
	Close()
}

type chat struct {
	users    []userslib.User
	listener net.Listener
}

func NewChat(port string) Chat {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server: %s", err)
		return nil
	}
	log.Printf("chat server started on %s", port)

	return &chat{ // TODO add singleton pattern (sync.Once)
		listener: listener,
		users:    []userslib.User{},
	}
}

func (c *chat) SendMessage(from userslib.User, msg string) {
	for _, anyUser := range c.users {
		if anyUser.RemoteAddr() != from.RemoteAddr() {
			msg = fmt.Sprintf("%s says: %s", from.Name(), msg)
			anyUser.ReceiveMessage(msg)
		}
	}
	return
}

func (c *chat) Broadcast(msg string) {
	for _, anyUser := range c.users {
		anyUser.ReceiveMessage(msg)
	}
	return
}

func (c *chat) RegisterUser(user userslib.User) error {
	for _, u := range c.users {
		if u.Name() == user.Name() {
			return errors.New(fmt.Sprintf("user \"%s\" already exists", user.Name())) // TODO: send to client
		}
	}
	c.users = append(c.users, user) // TODO: use mutex to avoid race condition
	return nil
}

func (c *chat) Listener() net.Listener {
	return c.listener
}

func (c *chat) Close() {
	c.listener.Close()
}
