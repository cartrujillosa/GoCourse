package main

import (
	"fmt"
	"log"
	"net"

	userslib "github.com/cartrujillosa/GoCourse/project/users"
)

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server: %s", err)
	}
	defer listener.Close()

	log.Printf("Chat server started on :8888")

	users := userslib.NewUserStorage()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err)
			continue
		}

		user, err := userslib.NewUser("carla", "canarias", conn)
		if err != nil {
			log.Println(err)
			continue
		}

		if err := users.Add(user); err != nil {
			log.Println(err)
			continue
		}
		user.ReceiveMessage(fmt.Sprintf("Hola %s, bienvenido al chat de GDG Marbella!\n", user.GetName()))

		go func() {
			for {
				msg, err := user.SendMessage()
				if err != nil {
					log.Println(err)
					continue
				}

				for _, anyUser := range users.FetchAll() {
					if anyUser.RemoteAddr() != user.RemoteAddr() {
						user.ReceiveMessage(msg)
					}
				}
			}
		}()
	}
}
