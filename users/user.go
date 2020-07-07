package users

import (
	"bufio"
	"errors"
	"io"
	"net"
)

type User interface {
	GetName() string
	ReceiveMessage(msg string)
	SendMessage() (string, error)
	RemoteAddr() string
}

type user struct {
	name     string
	location string
	conn     net.Conn
}

func NewUser(name, location string, conn net.Conn) (User, error) {
	if len(name) == 0 || len(location) == 0 {
		return nil, errors.New("name and location are mandatory")
	}
	return &user{
		name:     name,
		location: location,
		conn:     conn,
	}, nil
}

func (u user) GetName() string {
	return u.name
}

func (u user) ReceiveMessage(msg string) {
	io.WriteString(u.conn, msg)
	return
}

func (u user) SendMessage() (string, error) {
	return bufio.NewReader(u.conn).ReadString('\n')
}

func (u user) RemoteAddr() string {
	return u.conn.RemoteAddr().String()
}
