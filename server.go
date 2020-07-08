package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	userslib "github.com/cartrujillosa/GoCourse/project/users"
)

func main() { // TODO: shutdown friendly

	chat := NewChat(":8888")
	defer chat.Close()

	registeredUsers := make(chan userslib.User)
	defer close(registeredUsers)

	for {
		conn, err := chat.Listener().Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err)
			continue
		}

		go registerUser(chat, &conn, registeredUsers)

		select {
		case user := <-registeredUsers:
			go func() {
				for {
					msg, err := user.GetMessage()
					if err != nil {
						log.Println(err)
						return
					}
					chat.SendMessage(user, msg)
				}
			}()
		}
	}

}

func registerUser(chat Chat, conn *net.Conn, registeredUsers chan userslib.User) {
	var response *string
	if response = ask(conn, "¿cómo te llamas?\n"); response == nil {
		return
	}
	name := *response
	if response = ask(conn, "¿dónde vives?\n"); response == nil {
		return
	}
	location := *response

	user, err := userslib.NewUser(name, location, conn)
	if err != nil {
		log.Println(err)
		return
	}

	if err := chat.RegisterUser(user); err != nil {
		log.Println(err)
		return
	}
	user.ReceiveMessage(fmt.Sprintf("Hola %s, bienvenido al chat de GDG Marbella!\n", user.Name()))
	chat.Broadcast(fmt.Sprintf("Den la bienvenida a %s que viene con fuerza desde %s\n", user.Name(), user.Location()))
	registeredUsers <- user
	return
}

func ask(conn *net.Conn, question string) *string {
	if _, err := io.WriteString(*conn, question); err != nil {
		log.Println(err)
		return nil
	}

	input, err := bufio.NewReader(*conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if strings.HasSuffix(input, "\n") {
		input = input[:len(input)-len("\n")]
	}
	// response := make(chan string)
	// defer close(response)
	// go func() {
	// 	for {
	// 		input, err := bufio.NewReader(*conn).ReadString('\n')
	// 		if err != nil {
	// 			log.Fatal(err)
	// 			return
	// 		}

	// 		response <- input
	// 	}
	// 	return
	// }()

	// select {
	// case resp := <-response:
	// 	return &resp
	// case <-time.After(4 * time.Second):
	// 	io.WriteString(*conn, "timed out\n")
	// 	return nil
	// }
	return &input
}
