package users

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
)

type User interface {
	Name() string
	Location() string
	ReceiveMessage(msg string)
	GetMessage() (string, error)
	RemoteAddr() string
	Conn() net.Conn
}

type user struct {
	name     string
	location string
	conn     net.Conn
}

func NewUser(name, location string, conn *net.Conn) (User, error) {
	if len(name) == 0 || len(location) == 0 {
		return nil, errors.New("name and location are mandatory")
	}
	return &user{
		name:     name,
		location: location,
		conn:     *conn,
	}, nil
}

func (u user) Name() string {
	return u.name
}

func (u user) Location() string {
	return u.location
}

func (u user) ReceiveMessage(msg string) {
	_, err := io.WriteString(u.conn, msg)
	if err != nil {
		log.Print(err)
	}
	return
}

func (u user) GetMessage() (string, error) {
	return bufio.NewReader(u.conn).ReadString('\n')
}

func (u user) RemoteAddr() string {
	return u.conn.RemoteAddr().String()
}

func (u user) Conn() net.Conn {
	return u.conn
}
